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
	Bin: "0x608060405234801561000f575f80fd5b506120e98061001d5f395ff3fe608060405234801562000010575f80fd5b50600436106200002c575f3560e01c806306ef1a861462000030575b5f80fd5b6200004e6004803603810190620000489190620003a5565b62000066565b6040516200005d91906200047f565b60405180910390f35b5f808484846040516200007990620001f3565b62000087939291906200052f565b604051809103905ff080158015620000a1573d5f803e3d5ffd5b50905060f273ffffffffffffffffffffffffffffffffffffffff1663d126274a826040518263ffffffff1660e01b8152600401620000e091906200047f565b6020604051808303815f875af1158015620000fd573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190620001239190620005b2565b508073ffffffffffffffffffffffffffffffffffffffff1663f2fde38b336040518263ffffffff1660e01b81526004016200015f91906200047f565b5f604051808303815f87803b15801562000177575f80fd5b505af11580156200018a573d5f803e3d5ffd5b505050503373ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f85e892981b234101136bc30081e0a5c44345bebc0940193230c20a43b279e2d160405160405180910390a3809150509392505050565b611ad180620005e383390190565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b62000262826200021a565b810181811067ffffffffffffffff821117156200028457620002836200022a565b5b80604052505050565b5f6200029862000201565b9050620002a6828262000257565b919050565b5f67ffffffffffffffff821115620002c857620002c76200022a565b5b620002d3826200021a565b9050602081019050919050565b828183375f83830152505050565b5f62000304620002fe84620002ab565b6200028d565b90508281526020810184848401111562000323576200032262000216565b5b62000330848285620002e0565b509392505050565b5f82601f8301126200034f576200034e62000212565b5b813562000361848260208601620002ee565b91505092915050565b5f60ff82169050919050565b62000381816200036a565b81146200038c575f80fd5b50565b5f813590506200039f8162000376565b92915050565b5f805f60608486031215620003bf57620003be6200020a565b5b5f84013567ffffffffffffffff811115620003df57620003de6200020e565b5b620003ed8682870162000338565b935050602084013567ffffffffffffffff8111156200041157620004106200020e565b5b6200041f8682870162000338565b925050604062000432868287016200038f565b9150509250925092565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f62000467826200043c565b9050919050565b62000479816200045b565b82525050565b5f602082019050620004945f8301846200046e565b92915050565b5f81519050919050565b5f82825260208201905092915050565b5f5b83811015620004d3578082015181840152602081019050620004b6565b5f8484015250505050565b5f620004ea826200049a565b620004f68185620004a4565b935062000508818560208601620004b4565b62000513816200021a565b840191505092915050565b62000529816200036a565b82525050565b5f6060820190508181035f830152620005498186620004de565b905081810360208301526200055f8185620004de565b90506200057060408301846200051e565b949350505050565b5f8115159050919050565b6200058e8162000578565b811462000599575f80fd5b50565b5f81519050620005ac8162000583565b92915050565b5f60208284031215620005ca57620005c96200020a565b5b5f620005d9848285016200059c565b9150509291505056fe608060405234801562000010575f80fd5b5060405162001ad138038062001ad183398181016040528101906200003691906200027c565b335f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555082600390816200008691906200054a565b5081600490816200009891906200054a565b508060055f6101000a81548160ff021916908360ff1602179055505050506200062e565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6200011d82620000d5565b810181811067ffffffffffffffff821117156200013f576200013e620000e5565b5b80604052505050565b5f62000153620000bc565b905062000161828262000112565b919050565b5f67ffffffffffffffff821115620001835762000182620000e5565b5b6200018e82620000d5565b9050602081019050919050565b5f5b83811015620001ba5780820151818401526020810190506200019d565b5f8484015250505050565b5f620001db620001d58462000166565b62000148565b905082815260208101848484011115620001fa57620001f9620000d1565b5b620002078482856200019b565b509392505050565b5f82601f830112620002265762000225620000cd565b5b815162000238848260208601620001c5565b91505092915050565b5f60ff82169050919050565b620002588162000241565b811462000263575f80fd5b50565b5f8151905062000276816200024d565b92915050565b5f805f60608486031215620002965762000295620000c5565b5b5f84015167ffffffffffffffff811115620002b657620002b5620000c9565b5b620002c4868287016200020f565b935050602084015167ffffffffffffffff811115620002e857620002e7620000c9565b5b620002f6868287016200020f565b9250506040620003098682870162000266565b9150509250925092565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806200036257607f821691505b6020821081036200037857620003776200031d565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302620003dc7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826200039f565b620003e886836200039f565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f620004326200042c620004268462000400565b62000409565b62000400565b9050919050565b5f819050919050565b6200044d8362000412565b620004656200045c8262000439565b848454620003ab565b825550505050565b5f90565b6200047b6200046d565b6200048881848462000442565b505050565b5b81811015620004af57620004a35f8262000471565b6001810190506200048e565b5050565b601f821115620004fe57620004c8816200037e565b620004d38462000390565b81016020851015620004e3578190505b620004fb620004f28562000390565b8301826200048d565b50505b505050565b5f82821c905092915050565b5f620005205f198460080262000503565b1980831691505092915050565b5f6200053a83836200050f565b9150826002028217905092915050565b620005558262000313565b67ffffffffffffffff811115620005715762000570620000e5565b5b6200057d82546200034a565b6200058a828285620004b3565b5f60209050601f831160018114620005c0575f8415620005ab578287015190505b620005b785826200052d565b86555062000626565b601f198416620005d0866200037e565b5f5b82811015620005f957848901518255600182019150602085019450602081019050620005d2565b8683101562000619578489015162000615601f8916826200050f565b8355505b6001600288020188555050505b505050505050565b611495806200063c5f395ff3fe608060405234801561000f575f80fd5b50600436106100e8575f3560e01c806370a082311161008a5780639dc29fac116100645780639dc29fac1461025e578063a9059cbb1461027a578063dd62ed3e146102aa578063f2fde38b146102da576100e8565b806370a08231146101f25780638da5cb5b1461022257806395d89b4114610240576100e8565b806318160ddd116100c657806318160ddd1461016a57806323b872dd14610188578063313ce567146101b857806340c10f19146101d6576100e8565b806301ffc9a7146100ec57806306fdde031461011c578063095ea7b31461013a575b5f80fd5b6101066004803603810190610101919061100a565b6102f6565b604051610113919061104f565b60405180910390f35b61012461036f565b60405161013191906110f2565b60405180910390f35b610154600480360381019061014f919061119f565b6103fb565b604051610161919061104f565b60405180910390f35b6101726104e8565b60405161017f91906111ec565b60405180910390f35b6101a2600480360381019061019d9190611205565b6104ee565b6040516101af919061104f565b60405180910390f35b6101c061078e565b6040516101cd9190611270565b60405180910390f35b6101f060048036038101906101eb919061119f565b6107a0565b005b61020c60048036038101906102079190611289565b610804565b60405161021991906111ec565b60405180910390f35b61022a610819565b60405161023791906112c3565b60405180910390f35b61024861083c565b60405161025591906110f2565b60405180910390f35b6102786004803603810190610273919061119f565b6108c8565b005b610294600480360381019061028f919061119f565b61092c565b6040516102a1919061104f565b60405180910390f35b6102c460048036038101906102bf91906112dc565b610b3d565b6040516102d191906111ec565b60405180910390f35b6102f460048036038101906102ef9190611289565b610b5d565b005b5f7f942e8b22000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19161480610368575061036782610ca5565b5b9050919050565b6003805461037c90611347565b80601f01602080910402602001604051908101604052809291908181526020018280546103a890611347565b80156103f35780601f106103ca576101008083540402835291602001916103f3565b820191905f5260205f20905b8154815290600101906020018083116103d657829003601f168201915b505050505081565b5f8160025f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040516104d691906111ec565b60405180910390a36001905092915050565b60065481565b5f8260f273ffffffffffffffffffffffffffffffffffffffff16634e25ab64826040518263ffffffff1660e01b815260040161052a91906112c3565b602060405180830381865afa158015610545573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061056991906113a1565b6105e95760f273ffffffffffffffffffffffffffffffffffffffff1663ceeae52a826040518263ffffffff1660e01b81526004016105a791906112c3565b6020604051808303815f875af11580156105c3573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906105e791906113a1565b505b8260025f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825461067091906113f9565b925050819055508260015f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546106c391906113f9565b925050819055508260015f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254610716919061142c565b925050819055508373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8560405161077a91906111ec565b60405180910390a360019150509392505050565b60055f9054906101000a900460ff1681565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146107f6575f80fd5b6108008282610d0e565b5050565b6001602052805f5260405f205f915090505481565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6004805461084990611347565b80601f016020809104026020016040519081016040528092919081815260200182805461087590611347565b80156108c05780601f10610897576101008083540402835291602001916108c0565b820191905f5260205f20905b8154815290600101906020018083116108a357829003601f168201915b505050505081565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461091e575f80fd5b6109288282610edd565b5050565b5f8260f273ffffffffffffffffffffffffffffffffffffffff16634e25ab64826040518263ffffffff1660e01b815260040161096891906112c3565b602060405180830381865afa158015610983573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906109a791906113a1565b610a275760f273ffffffffffffffffffffffffffffffffffffffff1663ceeae52a826040518263ffffffff1660e01b81526004016109e591906112c3565b6020604051808303815f875af1158015610a01573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610a2591906113a1565b505b8260015f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254610a7391906113f9565b925050819055508260015f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254610ac6919061142c565b925050819055508373ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef85604051610b2a91906111ec565b60405180910390a3600191505092915050565b6002602052815f5260405f20602052805f5260405f205f91509150505481565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610bb3575f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610bea575f80fd5b8073ffffffffffffffffffffffffffffffffffffffff165f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3805f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b5f7f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b8160f273ffffffffffffffffffffffffffffffffffffffff16634e25ab64826040518263ffffffff1660e01b8152600401610d4991906112c3565b602060405180830381865afa158015610d64573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610d8891906113a1565b610e085760f273ffffffffffffffffffffffffffffffffffffffff1663ceeae52a826040518263ffffffff1660e01b8152600401610dc691906112c3565b6020604051808303815f875af1158015610de2573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610e0691906113a1565b505b8160015f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254610e54919061142c565b925050819055508160065f828254610e6c919061142c565b925050819055508273ffffffffffffffffffffffffffffffffffffffff165f73ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051610ed091906111ec565b60405180910390a3505050565b8060015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254610f2991906113f9565b925050819055508060065f828254610f4191906113f9565b925050819055505f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051610fa591906111ec565b60405180910390a35050565b5f80fd5b5f7fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b610fe981610fb5565b8114610ff3575f80fd5b50565b5f8135905061100481610fe0565b92915050565b5f6020828403121561101f5761101e610fb1565b5b5f61102c84828501610ff6565b91505092915050565b5f8115159050919050565b61104981611035565b82525050565b5f6020820190506110625f830184611040565b92915050565b5f81519050919050565b5f82825260208201905092915050565b5f5b8381101561109f578082015181840152602081019050611084565b5f8484015250505050565b5f601f19601f8301169050919050565b5f6110c482611068565b6110ce8185611072565b93506110de818560208601611082565b6110e7816110aa565b840191505092915050565b5f6020820190508181035f83015261110a81846110ba565b905092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61113b82611112565b9050919050565b61114b81611131565b8114611155575f80fd5b50565b5f8135905061116681611142565b92915050565b5f819050919050565b61117e8161116c565b8114611188575f80fd5b50565b5f8135905061119981611175565b92915050565b5f80604083850312156111b5576111b4610fb1565b5b5f6111c285828601611158565b92505060206111d38582860161118b565b9150509250929050565b6111e68161116c565b82525050565b5f6020820190506111ff5f8301846111dd565b92915050565b5f805f6060848603121561121c5761121b610fb1565b5b5f61122986828701611158565b935050602061123a86828701611158565b925050604061124b8682870161118b565b9150509250925092565b5f60ff82169050919050565b61126a81611255565b82525050565b5f6020820190506112835f830184611261565b92915050565b5f6020828403121561129e5761129d610fb1565b5b5f6112ab84828501611158565b91505092915050565b6112bd81611131565b82525050565b5f6020820190506112d65f8301846112b4565b92915050565b5f80604083850312156112f2576112f1610fb1565b5b5f6112ff85828601611158565b925050602061131085828601611158565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061135e57607f821691505b6020821081036113715761137061131a565b5b50919050565b61138081611035565b811461138a575f80fd5b50565b5f8151905061139b81611377565b92915050565b5f602082840312156113b6576113b5610fb1565b5b5f6113c38482850161138d565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6114038261116c565b915061140e8361116c565b9250828203905081811115611426576114256113cc565b5b92915050565b5f6114368261116c565b91506114418361116c565b9250828201905080821115611459576114586113cc565b5b9291505056fea264697066735822122067e09fb1c8f6c6d1f6c87e9baac07eeef8b81a3218e526c7e72f157bf344e58564736f6c63430008180033a2646970667358221220012e95eae571f8b8aa7d90dba694698d6bbe3e987c2b4c972f8e03c0a07e47fe64736f6c63430008180033",
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
