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
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"_decimals\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801562000010575f80fd5b5060405162001946380380620019468339818101604052810190620000369190620002da565b335f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060f273ffffffffffffffffffffffffffffffffffffffff16635e6c57596040518163ffffffff1660e01b81526004015f604051808303815f87803b158015620000bc575f80fd5b505af1158015620000cf573d5f803e3d5ffd5b505050508260039081620000e49190620005a8565b508160049081620000f69190620005a8565b508060055f6101000a81548160ff021916908360ff1602179055505050506200068c565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6200017b8262000133565b810181811067ffffffffffffffff821117156200019d576200019c62000143565b5b80604052505050565b5f620001b16200011a565b9050620001bf828262000170565b919050565b5f67ffffffffffffffff821115620001e157620001e062000143565b5b620001ec8262000133565b9050602081019050919050565b5f5b8381101562000218578082015181840152602081019050620001fb565b5f8484015250505050565b5f620002396200023384620001c4565b620001a6565b9050828152602081018484840111156200025857620002576200012f565b5b62000265848285620001f9565b509392505050565b5f82601f8301126200028457620002836200012b565b5b81516200029684826020860162000223565b91505092915050565b5f60ff82169050919050565b620002b6816200029f565b8114620002c1575f80fd5b50565b5f81519050620002d481620002ab565b92915050565b5f805f60608486031215620002f457620002f362000123565b5b5f84015167ffffffffffffffff81111562000314576200031362000127565b5b62000322868287016200026d565b935050602084015167ffffffffffffffff81111562000346576200034562000127565b5b62000354868287016200026d565b92505060406200036786828701620002c4565b9150509250925092565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680620003c057607f821691505b602082108103620003d657620003d56200037b565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026200043a7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82620003fd565b620004468683620003fd565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f620004906200048a62000484846200045e565b62000467565b6200045e565b9050919050565b5f819050919050565b620004ab8362000470565b620004c3620004ba8262000497565b84845462000409565b825550505050565b5f90565b620004d9620004cb565b620004e6818484620004a0565b505050565b5b818110156200050d57620005015f82620004cf565b600181019050620004ec565b5050565b601f8211156200055c576200052681620003dc565b6200053184620003ee565b8101602085101562000541578190505b620005596200055085620003ee565b830182620004eb565b50505b505050565b5f82821c905092915050565b5f6200057e5f198460080262000561565b1980831691505092915050565b5f6200059883836200056d565b9150826002028217905092915050565b620005b38262000371565b67ffffffffffffffff811115620005cf57620005ce62000143565b5b620005db8254620003a8565b620005e882828562000511565b5f60209050601f8311600181146200061e575f841562000609578287015190505b6200061585826200058b565b86555062000684565b601f1984166200062e86620003dc565b5f5b82811015620006575784890151825560018201915060208501945060208101905062000630565b8683101562000677578489015162000673601f8916826200056d565b8355505b6001600288020188555050505b505050505050565b6112ac806200069a5f395ff3fe608060405234801561000f575f80fd5b50600436106100cd575f3560e01c806370a082311161008a5780639dc29fac116100645780639dc29fac14610213578063a9059cbb1461022f578063dd62ed3e1461025f578063f2fde38b1461028f576100cd565b806370a08231146101a75780638da5cb5b146101d757806395d89b41146101f5576100cd565b806306fdde03146100d1578063095ea7b3146100ef57806318160ddd1461011f57806323b872dd1461013d578063313ce5671461016d57806340c10f191461018b575b5f80fd5b6100d96102ab565b6040516100e69190610ed2565b60405180910390f35b61010960048036038101906101049190610f83565b610337565b6040516101169190610fdb565b60405180910390f35b610127610424565b6040516101349190611003565b60405180910390f35b6101576004803603810190610152919061101c565b61042a565b6040516101649190610fdb565b60405180910390f35b6101756106b6565b6040516101829190611087565b60405180910390f35b6101a560048036038101906101a09190610f83565b6106c8565b005b6101c160048036038101906101bc91906110a0565b61072c565b6040516101ce9190611003565b60405180910390f35b6101df610741565b6040516101ec91906110da565b60405180910390f35b6101fd610764565b60405161020a9190610ed2565b60405180910390f35b61022d60048036038101906102289190610f83565b6107f0565b005b61024960048036038101906102449190610f83565b610854565b6040516102569190610fdb565b60405180910390f35b610279600480360381019061027491906110f3565b610a51565b6040516102869190611003565b60405180910390f35b6102a960048036038101906102a491906110a0565b610a71565b005b600380546102b89061115e565b80601f01602080910402602001604051908101604052809291908181526020018280546102e49061115e565b801561032f5780601f106103065761010080835404028352916020019161032f565b820191905f5260205f20905b81548152906001019060200180831161031257829003601f168201915b505050505081565b5f8160025f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040516104129190611003565b60405180910390a36001905092915050565b60065481565b5f8260f273ffffffffffffffffffffffffffffffffffffffff16634e25ab64826040518263ffffffff1660e01b815260040161046691906110da565b602060405180830381865afa158015610481573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906104a591906111b8565b6105115760f273ffffffffffffffffffffffffffffffffffffffff1663ceeae52a826040518263ffffffff1660e01b81526004016104e391906110da565b5f604051808303815f87803b1580156104fa575f80fd5b505af115801561050c573d5f803e3d5ffd5b505050505b8260025f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546105989190611210565b925050819055508260015f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546105eb9190611210565b925050819055508260015f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825461063e9190611243565b925050819055508373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef856040516106a29190611003565b60405180910390a360019150509392505050565b60055f9054906101000a900460ff1681565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461071e575f80fd5b6107288282610bb9565b5050565b6001602052805f5260405f205f915090505481565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600480546107719061115e565b80601f016020809104026020016040519081016040528092919081815260200182805461079d9061115e565b80156107e85780601f106107bf576101008083540402835291602001916107e8565b820191905f5260205f20905b8154815290600101906020018083116107cb57829003601f168201915b505050505081565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610846575f80fd5b6108508282610d74565b5050565b5f8260f273ffffffffffffffffffffffffffffffffffffffff16634e25ab64826040518263ffffffff1660e01b815260040161089091906110da565b602060405180830381865afa1580156108ab573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906108cf91906111b8565b61093b5760f273ffffffffffffffffffffffffffffffffffffffff1663ceeae52a826040518263ffffffff1660e01b815260040161090d91906110da565b5f604051808303815f87803b158015610924575f80fd5b505af1158015610936573d5f803e3d5ffd5b505050505b8260015f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546109879190611210565b925050819055508260015f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546109da9190611243565b925050819055508373ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef85604051610a3e9190611003565b60405180910390a3600191505092915050565b6002602052815f5260405f20602052805f5260405f205f91509150505481565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610ac7575f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610afe575f80fd5b8073ffffffffffffffffffffffffffffffffffffffff165f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3805f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b8160f273ffffffffffffffffffffffffffffffffffffffff16634e25ab64826040518263ffffffff1660e01b8152600401610bf491906110da565b602060405180830381865afa158015610c0f573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610c3391906111b8565b610c9f5760f273ffffffffffffffffffffffffffffffffffffffff1663ceeae52a826040518263ffffffff1660e01b8152600401610c7191906110da565b5f604051808303815f87803b158015610c88575f80fd5b505af1158015610c9a573d5f803e3d5ffd5b505050505b8160015f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254610ceb9190611243565b925050819055508160065f828254610d039190611243565b925050819055508273ffffffffffffffffffffffffffffffffffffffff165f73ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051610d679190611003565b60405180910390a3505050565b8060015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254610dc09190611210565b925050819055508060065f828254610dd89190611210565b925050819055505f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051610e3c9190611003565b60405180910390a35050565b5f81519050919050565b5f82825260208201905092915050565b5f5b83811015610e7f578082015181840152602081019050610e64565b5f8484015250505050565b5f601f19601f8301169050919050565b5f610ea482610e48565b610eae8185610e52565b9350610ebe818560208601610e62565b610ec781610e8a565b840191505092915050565b5f6020820190508181035f830152610eea8184610e9a565b905092915050565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610f1f82610ef6565b9050919050565b610f2f81610f15565b8114610f39575f80fd5b50565b5f81359050610f4a81610f26565b92915050565b5f819050919050565b610f6281610f50565b8114610f6c575f80fd5b50565b5f81359050610f7d81610f59565b92915050565b5f8060408385031215610f9957610f98610ef2565b5b5f610fa685828601610f3c565b9250506020610fb785828601610f6f565b9150509250929050565b5f8115159050919050565b610fd581610fc1565b82525050565b5f602082019050610fee5f830184610fcc565b92915050565b610ffd81610f50565b82525050565b5f6020820190506110165f830184610ff4565b92915050565b5f805f6060848603121561103357611032610ef2565b5b5f61104086828701610f3c565b935050602061105186828701610f3c565b925050604061106286828701610f6f565b9150509250925092565b5f60ff82169050919050565b6110818161106c565b82525050565b5f60208201905061109a5f830184611078565b92915050565b5f602082840312156110b5576110b4610ef2565b5b5f6110c284828501610f3c565b91505092915050565b6110d481610f15565b82525050565b5f6020820190506110ed5f8301846110cb565b92915050565b5f806040838503121561110957611108610ef2565b5b5f61111685828601610f3c565b925050602061112785828601610f3c565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061117557607f821691505b60208210810361118857611187611131565b5b50919050565b61119781610fc1565b81146111a1575f80fd5b50565b5f815190506111b28161118e565b92915050565b5f602082840312156111cd576111cc610ef2565b5b5f6111da848285016111a4565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61121a82610f50565b915061122583610f50565b925082820390508181111561123d5761123c6111e3565b5b92915050565b5f61124d82610f50565b915061125883610f50565b92508282019050808211156112705761126f6111e3565b5b9291505056fea26469706673582212200a27c26244aa71c5e398725dcc35daa636754ef6cd7aaa924611c155b684230364736f6c63430008180033",
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
