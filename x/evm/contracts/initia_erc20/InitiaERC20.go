// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package initia_erc20

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

// InitiaErc20MetaData contains all meta data concerning the InitiaErc20 contract.
var InitiaErc20MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"_decimals\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"sudoTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b50604051612479380380612479833981810160405281019061003191906102a5565b335f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060f273ffffffffffffffffffffffffffffffffffffffff16635e6c57596040518163ffffffff1660e01b81526004016020604051808303815f875af11580156100bb573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906100df9190610362565b5082600390816100ef919061059a565b5081600490816100ff919061059a565b508060055f6101000a81548160ff021916908360ff160217905550505050610669565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6101818261013b565b810181811067ffffffffffffffff821117156101a05761019f61014b565b5b80604052505050565b5f6101b2610122565b90506101be8282610178565b919050565b5f67ffffffffffffffff8211156101dd576101dc61014b565b5b6101e68261013b565b9050602081019050919050565b8281835e5f83830152505050565b5f61021361020e846101c3565b6101a9565b90508281526020810184848401111561022f5761022e610137565b5b61023a8482856101f3565b509392505050565b5f82601f83011261025657610255610133565b5b8151610266848260208601610201565b91505092915050565b5f60ff82169050919050565b6102848161026f565b811461028e575f80fd5b50565b5f8151905061029f8161027b565b92915050565b5f805f606084860312156102bc576102bb61012b565b5b5f84015167ffffffffffffffff8111156102d9576102d861012f565b5b6102e586828701610242565b935050602084015167ffffffffffffffff8111156103065761030561012f565b5b61031286828701610242565b925050604061032386828701610291565b9150509250925092565b5f8115159050919050565b6103418161032d565b811461034b575f80fd5b50565b5f8151905061035c81610338565b92915050565b5f602082840312156103775761037661012b565b5b5f6103848482850161034e565b91505092915050565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806103db57607f821691505b6020821081036103ee576103ed610397565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026104507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610415565b61045a8683610415565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f61049e61049961049484610472565b61047b565b610472565b9050919050565b5f819050919050565b6104b783610484565b6104cb6104c3826104a5565b848454610421565b825550505050565b5f90565b6104df6104d3565b6104ea8184846104ae565b505050565b5b8181101561050d576105025f826104d7565b6001810190506104f0565b5050565b601f82111561055257610523816103f4565b61052c84610406565b8101602085101561053b578190505b61054f61054785610406565b8301826104ef565b50505b505050565b5f82821c905092915050565b5f6105725f1984600802610557565b1980831691505092915050565b5f61058a8383610563565b9150826002028217905092915050565b6105a38261038d565b67ffffffffffffffff8111156105bc576105bb61014b565b5b6105c682546103c4565b6105d1828285610511565b5f60209050601f831160018114610602575f84156105f0578287015190505b6105fa858261057f565b865550610661565b601f198416610610866103f4565b5f5b8281101561063757848901518255600182019150602085019450602081019050610612565b868310156106545784890151610650601f891682610563565b8355505b6001600288020188555050505b505050505050565b611e03806106765f395ff3fe608060405234801561000f575f80fd5b50600436106100fe575f3560e01c806342966c681161009557806395d89b411161006457806395d89b41146102be578063a9059cbb146102dc578063dd62ed3e1461030c578063f2fde38b1461033c576100fe565b806342966c681461022457806370a082311461024057806379cc6790146102705780638da5cb5b146102a0576100fe565b80631988513b116100d15780631988513b1461019e57806323b872dd146101ba578063313ce567146101ea57806340c10f1914610208576100fe565b806301ffc9a71461010257806306fdde0314610132578063095ea7b31461015057806318160ddd14610180575b5f80fd5b61011c60048036038101906101179190611569565b610358565b60405161012991906115ae565b60405180910390f35b61013a6103d1565b6040516101479190611637565b60405180910390f35b61016a600480360381019061016591906116e4565b61045d565b60405161017791906115ae565b60405180910390f35b61018861054a565b6040516101959190611731565b60405180910390f35b6101b860048036038101906101b3919061174a565b610550565b005b6101d460048036038101906101cf919061174a565b6105cf565b6040516101e191906115ae565b60405180910390f35b6101f26107ea565b6040516101ff91906117b5565b60405180910390f35b610222600480360381019061021d91906116e4565b6107fc565b005b61023e600480360381019061023991906117ce565b61091b565b005b61025a600480360381019061025591906117f9565b6109e3565b6040516102679190611731565b60405180910390f35b61028a600480360381019061028591906116e4565b6109f8565b60405161029791906115ae565b60405180910390f35b6102a8610c11565b6040516102b59190611833565b60405180910390f35b6102c6610c34565b6040516102d39190611637565b60405180910390f35b6102f660048036038101906102f191906116e4565b610cc0565b60405161030391906115ae565b60405180910390f35b6103266004803603810190610321919061184c565b610d91565b6040516103339190611731565b60405180910390f35b610356600480360381019061035191906117f9565b610db1565b005b5f7f8da6da19000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614806103ca57506103c982610ef9565b5b9050919050565b600380546103de906118b7565b80601f016020809104026020016040519081016040528092919081815260200182805461040a906118b7565b80156104555780601f1061042c57610100808354040283529160200191610455565b820191905f5260205f20905b81548152906001019060200180831161043857829003601f168201915b505050505081565b5f8160025f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040516105389190611731565b60405180910390a36001905092915050565b60065481565b600173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146105bf576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105b690611931565b60405180910390fd5b6105ca838383610f62565b505050565b5f8260f173ffffffffffffffffffffffffffffffffffffffff1663f2af9ac9826040518263ffffffff1660e01b815260040161060b9190611833565b602060405180830381865afa158015610626573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061064a9190611979565b1561068a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161068190611a14565b60405180910390fd5b8260025f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20541015610745576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161073c90611aa2565b60405180910390fd5b8260025f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546107cc9190611aed565b925050819055506107de858585610f62565b60019150509392505050565b60055f9054906101000a900460ff1681565b8160f173ffffffffffffffffffffffffffffffffffffffff1663f2af9ac9826040518263ffffffff1660e01b81526004016108379190611833565b602060405180830381865afa158015610852573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906108769190611979565b156108b6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108ad90611b6a565b60405180910390fd5b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461090c575f80fd5b61091683836111ed565b505050565b3360f173ffffffffffffffffffffffffffffffffffffffff166360dc402f826040518263ffffffff1660e01b81526004016109569190611833565b602060405180830381865afa158015610971573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906109959190611979565b156109d5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109cc90611bd2565b60405180910390fd5b6109df33836113bc565b5050565b6001602052805f5260405f205f915090505481565b5f8260f173ffffffffffffffffffffffffffffffffffffffff166360dc402f826040518263ffffffff1660e01b8152600401610a349190611833565b602060405180830381865afa158015610a4f573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610a739190611979565b15610ab3576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610aaa90611bd2565b60405180910390fd5b8260025f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20541015610b6e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b6590611c60565b60405180910390fd5b8260025f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254610bf59190611aed565b92505081905550610c0684846113bc565b600191505092915050565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60048054610c41906118b7565b80601f0160208091040260200160405190810160405280929190818152602001828054610c6d906118b7565b8015610cb85780601f10610c8f57610100808354040283529160200191610cb8565b820191905f5260205f20905b815481529060010190602001808311610c9b57829003601f168201915b505050505081565b5f8260f173ffffffffffffffffffffffffffffffffffffffff1663f2af9ac9826040518263ffffffff1660e01b8152600401610cfc9190611833565b602060405180830381865afa158015610d17573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610d3b9190611979565b15610d7b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d7290611a14565b60405180910390fd5b610d86338585610f62565b600191505092915050565b6002602052815f5260405f20602052805f5260405f205f91509150505481565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610e07575f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610e3e575f80fd5b8073ffffffffffffffffffffffffffffffffffffffff165f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3805f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b5f7f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b8160f273ffffffffffffffffffffffffffffffffffffffff16634e25ab64826040518263ffffffff1660e01b8152600401610f9d9190611833565b602060405180830381865afa158015610fb8573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610fdc9190611979565b61105c5760f273ffffffffffffffffffffffffffffffffffffffff1663ceeae52a826040518263ffffffff1660e01b815260040161101a9190611833565b6020604051808303815f875af1158015611036573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061105a9190611979565b505b8160015f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205410156110dc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110d390611cee565b60405180910390fd5b8160015f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546111289190611aed565b925050819055508160015f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825461117b9190611d0c565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516111df9190611731565b60405180910390a350505050565b8160f273ffffffffffffffffffffffffffffffffffffffff16634e25ab64826040518263ffffffff1660e01b81526004016112289190611833565b602060405180830381865afa158015611243573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906112679190611979565b6112e75760f273ffffffffffffffffffffffffffffffffffffffff1663ceeae52a826040518263ffffffff1660e01b81526004016112a59190611833565b6020604051808303815f875af11580156112c1573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906112e59190611979565b505b8160015f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546113339190611d0c565b925050819055508160065f82825461134b9190611d0c565b925050819055508273ffffffffffffffffffffffffffffffffffffffff165f73ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516113af9190611731565b60405180910390a3505050565b8060015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054101561143c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161143390611daf565b60405180910390fd5b8060015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546114889190611aed565b925050819055508060065f8282546114a09190611aed565b925050819055505f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516115049190611731565b60405180910390a35050565b5f80fd5b5f7fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b61154881611514565b8114611552575f80fd5b50565b5f813590506115638161153f565b92915050565b5f6020828403121561157e5761157d611510565b5b5f61158b84828501611555565b91505092915050565b5f8115159050919050565b6115a881611594565b82525050565b5f6020820190506115c15f83018461159f565b92915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f611609826115c7565b61161381856115d1565b93506116238185602086016115e1565b61162c816115ef565b840191505092915050565b5f6020820190508181035f83015261164f81846115ff565b905092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61168082611657565b9050919050565b61169081611676565b811461169a575f80fd5b50565b5f813590506116ab81611687565b92915050565b5f819050919050565b6116c3816116b1565b81146116cd575f80fd5b50565b5f813590506116de816116ba565b92915050565b5f80604083850312156116fa576116f9611510565b5b5f6117078582860161169d565b9250506020611718858286016116d0565b9150509250929050565b61172b816116b1565b82525050565b5f6020820190506117445f830184611722565b92915050565b5f805f6060848603121561176157611760611510565b5b5f61176e8682870161169d565b935050602061177f8682870161169d565b9250506040611790868287016116d0565b9150509250925092565b5f60ff82169050919050565b6117af8161179a565b82525050565b5f6020820190506117c85f8301846117a6565b92915050565b5f602082840312156117e3576117e2611510565b5b5f6117f0848285016116d0565b91505092915050565b5f6020828403121561180e5761180d611510565b5b5f61181b8482850161169d565b91505092915050565b61182d81611676565b82525050565b5f6020820190506118465f830184611824565b92915050565b5f806040838503121561186257611861611510565b5b5f61186f8582860161169d565b92505060206118808582860161169d565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806118ce57607f821691505b6020821081036118e1576118e061188a565b5b50919050565b7f45524332303a2063616c6c6572206973206e6f742074686520636861696e00005f82015250565b5f61191b601e836115d1565b9150611926826118e7565b602082019050919050565b5f6020820190508181035f8301526119488161190f565b9050919050565b61195881611594565b8114611962575f80fd5b50565b5f815190506119738161194f565b92915050565b5f6020828403121561198e5761198d611510565b5b5f61199b84828501611965565b91505092915050565b7f45524332303a207472616e7366657220746f20626c6f636b65642061646472655f8201527f7373000000000000000000000000000000000000000000000000000000000000602082015250565b5f6119fe6022836115d1565b9150611a09826119a4565b604082019050919050565b5f6020820190508181035f830152611a2b816119f2565b9050919050565b7f45524332303a207472616e7366657220616d6f756e74206578636565647320615f8201527f6c6c6f77616e6365000000000000000000000000000000000000000000000000602082015250565b5f611a8c6028836115d1565b9150611a9782611a32565b604082019050919050565b5f6020820190508181035f830152611ab981611a80565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f611af7826116b1565b9150611b02836116b1565b9250828203905081811115611b1a57611b19611ac0565b5b92915050565b7f45524332303a206d696e7420746f20626c6f636b6564206164647265737300005f82015250565b5f611b54601e836115d1565b9150611b5f82611b20565b602082019050919050565b5f6020820190508181035f830152611b8181611b48565b9050919050565b7f45524332303a206275726e2066726f6d206d6f64756c652061646472657373005f82015250565b5f611bbc601f836115d1565b9150611bc782611b88565b602082019050919050565b5f6020820190508181035f830152611be981611bb0565b9050919050565b7f45524332303a206275726e20616d6f756e74206578636565647320616c6c6f775f8201527f616e636500000000000000000000000000000000000000000000000000000000602082015250565b5f611c4a6024836115d1565b9150611c5582611bf0565b604082019050919050565b5f6020820190508181035f830152611c7781611c3e565b9050919050565b7f45524332303a207472616e7366657220616d6f756e74206578636565647320625f8201527f616c616e63650000000000000000000000000000000000000000000000000000602082015250565b5f611cd86026836115d1565b9150611ce382611c7e565b604082019050919050565b5f6020820190508181035f830152611d0581611ccc565b9050919050565b5f611d16826116b1565b9150611d21836116b1565b9250828201905080821115611d3957611d38611ac0565b5b92915050565b7f45524332303a206275726e20616d6f756e7420657863656564732062616c616e5f8201527f6365000000000000000000000000000000000000000000000000000000000000602082015250565b5f611d996022836115d1565b9150611da482611d3f565b604082019050919050565b5f6020820190508181035f830152611dc681611d8d565b905091905056fea264697066735822122066dbab623b3de8459f53ebe5e3cb35fbd415d9fb7487a398b5e957a1692f6f4c64736f6c63430008190033",
}

// InitiaErc20ABI is the input ABI used to generate the binding from.
// Deprecated: Use InitiaErc20MetaData.ABI instead.
var InitiaErc20ABI = InitiaErc20MetaData.ABI

// InitiaErc20Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use InitiaErc20MetaData.Bin instead.
var InitiaErc20Bin = InitiaErc20MetaData.Bin

// DeployInitiaErc20 deploys a new Ethereum contract, binding an instance of InitiaErc20 to it.
func DeployInitiaErc20(auth *bind.TransactOpts, backend bind.ContractBackend, _name string, _symbol string, _decimals uint8) (common.Address, *types.Transaction, *InitiaErc20, error) {
	parsed, err := InitiaErc20MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(InitiaErc20Bin), backend, _name, _symbol, _decimals)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &InitiaErc20{InitiaErc20Caller: InitiaErc20Caller{contract: contract}, InitiaErc20Transactor: InitiaErc20Transactor{contract: contract}, InitiaErc20Filterer: InitiaErc20Filterer{contract: contract}}, nil
}

// InitiaErc20 is an auto generated Go binding around an Ethereum contract.
type InitiaErc20 struct {
	InitiaErc20Caller     // Read-only binding to the contract
	InitiaErc20Transactor // Write-only binding to the contract
	InitiaErc20Filterer   // Log filterer for contract events
}

// InitiaErc20Caller is an auto generated read-only Go binding around an Ethereum contract.
type InitiaErc20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InitiaErc20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type InitiaErc20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InitiaErc20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InitiaErc20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InitiaErc20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InitiaErc20Session struct {
	Contract     *InitiaErc20      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InitiaErc20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InitiaErc20CallerSession struct {
	Contract *InitiaErc20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// InitiaErc20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InitiaErc20TransactorSession struct {
	Contract     *InitiaErc20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// InitiaErc20Raw is an auto generated low-level Go binding around an Ethereum contract.
type InitiaErc20Raw struct {
	Contract *InitiaErc20 // Generic contract binding to access the raw methods on
}

// InitiaErc20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InitiaErc20CallerRaw struct {
	Contract *InitiaErc20Caller // Generic read-only contract binding to access the raw methods on
}

// InitiaErc20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InitiaErc20TransactorRaw struct {
	Contract *InitiaErc20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewInitiaErc20 creates a new instance of InitiaErc20, bound to a specific deployed contract.
func NewInitiaErc20(address common.Address, backend bind.ContractBackend) (*InitiaErc20, error) {
	contract, err := bindInitiaErc20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &InitiaErc20{InitiaErc20Caller: InitiaErc20Caller{contract: contract}, InitiaErc20Transactor: InitiaErc20Transactor{contract: contract}, InitiaErc20Filterer: InitiaErc20Filterer{contract: contract}}, nil
}

// NewInitiaErc20Caller creates a new read-only instance of InitiaErc20, bound to a specific deployed contract.
func NewInitiaErc20Caller(address common.Address, caller bind.ContractCaller) (*InitiaErc20Caller, error) {
	contract, err := bindInitiaErc20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InitiaErc20Caller{contract: contract}, nil
}

// NewInitiaErc20Transactor creates a new write-only instance of InitiaErc20, bound to a specific deployed contract.
func NewInitiaErc20Transactor(address common.Address, transactor bind.ContractTransactor) (*InitiaErc20Transactor, error) {
	contract, err := bindInitiaErc20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InitiaErc20Transactor{contract: contract}, nil
}

// NewInitiaErc20Filterer creates a new log filterer instance of InitiaErc20, bound to a specific deployed contract.
func NewInitiaErc20Filterer(address common.Address, filterer bind.ContractFilterer) (*InitiaErc20Filterer, error) {
	contract, err := bindInitiaErc20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InitiaErc20Filterer{contract: contract}, nil
}

// bindInitiaErc20 binds a generic wrapper to an already deployed contract.
func bindInitiaErc20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := InitiaErc20MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InitiaErc20 *InitiaErc20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InitiaErc20.Contract.InitiaErc20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InitiaErc20 *InitiaErc20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InitiaErc20.Contract.InitiaErc20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InitiaErc20 *InitiaErc20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InitiaErc20.Contract.InitiaErc20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InitiaErc20 *InitiaErc20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InitiaErc20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InitiaErc20 *InitiaErc20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InitiaErc20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InitiaErc20 *InitiaErc20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InitiaErc20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_InitiaErc20 *InitiaErc20Caller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _InitiaErc20.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_InitiaErc20 *InitiaErc20Session) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _InitiaErc20.Contract.Allowance(&_InitiaErc20.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_InitiaErc20 *InitiaErc20CallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _InitiaErc20.Contract.Allowance(&_InitiaErc20.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_InitiaErc20 *InitiaErc20Caller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _InitiaErc20.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_InitiaErc20 *InitiaErc20Session) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _InitiaErc20.Contract.BalanceOf(&_InitiaErc20.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_InitiaErc20 *InitiaErc20CallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _InitiaErc20.Contract.BalanceOf(&_InitiaErc20.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_InitiaErc20 *InitiaErc20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _InitiaErc20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_InitiaErc20 *InitiaErc20Session) Decimals() (uint8, error) {
	return _InitiaErc20.Contract.Decimals(&_InitiaErc20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_InitiaErc20 *InitiaErc20CallerSession) Decimals() (uint8, error) {
	return _InitiaErc20.Contract.Decimals(&_InitiaErc20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_InitiaErc20 *InitiaErc20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _InitiaErc20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_InitiaErc20 *InitiaErc20Session) Name() (string, error) {
	return _InitiaErc20.Contract.Name(&_InitiaErc20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_InitiaErc20 *InitiaErc20CallerSession) Name() (string, error) {
	return _InitiaErc20.Contract.Name(&_InitiaErc20.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_InitiaErc20 *InitiaErc20Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _InitiaErc20.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_InitiaErc20 *InitiaErc20Session) Owner() (common.Address, error) {
	return _InitiaErc20.Contract.Owner(&_InitiaErc20.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_InitiaErc20 *InitiaErc20CallerSession) Owner() (common.Address, error) {
	return _InitiaErc20.Contract.Owner(&_InitiaErc20.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_InitiaErc20 *InitiaErc20Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _InitiaErc20.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_InitiaErc20 *InitiaErc20Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _InitiaErc20.Contract.SupportsInterface(&_InitiaErc20.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_InitiaErc20 *InitiaErc20CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _InitiaErc20.Contract.SupportsInterface(&_InitiaErc20.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_InitiaErc20 *InitiaErc20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _InitiaErc20.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_InitiaErc20 *InitiaErc20Session) Symbol() (string, error) {
	return _InitiaErc20.Contract.Symbol(&_InitiaErc20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_InitiaErc20 *InitiaErc20CallerSession) Symbol() (string, error) {
	return _InitiaErc20.Contract.Symbol(&_InitiaErc20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_InitiaErc20 *InitiaErc20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InitiaErc20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_InitiaErc20 *InitiaErc20Session) TotalSupply() (*big.Int, error) {
	return _InitiaErc20.Contract.TotalSupply(&_InitiaErc20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_InitiaErc20 *InitiaErc20CallerSession) TotalSupply() (*big.Int, error) {
	return _InitiaErc20.Contract.TotalSupply(&_InitiaErc20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_InitiaErc20 *InitiaErc20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _InitiaErc20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_InitiaErc20 *InitiaErc20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _InitiaErc20.Contract.Approve(&_InitiaErc20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_InitiaErc20 *InitiaErc20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _InitiaErc20.Contract.Approve(&_InitiaErc20.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_InitiaErc20 *InitiaErc20Transactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _InitiaErc20.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_InitiaErc20 *InitiaErc20Session) Burn(amount *big.Int) (*types.Transaction, error) {
	return _InitiaErc20.Contract.Burn(&_InitiaErc20.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_InitiaErc20 *InitiaErc20TransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _InitiaErc20.Contract.Burn(&_InitiaErc20.TransactOpts, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address from, uint256 amount) returns(bool)
func (_InitiaErc20 *InitiaErc20Transactor) BurnFrom(opts *bind.TransactOpts, from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _InitiaErc20.contract.Transact(opts, "burnFrom", from, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address from, uint256 amount) returns(bool)
func (_InitiaErc20 *InitiaErc20Session) BurnFrom(from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _InitiaErc20.Contract.BurnFrom(&_InitiaErc20.TransactOpts, from, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address from, uint256 amount) returns(bool)
func (_InitiaErc20 *InitiaErc20TransactorSession) BurnFrom(from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _InitiaErc20.Contract.BurnFrom(&_InitiaErc20.TransactOpts, from, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_InitiaErc20 *InitiaErc20Transactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _InitiaErc20.contract.Transact(opts, "mint", to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_InitiaErc20 *InitiaErc20Session) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _InitiaErc20.Contract.Mint(&_InitiaErc20.TransactOpts, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_InitiaErc20 *InitiaErc20TransactorSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _InitiaErc20.Contract.Mint(&_InitiaErc20.TransactOpts, to, amount)
}

// SudoTransfer is a paid mutator transaction binding the contract method 0x1988513b.
//
// Solidity: function sudoTransfer(address sender, address recipient, uint256 amount) returns()
func (_InitiaErc20 *InitiaErc20Transactor) SudoTransfer(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _InitiaErc20.contract.Transact(opts, "sudoTransfer", sender, recipient, amount)
}

// SudoTransfer is a paid mutator transaction binding the contract method 0x1988513b.
//
// Solidity: function sudoTransfer(address sender, address recipient, uint256 amount) returns()
func (_InitiaErc20 *InitiaErc20Session) SudoTransfer(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _InitiaErc20.Contract.SudoTransfer(&_InitiaErc20.TransactOpts, sender, recipient, amount)
}

// SudoTransfer is a paid mutator transaction binding the contract method 0x1988513b.
//
// Solidity: function sudoTransfer(address sender, address recipient, uint256 amount) returns()
func (_InitiaErc20 *InitiaErc20TransactorSession) SudoTransfer(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _InitiaErc20.Contract.SudoTransfer(&_InitiaErc20.TransactOpts, sender, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_InitiaErc20 *InitiaErc20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _InitiaErc20.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_InitiaErc20 *InitiaErc20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _InitiaErc20.Contract.Transfer(&_InitiaErc20.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_InitiaErc20 *InitiaErc20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _InitiaErc20.Contract.Transfer(&_InitiaErc20.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_InitiaErc20 *InitiaErc20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _InitiaErc20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_InitiaErc20 *InitiaErc20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _InitiaErc20.Contract.TransferFrom(&_InitiaErc20.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_InitiaErc20 *InitiaErc20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _InitiaErc20.Contract.TransferFrom(&_InitiaErc20.TransactOpts, sender, recipient, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_InitiaErc20 *InitiaErc20Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _InitiaErc20.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_InitiaErc20 *InitiaErc20Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _InitiaErc20.Contract.TransferOwnership(&_InitiaErc20.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_InitiaErc20 *InitiaErc20TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _InitiaErc20.Contract.TransferOwnership(&_InitiaErc20.TransactOpts, newOwner)
}

// InitiaErc20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the InitiaErc20 contract.
type InitiaErc20ApprovalIterator struct {
	Event *InitiaErc20Approval // Event containing the contract specifics and raw log

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
func (it *InitiaErc20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InitiaErc20Approval)
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
		it.Event = new(InitiaErc20Approval)
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
func (it *InitiaErc20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InitiaErc20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InitiaErc20Approval represents a Approval event raised by the InitiaErc20 contract.
type InitiaErc20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_InitiaErc20 *InitiaErc20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*InitiaErc20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _InitiaErc20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &InitiaErc20ApprovalIterator{contract: _InitiaErc20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_InitiaErc20 *InitiaErc20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *InitiaErc20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _InitiaErc20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InitiaErc20Approval)
				if err := _InitiaErc20.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_InitiaErc20 *InitiaErc20Filterer) ParseApproval(log types.Log) (*InitiaErc20Approval, error) {
	event := new(InitiaErc20Approval)
	if err := _InitiaErc20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InitiaErc20OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the InitiaErc20 contract.
type InitiaErc20OwnershipTransferredIterator struct {
	Event *InitiaErc20OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *InitiaErc20OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InitiaErc20OwnershipTransferred)
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
		it.Event = new(InitiaErc20OwnershipTransferred)
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
func (it *InitiaErc20OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InitiaErc20OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InitiaErc20OwnershipTransferred represents a OwnershipTransferred event raised by the InitiaErc20 contract.
type InitiaErc20OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_InitiaErc20 *InitiaErc20Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*InitiaErc20OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _InitiaErc20.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &InitiaErc20OwnershipTransferredIterator{contract: _InitiaErc20.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_InitiaErc20 *InitiaErc20Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *InitiaErc20OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _InitiaErc20.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InitiaErc20OwnershipTransferred)
				if err := _InitiaErc20.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_InitiaErc20 *InitiaErc20Filterer) ParseOwnershipTransferred(log types.Log) (*InitiaErc20OwnershipTransferred, error) {
	event := new(InitiaErc20OwnershipTransferred)
	if err := _InitiaErc20.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InitiaErc20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the InitiaErc20 contract.
type InitiaErc20TransferIterator struct {
	Event *InitiaErc20Transfer // Event containing the contract specifics and raw log

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
func (it *InitiaErc20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InitiaErc20Transfer)
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
		it.Event = new(InitiaErc20Transfer)
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
func (it *InitiaErc20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InitiaErc20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InitiaErc20Transfer represents a Transfer event raised by the InitiaErc20 contract.
type InitiaErc20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_InitiaErc20 *InitiaErc20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*InitiaErc20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _InitiaErc20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &InitiaErc20TransferIterator{contract: _InitiaErc20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_InitiaErc20 *InitiaErc20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *InitiaErc20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _InitiaErc20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InitiaErc20Transfer)
				if err := _InitiaErc20.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_InitiaErc20 *InitiaErc20Filterer) ParseTransfer(log types.Log) (*InitiaErc20Transfer, error) {
	event := new(InitiaErc20Transfer)
	if err := _InitiaErc20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
