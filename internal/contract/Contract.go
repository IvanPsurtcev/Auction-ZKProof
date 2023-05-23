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

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_startingPrice\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"addressArray\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bidWin\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"bids\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"endAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ended\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_coefficient\",\"type\":\"uint256\"}],\"name\":\"makeBid\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"seller\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_started\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_ended\",\"type\":\"bool\"}],\"name\":\"setStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"started\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startingPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"winner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// AddressArray is a free data retrieval call binding the contract method 0x0f96cf34.
//
// Solidity: function addressArray(uint256 ) view returns(address)
func (_Contract *ContractCaller) AddressArray(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "addressArray", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressArray is a free data retrieval call binding the contract method 0x0f96cf34.
//
// Solidity: function addressArray(uint256 ) view returns(address)
func (_Contract *ContractSession) AddressArray(arg0 *big.Int) (common.Address, error) {
	return _Contract.Contract.AddressArray(&_Contract.CallOpts, arg0)
}

// AddressArray is a free data retrieval call binding the contract method 0x0f96cf34.
//
// Solidity: function addressArray(uint256 ) view returns(address)
func (_Contract *ContractCallerSession) AddressArray(arg0 *big.Int) (common.Address, error) {
	return _Contract.Contract.AddressArray(&_Contract.CallOpts, arg0)
}

// Bids is a free data retrieval call binding the contract method 0x62ea82db.
//
// Solidity: function bids(address ) view returns(uint256)
func (_Contract *ContractCaller) Bids(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "bids", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Bids is a free data retrieval call binding the contract method 0x62ea82db.
//
// Solidity: function bids(address ) view returns(uint256)
func (_Contract *ContractSession) Bids(arg0 common.Address) (*big.Int, error) {
	return _Contract.Contract.Bids(&_Contract.CallOpts, arg0)
}

// Bids is a free data retrieval call binding the contract method 0x62ea82db.
//
// Solidity: function bids(address ) view returns(uint256)
func (_Contract *ContractCallerSession) Bids(arg0 common.Address) (*big.Int, error) {
	return _Contract.Contract.Bids(&_Contract.CallOpts, arg0)
}

// Ended is a free data retrieval call binding the contract method 0x12fa6feb.
//
// Solidity: function ended() view returns(bool)
func (_Contract *ContractCaller) Ended(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "ended")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Ended is a free data retrieval call binding the contract method 0x12fa6feb.
//
// Solidity: function ended() view returns(bool)
func (_Contract *ContractSession) Ended() (bool, error) {
	return _Contract.Contract.Ended(&_Contract.CallOpts)
}

// Ended is a free data retrieval call binding the contract method 0x12fa6feb.
//
// Solidity: function ended() view returns(bool)
func (_Contract *ContractCallerSession) Ended() (bool, error) {
	return _Contract.Contract.Ended(&_Contract.CallOpts)
}

// Seller is a free data retrieval call binding the contract method 0x08551a53.
//
// Solidity: function seller() view returns(address)
func (_Contract *ContractCaller) Seller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "seller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Seller is a free data retrieval call binding the contract method 0x08551a53.
//
// Solidity: function seller() view returns(address)
func (_Contract *ContractSession) Seller() (common.Address, error) {
	return _Contract.Contract.Seller(&_Contract.CallOpts)
}

// Seller is a free data retrieval call binding the contract method 0x08551a53.
//
// Solidity: function seller() view returns(address)
func (_Contract *ContractCallerSession) Seller() (common.Address, error) {
	return _Contract.Contract.Seller(&_Contract.CallOpts)
}

// Started is a free data retrieval call binding the contract method 0x1f2698ab.
//
// Solidity: function started() view returns(bool)
func (_Contract *ContractCaller) Started(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "started")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Started is a free data retrieval call binding the contract method 0x1f2698ab.
//
// Solidity: function started() view returns(bool)
func (_Contract *ContractSession) Started() (bool, error) {
	return _Contract.Contract.Started(&_Contract.CallOpts)
}

// Started is a free data retrieval call binding the contract method 0x1f2698ab.
//
// Solidity: function started() view returns(bool)
func (_Contract *ContractCallerSession) Started() (bool, error) {
	return _Contract.Contract.Started(&_Contract.CallOpts)
}

// StartingPrice is a free data retrieval call binding the contract method 0xd6fbf202.
//
// Solidity: function startingPrice() view returns(uint256)
func (_Contract *ContractCaller) StartingPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "startingPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartingPrice is a free data retrieval call binding the contract method 0xd6fbf202.
//
// Solidity: function startingPrice() view returns(uint256)
func (_Contract *ContractSession) StartingPrice() (*big.Int, error) {
	return _Contract.Contract.StartingPrice(&_Contract.CallOpts)
}

// StartingPrice is a free data retrieval call binding the contract method 0xd6fbf202.
//
// Solidity: function startingPrice() view returns(uint256)
func (_Contract *ContractCallerSession) StartingPrice() (*big.Int, error) {
	return _Contract.Contract.StartingPrice(&_Contract.CallOpts)
}

// Winner is a free data retrieval call binding the contract method 0xdfbf53ae.
//
// Solidity: function winner() view returns(address)
func (_Contract *ContractCaller) Winner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "winner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Winner is a free data retrieval call binding the contract method 0xdfbf53ae.
//
// Solidity: function winner() view returns(address)
func (_Contract *ContractSession) Winner() (common.Address, error) {
	return _Contract.Contract.Winner(&_Contract.CallOpts)
}

// Winner is a free data retrieval call binding the contract method 0xdfbf53ae.
//
// Solidity: function winner() view returns(address)
func (_Contract *ContractCallerSession) Winner() (common.Address, error) {
	return _Contract.Contract.Winner(&_Contract.CallOpts)
}

// BidWin is a paid mutator transaction binding the contract method 0x45c9df02.
//
// Solidity: function bidWin() payable returns()
func (_Contract *ContractTransactor) BidWin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "bidWin")
}

// BidWin is a paid mutator transaction binding the contract method 0x45c9df02.
//
// Solidity: function bidWin() payable returns()
func (_Contract *ContractSession) BidWin() (*types.Transaction, error) {
	return _Contract.Contract.BidWin(&_Contract.TransactOpts)
}

// BidWin is a paid mutator transaction binding the contract method 0x45c9df02.
//
// Solidity: function bidWin() payable returns()
func (_Contract *ContractTransactorSession) BidWin() (*types.Transaction, error) {
	return _Contract.Contract.BidWin(&_Contract.TransactOpts)
}

// EndAuction is a paid mutator transaction binding the contract method 0xfe67a54b.
//
// Solidity: function endAuction() returns()
func (_Contract *ContractTransactor) EndAuction(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "endAuction")
}

// EndAuction is a paid mutator transaction binding the contract method 0xfe67a54b.
//
// Solidity: function endAuction() returns()
func (_Contract *ContractSession) EndAuction() (*types.Transaction, error) {
	return _Contract.Contract.EndAuction(&_Contract.TransactOpts)
}

// EndAuction is a paid mutator transaction binding the contract method 0xfe67a54b.
//
// Solidity: function endAuction() returns()
func (_Contract *ContractTransactorSession) EndAuction() (*types.Transaction, error) {
	return _Contract.Contract.EndAuction(&_Contract.TransactOpts)
}

// MakeBid is a paid mutator transaction binding the contract method 0x175b2304.
//
// Solidity: function makeBid(uint256 _coefficient) payable returns()
func (_Contract *ContractTransactor) MakeBid(opts *bind.TransactOpts, _coefficient *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "makeBid", _coefficient)
}

// MakeBid is a paid mutator transaction binding the contract method 0x175b2304.
//
// Solidity: function makeBid(uint256 _coefficient) payable returns()
func (_Contract *ContractSession) MakeBid(_coefficient *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.MakeBid(&_Contract.TransactOpts, _coefficient)
}

// MakeBid is a paid mutator transaction binding the contract method 0x175b2304.
//
// Solidity: function makeBid(uint256 _coefficient) payable returns()
func (_Contract *ContractTransactorSession) MakeBid(_coefficient *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.MakeBid(&_Contract.TransactOpts, _coefficient)
}

// SetStatus is a paid mutator transaction binding the contract method 0xb3c2eac1.
//
// Solidity: function setStatus(bool _started, bool _ended) returns()
func (_Contract *ContractTransactor) SetStatus(opts *bind.TransactOpts, _started bool, _ended bool) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setStatus", _started, _ended)
}

// SetStatus is a paid mutator transaction binding the contract method 0xb3c2eac1.
//
// Solidity: function setStatus(bool _started, bool _ended) returns()
func (_Contract *ContractSession) SetStatus(_started bool, _ended bool) (*types.Transaction, error) {
	return _Contract.Contract.SetStatus(&_Contract.TransactOpts, _started, _ended)
}

// SetStatus is a paid mutator transaction binding the contract method 0xb3c2eac1.
//
// Solidity: function setStatus(bool _started, bool _ended) returns()
func (_Contract *ContractTransactorSession) SetStatus(_started bool, _ended bool) (*types.Transaction, error) {
	return _Contract.Contract.SetStatus(&_Contract.TransactOpts, _started, _ended)
}

// StartAuction is a paid mutator transaction binding the contract method 0x6b64c769.
//
// Solidity: function startAuction() returns()
func (_Contract *ContractTransactor) StartAuction(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "startAuction")
}

// StartAuction is a paid mutator transaction binding the contract method 0x6b64c769.
//
// Solidity: function startAuction() returns()
func (_Contract *ContractSession) StartAuction() (*types.Transaction, error) {
	return _Contract.Contract.StartAuction(&_Contract.TransactOpts)
}

// StartAuction is a paid mutator transaction binding the contract method 0x6b64c769.
//
// Solidity: function startAuction() returns()
func (_Contract *ContractTransactorSession) StartAuction() (*types.Transaction, error) {
	return _Contract.Contract.StartAuction(&_Contract.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Contract *ContractTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Contract *ContractSession) Withdraw() (*types.Transaction, error) {
	return _Contract.Contract.Withdraw(&_Contract.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Contract *ContractTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Contract.Contract.Withdraw(&_Contract.TransactOpts)
}
