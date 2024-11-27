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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"StringsInsufficientHexLength\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"reverted\",\"type\":\"bool\"}],\"name\":\"execute_reverted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newCount\",\"type\":\"uint256\"}],\"name\":\"increased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"n\",\"type\":\"uint64\"}],\"name\":\"recursive_called\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"count\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"exec_msg\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"try_catch\",\"type\":\"bool\"}],\"name\":\"execute_cosmos\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"n\",\"type\":\"uint64\"}],\"name\":\"get_blockhash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"callback_id\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"name\":\"ibc_ack\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"callback_id\",\"type\":\"uint64\"}],\"name\":\"ibc_timeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"increase\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"num\",\"type\":\"uint64\"}],\"name\":\"increase_for_fuzz\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"path\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"req\",\"type\":\"string\"}],\"name\":\"query_cosmos\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"result\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"n\",\"type\":\"uint64\"}],\"name\":\"recursive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040526117c5806100115f395ff3fe608060405260043610610085575f3560e01c806331a503f01161005857806331a503f01461012b5780636193689514610153578063ac7fde5f1461017b578063cad23554146101b7578063e8927fbc146101f357610085565b806306661abd146100895780630d4f1f9d146100b357806324c68fce146100db5780632607baf814610103575b5f80fd5b348015610094575f80fd5b5061009d6101fd565b6040516100aa9190610c7c565b60405180910390f35b3480156100be575f80fd5b506100d960048036038101906100d49190610d18565b610202565b005b3480156100e6575f80fd5b5061010160048036038101906100fc9190610e92565b610249565b005b34801561010e575f80fd5b5061012960048036038101906101249190610eec565b6103c3565b005b348015610136575f80fd5b50610151600480360381019061014c9190610eec565b6103f6565b005b34801561015e575f80fd5b5061017960048036038101906101749190610eec565b61041a565b005b348015610186575f80fd5b506101a1600480360381019061019c9190610eec565b61056d565b6040516101ae9190610f2f565b60405180910390f35b3480156101c2575f80fd5b506101dd60048036038101906101d89190610f48565b610581565b6040516101ea919061101e565b60405180910390f35b6101fb61060a565b005b5f5481565b801561022e578167ffffffffffffffff165f80828254610222919061106b565b92505081905550610245565b5f8081548092919061023f9061109e565b91905055505b5050565b80156103435760f173ffffffffffffffffffffffffffffffffffffffff1663d46f64e6836040518263ffffffff1660e01b8152600401610289919061101e565b6020604051808303815f875af19250505080156102c457506040513d601f19601f820116820180604052508101906102c191906110f9565b60015b610305577f1a0d1eca1924bf9fc21c4b3f8b97d45c8fad7f5dbde70453a436a095844b3b6a60016040516102f89190611133565b60405180910390a161033e565b507f1a0d1eca1924bf9fc21c4b3f8b97d45c8fad7f5dbde70453a436a095844b3b6a5f6040516103359190611133565b60405180910390a15b6103bf565b60f173ffffffffffffffffffffffffffffffffffffffff1663d46f64e6836040518263ffffffff1660e01b815260040161037d919061101e565b6020604051808303815f875af1158015610399573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906103bd91906110f9565b505b5050565b5f8167ffffffffffffffff1603156103f3576103dd61060a565b6103f26001826103ed919061114c565b6103c3565b5b50565b8067ffffffffffffffff165f80828254610410919061106b565b9250508190555050565b7f4475bbd12ef452e28d39c4cb77494d85136c2d89ca1354b52188d4aaa8f4ba51816040516104499190611196565b60405180910390a15f8167ffffffffffffffff16031561056a5760f173ffffffffffffffffffffffffffffffffffffffff1663d46f64e661048983610669565b6040518263ffffffff1660e01b81526004016104a5919061101e565b6020604051808303815f875af11580156104c1573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906104e591906110f9565b5060f173ffffffffffffffffffffffffffffffffffffffff1663d46f64e661050c83610669565b6040518263ffffffff1660e01b8152600401610528919061101e565b6020604051808303815f875af1158015610544573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061056891906110f9565b505b50565b5f8167ffffffffffffffff16409050919050565b606060f173ffffffffffffffffffffffffffffffffffffffff1663cad2355484846040518363ffffffff1660e01b81526004016105bf9291906111af565b5f604051808303815f875af11580156105da573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f820116820180604052508101906106029190611252565b905092915050565b5f8081548092919061061b9061109e565b91905055507f61996fe196f72cb598c483e896a1221263a28bb630480aa89495f737d4a8e3df60015f5461064f9190611299565b5f5460405161065f9291906112cc565b60405180910390a1565b606060f173ffffffffffffffffffffffffffffffffffffffff16636af32a55306040518263ffffffff1660e01b81526004016106a59190611332565b5f604051808303815f875af11580156106c0573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f820116820180604052508101906106e89190611252565b6106f130610775565b61074d636193689560e01b600186610709919061114c565b6040516020016107199190611196565b6040516020818303038152906040526040516020016107399291906113da565b6040516020818303038152906040526107a2565b60405160200161075f93929190611667565b6040516020818303038152906040529050919050565b606061079b8273ffffffffffffffffffffffffffffffffffffffff16601460ff16610a26565b9050919050565b60605f60028084516107b491906116fa565b6107be919061106b565b67ffffffffffffffff8111156107d7576107d6610d6e565b5b6040519080825280601f01601f1916602001820160405280156108095781602001600182028036833780820191505090505b5090507f3000000000000000000000000000000000000000000000000000000000000000815f815181106108405761083f61173b565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a9053507f7800000000000000000000000000000000000000000000000000000000000000816001815181106108a3576108a261173b565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a9053505f5b8351811015610a1c575f8482815181106108f0576108ef61173b565b5b602001015160f81c60f81b60f81c90507f303132333435363738396162636465660000000000000000000000000000000060048260ff16901c60ff166010811061093d5761093c61173b565b5b1a60f81b83600280850201815181106109595761095861173b565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a9053507f3031323334353637383961626364656600000000000000000000000000000000600f821660ff16601081106109c0576109bf61173b565b5b1a60f81b8360026001600286020101815181106109e0576109df61173b565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a9053505080806001019150506108d3565b5080915050919050565b60605f8390505f6002846002610a3c91906116fa565b610a46919061106b565b67ffffffffffffffff811115610a5f57610a5e610d6e565b5b6040519080825280601f01601f191660200182016040528015610a915781602001600182028036833780820191505090505b5090507f3000000000000000000000000000000000000000000000000000000000000000815f81518110610ac857610ac761173b565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a9053507f780000000000000000000000000000000000000000000000000000000000000081600181518110610b2b57610b2a61173b565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a9053505f6001856002610b6991906116fa565b610b73919061106b565b90505b6001811115610c12577f3031323334353637383961626364656600000000000000000000000000000000600f841660108110610bb557610bb461173b565b5b1a60f81b828281518110610bcc57610bcb61173b565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a905350600483901c925080610c0b90611768565b9050610b76565b505f8214610c595784846040517fe22e27eb000000000000000000000000000000000000000000000000000000008152600401610c509291906112cc565b60405180910390fd5b809250505092915050565b5f819050919050565b610c7681610c64565b82525050565b5f602082019050610c8f5f830184610c6d565b92915050565b5f604051905090565b5f80fd5b5f80fd5b5f67ffffffffffffffff82169050919050565b610cc281610ca6565b8114610ccc575f80fd5b50565b5f81359050610cdd81610cb9565b92915050565b5f8115159050919050565b610cf781610ce3565b8114610d01575f80fd5b50565b5f81359050610d1281610cee565b92915050565b5f8060408385031215610d2e57610d2d610c9e565b5b5f610d3b85828601610ccf565b9250506020610d4c85828601610d04565b9150509250929050565b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610da482610d5e565b810181811067ffffffffffffffff82111715610dc357610dc2610d6e565b5b80604052505050565b5f610dd5610c95565b9050610de18282610d9b565b919050565b5f67ffffffffffffffff821115610e0057610dff610d6e565b5b610e0982610d5e565b9050602081019050919050565b828183375f83830152505050565b5f610e36610e3184610de6565b610dcc565b905082815260208101848484011115610e5257610e51610d5a565b5b610e5d848285610e16565b509392505050565b5f82601f830112610e7957610e78610d56565b5b8135610e89848260208601610e24565b91505092915050565b5f8060408385031215610ea857610ea7610c9e565b5b5f83013567ffffffffffffffff811115610ec557610ec4610ca2565b5b610ed185828601610e65565b9250506020610ee285828601610d04565b9150509250929050565b5f60208284031215610f0157610f00610c9e565b5b5f610f0e84828501610ccf565b91505092915050565b5f819050919050565b610f2981610f17565b82525050565b5f602082019050610f425f830184610f20565b92915050565b5f8060408385031215610f5e57610f5d610c9e565b5b5f83013567ffffffffffffffff811115610f7b57610f7a610ca2565b5b610f8785828601610e65565b925050602083013567ffffffffffffffff811115610fa857610fa7610ca2565b5b610fb485828601610e65565b9150509250929050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f610ff082610fbe565b610ffa8185610fc8565b935061100a818560208601610fd8565b61101381610d5e565b840191505092915050565b5f6020820190508181035f8301526110368184610fe6565b905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61107582610c64565b915061108083610c64565b92508282019050808211156110985761109761103e565b5b92915050565b5f6110a882610c64565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036110da576110d961103e565b5b600182019050919050565b5f815190506110f381610cee565b92915050565b5f6020828403121561110e5761110d610c9e565b5b5f61111b848285016110e5565b91505092915050565b61112d81610ce3565b82525050565b5f6020820190506111465f830184611124565b92915050565b5f61115682610ca6565b915061116183610ca6565b9250828203905067ffffffffffffffff8111156111815761118061103e565b5b92915050565b61119081610ca6565b82525050565b5f6020820190506111a95f830184611187565b92915050565b5f6040820190508181035f8301526111c78185610fe6565b905081810360208301526111db8184610fe6565b90509392505050565b5f6111f66111f184610de6565b610dcc565b90508281526020810184848401111561121257611211610d5a565b5b61121d848285610fd8565b509392505050565b5f82601f83011261123957611238610d56565b5b81516112498482602086016111e4565b91505092915050565b5f6020828403121561126757611266610c9e565b5b5f82015167ffffffffffffffff81111561128457611283610ca2565b5b61129084828501611225565b91505092915050565b5f6112a382610c64565b91506112ae83610c64565b92508282039050818111156112c6576112c561103e565b5b92915050565b5f6040820190506112df5f830185610c6d565b6112ec6020830184610c6d565b9392505050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61131c826112f3565b9050919050565b61132c81611312565b82525050565b5f6020820190506113455f830184611323565b92915050565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b5f819050919050565b61139061138b8261134b565b611376565b82525050565b5f81519050919050565b5f81905092915050565b5f6113b482611396565b6113be81856113a0565b93506113ce818560208601610fd8565b80840191505092915050565b5f6113e5828561137f565b6004820191506113f582846113aa565b91508190509392505050565b5f81905092915050565b7f7b224074797065223a20222f6d696e6965766d2e65766d2e76312e4d736743615f8201527f6c6c222c00000000000000000000000000000000000000000000000000000000602082015250565b5f611465602483611401565b91506114708261140b565b602482019050919050565b7f2273656e646572223a20220000000000000000000000000000000000000000005f82015250565b5f6114af600b83611401565b91506114ba8261147b565b600b82019050919050565b5f6114cf82610fbe565b6114d98185611401565b93506114e9818560208601610fd8565b80840191505092915050565b7f222c0000000000000000000000000000000000000000000000000000000000005f82015250565b5f611529600283611401565b9150611534826114f5565b600282019050919050565b7f22636f6e74726163745f61646472223a202200000000000000000000000000005f82015250565b5f611573601283611401565b915061157e8261153f565b601282019050919050565b7f22696e707574223a2022000000000000000000000000000000000000000000005f82015250565b5f6115bd600a83611401565b91506115c882611589565b600a82019050919050565b7f2276616c7565223a202230222c000000000000000000000000000000000000005f82015250565b5f611607600d83611401565b9150611612826115d3565b600d82019050919050565b7f226163636573735f6c697374223a205b5d7d00000000000000000000000000005f82015250565b5f611651601283611401565b915061165c8261161d565b601282019050919050565b5f61167182611459565b915061167c826114a3565b915061168882866114c5565b91506116938261151d565b915061169e82611567565b91506116aa82856114c5565b91506116b58261151d565b91506116c0826115b1565b91506116cc82846114c5565b91506116d78261151d565b91506116e2826115fb565b91506116ed82611645565b9150819050949350505050565b5f61170482610c64565b915061170f83610c64565b925082820261171d81610c64565b915082820484148315176117345761173361103e565b5b5092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f61177282610c64565b91505f82036117845761178361103e565b5b60018203905091905056fea26469706673582212203dcfd3e72d386299f9cc3c1e16d724be03a75bf28d1b18677fe42aefd8b2a5b864736f6c63430008190033",
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

// ExecuteCosmos is a paid mutator transaction binding the contract method 0x24c68fce.
//
// Solidity: function execute_cosmos(string exec_msg, bool try_catch) returns()
func (_Counter *CounterTransactor) ExecuteCosmos(opts *bind.TransactOpts, exec_msg string, try_catch bool) (*types.Transaction, error) {
	return _Counter.contract.Transact(opts, "execute_cosmos", exec_msg, try_catch)
}

// ExecuteCosmos is a paid mutator transaction binding the contract method 0x24c68fce.
//
// Solidity: function execute_cosmos(string exec_msg, bool try_catch) returns()
func (_Counter *CounterSession) ExecuteCosmos(exec_msg string, try_catch bool) (*types.Transaction, error) {
	return _Counter.Contract.ExecuteCosmos(&_Counter.TransactOpts, exec_msg, try_catch)
}

// ExecuteCosmos is a paid mutator transaction binding the contract method 0x24c68fce.
//
// Solidity: function execute_cosmos(string exec_msg, bool try_catch) returns()
func (_Counter *CounterTransactorSession) ExecuteCosmos(exec_msg string, try_catch bool) (*types.Transaction, error) {
	return _Counter.Contract.ExecuteCosmos(&_Counter.TransactOpts, exec_msg, try_catch)
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

// IncreaseForFuzz is a paid mutator transaction binding the contract method 0x2607baf8.
//
// Solidity: function increase_for_fuzz(uint64 num) returns()
func (_Counter *CounterTransactor) IncreaseForFuzz(opts *bind.TransactOpts, num uint64) (*types.Transaction, error) {
	return _Counter.contract.Transact(opts, "increase_for_fuzz", num)
}

// IncreaseForFuzz is a paid mutator transaction binding the contract method 0x2607baf8.
//
// Solidity: function increase_for_fuzz(uint64 num) returns()
func (_Counter *CounterSession) IncreaseForFuzz(num uint64) (*types.Transaction, error) {
	return _Counter.Contract.IncreaseForFuzz(&_Counter.TransactOpts, num)
}

// IncreaseForFuzz is a paid mutator transaction binding the contract method 0x2607baf8.
//
// Solidity: function increase_for_fuzz(uint64 num) returns()
func (_Counter *CounterTransactorSession) IncreaseForFuzz(num uint64) (*types.Transaction, error) {
	return _Counter.Contract.IncreaseForFuzz(&_Counter.TransactOpts, num)
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

// Recursive is a paid mutator transaction binding the contract method 0x61936895.
//
// Solidity: function recursive(uint64 n) returns()
func (_Counter *CounterTransactor) Recursive(opts *bind.TransactOpts, n uint64) (*types.Transaction, error) {
	return _Counter.contract.Transact(opts, "recursive", n)
}

// Recursive is a paid mutator transaction binding the contract method 0x61936895.
//
// Solidity: function recursive(uint64 n) returns()
func (_Counter *CounterSession) Recursive(n uint64) (*types.Transaction, error) {
	return _Counter.Contract.Recursive(&_Counter.TransactOpts, n)
}

// Recursive is a paid mutator transaction binding the contract method 0x61936895.
//
// Solidity: function recursive(uint64 n) returns()
func (_Counter *CounterTransactorSession) Recursive(n uint64) (*types.Transaction, error) {
	return _Counter.Contract.Recursive(&_Counter.TransactOpts, n)
}

// CounterExecuteRevertedIterator is returned from FilterExecuteReverted and is used to iterate over the raw logs and unpacked data for ExecuteReverted events raised by the Counter contract.
type CounterExecuteRevertedIterator struct {
	Event *CounterExecuteReverted // Event containing the contract specifics and raw log

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
func (it *CounterExecuteRevertedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CounterExecuteReverted)
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
		it.Event = new(CounterExecuteReverted)
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
func (it *CounterExecuteRevertedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CounterExecuteRevertedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CounterExecuteReverted represents a ExecuteReverted event raised by the Counter contract.
type CounterExecuteReverted struct {
	Reverted bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterExecuteReverted is a free log retrieval operation binding the contract event 0x1a0d1eca1924bf9fc21c4b3f8b97d45c8fad7f5dbde70453a436a095844b3b6a.
//
// Solidity: event execute_reverted(bool reverted)
func (_Counter *CounterFilterer) FilterExecuteReverted(opts *bind.FilterOpts) (*CounterExecuteRevertedIterator, error) {

	logs, sub, err := _Counter.contract.FilterLogs(opts, "execute_reverted")
	if err != nil {
		return nil, err
	}
	return &CounterExecuteRevertedIterator{contract: _Counter.contract, event: "execute_reverted", logs: logs, sub: sub}, nil
}

// WatchExecuteReverted is a free log subscription operation binding the contract event 0x1a0d1eca1924bf9fc21c4b3f8b97d45c8fad7f5dbde70453a436a095844b3b6a.
//
// Solidity: event execute_reverted(bool reverted)
func (_Counter *CounterFilterer) WatchExecuteReverted(opts *bind.WatchOpts, sink chan<- *CounterExecuteReverted) (event.Subscription, error) {

	logs, sub, err := _Counter.contract.WatchLogs(opts, "execute_reverted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CounterExecuteReverted)
				if err := _Counter.contract.UnpackLog(event, "execute_reverted", log); err != nil {
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

// ParseExecuteReverted is a log parse operation binding the contract event 0x1a0d1eca1924bf9fc21c4b3f8b97d45c8fad7f5dbde70453a436a095844b3b6a.
//
// Solidity: event execute_reverted(bool reverted)
func (_Counter *CounterFilterer) ParseExecuteReverted(log types.Log) (*CounterExecuteReverted, error) {
	event := new(CounterExecuteReverted)
	if err := _Counter.contract.UnpackLog(event, "execute_reverted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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

// CounterRecursiveCalledIterator is returned from FilterRecursiveCalled and is used to iterate over the raw logs and unpacked data for RecursiveCalled events raised by the Counter contract.
type CounterRecursiveCalledIterator struct {
	Event *CounterRecursiveCalled // Event containing the contract specifics and raw log

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
func (it *CounterRecursiveCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CounterRecursiveCalled)
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
		it.Event = new(CounterRecursiveCalled)
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
func (it *CounterRecursiveCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CounterRecursiveCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CounterRecursiveCalled represents a RecursiveCalled event raised by the Counter contract.
type CounterRecursiveCalled struct {
	N   uint64
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRecursiveCalled is a free log retrieval operation binding the contract event 0x4475bbd12ef452e28d39c4cb77494d85136c2d89ca1354b52188d4aaa8f4ba51.
//
// Solidity: event recursive_called(uint64 n)
func (_Counter *CounterFilterer) FilterRecursiveCalled(opts *bind.FilterOpts) (*CounterRecursiveCalledIterator, error) {

	logs, sub, err := _Counter.contract.FilterLogs(opts, "recursive_called")
	if err != nil {
		return nil, err
	}
	return &CounterRecursiveCalledIterator{contract: _Counter.contract, event: "recursive_called", logs: logs, sub: sub}, nil
}

// WatchRecursiveCalled is a free log subscription operation binding the contract event 0x4475bbd12ef452e28d39c4cb77494d85136c2d89ca1354b52188d4aaa8f4ba51.
//
// Solidity: event recursive_called(uint64 n)
func (_Counter *CounterFilterer) WatchRecursiveCalled(opts *bind.WatchOpts, sink chan<- *CounterRecursiveCalled) (event.Subscription, error) {

	logs, sub, err := _Counter.contract.WatchLogs(opts, "recursive_called")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CounterRecursiveCalled)
				if err := _Counter.contract.UnpackLog(event, "recursive_called", log); err != nil {
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

// ParseRecursiveCalled is a log parse operation binding the contract event 0x4475bbd12ef452e28d39c4cb77494d85136c2d89ca1354b52188d4aaa8f4ba51.
//
// Solidity: event recursive_called(uint64 n)
func (_Counter *CounterFilterer) ParseRecursiveCalled(log types.Log) (*CounterRecursiveCalled, error) {
	event := new(CounterRecursiveCalled)
	if err := _Counter.contract.UnpackLog(event, "recursive_called", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
