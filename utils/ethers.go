package utils

import (
	"CryptoYes/server/utils/contract"
	"context"
	"github.com/astaxie/beego"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/onrik/ethrpc"
	"math/big"
)

/**
@title 获取用户签名nonce
@param user 用户地址
@param chainid 链id(BSC)
@param contractId 合约地址
@param address 账户地址
*/
func GetSignNonce(addr common.Address, chainid, contractId string) (int64, error) {
	rpc := beego.AppConfig.String("rpc::" + chainid)
	//服务器地址
	conn, err := ethclient.Dial(rpc)
	if err != nil {
		beego.Error("Dial err", err)
		return 0, err
	}
	defer conn.Close()

	//创建合约对象
	token, err := contract.NewBossRaid(common.HexToAddress(contractId), conn)
	if err != nil {
		beego.Error("new contract error", err)
		return 0, err
	}

	res, err := token.GetSpendNonce(&bind.CallOpts{
		Pending:     false,
		From:        common.Address{},
		BlockNumber: nil,
		Context:     nil,
	}, addr)
	if err != nil {
		beego.Error("balance error", err)
		return 0, err
	}
	return res.Int64(), nil
}

func GetReceiptTransaction(chainid, hash string) (*ethrpc.TransactionReceipt, error) {
	rpc := beego.AppConfig.String("rpc::" + chainid)
	client := ethrpc.New(rpc)
	receipt, err := client.EthGetTransactionReceipt(hash)
	return receipt, err
}

func FormatAddress(addr string) string {
	return common.HexToAddress(addr).Hex()
}

func ProfitWithdraw(chainId, contractId, privkey, amount string, sign [][]byte) (*types.Receipt, error) {
	rpc := beego.AppConfig.String("rpc::" + chainId)
	chainid, _ := beego.AppConfig.Int64("chainid::" + chainId)
	//服务器地址
	conn, err := ethclient.Dial(rpc)
	if err != nil {
		beego.Error("Dial err", err)
		return nil, err
	}
	defer conn.Close()

	//创建合约对象
	contract, err := contract.NewBossRaid(common.HexToAddress(contractId), conn)
	if err != nil {
		beego.Error("new contract error", err)
		return nil, err
	}
	privateKey, err := crypto.HexToECDSA(privkey[2:])
	if err != nil {
		beego.Error("new contract error", err)
		return nil, err
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, new(big.Int).SetInt64(chainid))
	if err != nil {
		beego.Error("new contract error", err)
		return nil, err
	}

	rAmount, _ := new(big.Int).SetString(amount, 10)
	tx, err := contract.Harvest(&bind.TransactOpts{
		From: auth.From,
		//Nonce:     nil,
		Signer: auth.Signer,
		//Value:     nil,
		//GasPrice:  nil,
		//GasFeeCap: nil,
		//GasTipCap: nil,
		//GasLimit:  0,
		//Context:   nil,
		//NoSend:    false,
	}, rAmount, sign)

	if err != nil {
		beego.Error(err)
		return nil, err
	}
	//等待挖矿完成
	receipt, err := bind.WaitMined(context.Background(), conn, tx)
	if err != nil {
		beego.Error("WaitMined error", err)
		return receipt, err
	}
	return receipt, nil
}

func BossRaid(chainId, contractId, privkey string) error {
	rpc := beego.AppConfig.String("rpc::" + chainId)
	chainid, _ := beego.AppConfig.Int64("chainid::" + chainId)
	//服务器地址
	conn, err := ethclient.Dial(rpc)
	if err != nil {
		beego.Error("Dial err", err)
		return err
	}
	defer conn.Close()

	//创建合约对象
	contract, err := contract.NewSmallBossRaid(common.HexToAddress(contractId), conn)
	if err != nil {
		beego.Error("new contract error", err)
		return err
	}
	privateKey, err := crypto.HexToECDSA(privkey[2:])
	if err != nil {
		beego.Error("new contract error", err)
		return err
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, new(big.Int).SetInt64(chainid))
	if err != nil {
		beego.Error("new contract error", err)
		return err
	}

	res, err := contract.GetBossRaid(
		&bind.CallOpts{
			Pending:     false,
			From:        common.Address{},
			BlockNumber: nil,
			Context:     nil,
		})

	if err != nil {
		beego.Error("balance error", err)
		return err
	}

	beego.Info("===res===", res)

	if !res {
		return nil
	}

	tx, err := contract.BeginRound(&bind.TransactOpts{
		From: auth.From,
		//Nonce:     nil,
		Signer: auth.Signer,
		//Value:     nil,
		//GasPrice:  nil,
		//GasFeeCap: nil,
		//GasTipCap: nil,
		//GasLimit:  0,
		//Context:   nil,
		//NoSend:    false,
	})

	if err != nil {
		beego.Error(err)
		return nil
	}
	//等待挖矿完成
	_, err = bind.WaitMined(context.Background(), conn, tx)
	if err != nil {
		beego.Error("WaitMined error", err)
		return err
	}

	return nil
}

func GetNFT(chainid, contractId string, tokenid int64) (*contract.CryptoZooNFTCryptoZoon, error) {
	rpc := beego.AppConfig.String("rpc::" + chainid)
	//服务器地址
	conn, err := ethclient.Dial(rpc)
	if err != nil {
		beego.Error("Dial err", err)
		return nil, err
	}
	defer conn.Close()

	//创建合约对象
	nft, err := contract.NewCryptoZooNFT(common.HexToAddress(contractId), conn)
	if err != nil {
		beego.Error("new contract error", err)
		return nil, err
	}

	res, err := nft.GetZooner(&bind.CallOpts{
		Pending:     false,
		From:        common.Address{},
		BlockNumber: nil,
		Context:     nil,
	}, new(big.Int).SetInt64(tokenid))
	if err != nil {
		beego.Error("balance error", err)
		return nil, err
	}

	beego.Info("res====", res)

	return &res, nil
}

func GetPhysical(chainid, contractId string, tokenid int64) (int64, int64, error) {
	rpc := beego.AppConfig.String("rpc::" + chainid)
	//服务器地址
	conn, err := ethclient.Dial(rpc)
	if err != nil {
		beego.Error("Dial err", err)
		return 0, 0, err
	}
	defer conn.Close()

	//创建合约对象
	nft, err := contract.NewCryptoZooNFT(common.HexToAddress(contractId), conn)
	if err != nil {
		beego.Error("new contract error", err)
		return 0, 0, err
	}

	res, err := nft.GetPhysicalAndMax(&bind.CallOpts{
		Pending:     false,
		From:        common.Address{},
		BlockNumber: nil,
		Context:     nil,
	}, new(big.Int).SetInt64(tokenid))
	if err != nil {
		beego.Error("balance error", err)
		return 0, 0, err
	}

	beego.Info("res====", res)

	return res.Physical.Int64(), res.Total.Int64(), nil
}
