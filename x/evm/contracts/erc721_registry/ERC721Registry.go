// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package erc721_registry

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

// Erc721RegistryMetaData contains all meta data concerning the Erc721Registry contract.
var Erc721RegistryMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x6080604052348015600e575f80fd5b50603e80601a5f395ff3fe60806040525f80fdfea26469706673582212201fb579f749f166df13d0c2e4d5cf364bde2dd0d1f877ee7e0ef986d2d5e26e1564736f6c63430008190033",
}

// Erc721RegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use Erc721RegistryMetaData.ABI instead.
var Erc721RegistryABI = Erc721RegistryMetaData.ABI

// Erc721RegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Erc721RegistryMetaData.Bin instead.
var Erc721RegistryBin = Erc721RegistryMetaData.Bin

// DeployErc721Registry deploys a new Ethereum contract, binding an instance of Erc721Registry to it.
func DeployErc721Registry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Erc721Registry, error) {
	parsed, err := Erc721RegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Erc721RegistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Erc721Registry{Erc721RegistryCaller: Erc721RegistryCaller{contract: contract}, Erc721RegistryTransactor: Erc721RegistryTransactor{contract: contract}, Erc721RegistryFilterer: Erc721RegistryFilterer{contract: contract}}, nil
}

// Erc721Registry is an auto generated Go binding around an Ethereum contract.
type Erc721Registry struct {
	Erc721RegistryCaller     // Read-only binding to the contract
	Erc721RegistryTransactor // Write-only binding to the contract
	Erc721RegistryFilterer   // Log filterer for contract events
}

// Erc721RegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type Erc721RegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc721RegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Erc721RegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc721RegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Erc721RegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc721RegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Erc721RegistrySession struct {
	Contract     *Erc721Registry   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Erc721RegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Erc721RegistryCallerSession struct {
	Contract *Erc721RegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// Erc721RegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Erc721RegistryTransactorSession struct {
	Contract     *Erc721RegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// Erc721RegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type Erc721RegistryRaw struct {
	Contract *Erc721Registry // Generic contract binding to access the raw methods on
}

// Erc721RegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Erc721RegistryCallerRaw struct {
	Contract *Erc721RegistryCaller // Generic read-only contract binding to access the raw methods on
}

// Erc721RegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Erc721RegistryTransactorRaw struct {
	Contract *Erc721RegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewErc721Registry creates a new instance of Erc721Registry, bound to a specific deployed contract.
func NewErc721Registry(address common.Address, backend bind.ContractBackend) (*Erc721Registry, error) {
	contract, err := bindErc721Registry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Erc721Registry{Erc721RegistryCaller: Erc721RegistryCaller{contract: contract}, Erc721RegistryTransactor: Erc721RegistryTransactor{contract: contract}, Erc721RegistryFilterer: Erc721RegistryFilterer{contract: contract}}, nil
}

// NewErc721RegistryCaller creates a new read-only instance of Erc721Registry, bound to a specific deployed contract.
func NewErc721RegistryCaller(address common.Address, caller bind.ContractCaller) (*Erc721RegistryCaller, error) {
	contract, err := bindErc721Registry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Erc721RegistryCaller{contract: contract}, nil
}

// NewErc721RegistryTransactor creates a new write-only instance of Erc721Registry, bound to a specific deployed contract.
func NewErc721RegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*Erc721RegistryTransactor, error) {
	contract, err := bindErc721Registry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Erc721RegistryTransactor{contract: contract}, nil
}

// NewErc721RegistryFilterer creates a new log filterer instance of Erc721Registry, bound to a specific deployed contract.
func NewErc721RegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*Erc721RegistryFilterer, error) {
	contract, err := bindErc721Registry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Erc721RegistryFilterer{contract: contract}, nil
}

// bindErc721Registry binds a generic wrapper to an already deployed contract.
func bindErc721Registry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Erc721RegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc721Registry *Erc721RegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc721Registry.Contract.Erc721RegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc721Registry *Erc721RegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc721Registry.Contract.Erc721RegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc721Registry *Erc721RegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc721Registry.Contract.Erc721RegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc721Registry *Erc721RegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc721Registry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc721Registry *Erc721RegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc721Registry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc721Registry *Erc721RegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc721Registry.Contract.contract.Transact(opts, method, params...)
}
