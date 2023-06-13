package ord

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/astaxie/beego"
	"io/ioutil"
	"net/url"
	"strconv"
	"strings"
	"syncer/ord/parser"
	"time"
	"utils"
)

type Worker struct {
	wid        int
	baseURL    string
	uidChan    chan string
	resultChan chan (*result)
	stopC      chan struct{}
}

func (w *Worker) Start() {
	for {
		select {
		case uid := <-w.uidChan:
			beego.Debug("[worker]: processing inscription:", w.wid, uid)
			w.resultChan <- w.processInscription(uid)
		case <-w.stopC:
			beego.Debug("[worker]: stopping", w.wid)
			return
		}
	}
}

func (w *Worker) processInscription(uid string) *result {
	info, err := w.parseInscriptionInfo(uid)
	if info == nil {
		info = make(map[string]interface{})
	}
	info["uid"] = uid
	// FIXME: inscription_id is not always available
	inscriptionID, ok := info["inscription_id"].(int64)
	if !ok {
		if err == nil {
			beego.Error("info[inscription_id]:", info["inscription_id"])
			beego.Error("============info=============:", info)
			err = fmt.Errorf("failed to get inscription_id")
		}
		return &result{inscriptionUid: uid, inscriptionId: 0, info: info, err: err}
	}

	beego.Debug("[worker ] parsed inscription:", w.wid, inscriptionID)
	return &result{inscriptionUid: uid, inscriptionId: inscriptionID, info: info, err: err}
}

func (w *Worker) parseInscriptionInfo(uid string) (map[string]interface{}, error) {
	inscriptionURL, _ := url.JoinPath(w.baseURL, "inscription", uid)
	beego.Info("[worker ] fetching ...", w.wid, inscriptionURL)
	resp, err := utils.HttpGetResp(inscriptionURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	details := make(map[string]interface{})
	inscriptionIDText := doc.Find("h1").First().Text()
	beego.Info("inscriptionIDText:", inscriptionIDText)
	if strings.Contains(inscriptionIDText, "unstable") {
		details["inscription_id"] = int64(-1)
	} else {
		inscriptionIDText = strings.Replace(inscriptionIDText, "Inscription ", "", -1)
		// convert inscriptionID string to int64
		inscriptionID, err := strconv.ParseInt(inscriptionIDText, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to convert inscriptionID %s to int64: %v", inscriptionIDText, err)
		}
		details["inscription_id"] = inscriptionID
	}

	dtElements := doc.Find("dl dt")
	ddElements := doc.Find("dl dd")
	dtElements.Each(func(i int, dt *goquery.Selection) {
		key := dt.Text()
		dd := ddElements.Eq(i)
		value := dd.Text()
		if aTag := dd.Find("a"); aTag.Length() > 0 {
			value = aTag.Text()
		}
		key = strings.Replace(strings.ToLower(key), " ", "_", -1)
		switch key {
		case "id":
			details[key] = value
		case "output_value":
			v, _ := strconv.ParseUint(value, 10, 64)
			details[key] = v
		case "content_length":
			// conver "3440 bytes" to 3440
			value = strings.Replace(value, " bytes", "", -1)
			v, _ := strconv.ParseUint(value, 10, 64)
			details[key] = v
		case "timestamp":
			// convert "2023-05-28 03:28:17 UTC" to time.Time
			v, _ := time.Parse("2006-01-02 15:04:05 UTC", value)
			details[key] = v
		case "genesis_height":
			v, _ := strconv.ParseUint(value, 10, 64)
			details[key] = v
		case "genesis_fee":
			v, _ := strconv.ParseUint(value, 10, 64)
			details[key] = v
		case "offset":
			v, _ := strconv.ParseUint(value, 10, 64)
			details[key] = v
		case "content_type":
			details[key] = value
		default:
			details[key] = value
		}
	})

	err = w.parseContent(details)
	if err != nil {
		return nil, err
	}
	return details, nil
}

func (w *Worker) parseContent(info map[string]interface{}) error {
	contentURL, _ := url.JoinPath(w.baseURL, "content", info["id"].(string))
	beego.Info("[worker %d] fetching %s...", w.wid, contentURL)
	resp, err := utils.HttpGetResp(contentURL)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	content_type := info["content_type"].(string)

	if validateContentType(content_type) {
		if json.Valid(body) {
			domainParser := parser.NameDomainParser{}
			data, valid, err := domainParser.Parse(body)
			if err != nil || !valid {
				beego.Error("valid:", valid)
				beego.Info("body:", string(body))

				return nil
			}

			info["content"] = data
			info["content_parser"] = domainParser.Name()
		} else {
			count := strings.Count(string(body), ".")
			if count != 1 {
				return nil
			}

			names := strings.Split(string(body), ".")
			if len(names) != 2 {
				return nil
			}

			if len(names[1]) < 1 {
				return nil
			}

			content_length, ok := info["content_length"].(uint64)
			if ok {
				if content_length > 1024 {
					return nil
				}
			} else {
				if len(body) > 1024 {
					return nil
				}
			}

			info["content"] = string(body)
			info["content_parser"] = parser.NameDomain
		}
	}

	//var found bool
	//for _, p := range parser.ParserList() {
	//	data, valid, err := p.Parse(body)
	//	if err != nil {
	//		continue
	//	}
	//	if !valid {
	//		continue
	//	}
	//	found = true
	//	info["content"] = data
	//	info["content_parser"] = p.Name()
	//	break
	//}
	//if !found {
	//	info["content"] = body
	//	info["content_parser"] = "raw"
	//}
	return nil
}

func validateContentType(data string) bool {
	if strings.Contains(data, "text/plain") {
		return true
	}

	if strings.Contains(data, "application/json") {
		return true
	}

	return false
}
