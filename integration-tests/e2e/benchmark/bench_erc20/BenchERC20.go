// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bench_erc20

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

// BenchErc20MetaData contains all meta data concerning the BenchErc20 contract.
var BenchErc20MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b506105fe8061001c5f395ff3fe608060405234801561000f575f5ffd5b506004361061004a575f3560e01c806318160ddd1461004e57806340c10f191461006c57806370a0823114610088578063a9059cbb146100b8575b5f5ffd5b6100566100e8565b6040516100639190610380565b60405180910390f35b61008660048036038101906100819190610421565b6100ee565b005b6100a2600480360381019061009d919061045f565b6101c1565b6040516100af9190610380565b60405180910390f35b6100d260048036038101906100cd9190610421565b6101d5565b6040516100df91906104a4565b60405180910390f35b60015481565b805f5f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825461013991906104ea565b925050819055508060015f82825461015191906104ea565b925050819055508173ffffffffffffffffffffffffffffffffffffffff165f73ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516101b59190610380565b60405180910390a35050565b5f602052805f5260405f205f915090505481565b5f815f5f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20541015610255576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161024c90610577565b60405180910390fd5b815f5f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546102a09190610595565b92505081905550815f5f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546102f291906104ea565b925050819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516103569190610380565b60405180910390a36001905092915050565b5f819050919050565b61037a81610368565b82525050565b5f6020820190506103935f830184610371565b92915050565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6103c68261039d565b9050919050565b6103d6816103bc565b81146103e0575f5ffd5b50565b5f813590506103f1816103cd565b92915050565b61040081610368565b811461040a575f5ffd5b50565b5f8135905061041b816103f7565b92915050565b5f5f6040838503121561043757610436610399565b5b5f610444858286016103e3565b92505060206104558582860161040d565b9150509250929050565b5f6020828403121561047457610473610399565b5b5f610481848285016103e3565b91505092915050565b5f8115159050919050565b61049e8161048a565b82525050565b5f6020820190506104b75f830184610495565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6104f482610368565b91506104ff83610368565b9250828201905080821115610517576105166104bd565b5b92915050565b5f82825260208201905092915050565b7f696e73756666696369656e742062616c616e63650000000000000000000000005f82015250565b5f61056160148361051d565b915061056c8261052d565b602082019050919050565b5f6020820190508181035f83015261058e81610555565b9050919050565b5f61059f82610368565b91506105aa83610368565b92508282039050818111156105c2576105c16104bd565b5b9291505056fea264697066735822122048bccda8d2a5829d2b4ea34ec65e329301afc1d01fa658a167bce73b51359c2764736f6c63430008220033",
}

// BenchErc20ABI is the input ABI used to generate the binding from.
// Deprecated: Use BenchErc20MetaData.ABI instead.
var BenchErc20ABI = BenchErc20MetaData.ABI

// BenchErc20Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BenchErc20MetaData.Bin instead.
var BenchErc20Bin = BenchErc20MetaData.Bin

// DeployBenchErc20 deploys a new Ethereum contract, binding an instance of BenchErc20 to it.
func DeployBenchErc20(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BenchErc20, error) {
	parsed, err := BenchErc20MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BenchErc20Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BenchErc20{BenchErc20Caller: BenchErc20Caller{contract: contract}, BenchErc20Transactor: BenchErc20Transactor{contract: contract}, BenchErc20Filterer: BenchErc20Filterer{contract: contract}}, nil
}

// BenchErc20 is an auto generated Go binding around an Ethereum contract.
type BenchErc20 struct {
	BenchErc20Caller     // Read-only binding to the contract
	BenchErc20Transactor // Write-only binding to the contract
	BenchErc20Filterer   // Log filterer for contract events
}

// BenchErc20Caller is an auto generated read-only Go binding around an Ethereum contract.
type BenchErc20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BenchErc20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type BenchErc20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BenchErc20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BenchErc20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BenchErc20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BenchErc20Session struct {
	Contract     *BenchErc20       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BenchErc20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BenchErc20CallerSession struct {
	Contract *BenchErc20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// BenchErc20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BenchErc20TransactorSession struct {
	Contract     *BenchErc20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// BenchErc20Raw is an auto generated low-level Go binding around an Ethereum contract.
type BenchErc20Raw struct {
	Contract *BenchErc20 // Generic contract binding to access the raw methods on
}

// BenchErc20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BenchErc20CallerRaw struct {
	Contract *BenchErc20Caller // Generic read-only contract binding to access the raw methods on
}

// BenchErc20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BenchErc20TransactorRaw struct {
	Contract *BenchErc20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewBenchErc20 creates a new instance of BenchErc20, bound to a specific deployed contract.
func NewBenchErc20(address common.Address, backend bind.ContractBackend) (*BenchErc20, error) {
	contract, err := bindBenchErc20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BenchErc20{BenchErc20Caller: BenchErc20Caller{contract: contract}, BenchErc20Transactor: BenchErc20Transactor{contract: contract}, BenchErc20Filterer: BenchErc20Filterer{contract: contract}}, nil
}

// NewBenchErc20Caller creates a new read-only instance of BenchErc20, bound to a specific deployed contract.
func NewBenchErc20Caller(address common.Address, caller bind.ContractCaller) (*BenchErc20Caller, error) {
	contract, err := bindBenchErc20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BenchErc20Caller{contract: contract}, nil
}

// NewBenchErc20Transactor creates a new write-only instance of BenchErc20, bound to a specific deployed contract.
func NewBenchErc20Transactor(address common.Address, transactor bind.ContractTransactor) (*BenchErc20Transactor, error) {
	contract, err := bindBenchErc20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BenchErc20Transactor{contract: contract}, nil
}

// NewBenchErc20Filterer creates a new log filterer instance of BenchErc20, bound to a specific deployed contract.
func NewBenchErc20Filterer(address common.Address, filterer bind.ContractFilterer) (*BenchErc20Filterer, error) {
	contract, err := bindBenchErc20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BenchErc20Filterer{contract: contract}, nil
}

// bindBenchErc20 binds a generic wrapper to an already deployed contract.
func bindBenchErc20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BenchErc20MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BenchErc20 *BenchErc20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BenchErc20.Contract.BenchErc20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BenchErc20 *BenchErc20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BenchErc20.Contract.BenchErc20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BenchErc20 *BenchErc20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BenchErc20.Contract.BenchErc20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BenchErc20 *BenchErc20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BenchErc20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BenchErc20 *BenchErc20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BenchErc20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BenchErc20 *BenchErc20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BenchErc20.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_BenchErc20 *BenchErc20Caller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BenchErc20.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_BenchErc20 *BenchErc20Session) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _BenchErc20.Contract.BalanceOf(&_BenchErc20.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_BenchErc20 *BenchErc20CallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _BenchErc20.Contract.BalanceOf(&_BenchErc20.CallOpts, arg0)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BenchErc20 *BenchErc20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BenchErc20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BenchErc20 *BenchErc20Session) TotalSupply() (*big.Int, error) {
	return _BenchErc20.Contract.TotalSupply(&_BenchErc20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BenchErc20 *BenchErc20CallerSession) TotalSupply() (*big.Int, error) {
	return _BenchErc20.Contract.TotalSupply(&_BenchErc20.CallOpts)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_BenchErc20 *BenchErc20Transactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BenchErc20.contract.Transact(opts, "mint", to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_BenchErc20 *BenchErc20Session) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BenchErc20.Contract.Mint(&_BenchErc20.TransactOpts, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_BenchErc20 *BenchErc20TransactorSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BenchErc20.Contract.Mint(&_BenchErc20.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_BenchErc20 *BenchErc20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BenchErc20.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_BenchErc20 *BenchErc20Session) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BenchErc20.Contract.Transfer(&_BenchErc20.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_BenchErc20 *BenchErc20TransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BenchErc20.Contract.Transfer(&_BenchErc20.TransactOpts, to, amount)
}

// BenchErc20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the BenchErc20 contract.
type BenchErc20TransferIterator struct {
	Event *BenchErc20Transfer // Event containing the contract specifics and raw log

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
func (it *BenchErc20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BenchErc20Transfer)
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
		it.Event = new(BenchErc20Transfer)
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
func (it *BenchErc20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BenchErc20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BenchErc20Transfer represents a Transfer event raised by the BenchErc20 contract.
type BenchErc20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BenchErc20 *BenchErc20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BenchErc20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BenchErc20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BenchErc20TransferIterator{contract: _BenchErc20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BenchErc20 *BenchErc20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BenchErc20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BenchErc20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BenchErc20Transfer)
				if err := _BenchErc20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BenchErc20 *BenchErc20Filterer) ParseTransfer(log types.Log) (*BenchErc20Transfer, error) {
	event := new(BenchErc20Transfer)
	if err := _BenchErc20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PackMint returns the ABI-encoded calldata for mint(to, amount).
func PackMint(to common.Address, amount *big.Int) ([]byte, error) {
	parsed, err := BenchErc20MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return parsed.Pack("mint", to, amount)
}

// PackTransfer returns the ABI-encoded calldata for transfer(to, amount).
func PackTransfer(to common.Address, amount *big.Int) ([]byte, error) {
	parsed, err := BenchErc20MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return parsed.Pack("transfer", to, amount)
}
