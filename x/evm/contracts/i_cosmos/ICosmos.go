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

// ICosmosOptions is an auto generated low-level Go binding around an user-defined struct.
type ICosmosOptions struct {
	AllowFailure bool
	CallbackId   uint64
}

// ICosmosMetaData contains all meta data concerning the ICosmos contract.
var ICosmosMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"disable_execute_cosmos\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"dummy\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"msg\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"gas_limit\",\"type\":\"uint64\"}],\"name\":\"execute_cosmos\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"dummy\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"msg\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"gas_limit\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"allow_failure\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"callback_id\",\"type\":\"uint64\"}],\"internalType\":\"structICosmos.Options\",\"name\":\"options\",\"type\":\"tuple\"}],\"name\":\"execute_cosmos_with_options\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"dummy\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"is_authority_address\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"authority\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"is_blocked_address\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"blocked\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"is_module_address\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"module\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"path\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"req\",\"type\":\"string\"}],\"name\":\"query_cosmos\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"result\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"evm_address\",\"type\":\"address\"}],\"name\":\"to_cosmos_address\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"cosmos_address\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc20_address\",\"type\":\"address\"}],\"name\":\"to_denom\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"}],\"name\":\"to_erc20\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"erc20_address\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"cosmos_address\",\"type\":\"string\"}],\"name\":\"to_evm_address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"evm_address\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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

// IsAuthorityAddress is a free data retrieval call binding the contract method 0x01116078.
//
// Solidity: function is_authority_address(address account) view returns(bool authority)
func (_ICosmos *ICosmosCaller) IsAuthorityAddress(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _ICosmos.contract.Call(opts, &out, "is_authority_address", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAuthorityAddress is a free data retrieval call binding the contract method 0x01116078.
//
// Solidity: function is_authority_address(address account) view returns(bool authority)
func (_ICosmos *ICosmosSession) IsAuthorityAddress(account common.Address) (bool, error) {
	return _ICosmos.Contract.IsAuthorityAddress(&_ICosmos.CallOpts, account)
}

// IsAuthorityAddress is a free data retrieval call binding the contract method 0x01116078.
//
// Solidity: function is_authority_address(address account) view returns(bool authority)
func (_ICosmos *ICosmosCallerSession) IsAuthorityAddress(account common.Address) (bool, error) {
	return _ICosmos.Contract.IsAuthorityAddress(&_ICosmos.CallOpts, account)
}

// IsBlockedAddress is a free data retrieval call binding the contract method 0xf2af9ac9.
//
// Solidity: function is_blocked_address(address account) view returns(bool blocked)
func (_ICosmos *ICosmosCaller) IsBlockedAddress(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _ICosmos.contract.Call(opts, &out, "is_blocked_address", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBlockedAddress is a free data retrieval call binding the contract method 0xf2af9ac9.
//
// Solidity: function is_blocked_address(address account) view returns(bool blocked)
func (_ICosmos *ICosmosSession) IsBlockedAddress(account common.Address) (bool, error) {
	return _ICosmos.Contract.IsBlockedAddress(&_ICosmos.CallOpts, account)
}

// IsBlockedAddress is a free data retrieval call binding the contract method 0xf2af9ac9.
//
// Solidity: function is_blocked_address(address account) view returns(bool blocked)
func (_ICosmos *ICosmosCallerSession) IsBlockedAddress(account common.Address) (bool, error) {
	return _ICosmos.Contract.IsBlockedAddress(&_ICosmos.CallOpts, account)
}

// IsModuleAddress is a free data retrieval call binding the contract method 0x60dc402f.
//
// Solidity: function is_module_address(address account) view returns(bool module)
func (_ICosmos *ICosmosCaller) IsModuleAddress(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _ICosmos.contract.Call(opts, &out, "is_module_address", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsModuleAddress is a free data retrieval call binding the contract method 0x60dc402f.
//
// Solidity: function is_module_address(address account) view returns(bool module)
func (_ICosmos *ICosmosSession) IsModuleAddress(account common.Address) (bool, error) {
	return _ICosmos.Contract.IsModuleAddress(&_ICosmos.CallOpts, account)
}

// IsModuleAddress is a free data retrieval call binding the contract method 0x60dc402f.
//
// Solidity: function is_module_address(address account) view returns(bool module)
func (_ICosmos *ICosmosCallerSession) IsModuleAddress(account common.Address) (bool, error) {
	return _ICosmos.Contract.IsModuleAddress(&_ICosmos.CallOpts, account)
}

// QueryCosmos is a free data retrieval call binding the contract method 0xcad23554.
//
// Solidity: function query_cosmos(string path, string req) view returns(string result)
func (_ICosmos *ICosmosCaller) QueryCosmos(opts *bind.CallOpts, path string, req string) (string, error) {
	var out []interface{}
	err := _ICosmos.contract.Call(opts, &out, "query_cosmos", path, req)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// QueryCosmos is a free data retrieval call binding the contract method 0xcad23554.
//
// Solidity: function query_cosmos(string path, string req) view returns(string result)
func (_ICosmos *ICosmosSession) QueryCosmos(path string, req string) (string, error) {
	return _ICosmos.Contract.QueryCosmos(&_ICosmos.CallOpts, path, req)
}

// QueryCosmos is a free data retrieval call binding the contract method 0xcad23554.
//
// Solidity: function query_cosmos(string path, string req) view returns(string result)
func (_ICosmos *ICosmosCallerSession) QueryCosmos(path string, req string) (string, error) {
	return _ICosmos.Contract.QueryCosmos(&_ICosmos.CallOpts, path, req)
}

// ToCosmosAddress is a free data retrieval call binding the contract method 0x6af32a55.
//
// Solidity: function to_cosmos_address(address evm_address) view returns(string cosmos_address)
func (_ICosmos *ICosmosCaller) ToCosmosAddress(opts *bind.CallOpts, evm_address common.Address) (string, error) {
	var out []interface{}
	err := _ICosmos.contract.Call(opts, &out, "to_cosmos_address", evm_address)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ToCosmosAddress is a free data retrieval call binding the contract method 0x6af32a55.
//
// Solidity: function to_cosmos_address(address evm_address) view returns(string cosmos_address)
func (_ICosmos *ICosmosSession) ToCosmosAddress(evm_address common.Address) (string, error) {
	return _ICosmos.Contract.ToCosmosAddress(&_ICosmos.CallOpts, evm_address)
}

// ToCosmosAddress is a free data retrieval call binding the contract method 0x6af32a55.
//
// Solidity: function to_cosmos_address(address evm_address) view returns(string cosmos_address)
func (_ICosmos *ICosmosCallerSession) ToCosmosAddress(evm_address common.Address) (string, error) {
	return _ICosmos.Contract.ToCosmosAddress(&_ICosmos.CallOpts, evm_address)
}

// ToDenom is a free data retrieval call binding the contract method 0x81cf0f6a.
//
// Solidity: function to_denom(address erc20_address) view returns(string denom)
func (_ICosmos *ICosmosCaller) ToDenom(opts *bind.CallOpts, erc20_address common.Address) (string, error) {
	var out []interface{}
	err := _ICosmos.contract.Call(opts, &out, "to_denom", erc20_address)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ToDenom is a free data retrieval call binding the contract method 0x81cf0f6a.
//
// Solidity: function to_denom(address erc20_address) view returns(string denom)
func (_ICosmos *ICosmosSession) ToDenom(erc20_address common.Address) (string, error) {
	return _ICosmos.Contract.ToDenom(&_ICosmos.CallOpts, erc20_address)
}

// ToDenom is a free data retrieval call binding the contract method 0x81cf0f6a.
//
// Solidity: function to_denom(address erc20_address) view returns(string denom)
func (_ICosmos *ICosmosCallerSession) ToDenom(erc20_address common.Address) (string, error) {
	return _ICosmos.Contract.ToDenom(&_ICosmos.CallOpts, erc20_address)
}

// ToErc20 is a free data retrieval call binding the contract method 0x2b3324ce.
//
// Solidity: function to_erc20(string denom) view returns(address erc20_address)
func (_ICosmos *ICosmosCaller) ToErc20(opts *bind.CallOpts, denom string) (common.Address, error) {
	var out []interface{}
	err := _ICosmos.contract.Call(opts, &out, "to_erc20", denom)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ToErc20 is a free data retrieval call binding the contract method 0x2b3324ce.
//
// Solidity: function to_erc20(string denom) view returns(address erc20_address)
func (_ICosmos *ICosmosSession) ToErc20(denom string) (common.Address, error) {
	return _ICosmos.Contract.ToErc20(&_ICosmos.CallOpts, denom)
}

// ToErc20 is a free data retrieval call binding the contract method 0x2b3324ce.
//
// Solidity: function to_erc20(string denom) view returns(address erc20_address)
func (_ICosmos *ICosmosCallerSession) ToErc20(denom string) (common.Address, error) {
	return _ICosmos.Contract.ToErc20(&_ICosmos.CallOpts, denom)
}

// ToEvmAddress is a free data retrieval call binding the contract method 0x4f94a155.
//
// Solidity: function to_evm_address(string cosmos_address) view returns(address evm_address)
func (_ICosmos *ICosmosCaller) ToEvmAddress(opts *bind.CallOpts, cosmos_address string) (common.Address, error) {
	var out []interface{}
	err := _ICosmos.contract.Call(opts, &out, "to_evm_address", cosmos_address)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ToEvmAddress is a free data retrieval call binding the contract method 0x4f94a155.
//
// Solidity: function to_evm_address(string cosmos_address) view returns(address evm_address)
func (_ICosmos *ICosmosSession) ToEvmAddress(cosmos_address string) (common.Address, error) {
	return _ICosmos.Contract.ToEvmAddress(&_ICosmos.CallOpts, cosmos_address)
}

// ToEvmAddress is a free data retrieval call binding the contract method 0x4f94a155.
//
// Solidity: function to_evm_address(string cosmos_address) view returns(address evm_address)
func (_ICosmos *ICosmosCallerSession) ToEvmAddress(cosmos_address string) (common.Address, error) {
	return _ICosmos.Contract.ToEvmAddress(&_ICosmos.CallOpts, cosmos_address)
}

// DisableExecuteCosmos is a paid mutator transaction binding the contract method 0x8c1370cd.
//
// Solidity: function disable_execute_cosmos() returns(bool dummy)
func (_ICosmos *ICosmosTransactor) DisableExecuteCosmos(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICosmos.contract.Transact(opts, "disable_execute_cosmos")
}

// DisableExecuteCosmos is a paid mutator transaction binding the contract method 0x8c1370cd.
//
// Solidity: function disable_execute_cosmos() returns(bool dummy)
func (_ICosmos *ICosmosSession) DisableExecuteCosmos() (*types.Transaction, error) {
	return _ICosmos.Contract.DisableExecuteCosmos(&_ICosmos.TransactOpts)
}

// DisableExecuteCosmos is a paid mutator transaction binding the contract method 0x8c1370cd.
//
// Solidity: function disable_execute_cosmos() returns(bool dummy)
func (_ICosmos *ICosmosTransactorSession) DisableExecuteCosmos() (*types.Transaction, error) {
	return _ICosmos.Contract.DisableExecuteCosmos(&_ICosmos.TransactOpts)
}

// ExecuteCosmos is a paid mutator transaction binding the contract method 0x56c657a5.
//
// Solidity: function execute_cosmos(string msg, uint64 gas_limit) returns(bool dummy)
func (_ICosmos *ICosmosTransactor) ExecuteCosmos(opts *bind.TransactOpts, msg string, gas_limit uint64) (*types.Transaction, error) {
	return _ICosmos.contract.Transact(opts, "execute_cosmos", msg, gas_limit)
}

// ExecuteCosmos is a paid mutator transaction binding the contract method 0x56c657a5.
//
// Solidity: function execute_cosmos(string msg, uint64 gas_limit) returns(bool dummy)
func (_ICosmos *ICosmosSession) ExecuteCosmos(msg string, gas_limit uint64) (*types.Transaction, error) {
	return _ICosmos.Contract.ExecuteCosmos(&_ICosmos.TransactOpts, msg, gas_limit)
}

// ExecuteCosmos is a paid mutator transaction binding the contract method 0x56c657a5.
//
// Solidity: function execute_cosmos(string msg, uint64 gas_limit) returns(bool dummy)
func (_ICosmos *ICosmosTransactorSession) ExecuteCosmos(msg string, gas_limit uint64) (*types.Transaction, error) {
	return _ICosmos.Contract.ExecuteCosmos(&_ICosmos.TransactOpts, msg, gas_limit)
}

// ExecuteCosmosWithOptions is a paid mutator transaction binding the contract method 0xf1ed795d.
//
// Solidity: function execute_cosmos_with_options(string msg, uint64 gas_limit, (bool,uint64) options) returns(bool dummy)
func (_ICosmos *ICosmosTransactor) ExecuteCosmosWithOptions(opts *bind.TransactOpts, msg string, gas_limit uint64, options ICosmosOptions) (*types.Transaction, error) {
	return _ICosmos.contract.Transact(opts, "execute_cosmos_with_options", msg, gas_limit, options)
}

// ExecuteCosmosWithOptions is a paid mutator transaction binding the contract method 0xf1ed795d.
//
// Solidity: function execute_cosmos_with_options(string msg, uint64 gas_limit, (bool,uint64) options) returns(bool dummy)
func (_ICosmos *ICosmosSession) ExecuteCosmosWithOptions(msg string, gas_limit uint64, options ICosmosOptions) (*types.Transaction, error) {
	return _ICosmos.Contract.ExecuteCosmosWithOptions(&_ICosmos.TransactOpts, msg, gas_limit, options)
}

// ExecuteCosmosWithOptions is a paid mutator transaction binding the contract method 0xf1ed795d.
//
// Solidity: function execute_cosmos_with_options(string msg, uint64 gas_limit, (bool,uint64) options) returns(bool dummy)
func (_ICosmos *ICosmosTransactorSession) ExecuteCosmosWithOptions(msg string, gas_limit uint64, options ICosmosOptions) (*types.Transaction, error) {
	return _ICosmos.Contract.ExecuteCosmosWithOptions(&_ICosmos.TransactOpts, msg, gas_limit, options)
}
