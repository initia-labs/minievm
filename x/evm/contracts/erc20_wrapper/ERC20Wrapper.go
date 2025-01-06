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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc20Factory\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"StringsInsufficientHexLength\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"contractERC20Factory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"callback_id\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"name\":\"ibc_ack\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"callback_id\",\"type\":\"uint64\"}],\"name\":\"ibc_timeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newFactory\",\"type\":\"address\"}],\"name\":\"setFactory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"originToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wrappedAmt\",\"type\":\"uint256\"}],\"name\":\"unwrap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"channel\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"receiver\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"}],\"name\":\"wrap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wrappedTokens\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040525f8060146101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550348015610037575f80fd5b50604051612f37380380612f378339818101604052810190610059919061013c565b335f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508060015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050610167565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61010b826100e2565b9050919050565b61011b81610101565b8114610125575f80fd5b50565b5f8151905061013681610112565b92915050565b5f60208284031215610151576101506100de565b5b5f61015e84828501610128565b91505092915050565b612dc3806101745f395ff3fe608060405234801561000f575f80fd5b506004361061009c575f3560e01c80638da5cb5b116100645780638da5cb5b146101405780639a1114321461015e578063c45a01551461017a578063d5c6b50414610198578063f2fde38b146101c85761009c565b806301ffc9a7146100a05780630d4f1f9d146100d057806331a503f0146100ec5780635bb47808146101085780638cc7104f14610124575b5f80fd5b6100ba60048036038101906100b59190611a1f565b6101e4565b6040516100c79190611a64565b60405180910390f35b6100ea60048036038101906100e59190611ae4565b61024d565b005b61010660048036038101906101019190611b22565b6102ce565b005b610122600480360381019061011d9190611ba7565b610348565b005b61013e60048036038101906101399190611c05565b610468565b005b6101486106b0565b6040516101559190611c64565b60405180910390f35b61017860048036038101906101739190611db9565b6106d3565b005b610182610aec565b60405161018f9190611ec3565b60405180910390f35b6101b260048036038101906101ad9190611ba7565b610b11565b6040516101bf9190611c64565b60405180910390f35b6101e260048036038101906101dd9190611ba7565b610b41565b005b5f7f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b3073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146102bb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102b290611f5c565b60405180910390fd5b806102ca576102c982610c89565b5b5050565b3073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461033c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161033390611f5c565b60405180910390fd5b61034581610c89565b50565b600173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146103b7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103ae90611fc4565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610425576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161041c9061202c565b60405180910390fd5b8060015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b5f60025f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610536576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161052d90612094565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff166379cc679033846040518363ffffffff1660e01b81526004016105719291906120c1565b6020604051808303815f875af115801561058d573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906105b191906120fc565b505f61062b8360068773ffffffffffffffffffffffffffffffffffffffff1663313ce5676040518163ffffffff1660e01b8152600401602060405180830381865afa158015610602573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610626919061215d565b610fb7565b90508473ffffffffffffffffffffffffffffffffffffffff1663a9059cbb85836040518363ffffffff1660e01b81526004016106689291906120c1565b6020604051808303815f875af1158015610684573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906106a891906120fc565b505050505050565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6106dc84611086565b8373ffffffffffffffffffffffffffffffffffffffff166323b872dd3330856040518463ffffffff1660e01b815260040161071993929190612188565b6020604051808303815f875af1158015610735573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061075991906120fc565b505f6107d3838673ffffffffffffffffffffffffffffffffffffffff1663313ce5676040518163ffffffff1660e01b8152600401602060405180830381865afa1580156107a8573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906107cc919061215d565b6006610fb7565b905060025f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166340c10f1930836040518363ffffffff1660e01b815260040161086c9291906120c1565b5f604051808303815f87803b158015610883575f80fd5b505af1158015610895573d5f803e3d5ffd5b5050505060015f60148282829054906101000a900467ffffffffffffffff166108be91906121ea565b92506101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060405180606001604052803373ffffffffffffffffffffffffffffffffffffffff1681526020018673ffffffffffffffffffffffffffffffffffffffff1681526020018281525060035f8060149054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f820151815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550604082015181600201559050505f610a668760025f8973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff168486896113c1565b905060f173ffffffffffffffffffffffffffffffffffffffff1663d46f64e6826040518263ffffffff1660e01b8152600401610aa29190612275565b6020604051808303815f875af1158015610abe573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610ae291906120fc565b5050505050505050565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6002602052805f5260405f205f915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610b97575f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610bce575f80fd5b8073ffffffffffffffffffffffffffffffffffffffff165f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3805f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b5f60035f8367ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f206040518060600160405290815f82015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600182015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160028201548152505090505f60025f836020015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610e41576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e3890612094565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff166342966c6883604001516040518263ffffffff1660e01b8152600401610e7e9190612295565b5f604051808303815f87803b158015610e95575f80fd5b505af1158015610ea7573d5f803e3d5ffd5b505050505f610f2c83604001516006856020015173ffffffffffffffffffffffffffffffffffffffff1663313ce5676040518163ffffffff1660e01b8152600401602060405180830381865afa158015610f03573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610f27919061215d565b610fb7565b9050826020015173ffffffffffffffffffffffffffffffffffffffff1663a9059cbb845f0151836040518363ffffffff1660e01b8152600401610f709291906120c1565b6020604051808303815f875af1158015610f8c573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610fb091906120fc565b5050505050565b5f8160ff168360ff161115610ff8575f8284610fd391906122ae565b60ff16600a610fe29190612411565b90508085610ff09190612488565b91505061103d565b8160ff168360ff161015611038575f838361101391906122ae565b60ff16600a6110229190612411565b9050808561103091906124b8565b91505061103c565b8390505b5b5f810361107f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161107690612543565b60405180910390fd5b9392505050565b5f73ffffffffffffffffffffffffffffffffffffffff1660025f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16036113be575f60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166306ef1a866040518060400160405280600781526020017f57726170706564000000000000000000000000000000000000000000000000008152508473ffffffffffffffffffffffffffffffffffffffff166306fdde036040518163ffffffff1660e01b81526004015f60405180830381865afa1580156111d1573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f820116820180604052508101906111f991906125cf565b60405160200161120a929190612650565b6040516020818303038152906040526040518060400160405280600181526020017f57000000000000000000000000000000000000000000000000000000000000008152508573ffffffffffffffffffffffffffffffffffffffff166395d89b416040518163ffffffff1660e01b81526004015f60405180830381865afa158015611297573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f820116820180604052508101906112bf91906125cf565b6040516020016112d0929190612650565b60405160208183030381529060405260066040518463ffffffff1660e01b81526004016112ff93929190612682565b6020604051808303815f875af115801561131b573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061133f91906126d9565b90508060025f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505b50565b60608560f173ffffffffffffffffffffffffffffffffffffffff166381cf0f6a876040518263ffffffff1660e01b81526004016113fe9190611c64565b5f604051808303815f875af1158015611419573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f8201168201806040525081019061144191906125cf565b61144a86611533565b60f173ffffffffffffffffffffffffffffffffffffffff16636af32a55306040518263ffffffff1660e01b81526004016114849190611c64565b5f604051808303815f875af115801561149f573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f820116820180604052508101906114c791906125cf565b856114d188611533565b6114f95f60149054906101000a900467ffffffffffffffff1667ffffffffffffffff16611533565b611502306115fd565b604051602001611519989796959493929190612ba8565b604051602081830303815290604052905095945050505050565b60605f60016115418461162a565b0190505f8167ffffffffffffffff81111561155f5761155e611c95565b5b6040519080825280601f01601f1916602001820160405280156115915781602001600182028036833780820191505090505b5090505f82602001820190505b6001156115f2578080600190039150507f3031323334353637383961626364656600000000000000000000000000000000600a86061a8153600a85816115e7576115e661245b565b5b0494505f850361159e575b819350505050919050565b60606116238273ffffffffffffffffffffffffffffffffffffffff16601460ff1661177b565b9050919050565b5f805f90507a184f03e93ff9f4daa797ed6e38ed64bf6a1f0100000000000000008310611686577a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000838161167c5761167b61245b565b5b0492506040810190505b6d04ee2d6d415b85acef810000000083106116c3576d04ee2d6d415b85acef810000000083816116b9576116b861245b565b5b0492506020810190505b662386f26fc1000083106116f257662386f26fc1000083816116e8576116e761245b565b5b0492506010810190505b6305f5e100831061171b576305f5e10083816117115761171061245b565b5b0492506008810190505b61271083106117405761271083816117365761173561245b565b5b0492506004810190505b6064831061176357606483816117595761175861245b565b5b0492506002810190505b600a8310611772576001810190505b80915050919050565b60605f8390505f600284600261179191906124b8565b61179b9190612cdf565b67ffffffffffffffff8111156117b4576117b3611c95565b5b6040519080825280601f01601f1916602001820160405280156117e65781602001600182028036833780820191505090505b5090507f3000000000000000000000000000000000000000000000000000000000000000815f8151811061181d5761181c612d12565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a9053507f7800000000000000000000000000000000000000000000000000000000000000816001815181106118805761187f612d12565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a9053505f60018560026118be91906124b8565b6118c89190612cdf565b90505b6001811115611967577f3031323334353637383961626364656600000000000000000000000000000000600f84166010811061190a57611909612d12565b5b1a60f81b82828151811061192157611920612d12565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a905350600483901c92508061196090612d3f565b90506118cb565b505f82146119ae5784846040517fe22e27eb0000000000000000000000000000000000000000000000000000000081526004016119a5929190612d66565b60405180910390fd5b809250505092915050565b5f604051905090565b5f80fd5b5f80fd5b5f7fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b6119fe816119ca565b8114611a08575f80fd5b50565b5f81359050611a19816119f5565b92915050565b5f60208284031215611a3457611a336119c2565b5b5f611a4184828501611a0b565b91505092915050565b5f8115159050919050565b611a5e81611a4a565b82525050565b5f602082019050611a775f830184611a55565b92915050565b5f67ffffffffffffffff82169050919050565b611a9981611a7d565b8114611aa3575f80fd5b50565b5f81359050611ab481611a90565b92915050565b611ac381611a4a565b8114611acd575f80fd5b50565b5f81359050611ade81611aba565b92915050565b5f8060408385031215611afa57611af96119c2565b5b5f611b0785828601611aa6565b9250506020611b1885828601611ad0565b9150509250929050565b5f60208284031215611b3757611b366119c2565b5b5f611b4484828501611aa6565b91505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f611b7682611b4d565b9050919050565b611b8681611b6c565b8114611b90575f80fd5b50565b5f81359050611ba181611b7d565b92915050565b5f60208284031215611bbc57611bbb6119c2565b5b5f611bc984828501611b93565b91505092915050565b5f819050919050565b611be481611bd2565b8114611bee575f80fd5b50565b5f81359050611bff81611bdb565b92915050565b5f805f60608486031215611c1c57611c1b6119c2565b5b5f611c2986828701611b93565b9350506020611c3a86828701611b93565b9250506040611c4b86828701611bf1565b9150509250925092565b611c5e81611b6c565b82525050565b5f602082019050611c775f830184611c55565b92915050565b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b611ccb82611c85565b810181811067ffffffffffffffff82111715611cea57611ce9611c95565b5b80604052505050565b5f611cfc6119b9565b9050611d088282611cc2565b919050565b5f67ffffffffffffffff821115611d2757611d26611c95565b5b611d3082611c85565b9050602081019050919050565b828183375f83830152505050565b5f611d5d611d5884611d0d565b611cf3565b905082815260208101848484011115611d7957611d78611c81565b5b611d84848285611d3d565b509392505050565b5f82601f830112611da057611d9f611c7d565b5b8135611db0848260208601611d4b565b91505092915050565b5f805f805f60a08688031215611dd257611dd16119c2565b5b5f86013567ffffffffffffffff811115611def57611dee6119c6565b5b611dfb88828901611d8c565b9550506020611e0c88828901611b93565b945050604086013567ffffffffffffffff811115611e2d57611e2c6119c6565b5b611e3988828901611d8c565b9350506060611e4a88828901611bf1565b9250506080611e5b88828901611bf1565b9150509295509295909350565b5f819050919050565b5f611e8b611e86611e8184611b4d565b611e68565b611b4d565b9050919050565b5f611e9c82611e71565b9050919050565b5f611ead82611e92565b9050919050565b611ebd81611ea3565b82525050565b5f602082019050611ed65f830184611eb4565b92915050565b5f82825260208201905092915050565b7f6f6e6c792074686520636f6e747261637420697473656c662063616e2063616c5f8201527f6c20746869732066756e6374696f6e0000000000000000000000000000000000602082015250565b5f611f46602f83611edc565b9150611f5182611eec565b604082019050919050565b5f6020820190508181035f830152611f7381611f3a565b9050919050565b7f45524332303a2063616c6c6572206973206e6f742074686520636861696e00005f82015250565b5f611fae601e83611edc565b9150611fb982611f7a565b602082019050919050565b5f6020820190508181035f830152611fdb81611fa2565b9050919050565b7f696e76616c696420666163746f727920616464726573730000000000000000005f82015250565b5f612016601783611edc565b915061202182611fe2565b602082019050919050565b5f6020820190508181035f8301526120438161200a565b9050919050565b7f7772617070656420746f6b656e20646f65736e277420657869737400000000005f82015250565b5f61207e601b83611edc565b91506120898261204a565b602082019050919050565b5f6020820190508181035f8301526120ab81612072565b9050919050565b6120bb81611bd2565b82525050565b5f6040820190506120d45f830185611c55565b6120e160208301846120b2565b9392505050565b5f815190506120f681611aba565b92915050565b5f60208284031215612111576121106119c2565b5b5f61211e848285016120e8565b91505092915050565b5f60ff82169050919050565b61213c81612127565b8114612146575f80fd5b50565b5f8151905061215781612133565b92915050565b5f60208284031215612172576121716119c2565b5b5f61217f84828501612149565b91505092915050565b5f60608201905061219b5f830186611c55565b6121a86020830185611c55565b6121b560408301846120b2565b949350505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6121f482611a7d565b91506121ff83611a7d565b9250828201905067ffffffffffffffff81111561221f5761221e6121bd565b5b92915050565b5f81519050919050565b8281835e5f83830152505050565b5f61224782612225565b6122518185611edc565b935061226181856020860161222f565b61226a81611c85565b840191505092915050565b5f6020820190508181035f83015261228d818461223d565b905092915050565b5f6020820190506122a85f8301846120b2565b92915050565b5f6122b882612127565b91506122c383612127565b9250828203905060ff8111156122dc576122db6121bd565b5b92915050565b5f8160011c9050919050565b5f808291508390505b600185111561233757808604811115612313576123126121bd565b5b60018516156123225780820291505b8081029050612330856122e2565b94506122f7565b94509492505050565b5f8261234f576001905061240a565b8161235c575f905061240a565b8160018114612372576002811461237c576123ab565b600191505061240a565b60ff84111561238e5761238d6121bd565b5b8360020a9150848211156123a5576123a46121bd565b5b5061240a565b5060208310610133831016604e8410600b84101617156123e05782820a9050838111156123db576123da6121bd565b5b61240a565b6123ed84848460016122ee565b92509050818404811115612404576124036121bd565b5b81810290505b9392505050565b5f61241b82611bd2565b915061242683611bd2565b92506124537fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8484612340565b905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f61249282611bd2565b915061249d83611bd2565b9250826124ad576124ac61245b565b5b828204905092915050565b5f6124c282611bd2565b91506124cd83611bd2565b92508282026124db81611bd2565b915082820484148315176124f2576124f16121bd565b5b5092915050565b7f636f6e76657274656420616d6f756e74206973207a65726f00000000000000005f82015250565b5f61252d601883611edc565b9150612538826124f9565b602082019050919050565b5f6020820190508181035f83015261255a81612521565b9050919050565b5f61257361256e84611d0d565b611cf3565b90508281526020810184848401111561258f5761258e611c81565b5b61259a84828561222f565b509392505050565b5f82601f8301126125b6576125b5611c7d565b5b81516125c6848260208601612561565b91505092915050565b5f602082840312156125e4576125e36119c2565b5b5f82015167ffffffffffffffff811115612601576126006119c6565b5b61260d848285016125a2565b91505092915050565b5f81905092915050565b5f61262a82612225565b6126348185612616565b935061264481856020860161222f565b80840191505092915050565b5f61265b8285612620565b91506126678284612620565b91508190509392505050565b61267c81612127565b82525050565b5f6060820190508181035f83015261269a818661223d565b905081810360208301526126ae818561223d565b90506126bd6040830184612673565b949350505050565b5f815190506126d381611b7d565b92915050565b5f602082840312156126ee576126ed6119c2565b5b5f6126fb848285016126c5565b91505092915050565b7f7b224074797065223a20222f6962632e6170706c69636174696f6e732e7472615f8201527f6e736665722e76312e4d73675472616e73666572222c00000000000000000000602082015250565b5f61275e603683612616565b915061276982612704565b603682019050919050565b7f22736f757263655f706f7274223a20227472616e73666572222c0000000000005f82015250565b5f6127a8601a83612616565b91506127b382612774565b601a82019050919050565b7f22736f757263655f6368616e6e656c223a2022000000000000000000000000005f82015250565b5f6127f2601383612616565b91506127fd826127be565b601382019050919050565b7f222c0000000000000000000000000000000000000000000000000000000000005f82015250565b5f61283c600283612616565b915061284782612808565b600282019050919050565b7f22746f6b656e223a207b202264656e6f6d223a202200000000000000000000005f82015250565b5f612886601583612616565b915061289182612852565b601582019050919050565b7f22616d6f756e74223a20220000000000000000000000000000000000000000005f82015250565b5f6128d0600b83612616565b91506128db8261289c565b600b82019050919050565b7f227d2c00000000000000000000000000000000000000000000000000000000005f82015250565b5f61291a600383612616565b9150612925826128e6565b600382019050919050565b7f2273656e646572223a20220000000000000000000000000000000000000000005f82015250565b5f612964600b83612616565b915061296f82612930565b600b82019050919050565b7f227265636569766572223a2022000000000000000000000000000000000000005f82015250565b5f6129ae600d83612616565b91506129b98261297a565b600d82019050919050565b7f2274696d656f75745f686569676874223a207b227265766973696f6e5f6e756d5f8201527f626572223a202230222c227265766973696f6e5f686569676874223a2022302260208201527f7d2c000000000000000000000000000000000000000000000000000000000000604082015250565b5f612a44604283612616565b9150612a4f826129c4565b604282019050919050565b7f2274696d656f75745f74696d657374616d70223a2022000000000000000000005f82015250565b5f612a8e601683612616565b9150612a9982612a5a565b601682019050919050565b7f226d656d6f223a20227b5c2265766d5c223a207b5c226173796e635f63616c6c5f8201527f6261636b5c223a207b5c2269645c223a20000000000000000000000000000000602082015250565b5f612afe603183612616565b9150612b0982612aa4565b603182019050919050565b7f2c5c22636f6e74726163745f616464726573735c223a5c2200000000000000005f82015250565b5f612b48601883612616565b9150612b5382612b14565b601882019050919050565b7f5c227d7d7d227d000000000000000000000000000000000000000000000000005f82015250565b5f612b92600783612616565b9150612b9d82612b5e565b600782019050919050565b5f612bb282612752565b9150612bbd8261279c565b9150612bc8826127e6565b9150612bd4828b612620565b9150612bdf82612830565b9150612bea8261287a565b9150612bf6828a612620565b9150612c0182612830565b9150612c0c826128c4565b9150612c188289612620565b9150612c238261290e565b9150612c2e82612958565b9150612c3a8288612620565b9150612c4582612830565b9150612c50826129a2565b9150612c5c8287612620565b9150612c6782612830565b9150612c7282612a38565b9150612c7d82612a82565b9150612c898286612620565b9150612c9482612830565b9150612c9f82612af2565b9150612cab8285612620565b9150612cb682612b3c565b9150612cc28284612620565b9150612ccd82612b86565b91508190509998505050505050505050565b5f612ce982611bd2565b9150612cf483611bd2565b9250828201905080821115612d0c57612d0b6121bd565b5b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f612d4982611bd2565b91505f8203612d5b57612d5a6121bd565b5b600182039050919050565b5f604082019050612d795f8301856120b2565b612d8660208301846120b2565b939250505056fea2646970667358221220c92aa6411413595f1f3adde815cbd463b8b06039ef1ceb6acaaca4dc03ebf49264736f6c63430008190033",
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

// SetFactory is a paid mutator transaction binding the contract method 0x5bb47808.
//
// Solidity: function setFactory(address newFactory) returns()
func (_Erc20Wrapper *Erc20WrapperTransactor) SetFactory(opts *bind.TransactOpts, newFactory common.Address) (*types.Transaction, error) {
	return _Erc20Wrapper.contract.Transact(opts, "setFactory", newFactory)
}

// SetFactory is a paid mutator transaction binding the contract method 0x5bb47808.
//
// Solidity: function setFactory(address newFactory) returns()
func (_Erc20Wrapper *Erc20WrapperSession) SetFactory(newFactory common.Address) (*types.Transaction, error) {
	return _Erc20Wrapper.Contract.SetFactory(&_Erc20Wrapper.TransactOpts, newFactory)
}

// SetFactory is a paid mutator transaction binding the contract method 0x5bb47808.
//
// Solidity: function setFactory(address newFactory) returns()
func (_Erc20Wrapper *Erc20WrapperTransactorSession) SetFactory(newFactory common.Address) (*types.Transaction, error) {
	return _Erc20Wrapper.Contract.SetFactory(&_Erc20Wrapper.TransactOpts, newFactory)
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
