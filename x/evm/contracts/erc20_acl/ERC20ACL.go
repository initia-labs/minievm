// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package erc20_acl

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

// Erc20AclMetaData contains all meta data concerning the Erc20Acl contract.
var Erc20AclMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x6080604052348015600e575f80fd5b50603e80601a5f395ff3fe60806040525f80fdfea264697066735822122097f44a6309abb239e19ee1005989fb200d740afc2223e4d6c0fb82a45f806dec64736f6c63430008180033",
}

// Erc20AclABI is the input ABI used to generate the binding from.
// Deprecated: Use Erc20AclMetaData.ABI instead.
var Erc20AclABI = Erc20AclMetaData.ABI

// Erc20AclBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Erc20AclMetaData.Bin instead.
var Erc20AclBin = Erc20AclMetaData.Bin

// DeployErc20Acl deploys a new Ethereum contract, binding an instance of Erc20Acl to it.
func DeployErc20Acl(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Erc20Acl, error) {
	parsed, err := Erc20AclMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Erc20AclBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Erc20Acl{Erc20AclCaller: Erc20AclCaller{contract: contract}, Erc20AclTransactor: Erc20AclTransactor{contract: contract}, Erc20AclFilterer: Erc20AclFilterer{contract: contract}}, nil
}

// Erc20Acl is an auto generated Go binding around an Ethereum contract.
type Erc20Acl struct {
	Erc20AclCaller     // Read-only binding to the contract
	Erc20AclTransactor // Write-only binding to the contract
	Erc20AclFilterer   // Log filterer for contract events
}

// Erc20AclCaller is an auto generated read-only Go binding around an Ethereum contract.
type Erc20AclCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20AclTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Erc20AclTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20AclFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Erc20AclFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20AclSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Erc20AclSession struct {
	Contract     *Erc20Acl         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Erc20AclCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Erc20AclCallerSession struct {
	Contract *Erc20AclCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// Erc20AclTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Erc20AclTransactorSession struct {
	Contract     *Erc20AclTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// Erc20AclRaw is an auto generated low-level Go binding around an Ethereum contract.
type Erc20AclRaw struct {
	Contract *Erc20Acl // Generic contract binding to access the raw methods on
}

// Erc20AclCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Erc20AclCallerRaw struct {
	Contract *Erc20AclCaller // Generic read-only contract binding to access the raw methods on
}

// Erc20AclTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Erc20AclTransactorRaw struct {
	Contract *Erc20AclTransactor // Generic write-only contract binding to access the raw methods on
}

// NewErc20Acl creates a new instance of Erc20Acl, bound to a specific deployed contract.
func NewErc20Acl(address common.Address, backend bind.ContractBackend) (*Erc20Acl, error) {
	contract, err := bindErc20Acl(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Erc20Acl{Erc20AclCaller: Erc20AclCaller{contract: contract}, Erc20AclTransactor: Erc20AclTransactor{contract: contract}, Erc20AclFilterer: Erc20AclFilterer{contract: contract}}, nil
}

// NewErc20AclCaller creates a new read-only instance of Erc20Acl, bound to a specific deployed contract.
func NewErc20AclCaller(address common.Address, caller bind.ContractCaller) (*Erc20AclCaller, error) {
	contract, err := bindErc20Acl(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Erc20AclCaller{contract: contract}, nil
}

// NewErc20AclTransactor creates a new write-only instance of Erc20Acl, bound to a specific deployed contract.
func NewErc20AclTransactor(address common.Address, transactor bind.ContractTransactor) (*Erc20AclTransactor, error) {
	contract, err := bindErc20Acl(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Erc20AclTransactor{contract: contract}, nil
}

// NewErc20AclFilterer creates a new log filterer instance of Erc20Acl, bound to a specific deployed contract.
func NewErc20AclFilterer(address common.Address, filterer bind.ContractFilterer) (*Erc20AclFilterer, error) {
	contract, err := bindErc20Acl(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Erc20AclFilterer{contract: contract}, nil
}

// bindErc20Acl binds a generic wrapper to an already deployed contract.
func bindErc20Acl(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Erc20AclMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc20Acl *Erc20AclRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc20Acl.Contract.Erc20AclCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc20Acl *Erc20AclRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc20Acl.Contract.Erc20AclTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc20Acl *Erc20AclRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc20Acl.Contract.Erc20AclTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc20Acl *Erc20AclCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc20Acl.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc20Acl *Erc20AclTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc20Acl.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc20Acl *Erc20AclTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc20Acl.Contract.contract.Transact(opts, method, params...)
}
