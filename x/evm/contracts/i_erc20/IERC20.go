// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package i_erc20

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

// IErc20MetaData contains all meta data concerning the IErc20 contract.
var IErc20MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IErc20ABI is the input ABI used to generate the binding from.
// Deprecated: Use IErc20MetaData.ABI instead.
var IErc20ABI = IErc20MetaData.ABI

// IErc20 is an auto generated Go binding around an Ethereum contract.
type IErc20 struct {
	IErc20Caller     // Read-only binding to the contract
	IErc20Transactor // Write-only binding to the contract
	IErc20Filterer   // Log filterer for contract events
}

// IErc20Caller is an auto generated read-only Go binding around an Ethereum contract.
type IErc20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IErc20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IErc20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IErc20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IErc20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IErc20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IErc20Session struct {
	Contract     *IErc20           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IErc20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IErc20CallerSession struct {
	Contract *IErc20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IErc20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IErc20TransactorSession struct {
	Contract     *IErc20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IErc20Raw is an auto generated low-level Go binding around an Ethereum contract.
type IErc20Raw struct {
	Contract *IErc20 // Generic contract binding to access the raw methods on
}

// IErc20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IErc20CallerRaw struct {
	Contract *IErc20Caller // Generic read-only contract binding to access the raw methods on
}

// IErc20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IErc20TransactorRaw struct {
	Contract *IErc20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIErc20 creates a new instance of IErc20, bound to a specific deployed contract.
func NewIErc20(address common.Address, backend bind.ContractBackend) (*IErc20, error) {
	contract, err := bindIErc20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IErc20{IErc20Caller: IErc20Caller{contract: contract}, IErc20Transactor: IErc20Transactor{contract: contract}, IErc20Filterer: IErc20Filterer{contract: contract}}, nil
}

// NewIErc20Caller creates a new read-only instance of IErc20, bound to a specific deployed contract.
func NewIErc20Caller(address common.Address, caller bind.ContractCaller) (*IErc20Caller, error) {
	contract, err := bindIErc20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IErc20Caller{contract: contract}, nil
}

// NewIErc20Transactor creates a new write-only instance of IErc20, bound to a specific deployed contract.
func NewIErc20Transactor(address common.Address, transactor bind.ContractTransactor) (*IErc20Transactor, error) {
	contract, err := bindIErc20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IErc20Transactor{contract: contract}, nil
}

// NewIErc20Filterer creates a new log filterer instance of IErc20, bound to a specific deployed contract.
func NewIErc20Filterer(address common.Address, filterer bind.ContractFilterer) (*IErc20Filterer, error) {
	contract, err := bindIErc20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IErc20Filterer{contract: contract}, nil
}

// bindIErc20 binds a generic wrapper to an already deployed contract.
func bindIErc20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IErc20MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IErc20 *IErc20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IErc20.Contract.IErc20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IErc20 *IErc20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IErc20.Contract.IErc20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IErc20 *IErc20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IErc20.Contract.IErc20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IErc20 *IErc20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IErc20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IErc20 *IErc20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IErc20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IErc20 *IErc20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IErc20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IErc20 *IErc20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IErc20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IErc20 *IErc20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IErc20.Contract.Allowance(&_IErc20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IErc20 *IErc20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IErc20.Contract.Allowance(&_IErc20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IErc20 *IErc20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IErc20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IErc20 *IErc20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _IErc20.Contract.BalanceOf(&_IErc20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IErc20 *IErc20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IErc20.Contract.BalanceOf(&_IErc20.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IErc20 *IErc20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _IErc20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IErc20 *IErc20Session) Decimals() (uint8, error) {
	return _IErc20.Contract.Decimals(&_IErc20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IErc20 *IErc20CallerSession) Decimals() (uint8, error) {
	return _IErc20.Contract.Decimals(&_IErc20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IErc20 *IErc20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IErc20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IErc20 *IErc20Session) Name() (string, error) {
	return _IErc20.Contract.Name(&_IErc20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IErc20 *IErc20CallerSession) Name() (string, error) {
	return _IErc20.Contract.Name(&_IErc20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IErc20 *IErc20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IErc20.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IErc20 *IErc20Session) Symbol() (string, error) {
	return _IErc20.Contract.Symbol(&_IErc20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IErc20 *IErc20CallerSession) Symbol() (string, error) {
	return _IErc20.Contract.Symbol(&_IErc20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IErc20 *IErc20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IErc20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IErc20 *IErc20Session) TotalSupply() (*big.Int, error) {
	return _IErc20.Contract.TotalSupply(&_IErc20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IErc20 *IErc20CallerSession) TotalSupply() (*big.Int, error) {
	return _IErc20.Contract.TotalSupply(&_IErc20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IErc20 *IErc20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IErc20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IErc20 *IErc20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IErc20.Contract.Approve(&_IErc20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IErc20 *IErc20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IErc20.Contract.Approve(&_IErc20.TransactOpts, spender, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IErc20 *IErc20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IErc20.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IErc20 *IErc20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IErc20.Contract.Transfer(&_IErc20.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IErc20 *IErc20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IErc20.Contract.Transfer(&_IErc20.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IErc20 *IErc20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IErc20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IErc20 *IErc20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IErc20.Contract.TransferFrom(&_IErc20.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IErc20 *IErc20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IErc20.Contract.TransferFrom(&_IErc20.TransactOpts, sender, recipient, amount)
}
