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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"StringsInsufficientHexLength\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newCount\",\"type\":\"uint256\"}],\"name\":\"increased\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"count\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"n\",\"type\":\"uint64\"}],\"name\":\"get_blockhash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"callback_id\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"name\":\"ibc_ack\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"callback_id\",\"type\":\"uint64\"}],\"name\":\"ibc_timeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"increase\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"path\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"req\",\"type\":\"string\"}],\"name\":\"query_cosmos\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"result\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"n\",\"type\":\"uint64\"}],\"name\":\"recursive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040526114f9806100115f395ff3fe60806040526004361061006f575f3560e01c8063619368951161004d57806361936895146100ed578063ac7fde5f14610115578063cad2355414610151578063e8927fbc1461018d5761006f565b806306661abd146100735780630d4f1f9d1461009d57806331a503f0146100c5575b5f80fd5b34801561007e575f80fd5b50610087610197565b6040516100949190610a32565b60405180910390f35b3480156100a8575f80fd5b506100c360048036038101906100be9190610ace565b61019c565b005b3480156100d0575f80fd5b506100eb60048036038101906100e69190610b0c565b6101e3565b005b3480156100f8575f80fd5b50610113600480360381019061010e9190610b0c565b610207565b005b348015610120575f80fd5b5061013b60048036038101906101369190610b0c565b610323565b6040516101489190610b4f565b60405180910390f35b34801561015c575f80fd5b5061017760048036038101906101729190610ca4565b610337565b6040516101849190610d7a565b60405180910390f35b6101956103c0565b005b5f5481565b80156101c8578167ffffffffffffffff165f808282546101bc9190610dc7565b925050819055506101df565b5f808154809291906101d990610dfa565b91905055505b5050565b8067ffffffffffffffff165f808282546101fd9190610dc7565b9250508190555050565b5f8167ffffffffffffffff1603156103205760f173ffffffffffffffffffffffffffffffffffffffff1663d46f64e661023f8361041f565b6040518263ffffffff1660e01b815260040161025b9190610d7a565b6020604051808303815f875af1158015610277573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061029b9190610e55565b5060f173ffffffffffffffffffffffffffffffffffffffff1663d46f64e66102c28361041f565b6040518263ffffffff1660e01b81526004016102de9190610d7a565b6020604051808303815f875af11580156102fa573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061031e9190610e55565b505b50565b5f8167ffffffffffffffff16409050919050565b606060f173ffffffffffffffffffffffffffffffffffffffff1663cad2355484846040518363ffffffff1660e01b8152600401610375929190610e80565b5f604051808303815f875af1158015610390573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f820116820180604052508101906103b89190610f23565b905092915050565b5f808154809291906103d190610dfa565b91905055507f61996fe196f72cb598c483e896a1221263a28bb630480aa89495f737d4a8e3df60015f546104059190610f6a565b5f54604051610415929190610f9d565b60405180910390a1565b606060f173ffffffffffffffffffffffffffffffffffffffff16636af32a55306040518263ffffffff1660e01b815260040161045b9190611003565b5f604051808303815f875af1158015610476573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f8201168201806040525081019061049e9190610f23565b6104a73061052b565b610503636193689560e01b6001866104bf919061101c565b6040516020016104cf9190611066565b6040516020818303038152906040526040516020016104ef92919061110e565b604051602081830303815290604052610558565b6040516020016105159392919061139b565b6040516020818303038152906040529050919050565b60606105518273ffffffffffffffffffffffffffffffffffffffff16601460ff166107dc565b9050919050565b60605f600280845161056a919061142e565b6105749190610dc7565b67ffffffffffffffff81111561058d5761058c610b80565b5b6040519080825280601f01601f1916602001820160405280156105bf5781602001600182028036833780820191505090505b5090507f3000000000000000000000000000000000000000000000000000000000000000815f815181106105f6576105f561146f565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a9053507f7800000000000000000000000000000000000000000000000000000000000000816001815181106106595761065861146f565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a9053505f5b83518110156107d2575f8482815181106106a6576106a561146f565b5b602001015160f81c60f81b60f81c90507f303132333435363738396162636465660000000000000000000000000000000060048260ff16901c60ff16601081106106f3576106f261146f565b5b1a60f81b836002808502018151811061070f5761070e61146f565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a9053507f3031323334353637383961626364656600000000000000000000000000000000600f821660ff16601081106107765761077561146f565b5b1a60f81b8360026001600286020101815181106107965761079561146f565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a905350508080600101915050610689565b5080915050919050565b60605f8390505f60028460026107f2919061142e565b6107fc9190610dc7565b67ffffffffffffffff81111561081557610814610b80565b5b6040519080825280601f01601f1916602001820160405280156108475781602001600182028036833780820191505090505b5090507f3000000000000000000000000000000000000000000000000000000000000000815f8151811061087e5761087d61146f565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a9053507f7800000000000000000000000000000000000000000000000000000000000000816001815181106108e1576108e061146f565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a9053505f600185600261091f919061142e565b6109299190610dc7565b90505b60018111156109c8577f3031323334353637383961626364656600000000000000000000000000000000600f84166010811061096b5761096a61146f565b5b1a60f81b8282815181106109825761098161146f565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a905350600483901c9250806109c19061149c565b905061092c565b505f8214610a0f5784846040517fe22e27eb000000000000000000000000000000000000000000000000000000008152600401610a06929190610f9d565b60405180910390fd5b809250505092915050565b5f819050919050565b610a2c81610a1a565b82525050565b5f602082019050610a455f830184610a23565b92915050565b5f604051905090565b5f80fd5b5f80fd5b5f67ffffffffffffffff82169050919050565b610a7881610a5c565b8114610a82575f80fd5b50565b5f81359050610a9381610a6f565b92915050565b5f8115159050919050565b610aad81610a99565b8114610ab7575f80fd5b50565b5f81359050610ac881610aa4565b92915050565b5f8060408385031215610ae457610ae3610a54565b5b5f610af185828601610a85565b9250506020610b0285828601610aba565b9150509250929050565b5f60208284031215610b2157610b20610a54565b5b5f610b2e84828501610a85565b91505092915050565b5f819050919050565b610b4981610b37565b82525050565b5f602082019050610b625f830184610b40565b92915050565b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610bb682610b70565b810181811067ffffffffffffffff82111715610bd557610bd4610b80565b5b80604052505050565b5f610be7610a4b565b9050610bf38282610bad565b919050565b5f67ffffffffffffffff821115610c1257610c11610b80565b5b610c1b82610b70565b9050602081019050919050565b828183375f83830152505050565b5f610c48610c4384610bf8565b610bde565b905082815260208101848484011115610c6457610c63610b6c565b5b610c6f848285610c28565b509392505050565b5f82601f830112610c8b57610c8a610b68565b5b8135610c9b848260208601610c36565b91505092915050565b5f8060408385031215610cba57610cb9610a54565b5b5f83013567ffffffffffffffff811115610cd757610cd6610a58565b5b610ce385828601610c77565b925050602083013567ffffffffffffffff811115610d0457610d03610a58565b5b610d1085828601610c77565b9150509250929050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f610d4c82610d1a565b610d568185610d24565b9350610d66818560208601610d34565b610d6f81610b70565b840191505092915050565b5f6020820190508181035f830152610d928184610d42565b905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f610dd182610a1a565b9150610ddc83610a1a565b9250828201905080821115610df457610df3610d9a565b5b92915050565b5f610e0482610a1a565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610e3657610e35610d9a565b5b600182019050919050565b5f81519050610e4f81610aa4565b92915050565b5f60208284031215610e6a57610e69610a54565b5b5f610e7784828501610e41565b91505092915050565b5f6040820190508181035f830152610e988185610d42565b90508181036020830152610eac8184610d42565b90509392505050565b5f610ec7610ec284610bf8565b610bde565b905082815260208101848484011115610ee357610ee2610b6c565b5b610eee848285610d34565b509392505050565b5f82601f830112610f0a57610f09610b68565b5b8151610f1a848260208601610eb5565b91505092915050565b5f60208284031215610f3857610f37610a54565b5b5f82015167ffffffffffffffff811115610f5557610f54610a58565b5b610f6184828501610ef6565b91505092915050565b5f610f7482610a1a565b9150610f7f83610a1a565b9250828203905081811115610f9757610f96610d9a565b5b92915050565b5f604082019050610fb05f830185610a23565b610fbd6020830184610a23565b9392505050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610fed82610fc4565b9050919050565b610ffd81610fe3565b82525050565b5f6020820190506110165f830184610ff4565b92915050565b5f61102682610a5c565b915061103183610a5c565b9250828203905067ffffffffffffffff81111561105157611050610d9a565b5b92915050565b61106081610a5c565b82525050565b5f6020820190506110795f830184611057565b92915050565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b5f819050919050565b6110c46110bf8261107f565b6110aa565b82525050565b5f81519050919050565b5f81905092915050565b5f6110e8826110ca565b6110f281856110d4565b9350611102818560208601610d34565b80840191505092915050565b5f61111982856110b3565b60048201915061112982846110de565b91508190509392505050565b5f81905092915050565b7f7b224074797065223a20222f6d696e6965766d2e65766d2e76312e4d736743615f8201527f6c6c222c00000000000000000000000000000000000000000000000000000000602082015250565b5f611199602483611135565b91506111a48261113f565b602482019050919050565b7f2273656e646572223a20220000000000000000000000000000000000000000005f82015250565b5f6111e3600b83611135565b91506111ee826111af565b600b82019050919050565b5f61120382610d1a565b61120d8185611135565b935061121d818560208601610d34565b80840191505092915050565b7f222c0000000000000000000000000000000000000000000000000000000000005f82015250565b5f61125d600283611135565b915061126882611229565b600282019050919050565b7f22636f6e74726163745f61646472223a202200000000000000000000000000005f82015250565b5f6112a7601283611135565b91506112b282611273565b601282019050919050565b7f22696e707574223a2022000000000000000000000000000000000000000000005f82015250565b5f6112f1600a83611135565b91506112fc826112bd565b600a82019050919050565b7f2276616c7565223a202230222c000000000000000000000000000000000000005f82015250565b5f61133b600d83611135565b915061134682611307565b600d82019050919050565b7f226163636573735f6c697374223a205b5d7d00000000000000000000000000005f82015250565b5f611385601283611135565b915061139082611351565b601282019050919050565b5f6113a58261118d565b91506113b0826111d7565b91506113bc82866111f9565b91506113c782611251565b91506113d28261129b565b91506113de82856111f9565b91506113e982611251565b91506113f4826112e5565b915061140082846111f9565b915061140b82611251565b91506114168261132f565b915061142182611379565b9150819050949350505050565b5f61143882610a1a565b915061144383610a1a565b925082820261145181610a1a565b9150828204841483151761146857611467610d9a565b5b5092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f6114a682610a1a565b91505f82036114b8576114b7610d9a565b5b60018203905091905056fea2646970667358221220662342ef612cc67757c0cab65cd6080becaf1f776847bdde4dea3eff72d4342b64736f6c63430008190033",
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
