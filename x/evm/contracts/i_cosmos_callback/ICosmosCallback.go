// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package i_cosmos_callback

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

// ICosmosCallbackMetaData contains all meta data concerning the ICosmosCallback contract.
var ICosmosCallbackMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"callback_id\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"name\":\"callback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ICosmosCallbackABI is the input ABI used to generate the binding from.
// Deprecated: Use ICosmosCallbackMetaData.ABI instead.
var ICosmosCallbackABI = ICosmosCallbackMetaData.ABI

// ICosmosCallback is an auto generated Go binding around an Ethereum contract.
type ICosmosCallback struct {
	ICosmosCallbackCaller     // Read-only binding to the contract
	ICosmosCallbackTransactor // Write-only binding to the contract
	ICosmosCallbackFilterer   // Log filterer for contract events
}

// ICosmosCallbackCaller is an auto generated read-only Go binding around an Ethereum contract.
type ICosmosCallbackCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICosmosCallbackTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ICosmosCallbackTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICosmosCallbackFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ICosmosCallbackFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICosmosCallbackSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ICosmosCallbackSession struct {
	Contract     *ICosmosCallback  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ICosmosCallbackCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ICosmosCallbackCallerSession struct {
	Contract *ICosmosCallbackCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ICosmosCallbackTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ICosmosCallbackTransactorSession struct {
	Contract     *ICosmosCallbackTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ICosmosCallbackRaw is an auto generated low-level Go binding around an Ethereum contract.
type ICosmosCallbackRaw struct {
	Contract *ICosmosCallback // Generic contract binding to access the raw methods on
}

// ICosmosCallbackCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ICosmosCallbackCallerRaw struct {
	Contract *ICosmosCallbackCaller // Generic read-only contract binding to access the raw methods on
}

// ICosmosCallbackTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ICosmosCallbackTransactorRaw struct {
	Contract *ICosmosCallbackTransactor // Generic write-only contract binding to access the raw methods on
}

// NewICosmosCallback creates a new instance of ICosmosCallback, bound to a specific deployed contract.
func NewICosmosCallback(address common.Address, backend bind.ContractBackend) (*ICosmosCallback, error) {
	contract, err := bindICosmosCallback(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ICosmosCallback{ICosmosCallbackCaller: ICosmosCallbackCaller{contract: contract}, ICosmosCallbackTransactor: ICosmosCallbackTransactor{contract: contract}, ICosmosCallbackFilterer: ICosmosCallbackFilterer{contract: contract}}, nil
}

// NewICosmosCallbackCaller creates a new read-only instance of ICosmosCallback, bound to a specific deployed contract.
func NewICosmosCallbackCaller(address common.Address, caller bind.ContractCaller) (*ICosmosCallbackCaller, error) {
	contract, err := bindICosmosCallback(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ICosmosCallbackCaller{contract: contract}, nil
}

// NewICosmosCallbackTransactor creates a new write-only instance of ICosmosCallback, bound to a specific deployed contract.
func NewICosmosCallbackTransactor(address common.Address, transactor bind.ContractTransactor) (*ICosmosCallbackTransactor, error) {
	contract, err := bindICosmosCallback(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ICosmosCallbackTransactor{contract: contract}, nil
}

// NewICosmosCallbackFilterer creates a new log filterer instance of ICosmosCallback, bound to a specific deployed contract.
func NewICosmosCallbackFilterer(address common.Address, filterer bind.ContractFilterer) (*ICosmosCallbackFilterer, error) {
	contract, err := bindICosmosCallback(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ICosmosCallbackFilterer{contract: contract}, nil
}

// bindICosmosCallback binds a generic wrapper to an already deployed contract.
func bindICosmosCallback(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ICosmosCallbackMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICosmosCallback *ICosmosCallbackRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICosmosCallback.Contract.ICosmosCallbackCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICosmosCallback *ICosmosCallbackRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICosmosCallback.Contract.ICosmosCallbackTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICosmosCallback *ICosmosCallbackRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICosmosCallback.Contract.ICosmosCallbackTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICosmosCallback *ICosmosCallbackCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICosmosCallback.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICosmosCallback *ICosmosCallbackTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICosmosCallback.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICosmosCallback *ICosmosCallbackTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICosmosCallback.Contract.contract.Transact(opts, method, params...)
}

// Callback is a paid mutator transaction binding the contract method 0xc31925a7.
//
// Solidity: function callback(uint64 callback_id, bool success) returns()
func (_ICosmosCallback *ICosmosCallbackTransactor) Callback(opts *bind.TransactOpts, callback_id uint64, success bool) (*types.Transaction, error) {
	return _ICosmosCallback.contract.Transact(opts, "callback", callback_id, success)
}

// Callback is a paid mutator transaction binding the contract method 0xc31925a7.
//
// Solidity: function callback(uint64 callback_id, bool success) returns()
func (_ICosmosCallback *ICosmosCallbackSession) Callback(callback_id uint64, success bool) (*types.Transaction, error) {
	return _ICosmosCallback.Contract.Callback(&_ICosmosCallback.TransactOpts, callback_id, success)
}

// Callback is a paid mutator transaction binding the contract method 0xc31925a7.
//
// Solidity: function callback(uint64 callback_id, bool success) returns()
func (_ICosmosCallback *ICosmosCallbackTransactorSession) Callback(callback_id uint64, success bool) (*types.Transaction, error) {
	return _ICosmosCallback.Contract.Callback(&_ICosmosCallback.TransactOpts, callback_id, success)
}
