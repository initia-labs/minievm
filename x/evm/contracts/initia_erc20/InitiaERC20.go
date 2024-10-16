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
	Bin: "0x608060405234801561000f575f80fd5b50604051611fcb380380611fcb833981810160405281019061003191906102a5565b335f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060f273ffffffffffffffffffffffffffffffffffffffff16635e6c57596040518163ffffffff1660e01b81526004016020604051808303815f875af11580156100bb573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906100df9190610362565b5082600390816100ef919061059a565b5081600490816100ff919061059a565b508060055f6101000a81548160ff021916908360ff160217905550505050610669565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6101818261013b565b810181811067ffffffffffffffff821117156101a05761019f61014b565b5b80604052505050565b5f6101b2610122565b90506101be8282610178565b919050565b5f67ffffffffffffffff8211156101dd576101dc61014b565b5b6101e68261013b565b9050602081019050919050565b8281835e5f83830152505050565b5f61021361020e846101c3565b6101a9565b90508281526020810184848401111561022f5761022e610137565b5b61023a8482856101f3565b509392505050565b5f82601f83011261025657610255610133565b5b8151610266848260208601610201565b91505092915050565b5f60ff82169050919050565b6102848161026f565b811461028e575f80fd5b50565b5f8151905061029f8161027b565b92915050565b5f805f606084860312156102bc576102bb61012b565b5b5f84015167ffffffffffffffff8111156102d9576102d861012f565b5b6102e586828701610242565b935050602084015167ffffffffffffffff8111156103065761030561012f565b5b61031286828701610242565b925050604061032386828701610291565b9150509250925092565b5f8115159050919050565b6103418161032d565b811461034b575f80fd5b50565b5f8151905061035c81610338565b92915050565b5f602082840312156103775761037661012b565b5b5f6103848482850161034e565b91505092915050565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806103db57607f821691505b6020821081036103ee576103ed610397565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026104507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610415565b61045a8683610415565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f61049e61049961049484610472565b61047b565b610472565b9050919050565b5f819050919050565b6104b783610484565b6104cb6104c3826104a5565b848454610421565b825550505050565b5f90565b6104df6104d3565b6104ea8184846104ae565b505050565b5b8181101561050d576105025f826104d7565b6001810190506104f0565b5050565b601f82111561055257610523816103f4565b61052c84610406565b8101602085101561053b578190505b61054f61054785610406565b8301826104ef565b50505b505050565b5f82821c905092915050565b5f6105725f1984600802610557565b1980831691505092915050565b5f61058a8383610563565b9150826002028217905092915050565b6105a38261038d565b67ffffffffffffffff8111156105bc576105bb61014b565b5b6105c682546103c4565b6105d1828285610511565b5f60209050601f831160018114610602575f84156105f0578287015190505b6105fa858261057f565b865550610661565b601f198416610610866103f4565b5f5b8281101561063757848901518255600182019150602085019450602081019050610612565b868310156106545784890151610650601f891682610563565b8355505b6001600288020188555050505b505050505050565b611955806106765f395ff3fe608060405234801561000f575f80fd5b50600436106100fe575f3560e01c806342966c681161009557806395d89b411161006457806395d89b41146102be578063a9059cbb146102dc578063dd62ed3e1461030c578063f2fde38b1461033c576100fe565b806342966c681461022457806370a082311461024057806379cc6790146102705780638da5cb5b146102a0576100fe565b80631988513b116100d15780631988513b1461019e57806323b872dd146101ba578063313ce567146101ea57806340c10f1914610208576100fe565b806301ffc9a71461010257806306fdde0314610132578063095ea7b31461015057806318160ddd14610180575b5f80fd5b61011c600480360381019061011791906112f3565b610358565b6040516101299190611338565b60405180910390f35b61013a6103d1565b60405161014791906113c1565b60405180910390f35b61016a6004803603810190610165919061146e565b61045d565b6040516101779190611338565b60405180910390f35b61018861054a565b60405161019591906114bb565b60405180910390f35b6101b860048036038101906101b391906114d4565b610550565b005b6101d460048036038101906101cf91906114d4565b6105cf565b6040516101e19190611338565b60405180910390f35b6101f261072f565b6040516101ff919061153f565b60405180910390f35b610222600480360381019061021d919061146e565b610741565b005b61023e60048036038101906102399190611558565b610860565b005b61025a60048036038101906102559190611583565b610928565b60405161026791906114bb565b60405180910390f35b61028a6004803603810190610285919061146e565b61093d565b6040516102979190611338565b60405180910390f35b6102a8610a9b565b6040516102b591906115bd565b60405180910390f35b6102c6610abe565b6040516102d391906113c1565b60405180910390f35b6102f660048036038101906102f1919061146e565b610b4a565b6040516103039190611338565b60405180910390f35b610326600480360381019061032191906115d6565b610c1b565b60405161033391906114bb565b60405180910390f35b61035660048036038101906103519190611583565b610c3b565b005b5f7f8da6da19000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614806103ca57506103c982610d83565b5b9050919050565b600380546103de90611641565b80601f016020809104026020016040519081016040528092919081815260200182805461040a90611641565b80156104555780601f1061042c57610100808354040283529160200191610455565b820191905f5260205f20905b81548152906001019060200180831161043857829003601f168201915b505050505081565b5f8160025f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9258460405161053891906114bb565b60405180910390a36001905092915050565b60065481565b600173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146105bf576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105b6906116bb565b60405180910390fd5b6105ca838383610dec565b505050565b5f8260f173ffffffffffffffffffffffffffffffffffffffff1663f2af9ac9826040518263ffffffff1660e01b815260040161060b91906115bd565b602060405180830381865afa158015610626573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061064a9190611703565b1561068a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106819061179e565b60405180910390fd5b8260025f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825461071191906117e9565b92505081905550610723858585610dec565b60019150509392505050565b60055f9054906101000a900460ff1681565b8160f173ffffffffffffffffffffffffffffffffffffffff1663f2af9ac9826040518263ffffffff1660e01b815260040161077c91906115bd565b602060405180830381865afa158015610797573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906107bb9190611703565b156107fb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107f290611866565b60405180910390fd5b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610851575f80fd5b61085b8383610ff7565b505050565b3360f173ffffffffffffffffffffffffffffffffffffffff166360dc402f826040518263ffffffff1660e01b815260040161089b91906115bd565b602060405180830381865afa1580156108b6573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906108da9190611703565b1561091a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610911906118ce565b60405180910390fd5b61092433836111c6565b5050565b6001602052805f5260405f205f915090505481565b5f8260f173ffffffffffffffffffffffffffffffffffffffff166360dc402f826040518263ffffffff1660e01b815260040161097991906115bd565b602060405180830381865afa158015610994573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906109b89190611703565b156109f8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109ef906118ce565b60405180910390fd5b8260025f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254610a7f91906117e9565b92505081905550610a9084846111c6565b600191505092915050565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60048054610acb90611641565b80601f0160208091040260200160405190810160405280929190818152602001828054610af790611641565b8015610b425780601f10610b1957610100808354040283529160200191610b42565b820191905f5260205f20905b815481529060010190602001808311610b2557829003601f168201915b505050505081565b5f8260f173ffffffffffffffffffffffffffffffffffffffff1663f2af9ac9826040518263ffffffff1660e01b8152600401610b8691906115bd565b602060405180830381865afa158015610ba1573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610bc59190611703565b15610c05576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610bfc9061179e565b60405180910390fd5b610c10338585610dec565b600191505092915050565b6002602052815f5260405f20602052805f5260405f205f91509150505481565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610c91575f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610cc8575f80fd5b8073ffffffffffffffffffffffffffffffffffffffff165f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3805f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b5f7f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b8160f273ffffffffffffffffffffffffffffffffffffffff16634e25ab64826040518263ffffffff1660e01b8152600401610e2791906115bd565b602060405180830381865afa158015610e42573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610e669190611703565b610ee65760f273ffffffffffffffffffffffffffffffffffffffff1663ceeae52a826040518263ffffffff1660e01b8152600401610ea491906115bd565b6020604051808303815f875af1158015610ec0573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610ee49190611703565b505b8160015f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254610f3291906117e9565b925050819055508160015f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254610f8591906118ec565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051610fe991906114bb565b60405180910390a350505050565b8160f273ffffffffffffffffffffffffffffffffffffffff16634e25ab64826040518263ffffffff1660e01b815260040161103291906115bd565b602060405180830381865afa15801561104d573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906110719190611703565b6110f15760f273ffffffffffffffffffffffffffffffffffffffff1663ceeae52a826040518263ffffffff1660e01b81526004016110af91906115bd565b6020604051808303815f875af11580156110cb573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906110ef9190611703565b505b8160015f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825461113d91906118ec565b925050819055508160065f82825461115591906118ec565b925050819055508273ffffffffffffffffffffffffffffffffffffffff165f73ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516111b991906114bb565b60405180910390a3505050565b8060015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825461121291906117e9565b925050819055508060065f82825461122a91906117e9565b925050819055505f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8360405161128e91906114bb565b60405180910390a35050565b5f80fd5b5f7fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b6112d28161129e565b81146112dc575f80fd5b50565b5f813590506112ed816112c9565b92915050565b5f602082840312156113085761130761129a565b5b5f611315848285016112df565b91505092915050565b5f8115159050919050565b6113328161131e565b82525050565b5f60208201905061134b5f830184611329565b92915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f61139382611351565b61139d818561135b565b93506113ad81856020860161136b565b6113b681611379565b840191505092915050565b5f6020820190508181035f8301526113d98184611389565b905092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61140a826113e1565b9050919050565b61141a81611400565b8114611424575f80fd5b50565b5f8135905061143581611411565b92915050565b5f819050919050565b61144d8161143b565b8114611457575f80fd5b50565b5f8135905061146881611444565b92915050565b5f80604083850312156114845761148361129a565b5b5f61149185828601611427565b92505060206114a28582860161145a565b9150509250929050565b6114b58161143b565b82525050565b5f6020820190506114ce5f8301846114ac565b92915050565b5f805f606084860312156114eb576114ea61129a565b5b5f6114f886828701611427565b935050602061150986828701611427565b925050604061151a8682870161145a565b9150509250925092565b5f60ff82169050919050565b61153981611524565b82525050565b5f6020820190506115525f830184611530565b92915050565b5f6020828403121561156d5761156c61129a565b5b5f61157a8482850161145a565b91505092915050565b5f602082840312156115985761159761129a565b5b5f6115a584828501611427565b91505092915050565b6115b781611400565b82525050565b5f6020820190506115d05f8301846115ae565b92915050565b5f80604083850312156115ec576115eb61129a565b5b5f6115f985828601611427565b925050602061160a85828601611427565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061165857607f821691505b60208210810361166b5761166a611614565b5b50919050565b7f45524332303a2063616c6c6572206973206e6f742074686520636861696e00005f82015250565b5f6116a5601e8361135b565b91506116b082611671565b602082019050919050565b5f6020820190508181035f8301526116d281611699565b9050919050565b6116e28161131e565b81146116ec575f80fd5b50565b5f815190506116fd816116d9565b92915050565b5f602082840312156117185761171761129a565b5b5f611725848285016116ef565b91505092915050565b7f45524332303a207472616e7366657220746f20626c6f636b65642061646472655f8201527f7373000000000000000000000000000000000000000000000000000000000000602082015250565b5f61178860228361135b565b91506117938261172e565b604082019050919050565b5f6020820190508181035f8301526117b58161177c565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6117f38261143b565b91506117fe8361143b565b9250828203905081811115611816576118156117bc565b5b92915050565b7f45524332303a206d696e7420746f20626c6f636b6564206164647265737300005f82015250565b5f611850601e8361135b565b915061185b8261181c565b602082019050919050565b5f6020820190508181035f83015261187d81611844565b9050919050565b7f45524332303a206275726e2066726f6d206d6f64756c652061646472657373005f82015250565b5f6118b8601f8361135b565b91506118c382611884565b602082019050919050565b5f6020820190508181035f8301526118e5816118ac565b9050919050565b5f6118f68261143b565b91506119018361143b565b9250828201905080821115611919576119186117bc565b5b9291505056fea2646970667358221220ab53c51a607edd6b66f7bf9d9198aa0a03bf8009df19db00175dc5fb94aee2d764736f6c63430008190033",
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
