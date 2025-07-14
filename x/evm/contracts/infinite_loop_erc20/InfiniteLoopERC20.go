// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package infinite_loop_erc20

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

// InfiniteLoopErc20MetaData contains all meta data concerning the InfiniteLoopErc20 contract.
var InfiniteLoopErc20MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"__name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"__symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"__decimals\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"sudoTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b50604051612570380380612570833981810160405281019061003191906102a5565b335f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060f273ffffffffffffffffffffffffffffffffffffffff16635e6c57596040518163ffffffff1660e01b81526004016020604051808303815f875af11580156100bb573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906100df9190610362565b5082600390816100ef919061059a565b5081600490816100ff919061059a565b508060055f6101000a81548160ff021916908360ff160217905550505050610669565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6101818261013b565b810181811067ffffffffffffffff821117156101a05761019f61014b565b5b80604052505050565b5f6101b2610122565b90506101be8282610178565b919050565b5f67ffffffffffffffff8211156101dd576101dc61014b565b5b6101e68261013b565b9050602081019050919050565b8281835e5f83830152505050565b5f61021361020e846101c3565b6101a9565b90508281526020810184848401111561022f5761022e610137565b5b61023a8482856101f3565b509392505050565b5f82601f83011261025657610255610133565b5b8151610266848260208601610201565b91505092915050565b5f60ff82169050919050565b6102848161026f565b811461028e575f80fd5b50565b5f8151905061029f8161027b565b92915050565b5f805f606084860312156102bc576102bb61012b565b5b5f84015167ffffffffffffffff8111156102d9576102d861012f565b5b6102e586828701610242565b935050602084015167ffffffffffffffff8111156103065761030561012f565b5b61031286828701610242565b925050604061032386828701610291565b9150509250925092565b5f8115159050919050565b6103418161032d565b811461034b575f80fd5b50565b5f8151905061035c81610338565b92915050565b5f602082840312156103775761037661012b565b5b5f6103848482850161034e565b91505092915050565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806103db57607f821691505b6020821081036103ee576103ed610397565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026104507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610415565b61045a8683610415565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f61049e61049961049484610472565b61047b565b610472565b9050919050565b5f819050919050565b6104b783610484565b6104cb6104c3826104a5565b848454610421565b825550505050565b5f90565b6104df6104d3565b6104ea8184846104ae565b505050565b5b8181101561050d576105025f826104d7565b6001810190506104f0565b5050565b601f82111561055257610523816103f4565b61052c84610406565b8101602085101561053b578190505b61054f61054785610406565b8301826104ef565b50505b505050565b5f82821c905092915050565b5f6105725f1984600802610557565b1980831691505092915050565b5f61058a8383610563565b9150826002028217905092915050565b6105a38261038d565b67ffffffffffffffff8111156105bc576105bb61014b565b5b6105c682546103c4565b6105d1828285610511565b5f60209050601f831160018114610602575f84156105f0578287015190505b6105fa858261057f565b865550610661565b601f198416610610866103f4565b5f5b8281101561063757848901518255600182019150602085019450602081019050610612565b868310156106545784890151610650601f891682610563565b8355505b6001600288020188555050505b505050505050565b611efa806106765f395ff3fe608060405234801561000f575f80fd5b50600436106100fe575f3560e01c806342966c681161009557806395d89b411161006457806395d89b41146102be578063a9059cbb146102dc578063dd62ed3e1461030c578063f2fde38b1461033c576100fe565b806342966c681461022457806370a082311461024057806379cc6790146102705780638da5cb5b146102a0576100fe565b80631988513b116100d15780631988513b1461019e57806323b872dd146101ba578063313ce567146101ea57806340c10f1914610208576100fe565b806301ffc9a71461010257806306fdde0314610132578063095ea7b31461015057806318160ddd14610180575b5f80fd5b61011c60048036038101906101179190611642565b610358565b6040516101299190611687565b60405180910390f35b61013a6103d1565b6040516101479190611710565b60405180910390f35b61016a600480360381019061016591906117bd565b61047e565b6040516101779190611687565b60405180910390f35b61018861056b565b604051610195919061180a565b60405180910390f35b6101b860048036038101906101b39190611823565b610594565b005b6101d460048036038101906101cf9190611823565b610613565b6040516101e19190611687565b60405180910390f35b6101f261082e565b6040516101ff919061188e565b60405180910390f35b610222600480360381019061021d91906117bd565b610863565b005b61023e600480360381019061023991906118a7565b610982565b005b61025a600480360381019061025591906118d2565b610a4a565b604051610267919061180a565b60405180910390f35b61028a600480360381019061028591906117bd565b610ab0565b6040516102979190611687565b60405180910390f35b6102a8610cc9565b6040516102b5919061190c565b60405180910390f35b6102c6610cec565b6040516102d39190611710565b60405180910390f35b6102f660048036038101906102f191906117bd565b610d99565b6040516103039190611687565b60405180910390f35b61032660048036038101906103219190611925565b610e6a565b604051610333919061180a565b60405180910390f35b610356600480360381019061035191906118d2565b610e8a565b005b5f7f8da6da19000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614806103ca57506103c982610fd2565b5b9050919050565b60605f5b6001156103ef5780806103e790611990565b9150506103d5565b600380546103fc90611a04565b80601f016020809104026020016040519081016040528092919081815260200182805461042890611a04565b80156104735780601f1061044a57610100808354040283529160200191610473565b820191905f5260205f20905b81548152906001019060200180831161045657829003601f168201915b505050505091505090565b5f8160025f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92584604051610559919061180a565b60405180910390a36001905092915050565b5f805f90505b60011561058b57808061058390611990565b915050610571565b60065491505090565b600173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610603576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105fa90611a7e565b60405180910390fd5b61060e83838361103b565b505050565b5f8260f173ffffffffffffffffffffffffffffffffffffffff1663f2af9ac9826040518263ffffffff1660e01b815260040161064f919061190c565b602060405180830381865afa15801561066a573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061068e9190611ac6565b156106ce576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106c590611b61565b60405180910390fd5b8260025f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20541015610789576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161078090611bef565b60405180910390fd5b8260025f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546108109190611c0d565b9250508190555061082285858561103b565b60019150509392505050565b5f805f90505b60011561084e57808061084690611990565b915050610834565b60055f9054906101000a900460ff1691505090565b8160f173ffffffffffffffffffffffffffffffffffffffff1663f2af9ac9826040518263ffffffff1660e01b815260040161089e919061190c565b602060405180830381865afa1580156108b9573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906108dd9190611ac6565b1561091d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161091490611c8a565b60405180910390fd5b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610973575f80fd5b61097d83836112c6565b505050565b3360f173ffffffffffffffffffffffffffffffffffffffff166360dc402f826040518263ffffffff1660e01b81526004016109bd919061190c565b602060405180830381865afa1580156109d8573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906109fc9190611ac6565b15610a3c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a3390611cf2565b60405180910390fd5b610a463383611495565b5050565b5f805f90505b600115610a6a578080610a6290611990565b915050610a50565b60015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054915050919050565b5f8260f173ffffffffffffffffffffffffffffffffffffffff166360dc402f826040518263ffffffff1660e01b8152600401610aec919061190c565b602060405180830381865afa158015610b07573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610b2b9190611ac6565b15610b6b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b6290611cf2565b60405180910390fd5b8260025f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20541015610c26576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c1d90611d80565b60405180910390fd5b8260025f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254610cad9190611c0d565b92505081905550610cbe8484611495565b600191505092915050565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60605f5b600115610d0a578080610d0290611990565b915050610cf0565b60048054610d1790611a04565b80601f0160208091040260200160405190810160405280929190818152602001828054610d4390611a04565b8015610d8e5780601f10610d6557610100808354040283529160200191610d8e565b820191905f5260205f20905b815481529060010190602001808311610d7157829003601f168201915b505050505091505090565b5f8260f173ffffffffffffffffffffffffffffffffffffffff1663f2af9ac9826040518263ffffffff1660e01b8152600401610dd5919061190c565b602060405180830381865afa158015610df0573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610e149190611ac6565b15610e54576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e4b90611b61565b60405180910390fd5b610e5f33858561103b565b600191505092915050565b6002602052815f5260405f20602052805f5260405f205f91509150505481565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610ee0575f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610f17575f80fd5b8073ffffffffffffffffffffffffffffffffffffffff165f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3805f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b5f7f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b8160f273ffffffffffffffffffffffffffffffffffffffff16634e25ab64826040518263ffffffff1660e01b8152600401611076919061190c565b602060405180830381865afa158015611091573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906110b59190611ac6565b6111355760f273ffffffffffffffffffffffffffffffffffffffff1663ceeae52a826040518263ffffffff1660e01b81526004016110f3919061190c565b6020604051808303815f875af115801561110f573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906111339190611ac6565b505b8160015f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205410156111b5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016111ac90611e0e565b60405180910390fd5b8160015f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546112019190611c0d565b925050819055508160015f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546112549190611e2c565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516112b8919061180a565b60405180910390a350505050565b8160f273ffffffffffffffffffffffffffffffffffffffff16634e25ab64826040518263ffffffff1660e01b8152600401611301919061190c565b602060405180830381865afa15801561131c573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906113409190611ac6565b6113c05760f273ffffffffffffffffffffffffffffffffffffffff1663ceeae52a826040518263ffffffff1660e01b815260040161137e919061190c565b6020604051808303815f875af115801561139a573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906113be9190611ac6565b505b8160015f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825461140c9190611e2c565b925050819055508160065f8282546114249190611e2c565b925050819055508273ffffffffffffffffffffffffffffffffffffffff165f73ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051611488919061180a565b60405180910390a3505050565b8060015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20541015611515576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161150c90611ecf565b60405180910390fd5b8060015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546115619190611c0d565b925050819055508060065f8282546115799190611c0d565b925050819055505f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516115dd919061180a565b60405180910390a35050565b5f80fd5b5f7fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b611621816115ed565b811461162b575f80fd5b50565b5f8135905061163c81611618565b92915050565b5f60208284031215611657576116566115e9565b5b5f6116648482850161162e565b91505092915050565b5f8115159050919050565b6116818161166d565b82525050565b5f60208201905061169a5f830184611678565b92915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f6116e2826116a0565b6116ec81856116aa565b93506116fc8185602086016116ba565b611705816116c8565b840191505092915050565b5f6020820190508181035f83015261172881846116d8565b905092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61175982611730565b9050919050565b6117698161174f565b8114611773575f80fd5b50565b5f8135905061178481611760565b92915050565b5f819050919050565b61179c8161178a565b81146117a6575f80fd5b50565b5f813590506117b781611793565b92915050565b5f80604083850312156117d3576117d26115e9565b5b5f6117e085828601611776565b92505060206117f1858286016117a9565b9150509250929050565b6118048161178a565b82525050565b5f60208201905061181d5f8301846117fb565b92915050565b5f805f6060848603121561183a576118396115e9565b5b5f61184786828701611776565b935050602061185886828701611776565b9250506040611869868287016117a9565b9150509250925092565b5f60ff82169050919050565b61188881611873565b82525050565b5f6020820190506118a15f83018461187f565b92915050565b5f602082840312156118bc576118bb6115e9565b5b5f6118c9848285016117a9565b91505092915050565b5f602082840312156118e7576118e66115e9565b5b5f6118f484828501611776565b91505092915050565b6119068161174f565b82525050565b5f60208201905061191f5f8301846118fd565b92915050565b5f806040838503121561193b5761193a6115e9565b5b5f61194885828601611776565b925050602061195985828601611776565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61199a8261178a565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036119cc576119cb611963565b5b600182019050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680611a1b57607f821691505b602082108103611a2e57611a2d6119d7565b5b50919050565b7f45524332303a2063616c6c6572206973206e6f742074686520636861696e00005f82015250565b5f611a68601e836116aa565b9150611a7382611a34565b602082019050919050565b5f6020820190508181035f830152611a9581611a5c565b9050919050565b611aa58161166d565b8114611aaf575f80fd5b50565b5f81519050611ac081611a9c565b92915050565b5f60208284031215611adb57611ada6115e9565b5b5f611ae884828501611ab2565b91505092915050565b7f45524332303a207472616e7366657220746f20626c6f636b65642061646472655f8201527f7373000000000000000000000000000000000000000000000000000000000000602082015250565b5f611b4b6022836116aa565b9150611b5682611af1565b604082019050919050565b5f6020820190508181035f830152611b7881611b3f565b9050919050565b7f45524332303a207472616e7366657220616d6f756e74206578636565647320615f8201527f6c6c6f77616e6365000000000000000000000000000000000000000000000000602082015250565b5f611bd96028836116aa565b9150611be482611b7f565b604082019050919050565b5f6020820190508181035f830152611c0681611bcd565b9050919050565b5f611c178261178a565b9150611c228361178a565b9250828203905081811115611c3a57611c39611963565b5b92915050565b7f45524332303a206d696e7420746f20626c6f636b6564206164647265737300005f82015250565b5f611c74601e836116aa565b9150611c7f82611c40565b602082019050919050565b5f6020820190508181035f830152611ca181611c68565b9050919050565b7f45524332303a206275726e2066726f6d206d6f64756c652061646472657373005f82015250565b5f611cdc601f836116aa565b9150611ce782611ca8565b602082019050919050565b5f6020820190508181035f830152611d0981611cd0565b9050919050565b7f45524332303a206275726e20616d6f756e74206578636565647320616c6c6f775f8201527f616e636500000000000000000000000000000000000000000000000000000000602082015250565b5f611d6a6024836116aa565b9150611d7582611d10565b604082019050919050565b5f6020820190508181035f830152611d9781611d5e565b9050919050565b7f45524332303a207472616e7366657220616d6f756e74206578636565647320625f8201527f616c616e63650000000000000000000000000000000000000000000000000000602082015250565b5f611df86026836116aa565b9150611e0382611d9e565b604082019050919050565b5f6020820190508181035f830152611e2581611dec565b9050919050565b5f611e368261178a565b9150611e418361178a565b9250828201905080821115611e5957611e58611963565b5b92915050565b7f45524332303a206275726e20616d6f756e7420657863656564732062616c616e5f8201527f6365000000000000000000000000000000000000000000000000000000000000602082015250565b5f611eb96022836116aa565b9150611ec482611e5f565b604082019050919050565b5f6020820190508181035f830152611ee681611ead565b905091905056fea164736f6c6343000819000a",
}

// InfiniteLoopErc20ABI is the input ABI used to generate the binding from.
// Deprecated: Use InfiniteLoopErc20MetaData.ABI instead.
var InfiniteLoopErc20ABI = InfiniteLoopErc20MetaData.ABI

// InfiniteLoopErc20Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use InfiniteLoopErc20MetaData.Bin instead.
var InfiniteLoopErc20Bin = InfiniteLoopErc20MetaData.Bin

// DeployInfiniteLoopErc20 deploys a new Ethereum contract, binding an instance of InfiniteLoopErc20 to it.
func DeployInfiniteLoopErc20(auth *bind.TransactOpts, backend bind.ContractBackend, __name string, __symbol string, __decimals uint8) (common.Address, *types.Transaction, *InfiniteLoopErc20, error) {
	parsed, err := InfiniteLoopErc20MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(InfiniteLoopErc20Bin), backend, __name, __symbol, __decimals)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &InfiniteLoopErc20{InfiniteLoopErc20Caller: InfiniteLoopErc20Caller{contract: contract}, InfiniteLoopErc20Transactor: InfiniteLoopErc20Transactor{contract: contract}, InfiniteLoopErc20Filterer: InfiniteLoopErc20Filterer{contract: contract}}, nil
}

// InfiniteLoopErc20 is an auto generated Go binding around an Ethereum contract.
type InfiniteLoopErc20 struct {
	InfiniteLoopErc20Caller     // Read-only binding to the contract
	InfiniteLoopErc20Transactor // Write-only binding to the contract
	InfiniteLoopErc20Filterer   // Log filterer for contract events
}

// InfiniteLoopErc20Caller is an auto generated read-only Go binding around an Ethereum contract.
type InfiniteLoopErc20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InfiniteLoopErc20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type InfiniteLoopErc20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InfiniteLoopErc20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InfiniteLoopErc20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InfiniteLoopErc20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InfiniteLoopErc20Session struct {
	Contract     *InfiniteLoopErc20 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// InfiniteLoopErc20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InfiniteLoopErc20CallerSession struct {
	Contract *InfiniteLoopErc20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// InfiniteLoopErc20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InfiniteLoopErc20TransactorSession struct {
	Contract     *InfiniteLoopErc20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// InfiniteLoopErc20Raw is an auto generated low-level Go binding around an Ethereum contract.
type InfiniteLoopErc20Raw struct {
	Contract *InfiniteLoopErc20 // Generic contract binding to access the raw methods on
}

// InfiniteLoopErc20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InfiniteLoopErc20CallerRaw struct {
	Contract *InfiniteLoopErc20Caller // Generic read-only contract binding to access the raw methods on
}

// InfiniteLoopErc20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InfiniteLoopErc20TransactorRaw struct {
	Contract *InfiniteLoopErc20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewInfiniteLoopErc20 creates a new instance of InfiniteLoopErc20, bound to a specific deployed contract.
func NewInfiniteLoopErc20(address common.Address, backend bind.ContractBackend) (*InfiniteLoopErc20, error) {
	contract, err := bindInfiniteLoopErc20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &InfiniteLoopErc20{InfiniteLoopErc20Caller: InfiniteLoopErc20Caller{contract: contract}, InfiniteLoopErc20Transactor: InfiniteLoopErc20Transactor{contract: contract}, InfiniteLoopErc20Filterer: InfiniteLoopErc20Filterer{contract: contract}}, nil
}

// NewInfiniteLoopErc20Caller creates a new read-only instance of InfiniteLoopErc20, bound to a specific deployed contract.
func NewInfiniteLoopErc20Caller(address common.Address, caller bind.ContractCaller) (*InfiniteLoopErc20Caller, error) {
	contract, err := bindInfiniteLoopErc20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InfiniteLoopErc20Caller{contract: contract}, nil
}

// NewInfiniteLoopErc20Transactor creates a new write-only instance of InfiniteLoopErc20, bound to a specific deployed contract.
func NewInfiniteLoopErc20Transactor(address common.Address, transactor bind.ContractTransactor) (*InfiniteLoopErc20Transactor, error) {
	contract, err := bindInfiniteLoopErc20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InfiniteLoopErc20Transactor{contract: contract}, nil
}

// NewInfiniteLoopErc20Filterer creates a new log filterer instance of InfiniteLoopErc20, bound to a specific deployed contract.
func NewInfiniteLoopErc20Filterer(address common.Address, filterer bind.ContractFilterer) (*InfiniteLoopErc20Filterer, error) {
	contract, err := bindInfiniteLoopErc20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InfiniteLoopErc20Filterer{contract: contract}, nil
}

// bindInfiniteLoopErc20 binds a generic wrapper to an already deployed contract.
func bindInfiniteLoopErc20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := InfiniteLoopErc20MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InfiniteLoopErc20 *InfiniteLoopErc20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InfiniteLoopErc20.Contract.InfiniteLoopErc20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InfiniteLoopErc20 *InfiniteLoopErc20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InfiniteLoopErc20.Contract.InfiniteLoopErc20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InfiniteLoopErc20 *InfiniteLoopErc20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InfiniteLoopErc20.Contract.InfiniteLoopErc20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InfiniteLoopErc20 *InfiniteLoopErc20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InfiniteLoopErc20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InfiniteLoopErc20 *InfiniteLoopErc20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InfiniteLoopErc20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InfiniteLoopErc20 *InfiniteLoopErc20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InfiniteLoopErc20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_InfiniteLoopErc20 *InfiniteLoopErc20Caller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _InfiniteLoopErc20.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_InfiniteLoopErc20 *InfiniteLoopErc20Session) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _InfiniteLoopErc20.Contract.Allowance(&_InfiniteLoopErc20.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_InfiniteLoopErc20 *InfiniteLoopErc20CallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _InfiniteLoopErc20.Contract.Allowance(&_InfiniteLoopErc20.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_InfiniteLoopErc20 *InfiniteLoopErc20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _InfiniteLoopErc20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_InfiniteLoopErc20 *InfiniteLoopErc20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _InfiniteLoopErc20.Contract.BalanceOf(&_InfiniteLoopErc20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_InfiniteLoopErc20 *InfiniteLoopErc20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _InfiniteLoopErc20.Contract.BalanceOf(&_InfiniteLoopErc20.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_InfiniteLoopErc20 *InfiniteLoopErc20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _InfiniteLoopErc20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_InfiniteLoopErc20 *InfiniteLoopErc20Session) Decimals() (uint8, error) {
	return _InfiniteLoopErc20.Contract.Decimals(&_InfiniteLoopErc20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_InfiniteLoopErc20 *InfiniteLoopErc20CallerSession) Decimals() (uint8, error) {
	return _InfiniteLoopErc20.Contract.Decimals(&_InfiniteLoopErc20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_InfiniteLoopErc20 *InfiniteLoopErc20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _InfiniteLoopErc20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_InfiniteLoopErc20 *InfiniteLoopErc20Session) Name() (string, error) {
	return _InfiniteLoopErc20.Contract.Name(&_InfiniteLoopErc20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_InfiniteLoopErc20 *InfiniteLoopErc20CallerSession) Name() (string, error) {
	return _InfiniteLoopErc20.Contract.Name(&_InfiniteLoopErc20.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_InfiniteLoopErc20 *InfiniteLoopErc20Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _InfiniteLoopErc20.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_InfiniteLoopErc20 *InfiniteLoopErc20Session) Owner() (common.Address, error) {
	return _InfiniteLoopErc20.Contract.Owner(&_InfiniteLoopErc20.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_InfiniteLoopErc20 *InfiniteLoopErc20CallerSession) Owner() (common.Address, error) {
	return _InfiniteLoopErc20.Contract.Owner(&_InfiniteLoopErc20.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_InfiniteLoopErc20 *InfiniteLoopErc20Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _InfiniteLoopErc20.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_InfiniteLoopErc20 *InfiniteLoopErc20Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _InfiniteLoopErc20.Contract.SupportsInterface(&_InfiniteLoopErc20.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_InfiniteLoopErc20 *InfiniteLoopErc20CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _InfiniteLoopErc20.Contract.SupportsInterface(&_InfiniteLoopErc20.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_InfiniteLoopErc20 *InfiniteLoopErc20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _InfiniteLoopErc20.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_InfiniteLoopErc20 *InfiniteLoopErc20Session) Symbol() (string, error) {
	return _InfiniteLoopErc20.Contract.Symbol(&_InfiniteLoopErc20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_InfiniteLoopErc20 *InfiniteLoopErc20CallerSession) Symbol() (string, error) {
	return _InfiniteLoopErc20.Contract.Symbol(&_InfiniteLoopErc20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_InfiniteLoopErc20 *InfiniteLoopErc20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InfiniteLoopErc20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_InfiniteLoopErc20 *InfiniteLoopErc20Session) TotalSupply() (*big.Int, error) {
	return _InfiniteLoopErc20.Contract.TotalSupply(&_InfiniteLoopErc20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_InfiniteLoopErc20 *InfiniteLoopErc20CallerSession) TotalSupply() (*big.Int, error) {
	return _InfiniteLoopErc20.Contract.TotalSupply(&_InfiniteLoopErc20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_InfiniteLoopErc20 *InfiniteLoopErc20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _InfiniteLoopErc20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_InfiniteLoopErc20 *InfiniteLoopErc20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _InfiniteLoopErc20.Contract.Approve(&_InfiniteLoopErc20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_InfiniteLoopErc20 *InfiniteLoopErc20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _InfiniteLoopErc20.Contract.Approve(&_InfiniteLoopErc20.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_InfiniteLoopErc20 *InfiniteLoopErc20Transactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _InfiniteLoopErc20.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_InfiniteLoopErc20 *InfiniteLoopErc20Session) Burn(amount *big.Int) (*types.Transaction, error) {
	return _InfiniteLoopErc20.Contract.Burn(&_InfiniteLoopErc20.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_InfiniteLoopErc20 *InfiniteLoopErc20TransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _InfiniteLoopErc20.Contract.Burn(&_InfiniteLoopErc20.TransactOpts, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address from, uint256 amount) returns(bool)
func (_InfiniteLoopErc20 *InfiniteLoopErc20Transactor) BurnFrom(opts *bind.TransactOpts, from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _InfiniteLoopErc20.contract.Transact(opts, "burnFrom", from, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address from, uint256 amount) returns(bool)
func (_InfiniteLoopErc20 *InfiniteLoopErc20Session) BurnFrom(from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _InfiniteLoopErc20.Contract.BurnFrom(&_InfiniteLoopErc20.TransactOpts, from, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address from, uint256 amount) returns(bool)
func (_InfiniteLoopErc20 *InfiniteLoopErc20TransactorSession) BurnFrom(from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _InfiniteLoopErc20.Contract.BurnFrom(&_InfiniteLoopErc20.TransactOpts, from, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_InfiniteLoopErc20 *InfiniteLoopErc20Transactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _InfiniteLoopErc20.contract.Transact(opts, "mint", to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_InfiniteLoopErc20 *InfiniteLoopErc20Session) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _InfiniteLoopErc20.Contract.Mint(&_InfiniteLoopErc20.TransactOpts, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_InfiniteLoopErc20 *InfiniteLoopErc20TransactorSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _InfiniteLoopErc20.Contract.Mint(&_InfiniteLoopErc20.TransactOpts, to, amount)
}

// SudoTransfer is a paid mutator transaction binding the contract method 0x1988513b.
//
// Solidity: function sudoTransfer(address sender, address recipient, uint256 amount) returns()
func (_InfiniteLoopErc20 *InfiniteLoopErc20Transactor) SudoTransfer(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _InfiniteLoopErc20.contract.Transact(opts, "sudoTransfer", sender, recipient, amount)
}

// SudoTransfer is a paid mutator transaction binding the contract method 0x1988513b.
//
// Solidity: function sudoTransfer(address sender, address recipient, uint256 amount) returns()
func (_InfiniteLoopErc20 *InfiniteLoopErc20Session) SudoTransfer(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _InfiniteLoopErc20.Contract.SudoTransfer(&_InfiniteLoopErc20.TransactOpts, sender, recipient, amount)
}

// SudoTransfer is a paid mutator transaction binding the contract method 0x1988513b.
//
// Solidity: function sudoTransfer(address sender, address recipient, uint256 amount) returns()
func (_InfiniteLoopErc20 *InfiniteLoopErc20TransactorSession) SudoTransfer(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _InfiniteLoopErc20.Contract.SudoTransfer(&_InfiniteLoopErc20.TransactOpts, sender, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_InfiniteLoopErc20 *InfiniteLoopErc20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _InfiniteLoopErc20.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_InfiniteLoopErc20 *InfiniteLoopErc20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _InfiniteLoopErc20.Contract.Transfer(&_InfiniteLoopErc20.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_InfiniteLoopErc20 *InfiniteLoopErc20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _InfiniteLoopErc20.Contract.Transfer(&_InfiniteLoopErc20.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_InfiniteLoopErc20 *InfiniteLoopErc20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _InfiniteLoopErc20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_InfiniteLoopErc20 *InfiniteLoopErc20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _InfiniteLoopErc20.Contract.TransferFrom(&_InfiniteLoopErc20.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_InfiniteLoopErc20 *InfiniteLoopErc20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _InfiniteLoopErc20.Contract.TransferFrom(&_InfiniteLoopErc20.TransactOpts, sender, recipient, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_InfiniteLoopErc20 *InfiniteLoopErc20Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _InfiniteLoopErc20.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_InfiniteLoopErc20 *InfiniteLoopErc20Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _InfiniteLoopErc20.Contract.TransferOwnership(&_InfiniteLoopErc20.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_InfiniteLoopErc20 *InfiniteLoopErc20TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _InfiniteLoopErc20.Contract.TransferOwnership(&_InfiniteLoopErc20.TransactOpts, newOwner)
}

// InfiniteLoopErc20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the InfiniteLoopErc20 contract.
type InfiniteLoopErc20ApprovalIterator struct {
	Event *InfiniteLoopErc20Approval // Event containing the contract specifics and raw log

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
func (it *InfiniteLoopErc20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InfiniteLoopErc20Approval)
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
		it.Event = new(InfiniteLoopErc20Approval)
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
func (it *InfiniteLoopErc20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InfiniteLoopErc20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InfiniteLoopErc20Approval represents a Approval event raised by the InfiniteLoopErc20 contract.
type InfiniteLoopErc20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_InfiniteLoopErc20 *InfiniteLoopErc20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*InfiniteLoopErc20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _InfiniteLoopErc20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &InfiniteLoopErc20ApprovalIterator{contract: _InfiniteLoopErc20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_InfiniteLoopErc20 *InfiniteLoopErc20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *InfiniteLoopErc20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _InfiniteLoopErc20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InfiniteLoopErc20Approval)
				if err := _InfiniteLoopErc20.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_InfiniteLoopErc20 *InfiniteLoopErc20Filterer) ParseApproval(log types.Log) (*InfiniteLoopErc20Approval, error) {
	event := new(InfiniteLoopErc20Approval)
	if err := _InfiniteLoopErc20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InfiniteLoopErc20OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the InfiniteLoopErc20 contract.
type InfiniteLoopErc20OwnershipTransferredIterator struct {
	Event *InfiniteLoopErc20OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *InfiniteLoopErc20OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InfiniteLoopErc20OwnershipTransferred)
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
		it.Event = new(InfiniteLoopErc20OwnershipTransferred)
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
func (it *InfiniteLoopErc20OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InfiniteLoopErc20OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InfiniteLoopErc20OwnershipTransferred represents a OwnershipTransferred event raised by the InfiniteLoopErc20 contract.
type InfiniteLoopErc20OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_InfiniteLoopErc20 *InfiniteLoopErc20Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*InfiniteLoopErc20OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _InfiniteLoopErc20.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &InfiniteLoopErc20OwnershipTransferredIterator{contract: _InfiniteLoopErc20.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_InfiniteLoopErc20 *InfiniteLoopErc20Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *InfiniteLoopErc20OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _InfiniteLoopErc20.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InfiniteLoopErc20OwnershipTransferred)
				if err := _InfiniteLoopErc20.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_InfiniteLoopErc20 *InfiniteLoopErc20Filterer) ParseOwnershipTransferred(log types.Log) (*InfiniteLoopErc20OwnershipTransferred, error) {
	event := new(InfiniteLoopErc20OwnershipTransferred)
	if err := _InfiniteLoopErc20.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InfiniteLoopErc20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the InfiniteLoopErc20 contract.
type InfiniteLoopErc20TransferIterator struct {
	Event *InfiniteLoopErc20Transfer // Event containing the contract specifics and raw log

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
func (it *InfiniteLoopErc20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InfiniteLoopErc20Transfer)
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
		it.Event = new(InfiniteLoopErc20Transfer)
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
func (it *InfiniteLoopErc20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InfiniteLoopErc20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InfiniteLoopErc20Transfer represents a Transfer event raised by the InfiniteLoopErc20 contract.
type InfiniteLoopErc20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_InfiniteLoopErc20 *InfiniteLoopErc20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*InfiniteLoopErc20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _InfiniteLoopErc20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &InfiniteLoopErc20TransferIterator{contract: _InfiniteLoopErc20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_InfiniteLoopErc20 *InfiniteLoopErc20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *InfiniteLoopErc20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _InfiniteLoopErc20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InfiniteLoopErc20Transfer)
				if err := _InfiniteLoopErc20.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_InfiniteLoopErc20 *InfiniteLoopErc20Filterer) ParseTransfer(log types.Log) (*InfiniteLoopErc20Transfer, error) {
	event := new(InfiniteLoopErc20Transfer)
	if err := _InfiniteLoopErc20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
