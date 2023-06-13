package ord

import (
	"enum"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/astaxie/beego"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syncer/ord/parser"
	"syscall"
	"time"
	"utils"
	"utils/redis"
)

func init() {
	beego.LoadAppConfig("ini", "../conf/app.conf")
}

type result struct {
	inscriptionUid string
	inscriptionId  int64
	info           map[string]interface{}
	err            error
}

type uids []string

type Syncer struct {
	Concurrency           int64
	baseURL               string
	InscriptionIdStart    int64
	inscriptionUidChan    chan string
	resultChan            chan *result
	processChan           chan uids
	processFinishedChan   chan error
	eventChan             chan Event
	stopC                 chan struct{}
	lastInscriptionIdChan chan int64
}

var lastInscriptionIdFile = int64(0)

func NewSyncer() (*Syncer, error) {
	concurrency, _ := beego.AppConfig.Int64("worker::concurrency")
	baseURL := beego.AppConfig.String("ord::addr")
	inscription_id_start, _ := beego.AppConfig.Int64("ord::inscription_id_start")

	syncer := &Syncer{Concurrency: concurrency}

	syncer.InscriptionIdStart = inscription_id_start
	syncer.baseURL = baseURL
	syncer.inscriptionUidChan = make(chan string, concurrency)
	syncer.resultChan = make(chan *result, concurrency)
	syncer.processChan = make(chan uids)
	syncer.processFinishedChan = make(chan error)
	syncer.eventChan = make(chan Event, concurrency)
	syncer.stopC = make(chan struct{})
	syncer.lastInscriptionIdChan = make(chan int64)
	return syncer, nil
}

func (s *Syncer) Run() error {
	concurrency := s.Concurrency
	workers := make([]*Worker, concurrency)
	wg := &sync.WaitGroup{}
	wg.Add(int(concurrency))
	for i := 0; i < int(concurrency); i++ {
		workers[i] = &Worker{
			wid:        i,
			baseURL:    s.baseURL,
			uidChan:    s.inscriptionUidChan,
			resultChan: s.resultChan,
			stopC:      s.stopC,
		}

		go func(worker *Worker) {
			defer wg.Done()
			worker.Start()
		}(workers[i])
	}

	go func() {
		s.receveResult()
	}()

	go func() {
		for {
			lastInscriptionId, _ := s.getLastInscriptionId()
			s.lastInscriptionIdChan <- lastInscriptionId
			inscriptionURL, _ := url.JoinPath(s.baseURL, "inscriptions", fmt.Sprintf("%d", lastInscriptionId))
			beego.Info("start crawling from %s", inscriptionURL)
			_, err := s.parseInscriptions(inscriptionURL)
			if err != nil {
				beego.Error("failed to parse inscriptions: %v", err)
			}
			time.Sleep(60 * time.Second)
		}
	}()

	terminateSignals := make(chan os.Signal, 1)
	signal.Notify(terminateSignals, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM, os.Interrupt)
	signal := <-terminateSignals
	beego.Info("received signal %s, stopping workers...", signal)
	close(s.inscriptionUidChan)
	close(s.stopC)
	wg.Wait()
	beego.Info("all workers have been stopped")
	return nil
}

func (s *Syncer) receveResult() {
	insUids := make(uids, 0)
	results := make(map[string]*result)
	resultCount := 0
	var lastInscriptionId int64
	for {
		select {
		case lastInscriptionId = <-s.lastInscriptionIdChan:
			beego.Info("received lastInscriptionId: %d", lastInscriptionId)
		case result := <-s.resultChan:
			results[result.inscriptionUid] = result
			beego.Info("received result for inscription %d", result.inscriptionId)
			resultCount++
		case insUids = <-s.processChan:
			beego.Info("receiving %d inscriptions", len(insUids))
		case <-s.stopC:
			beego.Info("stopping result processor")
			return
		default:
			if resultCount == len(insUids) && len(insUids) > 0 {
				var err error
				processedResultCount := 0
				resultCount = 0
				resultsInOrder := make([]*result, len(insUids))
				for i := 0; i < len(insUids); i++ {
					resultsInOrder[i] = results[insUids[len(insUids)-i-1]]
					beego.Info("resultsInOrder: %v", resultsInOrder[i])
				}
				if len(resultsInOrder) != len(insUids) {
					err = fmt.Errorf("resultsInOrder length %d != insUids length %d", len(resultsInOrder), len(insUids))
				} else {
					// make sure to process results in ascending order of inscriptionId
					lastId := int64(0)
					for _, result := range resultsInOrder {
						if result.inscriptionId < 0 {
							continue
						}

						if result.inscriptionId < lastId {
							err = fmt.Errorf("results are not in order, lastId: %d, currentId: %d", lastId, result.inscriptionId)
							break
						}
						lastId = result.inscriptionId
					}
					if err == nil {
						processedResultCount, err = s.processResults(resultsInOrder, lastInscriptionId)
					}
				}
				if err != nil {
					beego.Error("failed to process results: %v", err)
					s.processFinishedChan <- err
				} else {
					beego.Info("processed %d results", processedResultCount)
					s.processFinishedChan <- nil
				}
			}
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func (s *Syncer) saveLastInscriptionId(lastInscriptionId int64) error {
	if lastInscriptionIdFile >= lastInscriptionId {
		return nil
	}

	_, err := redis.RedisSet(enum.LastInscriptionId, lastInscriptionId, -1)
	return err
}

func (s *Syncer) processResults(resultsInOrder []*result, lastInscriptionId int64) (int, error) {
	count := 0
	var lastSuccessInscriptionId int64
	for _, result := range resultsInOrder {
		if result.inscriptionId < 0 {
			continue
		}

		if result.inscriptionId < lastInscriptionId {
			beego.Info("inscription %d is less than lastInscriptionId %d, ignore", result.inscriptionId, lastInscriptionId)
			continue
		}
		if result.err != nil {
			beego.Error("failed to processResults: %v", result.err)
			return count, result.err
		}
		err := s.processResult(result)
		if err != nil {
			beego.Error("failed to processResult: %v", result.err)
			return count, err
		}
		beego.Info("processed inscription %d", result.inscriptionId)
		lastSuccessInscriptionId = result.inscriptionId
		count++
	}
	if lastSuccessInscriptionId > 0 && lastSuccessInscriptionId > lastInscriptionId {
		err := s.saveLastInscriptionId(lastSuccessInscriptionId)
		if err != nil {
			beego.Error("err:", err.Error())
		}
	}

	return count, nil
}

func (s *Syncer) processResult(result *result) error {
	inscriptionId := result.inscriptionId
	info := result.info
	switch info["content_parser"].(string) {
	case parser.NameDomain:
		err := s.processDomainMint(inscriptionId, info)
		if err != nil {
			return err
		}
	default:
	}
	return nil
}

func (s *Syncer) processDomainMint(inscriptionId int64, info map[string]interface{}) error {
	content := info["content"].(*parser.BRC721Mint)
	// check if the collection exists

	beego.Info("content:", content)
	return nil
}

//
//func (s *Syncer) processBRC721Deploy(inscriptionId int64, info map[string]interface{}) error {
//	o := info["content"].(*parser.BRC721Deploy)
//	// check if the collection already exists
//	collection, err := s.collectionUc.GetCollectionByTick(context.Background(), biz.ProtocolTypeBRC721, o.Tick)
//	if err != nil {
//		return err
//	}
//	if collection != nil {
//		if collection.InscriptionID > inscriptionId {
//			s.logger.Warnf("collection %s already exists, but inscriptionId %d is less than %d, ignore inscription %d", o.Tick, collection.InscriptionID, inscriptionId, inscriptionId)
//		} else {
//			s.logger.Infof("collection %s already exists, ignore inscription %d", o.Tick, inscriptionId)
//		}
//		return nil
//	}
//	// create the collection
//	collection = &biz.Collection{
//		P:      biz.ProtocolTypeBRC721,
//		Tick:   o.Tick,
//		Supply: 0,
//	}
//	max, err := strconv.ParseUint(o.Max, 10, 64)
//	if err != nil {
//		return err
//	}
//	collection.Max = max
//	if o.BaseURI != nil {
//		collection.BaseURI = *o.BaseURI
//	}
//	if o.Meta != nil {
//		collection.Name = o.Meta.Name
//		collection.Description = o.Meta.Description
//		collection.Image = o.Meta.Image
//		collection.Attributes = o.Meta.Attributes
//	}
//	collection.TxHash = info["genesis_transaction"].(string)
//	collection.BlockHeight = info["genesis_height"].(uint64)
//	collection.BlockTime = info["timestamp"].(time.Time)
//	collection.Address = info["address"].(string)
//	collection.InscriptionID = inscriptionId
//	collection.InscriptionUID = info["uid"].(string)
//	collection, err = s.collectionUc.CreateCollection(context.Background(), collection)
//	if err != nil {
//		return err
//	}
//	s.logger.Infof("created collection %s for inscription %d", o.Tick, inscriptionId)
//	return nil
//}
//
//func (s *Syncer) processBRC721Mint(inscriptionId int64, info map[string]interface{}) error {
//	o := info["content"].(*parser.BRC721Mint)
//	// check if the collection exists
//	collection, err := s.collectionUc.GetCollectionByTick(context.Background(), biz.ProtocolTypeBRC721, o.Tick)
//	if err != nil {
//		return err
//	}
//	if collection == nil {
//		s.logger.Infof("collection %s not found, ignore inscription %d", o.Tick, inscriptionId)
//		return nil
//	}
//	if collection.InscriptionID >= inscriptionId {
//		s.logger.Warnf("collection %s inscriptionId %d is greater than %d, ignore inscription %d", o.Tick, collection.InscriptionID, inscriptionId, inscriptionId)
//		return nil
//	}
//	s.logger.Debugf("collection: %+v", collection)
//	// check if supply is full
//	if collection.Supply >= collection.Max {
//		s.logger.Infof("collection %s supply is full, ignore inscription %d", o.Tick, inscriptionId)
//		return nil
//	}
//	t, err := s.tokenUc.FindByInscriptionID(context.Background(), inscriptionId)
//	if err != nil {
//		return err
//	}
//	if len(t) > 0 {
//		s.logger.Infof("token with inscription %d already processed, ignore", inscriptionId)
//		return nil
//	}
//	// create token
//	token := &biz.Token{
//		Tick:           o.Tick,
//		P:              biz.ProtocolTypeBRC721,
//		TokenID:        collection.Supply + 1,
//		TxHash:         info["genesis_transaction"].(string),
//		BlockHeight:    info["genesis_height"].(uint64),
//		BlockTime:      info["timestamp"].(time.Time),
//		Address:        info["address"].(string),
//		InscriptionID:  inscriptionId,
//		InscriptionUID: info["uid"].(string),
//		CollectionID:   collection.ID,
//	}
//	token, err = s.tokenUc.CreateToken(context.Background(), token)
//	if err != nil {
//		s.logger.Errorf("failed to create token: %T: %v", err, err)
//		return err
//	}
//	s.logger.Infof("created token %d for inscription %d", token.TokenID, inscriptionId)
//
//	collection.Supply++
//	collection, err = s.collectionUc.UpdateCollection(context.Background(), collection)
//	if err != nil {
//		return err
//	}
//	s.logger.Infof("updated collection %s supply to %d", o.Tick, collection.Supply)
//	return nil
//}
//
//func (s *Syncer) processBRC721Update(inscriptionId int64, info map[string]interface{}) error {
//	o := info["content"].(*parser.BRC721Update)
//	// check if the collection exists
//	collection, err := s.collectionUc.GetCollectionByTick(context.Background(), biz.ProtocolTypeBRC721, o.Tick)
//	if err != nil {
//		return err
//	}
//	if collection == nil {
//		s.logger.Infof("collection %s not found, ignore inscription %d", o.Tick, inscriptionId)
//		return nil
//	}
//	// update collection
//	if o.BaseURI != nil {
//		collection.BaseURI = *o.BaseURI
//	}
//	_, err = s.collectionUc.UpdateCollection(context.Background(), collection)
//	if err != nil {
//		return err
//	}
//	s.logger.Infof("updated collection %s", o.Tick)
//	return nil
//}

func (s *Syncer) getLastInscriptionId() (int64, error) {
	lastInscriptionId, err := s.readLastInscriptionId()
	if err == nil {
		beego.Info("get lastInscriptionId from file: %d", lastInscriptionId)
		return lastInscriptionId, nil
	}

	lastInscriptionId = s.InscriptionIdStart
	beego.Info("get lastInscriptionId from config: %d", lastInscriptionId)
	return lastInscriptionId, nil
}

func (s *Syncer) readLastInscriptionId() (int64, error) {
	id, _ := redis.RedisGet(enum.LastInscriptionId).Int64()

	return id, nil
}

func (s *Syncer) parseInscriptions(inscriptionURL string) (string, error) {
	if inscriptionURL == "" {
		return "", nil
	}
	beego.Info("fetching %s", inscriptionURL)
	resp, err := utils.HttpGetResp(inscriptionURL)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", err
	}

	var insUids uids
	links := doc.Find("div.thumbnails a")
	links.Each(func(i int, ss *goquery.Selection) {
		href, _ := ss.Attr("href")
		uid := strings.Replace(href, "/inscription/", "", -1)
		if uid == "" {
			return
		}

		beego.Info("uid %s", uid)
		// record the inscription uids, so that we can process them in order
		insUids = append(insUids, uid)
	})

	s.processChan <- insUids
	for _, insUid := range insUids {
		s.inscriptionUidChan <- insUid
	}
	// wait for the process to finish
	err = <-s.processFinishedChan
	if err != nil {
		return "", err
	}

	prevLink := doc.Find("a.next")
	if prevLink.Length() > 0 {
		href, _ := prevLink.Attr("href")
		inscriptionURL, _ = url.JoinPath(s.baseURL, href)
		return s.parseInscriptions(inscriptionURL)
	}

	return "", nil
}
