// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// BossRaidBossInfo is an auto generated low-level Go binding around an user-defined struct.
type BossRaidBossInfo struct {
	StartTime   *big.Int
	EndTime     *big.Int
	BloodVolume *big.Int
	TotalDamage *big.Int
	TotalValue  *big.Int
	Status      uint8
}

// BossRaidMetaData contains all meta data concerning the BossRaid contract.
var BossRaidMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"level\",\"type\":\"uint256\"}],\"name\":\"BattleBossEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"level\",\"type\":\"uint256\"}],\"name\":\"BattleEndEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"level\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalDamage\",\"type\":\"uint256\"}],\"name\":\"BattleRaidEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"level\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bloodVolume\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"BossBeginEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"HarvestEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BossId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"BossInfos\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bloodVolume\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalDamage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalValue\",\"type\":\"uint256\"},{\"internalType\":\"enumBossRaid.bossstatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"actives\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"battleBoss\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"beginBattle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"beginFirst\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"beginRound\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBossInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bloodVolume\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalDamage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalValue\",\"type\":\"uint256\"},{\"internalType\":\"enumBossRaid.bossstatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"internalType\":\"structBossRaid.BossInfo\",\"name\":\"_bossInfo\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBossRaid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_status\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"getInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bloodVolume\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalDamage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalValue\",\"type\":\"uint256\"},{\"internalType\":\"enumBossRaid.bossstatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"internalType\":\"structBossRaid.BossInfo\",\"name\":\"_bossInfo\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_level\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_userHarm\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nextTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getSpendNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"_sig\",\"type\":\"bytes[]\"}],\"name\":\"harvest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_owners\",\"type\":\"address[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"level\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_nftAddr\",\"type\":\"address\"}],\"name\":\"setAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"userRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"_sig\",\"type\":\"bytes[]\"}],\"name\":\"validSignature\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"zoonerNFT\",\"outputs\":[{\"internalType\":\"contractCryptoZooNFT\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// BossRaidABI is the input ABI used to generate the binding from.
// Deprecated: Use BossRaidMetaData.ABI instead.
var BossRaidABI = BossRaidMetaData.ABI

// BossRaid is an auto generated Go binding around an Ethereum contract.
type BossRaid struct {
	BossRaidCaller     // Read-only binding to the contract
	BossRaidTransactor // Write-only binding to the contract
	BossRaidFilterer   // Log filterer for contract events
}

// BossRaidCaller is an auto generated read-only Go binding around an Ethereum contract.
type BossRaidCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BossRaidTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BossRaidTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BossRaidFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BossRaidFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BossRaidSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BossRaidSession struct {
	Contract     *BossRaid         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BossRaidCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BossRaidCallerSession struct {
	Contract *BossRaidCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// BossRaidTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BossRaidTransactorSession struct {
	Contract     *BossRaidTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BossRaidRaw is an auto generated low-level Go binding around an Ethereum contract.
type BossRaidRaw struct {
	Contract *BossRaid // Generic contract binding to access the raw methods on
}

// BossRaidCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BossRaidCallerRaw struct {
	Contract *BossRaidCaller // Generic read-only contract binding to access the raw methods on
}

// BossRaidTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BossRaidTransactorRaw struct {
	Contract *BossRaidTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBossRaid creates a new instance of BossRaid, bound to a specific deployed contract.
func NewBossRaid(address common.Address, backend bind.ContractBackend) (*BossRaid, error) {
	contract, err := bindBossRaid(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BossRaid{BossRaidCaller: BossRaidCaller{contract: contract}, BossRaidTransactor: BossRaidTransactor{contract: contract}, BossRaidFilterer: BossRaidFilterer{contract: contract}}, nil
}

// NewBossRaidCaller creates a new read-only instance of BossRaid, bound to a specific deployed contract.
func NewBossRaidCaller(address common.Address, caller bind.ContractCaller) (*BossRaidCaller, error) {
	contract, err := bindBossRaid(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BossRaidCaller{contract: contract}, nil
}

// NewBossRaidTransactor creates a new write-only instance of BossRaid, bound to a specific deployed contract.
func NewBossRaidTransactor(address common.Address, transactor bind.ContractTransactor) (*BossRaidTransactor, error) {
	contract, err := bindBossRaid(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BossRaidTransactor{contract: contract}, nil
}

// NewBossRaidFilterer creates a new log filterer instance of BossRaid, bound to a specific deployed contract.
func NewBossRaidFilterer(address common.Address, filterer bind.ContractFilterer) (*BossRaidFilterer, error) {
	contract, err := bindBossRaid(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BossRaidFilterer{contract: contract}, nil
}

// bindBossRaid binds a generic wrapper to an already deployed contract.
func bindBossRaid(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BossRaidMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BossRaid *BossRaidRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BossRaid.Contract.BossRaidCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BossRaid *BossRaidRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BossRaid.Contract.BossRaidTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BossRaid *BossRaidRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BossRaid.Contract.BossRaidTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BossRaid *BossRaidCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BossRaid.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BossRaid *BossRaidTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BossRaid.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BossRaid *BossRaidTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BossRaid.Contract.contract.Transact(opts, method, params...)
}

// BossId is a free data retrieval call binding the contract method 0x7bd36711.
//
// Solidity: function BossId() view returns(uint256)
func (_BossRaid *BossRaidCaller) BossId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BossRaid.contract.Call(opts, &out, "BossId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BossId is a free data retrieval call binding the contract method 0x7bd36711.
//
// Solidity: function BossId() view returns(uint256)
func (_BossRaid *BossRaidSession) BossId() (*big.Int, error) {
	return _BossRaid.Contract.BossId(&_BossRaid.CallOpts)
}

// BossId is a free data retrieval call binding the contract method 0x7bd36711.
//
// Solidity: function BossId() view returns(uint256)
func (_BossRaid *BossRaidCallerSession) BossId() (*big.Int, error) {
	return _BossRaid.Contract.BossId(&_BossRaid.CallOpts)
}

// BossInfos is a free data retrieval call binding the contract method 0x13254c66.
//
// Solidity: function BossInfos(uint256 ) view returns(uint256 startTime, uint256 endTime, uint256 bloodVolume, uint256 totalDamage, uint256 totalValue, uint8 status)
func (_BossRaid *BossRaidCaller) BossInfos(opts *bind.CallOpts, arg0 *big.Int) (struct {
	StartTime   *big.Int
	EndTime     *big.Int
	BloodVolume *big.Int
	TotalDamage *big.Int
	TotalValue  *big.Int
	Status      uint8
}, error) {
	var out []interface{}
	err := _BossRaid.contract.Call(opts, &out, "BossInfos", arg0)

	outstruct := new(struct {
		StartTime   *big.Int
		EndTime     *big.Int
		BloodVolume *big.Int
		TotalDamage *big.Int
		TotalValue  *big.Int
		Status      uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StartTime = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.EndTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.BloodVolume = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.TotalDamage = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.TotalValue = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Status = *abi.ConvertType(out[5], new(uint8)).(*uint8)

	return *outstruct, err

}

// BossInfos is a free data retrieval call binding the contract method 0x13254c66.
//
// Solidity: function BossInfos(uint256 ) view returns(uint256 startTime, uint256 endTime, uint256 bloodVolume, uint256 totalDamage, uint256 totalValue, uint8 status)
func (_BossRaid *BossRaidSession) BossInfos(arg0 *big.Int) (struct {
	StartTime   *big.Int
	EndTime     *big.Int
	BloodVolume *big.Int
	TotalDamage *big.Int
	TotalValue  *big.Int
	Status      uint8
}, error) {
	return _BossRaid.Contract.BossInfos(&_BossRaid.CallOpts, arg0)
}

// BossInfos is a free data retrieval call binding the contract method 0x13254c66.
//
// Solidity: function BossInfos(uint256 ) view returns(uint256 startTime, uint256 endTime, uint256 bloodVolume, uint256 totalDamage, uint256 totalValue, uint8 status)
func (_BossRaid *BossRaidCallerSession) BossInfos(arg0 *big.Int) (struct {
	StartTime   *big.Int
	EndTime     *big.Int
	BloodVolume *big.Int
	TotalDamage *big.Int
	TotalValue  *big.Int
	Status      uint8
}, error) {
	return _BossRaid.Contract.BossInfos(&_BossRaid.CallOpts, arg0)
}

// Actives is a free data retrieval call binding the contract method 0x947535b4.
//
// Solidity: function actives(uint256 ) view returns(bool)
func (_BossRaid *BossRaidCaller) Actives(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _BossRaid.contract.Call(opts, &out, "actives", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Actives is a free data retrieval call binding the contract method 0x947535b4.
//
// Solidity: function actives(uint256 ) view returns(bool)
func (_BossRaid *BossRaidSession) Actives(arg0 *big.Int) (bool, error) {
	return _BossRaid.Contract.Actives(&_BossRaid.CallOpts, arg0)
}

// Actives is a free data retrieval call binding the contract method 0x947535b4.
//
// Solidity: function actives(uint256 ) view returns(bool)
func (_BossRaid *BossRaidCallerSession) Actives(arg0 *big.Int) (bool, error) {
	return _BossRaid.Contract.Actives(&_BossRaid.CallOpts, arg0)
}

// GetBossInfo is a free data retrieval call binding the contract method 0xff14c6b4.
//
// Solidity: function getBossInfo() view returns((uint256,uint256,uint256,uint256,uint256,uint8) _bossInfo)
func (_BossRaid *BossRaidCaller) GetBossInfo(opts *bind.CallOpts) (BossRaidBossInfo, error) {
	var out []interface{}
	err := _BossRaid.contract.Call(opts, &out, "getBossInfo")

	if err != nil {
		return *new(BossRaidBossInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(BossRaidBossInfo)).(*BossRaidBossInfo)

	return out0, err

}

// GetBossInfo is a free data retrieval call binding the contract method 0xff14c6b4.
//
// Solidity: function getBossInfo() view returns((uint256,uint256,uint256,uint256,uint256,uint8) _bossInfo)
func (_BossRaid *BossRaidSession) GetBossInfo() (BossRaidBossInfo, error) {
	return _BossRaid.Contract.GetBossInfo(&_BossRaid.CallOpts)
}

// GetBossInfo is a free data retrieval call binding the contract method 0xff14c6b4.
//
// Solidity: function getBossInfo() view returns((uint256,uint256,uint256,uint256,uint256,uint8) _bossInfo)
func (_BossRaid *BossRaidCallerSession) GetBossInfo() (BossRaidBossInfo, error) {
	return _BossRaid.Contract.GetBossInfo(&_BossRaid.CallOpts)
}

// GetBossRaid is a free data retrieval call binding the contract method 0x2bf89616.
//
// Solidity: function getBossRaid() view returns(bool _status)
func (_BossRaid *BossRaidCaller) GetBossRaid(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BossRaid.contract.Call(opts, &out, "getBossRaid")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetBossRaid is a free data retrieval call binding the contract method 0x2bf89616.
//
// Solidity: function getBossRaid() view returns(bool _status)
func (_BossRaid *BossRaidSession) GetBossRaid() (bool, error) {
	return _BossRaid.Contract.GetBossRaid(&_BossRaid.CallOpts)
}

// GetBossRaid is a free data retrieval call binding the contract method 0x2bf89616.
//
// Solidity: function getBossRaid() view returns(bool _status)
func (_BossRaid *BossRaidCallerSession) GetBossRaid() (bool, error) {
	return _BossRaid.Contract.GetBossRaid(&_BossRaid.CallOpts)
}

// GetInfo is a free data retrieval call binding the contract method 0xffdd5cf1.
//
// Solidity: function getInfo(address _sender) view returns((uint256,uint256,uint256,uint256,uint256,uint8) _bossInfo, uint256 _level, uint256 _userHarm, uint256 _nextTime)
func (_BossRaid *BossRaidCaller) GetInfo(opts *bind.CallOpts, _sender common.Address) (struct {
	BossInfo BossRaidBossInfo
	Level    *big.Int
	UserHarm *big.Int
	NextTime *big.Int
}, error) {
	var out []interface{}
	err := _BossRaid.contract.Call(opts, &out, "getInfo", _sender)

	outstruct := new(struct {
		BossInfo BossRaidBossInfo
		Level    *big.Int
		UserHarm *big.Int
		NextTime *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BossInfo = *abi.ConvertType(out[0], new(BossRaidBossInfo)).(*BossRaidBossInfo)
	outstruct.Level = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.UserHarm = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.NextTime = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetInfo is a free data retrieval call binding the contract method 0xffdd5cf1.
//
// Solidity: function getInfo(address _sender) view returns((uint256,uint256,uint256,uint256,uint256,uint8) _bossInfo, uint256 _level, uint256 _userHarm, uint256 _nextTime)
func (_BossRaid *BossRaidSession) GetInfo(_sender common.Address) (struct {
	BossInfo BossRaidBossInfo
	Level    *big.Int
	UserHarm *big.Int
	NextTime *big.Int
}, error) {
	return _BossRaid.Contract.GetInfo(&_BossRaid.CallOpts, _sender)
}

// GetInfo is a free data retrieval call binding the contract method 0xffdd5cf1.
//
// Solidity: function getInfo(address _sender) view returns((uint256,uint256,uint256,uint256,uint256,uint8) _bossInfo, uint256 _level, uint256 _userHarm, uint256 _nextTime)
func (_BossRaid *BossRaidCallerSession) GetInfo(_sender common.Address) (struct {
	BossInfo BossRaidBossInfo
	Level    *big.Int
	UserHarm *big.Int
	NextTime *big.Int
}, error) {
	return _BossRaid.Contract.GetInfo(&_BossRaid.CallOpts, _sender)
}

// GetSpendNonce is a free data retrieval call binding the contract method 0x241602f9.
//
// Solidity: function getSpendNonce(address _user) view returns(uint256)
func (_BossRaid *BossRaidCaller) GetSpendNonce(opts *bind.CallOpts, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BossRaid.contract.Call(opts, &out, "getSpendNonce", _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSpendNonce is a free data retrieval call binding the contract method 0x241602f9.
//
// Solidity: function getSpendNonce(address _user) view returns(uint256)
func (_BossRaid *BossRaidSession) GetSpendNonce(_user common.Address) (*big.Int, error) {
	return _BossRaid.Contract.GetSpendNonce(&_BossRaid.CallOpts, _user)
}

// GetSpendNonce is a free data retrieval call binding the contract method 0x241602f9.
//
// Solidity: function getSpendNonce(address _user) view returns(uint256)
func (_BossRaid *BossRaidCallerSession) GetSpendNonce(_user common.Address) (*big.Int, error) {
	return _BossRaid.Contract.GetSpendNonce(&_BossRaid.CallOpts, _user)
}

// Level is a free data retrieval call binding the contract method 0x6fd5ae15.
//
// Solidity: function level() view returns(uint256)
func (_BossRaid *BossRaidCaller) Level(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BossRaid.contract.Call(opts, &out, "level")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Level is a free data retrieval call binding the contract method 0x6fd5ae15.
//
// Solidity: function level() view returns(uint256)
func (_BossRaid *BossRaidSession) Level() (*big.Int, error) {
	return _BossRaid.Contract.Level(&_BossRaid.CallOpts)
}

// Level is a free data retrieval call binding the contract method 0x6fd5ae15.
//
// Solidity: function level() view returns(uint256)
func (_BossRaid *BossRaidCallerSession) Level() (*big.Int, error) {
	return _BossRaid.Contract.Level(&_BossRaid.CallOpts)
}

// UserRewards is a free data retrieval call binding the contract method 0x63c2a20a.
//
// Solidity: function userRewards(address , uint256 ) view returns(uint256)
func (_BossRaid *BossRaidCaller) UserRewards(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BossRaid.contract.Call(opts, &out, "userRewards", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserRewards is a free data retrieval call binding the contract method 0x63c2a20a.
//
// Solidity: function userRewards(address , uint256 ) view returns(uint256)
func (_BossRaid *BossRaidSession) UserRewards(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _BossRaid.Contract.UserRewards(&_BossRaid.CallOpts, arg0, arg1)
}

// UserRewards is a free data retrieval call binding the contract method 0x63c2a20a.
//
// Solidity: function userRewards(address , uint256 ) view returns(uint256)
func (_BossRaid *BossRaidCallerSession) UserRewards(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _BossRaid.Contract.UserRewards(&_BossRaid.CallOpts, arg0, arg1)
}

// ValidSignature is a free data retrieval call binding the contract method 0x178c47f5.
//
// Solidity: function validSignature(address _sender, uint256 _amount, bytes[] _sig) view returns(bool)
func (_BossRaid *BossRaidCaller) ValidSignature(opts *bind.CallOpts, _sender common.Address, _amount *big.Int, _sig [][]byte) (bool, error) {
	var out []interface{}
	err := _BossRaid.contract.Call(opts, &out, "validSignature", _sender, _amount, _sig)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidSignature is a free data retrieval call binding the contract method 0x178c47f5.
//
// Solidity: function validSignature(address _sender, uint256 _amount, bytes[] _sig) view returns(bool)
func (_BossRaid *BossRaidSession) ValidSignature(_sender common.Address, _amount *big.Int, _sig [][]byte) (bool, error) {
	return _BossRaid.Contract.ValidSignature(&_BossRaid.CallOpts, _sender, _amount, _sig)
}

// ValidSignature is a free data retrieval call binding the contract method 0x178c47f5.
//
// Solidity: function validSignature(address _sender, uint256 _amount, bytes[] _sig) view returns(bool)
func (_BossRaid *BossRaidCallerSession) ValidSignature(_sender common.Address, _amount *big.Int, _sig [][]byte) (bool, error) {
	return _BossRaid.Contract.ValidSignature(&_BossRaid.CallOpts, _sender, _amount, _sig)
}

// ZoonerNFT is a free data retrieval call binding the contract method 0xda940945.
//
// Solidity: function zoonerNFT() view returns(address)
func (_BossRaid *BossRaidCaller) ZoonerNFT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BossRaid.contract.Call(opts, &out, "zoonerNFT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ZoonerNFT is a free data retrieval call binding the contract method 0xda940945.
//
// Solidity: function zoonerNFT() view returns(address)
func (_BossRaid *BossRaidSession) ZoonerNFT() (common.Address, error) {
	return _BossRaid.Contract.ZoonerNFT(&_BossRaid.CallOpts)
}

// ZoonerNFT is a free data retrieval call binding the contract method 0xda940945.
//
// Solidity: function zoonerNFT() view returns(address)
func (_BossRaid *BossRaidCallerSession) ZoonerNFT() (common.Address, error) {
	return _BossRaid.Contract.ZoonerNFT(&_BossRaid.CallOpts)
}

// BattleBoss is a paid mutator transaction binding the contract method 0x68ccbfb2.
//
// Solidity: function battleBoss(uint256 _amount, uint256[] _tokenIds) returns()
func (_BossRaid *BossRaidTransactor) BattleBoss(opts *bind.TransactOpts, _amount *big.Int, _tokenIds []*big.Int) (*types.Transaction, error) {
	return _BossRaid.contract.Transact(opts, "battleBoss", _amount, _tokenIds)
}

// BattleBoss is a paid mutator transaction binding the contract method 0x68ccbfb2.
//
// Solidity: function battleBoss(uint256 _amount, uint256[] _tokenIds) returns()
func (_BossRaid *BossRaidSession) BattleBoss(_amount *big.Int, _tokenIds []*big.Int) (*types.Transaction, error) {
	return _BossRaid.Contract.BattleBoss(&_BossRaid.TransactOpts, _amount, _tokenIds)
}

// BattleBoss is a paid mutator transaction binding the contract method 0x68ccbfb2.
//
// Solidity: function battleBoss(uint256 _amount, uint256[] _tokenIds) returns()
func (_BossRaid *BossRaidTransactorSession) BattleBoss(_amount *big.Int, _tokenIds []*big.Int) (*types.Transaction, error) {
	return _BossRaid.Contract.BattleBoss(&_BossRaid.TransactOpts, _amount, _tokenIds)
}

// BeginBattle is a paid mutator transaction binding the contract method 0xccaf4430.
//
// Solidity: function beginBattle() returns()
func (_BossRaid *BossRaidTransactor) BeginBattle(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BossRaid.contract.Transact(opts, "beginBattle")
}

// BeginBattle is a paid mutator transaction binding the contract method 0xccaf4430.
//
// Solidity: function beginBattle() returns()
func (_BossRaid *BossRaidSession) BeginBattle() (*types.Transaction, error) {
	return _BossRaid.Contract.BeginBattle(&_BossRaid.TransactOpts)
}

// BeginBattle is a paid mutator transaction binding the contract method 0xccaf4430.
//
// Solidity: function beginBattle() returns()
func (_BossRaid *BossRaidTransactorSession) BeginBattle() (*types.Transaction, error) {
	return _BossRaid.Contract.BeginBattle(&_BossRaid.TransactOpts)
}

// BeginFirst is a paid mutator transaction binding the contract method 0x827ae457.
//
// Solidity: function beginFirst() returns()
func (_BossRaid *BossRaidTransactor) BeginFirst(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BossRaid.contract.Transact(opts, "beginFirst")
}

// BeginFirst is a paid mutator transaction binding the contract method 0x827ae457.
//
// Solidity: function beginFirst() returns()
func (_BossRaid *BossRaidSession) BeginFirst() (*types.Transaction, error) {
	return _BossRaid.Contract.BeginFirst(&_BossRaid.TransactOpts)
}

// BeginFirst is a paid mutator transaction binding the contract method 0x827ae457.
//
// Solidity: function beginFirst() returns()
func (_BossRaid *BossRaidTransactorSession) BeginFirst() (*types.Transaction, error) {
	return _BossRaid.Contract.BeginFirst(&_BossRaid.TransactOpts)
}

// BeginRound is a paid mutator transaction binding the contract method 0x65fdb08e.
//
// Solidity: function beginRound() returns()
func (_BossRaid *BossRaidTransactor) BeginRound(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BossRaid.contract.Transact(opts, "beginRound")
}

// BeginRound is a paid mutator transaction binding the contract method 0x65fdb08e.
//
// Solidity: function beginRound() returns()
func (_BossRaid *BossRaidSession) BeginRound() (*types.Transaction, error) {
	return _BossRaid.Contract.BeginRound(&_BossRaid.TransactOpts)
}

// BeginRound is a paid mutator transaction binding the contract method 0x65fdb08e.
//
// Solidity: function beginRound() returns()
func (_BossRaid *BossRaidTransactorSession) BeginRound() (*types.Transaction, error) {
	return _BossRaid.Contract.BeginRound(&_BossRaid.TransactOpts)
}

// Harvest is a paid mutator transaction binding the contract method 0xedb2f394.
//
// Solidity: function harvest(uint256 amount, bytes[] _sig) returns()
func (_BossRaid *BossRaidTransactor) Harvest(opts *bind.TransactOpts, amount *big.Int, _sig [][]byte) (*types.Transaction, error) {
	return _BossRaid.contract.Transact(opts, "harvest", amount, _sig)
}

// Harvest is a paid mutator transaction binding the contract method 0xedb2f394.
//
// Solidity: function harvest(uint256 amount, bytes[] _sig) returns()
func (_BossRaid *BossRaidSession) Harvest(amount *big.Int, _sig [][]byte) (*types.Transaction, error) {
	return _BossRaid.Contract.Harvest(&_BossRaid.TransactOpts, amount, _sig)
}

// Harvest is a paid mutator transaction binding the contract method 0xedb2f394.
//
// Solidity: function harvest(uint256 amount, bytes[] _sig) returns()
func (_BossRaid *BossRaidTransactorSession) Harvest(amount *big.Int, _sig [][]byte) (*types.Transaction, error) {
	return _BossRaid.Contract.Harvest(&_BossRaid.TransactOpts, amount, _sig)
}

// Initialize is a paid mutator transaction binding the contract method 0xa224cee7.
//
// Solidity: function initialize(address[] _owners) returns()
func (_BossRaid *BossRaidTransactor) Initialize(opts *bind.TransactOpts, _owners []common.Address) (*types.Transaction, error) {
	return _BossRaid.contract.Transact(opts, "initialize", _owners)
}

// Initialize is a paid mutator transaction binding the contract method 0xa224cee7.
//
// Solidity: function initialize(address[] _owners) returns()
func (_BossRaid *BossRaidSession) Initialize(_owners []common.Address) (*types.Transaction, error) {
	return _BossRaid.Contract.Initialize(&_BossRaid.TransactOpts, _owners)
}

// Initialize is a paid mutator transaction binding the contract method 0xa224cee7.
//
// Solidity: function initialize(address[] _owners) returns()
func (_BossRaid *BossRaidTransactorSession) Initialize(_owners []common.Address) (*types.Transaction, error) {
	return _BossRaid.Contract.Initialize(&_BossRaid.TransactOpts, _owners)
}

// SetAddr is a paid mutator transaction binding the contract method 0xebcb00e0.
//
// Solidity: function setAddr(address _tokenAddr, address _nftAddr) returns()
func (_BossRaid *BossRaidTransactor) SetAddr(opts *bind.TransactOpts, _tokenAddr common.Address, _nftAddr common.Address) (*types.Transaction, error) {
	return _BossRaid.contract.Transact(opts, "setAddr", _tokenAddr, _nftAddr)
}

// SetAddr is a paid mutator transaction binding the contract method 0xebcb00e0.
//
// Solidity: function setAddr(address _tokenAddr, address _nftAddr) returns()
func (_BossRaid *BossRaidSession) SetAddr(_tokenAddr common.Address, _nftAddr common.Address) (*types.Transaction, error) {
	return _BossRaid.Contract.SetAddr(&_BossRaid.TransactOpts, _tokenAddr, _nftAddr)
}

// SetAddr is a paid mutator transaction binding the contract method 0xebcb00e0.
//
// Solidity: function setAddr(address _tokenAddr, address _nftAddr) returns()
func (_BossRaid *BossRaidTransactorSession) SetAddr(_tokenAddr common.Address, _nftAddr common.Address) (*types.Transaction, error) {
	return _BossRaid.Contract.SetAddr(&_BossRaid.TransactOpts, _tokenAddr, _nftAddr)
}

// BossRaidBattleBossEventIterator is returned from FilterBattleBossEvent and is used to iterate over the raw logs and unpacked data for BattleBossEvent events raised by the BossRaid contract.
type BossRaidBattleBossEventIterator struct {
	Event *BossRaidBattleBossEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BossRaidBattleBossEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BossRaidBattleBossEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BossRaidBattleBossEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BossRaidBattleBossEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BossRaidBattleBossEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BossRaidBattleBossEvent represents a BattleBossEvent event raised by the BossRaid contract.
type BossRaidBattleBossEvent struct {
	Id          *big.Int
	User        common.Address
	Amount      *big.Int
	TotalAmount *big.Int
	Level       *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBattleBossEvent is a free log retrieval operation binding the contract event 0xbce97b44a23b9bc801561e6470ccf4655bfa812878a46ff5fda774ad3aec3572.
//
// Solidity: event BattleBossEvent(uint256 id, address user, uint256 amount, uint256 totalAmount, uint256 level)
func (_BossRaid *BossRaidFilterer) FilterBattleBossEvent(opts *bind.FilterOpts) (*BossRaidBattleBossEventIterator, error) {

	logs, sub, err := _BossRaid.contract.FilterLogs(opts, "BattleBossEvent")
	if err != nil {
		return nil, err
	}
	return &BossRaidBattleBossEventIterator{contract: _BossRaid.contract, event: "BattleBossEvent", logs: logs, sub: sub}, nil
}

// WatchBattleBossEvent is a free log subscription operation binding the contract event 0xbce97b44a23b9bc801561e6470ccf4655bfa812878a46ff5fda774ad3aec3572.
//
// Solidity: event BattleBossEvent(uint256 id, address user, uint256 amount, uint256 totalAmount, uint256 level)
func (_BossRaid *BossRaidFilterer) WatchBattleBossEvent(opts *bind.WatchOpts, sink chan<- *BossRaidBattleBossEvent) (event.Subscription, error) {

	logs, sub, err := _BossRaid.contract.WatchLogs(opts, "BattleBossEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BossRaidBattleBossEvent)
				if err := _BossRaid.contract.UnpackLog(event, "BattleBossEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBattleBossEvent is a log parse operation binding the contract event 0xbce97b44a23b9bc801561e6470ccf4655bfa812878a46ff5fda774ad3aec3572.
//
// Solidity: event BattleBossEvent(uint256 id, address user, uint256 amount, uint256 totalAmount, uint256 level)
func (_BossRaid *BossRaidFilterer) ParseBattleBossEvent(log types.Log) (*BossRaidBattleBossEvent, error) {
	event := new(BossRaidBattleBossEvent)
	if err := _BossRaid.contract.UnpackLog(event, "BattleBossEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BossRaidBattleEndEventIterator is returned from FilterBattleEndEvent and is used to iterate over the raw logs and unpacked data for BattleEndEvent events raised by the BossRaid contract.
type BossRaidBattleEndEventIterator struct {
	Event *BossRaidBattleEndEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BossRaidBattleEndEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BossRaidBattleEndEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BossRaidBattleEndEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BossRaidBattleEndEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BossRaidBattleEndEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BossRaidBattleEndEvent represents a BattleEndEvent event raised by the BossRaid contract.
type BossRaidBattleEndEvent struct {
	Id    *big.Int
	Level *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterBattleEndEvent is a free log retrieval operation binding the contract event 0x51f0ed65dace4c8f2bd6a9f3b3f929ca4322218cc5f8835307af742ada59e394.
//
// Solidity: event BattleEndEvent(uint256 id, uint256 level)
func (_BossRaid *BossRaidFilterer) FilterBattleEndEvent(opts *bind.FilterOpts) (*BossRaidBattleEndEventIterator, error) {

	logs, sub, err := _BossRaid.contract.FilterLogs(opts, "BattleEndEvent")
	if err != nil {
		return nil, err
	}
	return &BossRaidBattleEndEventIterator{contract: _BossRaid.contract, event: "BattleEndEvent", logs: logs, sub: sub}, nil
}

// WatchBattleEndEvent is a free log subscription operation binding the contract event 0x51f0ed65dace4c8f2bd6a9f3b3f929ca4322218cc5f8835307af742ada59e394.
//
// Solidity: event BattleEndEvent(uint256 id, uint256 level)
func (_BossRaid *BossRaidFilterer) WatchBattleEndEvent(opts *bind.WatchOpts, sink chan<- *BossRaidBattleEndEvent) (event.Subscription, error) {

	logs, sub, err := _BossRaid.contract.WatchLogs(opts, "BattleEndEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BossRaidBattleEndEvent)
				if err := _BossRaid.contract.UnpackLog(event, "BattleEndEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBattleEndEvent is a log parse operation binding the contract event 0x51f0ed65dace4c8f2bd6a9f3b3f929ca4322218cc5f8835307af742ada59e394.
//
// Solidity: event BattleEndEvent(uint256 id, uint256 level)
func (_BossRaid *BossRaidFilterer) ParseBattleEndEvent(log types.Log) (*BossRaidBattleEndEvent, error) {
	event := new(BossRaidBattleEndEvent)
	if err := _BossRaid.contract.UnpackLog(event, "BattleEndEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BossRaidBattleRaidEventIterator is returned from FilterBattleRaidEvent and is used to iterate over the raw logs and unpacked data for BattleRaidEvent events raised by the BossRaid contract.
type BossRaidBattleRaidEventIterator struct {
	Event *BossRaidBattleRaidEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BossRaidBattleRaidEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BossRaidBattleRaidEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BossRaidBattleRaidEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BossRaidBattleRaidEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BossRaidBattleRaidEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BossRaidBattleRaidEvent represents a BattleRaidEvent event raised by the BossRaid contract.
type BossRaidBattleRaidEvent struct {
	Id          *big.Int
	Level       *big.Int
	TotalDamage *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBattleRaidEvent is a free log retrieval operation binding the contract event 0xea9aa59269422fb735ba8a12f0351bcb87b4f4843d8d791844d4cf583eb67f43.
//
// Solidity: event BattleRaidEvent(uint256 id, uint256 level, uint256 totalDamage)
func (_BossRaid *BossRaidFilterer) FilterBattleRaidEvent(opts *bind.FilterOpts) (*BossRaidBattleRaidEventIterator, error) {

	logs, sub, err := _BossRaid.contract.FilterLogs(opts, "BattleRaidEvent")
	if err != nil {
		return nil, err
	}
	return &BossRaidBattleRaidEventIterator{contract: _BossRaid.contract, event: "BattleRaidEvent", logs: logs, sub: sub}, nil
}

// WatchBattleRaidEvent is a free log subscription operation binding the contract event 0xea9aa59269422fb735ba8a12f0351bcb87b4f4843d8d791844d4cf583eb67f43.
//
// Solidity: event BattleRaidEvent(uint256 id, uint256 level, uint256 totalDamage)
func (_BossRaid *BossRaidFilterer) WatchBattleRaidEvent(opts *bind.WatchOpts, sink chan<- *BossRaidBattleRaidEvent) (event.Subscription, error) {

	logs, sub, err := _BossRaid.contract.WatchLogs(opts, "BattleRaidEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BossRaidBattleRaidEvent)
				if err := _BossRaid.contract.UnpackLog(event, "BattleRaidEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBattleRaidEvent is a log parse operation binding the contract event 0xea9aa59269422fb735ba8a12f0351bcb87b4f4843d8d791844d4cf583eb67f43.
//
// Solidity: event BattleRaidEvent(uint256 id, uint256 level, uint256 totalDamage)
func (_BossRaid *BossRaidFilterer) ParseBattleRaidEvent(log types.Log) (*BossRaidBattleRaidEvent, error) {
	event := new(BossRaidBattleRaidEvent)
	if err := _BossRaid.contract.UnpackLog(event, "BattleRaidEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BossRaidBossBeginEventIterator is returned from FilterBossBeginEvent and is used to iterate over the raw logs and unpacked data for BossBeginEvent events raised by the BossRaid contract.
type BossRaidBossBeginEventIterator struct {
	Event *BossRaidBossBeginEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BossRaidBossBeginEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BossRaidBossBeginEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BossRaidBossBeginEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BossRaidBossBeginEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BossRaidBossBeginEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BossRaidBossBeginEvent represents a BossBeginEvent event raised by the BossRaid contract.
type BossRaidBossBeginEvent struct {
	Id          *big.Int
	Level       *big.Int
	StartTime   *big.Int
	EndTime     *big.Int
	BloodVolume *big.Int
	Value       *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBossBeginEvent is a free log retrieval operation binding the contract event 0x29149048ba6919815a437cb8fcdbb74d2d3535b21b01f4598d0275a2dde3e187.
//
// Solidity: event BossBeginEvent(uint256 id, uint256 level, uint256 startTime, uint256 endTime, uint256 bloodVolume, uint256 value)
func (_BossRaid *BossRaidFilterer) FilterBossBeginEvent(opts *bind.FilterOpts) (*BossRaidBossBeginEventIterator, error) {

	logs, sub, err := _BossRaid.contract.FilterLogs(opts, "BossBeginEvent")
	if err != nil {
		return nil, err
	}
	return &BossRaidBossBeginEventIterator{contract: _BossRaid.contract, event: "BossBeginEvent", logs: logs, sub: sub}, nil
}

// WatchBossBeginEvent is a free log subscription operation binding the contract event 0x29149048ba6919815a437cb8fcdbb74d2d3535b21b01f4598d0275a2dde3e187.
//
// Solidity: event BossBeginEvent(uint256 id, uint256 level, uint256 startTime, uint256 endTime, uint256 bloodVolume, uint256 value)
func (_BossRaid *BossRaidFilterer) WatchBossBeginEvent(opts *bind.WatchOpts, sink chan<- *BossRaidBossBeginEvent) (event.Subscription, error) {

	logs, sub, err := _BossRaid.contract.WatchLogs(opts, "BossBeginEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BossRaidBossBeginEvent)
				if err := _BossRaid.contract.UnpackLog(event, "BossBeginEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBossBeginEvent is a log parse operation binding the contract event 0x29149048ba6919815a437cb8fcdbb74d2d3535b21b01f4598d0275a2dde3e187.
//
// Solidity: event BossBeginEvent(uint256 id, uint256 level, uint256 startTime, uint256 endTime, uint256 bloodVolume, uint256 value)
func (_BossRaid *BossRaidFilterer) ParseBossBeginEvent(log types.Log) (*BossRaidBossBeginEvent, error) {
	event := new(BossRaidBossBeginEvent)
	if err := _BossRaid.contract.UnpackLog(event, "BossBeginEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BossRaidHarvestEventIterator is returned from FilterHarvestEvent and is used to iterate over the raw logs and unpacked data for HarvestEvent events raised by the BossRaid contract.
type BossRaidHarvestEventIterator struct {
	Event *BossRaidHarvestEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BossRaidHarvestEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BossRaidHarvestEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BossRaidHarvestEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BossRaidHarvestEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BossRaidHarvestEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BossRaidHarvestEvent represents a HarvestEvent event raised by the BossRaid contract.
type BossRaidHarvestEvent struct {
	Sender  common.Address
	OrderId *big.Int
	Time    *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterHarvestEvent is a free log retrieval operation binding the contract event 0x6d704856f5fcddffdcc513ff44645bd9829ac3205b60b48a44800219d73d3291.
//
// Solidity: event HarvestEvent(address sender, uint256 orderId, uint256 time)
func (_BossRaid *BossRaidFilterer) FilterHarvestEvent(opts *bind.FilterOpts) (*BossRaidHarvestEventIterator, error) {

	logs, sub, err := _BossRaid.contract.FilterLogs(opts, "HarvestEvent")
	if err != nil {
		return nil, err
	}
	return &BossRaidHarvestEventIterator{contract: _BossRaid.contract, event: "HarvestEvent", logs: logs, sub: sub}, nil
}

// WatchHarvestEvent is a free log subscription operation binding the contract event 0x6d704856f5fcddffdcc513ff44645bd9829ac3205b60b48a44800219d73d3291.
//
// Solidity: event HarvestEvent(address sender, uint256 orderId, uint256 time)
func (_BossRaid *BossRaidFilterer) WatchHarvestEvent(opts *bind.WatchOpts, sink chan<- *BossRaidHarvestEvent) (event.Subscription, error) {

	logs, sub, err := _BossRaid.contract.WatchLogs(opts, "HarvestEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BossRaidHarvestEvent)
				if err := _BossRaid.contract.UnpackLog(event, "HarvestEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseHarvestEvent is a log parse operation binding the contract event 0x6d704856f5fcddffdcc513ff44645bd9829ac3205b60b48a44800219d73d3291.
//
// Solidity: event HarvestEvent(address sender, uint256 orderId, uint256 time)
func (_BossRaid *BossRaidFilterer) ParseHarvestEvent(log types.Log) (*BossRaidHarvestEvent, error) {
	event := new(BossRaidHarvestEvent)
	if err := _BossRaid.contract.UnpackLog(event, "HarvestEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BossRaidInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the BossRaid contract.
type BossRaidInitializedIterator struct {
	Event *BossRaidInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BossRaidInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BossRaidInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BossRaidInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BossRaidInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BossRaidInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BossRaidInitialized represents a Initialized event raised by the BossRaid contract.
type BossRaidInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_BossRaid *BossRaidFilterer) FilterInitialized(opts *bind.FilterOpts) (*BossRaidInitializedIterator, error) {

	logs, sub, err := _BossRaid.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &BossRaidInitializedIterator{contract: _BossRaid.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_BossRaid *BossRaidFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *BossRaidInitialized) (event.Subscription, error) {

	logs, sub, err := _BossRaid.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BossRaidInitialized)
				if err := _BossRaid.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_BossRaid *BossRaidFilterer) ParseInitialized(log types.Log) (*BossRaidInitialized, error) {
	event := new(BossRaidInitialized)
	if err := _BossRaid.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
