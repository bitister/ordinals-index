package test

import (
	utils "CryptoYes/server/utils"
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/astaxie/beego"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
)

func init() {

}

func TestSign(t *testing.T) {

	types := []string{"address", "uint256"}
	values := []interface{}{"0x594fBfa5D8537382fdEBCCc7055Deab6416CEf59", fmt.Sprint(97)}
	initdata := solsha3.SoliditySHA3(types, values)
	beego.Info("Pack --- ", hexutil.Encode(initdata))
	// types = []string{"string", "bytes32"}
	// values = []interface{}{"\x19Ethereum Signed Message:\n32", initdata}
	// initdata = solsha3.SoliditySHA3(types, values)
	// beego.Info("Pack --- ", hexutil.Encode(initdata))
	// // // address := "0x594fBfa5D8537382fdEBCCc7055Deab6416CEf59"
	privkey := "0xb8faa392ff0cb04a9a623bcc521c4e2184882889aa69899f9bd12f7d1b184cea"
	// initdata := "077CBC0D37C94312225A11F8ED284DFC"
	// privkey := "0x9e6e1363895fe1531cfbef7710780780cc3dd2795f4312a588f02d899e5fa935"
	signature, _ := utils.SignClaim4Evm(common.BytesToHash(initdata), privkey)
	// signature, _ := utils.Sign(hexutil.Encode(initdata), privkey)
	beego.Info("signature ----", signature)

	V, R, S := utils.SigRSV(signature)

	beego.Info("V ----", V)
	beego.Info("R ----", hexutil.Encode(R[:]))
	beego.Info("S ----", hexutil.Encode(S[:]))

	// signature := "0x06753b7ec8c57d2c8803705b7b5c0ae986cb78e013d275ea242bba7d7ce8b3630450728462a9784ee91c54ecb0063ff9830ff973a95a95adfe503acb3e60f3ce1c"
	// s, err := utils.VerifyMessage(hexutil.Encode(initdata), signature)
	// beego.Info(s, err)
}

//
//func TestUlabContract(t *testing.T) {
//
//	//服务器地址
//	conn, err := ethclient.Dial("wss://bsc-testnet.nodereal.io/ws/v1/c92486b30c634586b6864cd4f4361440")
//	if err != nil {
//		beego.Error("Dial err", err)
//	}
//	defer conn.Close()
//
//	//创建合约对象
//	ulab, err := contract.NewUlab(common.HexToAddress("0x1498Ef9Fd9255B6967B510777F52B17E7d46666C"), conn)
//	if err != nil {
//		beego.Error("new contract error", err)
//	}
//	var start uint64
//	start = 24660900
//	loans := make(chan *contract.UlabLoan, 50000)
//
//	go func() {
//		for {
//			loan := <-loans
//			beego.Info("===", loan)
//		}
//	}()
//
//	opts := bind.WatchOpts{Start: &start, Context: context.Background()}
//	_, err = ulab.WatchLoan(&opts, loans)
//	if err != nil {
//		beego.Error("WatchLoan contract error", err)
//	}
//
//	select {}
//
//}
//
//func TestTime(t *testing.T) {
//	date := utils.GetStime()
//	beego.Warning(date)
//	beego.Info(0.0 + 0.1)
//	beego.Info(0.1 + 0.1)
//	beego.Info(0.2 + 0.1)
//}

func TestCPUID(t *testing.T) {
	beego.Info(utils.GET_CPUID())
}

func TestEncrypt(t *testing.T) {
	str := "ID:54060500FFFBEBBF"
	//解押文件
	//读取文件全部内容
	b, err := ioutil.ReadFile("E:/gowork/src/ulab/key/keyadmin")
	// b, err := ioutil.ReadFile(keypath)
	if err != nil {
		beego.Error("ioutil ReadFile error: ", err)
	}
	//解密
	h := sha256.New()
	h.Write([]byte(str))
	hash := fmt.Sprintf("%x", h.Sum(nil))
	//content := DecryptWithAES("123456", zip.Unzip([]byte(b)))
	content := utils.DecryptWithAES("gpJ1SXT]f~nU^c793Al[L-eD%WWAY2", utils.Unzip([]byte(b)))
	res := utils.DecryptWithAES("gpJ1SXT]f~nU^c793Al[L-eD%WWAY2"+hash, content)
	beego.Info("====", res)
}

func TestEncode(t *testing.T) {
	//str := "ID:54060500FFFBEBBF"
	//解押文件
	//读取文件全部内容
	//b, err := ioutil.ReadFile("E:/gowork/src/ulab/key/keyadmin")
	// b, err := ioutil.ReadFile(keypath)

	//if err != nil {
	//	beego.Error("ioutil ReadFile error: ", err)
	//}
	b := []byte("U2FsdGVkX182L2dEx4KKRUubtlBVsnvSyQsYwCuPI9JqpTCl+r/wzg6Tm6MCLyaWnP874ud+qO8enmnalSNL4LCnvdNFuSm/27KxDhsixz4h5L3UBQWdNlazjmSQfaQx")

	//解密
	//h := sha256.New()
	//h.Write([]byte(str))
	//hash := fmt.Sprintf("%x", h.Sum(nil))
	//content := DecryptWithAES("123456", zip.Unzip([]byte(b)))
	//content := utils.EncryptWithAES("gpJ1SXT]f~nU^c793Al[L-eD%WWAY2", string(utils.Zip("0xab834867Fe0b2FCb04fEAA1C8F1f3e8BB31844D0")))

	content := utils.DecryptWithAES("gpJ1SXT]f~nU^c793Al[L-eD%WWAY2", utils.Unzip([]byte(b)))
	//res := utils.DecryptWithAES("gpJ1SXT]f~nU^c793Al[L-eD%WWAY2"+hash, content)
	beego.Info("====", content)
}

func TestSwapQuotation(t *testing.T) {
	tokens := `\"0xbb4cdb9cbd36b01bd1cbaebf2de08d9173bc095c\",\"0x55d398326f99059ff775485246999027b3197955\",\"0x3c45a24d36ab6fc1925533c1f57bc7e1b6fba8a4\"`
	data := `{"query":"\n      query tokens {\n        now: tokens(\n      where: {id_in: [` + tokens + `]}\n      \n      orderBy: tradeVolumeUSD\n      orderDirection: desc\n    ) {\n      id\n     derivedUSD\n      }\n  \n        oneDayAgo: tokens(\n      where: {id_in: [` + tokens + `]}\n     \n   block: {number: 23645992}\n     orderBy: tradeVolumeUSD\n      orderDirection: desc\n    ) {\n      id\n     derivedUSD\n   }\n  \n                     }\n    ","operationName":"tokens"}`
	s, err := utils.HttpPostRaw("https://proxy-worker-dev.pancake-swap.workers.dev/bsc-exchange", data)
	beego.Info(s, err)
}
