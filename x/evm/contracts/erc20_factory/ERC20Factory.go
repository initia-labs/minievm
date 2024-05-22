// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package erc20_factory

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

// Erc20FactoryMetaData contains all meta data concerning the Erc20Factory contract.
var Erc20FactoryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"erc20\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC20Created\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"name\":\"createERC20\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b50611e818061001d5f395ff3fe608060405234801562000010575f80fd5b50600436106200002c575f3560e01c806306ef1a861462000030575b5f80fd5b6200004e600480360381019062000048919062000390565b62000066565b6040516200005d91906200046a565b60405180910390f35b5f808484846040516200007990620001de565b62000087939291906200051a565b604051809103905ff080158015620000a1573d5f803e3d5ffd5b50905060f273ffffffffffffffffffffffffffffffffffffffff1663d126274a826040518263ffffffff1660e01b8152600401620000e091906200046a565b5f604051808303815f87803b158015620000f8575f80fd5b505af11580156200010b573d5f803e3d5ffd5b505050508073ffffffffffffffffffffffffffffffffffffffff1663f2fde38b336040518263ffffffff1660e01b81526004016200014a91906200046a565b5f604051808303815f87803b15801562000162575f80fd5b505af115801562000175573d5f803e3d5ffd5b505050503373ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f85e892981b234101136bc30081e0a5c44345bebc0940193230c20a43b279e2d160405160405180910390a3809150509392505050565b6118e8806200056483390190565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6200024d8262000205565b810181811067ffffffffffffffff821117156200026f576200026e62000215565b5b80604052505050565b5f62000283620001ec565b905062000291828262000242565b919050565b5f67ffffffffffffffff821115620002b357620002b262000215565b5b620002be8262000205565b9050602081019050919050565b828183375f83830152505050565b5f620002ef620002e98462000296565b62000278565b9050828152602081018484840111156200030e576200030d62000201565b5b6200031b848285620002cb565b509392505050565b5f82601f8301126200033a5762000339620001fd565b5b81356200034c848260208601620002d9565b91505092915050565b5f60ff82169050919050565b6200036c8162000355565b811462000377575f80fd5b50565b5f813590506200038a8162000361565b92915050565b5f805f60608486031215620003aa57620003a9620001f5565b5b5f84013567ffffffffffffffff811115620003ca57620003c9620001f9565b5b620003d88682870162000323565b935050602084013567ffffffffffffffff811115620003fc57620003fb620001f9565b5b6200040a8682870162000323565b92505060406200041d868287016200037a565b9150509250925092565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f620004528262000427565b9050919050565b620004648162000446565b82525050565b5f6020820190506200047f5f83018462000459565b92915050565b5f81519050919050565b5f82825260208201905092915050565b5f5b83811015620004be578082015181840152602081019050620004a1565b5f8484015250505050565b5f620004d58262000485565b620004e181856200048f565b9350620004f38185602086016200049f565b620004fe8162000205565b840191505092915050565b620005148162000355565b82525050565b5f6060820190508181035f830152620005348186620004c9565b905081810360208301526200054a8185620004c9565b90506200055b604083018462000509565b94935050505056fe608060405234801562000010575f80fd5b50604051620018e8380380620018e883398181016040528101906200003691906200027c565b335f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555082600390816200008691906200054a565b5081600490816200009891906200054a565b508060055f6101000a81548160ff021916908360ff1602179055505050506200062e565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6200011d82620000d5565b810181811067ffffffffffffffff821117156200013f576200013e620000e5565b5b80604052505050565b5f62000153620000bc565b905062000161828262000112565b919050565b5f67ffffffffffffffff821115620001835762000182620000e5565b5b6200018e82620000d5565b9050602081019050919050565b5f5b83811015620001ba5780820151818401526020810190506200019d565b5f8484015250505050565b5f620001db620001d58462000166565b62000148565b905082815260208101848484011115620001fa57620001f9620000d1565b5b620002078482856200019b565b509392505050565b5f82601f830112620002265762000225620000cd565b5b815162000238848260208601620001c5565b91505092915050565b5f60ff82169050919050565b620002588162000241565b811462000263575f80fd5b50565b5f8151905062000276816200024d565b92915050565b5f805f60608486031215620002965762000295620000c5565b5b5f84015167ffffffffffffffff811115620002b657620002b5620000c9565b5b620002c4868287016200020f565b935050602084015167ffffffffffffffff811115620002e857620002e7620000c9565b5b620002f6868287016200020f565b9250506040620003098682870162000266565b9150509250925092565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806200036257607f821691505b6020821081036200037857620003776200031d565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302620003dc7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826200039f565b620003e886836200039f565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f620004326200042c620004268462000400565b62000409565b62000400565b9050919050565b5f819050919050565b6200044d8362000412565b620004656200045c8262000439565b848454620003ab565b825550505050565b5f90565b6200047b6200046d565b6200048881848462000442565b505050565b5b81811015620004af57620004a35f8262000471565b6001810190506200048e565b5050565b601f821115620004fe57620004c8816200037e565b620004d38462000390565b81016020851015620004e3578190505b620004fb620004f28562000390565b8301826200048d565b50505b505050565b5f82821c905092915050565b5f620005205f198460080262000503565b1980831691505092915050565b5f6200053a83836200050f565b9150826002028217905092915050565b620005558262000313565b67ffffffffffffffff811115620005715762000570620000e5565b5b6200057d82546200034a565b6200058a828285620004b3565b5f60209050601f831160018114620005c0575f8415620005ab578287015190505b620005b785826200052d565b86555062000626565b601f198416620005d0866200037e565b5f5b82811015620005f957848901518255600182019150602085019450602081019050620005d2565b8683101562000619578489015162000615601f8916826200050f565b8355505b6001600288020188555050505b505050505050565b6112ac806200063c5f395ff3fe608060405234801561000f575f80fd5b50600436106100cd575f3560e01c806370a082311161008a5780639dc29fac116100645780639dc29fac14610213578063a9059cbb1461022f578063dd62ed3e1461025f578063f2fde38b1461028f576100cd565b806370a08231146101a75780638da5cb5b146101d757806395d89b41146101f5576100cd565b806306fdde03146100d1578063095ea7b3146100ef57806318160ddd1461011f57806323b872dd1461013d578063313ce5671461016d57806340c10f191461018b575b5f80fd5b6100d96102ab565b6040516100e69190610ed2565b60405180910390f35b61010960048036038101906101049190610f83565b610337565b6040516101169190610fdb565b60405180910390f35b610127610424565b6040516101349190611003565b60405180910390f35b6101576004803603810190610152919061101c565b61042a565b6040516101649190610fdb565b60405180910390f35b6101756106b6565b6040516101829190611087565b60405180910390f35b6101a560048036038101906101a09190610f83565b6106c8565b005b6101c160048036038101906101bc91906110a0565b61072c565b6040516101ce9190611003565b60405180910390f35b6101df610741565b6040516101ec91906110da565b60405180910390f35b6101fd610764565b60405161020a9190610ed2565b60405180910390f35b61022d60048036038101906102289190610f83565b6107f0565b005b61024960048036038101906102449190610f83565b610854565b6040516102569190610fdb565b60405180910390f35b610279600480360381019061027491906110f3565b610a51565b6040516102869190611003565b60405180910390f35b6102a960048036038101906102a491906110a0565b610a71565b005b600380546102b89061115e565b80601f01602080910402602001604051908101604052809291908181526020018280546102e49061115e565b801561032f5780601f106103065761010080835404028352916020019161032f565b820191905f5260205f20905b81548152906001019060200180831161031257829003601f168201915b505050505081565b5f8160025f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040516104129190611003565b60405180910390a36001905092915050565b60065481565b5f8260f273ffffffffffffffffffffffffffffffffffffffff16634e25ab64826040518263ffffffff1660e01b815260040161046691906110da565b602060405180830381865afa158015610481573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906104a591906111b8565b6105115760f273ffffffffffffffffffffffffffffffffffffffff1663ceeae52a826040518263ffffffff1660e01b81526004016104e391906110da565b5f604051808303815f87803b1580156104fa575f80fd5b505af115801561050c573d5f803e3d5ffd5b505050505b8260025f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546105989190611210565b925050819055508260015f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546105eb9190611210565b925050819055508260015f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825461063e9190611243565b925050819055508373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef856040516106a29190611003565b60405180910390a360019150509392505050565b60055f9054906101000a900460ff1681565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461071e575f80fd5b6107288282610bb9565b5050565b6001602052805f5260405f205f915090505481565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600480546107719061115e565b80601f016020809104026020016040519081016040528092919081815260200182805461079d9061115e565b80156107e85780601f106107bf576101008083540402835291602001916107e8565b820191905f5260205f20905b8154815290600101906020018083116107cb57829003601f168201915b505050505081565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610846575f80fd5b6108508282610d74565b5050565b5f8260f273ffffffffffffffffffffffffffffffffffffffff16634e25ab64826040518263ffffffff1660e01b815260040161089091906110da565b602060405180830381865afa1580156108ab573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906108cf91906111b8565b61093b5760f273ffffffffffffffffffffffffffffffffffffffff1663ceeae52a826040518263ffffffff1660e01b815260040161090d91906110da565b5f604051808303815f87803b158015610924575f80fd5b505af1158015610936573d5f803e3d5ffd5b505050505b8260015f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546109879190611210565b925050819055508260015f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546109da9190611243565b925050819055508373ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef85604051610a3e9190611003565b60405180910390a3600191505092915050565b6002602052815f5260405f20602052805f5260405f205f91509150505481565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610ac7575f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610afe575f80fd5b8073ffffffffffffffffffffffffffffffffffffffff165f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3805f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b8160f273ffffffffffffffffffffffffffffffffffffffff16634e25ab64826040518263ffffffff1660e01b8152600401610bf491906110da565b602060405180830381865afa158015610c0f573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610c3391906111b8565b610c9f5760f273ffffffffffffffffffffffffffffffffffffffff1663ceeae52a826040518263ffffffff1660e01b8152600401610c7191906110da565b5f604051808303815f87803b158015610c88575f80fd5b505af1158015610c9a573d5f803e3d5ffd5b505050505b8160015f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254610ceb9190611243565b925050819055508160065f828254610d039190611243565b925050819055508273ffffffffffffffffffffffffffffffffffffffff165f73ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051610d679190611003565b60405180910390a3505050565b8060015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254610dc09190611210565b925050819055508060065f828254610dd89190611210565b925050819055505f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051610e3c9190611003565b60405180910390a35050565b5f81519050919050565b5f82825260208201905092915050565b5f5b83811015610e7f578082015181840152602081019050610e64565b5f8484015250505050565b5f601f19601f8301169050919050565b5f610ea482610e48565b610eae8185610e52565b9350610ebe818560208601610e62565b610ec781610e8a565b840191505092915050565b5f6020820190508181035f830152610eea8184610e9a565b905092915050565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610f1f82610ef6565b9050919050565b610f2f81610f15565b8114610f39575f80fd5b50565b5f81359050610f4a81610f26565b92915050565b5f819050919050565b610f6281610f50565b8114610f6c575f80fd5b50565b5f81359050610f7d81610f59565b92915050565b5f8060408385031215610f9957610f98610ef2565b5b5f610fa685828601610f3c565b9250506020610fb785828601610f6f565b9150509250929050565b5f8115159050919050565b610fd581610fc1565b82525050565b5f602082019050610fee5f830184610fcc565b92915050565b610ffd81610f50565b82525050565b5f6020820190506110165f830184610ff4565b92915050565b5f805f6060848603121561103357611032610ef2565b5b5f61104086828701610f3c565b935050602061105186828701610f3c565b925050604061106286828701610f6f565b9150509250925092565b5f60ff82169050919050565b6110818161106c565b82525050565b5f60208201905061109a5f830184611078565b92915050565b5f602082840312156110b5576110b4610ef2565b5b5f6110c284828501610f3c565b91505092915050565b6110d481610f15565b82525050565b5f6020820190506110ed5f8301846110cb565b92915050565b5f806040838503121561110957611108610ef2565b5b5f61111685828601610f3c565b925050602061112785828601610f3c565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061117557607f821691505b60208210810361118857611187611131565b5b50919050565b61119781610fc1565b81146111a1575f80fd5b50565b5f815190506111b28161118e565b92915050565b5f602082840312156111cd576111cc610ef2565b5b5f6111da848285016111a4565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61121a82610f50565b915061122583610f50565b925082820390508181111561123d5761123c6111e3565b5b92915050565b5f61124d82610f50565b915061125883610f50565b92508282019050808211156112705761126f6111e3565b5b9291505056fea264697066735822122075dfce4979d3f63ac4a6d3d47c5370ed955d4050130854ee823177a6231faf1264736f6c63430008180033a2646970667358221220d0ec928f08268f995edf07fea115f98565a542f202405a2c74f10e1d2450155064736f6c63430008180033",
}

// Erc20FactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use Erc20FactoryMetaData.ABI instead.
var Erc20FactoryABI = Erc20FactoryMetaData.ABI

// Erc20FactoryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Erc20FactoryMetaData.Bin instead.
var Erc20FactoryBin = Erc20FactoryMetaData.Bin

// DeployErc20Factory deploys a new Ethereum contract, binding an instance of Erc20Factory to it.
func DeployErc20Factory(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Erc20Factory, error) {
	parsed, err := Erc20FactoryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Erc20FactoryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Erc20Factory{Erc20FactoryCaller: Erc20FactoryCaller{contract: contract}, Erc20FactoryTransactor: Erc20FactoryTransactor{contract: contract}, Erc20FactoryFilterer: Erc20FactoryFilterer{contract: contract}}, nil
}

// Erc20Factory is an auto generated Go binding around an Ethereum contract.
type Erc20Factory struct {
	Erc20FactoryCaller     // Read-only binding to the contract
	Erc20FactoryTransactor // Write-only binding to the contract
	Erc20FactoryFilterer   // Log filterer for contract events
}

// Erc20FactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type Erc20FactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20FactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Erc20FactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20FactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Erc20FactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20FactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Erc20FactorySession struct {
	Contract     *Erc20Factory     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Erc20FactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Erc20FactoryCallerSession struct {
	Contract *Erc20FactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// Erc20FactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Erc20FactoryTransactorSession struct {
	Contract     *Erc20FactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// Erc20FactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type Erc20FactoryRaw struct {
	Contract *Erc20Factory // Generic contract binding to access the raw methods on
}

// Erc20FactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Erc20FactoryCallerRaw struct {
	Contract *Erc20FactoryCaller // Generic read-only contract binding to access the raw methods on
}

// Erc20FactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Erc20FactoryTransactorRaw struct {
	Contract *Erc20FactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewErc20Factory creates a new instance of Erc20Factory, bound to a specific deployed contract.
func NewErc20Factory(address common.Address, backend bind.ContractBackend) (*Erc20Factory, error) {
	contract, err := bindErc20Factory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Erc20Factory{Erc20FactoryCaller: Erc20FactoryCaller{contract: contract}, Erc20FactoryTransactor: Erc20FactoryTransactor{contract: contract}, Erc20FactoryFilterer: Erc20FactoryFilterer{contract: contract}}, nil
}

// NewErc20FactoryCaller creates a new read-only instance of Erc20Factory, bound to a specific deployed contract.
func NewErc20FactoryCaller(address common.Address, caller bind.ContractCaller) (*Erc20FactoryCaller, error) {
	contract, err := bindErc20Factory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Erc20FactoryCaller{contract: contract}, nil
}

// NewErc20FactoryTransactor creates a new write-only instance of Erc20Factory, bound to a specific deployed contract.
func NewErc20FactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*Erc20FactoryTransactor, error) {
	contract, err := bindErc20Factory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Erc20FactoryTransactor{contract: contract}, nil
}

// NewErc20FactoryFilterer creates a new log filterer instance of Erc20Factory, bound to a specific deployed contract.
func NewErc20FactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*Erc20FactoryFilterer, error) {
	contract, err := bindErc20Factory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Erc20FactoryFilterer{contract: contract}, nil
}

// bindErc20Factory binds a generic wrapper to an already deployed contract.
func bindErc20Factory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Erc20FactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc20Factory *Erc20FactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc20Factory.Contract.Erc20FactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc20Factory *Erc20FactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc20Factory.Contract.Erc20FactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc20Factory *Erc20FactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc20Factory.Contract.Erc20FactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc20Factory *Erc20FactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc20Factory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc20Factory *Erc20FactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc20Factory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc20Factory *Erc20FactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc20Factory.Contract.contract.Transact(opts, method, params...)
}

// CreateERC20 is a paid mutator transaction binding the contract method 0x06ef1a86.
//
// Solidity: function createERC20(string name, string symbol, uint8 decimals) returns(address)
func (_Erc20Factory *Erc20FactoryTransactor) CreateERC20(opts *bind.TransactOpts, name string, symbol string, decimals uint8) (*types.Transaction, error) {
	return _Erc20Factory.contract.Transact(opts, "createERC20", name, symbol, decimals)
}

// CreateERC20 is a paid mutator transaction binding the contract method 0x06ef1a86.
//
// Solidity: function createERC20(string name, string symbol, uint8 decimals) returns(address)
func (_Erc20Factory *Erc20FactorySession) CreateERC20(name string, symbol string, decimals uint8) (*types.Transaction, error) {
	return _Erc20Factory.Contract.CreateERC20(&_Erc20Factory.TransactOpts, name, symbol, decimals)
}

// CreateERC20 is a paid mutator transaction binding the contract method 0x06ef1a86.
//
// Solidity: function createERC20(string name, string symbol, uint8 decimals) returns(address)
func (_Erc20Factory *Erc20FactoryTransactorSession) CreateERC20(name string, symbol string, decimals uint8) (*types.Transaction, error) {
	return _Erc20Factory.Contract.CreateERC20(&_Erc20Factory.TransactOpts, name, symbol, decimals)
}

// Erc20FactoryERC20CreatedIterator is returned from FilterERC20Created and is used to iterate over the raw logs and unpacked data for ERC20Created events raised by the Erc20Factory contract.
type Erc20FactoryERC20CreatedIterator struct {
	Event *Erc20FactoryERC20Created // Event containing the contract specifics and raw log

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
func (it *Erc20FactoryERC20CreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc20FactoryERC20Created)
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
		it.Event = new(Erc20FactoryERC20Created)
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
func (it *Erc20FactoryERC20CreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc20FactoryERC20CreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc20FactoryERC20Created represents a ERC20Created event raised by the Erc20Factory contract.
type Erc20FactoryERC20Created struct {
	Erc20 common.Address
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterERC20Created is a free log retrieval operation binding the contract event 0x85e892981b234101136bc30081e0a5c44345bebc0940193230c20a43b279e2d1.
//
// Solidity: event ERC20Created(address indexed erc20, address indexed owner)
func (_Erc20Factory *Erc20FactoryFilterer) FilterERC20Created(opts *bind.FilterOpts, erc20 []common.Address, owner []common.Address) (*Erc20FactoryERC20CreatedIterator, error) {

	var erc20Rule []interface{}
	for _, erc20Item := range erc20 {
		erc20Rule = append(erc20Rule, erc20Item)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Erc20Factory.contract.FilterLogs(opts, "ERC20Created", erc20Rule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &Erc20FactoryERC20CreatedIterator{contract: _Erc20Factory.contract, event: "ERC20Created", logs: logs, sub: sub}, nil
}

// WatchERC20Created is a free log subscription operation binding the contract event 0x85e892981b234101136bc30081e0a5c44345bebc0940193230c20a43b279e2d1.
//
// Solidity: event ERC20Created(address indexed erc20, address indexed owner)
func (_Erc20Factory *Erc20FactoryFilterer) WatchERC20Created(opts *bind.WatchOpts, sink chan<- *Erc20FactoryERC20Created, erc20 []common.Address, owner []common.Address) (event.Subscription, error) {

	var erc20Rule []interface{}
	for _, erc20Item := range erc20 {
		erc20Rule = append(erc20Rule, erc20Item)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Erc20Factory.contract.WatchLogs(opts, "ERC20Created", erc20Rule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc20FactoryERC20Created)
				if err := _Erc20Factory.contract.UnpackLog(event, "ERC20Created", log); err != nil {
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

// ParseERC20Created is a log parse operation binding the contract event 0x85e892981b234101136bc30081e0a5c44345bebc0940193230c20a43b279e2d1.
//
// Solidity: event ERC20Created(address indexed erc20, address indexed owner)
func (_Erc20Factory *Erc20FactoryFilterer) ParseERC20Created(log types.Log) (*Erc20FactoryERC20Created, error) {
	event := new(Erc20FactoryERC20Created)
	if err := _Erc20Factory.contract.UnpackLog(event, "ERC20Created", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
