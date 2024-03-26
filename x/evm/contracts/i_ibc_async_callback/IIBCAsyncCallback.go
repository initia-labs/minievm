// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package i_ibc_async_callback

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

// IIbcAsyncCallbackMetaData contains all meta data concerning the IIbcAsyncCallback contract.
var IIbcAsyncCallbackMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"callback_id\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"name\":\"ibc_ack\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"callback_id\",\"type\":\"uint64\"}],\"name\":\"ibc_timeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IIbcAsyncCallbackABI is the input ABI used to generate the binding from.
// Deprecated: Use IIbcAsyncCallbackMetaData.ABI instead.
var IIbcAsyncCallbackABI = IIbcAsyncCallbackMetaData.ABI

// IIbcAsyncCallback is an auto generated Go binding around an Ethereum contract.
type IIbcAsyncCallback struct {
	IIbcAsyncCallbackCaller     // Read-only binding to the contract
	IIbcAsyncCallbackTransactor // Write-only binding to the contract
	IIbcAsyncCallbackFilterer   // Log filterer for contract events
}

// IIbcAsyncCallbackCaller is an auto generated read-only Go binding around an Ethereum contract.
type IIbcAsyncCallbackCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IIbcAsyncCallbackTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IIbcAsyncCallbackTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IIbcAsyncCallbackFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IIbcAsyncCallbackFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IIbcAsyncCallbackSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IIbcAsyncCallbackSession struct {
	Contract     *IIbcAsyncCallback // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IIbcAsyncCallbackCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IIbcAsyncCallbackCallerSession struct {
	Contract *IIbcAsyncCallbackCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// IIbcAsyncCallbackTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IIbcAsyncCallbackTransactorSession struct {
	Contract     *IIbcAsyncCallbackTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// IIbcAsyncCallbackRaw is an auto generated low-level Go binding around an Ethereum contract.
type IIbcAsyncCallbackRaw struct {
	Contract *IIbcAsyncCallback // Generic contract binding to access the raw methods on
}

// IIbcAsyncCallbackCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IIbcAsyncCallbackCallerRaw struct {
	Contract *IIbcAsyncCallbackCaller // Generic read-only contract binding to access the raw methods on
}

// IIbcAsyncCallbackTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IIbcAsyncCallbackTransactorRaw struct {
	Contract *IIbcAsyncCallbackTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIIbcAsyncCallback creates a new instance of IIbcAsyncCallback, bound to a specific deployed contract.
func NewIIbcAsyncCallback(address common.Address, backend bind.ContractBackend) (*IIbcAsyncCallback, error) {
	contract, err := bindIIbcAsyncCallback(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IIbcAsyncCallback{IIbcAsyncCallbackCaller: IIbcAsyncCallbackCaller{contract: contract}, IIbcAsyncCallbackTransactor: IIbcAsyncCallbackTransactor{contract: contract}, IIbcAsyncCallbackFilterer: IIbcAsyncCallbackFilterer{contract: contract}}, nil
}

// NewIIbcAsyncCallbackCaller creates a new read-only instance of IIbcAsyncCallback, bound to a specific deployed contract.
func NewIIbcAsyncCallbackCaller(address common.Address, caller bind.ContractCaller) (*IIbcAsyncCallbackCaller, error) {
	contract, err := bindIIbcAsyncCallback(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IIbcAsyncCallbackCaller{contract: contract}, nil
}

// NewIIbcAsyncCallbackTransactor creates a new write-only instance of IIbcAsyncCallback, bound to a specific deployed contract.
func NewIIbcAsyncCallbackTransactor(address common.Address, transactor bind.ContractTransactor) (*IIbcAsyncCallbackTransactor, error) {
	contract, err := bindIIbcAsyncCallback(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IIbcAsyncCallbackTransactor{contract: contract}, nil
}

// NewIIbcAsyncCallbackFilterer creates a new log filterer instance of IIbcAsyncCallback, bound to a specific deployed contract.
func NewIIbcAsyncCallbackFilterer(address common.Address, filterer bind.ContractFilterer) (*IIbcAsyncCallbackFilterer, error) {
	contract, err := bindIIbcAsyncCallback(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IIbcAsyncCallbackFilterer{contract: contract}, nil
}

// bindIIbcAsyncCallback binds a generic wrapper to an already deployed contract.
func bindIIbcAsyncCallback(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IIbcAsyncCallbackMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IIbcAsyncCallback *IIbcAsyncCallbackRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IIbcAsyncCallback.Contract.IIbcAsyncCallbackCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IIbcAsyncCallback *IIbcAsyncCallbackRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IIbcAsyncCallback.Contract.IIbcAsyncCallbackTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IIbcAsyncCallback *IIbcAsyncCallbackRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IIbcAsyncCallback.Contract.IIbcAsyncCallbackTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IIbcAsyncCallback *IIbcAsyncCallbackCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IIbcAsyncCallback.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IIbcAsyncCallback *IIbcAsyncCallbackTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IIbcAsyncCallback.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IIbcAsyncCallback *IIbcAsyncCallbackTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IIbcAsyncCallback.Contract.contract.Transact(opts, method, params...)
}

// IbcAck is a paid mutator transaction binding the contract method 0x0d4f1f9d.
//
// Solidity: function ibc_ack(uint64 callback_id, bool success) returns()
func (_IIbcAsyncCallback *IIbcAsyncCallbackTransactor) IbcAck(opts *bind.TransactOpts, callback_id uint64, success bool) (*types.Transaction, error) {
	return _IIbcAsyncCallback.contract.Transact(opts, "ibc_ack", callback_id, success)
}

// IbcAck is a paid mutator transaction binding the contract method 0x0d4f1f9d.
//
// Solidity: function ibc_ack(uint64 callback_id, bool success) returns()
func (_IIbcAsyncCallback *IIbcAsyncCallbackSession) IbcAck(callback_id uint64, success bool) (*types.Transaction, error) {
	return _IIbcAsyncCallback.Contract.IbcAck(&_IIbcAsyncCallback.TransactOpts, callback_id, success)
}

// IbcAck is a paid mutator transaction binding the contract method 0x0d4f1f9d.
//
// Solidity: function ibc_ack(uint64 callback_id, bool success) returns()
func (_IIbcAsyncCallback *IIbcAsyncCallbackTransactorSession) IbcAck(callback_id uint64, success bool) (*types.Transaction, error) {
	return _IIbcAsyncCallback.Contract.IbcAck(&_IIbcAsyncCallback.TransactOpts, callback_id, success)
}

// IbcTimeout is a paid mutator transaction binding the contract method 0x31a503f0.
//
// Solidity: function ibc_timeout(uint64 callback_id) returns()
func (_IIbcAsyncCallback *IIbcAsyncCallbackTransactor) IbcTimeout(opts *bind.TransactOpts, callback_id uint64) (*types.Transaction, error) {
	return _IIbcAsyncCallback.contract.Transact(opts, "ibc_timeout", callback_id)
}

// IbcTimeout is a paid mutator transaction binding the contract method 0x31a503f0.
//
// Solidity: function ibc_timeout(uint64 callback_id) returns()
func (_IIbcAsyncCallback *IIbcAsyncCallbackSession) IbcTimeout(callback_id uint64) (*types.Transaction, error) {
	return _IIbcAsyncCallback.Contract.IbcTimeout(&_IIbcAsyncCallback.TransactOpts, callback_id)
}

// IbcTimeout is a paid mutator transaction binding the contract method 0x31a503f0.
//
// Solidity: function ibc_timeout(uint64 callback_id) returns()
func (_IIbcAsyncCallback *IIbcAsyncCallbackTransactorSession) IbcTimeout(callback_id uint64) (*types.Transaction, error) {
	return _IIbcAsyncCallback.Contract.IbcTimeout(&_IIbcAsyncCallback.TransactOpts, callback_id)
}
