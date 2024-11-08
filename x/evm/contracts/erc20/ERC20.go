// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package erc20

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

// Erc20MetaData contains all meta data concerning the Erc20 contract.
var Erc20MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"_decimals\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"metadataSealed\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"_metadataSealed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"sudoBurn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"sudoMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"sudoTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"_decimals\",\"type\":\"uint8\"}],\"name\":\"updateMetadata\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b50604051612c63380380612c6383398181016040528101906100319190610284565b335f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550836003908161007f919061052d565b50826004908161008f919061052d565b508160055f6101000a81548160ff021916908360ff1602179055508060075f6101000a81548160ff021916908315150217905550505050506105fc565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b61012b826100e5565b810181811067ffffffffffffffff8211171561014a576101496100f5565b5b80604052505050565b5f61015c6100cc565b90506101688282610122565b919050565b5f67ffffffffffffffff821115610187576101866100f5565b5b610190826100e5565b9050602081019050919050565b8281835e5f83830152505050565b5f6101bd6101b88461016d565b610153565b9050828152602081018484840111156101d9576101d86100e1565b5b6101e484828561019d565b509392505050565b5f82601f830112610200576101ff6100dd565b5b81516102108482602086016101ab565b91505092915050565b5f60ff82169050919050565b61022e81610219565b8114610238575f80fd5b50565b5f8151905061024981610225565b92915050565b5f8115159050919050565b6102638161024f565b811461026d575f80fd5b50565b5f8151905061027e8161025a565b92915050565b5f805f806080858703121561029c5761029b6100d5565b5b5f85015167ffffffffffffffff8111156102b9576102b86100d9565b5b6102c5878288016101ec565b945050602085015167ffffffffffffffff8111156102e6576102e56100d9565b5b6102f2878288016101ec565b93505060406103038782880161023b565b925050606061031487828801610270565b91505092959194509250565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061036e57607f821691505b6020821081036103815761038061032a565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026103e37fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826103a8565b6103ed86836103a8565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f61043161042c61042784610405565b61040e565b610405565b9050919050565b5f819050919050565b61044a83610417565b61045e61045682610438565b8484546103b4565b825550505050565b5f90565b610472610466565b61047d818484610441565b505050565b5b818110156104a0576104955f8261046a565b600181019050610483565b5050565b601f8211156104e5576104b681610387565b6104bf84610399565b810160208510156104ce578190505b6104e26104da85610399565b830182610482565b50505b505050565b5f82821c905092915050565b5f6105055f19846008026104ea565b1980831691505092915050565b5f61051d83836104f6565b9150826002028217905092915050565b61053682610320565b67ffffffffffffffff81111561054f5761054e6100f5565b5b6105598254610357565b6105648282856104a4565b5f60209050601f831160018114610595575f8415610583578287015190505b61058d8582610512565b8655506105f4565b601f1984166105a386610387565b5f5b828110156105ca578489015182556001820191506020850194506020810190506105a5565b868310156105e757848901516105e3601f8916826104f6565b8355505b6001600288020188555050505b505050505050565b61265a806106095f395ff3fe608060405234801561000f575f80fd5b506004361061012a575f3560e01c80636312013f116100ab5780639f1914841161006f5780639f19148414610342578063a9059cbb1461035e578063dd62ed3e1461038e578063f2fde38b146103be578063fe1195ec146103da5761012a565b80636312013f1461028857806370a08231146102a657806379cc6790146102d65780638da5cb5b1461030657806395d89b41146103245761012a565b806323b872dd116100f257806323b872dd146101e65780632d688ca814610216578063313ce5671461023257806340c10f191461025057806342966c681461026c5761012a565b806301ffc9a71461012e57806306fdde031461015e578063095ea7b31461017c57806318160ddd146101ac5780631988513b146101ca575b5f80fd5b61014860048036038101906101439190611880565b6103f6565b60405161015591906118c5565b60405180910390f35b61016661046f565b604051610173919061194e565b60405180910390f35b610196600480360381019061019191906119fb565b6104fb565b6040516101a391906118c5565b60405180910390f35b6101b46105e8565b6040516101c19190611a48565b60405180910390f35b6101e460048036038101906101df9190611a61565b6105ee565b005b61020060048036038101906101fb9190611a61565b61066d565b60405161020d91906118c5565b60405180910390f35b610230600480360381019061022b91906119fb565b610888565b005b61023a610905565b6040516102479190611acc565b60405180910390f35b61026a600480360381019061026591906119fb565b610917565b005b61028660048036038101906102819190611ae5565b610a36565b005b610290610afe565b60405161029d91906118c5565b60405180910390f35b6102c060048036038101906102bb9190611b10565b610b10565b6040516102cd9190611a48565b60405180910390f35b6102f060048036038101906102eb91906119fb565b610b25565b6040516102fd91906118c5565b60405180910390f35b61030e610d3e565b60405161031b9190611b4a565b60405180910390f35b61032c610d61565b604051610339919061194e565b60405180910390f35b61035c60048036038101906103579190611cb9565b610ded565b005b610378600480360381019061037391906119fb565b610f4d565b60405161038591906118c5565b60405180910390f35b6103a860048036038101906103a39190611d41565b61101e565b6040516103b59190611a48565b60405180910390f35b6103d860048036038101906103d39190611b10565b61103e565b005b6103f460048036038101906103ef91906119fb565b611186565b005b5f7f8da6da19000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19161480610468575061046782611203565b5b9050919050565b6003805461047c90611dac565b80601f01602080910402602001604051908101604052809291908181526020018280546104a890611dac565b80156104f35780601f106104ca576101008083540402835291602001916104f3565b820191905f5260205f20905b8154815290600101906020018083116104d657829003601f168201915b505050505081565b5f8160025f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040516105d69190611a48565b60405180910390a36001905092915050565b60065481565b600173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461065d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161065490611e26565b60405180910390fd5b61066883838361126c565b505050565b5f8260f173ffffffffffffffffffffffffffffffffffffffff1663f2af9ac9826040518263ffffffff1660e01b81526004016106a99190611b4a565b602060405180830381865afa1580156106c4573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906106e89190611e6e565b15610728576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161071f90611f09565b60405180910390fd5b8260025f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205410156107e3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107da90611f97565b60405180910390fd5b8260025f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825461086a9190611fe2565b9250508190555061087c85858561126c565b60019150509392505050565b600173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146108f7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108ee90611e26565b60405180910390fd5b61090182826114f7565b5050565b60055f9054906101000a900460ff1681565b8160f173ffffffffffffffffffffffffffffffffffffffff1663f2af9ac9826040518263ffffffff1660e01b81526004016109529190611b4a565b602060405180830381865afa15801561096d573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906109919190611e6e565b156109d1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109c89061205f565b60405180910390fd5b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610a27575f80fd5b610a3183836114f7565b505050565b3360f173ffffffffffffffffffffffffffffffffffffffff166360dc402f826040518263ffffffff1660e01b8152600401610a719190611b4a565b602060405180830381865afa158015610a8c573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610ab09190611e6e565b15610af0576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ae7906120c7565b60405180910390fd5b610afa33836116c6565b5050565b60075f9054906101000a900460ff1681565b6001602052805f5260405f205f915090505481565b5f8260f173ffffffffffffffffffffffffffffffffffffffff166360dc402f826040518263ffffffff1660e01b8152600401610b619190611b4a565b602060405180830381865afa158015610b7c573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610ba09190611e6e565b15610be0576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610bd7906120c7565b60405180910390fd5b8260025f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20541015610c9b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c9290612155565b60405180910390fd5b8260025f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254610d229190611fe2565b92505081905550610d3384846116c6565b600191505092915050565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60048054610d6e90611dac565b80601f0160208091040260200160405190810160405280929190818152602001828054610d9a90611dac565b8015610de55780601f10610dbc57610100808354040283529160200191610de5565b820191905f5260205f20905b815481529060010190602001808311610dc857829003601f168201915b505050505081565b60f173ffffffffffffffffffffffffffffffffffffffff166301116078336040518263ffffffff1660e01b8152600401610e279190611b4a565b602060405180830381865afa158015610e42573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610e669190611e6e565b610ea5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e9c906121e3565b60405180910390fd5b60075f9054906101000a900460ff1615610ef4576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610eeb9061224b565b60405180910390fd5b8260039081610f039190612406565b508160049081610f139190612406565b508060055f6101000a81548160ff021916908360ff160217905550600160075f6101000a81548160ff021916908315150217905550505050565b5f8260f173ffffffffffffffffffffffffffffffffffffffff1663f2af9ac9826040518263ffffffff1660e01b8152600401610f899190611b4a565b602060405180830381865afa158015610fa4573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610fc89190611e6e565b15611008576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610fff90611f09565b60405180910390fd5b61101333858561126c565b600191505092915050565b6002602052815f5260405f20602052805f5260405f205f91509150505481565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611094575f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036110cb575f80fd5b8073ffffffffffffffffffffffffffffffffffffffff165f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3805f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146111f5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016111ec90611e26565b60405180910390fd5b6111ff82826116c6565b5050565b5f7f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b8160f273ffffffffffffffffffffffffffffffffffffffff16634e25ab64826040518263ffffffff1660e01b81526004016112a79190611b4a565b602060405180830381865afa1580156112c2573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906112e69190611e6e565b6113665760f273ffffffffffffffffffffffffffffffffffffffff1663ceeae52a826040518263ffffffff1660e01b81526004016113249190611b4a565b6020604051808303815f875af1158015611340573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906113649190611e6e565b505b8160015f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205410156113e6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016113dd90612545565b60405180910390fd5b8160015f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546114329190611fe2565b925050819055508160015f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546114859190612563565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516114e99190611a48565b60405180910390a350505050565b8160f273ffffffffffffffffffffffffffffffffffffffff16634e25ab64826040518263ffffffff1660e01b81526004016115329190611b4a565b602060405180830381865afa15801561154d573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906115719190611e6e565b6115f15760f273ffffffffffffffffffffffffffffffffffffffff1663ceeae52a826040518263ffffffff1660e01b81526004016115af9190611b4a565b6020604051808303815f875af11580156115cb573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906115ef9190611e6e565b505b8160015f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825461163d9190612563565b925050819055508160065f8282546116559190612563565b925050819055508273ffffffffffffffffffffffffffffffffffffffff165f73ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516116b99190611a48565b60405180910390a3505050565b8060015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20541015611746576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161173d90612606565b60405180910390fd5b8060015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546117929190611fe2565b925050819055508060065f8282546117aa9190611fe2565b925050819055505f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8360405161180e9190611a48565b60405180910390a35050565b5f604051905090565b5f80fd5b5f80fd5b5f7fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b61185f8161182b565b8114611869575f80fd5b50565b5f8135905061187a81611856565b92915050565b5f6020828403121561189557611894611823565b5b5f6118a28482850161186c565b91505092915050565b5f8115159050919050565b6118bf816118ab565b82525050565b5f6020820190506118d85f8301846118b6565b92915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f611920826118de565b61192a81856118e8565b935061193a8185602086016118f8565b61194381611906565b840191505092915050565b5f6020820190508181035f8301526119668184611916565b905092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6119978261196e565b9050919050565b6119a78161198d565b81146119b1575f80fd5b50565b5f813590506119c28161199e565b92915050565b5f819050919050565b6119da816119c8565b81146119e4575f80fd5b50565b5f813590506119f5816119d1565b92915050565b5f8060408385031215611a1157611a10611823565b5b5f611a1e858286016119b4565b9250506020611a2f858286016119e7565b9150509250929050565b611a42816119c8565b82525050565b5f602082019050611a5b5f830184611a39565b92915050565b5f805f60608486031215611a7857611a77611823565b5b5f611a85868287016119b4565b9350506020611a96868287016119b4565b9250506040611aa7868287016119e7565b9150509250925092565b5f60ff82169050919050565b611ac681611ab1565b82525050565b5f602082019050611adf5f830184611abd565b92915050565b5f60208284031215611afa57611af9611823565b5b5f611b07848285016119e7565b91505092915050565b5f60208284031215611b2557611b24611823565b5b5f611b32848285016119b4565b91505092915050565b611b448161198d565b82525050565b5f602082019050611b5d5f830184611b3b565b92915050565b5f80fd5b5f80fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b611ba182611906565b810181811067ffffffffffffffff82111715611bc057611bbf611b6b565b5b80604052505050565b5f611bd261181a565b9050611bde8282611b98565b919050565b5f67ffffffffffffffff821115611bfd57611bfc611b6b565b5b611c0682611906565b9050602081019050919050565b828183375f83830152505050565b5f611c33611c2e84611be3565b611bc9565b905082815260208101848484011115611c4f57611c4e611b67565b5b611c5a848285611c13565b509392505050565b5f82601f830112611c7657611c75611b63565b5b8135611c86848260208601611c21565b91505092915050565b611c9881611ab1565b8114611ca2575f80fd5b50565b5f81359050611cb381611c8f565b92915050565b5f805f60608486031215611cd057611ccf611823565b5b5f84013567ffffffffffffffff811115611ced57611cec611827565b5b611cf986828701611c62565b935050602084013567ffffffffffffffff811115611d1a57611d19611827565b5b611d2686828701611c62565b9250506040611d3786828701611ca5565b9150509250925092565b5f8060408385031215611d5757611d56611823565b5b5f611d64858286016119b4565b9250506020611d75858286016119b4565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680611dc357607f821691505b602082108103611dd657611dd5611d7f565b5b50919050565b7f45524332303a2063616c6c6572206973206e6f742074686520636861696e00005f82015250565b5f611e10601e836118e8565b9150611e1b82611ddc565b602082019050919050565b5f6020820190508181035f830152611e3d81611e04565b9050919050565b611e4d816118ab565b8114611e57575f80fd5b50565b5f81519050611e6881611e44565b92915050565b5f60208284031215611e8357611e82611823565b5b5f611e9084828501611e5a565b91505092915050565b7f45524332303a207472616e7366657220746f20626c6f636b65642061646472655f8201527f7373000000000000000000000000000000000000000000000000000000000000602082015250565b5f611ef36022836118e8565b9150611efe82611e99565b604082019050919050565b5f6020820190508181035f830152611f2081611ee7565b9050919050565b7f45524332303a207472616e7366657220616d6f756e74206578636565647320615f8201527f6c6c6f77616e6365000000000000000000000000000000000000000000000000602082015250565b5f611f816028836118e8565b9150611f8c82611f27565b604082019050919050565b5f6020820190508181035f830152611fae81611f75565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f611fec826119c8565b9150611ff7836119c8565b925082820390508181111561200f5761200e611fb5565b5b92915050565b7f45524332303a206d696e7420746f20626c6f636b6564206164647265737300005f82015250565b5f612049601e836118e8565b915061205482612015565b602082019050919050565b5f6020820190508181035f8301526120768161203d565b9050919050565b7f45524332303a206275726e2066726f6d206d6f64756c652061646472657373005f82015250565b5f6120b1601f836118e8565b91506120bc8261207d565b602082019050919050565b5f6020820190508181035f8301526120de816120a5565b9050919050565b7f45524332303a206275726e20616d6f756e74206578636565647320616c6c6f775f8201527f616e636500000000000000000000000000000000000000000000000000000000602082015250565b5f61213f6024836118e8565b915061214a826120e5565b604082019050919050565b5f6020820190508181035f83015261216c81612133565b9050919050565b7f45524332303a2063616c6c6572206973206e6f742074686520617574686f72695f8201527f7479000000000000000000000000000000000000000000000000000000000000602082015250565b5f6121cd6022836118e8565b91506121d882612173565b604082019050919050565b5f6020820190508181035f8301526121fa816121c1565b9050919050565b7f45524332303a206d65746164617461207365616c6564000000000000000000005f82015250565b5f6122356016836118e8565b915061224082612201565b602082019050919050565b5f6020820190508181035f83015261226281612229565b9050919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026122c57fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8261228a565b6122cf868361228a565b95508019841693508086168417925050509392505050565b5f819050919050565b5f61230a612305612300846119c8565b6122e7565b6119c8565b9050919050565b5f819050919050565b612323836122f0565b61233761232f82612311565b848454612296565b825550505050565b5f90565b61234b61233f565b61235681848461231a565b505050565b5b818110156123795761236e5f82612343565b60018101905061235c565b5050565b601f8211156123be5761238f81612269565b6123988461227b565b810160208510156123a7578190505b6123bb6123b38561227b565b83018261235b565b50505b505050565b5f82821c905092915050565b5f6123de5f19846008026123c3565b1980831691505092915050565b5f6123f683836123cf565b9150826002028217905092915050565b61240f826118de565b67ffffffffffffffff81111561242857612427611b6b565b5b6124328254611dac565b61243d82828561237d565b5f60209050601f83116001811461246e575f841561245c578287015190505b61246685826123eb565b8655506124cd565b601f19841661247c86612269565b5f5b828110156124a35784890151825560018201915060208501945060208101905061247e565b868310156124c057848901516124bc601f8916826123cf565b8355505b6001600288020188555050505b505050505050565b7f45524332303a207472616e7366657220616d6f756e74206578636565647320625f8201527f616c616e63650000000000000000000000000000000000000000000000000000602082015250565b5f61252f6026836118e8565b915061253a826124d5565b604082019050919050565b5f6020820190508181035f83015261255c81612523565b9050919050565b5f61256d826119c8565b9150612578836119c8565b92508282019050808211156125905761258f611fb5565b5b92915050565b7f45524332303a206275726e20616d6f756e7420657863656564732062616c616e5f8201527f6365000000000000000000000000000000000000000000000000000000000000602082015250565b5f6125f06022836118e8565b91506125fb82612596565b604082019050919050565b5f6020820190508181035f83015261261d816125e4565b905091905056fea264697066735822122052cf15d32a984b09e97731db92fec90f8955cbb2a6d66c568f7e9f5043fdfa8864736f6c63430008190033",
}

// Erc20ABI is the input ABI used to generate the binding from.
// Deprecated: Use Erc20MetaData.ABI instead.
var Erc20ABI = Erc20MetaData.ABI

// Erc20Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Erc20MetaData.Bin instead.
var Erc20Bin = Erc20MetaData.Bin

// DeployErc20 deploys a new Ethereum contract, binding an instance of Erc20 to it.
func DeployErc20(auth *bind.TransactOpts, backend bind.ContractBackend, _name string, _symbol string, _decimals uint8, metadataSealed bool) (common.Address, *types.Transaction, *Erc20, error) {
	parsed, err := Erc20MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Erc20Bin), backend, _name, _symbol, _decimals, metadataSealed)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Erc20{Erc20Caller: Erc20Caller{contract: contract}, Erc20Transactor: Erc20Transactor{contract: contract}, Erc20Filterer: Erc20Filterer{contract: contract}}, nil
}

// Erc20 is an auto generated Go binding around an Ethereum contract.
type Erc20 struct {
	Erc20Caller     // Read-only binding to the contract
	Erc20Transactor // Write-only binding to the contract
	Erc20Filterer   // Log filterer for contract events
}

// Erc20Caller is an auto generated read-only Go binding around an Ethereum contract.
type Erc20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Erc20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Erc20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Erc20Session struct {
	Contract     *Erc20            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Erc20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Erc20CallerSession struct {
	Contract *Erc20Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// Erc20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Erc20TransactorSession struct {
	Contract     *Erc20Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Erc20Raw is an auto generated low-level Go binding around an Ethereum contract.
type Erc20Raw struct {
	Contract *Erc20 // Generic contract binding to access the raw methods on
}

// Erc20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Erc20CallerRaw struct {
	Contract *Erc20Caller // Generic read-only contract binding to access the raw methods on
}

// Erc20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Erc20TransactorRaw struct {
	Contract *Erc20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewErc20 creates a new instance of Erc20, bound to a specific deployed contract.
func NewErc20(address common.Address, backend bind.ContractBackend) (*Erc20, error) {
	contract, err := bindErc20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Erc20{Erc20Caller: Erc20Caller{contract: contract}, Erc20Transactor: Erc20Transactor{contract: contract}, Erc20Filterer: Erc20Filterer{contract: contract}}, nil
}

// NewErc20Caller creates a new read-only instance of Erc20, bound to a specific deployed contract.
func NewErc20Caller(address common.Address, caller bind.ContractCaller) (*Erc20Caller, error) {
	contract, err := bindErc20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Erc20Caller{contract: contract}, nil
}

// NewErc20Transactor creates a new write-only instance of Erc20, bound to a specific deployed contract.
func NewErc20Transactor(address common.Address, transactor bind.ContractTransactor) (*Erc20Transactor, error) {
	contract, err := bindErc20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Erc20Transactor{contract: contract}, nil
}

// NewErc20Filterer creates a new log filterer instance of Erc20, bound to a specific deployed contract.
func NewErc20Filterer(address common.Address, filterer bind.ContractFilterer) (*Erc20Filterer, error) {
	contract, err := bindErc20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Erc20Filterer{contract: contract}, nil
}

// bindErc20 binds a generic wrapper to an already deployed contract.
func bindErc20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Erc20MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc20 *Erc20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc20.Contract.Erc20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc20 *Erc20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc20.Contract.Erc20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc20 *Erc20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc20.Contract.Erc20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc20 *Erc20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc20 *Erc20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc20 *Erc20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc20.Contract.contract.Transact(opts, method, params...)
}

// MetadataSealed is a free data retrieval call binding the contract method 0x6312013f.
//
// Solidity: function _metadataSealed() view returns(bool)
func (_Erc20 *Erc20Caller) MetadataSealed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "_metadataSealed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MetadataSealed is a free data retrieval call binding the contract method 0x6312013f.
//
// Solidity: function _metadataSealed() view returns(bool)
func (_Erc20 *Erc20Session) MetadataSealed() (bool, error) {
	return _Erc20.Contract.MetadataSealed(&_Erc20.CallOpts)
}

// MetadataSealed is a free data retrieval call binding the contract method 0x6312013f.
//
// Solidity: function _metadataSealed() view returns(bool)
func (_Erc20 *Erc20CallerSession) MetadataSealed() (bool, error) {
	return _Erc20.Contract.MetadataSealed(&_Erc20.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Erc20 *Erc20Caller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Erc20 *Erc20Session) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Erc20.Contract.Allowance(&_Erc20.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Erc20 *Erc20CallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Erc20.Contract.Allowance(&_Erc20.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Erc20 *Erc20Caller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Erc20 *Erc20Session) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Erc20.Contract.BalanceOf(&_Erc20.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Erc20 *Erc20CallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Erc20.Contract.BalanceOf(&_Erc20.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Erc20 *Erc20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Erc20 *Erc20Session) Decimals() (uint8, error) {
	return _Erc20.Contract.Decimals(&_Erc20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Erc20 *Erc20CallerSession) Decimals() (uint8, error) {
	return _Erc20.Contract.Decimals(&_Erc20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Erc20 *Erc20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Erc20 *Erc20Session) Name() (string, error) {
	return _Erc20.Contract.Name(&_Erc20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Erc20 *Erc20CallerSession) Name() (string, error) {
	return _Erc20.Contract.Name(&_Erc20.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Erc20 *Erc20Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Erc20 *Erc20Session) Owner() (common.Address, error) {
	return _Erc20.Contract.Owner(&_Erc20.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Erc20 *Erc20CallerSession) Owner() (common.Address, error) {
	return _Erc20.Contract.Owner(&_Erc20.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Erc20 *Erc20Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Erc20 *Erc20Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Erc20.Contract.SupportsInterface(&_Erc20.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Erc20 *Erc20CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Erc20.Contract.SupportsInterface(&_Erc20.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Erc20 *Erc20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Erc20 *Erc20Session) Symbol() (string, error) {
	return _Erc20.Contract.Symbol(&_Erc20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Erc20 *Erc20CallerSession) Symbol() (string, error) {
	return _Erc20.Contract.Symbol(&_Erc20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Erc20 *Erc20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Erc20 *Erc20Session) TotalSupply() (*big.Int, error) {
	return _Erc20.Contract.TotalSupply(&_Erc20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Erc20 *Erc20CallerSession) TotalSupply() (*big.Int, error) {
	return _Erc20.Contract.TotalSupply(&_Erc20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Erc20 *Erc20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Erc20 *Erc20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.Approve(&_Erc20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Erc20 *Erc20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.Approve(&_Erc20.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_Erc20 *Erc20Transactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_Erc20 *Erc20Session) Burn(amount *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.Burn(&_Erc20.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_Erc20 *Erc20TransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.Burn(&_Erc20.TransactOpts, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address from, uint256 amount) returns(bool)
func (_Erc20 *Erc20Transactor) BurnFrom(opts *bind.TransactOpts, from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "burnFrom", from, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address from, uint256 amount) returns(bool)
func (_Erc20 *Erc20Session) BurnFrom(from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.BurnFrom(&_Erc20.TransactOpts, from, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address from, uint256 amount) returns(bool)
func (_Erc20 *Erc20TransactorSession) BurnFrom(from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.BurnFrom(&_Erc20.TransactOpts, from, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_Erc20 *Erc20Transactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "mint", to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_Erc20 *Erc20Session) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.Mint(&_Erc20.TransactOpts, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_Erc20 *Erc20TransactorSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.Mint(&_Erc20.TransactOpts, to, amount)
}

// SudoBurn is a paid mutator transaction binding the contract method 0xfe1195ec.
//
// Solidity: function sudoBurn(address from, uint256 amount) returns()
func (_Erc20 *Erc20Transactor) SudoBurn(opts *bind.TransactOpts, from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "sudoBurn", from, amount)
}

// SudoBurn is a paid mutator transaction binding the contract method 0xfe1195ec.
//
// Solidity: function sudoBurn(address from, uint256 amount) returns()
func (_Erc20 *Erc20Session) SudoBurn(from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.SudoBurn(&_Erc20.TransactOpts, from, amount)
}

// SudoBurn is a paid mutator transaction binding the contract method 0xfe1195ec.
//
// Solidity: function sudoBurn(address from, uint256 amount) returns()
func (_Erc20 *Erc20TransactorSession) SudoBurn(from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.SudoBurn(&_Erc20.TransactOpts, from, amount)
}

// SudoMint is a paid mutator transaction binding the contract method 0x2d688ca8.
//
// Solidity: function sudoMint(address to, uint256 amount) returns()
func (_Erc20 *Erc20Transactor) SudoMint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "sudoMint", to, amount)
}

// SudoMint is a paid mutator transaction binding the contract method 0x2d688ca8.
//
// Solidity: function sudoMint(address to, uint256 amount) returns()
func (_Erc20 *Erc20Session) SudoMint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.SudoMint(&_Erc20.TransactOpts, to, amount)
}

// SudoMint is a paid mutator transaction binding the contract method 0x2d688ca8.
//
// Solidity: function sudoMint(address to, uint256 amount) returns()
func (_Erc20 *Erc20TransactorSession) SudoMint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.SudoMint(&_Erc20.TransactOpts, to, amount)
}

// SudoTransfer is a paid mutator transaction binding the contract method 0x1988513b.
//
// Solidity: function sudoTransfer(address sender, address recipient, uint256 amount) returns()
func (_Erc20 *Erc20Transactor) SudoTransfer(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "sudoTransfer", sender, recipient, amount)
}

// SudoTransfer is a paid mutator transaction binding the contract method 0x1988513b.
//
// Solidity: function sudoTransfer(address sender, address recipient, uint256 amount) returns()
func (_Erc20 *Erc20Session) SudoTransfer(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.SudoTransfer(&_Erc20.TransactOpts, sender, recipient, amount)
}

// SudoTransfer is a paid mutator transaction binding the contract method 0x1988513b.
//
// Solidity: function sudoTransfer(address sender, address recipient, uint256 amount) returns()
func (_Erc20 *Erc20TransactorSession) SudoTransfer(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.SudoTransfer(&_Erc20.TransactOpts, sender, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Erc20 *Erc20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Erc20 *Erc20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.Transfer(&_Erc20.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Erc20 *Erc20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.Transfer(&_Erc20.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Erc20 *Erc20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Erc20 *Erc20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.TransferFrom(&_Erc20.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Erc20 *Erc20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.TransferFrom(&_Erc20.TransactOpts, sender, recipient, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Erc20 *Erc20Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Erc20 *Erc20Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Erc20.Contract.TransferOwnership(&_Erc20.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Erc20 *Erc20TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Erc20.Contract.TransferOwnership(&_Erc20.TransactOpts, newOwner)
}

// UpdateMetadata is a paid mutator transaction binding the contract method 0x9f191484.
//
// Solidity: function updateMetadata(string _name, string _symbol, uint8 _decimals) returns()
func (_Erc20 *Erc20Transactor) UpdateMetadata(opts *bind.TransactOpts, _name string, _symbol string, _decimals uint8) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "updateMetadata", _name, _symbol, _decimals)
}

// UpdateMetadata is a paid mutator transaction binding the contract method 0x9f191484.
//
// Solidity: function updateMetadata(string _name, string _symbol, uint8 _decimals) returns()
func (_Erc20 *Erc20Session) UpdateMetadata(_name string, _symbol string, _decimals uint8) (*types.Transaction, error) {
	return _Erc20.Contract.UpdateMetadata(&_Erc20.TransactOpts, _name, _symbol, _decimals)
}

// UpdateMetadata is a paid mutator transaction binding the contract method 0x9f191484.
//
// Solidity: function updateMetadata(string _name, string _symbol, uint8 _decimals) returns()
func (_Erc20 *Erc20TransactorSession) UpdateMetadata(_name string, _symbol string, _decimals uint8) (*types.Transaction, error) {
	return _Erc20.Contract.UpdateMetadata(&_Erc20.TransactOpts, _name, _symbol, _decimals)
}

// Erc20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Erc20 contract.
type Erc20ApprovalIterator struct {
	Event *Erc20Approval // Event containing the contract specifics and raw log

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
func (it *Erc20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc20Approval)
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
		it.Event = new(Erc20Approval)
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
func (it *Erc20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc20Approval represents a Approval event raised by the Erc20 contract.
type Erc20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Erc20 *Erc20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*Erc20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Erc20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &Erc20ApprovalIterator{contract: _Erc20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Erc20 *Erc20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *Erc20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Erc20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc20Approval)
				if err := _Erc20.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Erc20 *Erc20Filterer) ParseApproval(log types.Log) (*Erc20Approval, error) {
	event := new(Erc20Approval)
	if err := _Erc20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc20OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Erc20 contract.
type Erc20OwnershipTransferredIterator struct {
	Event *Erc20OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *Erc20OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc20OwnershipTransferred)
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
		it.Event = new(Erc20OwnershipTransferred)
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
func (it *Erc20OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc20OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc20OwnershipTransferred represents a OwnershipTransferred event raised by the Erc20 contract.
type Erc20OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Erc20 *Erc20Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*Erc20OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Erc20.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &Erc20OwnershipTransferredIterator{contract: _Erc20.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Erc20 *Erc20Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *Erc20OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Erc20.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc20OwnershipTransferred)
				if err := _Erc20.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Erc20 *Erc20Filterer) ParseOwnershipTransferred(log types.Log) (*Erc20OwnershipTransferred, error) {
	event := new(Erc20OwnershipTransferred)
	if err := _Erc20.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Erc20 contract.
type Erc20TransferIterator struct {
	Event *Erc20Transfer // Event containing the contract specifics and raw log

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
func (it *Erc20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc20Transfer)
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
		it.Event = new(Erc20Transfer)
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
func (it *Erc20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc20Transfer represents a Transfer event raised by the Erc20 contract.
type Erc20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Erc20 *Erc20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*Erc20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Erc20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Erc20TransferIterator{contract: _Erc20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Erc20 *Erc20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *Erc20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Erc20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc20Transfer)
				if err := _Erc20.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Erc20 *Erc20Filterer) ParseTransfer(log types.Log) (*Erc20Transfer, error) {
	event := new(Erc20Transfer)
	if err := _Erc20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
