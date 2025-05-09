// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package create2

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

// Create2MetaData contains all meta data concerning the Create2 contract.
var Create2MetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x602c604b600b8282823980515f1a607314603f577f4e487b71000000000000000000000000000000000000000000000000000000005f525f60045260245ffd5b305f52607381538281f3fe730000000000000000000000000000000000000000301460806040525f80fdfea164736f6c6343000819000a",
}

// Create2ABI is the input ABI used to generate the binding from.
// Deprecated: Use Create2MetaData.ABI instead.
var Create2ABI = Create2MetaData.ABI

// Create2Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Create2MetaData.Bin instead.
var Create2Bin = Create2MetaData.Bin

// DeployCreate2 deploys a new Ethereum contract, binding an instance of Create2 to it.
func DeployCreate2(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Create2, error) {
	parsed, err := Create2MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Create2Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Create2{Create2Caller: Create2Caller{contract: contract}, Create2Transactor: Create2Transactor{contract: contract}, Create2Filterer: Create2Filterer{contract: contract}}, nil
}

// Create2 is an auto generated Go binding around an Ethereum contract.
type Create2 struct {
	Create2Caller     // Read-only binding to the contract
	Create2Transactor // Write-only binding to the contract
	Create2Filterer   // Log filterer for contract events
}

// Create2Caller is an auto generated read-only Go binding around an Ethereum contract.
type Create2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Create2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Create2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Create2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Create2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Create2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Create2Session struct {
	Contract     *Create2          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Create2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Create2CallerSession struct {
	Contract *Create2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// Create2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Create2TransactorSession struct {
	Contract     *Create2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// Create2Raw is an auto generated low-level Go binding around an Ethereum contract.
type Create2Raw struct {
	Contract *Create2 // Generic contract binding to access the raw methods on
}

// Create2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Create2CallerRaw struct {
	Contract *Create2Caller // Generic read-only contract binding to access the raw methods on
}

// Create2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Create2TransactorRaw struct {
	Contract *Create2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewCreate2 creates a new instance of Create2, bound to a specific deployed contract.
func NewCreate2(address common.Address, backend bind.ContractBackend) (*Create2, error) {
	contract, err := bindCreate2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Create2{Create2Caller: Create2Caller{contract: contract}, Create2Transactor: Create2Transactor{contract: contract}, Create2Filterer: Create2Filterer{contract: contract}}, nil
}

// NewCreate2Caller creates a new read-only instance of Create2, bound to a specific deployed contract.
func NewCreate2Caller(address common.Address, caller bind.ContractCaller) (*Create2Caller, error) {
	contract, err := bindCreate2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Create2Caller{contract: contract}, nil
}

// NewCreate2Transactor creates a new write-only instance of Create2, bound to a specific deployed contract.
func NewCreate2Transactor(address common.Address, transactor bind.ContractTransactor) (*Create2Transactor, error) {
	contract, err := bindCreate2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Create2Transactor{contract: contract}, nil
}

// NewCreate2Filterer creates a new log filterer instance of Create2, bound to a specific deployed contract.
func NewCreate2Filterer(address common.Address, filterer bind.ContractFilterer) (*Create2Filterer, error) {
	contract, err := bindCreate2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Create2Filterer{contract: contract}, nil
}

// bindCreate2 binds a generic wrapper to an already deployed contract.
func bindCreate2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Create2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Create2 *Create2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Create2.Contract.Create2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Create2 *Create2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Create2.Contract.Create2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Create2 *Create2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Create2.Contract.Create2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Create2 *Create2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Create2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Create2 *Create2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Create2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Create2 *Create2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Create2.Contract.contract.Transact(opts, method, params...)
}
