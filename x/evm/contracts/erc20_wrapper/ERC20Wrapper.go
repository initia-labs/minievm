// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package erc20_wrapper

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

// Erc20WrapperMetaData contains all meta data concerning the Erc20Wrapper contract.
var Erc20WrapperMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc20Factory\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"StringsInsufficientHexLength\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"contractERC20Factory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"callback_id\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"name\":\"ibc_ack\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"callback_id\",\"type\":\"uint64\"}],\"name\":\"ibc_timeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"originToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wrappedAmt\",\"type\":\"uint256\"}],\"name\":\"unwrap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"channel\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"receiver\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"}],\"name\":\"wrap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wrappedTokens\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a06040525f5f60146101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550348015610037575f5ffd5b50604051612d27380380612d2783398181016040528101906100599190610130565b335f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff16815250505061015b565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6100ff826100d6565b9050919050565b61010f816100f5565b8114610119575f5ffd5b50565b5f8151905061012a81610106565b92915050565b5f60208284031215610145576101446100d2565b5b5f6101528482850161011c565b91505092915050565b608051612bad61017a5f395f81816109a80152610fd30152612bad5ff3fe608060405234801561000f575f5ffd5b5060043610610091575f3560e01c80638da5cb5b116100645780638da5cb5b146101195780639a11143214610137578063c45a015514610153578063d5c6b50414610171578063f2fde38b146101a157610091565b806301ffc9a7146100955780630d4f1f9d146100c557806331a503f0146100e15780638cc7104f146100fd575b5f5ffd5b6100af60048036038101906100aa91906118d9565b6101bd565b6040516100bc919061191e565b60405180910390f35b6100df60048036038101906100da919061199e565b610226565b005b6100fb60048036038101906100f691906119dc565b6102a7565b005b61011760048036038101906101129190611a94565b610321565b005b610121610569565b60405161012e9190611af3565b60405180910390f35b610151600480360381019061014c9190611c48565b61058d565b005b61015b6109a6565b6040516101689190611d52565b60405180910390f35b61018b60048036038101906101869190611d6b565b6109ca565b6040516101989190611af3565b60405180910390f35b6101bb60048036038101906101b69190611d6b565b6109fa565b005b5f7f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b3073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610294576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161028b90611e16565b60405180910390fd5b806102a3576102a282610b44565b5b5050565b3073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610315576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161030c90611e16565b60405180910390fd5b61031e81610b44565b50565b5f60015f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036103ef576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103e690611e7e565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff166379cc679033846040518363ffffffff1660e01b815260040161042a929190611eab565b6020604051808303815f875af1158015610446573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061046a9190611ee6565b505f6104e48360068773ffffffffffffffffffffffffffffffffffffffff1663313ce5676040518163ffffffff1660e01b8152600401602060405180830381865afa1580156104bb573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906104df9190611f47565b610e72565b90508473ffffffffffffffffffffffffffffffffffffffff1663a9059cbb85836040518363ffffffff1660e01b8152600401610521929190611eab565b6020604051808303815f875af115801561053d573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906105619190611ee6565b505050505050565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b61059684610f41565b8373ffffffffffffffffffffffffffffffffffffffff166323b872dd3330856040518463ffffffff1660e01b81526004016105d393929190611f72565b6020604051808303815f875af11580156105ef573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906106139190611ee6565b505f61068d838673ffffffffffffffffffffffffffffffffffffffff1663313ce5676040518163ffffffff1660e01b8152600401602060405180830381865afa158015610662573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906106869190611f47565b6006610e72565b905060015f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166340c10f1930836040518363ffffffff1660e01b8152600401610726929190611eab565b5f604051808303815f87803b15801561073d575f5ffd5b505af115801561074f573d5f5f3e3d5ffd5b5050505060015f60148282829054906101000a900467ffffffffffffffff166107789190611fd4565b92506101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060405180606001604052803373ffffffffffffffffffffffffffffffffffffffff1681526020018673ffffffffffffffffffffffffffffffffffffffff1681526020018281525060025f5f60149054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f820151815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550604082015181600201559050505f6109208760015f8973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1684868961127b565b905060f173ffffffffffffffffffffffffffffffffffffffff1663d46f64e6826040518263ffffffff1660e01b815260040161095c919061205f565b6020604051808303815f875af1158015610978573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061099c9190611ee6565b5050505050505050565b7f000000000000000000000000000000000000000000000000000000000000000081565b6001602052805f5260405f205f915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610a51575f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610a88575f5ffd5b8073ffffffffffffffffffffffffffffffffffffffff165f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3805f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b5f60025f8367ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f206040518060600160405290815f82015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600182015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160028201548152505090505f60015f836020015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610cfc576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610cf390611e7e565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff166342966c6883604001516040518263ffffffff1660e01b8152600401610d39919061207f565b5f604051808303815f87803b158015610d50575f5ffd5b505af1158015610d62573d5f5f3e3d5ffd5b505050505f610de783604001516006856020015173ffffffffffffffffffffffffffffffffffffffff1663313ce5676040518163ffffffff1660e01b8152600401602060405180830381865afa158015610dbe573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610de29190611f47565b610e72565b9050826020015173ffffffffffffffffffffffffffffffffffffffff1663a9059cbb845f0151836040518363ffffffff1660e01b8152600401610e2b929190611eab565b6020604051808303815f875af1158015610e47573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610e6b9190611ee6565b5050505050565b5f8160ff168360ff161115610eb3575f8284610e8e9190612098565b60ff16600a610e9d91906121fb565b90508085610eab9190612272565b915050610ef8565b8160ff168360ff161015610ef3575f8383610ece9190612098565b60ff16600a610edd91906121fb565b90508085610eeb91906122a2565b915050610ef7565b8390505b5b5f8103610f3a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f319061232d565b60405180910390fd5b9392505050565b5f73ffffffffffffffffffffffffffffffffffffffff1660015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603611278575f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166306ef1a866040518060400160405280600781526020017f57726170706564000000000000000000000000000000000000000000000000008152508473ffffffffffffffffffffffffffffffffffffffff166306fdde036040518163ffffffff1660e01b81526004015f60405180830381865afa15801561108b573d5f5f3e3d5ffd5b505050506040513d5f823e3d601f19601f820116820180604052508101906110b391906123b9565b6040516020016110c492919061243a565b6040516020818303038152906040526040518060400160405280600181526020017f57000000000000000000000000000000000000000000000000000000000000008152508573ffffffffffffffffffffffffffffffffffffffff166395d89b416040518163ffffffff1660e01b81526004015f60405180830381865afa158015611151573d5f5f3e3d5ffd5b505050506040513d5f823e3d601f19601f8201168201806040525081019061117991906123b9565b60405160200161118a92919061243a565b60405160208183030381529060405260066040518463ffffffff1660e01b81526004016111b99392919061246c565b6020604051808303815f875af11580156111d5573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906111f991906124c3565b90508060015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505b50565b60608560f173ffffffffffffffffffffffffffffffffffffffff166381cf0f6a876040518263ffffffff1660e01b81526004016112b89190611af3565b5f604051808303815f875af11580156112d3573d5f5f3e3d5ffd5b505050506040513d5f823e3d601f19601f820116820180604052508101906112fb91906123b9565b611304866113ed565b60f173ffffffffffffffffffffffffffffffffffffffff16636af32a55306040518263ffffffff1660e01b815260040161133e9190611af3565b5f604051808303815f875af1158015611359573d5f5f3e3d5ffd5b505050506040513d5f823e3d601f19601f8201168201806040525081019061138191906123b9565b8561138b886113ed565b6113b35f60149054906101000a900467ffffffffffffffff1667ffffffffffffffff166113ed565b6113bc306114b7565b6040516020016113d3989796959493929190612992565b604051602081830303815290604052905095945050505050565b60605f60016113fb846114e4565b0190505f8167ffffffffffffffff81111561141957611418611b24565b5b6040519080825280601f01601f19166020018201604052801561144b5781602001600182028036833780820191505090505b5090505f82602001820190505b6001156114ac578080600190039150507f3031323334353637383961626364656600000000000000000000000000000000600a86061a8153600a85816114a1576114a0612245565b5b0494505f8503611458575b819350505050919050565b60606114dd8273ffffffffffffffffffffffffffffffffffffffff16601460ff16611635565b9050919050565b5f5f5f90507a184f03e93ff9f4daa797ed6e38ed64bf6a1f0100000000000000008310611540577a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000838161153657611535612245565b5b0492506040810190505b6d04ee2d6d415b85acef8100000000831061157d576d04ee2d6d415b85acef8100000000838161157357611572612245565b5b0492506020810190505b662386f26fc1000083106115ac57662386f26fc1000083816115a2576115a1612245565b5b0492506010810190505b6305f5e10083106115d5576305f5e10083816115cb576115ca612245565b5b0492506008810190505b61271083106115fa5761271083816115f0576115ef612245565b5b0492506004810190505b6064831061161d576064838161161357611612612245565b5b0492506002810190505b600a831061162c576001810190505b80915050919050565b60605f8390505f600284600261164b91906122a2565b6116559190612ac9565b67ffffffffffffffff81111561166e5761166d611b24565b5b6040519080825280601f01601f1916602001820160405280156116a05781602001600182028036833780820191505090505b5090507f3000000000000000000000000000000000000000000000000000000000000000815f815181106116d7576116d6612afc565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a9053507f78000000000000000000000000000000000000000000000000000000000000008160018151811061173a57611739612afc565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a9053505f600185600261177891906122a2565b6117829190612ac9565b90505b6001811115611821577f3031323334353637383961626364656600000000000000000000000000000000600f8416601081106117c4576117c3612afc565b5b1a60f81b8282815181106117db576117da612afc565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a905350600483901c92508061181a90612b29565b9050611785565b505f82146118685784846040517fe22e27eb00000000000000000000000000000000000000000000000000000000815260040161185f929190612b50565b60405180910390fd5b809250505092915050565b5f604051905090565b5f5ffd5b5f5ffd5b5f7fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b6118b881611884565b81146118c2575f5ffd5b50565b5f813590506118d3816118af565b92915050565b5f602082840312156118ee576118ed61187c565b5b5f6118fb848285016118c5565b91505092915050565b5f8115159050919050565b61191881611904565b82525050565b5f6020820190506119315f83018461190f565b92915050565b5f67ffffffffffffffff82169050919050565b61195381611937565b811461195d575f5ffd5b50565b5f8135905061196e8161194a565b92915050565b61197d81611904565b8114611987575f5ffd5b50565b5f8135905061199881611974565b92915050565b5f5f604083850312156119b4576119b361187c565b5b5f6119c185828601611960565b92505060206119d28582860161198a565b9150509250929050565b5f602082840312156119f1576119f061187c565b5b5f6119fe84828501611960565b91505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f611a3082611a07565b9050919050565b611a4081611a26565b8114611a4a575f5ffd5b50565b5f81359050611a5b81611a37565b92915050565b5f819050919050565b611a7381611a61565b8114611a7d575f5ffd5b50565b5f81359050611a8e81611a6a565b92915050565b5f5f5f60608486031215611aab57611aaa61187c565b5b5f611ab886828701611a4d565b9350506020611ac986828701611a4d565b9250506040611ada86828701611a80565b9150509250925092565b611aed81611a26565b82525050565b5f602082019050611b065f830184611ae4565b92915050565b5f5ffd5b5f5ffd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b611b5a82611b14565b810181811067ffffffffffffffff82111715611b7957611b78611b24565b5b80604052505050565b5f611b8b611873565b9050611b978282611b51565b919050565b5f67ffffffffffffffff821115611bb657611bb5611b24565b5b611bbf82611b14565b9050602081019050919050565b828183375f83830152505050565b5f611bec611be784611b9c565b611b82565b905082815260208101848484011115611c0857611c07611b10565b5b611c13848285611bcc565b509392505050565b5f82601f830112611c2f57611c2e611b0c565b5b8135611c3f848260208601611bda565b91505092915050565b5f5f5f5f5f60a08688031215611c6157611c6061187c565b5b5f86013567ffffffffffffffff811115611c7e57611c7d611880565b5b611c8a88828901611c1b565b9550506020611c9b88828901611a4d565b945050604086013567ffffffffffffffff811115611cbc57611cbb611880565b5b611cc888828901611c1b565b9350506060611cd988828901611a80565b9250506080611cea88828901611a80565b9150509295509295909350565b5f819050919050565b5f611d1a611d15611d1084611a07565b611cf7565b611a07565b9050919050565b5f611d2b82611d00565b9050919050565b5f611d3c82611d21565b9050919050565b611d4c81611d32565b82525050565b5f602082019050611d655f830184611d43565b92915050565b5f60208284031215611d8057611d7f61187c565b5b5f611d8d84828501611a4d565b91505092915050565b5f82825260208201905092915050565b7f6f6e6c792074686520636f6e747261637420697473656c662063616e2063616c5f8201527f6c20746869732066756e6374696f6e0000000000000000000000000000000000602082015250565b5f611e00602f83611d96565b9150611e0b82611da6565b604082019050919050565b5f6020820190508181035f830152611e2d81611df4565b9050919050565b7f7772617070656420746f6b656e20646f65736e277420657869737400000000005f82015250565b5f611e68601b83611d96565b9150611e7382611e34565b602082019050919050565b5f6020820190508181035f830152611e9581611e5c565b9050919050565b611ea581611a61565b82525050565b5f604082019050611ebe5f830185611ae4565b611ecb6020830184611e9c565b9392505050565b5f81519050611ee081611974565b92915050565b5f60208284031215611efb57611efa61187c565b5b5f611f0884828501611ed2565b91505092915050565b5f60ff82169050919050565b611f2681611f11565b8114611f30575f5ffd5b50565b5f81519050611f4181611f1d565b92915050565b5f60208284031215611f5c57611f5b61187c565b5b5f611f6984828501611f33565b91505092915050565b5f606082019050611f855f830186611ae4565b611f926020830185611ae4565b611f9f6040830184611e9c565b949350505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f611fde82611937565b9150611fe983611937565b9250828201905067ffffffffffffffff81111561200957612008611fa7565b5b92915050565b5f81519050919050565b8281835e5f83830152505050565b5f6120318261200f565b61203b8185611d96565b935061204b818560208601612019565b61205481611b14565b840191505092915050565b5f6020820190508181035f8301526120778184612027565b905092915050565b5f6020820190506120925f830184611e9c565b92915050565b5f6120a282611f11565b91506120ad83611f11565b9250828203905060ff8111156120c6576120c5611fa7565b5b92915050565b5f8160011c9050919050565b5f5f8291508390505b6001851115612121578086048111156120fd576120fc611fa7565b5b600185161561210c5780820291505b808102905061211a856120cc565b94506120e1565b94509492505050565b5f8261213957600190506121f4565b81612146575f90506121f4565b816001811461215c576002811461216657612195565b60019150506121f4565b60ff84111561217857612177611fa7565b5b8360020a91508482111561218f5761218e611fa7565b5b506121f4565b5060208310610133831016604e8410600b84101617156121ca5782820a9050838111156121c5576121c4611fa7565b5b6121f4565b6121d784848460016120d8565b925090508184048111156121ee576121ed611fa7565b5b81810290505b9392505050565b5f61220582611a61565b915061221083611a61565b925061223d7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff848461212a565b905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f61227c82611a61565b915061228783611a61565b92508261229757612296612245565b5b828204905092915050565b5f6122ac82611a61565b91506122b783611a61565b92508282026122c581611a61565b915082820484148315176122dc576122db611fa7565b5b5092915050565b7f636f6e76657274656420616d6f756e74206973207a65726f00000000000000005f82015250565b5f612317601883611d96565b9150612322826122e3565b602082019050919050565b5f6020820190508181035f8301526123448161230b565b9050919050565b5f61235d61235884611b9c565b611b82565b90508281526020810184848401111561237957612378611b10565b5b612384848285612019565b509392505050565b5f82601f8301126123a05761239f611b0c565b5b81516123b084826020860161234b565b91505092915050565b5f602082840312156123ce576123cd61187c565b5b5f82015167ffffffffffffffff8111156123eb576123ea611880565b5b6123f78482850161238c565b91505092915050565b5f81905092915050565b5f6124148261200f565b61241e8185612400565b935061242e818560208601612019565b80840191505092915050565b5f612445828561240a565b9150612451828461240a565b91508190509392505050565b61246681611f11565b82525050565b5f6060820190508181035f8301526124848186612027565b905081810360208301526124988185612027565b90506124a7604083018461245d565b949350505050565b5f815190506124bd81611a37565b92915050565b5f602082840312156124d8576124d761187c565b5b5f6124e5848285016124af565b91505092915050565b7f7b224074797065223a20222f6962632e6170706c69636174696f6e732e7472615f8201527f6e736665722e76312e4d73675472616e73666572222c00000000000000000000602082015250565b5f612548603683612400565b9150612553826124ee565b603682019050919050565b7f22736f757263655f706f7274223a20227472616e73666572222c0000000000005f82015250565b5f612592601a83612400565b915061259d8261255e565b601a82019050919050565b7f22736f757263655f6368616e6e656c223a2022000000000000000000000000005f82015250565b5f6125dc601383612400565b91506125e7826125a8565b601382019050919050565b7f222c0000000000000000000000000000000000000000000000000000000000005f82015250565b5f612626600283612400565b9150612631826125f2565b600282019050919050565b7f22746f6b656e223a207b202264656e6f6d223a202200000000000000000000005f82015250565b5f612670601583612400565b915061267b8261263c565b601582019050919050565b7f22616d6f756e74223a20220000000000000000000000000000000000000000005f82015250565b5f6126ba600b83612400565b91506126c582612686565b600b82019050919050565b7f227d2c00000000000000000000000000000000000000000000000000000000005f82015250565b5f612704600383612400565b915061270f826126d0565b600382019050919050565b7f2273656e646572223a20220000000000000000000000000000000000000000005f82015250565b5f61274e600b83612400565b91506127598261271a565b600b82019050919050565b7f227265636569766572223a2022000000000000000000000000000000000000005f82015250565b5f612798600d83612400565b91506127a382612764565b600d82019050919050565b7f2274696d656f75745f686569676874223a207b227265766973696f6e5f6e756d5f8201527f626572223a202230222c227265766973696f6e5f686569676874223a2022302260208201527f7d2c000000000000000000000000000000000000000000000000000000000000604082015250565b5f61282e604283612400565b9150612839826127ae565b604282019050919050565b7f2274696d656f75745f74696d657374616d70223a2022000000000000000000005f82015250565b5f612878601683612400565b915061288382612844565b601682019050919050565b7f226d656d6f223a20227b5c2265766d5c223a207b5c226173796e635f63616c6c5f8201527f6261636b5c223a207b5c2269645c223a20000000000000000000000000000000602082015250565b5f6128e8603183612400565b91506128f38261288e565b603182019050919050565b7f2c5c22636f6e74726163745f616464726573735c223a5c2200000000000000005f82015250565b5f612932601883612400565b915061293d826128fe565b601882019050919050565b7f5c227d7d7d227d000000000000000000000000000000000000000000000000005f82015250565b5f61297c600783612400565b915061298782612948565b600782019050919050565b5f61299c8261253c565b91506129a782612586565b91506129b2826125d0565b91506129be828b61240a565b91506129c98261261a565b91506129d482612664565b91506129e0828a61240a565b91506129eb8261261a565b91506129f6826126ae565b9150612a02828961240a565b9150612a0d826126f8565b9150612a1882612742565b9150612a24828861240a565b9150612a2f8261261a565b9150612a3a8261278c565b9150612a46828761240a565b9150612a518261261a565b9150612a5c82612822565b9150612a678261286c565b9150612a73828661240a565b9150612a7e8261261a565b9150612a89826128dc565b9150612a95828561240a565b9150612aa082612926565b9150612aac828461240a565b9150612ab782612970565b91508190509998505050505050505050565b5f612ad382611a61565b9150612ade83611a61565b9250828201905080821115612af657612af5611fa7565b5b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f612b3382611a61565b91505f8203612b4557612b44611fa7565b5b600182039050919050565b5f604082019050612b635f830185611e9c565b612b706020830184611e9c565b939250505056fea2646970667358221220afc6324e939ed2eb8739e6cb65f004e81c216eaec4d1813ebc44733404dca01c64736f6c634300081c0033",
}

// Erc20WrapperABI is the input ABI used to generate the binding from.
// Deprecated: Use Erc20WrapperMetaData.ABI instead.
var Erc20WrapperABI = Erc20WrapperMetaData.ABI

// Erc20WrapperBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Erc20WrapperMetaData.Bin instead.
var Erc20WrapperBin = Erc20WrapperMetaData.Bin

// DeployErc20Wrapper deploys a new Ethereum contract, binding an instance of Erc20Wrapper to it.
func DeployErc20Wrapper(auth *bind.TransactOpts, backend bind.ContractBackend, erc20Factory common.Address) (common.Address, *types.Transaction, *Erc20Wrapper, error) {
	parsed, err := Erc20WrapperMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Erc20WrapperBin), backend, erc20Factory)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Erc20Wrapper{Erc20WrapperCaller: Erc20WrapperCaller{contract: contract}, Erc20WrapperTransactor: Erc20WrapperTransactor{contract: contract}, Erc20WrapperFilterer: Erc20WrapperFilterer{contract: contract}}, nil
}

// Erc20Wrapper is an auto generated Go binding around an Ethereum contract.
type Erc20Wrapper struct {
	Erc20WrapperCaller     // Read-only binding to the contract
	Erc20WrapperTransactor // Write-only binding to the contract
	Erc20WrapperFilterer   // Log filterer for contract events
}

// Erc20WrapperCaller is an auto generated read-only Go binding around an Ethereum contract.
type Erc20WrapperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20WrapperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Erc20WrapperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20WrapperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Erc20WrapperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20WrapperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Erc20WrapperSession struct {
	Contract     *Erc20Wrapper     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Erc20WrapperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Erc20WrapperCallerSession struct {
	Contract *Erc20WrapperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// Erc20WrapperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Erc20WrapperTransactorSession struct {
	Contract     *Erc20WrapperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// Erc20WrapperRaw is an auto generated low-level Go binding around an Ethereum contract.
type Erc20WrapperRaw struct {
	Contract *Erc20Wrapper // Generic contract binding to access the raw methods on
}

// Erc20WrapperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Erc20WrapperCallerRaw struct {
	Contract *Erc20WrapperCaller // Generic read-only contract binding to access the raw methods on
}

// Erc20WrapperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Erc20WrapperTransactorRaw struct {
	Contract *Erc20WrapperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewErc20Wrapper creates a new instance of Erc20Wrapper, bound to a specific deployed contract.
func NewErc20Wrapper(address common.Address, backend bind.ContractBackend) (*Erc20Wrapper, error) {
	contract, err := bindErc20Wrapper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Erc20Wrapper{Erc20WrapperCaller: Erc20WrapperCaller{contract: contract}, Erc20WrapperTransactor: Erc20WrapperTransactor{contract: contract}, Erc20WrapperFilterer: Erc20WrapperFilterer{contract: contract}}, nil
}

// NewErc20WrapperCaller creates a new read-only instance of Erc20Wrapper, bound to a specific deployed contract.
func NewErc20WrapperCaller(address common.Address, caller bind.ContractCaller) (*Erc20WrapperCaller, error) {
	contract, err := bindErc20Wrapper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Erc20WrapperCaller{contract: contract}, nil
}

// NewErc20WrapperTransactor creates a new write-only instance of Erc20Wrapper, bound to a specific deployed contract.
func NewErc20WrapperTransactor(address common.Address, transactor bind.ContractTransactor) (*Erc20WrapperTransactor, error) {
	contract, err := bindErc20Wrapper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Erc20WrapperTransactor{contract: contract}, nil
}

// NewErc20WrapperFilterer creates a new log filterer instance of Erc20Wrapper, bound to a specific deployed contract.
func NewErc20WrapperFilterer(address common.Address, filterer bind.ContractFilterer) (*Erc20WrapperFilterer, error) {
	contract, err := bindErc20Wrapper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Erc20WrapperFilterer{contract: contract}, nil
}

// bindErc20Wrapper binds a generic wrapper to an already deployed contract.
func bindErc20Wrapper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Erc20WrapperMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc20Wrapper *Erc20WrapperRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc20Wrapper.Contract.Erc20WrapperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc20Wrapper *Erc20WrapperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc20Wrapper.Contract.Erc20WrapperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc20Wrapper *Erc20WrapperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc20Wrapper.Contract.Erc20WrapperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc20Wrapper *Erc20WrapperCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc20Wrapper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc20Wrapper *Erc20WrapperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc20Wrapper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc20Wrapper *Erc20WrapperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc20Wrapper.Contract.contract.Transact(opts, method, params...)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Erc20Wrapper *Erc20WrapperCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Erc20Wrapper.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Erc20Wrapper *Erc20WrapperSession) Factory() (common.Address, error) {
	return _Erc20Wrapper.Contract.Factory(&_Erc20Wrapper.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Erc20Wrapper *Erc20WrapperCallerSession) Factory() (common.Address, error) {
	return _Erc20Wrapper.Contract.Factory(&_Erc20Wrapper.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Erc20Wrapper *Erc20WrapperCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Erc20Wrapper.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Erc20Wrapper *Erc20WrapperSession) Owner() (common.Address, error) {
	return _Erc20Wrapper.Contract.Owner(&_Erc20Wrapper.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Erc20Wrapper *Erc20WrapperCallerSession) Owner() (common.Address, error) {
	return _Erc20Wrapper.Contract.Owner(&_Erc20Wrapper.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Erc20Wrapper *Erc20WrapperCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Erc20Wrapper.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Erc20Wrapper *Erc20WrapperSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Erc20Wrapper.Contract.SupportsInterface(&_Erc20Wrapper.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Erc20Wrapper *Erc20WrapperCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Erc20Wrapper.Contract.SupportsInterface(&_Erc20Wrapper.CallOpts, interfaceId)
}

// WrappedTokens is a free data retrieval call binding the contract method 0xd5c6b504.
//
// Solidity: function wrappedTokens(address ) view returns(address)
func (_Erc20Wrapper *Erc20WrapperCaller) WrappedTokens(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Erc20Wrapper.contract.Call(opts, &out, "wrappedTokens", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WrappedTokens is a free data retrieval call binding the contract method 0xd5c6b504.
//
// Solidity: function wrappedTokens(address ) view returns(address)
func (_Erc20Wrapper *Erc20WrapperSession) WrappedTokens(arg0 common.Address) (common.Address, error) {
	return _Erc20Wrapper.Contract.WrappedTokens(&_Erc20Wrapper.CallOpts, arg0)
}

// WrappedTokens is a free data retrieval call binding the contract method 0xd5c6b504.
//
// Solidity: function wrappedTokens(address ) view returns(address)
func (_Erc20Wrapper *Erc20WrapperCallerSession) WrappedTokens(arg0 common.Address) (common.Address, error) {
	return _Erc20Wrapper.Contract.WrappedTokens(&_Erc20Wrapper.CallOpts, arg0)
}

// IbcAck is a paid mutator transaction binding the contract method 0x0d4f1f9d.
//
// Solidity: function ibc_ack(uint64 callback_id, bool success) returns()
func (_Erc20Wrapper *Erc20WrapperTransactor) IbcAck(opts *bind.TransactOpts, callback_id uint64, success bool) (*types.Transaction, error) {
	return _Erc20Wrapper.contract.Transact(opts, "ibc_ack", callback_id, success)
}

// IbcAck is a paid mutator transaction binding the contract method 0x0d4f1f9d.
//
// Solidity: function ibc_ack(uint64 callback_id, bool success) returns()
func (_Erc20Wrapper *Erc20WrapperSession) IbcAck(callback_id uint64, success bool) (*types.Transaction, error) {
	return _Erc20Wrapper.Contract.IbcAck(&_Erc20Wrapper.TransactOpts, callback_id, success)
}

// IbcAck is a paid mutator transaction binding the contract method 0x0d4f1f9d.
//
// Solidity: function ibc_ack(uint64 callback_id, bool success) returns()
func (_Erc20Wrapper *Erc20WrapperTransactorSession) IbcAck(callback_id uint64, success bool) (*types.Transaction, error) {
	return _Erc20Wrapper.Contract.IbcAck(&_Erc20Wrapper.TransactOpts, callback_id, success)
}

// IbcTimeout is a paid mutator transaction binding the contract method 0x31a503f0.
//
// Solidity: function ibc_timeout(uint64 callback_id) returns()
func (_Erc20Wrapper *Erc20WrapperTransactor) IbcTimeout(opts *bind.TransactOpts, callback_id uint64) (*types.Transaction, error) {
	return _Erc20Wrapper.contract.Transact(opts, "ibc_timeout", callback_id)
}

// IbcTimeout is a paid mutator transaction binding the contract method 0x31a503f0.
//
// Solidity: function ibc_timeout(uint64 callback_id) returns()
func (_Erc20Wrapper *Erc20WrapperSession) IbcTimeout(callback_id uint64) (*types.Transaction, error) {
	return _Erc20Wrapper.Contract.IbcTimeout(&_Erc20Wrapper.TransactOpts, callback_id)
}

// IbcTimeout is a paid mutator transaction binding the contract method 0x31a503f0.
//
// Solidity: function ibc_timeout(uint64 callback_id) returns()
func (_Erc20Wrapper *Erc20WrapperTransactorSession) IbcTimeout(callback_id uint64) (*types.Transaction, error) {
	return _Erc20Wrapper.Contract.IbcTimeout(&_Erc20Wrapper.TransactOpts, callback_id)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Erc20Wrapper *Erc20WrapperTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Erc20Wrapper.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Erc20Wrapper *Erc20WrapperSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Erc20Wrapper.Contract.TransferOwnership(&_Erc20Wrapper.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Erc20Wrapper *Erc20WrapperTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Erc20Wrapper.Contract.TransferOwnership(&_Erc20Wrapper.TransactOpts, newOwner)
}

// Unwrap is a paid mutator transaction binding the contract method 0x8cc7104f.
//
// Solidity: function unwrap(address originToken, address receiver, uint256 wrappedAmt) returns()
func (_Erc20Wrapper *Erc20WrapperTransactor) Unwrap(opts *bind.TransactOpts, originToken common.Address, receiver common.Address, wrappedAmt *big.Int) (*types.Transaction, error) {
	return _Erc20Wrapper.contract.Transact(opts, "unwrap", originToken, receiver, wrappedAmt)
}

// Unwrap is a paid mutator transaction binding the contract method 0x8cc7104f.
//
// Solidity: function unwrap(address originToken, address receiver, uint256 wrappedAmt) returns()
func (_Erc20Wrapper *Erc20WrapperSession) Unwrap(originToken common.Address, receiver common.Address, wrappedAmt *big.Int) (*types.Transaction, error) {
	return _Erc20Wrapper.Contract.Unwrap(&_Erc20Wrapper.TransactOpts, originToken, receiver, wrappedAmt)
}

// Unwrap is a paid mutator transaction binding the contract method 0x8cc7104f.
//
// Solidity: function unwrap(address originToken, address receiver, uint256 wrappedAmt) returns()
func (_Erc20Wrapper *Erc20WrapperTransactorSession) Unwrap(originToken common.Address, receiver common.Address, wrappedAmt *big.Int) (*types.Transaction, error) {
	return _Erc20Wrapper.Contract.Unwrap(&_Erc20Wrapper.TransactOpts, originToken, receiver, wrappedAmt)
}

// Wrap is a paid mutator transaction binding the contract method 0x9a111432.
//
// Solidity: function wrap(string channel, address token, string receiver, uint256 amount, uint256 timeout) returns()
func (_Erc20Wrapper *Erc20WrapperTransactor) Wrap(opts *bind.TransactOpts, channel string, token common.Address, receiver string, amount *big.Int, timeout *big.Int) (*types.Transaction, error) {
	return _Erc20Wrapper.contract.Transact(opts, "wrap", channel, token, receiver, amount, timeout)
}

// Wrap is a paid mutator transaction binding the contract method 0x9a111432.
//
// Solidity: function wrap(string channel, address token, string receiver, uint256 amount, uint256 timeout) returns()
func (_Erc20Wrapper *Erc20WrapperSession) Wrap(channel string, token common.Address, receiver string, amount *big.Int, timeout *big.Int) (*types.Transaction, error) {
	return _Erc20Wrapper.Contract.Wrap(&_Erc20Wrapper.TransactOpts, channel, token, receiver, amount, timeout)
}

// Wrap is a paid mutator transaction binding the contract method 0x9a111432.
//
// Solidity: function wrap(string channel, address token, string receiver, uint256 amount, uint256 timeout) returns()
func (_Erc20Wrapper *Erc20WrapperTransactorSession) Wrap(channel string, token common.Address, receiver string, amount *big.Int, timeout *big.Int) (*types.Transaction, error) {
	return _Erc20Wrapper.Contract.Wrap(&_Erc20Wrapper.TransactOpts, channel, token, receiver, amount, timeout)
}

// Erc20WrapperOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Erc20Wrapper contract.
type Erc20WrapperOwnershipTransferredIterator struct {
	Event *Erc20WrapperOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *Erc20WrapperOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc20WrapperOwnershipTransferred)
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
		it.Event = new(Erc20WrapperOwnershipTransferred)
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
func (it *Erc20WrapperOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc20WrapperOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc20WrapperOwnershipTransferred represents a OwnershipTransferred event raised by the Erc20Wrapper contract.
type Erc20WrapperOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Erc20Wrapper *Erc20WrapperFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*Erc20WrapperOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Erc20Wrapper.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &Erc20WrapperOwnershipTransferredIterator{contract: _Erc20Wrapper.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Erc20Wrapper *Erc20WrapperFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *Erc20WrapperOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Erc20Wrapper.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc20WrapperOwnershipTransferred)
				if err := _Erc20Wrapper.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Erc20Wrapper *Erc20WrapperFilterer) ParseOwnershipTransferred(log types.Log) (*Erc20WrapperOwnershipTransferred, error) {
	event := new(Erc20WrapperOwnershipTransferred)
	if err := _Erc20Wrapper.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
