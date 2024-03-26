// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package i_erc20_registry

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

// IErc20RegistryMetaData contains all meta data concerning the IErc20Registry contract.
var IErc20RegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"is_erc20_store_registered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"registered\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"register_erc20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"register_erc20_store\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IErc20RegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use IErc20RegistryMetaData.ABI instead.
var IErc20RegistryABI = IErc20RegistryMetaData.ABI

// IErc20Registry is an auto generated Go binding around an Ethereum contract.
type IErc20Registry struct {
	IErc20RegistryCaller     // Read-only binding to the contract
	IErc20RegistryTransactor // Write-only binding to the contract
	IErc20RegistryFilterer   // Log filterer for contract events
}

// IErc20RegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type IErc20RegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IErc20RegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IErc20RegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IErc20RegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IErc20RegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IErc20RegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IErc20RegistrySession struct {
	Contract     *IErc20Registry   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IErc20RegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IErc20RegistryCallerSession struct {
	Contract *IErc20RegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IErc20RegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IErc20RegistryTransactorSession struct {
	Contract     *IErc20RegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IErc20RegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type IErc20RegistryRaw struct {
	Contract *IErc20Registry // Generic contract binding to access the raw methods on
}

// IErc20RegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IErc20RegistryCallerRaw struct {
	Contract *IErc20RegistryCaller // Generic read-only contract binding to access the raw methods on
}

// IErc20RegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IErc20RegistryTransactorRaw struct {
	Contract *IErc20RegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIErc20Registry creates a new instance of IErc20Registry, bound to a specific deployed contract.
func NewIErc20Registry(address common.Address, backend bind.ContractBackend) (*IErc20Registry, error) {
	contract, err := bindIErc20Registry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IErc20Registry{IErc20RegistryCaller: IErc20RegistryCaller{contract: contract}, IErc20RegistryTransactor: IErc20RegistryTransactor{contract: contract}, IErc20RegistryFilterer: IErc20RegistryFilterer{contract: contract}}, nil
}

// NewIErc20RegistryCaller creates a new read-only instance of IErc20Registry, bound to a specific deployed contract.
func NewIErc20RegistryCaller(address common.Address, caller bind.ContractCaller) (*IErc20RegistryCaller, error) {
	contract, err := bindIErc20Registry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IErc20RegistryCaller{contract: contract}, nil
}

// NewIErc20RegistryTransactor creates a new write-only instance of IErc20Registry, bound to a specific deployed contract.
func NewIErc20RegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*IErc20RegistryTransactor, error) {
	contract, err := bindIErc20Registry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IErc20RegistryTransactor{contract: contract}, nil
}

// NewIErc20RegistryFilterer creates a new log filterer instance of IErc20Registry, bound to a specific deployed contract.
func NewIErc20RegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*IErc20RegistryFilterer, error) {
	contract, err := bindIErc20Registry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IErc20RegistryFilterer{contract: contract}, nil
}

// bindIErc20Registry binds a generic wrapper to an already deployed contract.
func bindIErc20Registry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IErc20RegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IErc20Registry *IErc20RegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IErc20Registry.Contract.IErc20RegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IErc20Registry *IErc20RegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IErc20Registry.Contract.IErc20RegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IErc20Registry *IErc20RegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IErc20Registry.Contract.IErc20RegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IErc20Registry *IErc20RegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IErc20Registry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IErc20Registry *IErc20RegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IErc20Registry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IErc20Registry *IErc20RegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IErc20Registry.Contract.contract.Transact(opts, method, params...)
}

// IsErc20StoreRegistered is a free data retrieval call binding the contract method 0x4e25ab64.
//
// Solidity: function is_erc20_store_registered(address account) view returns(bool registered)
func (_IErc20Registry *IErc20RegistryCaller) IsErc20StoreRegistered(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _IErc20Registry.contract.Call(opts, &out, "is_erc20_store_registered", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsErc20StoreRegistered is a free data retrieval call binding the contract method 0x4e25ab64.
//
// Solidity: function is_erc20_store_registered(address account) view returns(bool registered)
func (_IErc20Registry *IErc20RegistrySession) IsErc20StoreRegistered(account common.Address) (bool, error) {
	return _IErc20Registry.Contract.IsErc20StoreRegistered(&_IErc20Registry.CallOpts, account)
}

// IsErc20StoreRegistered is a free data retrieval call binding the contract method 0x4e25ab64.
//
// Solidity: function is_erc20_store_registered(address account) view returns(bool registered)
func (_IErc20Registry *IErc20RegistryCallerSession) IsErc20StoreRegistered(account common.Address) (bool, error) {
	return _IErc20Registry.Contract.IsErc20StoreRegistered(&_IErc20Registry.CallOpts, account)
}

// RegisterErc20 is a paid mutator transaction binding the contract method 0x5e6c5759.
//
// Solidity: function register_erc20() returns()
func (_IErc20Registry *IErc20RegistryTransactor) RegisterErc20(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IErc20Registry.contract.Transact(opts, "register_erc20")
}

// RegisterErc20 is a paid mutator transaction binding the contract method 0x5e6c5759.
//
// Solidity: function register_erc20() returns()
func (_IErc20Registry *IErc20RegistrySession) RegisterErc20() (*types.Transaction, error) {
	return _IErc20Registry.Contract.RegisterErc20(&_IErc20Registry.TransactOpts)
}

// RegisterErc20 is a paid mutator transaction binding the contract method 0x5e6c5759.
//
// Solidity: function register_erc20() returns()
func (_IErc20Registry *IErc20RegistryTransactorSession) RegisterErc20() (*types.Transaction, error) {
	return _IErc20Registry.Contract.RegisterErc20(&_IErc20Registry.TransactOpts)
}

// RegisterErc20Store is a paid mutator transaction binding the contract method 0xceeae52a.
//
// Solidity: function register_erc20_store(address account) returns()
func (_IErc20Registry *IErc20RegistryTransactor) RegisterErc20Store(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _IErc20Registry.contract.Transact(opts, "register_erc20_store", account)
}

// RegisterErc20Store is a paid mutator transaction binding the contract method 0xceeae52a.
//
// Solidity: function register_erc20_store(address account) returns()
func (_IErc20Registry *IErc20RegistrySession) RegisterErc20Store(account common.Address) (*types.Transaction, error) {
	return _IErc20Registry.Contract.RegisterErc20Store(&_IErc20Registry.TransactOpts, account)
}

// RegisterErc20Store is a paid mutator transaction binding the contract method 0xceeae52a.
//
// Solidity: function register_erc20_store(address account) returns()
func (_IErc20Registry *IErc20RegistryTransactorSession) RegisterErc20Store(account common.Address) (*types.Transaction, error) {
	return _IErc20Registry.Contract.RegisterErc20Store(&_IErc20Registry.TransactOpts, account)
}
