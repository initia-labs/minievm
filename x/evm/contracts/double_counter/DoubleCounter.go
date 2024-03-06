// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package double_counter

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

// DoubleCounterMetaData contains all meta data concerning the DoubleCounter contract.
var DoubleCounterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"count\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"increase\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b506101468061001d5f395ff3fe608060405234801561000f575f80fd5b5060043610610034575f3560e01c806306661abd14610038578063e8927fbc14610056575b5f80fd5b610040610060565b60405161004d9190610097565b60405180910390f35b61005e610065565b005b5f5481565b60025f8082825461007691906100dd565b92505081905550565b5f819050919050565b6100918161007f565b82525050565b5f6020820190506100aa5f830184610088565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6100e78261007f565b91506100f28361007f565b925082820190508082111561010a576101096100b0565b5b9291505056fea26469706673582212207f99c62aef3d87ded350e2034fe1971ea63985a0b8ab02c2d1a51417f74f660864736f6c63430008180033",
}

// DoubleCounterABI is the input ABI used to generate the binding from.
// Deprecated: Use DoubleCounterMetaData.ABI instead.
var DoubleCounterABI = DoubleCounterMetaData.ABI

// DoubleCounterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DoubleCounterMetaData.Bin instead.
var DoubleCounterBin = DoubleCounterMetaData.Bin

// DeployDoubleCounter deploys a new Ethereum contract, binding an instance of DoubleCounter to it.
func DeployDoubleCounter(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DoubleCounter, error) {
	parsed, err := DoubleCounterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DoubleCounterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DoubleCounter{DoubleCounterCaller: DoubleCounterCaller{contract: contract}, DoubleCounterTransactor: DoubleCounterTransactor{contract: contract}, DoubleCounterFilterer: DoubleCounterFilterer{contract: contract}}, nil
}

// DoubleCounter is an auto generated Go binding around an Ethereum contract.
type DoubleCounter struct {
	DoubleCounterCaller     // Read-only binding to the contract
	DoubleCounterTransactor // Write-only binding to the contract
	DoubleCounterFilterer   // Log filterer for contract events
}

// DoubleCounterCaller is an auto generated read-only Go binding around an Ethereum contract.
type DoubleCounterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DoubleCounterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DoubleCounterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DoubleCounterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DoubleCounterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DoubleCounterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DoubleCounterSession struct {
	Contract     *DoubleCounter    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DoubleCounterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DoubleCounterCallerSession struct {
	Contract *DoubleCounterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// DoubleCounterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DoubleCounterTransactorSession struct {
	Contract     *DoubleCounterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// DoubleCounterRaw is an auto generated low-level Go binding around an Ethereum contract.
type DoubleCounterRaw struct {
	Contract *DoubleCounter // Generic contract binding to access the raw methods on
}

// DoubleCounterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DoubleCounterCallerRaw struct {
	Contract *DoubleCounterCaller // Generic read-only contract binding to access the raw methods on
}

// DoubleCounterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DoubleCounterTransactorRaw struct {
	Contract *DoubleCounterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDoubleCounter creates a new instance of DoubleCounter, bound to a specific deployed contract.
func NewDoubleCounter(address common.Address, backend bind.ContractBackend) (*DoubleCounter, error) {
	contract, err := bindDoubleCounter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DoubleCounter{DoubleCounterCaller: DoubleCounterCaller{contract: contract}, DoubleCounterTransactor: DoubleCounterTransactor{contract: contract}, DoubleCounterFilterer: DoubleCounterFilterer{contract: contract}}, nil
}

// NewDoubleCounterCaller creates a new read-only instance of DoubleCounter, bound to a specific deployed contract.
func NewDoubleCounterCaller(address common.Address, caller bind.ContractCaller) (*DoubleCounterCaller, error) {
	contract, err := bindDoubleCounter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DoubleCounterCaller{contract: contract}, nil
}

// NewDoubleCounterTransactor creates a new write-only instance of DoubleCounter, bound to a specific deployed contract.
func NewDoubleCounterTransactor(address common.Address, transactor bind.ContractTransactor) (*DoubleCounterTransactor, error) {
	contract, err := bindDoubleCounter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DoubleCounterTransactor{contract: contract}, nil
}

// NewDoubleCounterFilterer creates a new log filterer instance of DoubleCounter, bound to a specific deployed contract.
func NewDoubleCounterFilterer(address common.Address, filterer bind.ContractFilterer) (*DoubleCounterFilterer, error) {
	contract, err := bindDoubleCounter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DoubleCounterFilterer{contract: contract}, nil
}

// bindDoubleCounter binds a generic wrapper to an already deployed contract.
func bindDoubleCounter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DoubleCounterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DoubleCounter *DoubleCounterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DoubleCounter.Contract.DoubleCounterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DoubleCounter *DoubleCounterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DoubleCounter.Contract.DoubleCounterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DoubleCounter *DoubleCounterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DoubleCounter.Contract.DoubleCounterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DoubleCounter *DoubleCounterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DoubleCounter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DoubleCounter *DoubleCounterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DoubleCounter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DoubleCounter *DoubleCounterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DoubleCounter.Contract.contract.Transact(opts, method, params...)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_DoubleCounter *DoubleCounterCaller) Count(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DoubleCounter.contract.Call(opts, &out, "count")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_DoubleCounter *DoubleCounterSession) Count() (*big.Int, error) {
	return _DoubleCounter.Contract.Count(&_DoubleCounter.CallOpts)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_DoubleCounter *DoubleCounterCallerSession) Count() (*big.Int, error) {
	return _DoubleCounter.Contract.Count(&_DoubleCounter.CallOpts)
}

// Increase is a paid mutator transaction binding the contract method 0xe8927fbc.
//
// Solidity: function increase() returns()
func (_DoubleCounter *DoubleCounterTransactor) Increase(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DoubleCounter.contract.Transact(opts, "increase")
}

// Increase is a paid mutator transaction binding the contract method 0xe8927fbc.
//
// Solidity: function increase() returns()
func (_DoubleCounter *DoubleCounterSession) Increase() (*types.Transaction, error) {
	return _DoubleCounter.Contract.Increase(&_DoubleCounter.TransactOpts)
}

// Increase is a paid mutator transaction binding the contract method 0xe8927fbc.
//
// Solidity: function increase() returns()
func (_DoubleCounter *DoubleCounterTransactorSession) Increase() (*types.Transaction, error) {
	return _DoubleCounter.Contract.Increase(&_DoubleCounter.TransactOpts)
}
