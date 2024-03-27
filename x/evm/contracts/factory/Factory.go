// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package factory

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

// FactoryMetaData contains all meta data concerning the Factory contract.
var FactoryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"ERC20Created\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"name\":\"deployNewERC20\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b50611d4b8061001d5f395ff3fe608060405234801562000010575f80fd5b50600436106200002c575f3560e01c8063ba2b9c441462000030575b5f80fd5b6200004e600480360381019062000048919062000213565b62000066565b6040516200005d9190620002f6565b60405180910390f35b5f8086868686866040516200007b9062000158565b6200008b95949392919062000382565b604051809103905ff080158015620000a5573d5f803e3d5ffd5b5090508073ffffffffffffffffffffffffffffffffffffffff1663f2fde38b336040518263ffffffff1660e01b8152600401620000e39190620002f6565b5f604051808303815f87803b158015620000fb575f80fd5b505af11580156200010e573d5f803e3d5ffd5b505050507f41039c8b82bb5a365f56b4f9d87cee99e069b450bfdc3d87696fcd87783f24c181604051620001439190620002f6565b60405180910390a18091505095945050505050565b61194680620003d083390190565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f8083601f8401126200019257620001916200016e565b5b8235905067ffffffffffffffff811115620001b257620001b162000172565b5b602083019150836001820283011115620001d157620001d062000176565b5b9250929050565b5f60ff82169050919050565b620001ef81620001d8565b8114620001fa575f80fd5b50565b5f813590506200020d81620001e4565b92915050565b5f805f805f606086880312156200022f576200022e62000166565b5b5f86013567ffffffffffffffff8111156200024f576200024e6200016a565b5b6200025d888289016200017a565b9550955050602086013567ffffffffffffffff8111156200028357620002826200016a565b5b62000291888289016200017a565b93509350506040620002a688828901620001fd565b9150509295509295909350565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f620002de82620002b3565b9050919050565b620002f081620002d2565b82525050565b5f6020820190506200030b5f830184620002e5565b92915050565b5f82825260208201905092915050565b828183375f83830152505050565b5f601f19601f8301169050919050565b5f6200034c838562000311565b93506200035b83858462000321565b62000366836200032f565b840190509392505050565b6200037c81620001d8565b82525050565b5f6060820190508181035f8301526200039d8187896200033f565b90508181036020830152620003b48185876200033f565b9050620003c5604083018462000371565b969550505050505056fe608060405234801562000010575f80fd5b5060405162001946380380620019468339818101604052810190620000369190620002da565b335f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060f173ffffffffffffffffffffffffffffffffffffffff16635e6c57596040518163ffffffff1660e01b81526004015f604051808303815f87803b158015620000bc575f80fd5b505af1158015620000cf573d5f803e3d5ffd5b505050508260039081620000e49190620005a8565b508160049081620000f69190620005a8565b508060055f6101000a81548160ff021916908360ff1602179055505050506200068c565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6200017b8262000133565b810181811067ffffffffffffffff821117156200019d576200019c62000143565b5b80604052505050565b5f620001b16200011a565b9050620001bf828262000170565b919050565b5f67ffffffffffffffff821115620001e157620001e062000143565b5b620001ec8262000133565b9050602081019050919050565b5f5b8381101562000218578082015181840152602081019050620001fb565b5f8484015250505050565b5f620002396200023384620001c4565b620001a6565b9050828152602081018484840111156200025857620002576200012f565b5b62000265848285620001f9565b509392505050565b5f82601f8301126200028457620002836200012b565b5b81516200029684826020860162000223565b91505092915050565b5f60ff82169050919050565b620002b6816200029f565b8114620002c1575f80fd5b50565b5f81519050620002d481620002ab565b92915050565b5f805f60608486031215620002f457620002f362000123565b5b5f84015167ffffffffffffffff81111562000314576200031362000127565b5b62000322868287016200026d565b935050602084015167ffffffffffffffff81111562000346576200034562000127565b5b62000354868287016200026d565b92505060406200036786828701620002c4565b9150509250925092565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680620003c057607f821691505b602082108103620003d657620003d56200037b565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026200043a7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82620003fd565b620004468683620003fd565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f620004906200048a62000484846200045e565b62000467565b6200045e565b9050919050565b5f819050919050565b620004ab8362000470565b620004c3620004ba8262000497565b84845462000409565b825550505050565b5f90565b620004d9620004cb565b620004e6818484620004a0565b505050565b5b818110156200050d57620005015f82620004cf565b600181019050620004ec565b5050565b601f8211156200055c576200052681620003dc565b6200053184620003ee565b8101602085101562000541578190505b620005596200055085620003ee565b830182620004eb565b50505b505050565b5f82821c905092915050565b5f6200057e5f198460080262000561565b1980831691505092915050565b5f6200059883836200056d565b9150826002028217905092915050565b620005b38262000371565b67ffffffffffffffff811115620005cf57620005ce62000143565b5b620005db8254620003a8565b620005e882828562000511565b5f60209050601f8311600181146200061e575f841562000609578287015190505b6200061585826200058b565b86555062000684565b601f1984166200062e86620003dc565b5f5b82811015620006575784890151825560018201915060208501945060208101905062000630565b8683101562000677578489015162000673601f8916826200056d565b8355505b6001600288020188555050505b505050505050565b6112ac806200069a5f395ff3fe608060405234801561000f575f80fd5b50600436106100cd575f3560e01c806370a082311161008a5780639dc29fac116100645780639dc29fac14610213578063a9059cbb1461022f578063dd62ed3e1461025f578063f2fde38b1461028f576100cd565b806370a08231146101a75780638da5cb5b146101d757806395d89b41146101f5576100cd565b806306fdde03146100d1578063095ea7b3146100ef57806318160ddd1461011f57806323b872dd1461013d578063313ce5671461016d57806340c10f191461018b575b5f80fd5b6100d96102ab565b6040516100e69190610ed2565b60405180910390f35b61010960048036038101906101049190610f83565b610337565b6040516101169190610fdb565b60405180910390f35b610127610424565b6040516101349190611003565b60405180910390f35b6101576004803603810190610152919061101c565b61042a565b6040516101649190610fdb565b60405180910390f35b6101756106b6565b6040516101829190611087565b60405180910390f35b6101a560048036038101906101a09190610f83565b6106c8565b005b6101c160048036038101906101bc91906110a0565b61072c565b6040516101ce9190611003565b60405180910390f35b6101df610741565b6040516101ec91906110da565b60405180910390f35b6101fd610764565b60405161020a9190610ed2565b60405180910390f35b61022d60048036038101906102289190610f83565b6107f0565b005b61024960048036038101906102449190610f83565b610854565b6040516102569190610fdb565b60405180910390f35b610279600480360381019061027491906110f3565b610a51565b6040516102869190611003565b60405180910390f35b6102a960048036038101906102a491906110a0565b610a71565b005b600380546102b89061115e565b80601f01602080910402602001604051908101604052809291908181526020018280546102e49061115e565b801561032f5780601f106103065761010080835404028352916020019161032f565b820191905f5260205f20905b81548152906001019060200180831161031257829003601f168201915b505050505081565b5f8160025f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040516104129190611003565b60405180910390a36001905092915050565b60065481565b5f8260f173ffffffffffffffffffffffffffffffffffffffff16634e25ab64826040518263ffffffff1660e01b815260040161046691906110da565b602060405180830381865afa158015610481573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906104a591906111b8565b6105115760f173ffffffffffffffffffffffffffffffffffffffff1663ceeae52a826040518263ffffffff1660e01b81526004016104e391906110da565b5f604051808303815f87803b1580156104fa575f80fd5b505af115801561050c573d5f803e3d5ffd5b505050505b8260025f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546105989190611210565b925050819055508260015f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546105eb9190611210565b925050819055508260015f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825461063e9190611243565b925050819055508373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef856040516106a29190611003565b60405180910390a360019150509392505050565b60055f9054906101000a900460ff1681565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461071e575f80fd5b6107288282610bb9565b5050565b6001602052805f5260405f205f915090505481565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600480546107719061115e565b80601f016020809104026020016040519081016040528092919081815260200182805461079d9061115e565b80156107e85780601f106107bf576101008083540402835291602001916107e8565b820191905f5260205f20905b8154815290600101906020018083116107cb57829003601f168201915b505050505081565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610846575f80fd5b6108508282610d74565b5050565b5f8260f173ffffffffffffffffffffffffffffffffffffffff16634e25ab64826040518263ffffffff1660e01b815260040161089091906110da565b602060405180830381865afa1580156108ab573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906108cf91906111b8565b61093b5760f173ffffffffffffffffffffffffffffffffffffffff1663ceeae52a826040518263ffffffff1660e01b815260040161090d91906110da565b5f604051808303815f87803b158015610924575f80fd5b505af1158015610936573d5f803e3d5ffd5b505050505b8260015f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546109879190611210565b925050819055508260015f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546109da9190611243565b925050819055508373ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef85604051610a3e9190611003565b60405180910390a3600191505092915050565b6002602052815f5260405f20602052805f5260405f205f91509150505481565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610ac7575f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610afe575f80fd5b8073ffffffffffffffffffffffffffffffffffffffff165f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3805f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b8160f173ffffffffffffffffffffffffffffffffffffffff16634e25ab64826040518263ffffffff1660e01b8152600401610bf491906110da565b602060405180830381865afa158015610c0f573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610c3391906111b8565b610c9f5760f173ffffffffffffffffffffffffffffffffffffffff1663ceeae52a826040518263ffffffff1660e01b8152600401610c7191906110da565b5f604051808303815f87803b158015610c88575f80fd5b505af1158015610c9a573d5f803e3d5ffd5b505050505b8160015f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254610ceb9190611243565b925050819055508160065f828254610d039190611243565b925050819055508273ffffffffffffffffffffffffffffffffffffffff165f73ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051610d679190611003565b60405180910390a3505050565b8060015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254610dc09190611210565b925050819055508060065f828254610dd89190611210565b925050819055505f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051610e3c9190611003565b60405180910390a35050565b5f81519050919050565b5f82825260208201905092915050565b5f5b83811015610e7f578082015181840152602081019050610e64565b5f8484015250505050565b5f601f19601f8301169050919050565b5f610ea482610e48565b610eae8185610e52565b9350610ebe818560208601610e62565b610ec781610e8a565b840191505092915050565b5f6020820190508181035f830152610eea8184610e9a565b905092915050565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610f1f82610ef6565b9050919050565b610f2f81610f15565b8114610f39575f80fd5b50565b5f81359050610f4a81610f26565b92915050565b5f819050919050565b610f6281610f50565b8114610f6c575f80fd5b50565b5f81359050610f7d81610f59565b92915050565b5f8060408385031215610f9957610f98610ef2565b5b5f610fa685828601610f3c565b9250506020610fb785828601610f6f565b9150509250929050565b5f8115159050919050565b610fd581610fc1565b82525050565b5f602082019050610fee5f830184610fcc565b92915050565b610ffd81610f50565b82525050565b5f6020820190506110165f830184610ff4565b92915050565b5f805f6060848603121561103357611032610ef2565b5b5f61104086828701610f3c565b935050602061105186828701610f3c565b925050604061106286828701610f6f565b9150509250925092565b5f60ff82169050919050565b6110818161106c565b82525050565b5f60208201905061109a5f830184611078565b92915050565b5f602082840312156110b5576110b4610ef2565b5b5f6110c284828501610f3c565b91505092915050565b6110d481610f15565b82525050565b5f6020820190506110ed5f8301846110cb565b92915050565b5f806040838503121561110957611108610ef2565b5b5f61111685828601610f3c565b925050602061112785828601610f3c565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061117557607f821691505b60208210810361118857611187611131565b5b50919050565b61119781610fc1565b81146111a1575f80fd5b50565b5f815190506111b28161118e565b92915050565b5f602082840312156111cd576111cc610ef2565b5b5f6111da848285016111a4565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61121a82610f50565b915061122583610f50565b925082820390508181111561123d5761123c6111e3565b5b92915050565b5f61124d82610f50565b915061125883610f50565b92508282019050808211156112705761126f6111e3565b5b9291505056fea26469706673582212201654ff48dacb1ac69db2c0dc8641bbd065c12ffafc824dd3b317fb1ba237e19f64736f6c63430008180033a2646970667358221220eae2db968704a40aecf7b994210f1b458649fff8cb2925b940404c793cfe18ba64736f6c63430008180033",
}

// FactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use FactoryMetaData.ABI instead.
var FactoryABI = FactoryMetaData.ABI

// FactoryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use FactoryMetaData.Bin instead.
var FactoryBin = FactoryMetaData.Bin

// DeployFactory deploys a new Ethereum contract, binding an instance of Factory to it.
func DeployFactory(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Factory, error) {
	parsed, err := FactoryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(FactoryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Factory{FactoryCaller: FactoryCaller{contract: contract}, FactoryTransactor: FactoryTransactor{contract: contract}, FactoryFilterer: FactoryFilterer{contract: contract}}, nil
}

// Factory is an auto generated Go binding around an Ethereum contract.
type Factory struct {
	FactoryCaller     // Read-only binding to the contract
	FactoryTransactor // Write-only binding to the contract
	FactoryFilterer   // Log filterer for contract events
}

// FactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type FactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FactorySession struct {
	Contract     *Factory          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FactoryCallerSession struct {
	Contract *FactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// FactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FactoryTransactorSession struct {
	Contract     *FactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// FactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type FactoryRaw struct {
	Contract *Factory // Generic contract binding to access the raw methods on
}

// FactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FactoryCallerRaw struct {
	Contract *FactoryCaller // Generic read-only contract binding to access the raw methods on
}

// FactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FactoryTransactorRaw struct {
	Contract *FactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFactory creates a new instance of Factory, bound to a specific deployed contract.
func NewFactory(address common.Address, backend bind.ContractBackend) (*Factory, error) {
	contract, err := bindFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Factory{FactoryCaller: FactoryCaller{contract: contract}, FactoryTransactor: FactoryTransactor{contract: contract}, FactoryFilterer: FactoryFilterer{contract: contract}}, nil
}

// NewFactoryCaller creates a new read-only instance of Factory, bound to a specific deployed contract.
func NewFactoryCaller(address common.Address, caller bind.ContractCaller) (*FactoryCaller, error) {
	contract, err := bindFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FactoryCaller{contract: contract}, nil
}

// NewFactoryTransactor creates a new write-only instance of Factory, bound to a specific deployed contract.
func NewFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*FactoryTransactor, error) {
	contract, err := bindFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FactoryTransactor{contract: contract}, nil
}

// NewFactoryFilterer creates a new log filterer instance of Factory, bound to a specific deployed contract.
func NewFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*FactoryFilterer, error) {
	contract, err := bindFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FactoryFilterer{contract: contract}, nil
}

// bindFactory binds a generic wrapper to an already deployed contract.
func bindFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Factory *FactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Factory.Contract.FactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Factory *FactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Factory.Contract.FactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Factory *FactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Factory.Contract.FactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Factory *FactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Factory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Factory *FactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Factory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Factory *FactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Factory.Contract.contract.Transact(opts, method, params...)
}

// DeployNewERC20 is a paid mutator transaction binding the contract method 0xba2b9c44.
//
// Solidity: function deployNewERC20(string name, string symbol, uint8 decimals) returns(address)
func (_Factory *FactoryTransactor) DeployNewERC20(opts *bind.TransactOpts, name string, symbol string, decimals uint8) (*types.Transaction, error) {
	return _Factory.contract.Transact(opts, "deployNewERC20", name, symbol, decimals)
}

// DeployNewERC20 is a paid mutator transaction binding the contract method 0xba2b9c44.
//
// Solidity: function deployNewERC20(string name, string symbol, uint8 decimals) returns(address)
func (_Factory *FactorySession) DeployNewERC20(name string, symbol string, decimals uint8) (*types.Transaction, error) {
	return _Factory.Contract.DeployNewERC20(&_Factory.TransactOpts, name, symbol, decimals)
}

// DeployNewERC20 is a paid mutator transaction binding the contract method 0xba2b9c44.
//
// Solidity: function deployNewERC20(string name, string symbol, uint8 decimals) returns(address)
func (_Factory *FactoryTransactorSession) DeployNewERC20(name string, symbol string, decimals uint8) (*types.Transaction, error) {
	return _Factory.Contract.DeployNewERC20(&_Factory.TransactOpts, name, symbol, decimals)
}

// FactoryERC20CreatedIterator is returned from FilterERC20Created and is used to iterate over the raw logs and unpacked data for ERC20Created events raised by the Factory contract.
type FactoryERC20CreatedIterator struct {
	Event *FactoryERC20Created // Event containing the contract specifics and raw log

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
func (it *FactoryERC20CreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FactoryERC20Created)
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
		it.Event = new(FactoryERC20Created)
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
func (it *FactoryERC20CreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FactoryERC20CreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FactoryERC20Created represents a ERC20Created event raised by the Factory contract.
type FactoryERC20Created struct {
	TokenAddress common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterERC20Created is a free log retrieval operation binding the contract event 0x41039c8b82bb5a365f56b4f9d87cee99e069b450bfdc3d87696fcd87783f24c1.
//
// Solidity: event ERC20Created(address tokenAddress)
func (_Factory *FactoryFilterer) FilterERC20Created(opts *bind.FilterOpts) (*FactoryERC20CreatedIterator, error) {

	logs, sub, err := _Factory.contract.FilterLogs(opts, "ERC20Created")
	if err != nil {
		return nil, err
	}
	return &FactoryERC20CreatedIterator{contract: _Factory.contract, event: "ERC20Created", logs: logs, sub: sub}, nil
}

// WatchERC20Created is a free log subscription operation binding the contract event 0x41039c8b82bb5a365f56b4f9d87cee99e069b450bfdc3d87696fcd87783f24c1.
//
// Solidity: event ERC20Created(address tokenAddress)
func (_Factory *FactoryFilterer) WatchERC20Created(opts *bind.WatchOpts, sink chan<- *FactoryERC20Created) (event.Subscription, error) {

	logs, sub, err := _Factory.contract.WatchLogs(opts, "ERC20Created")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FactoryERC20Created)
				if err := _Factory.contract.UnpackLog(event, "ERC20Created", log); err != nil {
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

// ParseERC20Created is a log parse operation binding the contract event 0x41039c8b82bb5a365f56b4f9d87cee99e069b450bfdc3d87696fcd87783f24c1.
//
// Solidity: event ERC20Created(address tokenAddress)
func (_Factory *FactoryFilterer) ParseERC20Created(log types.Log) (*FactoryERC20Created, error) {
	event := new(FactoryERC20Created)
	if err := _Factory.contract.UnpackLog(event, "ERC20Created", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
