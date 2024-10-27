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
	Bin: "0x60a06040525f8060146101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550348015610037575f80fd5b50604051612add380380612add83398181016040528101906100599190610130565b335f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff16815250505061015b565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6100ff826100d6565b9050919050565b61010f816100f5565b8114610119575f80fd5b50565b5f8151905061012a81610106565b92915050565b5f60208284031215610145576101446100d2565b5b5f6101528482850161011c565b91505092915050565b60805161296361017a5f395f81816109a70152610da201526129635ff3fe608060405234801561000f575f80fd5b5060043610610091575f3560e01c80638da5cb5b116100645780638da5cb5b146101195780639a11143214610137578063c45a015514610153578063d5c6b50414610171578063f2fde38b146101a157610091565b806301ffc9a7146100955780630d4f1f9d146100c557806331a503f0146100e15780638cc7104f146100fd575b5f80fd5b6100af60048036038101906100aa91906116a8565b6101bd565b6040516100bc91906116ed565b60405180910390f35b6100df60048036038101906100da919061176d565b610226565b005b6100fb60048036038101906100f691906117ab565b6102a7565b005b61011760048036038101906101129190611863565b610321565b005b610121610569565b60405161012e91906118c2565b60405180910390f35b610151600480360381019061014c9190611a17565b61058c565b005b61015b6109a5565b6040516101689190611b21565b60405180910390f35b61018b60048036038101906101869190611b3a565b6109c9565b60405161019891906118c2565b60405180910390f35b6101bb60048036038101906101b69190611b3a565b6109f9565b005b5f7f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b3073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610294576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161028b90611be5565b60405180910390fd5b806102a3576102a282610b41565b5b5050565b3073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610315576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161030c90611be5565b60405180910390fd5b61031e81610b41565b50565b5f60015f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036103ef576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103e690611c4d565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff166379cc679033846040518363ffffffff1660e01b815260040161042a929190611c7a565b6020604051808303815f875af1158015610446573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061046a9190611cb5565b505f6104e48360068773ffffffffffffffffffffffffffffffffffffffff1663313ce5676040518163ffffffff1660e01b8152600401602060405180830381865afa1580156104bb573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906104df9190611d16565b610c41565b90508473ffffffffffffffffffffffffffffffffffffffff1663a9059cbb85836040518363ffffffff1660e01b8152600401610521929190611c7a565b6020604051808303815f875af115801561053d573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906105619190611cb5565b505050505050565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b61059584610d10565b8373ffffffffffffffffffffffffffffffffffffffff166323b872dd3330856040518463ffffffff1660e01b81526004016105d293929190611d41565b6020604051808303815f875af11580156105ee573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906106129190611cb5565b505f61068c838673ffffffffffffffffffffffffffffffffffffffff1663313ce5676040518163ffffffff1660e01b8152600401602060405180830381865afa158015610661573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906106859190611d16565b6006610c41565b905060015f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166340c10f1930836040518363ffffffff1660e01b8152600401610725929190611c7a565b5f604051808303815f87803b15801561073c575f80fd5b505af115801561074e573d5f803e3d5ffd5b5050505060015f60148282829054906101000a900467ffffffffffffffff166107779190611da3565b92506101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060405180606001604052803373ffffffffffffffffffffffffffffffffffffffff1681526020018673ffffffffffffffffffffffffffffffffffffffff1681526020018281525060025f8060149054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f820151815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550604082015181600201559050505f61091f8760015f8973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1684868961104a565b905060f173ffffffffffffffffffffffffffffffffffffffff1663d46f64e6826040518263ffffffff1660e01b815260040161095b9190611e2e565b6020604051808303815f875af1158015610977573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061099b9190611cb5565b5050505050505050565b7f000000000000000000000000000000000000000000000000000000000000000081565b6001602052805f5260405f205f915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610a4f575f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610a86575f80fd5b8073ffffffffffffffffffffffffffffffffffffffff165f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3805f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b5f60025f8367ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f206040518060600160405290815f82015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600182015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820154815250509050610c3d8160200151825f01518360400151610321565b5050565b5f8160ff168360ff161115610c82575f8284610c5d9190611e4e565b60ff16600a610c6c9190611fb1565b90508085610c7a9190612028565b915050610cc7565b8160ff168360ff161015610cc2575f8383610c9d9190611e4e565b60ff16600a610cac9190611fb1565b90508085610cba9190612058565b915050610cc6565b8390505b5b5f8103610d09576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d00906120e3565b60405180910390fd5b9392505050565b5f73ffffffffffffffffffffffffffffffffffffffff1660015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603611047575f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166306ef1a866040518060400160405280600781526020017f57726170706564000000000000000000000000000000000000000000000000008152508473ffffffffffffffffffffffffffffffffffffffff166306fdde036040518163ffffffff1660e01b81526004015f60405180830381865afa158015610e5a573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f82011682018060405250810190610e82919061216f565b604051602001610e939291906121f0565b6040516020818303038152906040526040518060400160405280600181526020017f57000000000000000000000000000000000000000000000000000000000000008152508573ffffffffffffffffffffffffffffffffffffffff166395d89b416040518163ffffffff1660e01b81526004015f60405180830381865afa158015610f20573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f82011682018060405250810190610f48919061216f565b604051602001610f599291906121f0565b60405160208183030381529060405260066040518463ffffffff1660e01b8152600401610f8893929190612222565b6020604051808303815f875af1158015610fa4573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610fc89190612279565b90508060015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505b50565b60608560f173ffffffffffffffffffffffffffffffffffffffff166381cf0f6a876040518263ffffffff1660e01b815260040161108791906118c2565b5f604051808303815f875af11580156110a2573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f820116820180604052508101906110ca919061216f565b6110d3866111bc565b60f173ffffffffffffffffffffffffffffffffffffffff16636af32a55306040518263ffffffff1660e01b815260040161110d91906118c2565b5f604051808303815f875af1158015611128573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f82011682018060405250810190611150919061216f565b8561115a886111bc565b6111825f60149054906101000a900467ffffffffffffffff1667ffffffffffffffff166111bc565b61118b30611286565b6040516020016111a2989796959493929190612748565b604051602081830303815290604052905095945050505050565b60605f60016111ca846112b3565b0190505f8167ffffffffffffffff8111156111e8576111e76118f3565b5b6040519080825280601f01601f19166020018201604052801561121a5781602001600182028036833780820191505090505b5090505f82602001820190505b60011561127b578080600190039150507f3031323334353637383961626364656600000000000000000000000000000000600a86061a8153600a85816112705761126f611ffb565b5b0494505f8503611227575b819350505050919050565b60606112ac8273ffffffffffffffffffffffffffffffffffffffff16601460ff16611404565b9050919050565b5f805f90507a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000831061130f577a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000838161130557611304611ffb565b5b0492506040810190505b6d04ee2d6d415b85acef8100000000831061134c576d04ee2d6d415b85acef8100000000838161134257611341611ffb565b5b0492506020810190505b662386f26fc10000831061137b57662386f26fc10000838161137157611370611ffb565b5b0492506010810190505b6305f5e10083106113a4576305f5e100838161139a57611399611ffb565b5b0492506008810190505b61271083106113c95761271083816113bf576113be611ffb565b5b0492506004810190505b606483106113ec57606483816113e2576113e1611ffb565b5b0492506002810190505b600a83106113fb576001810190505b80915050919050565b60605f8390505f600284600261141a9190612058565b611424919061287f565b67ffffffffffffffff81111561143d5761143c6118f3565b5b6040519080825280601f01601f19166020018201604052801561146f5781602001600182028036833780820191505090505b5090507f3000000000000000000000000000000000000000000000000000000000000000815f815181106114a6576114a56128b2565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a9053507f780000000000000000000000000000000000000000000000000000000000000081600181518110611509576115086128b2565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a9053505f60018560026115479190612058565b611551919061287f565b90505b60018111156115f0577f3031323334353637383961626364656600000000000000000000000000000000600f841660108110611593576115926128b2565b5b1a60f81b8282815181106115aa576115a96128b2565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a905350600483901c9250806115e9906128df565b9050611554565b505f82146116375784846040517fe22e27eb00000000000000000000000000000000000000000000000000000000815260040161162e929190612906565b60405180910390fd5b809250505092915050565b5f604051905090565b5f80fd5b5f80fd5b5f7fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b61168781611653565b8114611691575f80fd5b50565b5f813590506116a28161167e565b92915050565b5f602082840312156116bd576116bc61164b565b5b5f6116ca84828501611694565b91505092915050565b5f8115159050919050565b6116e7816116d3565b82525050565b5f6020820190506117005f8301846116de565b92915050565b5f67ffffffffffffffff82169050919050565b61172281611706565b811461172c575f80fd5b50565b5f8135905061173d81611719565b92915050565b61174c816116d3565b8114611756575f80fd5b50565b5f8135905061176781611743565b92915050565b5f80604083850312156117835761178261164b565b5b5f6117908582860161172f565b92505060206117a185828601611759565b9150509250929050565b5f602082840312156117c0576117bf61164b565b5b5f6117cd8482850161172f565b91505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6117ff826117d6565b9050919050565b61180f816117f5565b8114611819575f80fd5b50565b5f8135905061182a81611806565b92915050565b5f819050919050565b61184281611830565b811461184c575f80fd5b50565b5f8135905061185d81611839565b92915050565b5f805f6060848603121561187a5761187961164b565b5b5f6118878682870161181c565b93505060206118988682870161181c565b92505060406118a98682870161184f565b9150509250925092565b6118bc816117f5565b82525050565b5f6020820190506118d55f8301846118b3565b92915050565b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b611929826118e3565b810181811067ffffffffffffffff82111715611948576119476118f3565b5b80604052505050565b5f61195a611642565b90506119668282611920565b919050565b5f67ffffffffffffffff821115611985576119846118f3565b5b61198e826118e3565b9050602081019050919050565b828183375f83830152505050565b5f6119bb6119b68461196b565b611951565b9050828152602081018484840111156119d7576119d66118df565b5b6119e284828561199b565b509392505050565b5f82601f8301126119fe576119fd6118db565b5b8135611a0e8482602086016119a9565b91505092915050565b5f805f805f60a08688031215611a3057611a2f61164b565b5b5f86013567ffffffffffffffff811115611a4d57611a4c61164f565b5b611a59888289016119ea565b9550506020611a6a8882890161181c565b945050604086013567ffffffffffffffff811115611a8b57611a8a61164f565b5b611a97888289016119ea565b9350506060611aa88882890161184f565b9250506080611ab98882890161184f565b9150509295509295909350565b5f819050919050565b5f611ae9611ae4611adf846117d6565b611ac6565b6117d6565b9050919050565b5f611afa82611acf565b9050919050565b5f611b0b82611af0565b9050919050565b611b1b81611b01565b82525050565b5f602082019050611b345f830184611b12565b92915050565b5f60208284031215611b4f57611b4e61164b565b5b5f611b5c8482850161181c565b91505092915050565b5f82825260208201905092915050565b7f4f6e6c792074686520636f6e74726163742063616e2063616c6c2074686973205f8201527f66756e6374696f6e000000000000000000000000000000000000000000000000602082015250565b5f611bcf602883611b65565b9150611bda82611b75565b604082019050919050565b5f6020820190508181035f830152611bfc81611bc3565b9050919050565b7f7772617070656420746f6b656e20646f65736e277420657869737400000000005f82015250565b5f611c37601b83611b65565b9150611c4282611c03565b602082019050919050565b5f6020820190508181035f830152611c6481611c2b565b9050919050565b611c7481611830565b82525050565b5f604082019050611c8d5f8301856118b3565b611c9a6020830184611c6b565b9392505050565b5f81519050611caf81611743565b92915050565b5f60208284031215611cca57611cc961164b565b5b5f611cd784828501611ca1565b91505092915050565b5f60ff82169050919050565b611cf581611ce0565b8114611cff575f80fd5b50565b5f81519050611d1081611cec565b92915050565b5f60208284031215611d2b57611d2a61164b565b5b5f611d3884828501611d02565b91505092915050565b5f606082019050611d545f8301866118b3565b611d6160208301856118b3565b611d6e6040830184611c6b565b949350505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f611dad82611706565b9150611db883611706565b9250828201905067ffffffffffffffff811115611dd857611dd7611d76565b5b92915050565b5f81519050919050565b8281835e5f83830152505050565b5f611e0082611dde565b611e0a8185611b65565b9350611e1a818560208601611de8565b611e23816118e3565b840191505092915050565b5f6020820190508181035f830152611e468184611df6565b905092915050565b5f611e5882611ce0565b9150611e6383611ce0565b9250828203905060ff811115611e7c57611e7b611d76565b5b92915050565b5f8160011c9050919050565b5f808291508390505b6001851115611ed757808604811115611eb357611eb2611d76565b5b6001851615611ec25780820291505b8081029050611ed085611e82565b9450611e97565b94509492505050565b5f82611eef5760019050611faa565b81611efc575f9050611faa565b8160018114611f125760028114611f1c57611f4b565b6001915050611faa565b60ff841115611f2e57611f2d611d76565b5b8360020a915084821115611f4557611f44611d76565b5b50611faa565b5060208310610133831016604e8410600b8410161715611f805782820a905083811115611f7b57611f7a611d76565b5b611faa565b611f8d8484846001611e8e565b92509050818404811115611fa457611fa3611d76565b5b81810290505b9392505050565b5f611fbb82611830565b9150611fc683611830565b9250611ff37fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8484611ee0565b905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f61203282611830565b915061203d83611830565b92508261204d5761204c611ffb565b5b828204905092915050565b5f61206282611830565b915061206d83611830565b925082820261207b81611830565b9150828204841483151761209257612091611d76565b5b5092915050565b7f636f6e76657274656420616d6f756e74206973207a65726f00000000000000005f82015250565b5f6120cd601883611b65565b91506120d882612099565b602082019050919050565b5f6020820190508181035f8301526120fa816120c1565b9050919050565b5f61211361210e8461196b565b611951565b90508281526020810184848401111561212f5761212e6118df565b5b61213a848285611de8565b509392505050565b5f82601f830112612156576121556118db565b5b8151612166848260208601612101565b91505092915050565b5f602082840312156121845761218361164b565b5b5f82015167ffffffffffffffff8111156121a1576121a061164f565b5b6121ad84828501612142565b91505092915050565b5f81905092915050565b5f6121ca82611dde565b6121d481856121b6565b93506121e4818560208601611de8565b80840191505092915050565b5f6121fb82856121c0565b915061220782846121c0565b91508190509392505050565b61221c81611ce0565b82525050565b5f6060820190508181035f83015261223a8186611df6565b9050818103602083015261224e8185611df6565b905061225d6040830184612213565b949350505050565b5f8151905061227381611806565b92915050565b5f6020828403121561228e5761228d61164b565b5b5f61229b84828501612265565b91505092915050565b7f7b224074797065223a20222f6962632e6170706c69636174696f6e732e7472615f8201527f6e736665722e76312e4d73675472616e73666572222c00000000000000000000602082015250565b5f6122fe6036836121b6565b9150612309826122a4565b603682019050919050565b7f22736f757263655f706f7274223a20227472616e73666572222c0000000000005f82015250565b5f612348601a836121b6565b915061235382612314565b601a82019050919050565b7f22736f757263655f6368616e6e656c223a2022000000000000000000000000005f82015250565b5f6123926013836121b6565b915061239d8261235e565b601382019050919050565b7f222c0000000000000000000000000000000000000000000000000000000000005f82015250565b5f6123dc6002836121b6565b91506123e7826123a8565b600282019050919050565b7f22746f6b656e223a207b202264656e6f6d223a202200000000000000000000005f82015250565b5f6124266015836121b6565b9150612431826123f2565b601582019050919050565b7f22616d6f756e74223a20220000000000000000000000000000000000000000005f82015250565b5f612470600b836121b6565b915061247b8261243c565b600b82019050919050565b7f227d2c00000000000000000000000000000000000000000000000000000000005f82015250565b5f6124ba6003836121b6565b91506124c582612486565b600382019050919050565b7f2273656e646572223a20220000000000000000000000000000000000000000005f82015250565b5f612504600b836121b6565b915061250f826124d0565b600b82019050919050565b7f227265636569766572223a2022000000000000000000000000000000000000005f82015250565b5f61254e600d836121b6565b91506125598261251a565b600d82019050919050565b7f2274696d656f75745f686569676874223a207b227265766973696f6e5f6e756d5f8201527f626572223a202230222c227265766973696f6e5f686569676874223a2022302260208201527f7d2c000000000000000000000000000000000000000000000000000000000000604082015250565b5f6125e46042836121b6565b91506125ef82612564565b604282019050919050565b7f2274696d656f75745f74696d657374616d70223a2022000000000000000000005f82015250565b5f61262e6016836121b6565b9150612639826125fa565b601682019050919050565b7f226d656d6f223a20227b5c2265766d5c223a207b5c226173796e635f63616c6c5f8201527f6261636b5c223a207b5c2269645c223a20000000000000000000000000000000602082015250565b5f61269e6031836121b6565b91506126a982612644565b603182019050919050565b7f2c5c22636f6e74726163745f616464726573735c223a5c2200000000000000005f82015250565b5f6126e86018836121b6565b91506126f3826126b4565b601882019050919050565b7f5c227d7d7d227d000000000000000000000000000000000000000000000000005f82015250565b5f6127326007836121b6565b915061273d826126fe565b600782019050919050565b5f612752826122f2565b915061275d8261233c565b915061276882612386565b9150612774828b6121c0565b915061277f826123d0565b915061278a8261241a565b9150612796828a6121c0565b91506127a1826123d0565b91506127ac82612464565b91506127b882896121c0565b91506127c3826124ae565b91506127ce826124f8565b91506127da82886121c0565b91506127e5826123d0565b91506127f082612542565b91506127fc82876121c0565b9150612807826123d0565b9150612812826125d8565b915061281d82612622565b915061282982866121c0565b9150612834826123d0565b915061283f82612692565b915061284b82856121c0565b9150612856826126dc565b915061286282846121c0565b915061286d82612726565b91508190509998505050505050505050565b5f61288982611830565b915061289483611830565b92508282019050808211156128ac576128ab611d76565b5b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f6128e982611830565b91505f82036128fb576128fa611d76565b5b600182039050919050565b5f6040820190506129195f830185611c6b565b6129266020830184611c6b565b939250505056fea264697066735822122002565a023112daa953f4c45aa3cfc8d10535e7cfecbf8e1efa2252cff9cd7d1064736f6c63430008190033",
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
