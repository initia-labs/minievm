// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package i_erc721_registry

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

// IErc721RegistryMetaData contains all meta data concerning the IErc721Registry contract.
var IErc721RegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"is_erc721_store_registered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"registered\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"register_erc721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"register_erc721_store\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IErc721RegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use IErc721RegistryMetaData.ABI instead.
var IErc721RegistryABI = IErc721RegistryMetaData.ABI

// IErc721Registry is an auto generated Go binding around an Ethereum contract.
type IErc721Registry struct {
	IErc721RegistryCaller     // Read-only binding to the contract
	IErc721RegistryTransactor // Write-only binding to the contract
	IErc721RegistryFilterer   // Log filterer for contract events
}

// IErc721RegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type IErc721RegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IErc721RegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IErc721RegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IErc721RegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IErc721RegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IErc721RegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IErc721RegistrySession struct {
	Contract     *IErc721Registry  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IErc721RegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IErc721RegistryCallerSession struct {
	Contract *IErc721RegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// IErc721RegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IErc721RegistryTransactorSession struct {
	Contract     *IErc721RegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IErc721RegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type IErc721RegistryRaw struct {
	Contract *IErc721Registry // Generic contract binding to access the raw methods on
}

// IErc721RegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IErc721RegistryCallerRaw struct {
	Contract *IErc721RegistryCaller // Generic read-only contract binding to access the raw methods on
}

// IErc721RegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IErc721RegistryTransactorRaw struct {
	Contract *IErc721RegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIErc721Registry creates a new instance of IErc721Registry, bound to a specific deployed contract.
func NewIErc721Registry(address common.Address, backend bind.ContractBackend) (*IErc721Registry, error) {
	contract, err := bindIErc721Registry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IErc721Registry{IErc721RegistryCaller: IErc721RegistryCaller{contract: contract}, IErc721RegistryTransactor: IErc721RegistryTransactor{contract: contract}, IErc721RegistryFilterer: IErc721RegistryFilterer{contract: contract}}, nil
}

// NewIErc721RegistryCaller creates a new read-only instance of IErc721Registry, bound to a specific deployed contract.
func NewIErc721RegistryCaller(address common.Address, caller bind.ContractCaller) (*IErc721RegistryCaller, error) {
	contract, err := bindIErc721Registry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IErc721RegistryCaller{contract: contract}, nil
}

// NewIErc721RegistryTransactor creates a new write-only instance of IErc721Registry, bound to a specific deployed contract.
func NewIErc721RegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*IErc721RegistryTransactor, error) {
	contract, err := bindIErc721Registry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IErc721RegistryTransactor{contract: contract}, nil
}

// NewIErc721RegistryFilterer creates a new log filterer instance of IErc721Registry, bound to a specific deployed contract.
func NewIErc721RegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*IErc721RegistryFilterer, error) {
	contract, err := bindIErc721Registry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IErc721RegistryFilterer{contract: contract}, nil
}

// bindIErc721Registry binds a generic wrapper to an already deployed contract.
func bindIErc721Registry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IErc721RegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IErc721Registry *IErc721RegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IErc721Registry.Contract.IErc721RegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IErc721Registry *IErc721RegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IErc721Registry.Contract.IErc721RegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IErc721Registry *IErc721RegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IErc721Registry.Contract.IErc721RegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IErc721Registry *IErc721RegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IErc721Registry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IErc721Registry *IErc721RegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IErc721Registry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IErc721Registry *IErc721RegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IErc721Registry.Contract.contract.Transact(opts, method, params...)
}

// IsErc721StoreRegistered is a free data retrieval call binding the contract method 0xfa75f257.
//
// Solidity: function is_erc721_store_registered(address account) view returns(bool registered)
func (_IErc721Registry *IErc721RegistryCaller) IsErc721StoreRegistered(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _IErc721Registry.contract.Call(opts, &out, "is_erc721_store_registered", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsErc721StoreRegistered is a free data retrieval call binding the contract method 0xfa75f257.
//
// Solidity: function is_erc721_store_registered(address account) view returns(bool registered)
func (_IErc721Registry *IErc721RegistrySession) IsErc721StoreRegistered(account common.Address) (bool, error) {
	return _IErc721Registry.Contract.IsErc721StoreRegistered(&_IErc721Registry.CallOpts, account)
}

// IsErc721StoreRegistered is a free data retrieval call binding the contract method 0xfa75f257.
//
// Solidity: function is_erc721_store_registered(address account) view returns(bool registered)
func (_IErc721Registry *IErc721RegistryCallerSession) IsErc721StoreRegistered(account common.Address) (bool, error) {
	return _IErc721Registry.Contract.IsErc721StoreRegistered(&_IErc721Registry.CallOpts, account)
}

// RegisterErc721 is a paid mutator transaction binding the contract method 0x379da846.
//
// Solidity: function register_erc721() returns()
func (_IErc721Registry *IErc721RegistryTransactor) RegisterErc721(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IErc721Registry.contract.Transact(opts, "register_erc721")
}

// RegisterErc721 is a paid mutator transaction binding the contract method 0x379da846.
//
// Solidity: function register_erc721() returns()
func (_IErc721Registry *IErc721RegistrySession) RegisterErc721() (*types.Transaction, error) {
	return _IErc721Registry.Contract.RegisterErc721(&_IErc721Registry.TransactOpts)
}

// RegisterErc721 is a paid mutator transaction binding the contract method 0x379da846.
//
// Solidity: function register_erc721() returns()
func (_IErc721Registry *IErc721RegistryTransactorSession) RegisterErc721() (*types.Transaction, error) {
	return _IErc721Registry.Contract.RegisterErc721(&_IErc721Registry.TransactOpts)
}

// RegisterErc721Store is a paid mutator transaction binding the contract method 0xd6e69551.
//
// Solidity: function register_erc721_store(address account) returns()
func (_IErc721Registry *IErc721RegistryTransactor) RegisterErc721Store(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _IErc721Registry.contract.Transact(opts, "register_erc721_store", account)
}

// RegisterErc721Store is a paid mutator transaction binding the contract method 0xd6e69551.
//
// Solidity: function register_erc721_store(address account) returns()
func (_IErc721Registry *IErc721RegistrySession) RegisterErc721Store(account common.Address) (*types.Transaction, error) {
	return _IErc721Registry.Contract.RegisterErc721Store(&_IErc721Registry.TransactOpts, account)
}

// RegisterErc721Store is a paid mutator transaction binding the contract method 0xd6e69551.
//
// Solidity: function register_erc721_store(address account) returns()
func (_IErc721Registry *IErc721RegistryTransactorSession) RegisterErc721Store(account common.Address) (*types.Transaction, error) {
	return _IErc721Registry.Contract.RegisterErc721Store(&_IErc721Registry.TransactOpts, account)
}
