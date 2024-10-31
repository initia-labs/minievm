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
	Bin: "0x60a06040525f8060146101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550348015610037575f80fd5b50604051612d24380380612d2483398181016040528101906100599190610130565b335f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff16815250505061015b565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6100ff826100d6565b9050919050565b61010f816100f5565b8114610119575f80fd5b50565b5f8151905061012a81610106565b92915050565b5f60208284031215610145576101446100d2565b5b5f6101528482850161011c565b91505092915050565b608051612baa61017a5f395f81816109a70152610fd00152612baa5ff3fe608060405234801561000f575f80fd5b5060043610610091575f3560e01c80638da5cb5b116100645780638da5cb5b146101195780639a11143214610137578063c45a015514610153578063d5c6b50414610171578063f2fde38b146101a157610091565b806301ffc9a7146100955780630d4f1f9d146100c557806331a503f0146100e15780638cc7104f146100fd575b5f80fd5b6100af60048036038101906100aa91906118d6565b6101bd565b6040516100bc919061191b565b60405180910390f35b6100df60048036038101906100da919061199b565b610226565b005b6100fb60048036038101906100f691906119d9565b6102a7565b005b61011760048036038101906101129190611a91565b610321565b005b610121610569565b60405161012e9190611af0565b60405180910390f35b610151600480360381019061014c9190611c45565b61058c565b005b61015b6109a5565b6040516101689190611d4f565b60405180910390f35b61018b60048036038101906101869190611d68565b6109c9565b6040516101989190611af0565b60405180910390f35b6101bb60048036038101906101b69190611d68565b6109f9565b005b5f7f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b3073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610294576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161028b90611e13565b60405180910390fd5b806102a3576102a282610b41565b5b5050565b3073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610315576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161030c90611e13565b60405180910390fd5b61031e81610b41565b50565b5f60015f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036103ef576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103e690611e7b565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff166379cc679033846040518363ffffffff1660e01b815260040161042a929190611ea8565b6020604051808303815f875af1158015610446573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061046a9190611ee3565b505f6104e48360068773ffffffffffffffffffffffffffffffffffffffff1663313ce5676040518163ffffffff1660e01b8152600401602060405180830381865afa1580156104bb573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906104df9190611f44565b610e6f565b90508473ffffffffffffffffffffffffffffffffffffffff1663a9059cbb85836040518363ffffffff1660e01b8152600401610521929190611ea8565b6020604051808303815f875af115801561053d573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906105619190611ee3565b505050505050565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b61059584610f3e565b8373ffffffffffffffffffffffffffffffffffffffff166323b872dd3330856040518463ffffffff1660e01b81526004016105d293929190611f6f565b6020604051808303815f875af11580156105ee573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906106129190611ee3565b505f61068c838673ffffffffffffffffffffffffffffffffffffffff1663313ce5676040518163ffffffff1660e01b8152600401602060405180830381865afa158015610661573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906106859190611f44565b6006610e6f565b905060015f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166340c10f1930836040518363ffffffff1660e01b8152600401610725929190611ea8565b5f604051808303815f87803b15801561073c575f80fd5b505af115801561074e573d5f803e3d5ffd5b5050505060015f60148282829054906101000a900467ffffffffffffffff166107779190611fd1565b92506101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060405180606001604052803373ffffffffffffffffffffffffffffffffffffffff1681526020018673ffffffffffffffffffffffffffffffffffffffff1681526020018281525060025f8060149054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f820151815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550604082015181600201559050505f61091f8760015f8973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16848689611278565b905060f173ffffffffffffffffffffffffffffffffffffffff1663d46f64e6826040518263ffffffff1660e01b815260040161095b919061205c565b6020604051808303815f875af1158015610977573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061099b9190611ee3565b5050505050505050565b7f000000000000000000000000000000000000000000000000000000000000000081565b6001602052805f5260405f205f915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610a4f575f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610a86575f80fd5b8073ffffffffffffffffffffffffffffffffffffffff165f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3805f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b5f60025f8367ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f206040518060600160405290815f82015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600182015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160028201548152505090505f60015f836020015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610cf9576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610cf090611e7b565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff166342966c6883604001516040518263ffffffff1660e01b8152600401610d36919061207c565b5f604051808303815f87803b158015610d4d575f80fd5b505af1158015610d5f573d5f803e3d5ffd5b505050505f610de483604001516006856020015173ffffffffffffffffffffffffffffffffffffffff1663313ce5676040518163ffffffff1660e01b8152600401602060405180830381865afa158015610dbb573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610ddf9190611f44565b610e6f565b9050826020015173ffffffffffffffffffffffffffffffffffffffff1663a9059cbb845f0151836040518363ffffffff1660e01b8152600401610e28929190611ea8565b6020604051808303815f875af1158015610e44573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610e689190611ee3565b5050505050565b5f8160ff168360ff161115610eb0575f8284610e8b9190612095565b60ff16600a610e9a91906121f8565b90508085610ea8919061226f565b915050610ef5565b8160ff168360ff161015610ef0575f8383610ecb9190612095565b60ff16600a610eda91906121f8565b90508085610ee8919061229f565b915050610ef4565b8390505b5b5f8103610f37576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f2e9061232a565b60405180910390fd5b9392505050565b5f73ffffffffffffffffffffffffffffffffffffffff1660015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603611275575f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166306ef1a866040518060400160405280600781526020017f57726170706564000000000000000000000000000000000000000000000000008152508473ffffffffffffffffffffffffffffffffffffffff166306fdde036040518163ffffffff1660e01b81526004015f60405180830381865afa158015611088573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f820116820180604052508101906110b091906123b6565b6040516020016110c1929190612437565b6040516020818303038152906040526040518060400160405280600181526020017f57000000000000000000000000000000000000000000000000000000000000008152508573ffffffffffffffffffffffffffffffffffffffff166395d89b416040518163ffffffff1660e01b81526004015f60405180830381865afa15801561114e573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f8201168201806040525081019061117691906123b6565b604051602001611187929190612437565b60405160208183030381529060405260066040518463ffffffff1660e01b81526004016111b693929190612469565b6020604051808303815f875af11580156111d2573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906111f691906124c0565b90508060015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505b50565b60608560f173ffffffffffffffffffffffffffffffffffffffff166381cf0f6a876040518263ffffffff1660e01b81526004016112b59190611af0565b5f604051808303815f875af11580156112d0573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f820116820180604052508101906112f891906123b6565b611301866113ea565b60f173ffffffffffffffffffffffffffffffffffffffff16636af32a55306040518263ffffffff1660e01b815260040161133b9190611af0565b5f604051808303815f875af1158015611356573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f8201168201806040525081019061137e91906123b6565b85611388886113ea565b6113b05f60149054906101000a900467ffffffffffffffff1667ffffffffffffffff166113ea565b6113b9306114b4565b6040516020016113d098979695949392919061298f565b604051602081830303815290604052905095945050505050565b60605f60016113f8846114e1565b0190505f8167ffffffffffffffff81111561141657611415611b21565b5b6040519080825280601f01601f1916602001820160405280156114485781602001600182028036833780820191505090505b5090505f82602001820190505b6001156114a9578080600190039150507f3031323334353637383961626364656600000000000000000000000000000000600a86061a8153600a858161149e5761149d612242565b5b0494505f8503611455575b819350505050919050565b60606114da8273ffffffffffffffffffffffffffffffffffffffff16601460ff16611632565b9050919050565b5f805f90507a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000831061153d577a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000838161153357611532612242565b5b0492506040810190505b6d04ee2d6d415b85acef8100000000831061157a576d04ee2d6d415b85acef810000000083816115705761156f612242565b5b0492506020810190505b662386f26fc1000083106115a957662386f26fc10000838161159f5761159e612242565b5b0492506010810190505b6305f5e10083106115d2576305f5e10083816115c8576115c7612242565b5b0492506008810190505b61271083106115f75761271083816115ed576115ec612242565b5b0492506004810190505b6064831061161a57606483816116105761160f612242565b5b0492506002810190505b600a8310611629576001810190505b80915050919050565b60605f8390505f6002846002611648919061229f565b6116529190612ac6565b67ffffffffffffffff81111561166b5761166a611b21565b5b6040519080825280601f01601f19166020018201604052801561169d5781602001600182028036833780820191505090505b5090507f3000000000000000000000000000000000000000000000000000000000000000815f815181106116d4576116d3612af9565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a9053507f78000000000000000000000000000000000000000000000000000000000000008160018151811061173757611736612af9565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a9053505f6001856002611775919061229f565b61177f9190612ac6565b90505b600181111561181e577f3031323334353637383961626364656600000000000000000000000000000000600f8416601081106117c1576117c0612af9565b5b1a60f81b8282815181106117d8576117d7612af9565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a905350600483901c92508061181790612b26565b9050611782565b505f82146118655784846040517fe22e27eb00000000000000000000000000000000000000000000000000000000815260040161185c929190612b4d565b60405180910390fd5b809250505092915050565b5f604051905090565b5f80fd5b5f80fd5b5f7fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b6118b581611881565b81146118bf575f80fd5b50565b5f813590506118d0816118ac565b92915050565b5f602082840312156118eb576118ea611879565b5b5f6118f8848285016118c2565b91505092915050565b5f8115159050919050565b61191581611901565b82525050565b5f60208201905061192e5f83018461190c565b92915050565b5f67ffffffffffffffff82169050919050565b61195081611934565b811461195a575f80fd5b50565b5f8135905061196b81611947565b92915050565b61197a81611901565b8114611984575f80fd5b50565b5f8135905061199581611971565b92915050565b5f80604083850312156119b1576119b0611879565b5b5f6119be8582860161195d565b92505060206119cf85828601611987565b9150509250929050565b5f602082840312156119ee576119ed611879565b5b5f6119fb8482850161195d565b91505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f611a2d82611a04565b9050919050565b611a3d81611a23565b8114611a47575f80fd5b50565b5f81359050611a5881611a34565b92915050565b5f819050919050565b611a7081611a5e565b8114611a7a575f80fd5b50565b5f81359050611a8b81611a67565b92915050565b5f805f60608486031215611aa857611aa7611879565b5b5f611ab586828701611a4a565b9350506020611ac686828701611a4a565b9250506040611ad786828701611a7d565b9150509250925092565b611aea81611a23565b82525050565b5f602082019050611b035f830184611ae1565b92915050565b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b611b5782611b11565b810181811067ffffffffffffffff82111715611b7657611b75611b21565b5b80604052505050565b5f611b88611870565b9050611b948282611b4e565b919050565b5f67ffffffffffffffff821115611bb357611bb2611b21565b5b611bbc82611b11565b9050602081019050919050565b828183375f83830152505050565b5f611be9611be484611b99565b611b7f565b905082815260208101848484011115611c0557611c04611b0d565b5b611c10848285611bc9565b509392505050565b5f82601f830112611c2c57611c2b611b09565b5b8135611c3c848260208601611bd7565b91505092915050565b5f805f805f60a08688031215611c5e57611c5d611879565b5b5f86013567ffffffffffffffff811115611c7b57611c7a61187d565b5b611c8788828901611c18565b9550506020611c9888828901611a4a565b945050604086013567ffffffffffffffff811115611cb957611cb861187d565b5b611cc588828901611c18565b9350506060611cd688828901611a7d565b9250506080611ce788828901611a7d565b9150509295509295909350565b5f819050919050565b5f611d17611d12611d0d84611a04565b611cf4565b611a04565b9050919050565b5f611d2882611cfd565b9050919050565b5f611d3982611d1e565b9050919050565b611d4981611d2f565b82525050565b5f602082019050611d625f830184611d40565b92915050565b5f60208284031215611d7d57611d7c611879565b5b5f611d8a84828501611a4a565b91505092915050565b5f82825260208201905092915050565b7f6f6e6c792074686520636f6e747261637420697473656c662063616e2063616c5f8201527f6c20746869732066756e6374696f6e0000000000000000000000000000000000602082015250565b5f611dfd602f83611d93565b9150611e0882611da3565b604082019050919050565b5f6020820190508181035f830152611e2a81611df1565b9050919050565b7f7772617070656420746f6b656e20646f65736e277420657869737400000000005f82015250565b5f611e65601b83611d93565b9150611e7082611e31565b602082019050919050565b5f6020820190508181035f830152611e9281611e59565b9050919050565b611ea281611a5e565b82525050565b5f604082019050611ebb5f830185611ae1565b611ec86020830184611e99565b9392505050565b5f81519050611edd81611971565b92915050565b5f60208284031215611ef857611ef7611879565b5b5f611f0584828501611ecf565b91505092915050565b5f60ff82169050919050565b611f2381611f0e565b8114611f2d575f80fd5b50565b5f81519050611f3e81611f1a565b92915050565b5f60208284031215611f5957611f58611879565b5b5f611f6684828501611f30565b91505092915050565b5f606082019050611f825f830186611ae1565b611f8f6020830185611ae1565b611f9c6040830184611e99565b949350505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f611fdb82611934565b9150611fe683611934565b9250828201905067ffffffffffffffff81111561200657612005611fa4565b5b92915050565b5f81519050919050565b8281835e5f83830152505050565b5f61202e8261200c565b6120388185611d93565b9350612048818560208601612016565b61205181611b11565b840191505092915050565b5f6020820190508181035f8301526120748184612024565b905092915050565b5f60208201905061208f5f830184611e99565b92915050565b5f61209f82611f0e565b91506120aa83611f0e565b9250828203905060ff8111156120c3576120c2611fa4565b5b92915050565b5f8160011c9050919050565b5f808291508390505b600185111561211e578086048111156120fa576120f9611fa4565b5b60018516156121095780820291505b8081029050612117856120c9565b94506120de565b94509492505050565b5f8261213657600190506121f1565b81612143575f90506121f1565b8160018114612159576002811461216357612192565b60019150506121f1565b60ff84111561217557612174611fa4565b5b8360020a91508482111561218c5761218b611fa4565b5b506121f1565b5060208310610133831016604e8410600b84101617156121c75782820a9050838111156121c2576121c1611fa4565b5b6121f1565b6121d484848460016120d5565b925090508184048111156121eb576121ea611fa4565b5b81810290505b9392505050565b5f61220282611a5e565b915061220d83611a5e565b925061223a7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8484612127565b905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f61227982611a5e565b915061228483611a5e565b92508261229457612293612242565b5b828204905092915050565b5f6122a982611a5e565b91506122b483611a5e565b92508282026122c281611a5e565b915082820484148315176122d9576122d8611fa4565b5b5092915050565b7f636f6e76657274656420616d6f756e74206973207a65726f00000000000000005f82015250565b5f612314601883611d93565b915061231f826122e0565b602082019050919050565b5f6020820190508181035f83015261234181612308565b9050919050565b5f61235a61235584611b99565b611b7f565b90508281526020810184848401111561237657612375611b0d565b5b612381848285612016565b509392505050565b5f82601f83011261239d5761239c611b09565b5b81516123ad848260208601612348565b91505092915050565b5f602082840312156123cb576123ca611879565b5b5f82015167ffffffffffffffff8111156123e8576123e761187d565b5b6123f484828501612389565b91505092915050565b5f81905092915050565b5f6124118261200c565b61241b81856123fd565b935061242b818560208601612016565b80840191505092915050565b5f6124428285612407565b915061244e8284612407565b91508190509392505050565b61246381611f0e565b82525050565b5f6060820190508181035f8301526124818186612024565b905081810360208301526124958185612024565b90506124a4604083018461245a565b949350505050565b5f815190506124ba81611a34565b92915050565b5f602082840312156124d5576124d4611879565b5b5f6124e2848285016124ac565b91505092915050565b7f7b224074797065223a20222f6962632e6170706c69636174696f6e732e7472615f8201527f6e736665722e76312e4d73675472616e73666572222c00000000000000000000602082015250565b5f6125456036836123fd565b9150612550826124eb565b603682019050919050565b7f22736f757263655f706f7274223a20227472616e73666572222c0000000000005f82015250565b5f61258f601a836123fd565b915061259a8261255b565b601a82019050919050565b7f22736f757263655f6368616e6e656c223a2022000000000000000000000000005f82015250565b5f6125d96013836123fd565b91506125e4826125a5565b601382019050919050565b7f222c0000000000000000000000000000000000000000000000000000000000005f82015250565b5f6126236002836123fd565b915061262e826125ef565b600282019050919050565b7f22746f6b656e223a207b202264656e6f6d223a202200000000000000000000005f82015250565b5f61266d6015836123fd565b915061267882612639565b601582019050919050565b7f22616d6f756e74223a20220000000000000000000000000000000000000000005f82015250565b5f6126b7600b836123fd565b91506126c282612683565b600b82019050919050565b7f227d2c00000000000000000000000000000000000000000000000000000000005f82015250565b5f6127016003836123fd565b915061270c826126cd565b600382019050919050565b7f2273656e646572223a20220000000000000000000000000000000000000000005f82015250565b5f61274b600b836123fd565b915061275682612717565b600b82019050919050565b7f227265636569766572223a2022000000000000000000000000000000000000005f82015250565b5f612795600d836123fd565b91506127a082612761565b600d82019050919050565b7f2274696d656f75745f686569676874223a207b227265766973696f6e5f6e756d5f8201527f626572223a202230222c227265766973696f6e5f686569676874223a2022302260208201527f7d2c000000000000000000000000000000000000000000000000000000000000604082015250565b5f61282b6042836123fd565b9150612836826127ab565b604282019050919050565b7f2274696d656f75745f74696d657374616d70223a2022000000000000000000005f82015250565b5f6128756016836123fd565b915061288082612841565b601682019050919050565b7f226d656d6f223a20227b5c2265766d5c223a207b5c226173796e635f63616c6c5f8201527f6261636b5c223a207b5c2269645c223a20000000000000000000000000000000602082015250565b5f6128e56031836123fd565b91506128f08261288b565b603182019050919050565b7f2c5c22636f6e74726163745f616464726573735c223a5c2200000000000000005f82015250565b5f61292f6018836123fd565b915061293a826128fb565b601882019050919050565b7f5c227d7d7d227d000000000000000000000000000000000000000000000000005f82015250565b5f6129796007836123fd565b915061298482612945565b600782019050919050565b5f61299982612539565b91506129a482612583565b91506129af826125cd565b91506129bb828b612407565b91506129c682612617565b91506129d182612661565b91506129dd828a612407565b91506129e882612617565b91506129f3826126ab565b91506129ff8289612407565b9150612a0a826126f5565b9150612a158261273f565b9150612a218288612407565b9150612a2c82612617565b9150612a3782612789565b9150612a438287612407565b9150612a4e82612617565b9150612a598261281f565b9150612a6482612869565b9150612a708286612407565b9150612a7b82612617565b9150612a86826128d9565b9150612a928285612407565b9150612a9d82612923565b9150612aa98284612407565b9150612ab48261296d565b91508190509998505050505050505050565b5f612ad082611a5e565b9150612adb83611a5e565b9250828201905080821115612af357612af2611fa4565b5b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f612b3082611a5e565b91505f8203612b4257612b41611fa4565b5b600182039050919050565b5f604082019050612b605f830185611e99565b612b6d6020830184611e99565b939250505056fea2646970667358221220ccd74c9a245399d217d5469cc62fd2be22f8c19446bbf507e58a72f47cb20cbc64736f6c63430008190033",
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
