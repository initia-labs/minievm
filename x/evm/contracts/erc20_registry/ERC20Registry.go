// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package erc20_registry

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

// Erc20RegistryMetaData contains all meta data concerning the Erc20Registry contract.
var Erc20RegistryMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x6080604052348015600e575f80fd5b50603e80601a5f395ff3fe60806040525f80fdfea264697066735822122030aa80545713f6ecc82346bb6aa0d43d3e45ae18fdeceb8d6e53b211854fb54064736f6c63430008180033",
}

// Erc20RegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use Erc20RegistryMetaData.ABI instead.
var Erc20RegistryABI = Erc20RegistryMetaData.ABI

// Erc20RegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Erc20RegistryMetaData.Bin instead.
var Erc20RegistryBin = Erc20RegistryMetaData.Bin

// DeployErc20Registry deploys a new Ethereum contract, binding an instance of Erc20Registry to it.
func DeployErc20Registry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Erc20Registry, error) {
	parsed, err := Erc20RegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Erc20RegistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Erc20Registry{Erc20RegistryCaller: Erc20RegistryCaller{contract: contract}, Erc20RegistryTransactor: Erc20RegistryTransactor{contract: contract}, Erc20RegistryFilterer: Erc20RegistryFilterer{contract: contract}}, nil
}

// Erc20Registry is an auto generated Go binding around an Ethereum contract.
type Erc20Registry struct {
	Erc20RegistryCaller     // Read-only binding to the contract
	Erc20RegistryTransactor // Write-only binding to the contract
	Erc20RegistryFilterer   // Log filterer for contract events
}

// Erc20RegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type Erc20RegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20RegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Erc20RegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20RegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Erc20RegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20RegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Erc20RegistrySession struct {
	Contract     *Erc20Registry    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Erc20RegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Erc20RegistryCallerSession struct {
	Contract *Erc20RegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// Erc20RegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Erc20RegistryTransactorSession struct {
	Contract     *Erc20RegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// Erc20RegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type Erc20RegistryRaw struct {
	Contract *Erc20Registry // Generic contract binding to access the raw methods on
}

// Erc20RegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Erc20RegistryCallerRaw struct {
	Contract *Erc20RegistryCaller // Generic read-only contract binding to access the raw methods on
}

// Erc20RegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Erc20RegistryTransactorRaw struct {
	Contract *Erc20RegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewErc20Registry creates a new instance of Erc20Registry, bound to a specific deployed contract.
func NewErc20Registry(address common.Address, backend bind.ContractBackend) (*Erc20Registry, error) {
	contract, err := bindErc20Registry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Erc20Registry{Erc20RegistryCaller: Erc20RegistryCaller{contract: contract}, Erc20RegistryTransactor: Erc20RegistryTransactor{contract: contract}, Erc20RegistryFilterer: Erc20RegistryFilterer{contract: contract}}, nil
}

// NewErc20RegistryCaller creates a new read-only instance of Erc20Registry, bound to a specific deployed contract.
func NewErc20RegistryCaller(address common.Address, caller bind.ContractCaller) (*Erc20RegistryCaller, error) {
	contract, err := bindErc20Registry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Erc20RegistryCaller{contract: contract}, nil
}

// NewErc20RegistryTransactor creates a new write-only instance of Erc20Registry, bound to a specific deployed contract.
func NewErc20RegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*Erc20RegistryTransactor, error) {
	contract, err := bindErc20Registry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Erc20RegistryTransactor{contract: contract}, nil
}

// NewErc20RegistryFilterer creates a new log filterer instance of Erc20Registry, bound to a specific deployed contract.
func NewErc20RegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*Erc20RegistryFilterer, error) {
	contract, err := bindErc20Registry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Erc20RegistryFilterer{contract: contract}, nil
}

// bindErc20Registry binds a generic wrapper to an already deployed contract.
func bindErc20Registry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Erc20RegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc20Registry *Erc20RegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc20Registry.Contract.Erc20RegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc20Registry *Erc20RegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc20Registry.Contract.Erc20RegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc20Registry *Erc20RegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc20Registry.Contract.Erc20RegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc20Registry *Erc20RegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc20Registry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc20Registry *Erc20RegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc20Registry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc20Registry *Erc20RegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc20Registry.Contract.contract.Transact(opts, method, params...)
}
