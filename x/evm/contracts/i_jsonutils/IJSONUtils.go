// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package i_jsonutils

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

// IJsonutilsMetaData contains all meta data concerning the IJsonutils contract.
var IJsonutilsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dst_json\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"src_json\",\"type\":\"string\"}],\"name\":\"merge_json\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"json\",\"type\":\"string\"}],\"name\":\"stringify_json\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IJsonutilsABI is the input ABI used to generate the binding from.
// Deprecated: Use IJsonutilsMetaData.ABI instead.
var IJsonutilsABI = IJsonutilsMetaData.ABI

// IJsonutils is an auto generated Go binding around an Ethereum contract.
type IJsonutils struct {
	IJsonutilsCaller     // Read-only binding to the contract
	IJsonutilsTransactor // Write-only binding to the contract
	IJsonutilsFilterer   // Log filterer for contract events
}

// IJsonutilsCaller is an auto generated read-only Go binding around an Ethereum contract.
type IJsonutilsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IJsonutilsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IJsonutilsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IJsonutilsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IJsonutilsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IJsonutilsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IJsonutilsSession struct {
	Contract     *IJsonutils       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IJsonutilsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IJsonutilsCallerSession struct {
	Contract *IJsonutilsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// IJsonutilsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IJsonutilsTransactorSession struct {
	Contract     *IJsonutilsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// IJsonutilsRaw is an auto generated low-level Go binding around an Ethereum contract.
type IJsonutilsRaw struct {
	Contract *IJsonutils // Generic contract binding to access the raw methods on
}

// IJsonutilsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IJsonutilsCallerRaw struct {
	Contract *IJsonutilsCaller // Generic read-only contract binding to access the raw methods on
}

// IJsonutilsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IJsonutilsTransactorRaw struct {
	Contract *IJsonutilsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIJsonutils creates a new instance of IJsonutils, bound to a specific deployed contract.
func NewIJsonutils(address common.Address, backend bind.ContractBackend) (*IJsonutils, error) {
	contract, err := bindIJsonutils(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IJsonutils{IJsonutilsCaller: IJsonutilsCaller{contract: contract}, IJsonutilsTransactor: IJsonutilsTransactor{contract: contract}, IJsonutilsFilterer: IJsonutilsFilterer{contract: contract}}, nil
}

// NewIJsonutilsCaller creates a new read-only instance of IJsonutils, bound to a specific deployed contract.
func NewIJsonutilsCaller(address common.Address, caller bind.ContractCaller) (*IJsonutilsCaller, error) {
	contract, err := bindIJsonutils(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IJsonutilsCaller{contract: contract}, nil
}

// NewIJsonutilsTransactor creates a new write-only instance of IJsonutils, bound to a specific deployed contract.
func NewIJsonutilsTransactor(address common.Address, transactor bind.ContractTransactor) (*IJsonutilsTransactor, error) {
	contract, err := bindIJsonutils(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IJsonutilsTransactor{contract: contract}, nil
}

// NewIJsonutilsFilterer creates a new log filterer instance of IJsonutils, bound to a specific deployed contract.
func NewIJsonutilsFilterer(address common.Address, filterer bind.ContractFilterer) (*IJsonutilsFilterer, error) {
	contract, err := bindIJsonutils(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IJsonutilsFilterer{contract: contract}, nil
}

// bindIJsonutils binds a generic wrapper to an already deployed contract.
func bindIJsonutils(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IJsonutilsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IJsonutils *IJsonutilsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IJsonutils.Contract.IJsonutilsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IJsonutils *IJsonutilsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IJsonutils.Contract.IJsonutilsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IJsonutils *IJsonutilsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IJsonutils.Contract.IJsonutilsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IJsonutils *IJsonutilsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IJsonutils.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IJsonutils *IJsonutilsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IJsonutils.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IJsonutils *IJsonutilsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IJsonutils.Contract.contract.Transact(opts, method, params...)
}

// MergeJson is a free data retrieval call binding the contract method 0x5cc855e3.
//
// Solidity: function merge_json(string dst_json, string src_json) view returns(string)
func (_IJsonutils *IJsonutilsCaller) MergeJson(opts *bind.CallOpts, dst_json string, src_json string) (string, error) {
	var out []interface{}
	err := _IJsonutils.contract.Call(opts, &out, "merge_json", dst_json, src_json)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// MergeJson is a free data retrieval call binding the contract method 0x5cc855e3.
//
// Solidity: function merge_json(string dst_json, string src_json) view returns(string)
func (_IJsonutils *IJsonutilsSession) MergeJson(dst_json string, src_json string) (string, error) {
	return _IJsonutils.Contract.MergeJson(&_IJsonutils.CallOpts, dst_json, src_json)
}

// MergeJson is a free data retrieval call binding the contract method 0x5cc855e3.
//
// Solidity: function merge_json(string dst_json, string src_json) view returns(string)
func (_IJsonutils *IJsonutilsCallerSession) MergeJson(dst_json string, src_json string) (string, error) {
	return _IJsonutils.Contract.MergeJson(&_IJsonutils.CallOpts, dst_json, src_json)
}

// StringifyJson is a free data retrieval call binding the contract method 0x8d5c8817.
//
// Solidity: function stringify_json(string json) view returns(string)
func (_IJsonutils *IJsonutilsCaller) StringifyJson(opts *bind.CallOpts, json string) (string, error) {
	var out []interface{}
	err := _IJsonutils.contract.Call(opts, &out, "stringify_json", json)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// StringifyJson is a free data retrieval call binding the contract method 0x8d5c8817.
//
// Solidity: function stringify_json(string json) view returns(string)
func (_IJsonutils *IJsonutilsSession) StringifyJson(json string) (string, error) {
	return _IJsonutils.Contract.StringifyJson(&_IJsonutils.CallOpts, json)
}

// StringifyJson is a free data retrieval call binding the contract method 0x8d5c8817.
//
// Solidity: function stringify_json(string json) view returns(string)
func (_IJsonutils *IJsonutilsCallerSession) StringifyJson(json string) (string, error) {
	return _IJsonutils.Contract.StringifyJson(&_IJsonutils.CallOpts, json)
}
