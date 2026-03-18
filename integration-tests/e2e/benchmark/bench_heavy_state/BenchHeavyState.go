// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bench_heavy_state

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

// BenchHeavyStateMetaData contains all meta data concerning the BenchHeavyState contract.
var BenchHeavyStateMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"localState\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"senderNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sharedState\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalCalls\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sharedCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"localCount\",\"type\":\"uint256\"}],\"name\":\"writeMixed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b506105988061001c5f395ff3fe608060405234801561000f575f5ffd5b5060043610610055575f3560e01c80633af3f24f146100595780636712fc83146100775780637746c1d2146100a75780639c90b443146100c3578063fb0756bb146100f3575b5f5ffd5b610061610123565b60405161006e9190610307565b60405180910390f35b610091600480360381019061008c919061034e565b610129565b60405161009e9190610307565b60405180910390f35b6100c160048036038101906100bc9190610379565b61013d565b005b6100dd60048036038101906100d89190610411565b6102ba565b6040516100ea9190610307565b60405180910390f35b61010d6004803603810190610108919061043c565b6102cf565b60405161011a9190610307565b60405180910390f35b60035481565b5f602052805f5260405f205f915090505481565b5f60025f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054905060018161018b91906104a7565b60025f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055505f83826101d991906104da565b90505f5f90505b8481101561021557435f5f83856101f791906104a7565b81526020019081526020015f208190555080806001019150506101e0565b505f838361022391906104da565b90505f5f90505b8481101561029b574360015f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f838561027d91906104a7565b81526020019081526020015f2081905550808060010191505061022a565b5060035f8154809291906102ae9061051b565b91905055505050505050565b6002602052805f5260405f205f915090505481565b6001602052815f5260405f20602052805f5260405f205f91509150505481565b5f819050919050565b610301816102ef565b82525050565b5f60208201905061031a5f8301846102f8565b92915050565b5f5ffd5b61032d816102ef565b8114610337575f5ffd5b50565b5f8135905061034881610324565b92915050565b5f6020828403121561036357610362610320565b5b5f6103708482850161033a565b91505092915050565b5f5f6040838503121561038f5761038e610320565b5b5f61039c8582860161033a565b92505060206103ad8582860161033a565b9150509250929050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6103e0826103b7565b9050919050565b6103f0816103d6565b81146103fa575f5ffd5b50565b5f8135905061040b816103e7565b92915050565b5f6020828403121561042657610425610320565b5b5f610433848285016103fd565b91505092915050565b5f5f6040838503121561045257610451610320565b5b5f61045f858286016103fd565b92505060206104708582860161033a565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6104b1826102ef565b91506104bc836102ef565b92508282019050808211156104d4576104d361047a565b5b92915050565b5f6104e4826102ef565b91506104ef836102ef565b92508282026104fd816102ef565b915082820484148315176105145761051361047a565b5b5092915050565b5f610525826102ef565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036105575761055661047a565b5b60018201905091905056fea264697066735822122078d1444998d26840245c12f320170319e8cd0de5f8638a5431cbac7ffb10033164736f6c63430008220033",
}

// BenchHeavyStateABI is the input ABI used to generate the binding from.
// Deprecated: Use BenchHeavyStateMetaData.ABI instead.
var BenchHeavyStateABI = BenchHeavyStateMetaData.ABI

// BenchHeavyStateBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BenchHeavyStateMetaData.Bin instead.
var BenchHeavyStateBin = BenchHeavyStateMetaData.Bin

// DeployBenchHeavyState deploys a new Ethereum contract, binding an instance of BenchHeavyState to it.
func DeployBenchHeavyState(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BenchHeavyState, error) {
	parsed, err := BenchHeavyStateMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BenchHeavyStateBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BenchHeavyState{BenchHeavyStateCaller: BenchHeavyStateCaller{contract: contract}, BenchHeavyStateTransactor: BenchHeavyStateTransactor{contract: contract}, BenchHeavyStateFilterer: BenchHeavyStateFilterer{contract: contract}}, nil
}

// BenchHeavyState is an auto generated Go binding around an Ethereum contract.
type BenchHeavyState struct {
	BenchHeavyStateCaller     // Read-only binding to the contract
	BenchHeavyStateTransactor // Write-only binding to the contract
	BenchHeavyStateFilterer   // Log filterer for contract events
}

// BenchHeavyStateCaller is an auto generated read-only Go binding around an Ethereum contract.
type BenchHeavyStateCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BenchHeavyStateTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BenchHeavyStateTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BenchHeavyStateFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BenchHeavyStateFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BenchHeavyStateSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BenchHeavyStateSession struct {
	Contract     *BenchHeavyState  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BenchHeavyStateCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BenchHeavyStateCallerSession struct {
	Contract *BenchHeavyStateCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// BenchHeavyStateTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BenchHeavyStateTransactorSession struct {
	Contract     *BenchHeavyStateTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// BenchHeavyStateRaw is an auto generated low-level Go binding around an Ethereum contract.
type BenchHeavyStateRaw struct {
	Contract *BenchHeavyState // Generic contract binding to access the raw methods on
}

// BenchHeavyStateCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BenchHeavyStateCallerRaw struct {
	Contract *BenchHeavyStateCaller // Generic read-only contract binding to access the raw methods on
}

// BenchHeavyStateTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BenchHeavyStateTransactorRaw struct {
	Contract *BenchHeavyStateTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBenchHeavyState creates a new instance of BenchHeavyState, bound to a specific deployed contract.
func NewBenchHeavyState(address common.Address, backend bind.ContractBackend) (*BenchHeavyState, error) {
	contract, err := bindBenchHeavyState(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BenchHeavyState{BenchHeavyStateCaller: BenchHeavyStateCaller{contract: contract}, BenchHeavyStateTransactor: BenchHeavyStateTransactor{contract: contract}, BenchHeavyStateFilterer: BenchHeavyStateFilterer{contract: contract}}, nil
}

// NewBenchHeavyStateCaller creates a new read-only instance of BenchHeavyState, bound to a specific deployed contract.
func NewBenchHeavyStateCaller(address common.Address, caller bind.ContractCaller) (*BenchHeavyStateCaller, error) {
	contract, err := bindBenchHeavyState(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BenchHeavyStateCaller{contract: contract}, nil
}

// NewBenchHeavyStateTransactor creates a new write-only instance of BenchHeavyState, bound to a specific deployed contract.
func NewBenchHeavyStateTransactor(address common.Address, transactor bind.ContractTransactor) (*BenchHeavyStateTransactor, error) {
	contract, err := bindBenchHeavyState(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BenchHeavyStateTransactor{contract: contract}, nil
}

// NewBenchHeavyStateFilterer creates a new log filterer instance of BenchHeavyState, bound to a specific deployed contract.
func NewBenchHeavyStateFilterer(address common.Address, filterer bind.ContractFilterer) (*BenchHeavyStateFilterer, error) {
	contract, err := bindBenchHeavyState(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BenchHeavyStateFilterer{contract: contract}, nil
}

// bindBenchHeavyState binds a generic wrapper to an already deployed contract.
func bindBenchHeavyState(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BenchHeavyStateMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BenchHeavyState *BenchHeavyStateRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BenchHeavyState.Contract.BenchHeavyStateCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BenchHeavyState *BenchHeavyStateRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BenchHeavyState.Contract.BenchHeavyStateTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BenchHeavyState *BenchHeavyStateRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BenchHeavyState.Contract.BenchHeavyStateTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BenchHeavyState *BenchHeavyStateCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BenchHeavyState.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BenchHeavyState *BenchHeavyStateTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BenchHeavyState.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BenchHeavyState *BenchHeavyStateTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BenchHeavyState.Contract.contract.Transact(opts, method, params...)
}

// LocalState is a free data retrieval call binding the contract method 0xfb0756bb.
//
// Solidity: function localState(address , uint256 ) view returns(uint256)
func (_BenchHeavyState *BenchHeavyStateCaller) LocalState(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BenchHeavyState.contract.Call(opts, &out, "localState", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LocalState is a free data retrieval call binding the contract method 0xfb0756bb.
//
// Solidity: function localState(address , uint256 ) view returns(uint256)
func (_BenchHeavyState *BenchHeavyStateSession) LocalState(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _BenchHeavyState.Contract.LocalState(&_BenchHeavyState.CallOpts, arg0, arg1)
}

// LocalState is a free data retrieval call binding the contract method 0xfb0756bb.
//
// Solidity: function localState(address , uint256 ) view returns(uint256)
func (_BenchHeavyState *BenchHeavyStateCallerSession) LocalState(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _BenchHeavyState.Contract.LocalState(&_BenchHeavyState.CallOpts, arg0, arg1)
}

// SenderNonce is a free data retrieval call binding the contract method 0x9c90b443.
//
// Solidity: function senderNonce(address ) view returns(uint256)
func (_BenchHeavyState *BenchHeavyStateCaller) SenderNonce(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BenchHeavyState.contract.Call(opts, &out, "senderNonce", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SenderNonce is a free data retrieval call binding the contract method 0x9c90b443.
//
// Solidity: function senderNonce(address ) view returns(uint256)
func (_BenchHeavyState *BenchHeavyStateSession) SenderNonce(arg0 common.Address) (*big.Int, error) {
	return _BenchHeavyState.Contract.SenderNonce(&_BenchHeavyState.CallOpts, arg0)
}

// SenderNonce is a free data retrieval call binding the contract method 0x9c90b443.
//
// Solidity: function senderNonce(address ) view returns(uint256)
func (_BenchHeavyState *BenchHeavyStateCallerSession) SenderNonce(arg0 common.Address) (*big.Int, error) {
	return _BenchHeavyState.Contract.SenderNonce(&_BenchHeavyState.CallOpts, arg0)
}

// SharedState is a free data retrieval call binding the contract method 0x6712fc83.
//
// Solidity: function sharedState(uint256 ) view returns(uint256)
func (_BenchHeavyState *BenchHeavyStateCaller) SharedState(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BenchHeavyState.contract.Call(opts, &out, "sharedState", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SharedState is a free data retrieval call binding the contract method 0x6712fc83.
//
// Solidity: function sharedState(uint256 ) view returns(uint256)
func (_BenchHeavyState *BenchHeavyStateSession) SharedState(arg0 *big.Int) (*big.Int, error) {
	return _BenchHeavyState.Contract.SharedState(&_BenchHeavyState.CallOpts, arg0)
}

// SharedState is a free data retrieval call binding the contract method 0x6712fc83.
//
// Solidity: function sharedState(uint256 ) view returns(uint256)
func (_BenchHeavyState *BenchHeavyStateCallerSession) SharedState(arg0 *big.Int) (*big.Int, error) {
	return _BenchHeavyState.Contract.SharedState(&_BenchHeavyState.CallOpts, arg0)
}

// TotalCalls is a free data retrieval call binding the contract method 0x3af3f24f.
//
// Solidity: function totalCalls() view returns(uint256)
func (_BenchHeavyState *BenchHeavyStateCaller) TotalCalls(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BenchHeavyState.contract.Call(opts, &out, "totalCalls")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalCalls is a free data retrieval call binding the contract method 0x3af3f24f.
//
// Solidity: function totalCalls() view returns(uint256)
func (_BenchHeavyState *BenchHeavyStateSession) TotalCalls() (*big.Int, error) {
	return _BenchHeavyState.Contract.TotalCalls(&_BenchHeavyState.CallOpts)
}

// TotalCalls is a free data retrieval call binding the contract method 0x3af3f24f.
//
// Solidity: function totalCalls() view returns(uint256)
func (_BenchHeavyState *BenchHeavyStateCallerSession) TotalCalls() (*big.Int, error) {
	return _BenchHeavyState.Contract.TotalCalls(&_BenchHeavyState.CallOpts)
}

// WriteMixed is a paid mutator transaction binding the contract method 0x7746c1d2.
//
// Solidity: function writeMixed(uint256 sharedCount, uint256 localCount) returns()
func (_BenchHeavyState *BenchHeavyStateTransactor) WriteMixed(opts *bind.TransactOpts, sharedCount *big.Int, localCount *big.Int) (*types.Transaction, error) {
	return _BenchHeavyState.contract.Transact(opts, "writeMixed", sharedCount, localCount)
}

// WriteMixed is a paid mutator transaction binding the contract method 0x7746c1d2.
//
// Solidity: function writeMixed(uint256 sharedCount, uint256 localCount) returns()
func (_BenchHeavyState *BenchHeavyStateSession) WriteMixed(sharedCount *big.Int, localCount *big.Int) (*types.Transaction, error) {
	return _BenchHeavyState.Contract.WriteMixed(&_BenchHeavyState.TransactOpts, sharedCount, localCount)
}

// WriteMixed is a paid mutator transaction binding the contract method 0x7746c1d2.
//
// Solidity: function writeMixed(uint256 sharedCount, uint256 localCount) returns()
func (_BenchHeavyState *BenchHeavyStateTransactorSession) WriteMixed(sharedCount *big.Int, localCount *big.Int) (*types.Transaction, error) {
	return _BenchHeavyState.Contract.WriteMixed(&_BenchHeavyState.TransactOpts, sharedCount, localCount)
}

// PackWriteMixed returns the ABI-encoded calldata for writeMixed(sharedCount, localCount).
func PackWriteMixed(sharedCount, localCount int64) ([]byte, error) {
	parsed, err := BenchHeavyStateMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return parsed.Pack("writeMixed", big.NewInt(sharedCount), big.NewInt(localCount))
}
