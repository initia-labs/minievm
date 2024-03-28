// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package i_cosmos

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

// ICosmosMetaData contains all meta data concerning the ICosmos contract.
var ICosmosMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"msg\",\"type\":\"string\"}],\"name\":\"execute_cosmos_message\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"evm_address\",\"type\":\"address\"}],\"name\":\"to_cosmos_address\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"cosmos_address\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"cosmos_address\",\"type\":\"string\"}],\"name\":\"to_evm_address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"evm_address\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ICosmosABI is the input ABI used to generate the binding from.
// Deprecated: Use ICosmosMetaData.ABI instead.
var ICosmosABI = ICosmosMetaData.ABI

// ICosmos is an auto generated Go binding around an Ethereum contract.
type ICosmos struct {
	ICosmosCaller     // Read-only binding to the contract
	ICosmosTransactor // Write-only binding to the contract
	ICosmosFilterer   // Log filterer for contract events
}

// ICosmosCaller is an auto generated read-only Go binding around an Ethereum contract.
type ICosmosCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICosmosTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ICosmosTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICosmosFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ICosmosFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICosmosSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ICosmosSession struct {
	Contract     *ICosmos          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ICosmosCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ICosmosCallerSession struct {
	Contract *ICosmosCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ICosmosTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ICosmosTransactorSession struct {
	Contract     *ICosmosTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ICosmosRaw is an auto generated low-level Go binding around an Ethereum contract.
type ICosmosRaw struct {
	Contract *ICosmos // Generic contract binding to access the raw methods on
}

// ICosmosCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ICosmosCallerRaw struct {
	Contract *ICosmosCaller // Generic read-only contract binding to access the raw methods on
}

// ICosmosTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ICosmosTransactorRaw struct {
	Contract *ICosmosTransactor // Generic write-only contract binding to access the raw methods on
}

// NewICosmos creates a new instance of ICosmos, bound to a specific deployed contract.
func NewICosmos(address common.Address, backend bind.ContractBackend) (*ICosmos, error) {
	contract, err := bindICosmos(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ICosmos{ICosmosCaller: ICosmosCaller{contract: contract}, ICosmosTransactor: ICosmosTransactor{contract: contract}, ICosmosFilterer: ICosmosFilterer{contract: contract}}, nil
}

// NewICosmosCaller creates a new read-only instance of ICosmos, bound to a specific deployed contract.
func NewICosmosCaller(address common.Address, caller bind.ContractCaller) (*ICosmosCaller, error) {
	contract, err := bindICosmos(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ICosmosCaller{contract: contract}, nil
}

// NewICosmosTransactor creates a new write-only instance of ICosmos, bound to a specific deployed contract.
func NewICosmosTransactor(address common.Address, transactor bind.ContractTransactor) (*ICosmosTransactor, error) {
	contract, err := bindICosmos(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ICosmosTransactor{contract: contract}, nil
}

// NewICosmosFilterer creates a new log filterer instance of ICosmos, bound to a specific deployed contract.
func NewICosmosFilterer(address common.Address, filterer bind.ContractFilterer) (*ICosmosFilterer, error) {
	contract, err := bindICosmos(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ICosmosFilterer{contract: contract}, nil
}

// bindICosmos binds a generic wrapper to an already deployed contract.
func bindICosmos(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ICosmosMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICosmos *ICosmosRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICosmos.Contract.ICosmosCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICosmos *ICosmosRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICosmos.Contract.ICosmosTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICosmos *ICosmosRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICosmos.Contract.ICosmosTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICosmos *ICosmosCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICosmos.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICosmos *ICosmosTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICosmos.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICosmos *ICosmosTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICosmos.Contract.contract.Transact(opts, method, params...)
}

// ExecuteCosmosMessage is a paid mutator transaction binding the contract method 0x300da594.
//
// Solidity: function execute_cosmos_message(string msg) returns()
func (_ICosmos *ICosmosTransactor) ExecuteCosmosMessage(opts *bind.TransactOpts, msg string) (*types.Transaction, error) {
	return _ICosmos.contract.Transact(opts, "execute_cosmos_message", msg)
}

// ExecuteCosmosMessage is a paid mutator transaction binding the contract method 0x300da594.
//
// Solidity: function execute_cosmos_message(string msg) returns()
func (_ICosmos *ICosmosSession) ExecuteCosmosMessage(msg string) (*types.Transaction, error) {
	return _ICosmos.Contract.ExecuteCosmosMessage(&_ICosmos.TransactOpts, msg)
}

// ExecuteCosmosMessage is a paid mutator transaction binding the contract method 0x300da594.
//
// Solidity: function execute_cosmos_message(string msg) returns()
func (_ICosmos *ICosmosTransactorSession) ExecuteCosmosMessage(msg string) (*types.Transaction, error) {
	return _ICosmos.Contract.ExecuteCosmosMessage(&_ICosmos.TransactOpts, msg)
}

// ToCosmosAddress is a paid mutator transaction binding the contract method 0x6af32a55.
//
// Solidity: function to_cosmos_address(address evm_address) returns(string cosmos_address)
func (_ICosmos *ICosmosTransactor) ToCosmosAddress(opts *bind.TransactOpts, evm_address common.Address) (*types.Transaction, error) {
	return _ICosmos.contract.Transact(opts, "to_cosmos_address", evm_address)
}

// ToCosmosAddress is a paid mutator transaction binding the contract method 0x6af32a55.
//
// Solidity: function to_cosmos_address(address evm_address) returns(string cosmos_address)
func (_ICosmos *ICosmosSession) ToCosmosAddress(evm_address common.Address) (*types.Transaction, error) {
	return _ICosmos.Contract.ToCosmosAddress(&_ICosmos.TransactOpts, evm_address)
}

// ToCosmosAddress is a paid mutator transaction binding the contract method 0x6af32a55.
//
// Solidity: function to_cosmos_address(address evm_address) returns(string cosmos_address)
func (_ICosmos *ICosmosTransactorSession) ToCosmosAddress(evm_address common.Address) (*types.Transaction, error) {
	return _ICosmos.Contract.ToCosmosAddress(&_ICosmos.TransactOpts, evm_address)
}

// ToEvmAddress is a paid mutator transaction binding the contract method 0x4f94a155.
//
// Solidity: function to_evm_address(string cosmos_address) returns(address evm_address)
func (_ICosmos *ICosmosTransactor) ToEvmAddress(opts *bind.TransactOpts, cosmos_address string) (*types.Transaction, error) {
	return _ICosmos.contract.Transact(opts, "to_evm_address", cosmos_address)
}

// ToEvmAddress is a paid mutator transaction binding the contract method 0x4f94a155.
//
// Solidity: function to_evm_address(string cosmos_address) returns(address evm_address)
func (_ICosmos *ICosmosSession) ToEvmAddress(cosmos_address string) (*types.Transaction, error) {
	return _ICosmos.Contract.ToEvmAddress(&_ICosmos.TransactOpts, cosmos_address)
}

// ToEvmAddress is a paid mutator transaction binding the contract method 0x4f94a155.
//
// Solidity: function to_evm_address(string cosmos_address) returns(address evm_address)
func (_ICosmos *ICosmosTransactorSession) ToEvmAddress(cosmos_address string) (*types.Transaction, error) {
	return _ICosmos.Contract.ToEvmAddress(&_ICosmos.TransactOpts, cosmos_address)
}
