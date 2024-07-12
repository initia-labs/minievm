// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package custom_erc20

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

// CustomErc20MetaData contains all meta data concerning the CustomErc20 contract.
var CustomErc20MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"_decimals\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801562000010575f80fd5b5060405162001bae38038062001bae8339818101604052810190620000369190620002ef565b335f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060f273ffffffffffffffffffffffffffffffffffffffff16635e6c57596040518163ffffffff1660e01b81526004016020604051808303815f875af1158015620000c1573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190620000e79190620003c0565b508260039081620000f9919062000627565b5081600490816200010b919062000627565b508060055f6101000a81548160ff021916908360ff1602179055505050506200070b565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b620001908262000148565b810181811067ffffffffffffffff82111715620001b257620001b162000158565b5b80604052505050565b5f620001c66200012f565b9050620001d4828262000185565b919050565b5f67ffffffffffffffff821115620001f657620001f562000158565b5b620002018262000148565b9050602081019050919050565b5f5b838110156200022d57808201518184015260208101905062000210565b5f8484015250505050565b5f6200024e6200024884620001d9565b620001bb565b9050828152602081018484840111156200026d576200026c62000144565b5b6200027a8482856200020e565b509392505050565b5f82601f83011262000299576200029862000140565b5b8151620002ab84826020860162000238565b91505092915050565b5f60ff82169050919050565b620002cb81620002b4565b8114620002d6575f80fd5b50565b5f81519050620002e981620002c0565b92915050565b5f805f6060848603121562000309576200030862000138565b5b5f84015167ffffffffffffffff8111156200032957620003286200013c565b5b620003378682870162000282565b935050602084015167ffffffffffffffff8111156200035b576200035a6200013c565b5b620003698682870162000282565b92505060406200037c86828701620002d9565b9150509250925092565b5f8115159050919050565b6200039c8162000386565b8114620003a7575f80fd5b50565b5f81519050620003ba8162000391565b92915050565b5f60208284031215620003d857620003d762000138565b5b5f620003e784828501620003aa565b91505092915050565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806200043f57607f821691505b602082108103620004555762000454620003fa565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302620004b97fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826200047c565b620004c586836200047c565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f6200050f620005096200050384620004dd565b620004e6565b620004dd565b9050919050565b5f819050919050565b6200052a83620004ef565b62000542620005398262000516565b84845462000488565b825550505050565b5f90565b620005586200054a565b620005658184846200051f565b505050565b5b818110156200058c57620005805f826200054e565b6001810190506200056b565b5050565b601f821115620005db57620005a5816200045b565b620005b0846200046d565b81016020851015620005c0578190505b620005d8620005cf856200046d565b8301826200056a565b50505b505050565b5f82821c905092915050565b5f620005fd5f1984600802620005e0565b1980831691505092915050565b5f620006178383620005ec565b9150826002028217905092915050565b6200063282620003f0565b67ffffffffffffffff8111156200064e576200064d62000158565b5b6200065a825462000427565b6200066782828562000590565b5f60209050601f8311600181146200069d575f841562000688578287015190505b6200069485826200060a565b86555062000703565b601f198416620006ad866200045b565b5f5b82811015620006d657848901518255600182019150602085019450602081019050620006af565b86831015620006f65784890151620006f2601f891682620005ec565b8355505b6001600288020188555050505b505050505050565b61149580620007195f395ff3fe608060405234801561000f575f80fd5b50600436106100e8575f3560e01c806370a082311161008a5780639dc29fac116100645780639dc29fac1461025e578063a9059cbb1461027a578063dd62ed3e146102aa578063f2fde38b146102da576100e8565b806370a08231146101f25780638da5cb5b1461022257806395d89b4114610240576100e8565b806318160ddd116100c657806318160ddd1461016a57806323b872dd14610188578063313ce567146101b857806340c10f19146101d6576100e8565b806301ffc9a7146100ec57806306fdde031461011c578063095ea7b31461013a575b5f80fd5b6101066004803603810190610101919061100a565b6102f6565b604051610113919061104f565b60405180910390f35b61012461036f565b60405161013191906110f2565b60405180910390f35b610154600480360381019061014f919061119f565b6103fb565b604051610161919061104f565b60405180910390f35b6101726104e8565b60405161017f91906111ec565b60405180910390f35b6101a2600480360381019061019d9190611205565b6104ee565b6040516101af919061104f565b60405180910390f35b6101c061078e565b6040516101cd9190611270565b60405180910390f35b6101f060048036038101906101eb919061119f565b6107a0565b005b61020c60048036038101906102079190611289565b610804565b60405161021991906111ec565b60405180910390f35b61022a610819565b60405161023791906112c3565b60405180910390f35b61024861083c565b60405161025591906110f2565b60405180910390f35b6102786004803603810190610273919061119f565b6108c8565b005b610294600480360381019061028f919061119f565b61092c565b6040516102a1919061104f565b60405180910390f35b6102c460048036038101906102bf91906112dc565b610b3d565b6040516102d191906111ec565b60405180910390f35b6102f460048036038101906102ef9190611289565b610b5d565b005b5f7f942e8b22000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19161480610368575061036782610ca5565b5b9050919050565b6003805461037c90611347565b80601f01602080910402602001604051908101604052809291908181526020018280546103a890611347565b80156103f35780601f106103ca576101008083540402835291602001916103f3565b820191905f5260205f20905b8154815290600101906020018083116103d657829003601f168201915b505050505081565b5f8160025f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040516104d691906111ec565b60405180910390a36001905092915050565b60065481565b5f8260f273ffffffffffffffffffffffffffffffffffffffff16634e25ab64826040518263ffffffff1660e01b815260040161052a91906112c3565b602060405180830381865afa158015610545573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061056991906113a1565b6105e95760f273ffffffffffffffffffffffffffffffffffffffff1663ceeae52a826040518263ffffffff1660e01b81526004016105a791906112c3565b6020604051808303815f875af11580156105c3573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906105e791906113a1565b505b8260025f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825461067091906113f9565b925050819055508260015f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546106c391906113f9565b925050819055508260015f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254610716919061142c565b925050819055508373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8560405161077a91906111ec565b60405180910390a360019150509392505050565b60055f9054906101000a900460ff1681565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146107f6575f80fd5b6108008282610d0e565b5050565b6001602052805f5260405f205f915090505481565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6004805461084990611347565b80601f016020809104026020016040519081016040528092919081815260200182805461087590611347565b80156108c05780601f10610897576101008083540402835291602001916108c0565b820191905f5260205f20905b8154815290600101906020018083116108a357829003601f168201915b505050505081565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461091e575f80fd5b6109288282610edd565b5050565b5f8260f273ffffffffffffffffffffffffffffffffffffffff16634e25ab64826040518263ffffffff1660e01b815260040161096891906112c3565b602060405180830381865afa158015610983573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906109a791906113a1565b610a275760f273ffffffffffffffffffffffffffffffffffffffff1663ceeae52a826040518263ffffffff1660e01b81526004016109e591906112c3565b6020604051808303815f875af1158015610a01573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610a2591906113a1565b505b8260015f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254610a7391906113f9565b925050819055508260015f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254610ac6919061142c565b925050819055508373ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef85604051610b2a91906111ec565b60405180910390a3600191505092915050565b6002602052815f5260405f20602052805f5260405f205f91509150505481565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610bb3575f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610bea575f80fd5b8073ffffffffffffffffffffffffffffffffffffffff165f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3805f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b5f7f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b8160f273ffffffffffffffffffffffffffffffffffffffff16634e25ab64826040518263ffffffff1660e01b8152600401610d4991906112c3565b602060405180830381865afa158015610d64573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610d8891906113a1565b610e085760f273ffffffffffffffffffffffffffffffffffffffff1663ceeae52a826040518263ffffffff1660e01b8152600401610dc691906112c3565b6020604051808303815f875af1158015610de2573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610e0691906113a1565b505b8160015f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254610e54919061142c565b925050819055508160065f828254610e6c919061142c565b925050819055508273ffffffffffffffffffffffffffffffffffffffff165f73ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051610ed091906111ec565b60405180910390a3505050565b8060015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254610f2991906113f9565b925050819055508060065f828254610f4191906113f9565b925050819055505f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051610fa591906111ec565b60405180910390a35050565b5f80fd5b5f7fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b610fe981610fb5565b8114610ff3575f80fd5b50565b5f8135905061100481610fe0565b92915050565b5f6020828403121561101f5761101e610fb1565b5b5f61102c84828501610ff6565b91505092915050565b5f8115159050919050565b61104981611035565b82525050565b5f6020820190506110625f830184611040565b92915050565b5f81519050919050565b5f82825260208201905092915050565b5f5b8381101561109f578082015181840152602081019050611084565b5f8484015250505050565b5f601f19601f8301169050919050565b5f6110c482611068565b6110ce8185611072565b93506110de818560208601611082565b6110e7816110aa565b840191505092915050565b5f6020820190508181035f83015261110a81846110ba565b905092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61113b82611112565b9050919050565b61114b81611131565b8114611155575f80fd5b50565b5f8135905061116681611142565b92915050565b5f819050919050565b61117e8161116c565b8114611188575f80fd5b50565b5f8135905061119981611175565b92915050565b5f80604083850312156111b5576111b4610fb1565b5b5f6111c285828601611158565b92505060206111d38582860161118b565b9150509250929050565b6111e68161116c565b82525050565b5f6020820190506111ff5f8301846111dd565b92915050565b5f805f6060848603121561121c5761121b610fb1565b5b5f61122986828701611158565b935050602061123a86828701611158565b925050604061124b8682870161118b565b9150509250925092565b5f60ff82169050919050565b61126a81611255565b82525050565b5f6020820190506112835f830184611261565b92915050565b5f6020828403121561129e5761129d610fb1565b5b5f6112ab84828501611158565b91505092915050565b6112bd81611131565b82525050565b5f6020820190506112d65f8301846112b4565b92915050565b5f80604083850312156112f2576112f1610fb1565b5b5f6112ff85828601611158565b925050602061131085828601611158565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061135e57607f821691505b6020821081036113715761137061131a565b5b50919050565b61138081611035565b811461138a575f80fd5b50565b5f8151905061139b81611377565b92915050565b5f602082840312156113b6576113b5610fb1565b5b5f6113c38482850161138d565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6114038261116c565b915061140e8361116c565b9250828203905081811115611426576114256113cc565b5b92915050565b5f6114368261116c565b91506114418361116c565b9250828201905080821115611459576114586113cc565b5b9291505056fea2646970667358221220b4dd1d8802dffa0485bc89cd679dc48dda5cbc64073a4cd2bcc4ac0dc6cf8c8f64736f6c63430008180033",
}

// CustomErc20ABI is the input ABI used to generate the binding from.
// Deprecated: Use CustomErc20MetaData.ABI instead.
var CustomErc20ABI = CustomErc20MetaData.ABI

// CustomErc20Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CustomErc20MetaData.Bin instead.
var CustomErc20Bin = CustomErc20MetaData.Bin

// DeployCustomErc20 deploys a new Ethereum contract, binding an instance of CustomErc20 to it.
func DeployCustomErc20(auth *bind.TransactOpts, backend bind.ContractBackend, _name string, _symbol string, _decimals uint8) (common.Address, *types.Transaction, *CustomErc20, error) {
	parsed, err := CustomErc20MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CustomErc20Bin), backend, _name, _symbol, _decimals)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CustomErc20{CustomErc20Caller: CustomErc20Caller{contract: contract}, CustomErc20Transactor: CustomErc20Transactor{contract: contract}, CustomErc20Filterer: CustomErc20Filterer{contract: contract}}, nil
}

// CustomErc20 is an auto generated Go binding around an Ethereum contract.
type CustomErc20 struct {
	CustomErc20Caller     // Read-only binding to the contract
	CustomErc20Transactor // Write-only binding to the contract
	CustomErc20Filterer   // Log filterer for contract events
}

// CustomErc20Caller is an auto generated read-only Go binding around an Ethereum contract.
type CustomErc20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CustomErc20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type CustomErc20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CustomErc20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CustomErc20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CustomErc20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CustomErc20Session struct {
	Contract     *CustomErc20      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CustomErc20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CustomErc20CallerSession struct {
	Contract *CustomErc20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// CustomErc20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CustomErc20TransactorSession struct {
	Contract     *CustomErc20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// CustomErc20Raw is an auto generated low-level Go binding around an Ethereum contract.
type CustomErc20Raw struct {
	Contract *CustomErc20 // Generic contract binding to access the raw methods on
}

// CustomErc20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CustomErc20CallerRaw struct {
	Contract *CustomErc20Caller // Generic read-only contract binding to access the raw methods on
}

// CustomErc20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CustomErc20TransactorRaw struct {
	Contract *CustomErc20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewCustomErc20 creates a new instance of CustomErc20, bound to a specific deployed contract.
func NewCustomErc20(address common.Address, backend bind.ContractBackend) (*CustomErc20, error) {
	contract, err := bindCustomErc20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CustomErc20{CustomErc20Caller: CustomErc20Caller{contract: contract}, CustomErc20Transactor: CustomErc20Transactor{contract: contract}, CustomErc20Filterer: CustomErc20Filterer{contract: contract}}, nil
}

// NewCustomErc20Caller creates a new read-only instance of CustomErc20, bound to a specific deployed contract.
func NewCustomErc20Caller(address common.Address, caller bind.ContractCaller) (*CustomErc20Caller, error) {
	contract, err := bindCustomErc20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CustomErc20Caller{contract: contract}, nil
}

// NewCustomErc20Transactor creates a new write-only instance of CustomErc20, bound to a specific deployed contract.
func NewCustomErc20Transactor(address common.Address, transactor bind.ContractTransactor) (*CustomErc20Transactor, error) {
	contract, err := bindCustomErc20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CustomErc20Transactor{contract: contract}, nil
}

// NewCustomErc20Filterer creates a new log filterer instance of CustomErc20, bound to a specific deployed contract.
func NewCustomErc20Filterer(address common.Address, filterer bind.ContractFilterer) (*CustomErc20Filterer, error) {
	contract, err := bindCustomErc20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CustomErc20Filterer{contract: contract}, nil
}

// bindCustomErc20 binds a generic wrapper to an already deployed contract.
func bindCustomErc20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CustomErc20MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CustomErc20 *CustomErc20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CustomErc20.Contract.CustomErc20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CustomErc20 *CustomErc20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CustomErc20.Contract.CustomErc20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CustomErc20 *CustomErc20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CustomErc20.Contract.CustomErc20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CustomErc20 *CustomErc20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CustomErc20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CustomErc20 *CustomErc20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CustomErc20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CustomErc20 *CustomErc20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CustomErc20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_CustomErc20 *CustomErc20Caller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CustomErc20.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_CustomErc20 *CustomErc20Session) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _CustomErc20.Contract.Allowance(&_CustomErc20.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_CustomErc20 *CustomErc20CallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _CustomErc20.Contract.Allowance(&_CustomErc20.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_CustomErc20 *CustomErc20Caller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CustomErc20.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_CustomErc20 *CustomErc20Session) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _CustomErc20.Contract.BalanceOf(&_CustomErc20.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_CustomErc20 *CustomErc20CallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _CustomErc20.Contract.BalanceOf(&_CustomErc20.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_CustomErc20 *CustomErc20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _CustomErc20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_CustomErc20 *CustomErc20Session) Decimals() (uint8, error) {
	return _CustomErc20.Contract.Decimals(&_CustomErc20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_CustomErc20 *CustomErc20CallerSession) Decimals() (uint8, error) {
	return _CustomErc20.Contract.Decimals(&_CustomErc20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CustomErc20 *CustomErc20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CustomErc20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CustomErc20 *CustomErc20Session) Name() (string, error) {
	return _CustomErc20.Contract.Name(&_CustomErc20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CustomErc20 *CustomErc20CallerSession) Name() (string, error) {
	return _CustomErc20.Contract.Name(&_CustomErc20.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CustomErc20 *CustomErc20Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CustomErc20.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CustomErc20 *CustomErc20Session) Owner() (common.Address, error) {
	return _CustomErc20.Contract.Owner(&_CustomErc20.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CustomErc20 *CustomErc20CallerSession) Owner() (common.Address, error) {
	return _CustomErc20.Contract.Owner(&_CustomErc20.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CustomErc20 *CustomErc20Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _CustomErc20.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CustomErc20 *CustomErc20Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CustomErc20.Contract.SupportsInterface(&_CustomErc20.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CustomErc20 *CustomErc20CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CustomErc20.Contract.SupportsInterface(&_CustomErc20.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CustomErc20 *CustomErc20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CustomErc20.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CustomErc20 *CustomErc20Session) Symbol() (string, error) {
	return _CustomErc20.Contract.Symbol(&_CustomErc20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CustomErc20 *CustomErc20CallerSession) Symbol() (string, error) {
	return _CustomErc20.Contract.Symbol(&_CustomErc20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CustomErc20 *CustomErc20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CustomErc20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CustomErc20 *CustomErc20Session) TotalSupply() (*big.Int, error) {
	return _CustomErc20.Contract.TotalSupply(&_CustomErc20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CustomErc20 *CustomErc20CallerSession) TotalSupply() (*big.Int, error) {
	return _CustomErc20.Contract.TotalSupply(&_CustomErc20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_CustomErc20 *CustomErc20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CustomErc20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_CustomErc20 *CustomErc20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CustomErc20.Contract.Approve(&_CustomErc20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_CustomErc20 *CustomErc20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CustomErc20.Contract.Approve(&_CustomErc20.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address from, uint256 amount) returns()
func (_CustomErc20 *CustomErc20Transactor) Burn(opts *bind.TransactOpts, from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CustomErc20.contract.Transact(opts, "burn", from, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address from, uint256 amount) returns()
func (_CustomErc20 *CustomErc20Session) Burn(from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CustomErc20.Contract.Burn(&_CustomErc20.TransactOpts, from, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address from, uint256 amount) returns()
func (_CustomErc20 *CustomErc20TransactorSession) Burn(from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CustomErc20.Contract.Burn(&_CustomErc20.TransactOpts, from, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_CustomErc20 *CustomErc20Transactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CustomErc20.contract.Transact(opts, "mint", to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_CustomErc20 *CustomErc20Session) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CustomErc20.Contract.Mint(&_CustomErc20.TransactOpts, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_CustomErc20 *CustomErc20TransactorSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CustomErc20.Contract.Mint(&_CustomErc20.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_CustomErc20 *CustomErc20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CustomErc20.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_CustomErc20 *CustomErc20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CustomErc20.Contract.Transfer(&_CustomErc20.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_CustomErc20 *CustomErc20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CustomErc20.Contract.Transfer(&_CustomErc20.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_CustomErc20 *CustomErc20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CustomErc20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_CustomErc20 *CustomErc20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CustomErc20.Contract.TransferFrom(&_CustomErc20.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_CustomErc20 *CustomErc20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CustomErc20.Contract.TransferFrom(&_CustomErc20.TransactOpts, sender, recipient, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CustomErc20 *CustomErc20Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _CustomErc20.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CustomErc20 *CustomErc20Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CustomErc20.Contract.TransferOwnership(&_CustomErc20.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CustomErc20 *CustomErc20TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CustomErc20.Contract.TransferOwnership(&_CustomErc20.TransactOpts, newOwner)
}

// CustomErc20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the CustomErc20 contract.
type CustomErc20ApprovalIterator struct {
	Event *CustomErc20Approval // Event containing the contract specifics and raw log

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
func (it *CustomErc20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CustomErc20Approval)
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
		it.Event = new(CustomErc20Approval)
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
func (it *CustomErc20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CustomErc20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CustomErc20Approval represents a Approval event raised by the CustomErc20 contract.
type CustomErc20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_CustomErc20 *CustomErc20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*CustomErc20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _CustomErc20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &CustomErc20ApprovalIterator{contract: _CustomErc20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_CustomErc20 *CustomErc20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *CustomErc20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _CustomErc20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CustomErc20Approval)
				if err := _CustomErc20.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_CustomErc20 *CustomErc20Filterer) ParseApproval(log types.Log) (*CustomErc20Approval, error) {
	event := new(CustomErc20Approval)
	if err := _CustomErc20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CustomErc20OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the CustomErc20 contract.
type CustomErc20OwnershipTransferredIterator struct {
	Event *CustomErc20OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CustomErc20OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CustomErc20OwnershipTransferred)
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
		it.Event = new(CustomErc20OwnershipTransferred)
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
func (it *CustomErc20OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CustomErc20OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CustomErc20OwnershipTransferred represents a OwnershipTransferred event raised by the CustomErc20 contract.
type CustomErc20OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CustomErc20 *CustomErc20Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CustomErc20OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CustomErc20.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CustomErc20OwnershipTransferredIterator{contract: _CustomErc20.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CustomErc20 *CustomErc20Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CustomErc20OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CustomErc20.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CustomErc20OwnershipTransferred)
				if err := _CustomErc20.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CustomErc20 *CustomErc20Filterer) ParseOwnershipTransferred(log types.Log) (*CustomErc20OwnershipTransferred, error) {
	event := new(CustomErc20OwnershipTransferred)
	if err := _CustomErc20.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CustomErc20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the CustomErc20 contract.
type CustomErc20TransferIterator struct {
	Event *CustomErc20Transfer // Event containing the contract specifics and raw log

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
func (it *CustomErc20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CustomErc20Transfer)
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
		it.Event = new(CustomErc20Transfer)
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
func (it *CustomErc20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CustomErc20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CustomErc20Transfer represents a Transfer event raised by the CustomErc20 contract.
type CustomErc20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_CustomErc20 *CustomErc20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CustomErc20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CustomErc20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CustomErc20TransferIterator{contract: _CustomErc20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_CustomErc20 *CustomErc20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *CustomErc20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CustomErc20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CustomErc20Transfer)
				if err := _CustomErc20.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_CustomErc20 *CustomErc20Filterer) ParseTransfer(log types.Log) (*CustomErc20Transfer, error) {
	event := new(CustomErc20Transfer)
	if err := _CustomErc20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
