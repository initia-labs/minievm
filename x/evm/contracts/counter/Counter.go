// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package counter

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

// CounterMetaData contains all meta data concerning the Counter contract.
var CounterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newCount\",\"type\":\"uint256\"}],\"name\":\"increased\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"count\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"n\",\"type\":\"uint64\"}],\"name\":\"get_blockhash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"callback_id\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"name\":\"ibc_ack\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"callback_id\",\"type\":\"uint64\"}],\"name\":\"ibc_timeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"increase\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"path\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"req\",\"type\":\"string\"}],\"name\":\"query_cosmos\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"result\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052610861806100115f395ff3fe608060405260043610610054575f3560e01c806306661abd146100585780630d4f1f9d1461008257806331a503f0146100aa578063ac7fde5f146100d2578063cad235541461010e578063e8927fbc1461014a575b5f80fd5b348015610063575f80fd5b5061006c610154565b60405161007991906102d8565b60405180910390f35b34801561008d575f80fd5b506100a860048036038101906100a39190610374565b610159565b005b3480156100b5575f80fd5b506100d060048036038101906100cb91906103b2565b6101a0565b005b3480156100dd575f80fd5b506100f860048036038101906100f391906103b2565b6101c4565b60405161010591906103f5565b60405180910390f35b348015610119575f80fd5b50610134600480360381019061012f919061054a565b6101d8565b6040516101419190610620565b60405180910390f35b610152610261565b005b5f5481565b8015610185578167ffffffffffffffff165f80828254610179919061066d565b9250508190555061019c565b5f80815480929190610196906106a0565b91905055505b5050565b8067ffffffffffffffff165f808282546101ba919061066d565b9250508190555050565b5f8167ffffffffffffffff16409050919050565b606060f173ffffffffffffffffffffffffffffffffffffffff1663cad2355484846040518363ffffffff1660e01b81526004016102169291906106e7565b5f604051808303815f875af1158015610231573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f82011682018060405250810190610259919061078a565b905092915050565b5f80815480929190610272906106a0565b91905055507f61996fe196f72cb598c483e896a1221263a28bb630480aa89495f737d4a8e3df60015f546102a691906107d1565b5f546040516102b6929190610804565b60405180910390a1565b5f819050919050565b6102d2816102c0565b82525050565b5f6020820190506102eb5f8301846102c9565b92915050565b5f604051905090565b5f80fd5b5f80fd5b5f67ffffffffffffffff82169050919050565b61031e81610302565b8114610328575f80fd5b50565b5f8135905061033981610315565b92915050565b5f8115159050919050565b6103538161033f565b811461035d575f80fd5b50565b5f8135905061036e8161034a565b92915050565b5f806040838503121561038a576103896102fa565b5b5f6103978582860161032b565b92505060206103a885828601610360565b9150509250929050565b5f602082840312156103c7576103c66102fa565b5b5f6103d48482850161032b565b91505092915050565b5f819050919050565b6103ef816103dd565b82525050565b5f6020820190506104085f8301846103e6565b92915050565b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b61045c82610416565b810181811067ffffffffffffffff8211171561047b5761047a610426565b5b80604052505050565b5f61048d6102f1565b90506104998282610453565b919050565b5f67ffffffffffffffff8211156104b8576104b7610426565b5b6104c182610416565b9050602081019050919050565b828183375f83830152505050565b5f6104ee6104e98461049e565b610484565b90508281526020810184848401111561050a57610509610412565b5b6105158482856104ce565b509392505050565b5f82601f8301126105315761053061040e565b5b81356105418482602086016104dc565b91505092915050565b5f80604083850312156105605761055f6102fa565b5b5f83013567ffffffffffffffff81111561057d5761057c6102fe565b5b6105898582860161051d565b925050602083013567ffffffffffffffff8111156105aa576105a96102fe565b5b6105b68582860161051d565b9150509250929050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f6105f2826105c0565b6105fc81856105ca565b935061060c8185602086016105da565b61061581610416565b840191505092915050565b5f6020820190508181035f83015261063881846105e8565b905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f610677826102c0565b9150610682836102c0565b925082820190508082111561069a57610699610640565b5b92915050565b5f6106aa826102c0565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036106dc576106db610640565b5b600182019050919050565b5f6040820190508181035f8301526106ff81856105e8565b9050818103602083015261071381846105e8565b90509392505050565b5f61072e6107298461049e565b610484565b90508281526020810184848401111561074a57610749610412565b5b6107558482856105da565b509392505050565b5f82601f8301126107715761077061040e565b5b815161078184826020860161071c565b91505092915050565b5f6020828403121561079f5761079e6102fa565b5b5f82015167ffffffffffffffff8111156107bc576107bb6102fe565b5b6107c88482850161075d565b91505092915050565b5f6107db826102c0565b91506107e6836102c0565b92508282039050818111156107fe576107fd610640565b5b92915050565b5f6040820190506108175f8301856102c9565b61082460208301846102c9565b939250505056fea2646970667358221220edf83c9869e7b1f6b208a0b95c5ff4dc0d6020752d0251d1ccb8b12b3d8fabfd64736f6c63430008190033",
}

// CounterABI is the input ABI used to generate the binding from.
// Deprecated: Use CounterMetaData.ABI instead.
var CounterABI = CounterMetaData.ABI

// CounterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CounterMetaData.Bin instead.
var CounterBin = CounterMetaData.Bin

// DeployCounter deploys a new Ethereum contract, binding an instance of Counter to it.
func DeployCounter(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Counter, error) {
	parsed, err := CounterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CounterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Counter{CounterCaller: CounterCaller{contract: contract}, CounterTransactor: CounterTransactor{contract: contract}, CounterFilterer: CounterFilterer{contract: contract}}, nil
}

// Counter is an auto generated Go binding around an Ethereum contract.
type Counter struct {
	CounterCaller     // Read-only binding to the contract
	CounterTransactor // Write-only binding to the contract
	CounterFilterer   // Log filterer for contract events
}

// CounterCaller is an auto generated read-only Go binding around an Ethereum contract.
type CounterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CounterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CounterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CounterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CounterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CounterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CounterSession struct {
	Contract     *Counter          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CounterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CounterCallerSession struct {
	Contract *CounterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// CounterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CounterTransactorSession struct {
	Contract     *CounterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// CounterRaw is an auto generated low-level Go binding around an Ethereum contract.
type CounterRaw struct {
	Contract *Counter // Generic contract binding to access the raw methods on
}

// CounterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CounterCallerRaw struct {
	Contract *CounterCaller // Generic read-only contract binding to access the raw methods on
}

// CounterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CounterTransactorRaw struct {
	Contract *CounterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCounter creates a new instance of Counter, bound to a specific deployed contract.
func NewCounter(address common.Address, backend bind.ContractBackend) (*Counter, error) {
	contract, err := bindCounter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Counter{CounterCaller: CounterCaller{contract: contract}, CounterTransactor: CounterTransactor{contract: contract}, CounterFilterer: CounterFilterer{contract: contract}}, nil
}

// NewCounterCaller creates a new read-only instance of Counter, bound to a specific deployed contract.
func NewCounterCaller(address common.Address, caller bind.ContractCaller) (*CounterCaller, error) {
	contract, err := bindCounter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CounterCaller{contract: contract}, nil
}

// NewCounterTransactor creates a new write-only instance of Counter, bound to a specific deployed contract.
func NewCounterTransactor(address common.Address, transactor bind.ContractTransactor) (*CounterTransactor, error) {
	contract, err := bindCounter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CounterTransactor{contract: contract}, nil
}

// NewCounterFilterer creates a new log filterer instance of Counter, bound to a specific deployed contract.
func NewCounterFilterer(address common.Address, filterer bind.ContractFilterer) (*CounterFilterer, error) {
	contract, err := bindCounter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CounterFilterer{contract: contract}, nil
}

// bindCounter binds a generic wrapper to an already deployed contract.
func bindCounter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CounterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Counter *CounterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Counter.Contract.CounterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Counter *CounterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Counter.Contract.CounterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Counter *CounterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Counter.Contract.CounterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Counter *CounterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Counter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Counter *CounterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Counter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Counter *CounterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Counter.Contract.contract.Transact(opts, method, params...)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_Counter *CounterCaller) Count(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Counter.contract.Call(opts, &out, "count")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_Counter *CounterSession) Count() (*big.Int, error) {
	return _Counter.Contract.Count(&_Counter.CallOpts)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_Counter *CounterCallerSession) Count() (*big.Int, error) {
	return _Counter.Contract.Count(&_Counter.CallOpts)
}

// GetBlockhash is a free data retrieval call binding the contract method 0xac7fde5f.
//
// Solidity: function get_blockhash(uint64 n) view returns(bytes32)
func (_Counter *CounterCaller) GetBlockhash(opts *bind.CallOpts, n uint64) ([32]byte, error) {
	var out []interface{}
	err := _Counter.contract.Call(opts, &out, "get_blockhash", n)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetBlockhash is a free data retrieval call binding the contract method 0xac7fde5f.
//
// Solidity: function get_blockhash(uint64 n) view returns(bytes32)
func (_Counter *CounterSession) GetBlockhash(n uint64) ([32]byte, error) {
	return _Counter.Contract.GetBlockhash(&_Counter.CallOpts, n)
}

// GetBlockhash is a free data retrieval call binding the contract method 0xac7fde5f.
//
// Solidity: function get_blockhash(uint64 n) view returns(bytes32)
func (_Counter *CounterCallerSession) GetBlockhash(n uint64) ([32]byte, error) {
	return _Counter.Contract.GetBlockhash(&_Counter.CallOpts, n)
}

// IbcAck is a paid mutator transaction binding the contract method 0x0d4f1f9d.
//
// Solidity: function ibc_ack(uint64 callback_id, bool success) returns()
func (_Counter *CounterTransactor) IbcAck(opts *bind.TransactOpts, callback_id uint64, success bool) (*types.Transaction, error) {
	return _Counter.contract.Transact(opts, "ibc_ack", callback_id, success)
}

// IbcAck is a paid mutator transaction binding the contract method 0x0d4f1f9d.
//
// Solidity: function ibc_ack(uint64 callback_id, bool success) returns()
func (_Counter *CounterSession) IbcAck(callback_id uint64, success bool) (*types.Transaction, error) {
	return _Counter.Contract.IbcAck(&_Counter.TransactOpts, callback_id, success)
}

// IbcAck is a paid mutator transaction binding the contract method 0x0d4f1f9d.
//
// Solidity: function ibc_ack(uint64 callback_id, bool success) returns()
func (_Counter *CounterTransactorSession) IbcAck(callback_id uint64, success bool) (*types.Transaction, error) {
	return _Counter.Contract.IbcAck(&_Counter.TransactOpts, callback_id, success)
}

// IbcTimeout is a paid mutator transaction binding the contract method 0x31a503f0.
//
// Solidity: function ibc_timeout(uint64 callback_id) returns()
func (_Counter *CounterTransactor) IbcTimeout(opts *bind.TransactOpts, callback_id uint64) (*types.Transaction, error) {
	return _Counter.contract.Transact(opts, "ibc_timeout", callback_id)
}

// IbcTimeout is a paid mutator transaction binding the contract method 0x31a503f0.
//
// Solidity: function ibc_timeout(uint64 callback_id) returns()
func (_Counter *CounterSession) IbcTimeout(callback_id uint64) (*types.Transaction, error) {
	return _Counter.Contract.IbcTimeout(&_Counter.TransactOpts, callback_id)
}

// IbcTimeout is a paid mutator transaction binding the contract method 0x31a503f0.
//
// Solidity: function ibc_timeout(uint64 callback_id) returns()
func (_Counter *CounterTransactorSession) IbcTimeout(callback_id uint64) (*types.Transaction, error) {
	return _Counter.Contract.IbcTimeout(&_Counter.TransactOpts, callback_id)
}

// Increase is a paid mutator transaction binding the contract method 0xe8927fbc.
//
// Solidity: function increase() payable returns()
func (_Counter *CounterTransactor) Increase(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Counter.contract.Transact(opts, "increase")
}

// Increase is a paid mutator transaction binding the contract method 0xe8927fbc.
//
// Solidity: function increase() payable returns()
func (_Counter *CounterSession) Increase() (*types.Transaction, error) {
	return _Counter.Contract.Increase(&_Counter.TransactOpts)
}

// Increase is a paid mutator transaction binding the contract method 0xe8927fbc.
//
// Solidity: function increase() payable returns()
func (_Counter *CounterTransactorSession) Increase() (*types.Transaction, error) {
	return _Counter.Contract.Increase(&_Counter.TransactOpts)
}

// QueryCosmos is a paid mutator transaction binding the contract method 0xcad23554.
//
// Solidity: function query_cosmos(string path, string req) returns(string result)
func (_Counter *CounterTransactor) QueryCosmos(opts *bind.TransactOpts, path string, req string) (*types.Transaction, error) {
	return _Counter.contract.Transact(opts, "query_cosmos", path, req)
}

// QueryCosmos is a paid mutator transaction binding the contract method 0xcad23554.
//
// Solidity: function query_cosmos(string path, string req) returns(string result)
func (_Counter *CounterSession) QueryCosmos(path string, req string) (*types.Transaction, error) {
	return _Counter.Contract.QueryCosmos(&_Counter.TransactOpts, path, req)
}

// QueryCosmos is a paid mutator transaction binding the contract method 0xcad23554.
//
// Solidity: function query_cosmos(string path, string req) returns(string result)
func (_Counter *CounterTransactorSession) QueryCosmos(path string, req string) (*types.Transaction, error) {
	return _Counter.Contract.QueryCosmos(&_Counter.TransactOpts, path, req)
}

// CounterIncreasedIterator is returned from FilterIncreased and is used to iterate over the raw logs and unpacked data for Increased events raised by the Counter contract.
type CounterIncreasedIterator struct {
	Event *CounterIncreased // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CounterIncreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CounterIncreased)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CounterIncreased)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CounterIncreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CounterIncreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CounterIncreased represents a Increased event raised by the Counter contract.
type CounterIncreased struct {
	OldCount *big.Int
	NewCount *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterIncreased is a free log retrieval operation binding the contract event 0x61996fe196f72cb598c483e896a1221263a28bb630480aa89495f737d4a8e3df.
//
// Solidity: event increased(uint256 oldCount, uint256 newCount)
func (_Counter *CounterFilterer) FilterIncreased(opts *bind.FilterOpts) (*CounterIncreasedIterator, error) {

	logs, sub, err := _Counter.contract.FilterLogs(opts, "increased")
	if err != nil {
		return nil, err
	}
	return &CounterIncreasedIterator{contract: _Counter.contract, event: "increased", logs: logs, sub: sub}, nil
}

// WatchIncreased is a free log subscription operation binding the contract event 0x61996fe196f72cb598c483e896a1221263a28bb630480aa89495f737d4a8e3df.
//
// Solidity: event increased(uint256 oldCount, uint256 newCount)
func (_Counter *CounterFilterer) WatchIncreased(opts *bind.WatchOpts, sink chan<- *CounterIncreased) (event.Subscription, error) {

	logs, sub, err := _Counter.contract.WatchLogs(opts, "increased")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CounterIncreased)
				if err := _Counter.contract.UnpackLog(event, "increased", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseIncreased is a log parse operation binding the contract event 0x61996fe196f72cb598c483e896a1221263a28bb630480aa89495f737d4a8e3df.
//
// Solidity: event increased(uint256 oldCount, uint256 newCount)
func (_Counter *CounterFilterer) ParseIncreased(log types.Log) (*CounterIncreased, error) {
	event := new(CounterIncreased)
	if err := _Counter.contract.UnpackLog(event, "increased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
