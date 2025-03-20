// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package counter

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

// CounterMetaData contains all meta data concerning the Counter contract.
var CounterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"StringsInsufficientHexLength\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"callback_id\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"name\":\"callback_received\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newCount\",\"type\":\"uint256\"}],\"name\":\"increased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"n\",\"type\":\"uint64\"}],\"name\":\"recursive_called\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"callback_id\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"name\":\"callback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"count\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"exec_msg\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"gas_limit\",\"type\":\"uint64\"}],\"name\":\"disable_execute_cosmos\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"exec_msg\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"gas_limit\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"call_revert\",\"type\":\"bool\"}],\"name\":\"execute_cosmos\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"exec_msg\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"gas_limit\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"allow_failure\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"callback_id\",\"type\":\"uint64\"}],\"name\":\"execute_cosmos_with_options\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"n\",\"type\":\"uint64\"}],\"name\":\"get_blockhash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"callback_id\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"name\":\"ibc_ack\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"callback_id\",\"type\":\"uint64\"}],\"name\":\"ibc_timeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"increase\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"num\",\"type\":\"uint64\"}],\"name\":\"increase_for_fuzz\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"n\",\"type\":\"uint64\"}],\"name\":\"nested_recursive_revert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"path\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"req\",\"type\":\"string\"}],\"name\":\"query_cosmos\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"result\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"n\",\"type\":\"uint64\"}],\"name\":\"recursive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"n\",\"type\":\"uint64\"}],\"name\":\"recursive_revert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052612072806100115f395ff3fe6080604052600436106100dc575f3560e01c80637aa4d9cc1161007e578063c31925a711610058578063c31925a714610286578063cad23554146102ae578063df3f7250146102ea578063e8927fbc14610312576100dc565b80637aa4d9cc146101fa57806395eb891214610222578063ac7fde5f1461024a576100dc565b806331a503f0116100ba57806331a503f01461015a57806352dadc5a146101825780635b133d02146101aa57806361936895146101d2576100dc565b806306661abd146100e05780630d4f1f9d1461010a5780632607baf814610132575b5f80fd5b3480156100eb575f80fd5b506100f461031c565b6040516101019190611125565b60405180910390f35b348015610115575f80fd5b50610130600480360381019061012b91906111c1565b610321565b005b34801561013d575f80fd5b50610158600480360381019061015391906111ff565b610368565b005b348015610165575f80fd5b50610180600480360381019061017b91906111ff565b61039b565b005b34801561018d575f80fd5b506101a860048036038101906101a39190611366565b6103bf565b005b3480156101b5575f80fd5b506101d060048036038101906101cb91906111ff565b610464565b005b3480156101dd575f80fd5b506101f860048036038101906101f391906111ff565b610513565b005b348015610205575f80fd5b50610220600480360381019061021b91906113e6565b610738565b005b34801561022d575f80fd5b5061024860048036038101906102439190611452565b6107fb565b005b348015610255575f80fd5b50610270600480360381019061026b91906111ff565b6108ec565b60405161027d91906114c4565b60405180910390f35b348015610291575f80fd5b506102ac60048036038101906102a791906111c1565b610900565b005b3480156102b9575f80fd5b506102d460048036038101906102cf91906114dd565b61093d565b6040516102e191906115b3565b60405180910390f35b3480156102f5575f80fd5b50610310600480360381019061030b91906111ff565b6109c5565b005b61031a610ab4565b005b5f5481565b801561034d578167ffffffffffffffff165f808282546103419190611600565b92505081905550610364565b5f8081548092919061035e90611633565b91905055505b5050565b5f8167ffffffffffffffff16031561039857610382610ab4565b610397600182610392919061167a565b610368565b5b50565b8067ffffffffffffffff165f808282546103b59190611600565b9250508190555050565b60f173ffffffffffffffffffffffffffffffffffffffff1663f1ed795d8585604051806040016040528087151581526020018667ffffffffffffffff168152506040518463ffffffff1660e01b815260040161041d9392919061170f565b6020604051808303815f875af1158015610439573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061045d919061175f565b5050505050565b7f4475bbd12ef452e28d39c4cb77494d85136c2d89ca1354b52188d4aaa8f4ba5181604051610493919061178a565b60405180910390a15f8167ffffffffffffffff160315610510573073ffffffffffffffffffffffffffffffffffffffff1663df3f7250826040518263ffffffff1660e01b81526004016104e6919061178a565b5f604051808303815f87803b1580156104fd575f80fd5b505af192505050801561050e575060015b505b50565b7f4475bbd12ef452e28d39c4cb77494d85136c2d89ca1354b52188d4aaa8f4ba5181604051610542919061178a565b60405180910390a15f8167ffffffffffffffff1603156107355760f173ffffffffffffffffffffffffffffffffffffffff166356c657a561058283610b13565b8361271061059091906117a3565b61753061059d91906117df565b67ffffffffffffffff16600180866105b591906117df565b60026105c19190611949565b6105cb9190611993565b8567ffffffffffffffff166105e091906119c6565b6105ea91906119c6565b6040518363ffffffff1660e01b8152600401610607929190611a07565b6020604051808303815f875af1158015610623573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610647919061175f565b5060f173ffffffffffffffffffffffffffffffffffffffff166356c657a561066e83610b13565b8361271061067c91906117a3565b61753061068991906117df565b67ffffffffffffffff16600180866106a191906117df565b60026106ad9190611949565b6106b79190611993565b8567ffffffffffffffff166106cc91906119c6565b6106d691906119c6565b6040518363ffffffff1660e01b81526004016106f3929190611a07565b6020604051808303815f875af115801561070f573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610733919061175f565b505b50565b60f173ffffffffffffffffffffffffffffffffffffffff166356c657a584846040518363ffffffff1660e01b8152600401610774929190611a07565b6020604051808303815f875af1158015610790573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906107b4919061175f565b5080156107f6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107ed90611aa5565b60405180910390fd5b505050565b60f173ffffffffffffffffffffffffffffffffffffffff16638c1370cd6040518163ffffffff1660e01b81526004016020604051808303815f875af1158015610846573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061086a919061175f565b5060f173ffffffffffffffffffffffffffffffffffffffff166356c657a583836040518363ffffffff1660e01b81526004016108a7929190611a07565b6020604051808303815f875af11580156108c3573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906108e7919061175f565b505050565b5f8167ffffffffffffffff16409050919050565b7fa019c7431cdfd7ba63501ffa1ba7d8f2a028e447653a5af5a96077e5038e03398282604051610931929190611ad2565b60405180910390a15050565b606060f173ffffffffffffffffffffffffffffffffffffffff1663cad2355484846040518363ffffffff1660e01b815260040161097b929190611af9565b5f60405180830381865afa158015610995573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f820116820180604052508101906109bd9190611b9c565b905092915050565b60f173ffffffffffffffffffffffffffffffffffffffff166356c657a56109eb83610b13565b836127106109f991906117a3565b617530610a0691906117df565b67ffffffffffffffff1660018086610a1e91906117df565b6002610a2a9190611949565b610a349190611993565b8567ffffffffffffffff16610a4991906119c6565b610a5391906119c6565b6040518363ffffffff1660e01b8152600401610a70929190611a07565b6020604051808303815f875af1158015610a8c573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610ab0919061175f565b5f80fd5b5f80815480929190610ac590611633565b91905055507f61996fe196f72cb598c483e896a1221263a28bb630480aa89495f737d4a8e3df60015f54610af99190611993565b5f54604051610b09929190611be3565b60405180910390a1565b606060f173ffffffffffffffffffffffffffffffffffffffff16636af32a55306040518263ffffffff1660e01b8152600401610b4f9190611c49565b5f60405180830381865afa158015610b69573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f82011682018060405250810190610b919190611b9c565b610b9a30610c1e565b610bf6636193689560e01b600186610bb2919061167a565b604051602001610bc2919061178a565b604051602081830303815290604052604051602001610be2929190611cf1565b604051602081830303815290604052610c4b565b604051602001610c0893929190611f7e565b6040516020818303038152906040529050919050565b6060610c448273ffffffffffffffffffffffffffffffffffffffff16601460ff16610ecf565b9050919050565b60605f6002808451610c5d91906119c6565b610c679190611600565b67ffffffffffffffff811115610c8057610c7f611242565b5b6040519080825280601f01601f191660200182016040528015610cb25781602001600182028036833780820191505090505b5090507f3000000000000000000000000000000000000000000000000000000000000000815f81518110610ce957610ce8612011565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a9053507f780000000000000000000000000000000000000000000000000000000000000081600181518110610d4c57610d4b612011565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a9053505f5b8351811015610ec5575f848281518110610d9957610d98612011565b5b602001015160f81c60f81b60f81c90507f303132333435363738396162636465660000000000000000000000000000000060048260ff16901c60ff1660108110610de657610de5612011565b5b1a60f81b8360028085020181518110610e0257610e01612011565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a9053507f3031323334353637383961626364656600000000000000000000000000000000600f821660ff1660108110610e6957610e68612011565b5b1a60f81b836002600160028602010181518110610e8957610e88612011565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a905350508080600101915050610d7c565b5080915050919050565b60605f8390505f6002846002610ee591906119c6565b610eef9190611600565b67ffffffffffffffff811115610f0857610f07611242565b5b6040519080825280601f01601f191660200182016040528015610f3a5781602001600182028036833780820191505090505b5090507f3000000000000000000000000000000000000000000000000000000000000000815f81518110610f7157610f70612011565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a9053507f780000000000000000000000000000000000000000000000000000000000000081600181518110610fd457610fd3612011565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a9053505f600185600261101291906119c6565b61101c9190611600565b90505b60018111156110bb577f3031323334353637383961626364656600000000000000000000000000000000600f84166010811061105e5761105d612011565b5b1a60f81b82828151811061107557611074612011565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a905350600483901c9250806110b49061203e565b905061101f565b505f82146111025784846040517fe22e27eb0000000000000000000000000000000000000000000000000000000081526004016110f9929190611be3565b60405180910390fd5b809250505092915050565b5f819050919050565b61111f8161110d565b82525050565b5f6020820190506111385f830184611116565b92915050565b5f604051905090565b5f80fd5b5f80fd5b5f67ffffffffffffffff82169050919050565b61116b8161114f565b8114611175575f80fd5b50565b5f8135905061118681611162565b92915050565b5f8115159050919050565b6111a08161118c565b81146111aa575f80fd5b50565b5f813590506111bb81611197565b92915050565b5f80604083850312156111d7576111d6611147565b5b5f6111e485828601611178565b92505060206111f5858286016111ad565b9150509250929050565b5f6020828403121561121457611213611147565b5b5f61122184828501611178565b91505092915050565b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b61127882611232565b810181811067ffffffffffffffff8211171561129757611296611242565b5b80604052505050565b5f6112a961113e565b90506112b5828261126f565b919050565b5f67ffffffffffffffff8211156112d4576112d3611242565b5b6112dd82611232565b9050602081019050919050565b828183375f83830152505050565b5f61130a611305846112ba565b6112a0565b9050828152602081018484840111156113265761132561122e565b5b6113318482856112ea565b509392505050565b5f82601f83011261134d5761134c61122a565b5b813561135d8482602086016112f8565b91505092915050565b5f805f806080858703121561137e5761137d611147565b5b5f85013567ffffffffffffffff81111561139b5761139a61114b565b5b6113a787828801611339565b94505060206113b887828801611178565b93505060406113c9878288016111ad565b92505060606113da87828801611178565b91505092959194509250565b5f805f606084860312156113fd576113fc611147565b5b5f84013567ffffffffffffffff81111561141a5761141961114b565b5b61142686828701611339565b935050602061143786828701611178565b9250506040611448868287016111ad565b9150509250925092565b5f806040838503121561146857611467611147565b5b5f83013567ffffffffffffffff8111156114855761148461114b565b5b61149185828601611339565b92505060206114a285828601611178565b9150509250929050565b5f819050919050565b6114be816114ac565b82525050565b5f6020820190506114d75f8301846114b5565b92915050565b5f80604083850312156114f3576114f2611147565b5b5f83013567ffffffffffffffff8111156115105761150f61114b565b5b61151c85828601611339565b925050602083013567ffffffffffffffff81111561153d5761153c61114b565b5b61154985828601611339565b9150509250929050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f61158582611553565b61158f818561155d565b935061159f81856020860161156d565b6115a881611232565b840191505092915050565b5f6020820190508181035f8301526115cb818461157b565b905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61160a8261110d565b91506116158361110d565b925082820190508082111561162d5761162c6115d3565b5b92915050565b5f61163d8261110d565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361166f5761166e6115d3565b5b600182019050919050565b5f6116848261114f565b915061168f8361114f565b9250828203905067ffffffffffffffff8111156116af576116ae6115d3565b5b92915050565b6116be8161114f565b82525050565b6116cd8161118c565b82525050565b6116dc8161114f565b82525050565b604082015f8201516116f65f8501826116c4565b50602082015161170960208501826116d3565b50505050565b5f6080820190508181035f830152611727818661157b565b905061173660208301856116b5565b61174360408301846116e2565b949350505050565b5f8151905061175981611197565b92915050565b5f6020828403121561177457611773611147565b5b5f6117818482850161174b565b91505092915050565b5f60208201905061179d5f8301846116b5565b92915050565b5f6117ad8261114f565b91506117b88361114f565b92508282026117c68161114f565b91508082146117d8576117d76115d3565b5b5092915050565b5f6117e98261114f565b91506117f48361114f565b9250828201905067ffffffffffffffff811115611814576118136115d3565b5b92915050565b5f8160011c9050919050565b5f808291508390505b600185111561186f5780860481111561184b5761184a6115d3565b5b600185161561185a5780820291505b80810290506118688561181a565b945061182f565b94509492505050565b5f826118875760019050611942565b81611894575f9050611942565b81600181146118aa57600281146118b4576118e3565b6001915050611942565b60ff8411156118c6576118c56115d3565b5b8360020a9150848211156118dd576118dc6115d3565b5b50611942565b5060208310610133831016604e8410600b84101617156119185782820a905083811115611913576119126115d3565b5b611942565b6119258484846001611826565b9250905081840481111561193c5761193b6115d3565b5b81810290505b9392505050565b5f6119538261110d565b915061195e8361114f565b925061198b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8484611878565b905092915050565b5f61199d8261110d565b91506119a88361110d565b92508282039050818111156119c0576119bf6115d3565b5b92915050565b5f6119d08261110d565b91506119db8361110d565b92508282026119e98161110d565b91508282048414831517611a00576119ff6115d3565b5b5092915050565b5f6040820190508181035f830152611a1f818561157b565b9050611a2e60208301846116b5565b9392505050565b7f72657665727420726561736f6e2064756d6d792076616c756520666f722074655f8201527f7374000000000000000000000000000000000000000000000000000000000000602082015250565b5f611a8f60228361155d565b9150611a9a82611a35565b604082019050919050565b5f6020820190508181035f830152611abc81611a83565b9050919050565b611acc8161118c565b82525050565b5f604082019050611ae55f8301856116b5565b611af26020830184611ac3565b9392505050565b5f6040820190508181035f830152611b11818561157b565b90508181036020830152611b25818461157b565b90509392505050565b5f611b40611b3b846112ba565b6112a0565b905082815260208101848484011115611b5c57611b5b61122e565b5b611b6784828561156d565b509392505050565b5f82601f830112611b8357611b8261122a565b5b8151611b93848260208601611b2e565b91505092915050565b5f60208284031215611bb157611bb0611147565b5b5f82015167ffffffffffffffff811115611bce57611bcd61114b565b5b611bda84828501611b6f565b91505092915050565b5f604082019050611bf65f830185611116565b611c036020830184611116565b9392505050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f611c3382611c0a565b9050919050565b611c4381611c29565b82525050565b5f602082019050611c5c5f830184611c3a565b92915050565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b5f819050919050565b611ca7611ca282611c62565b611c8d565b82525050565b5f81519050919050565b5f81905092915050565b5f611ccb82611cad565b611cd58185611cb7565b9350611ce581856020860161156d565b80840191505092915050565b5f611cfc8285611c96565b600482019150611d0c8284611cc1565b91508190509392505050565b5f81905092915050565b7f7b224074797065223a20222f6d696e6965766d2e65766d2e76312e4d736743615f8201527f6c6c222c00000000000000000000000000000000000000000000000000000000602082015250565b5f611d7c602483611d18565b9150611d8782611d22565b602482019050919050565b7f2273656e646572223a20220000000000000000000000000000000000000000005f82015250565b5f611dc6600b83611d18565b9150611dd182611d92565b600b82019050919050565b5f611de682611553565b611df08185611d18565b9350611e0081856020860161156d565b80840191505092915050565b7f222c0000000000000000000000000000000000000000000000000000000000005f82015250565b5f611e40600283611d18565b9150611e4b82611e0c565b600282019050919050565b7f22636f6e74726163745f61646472223a202200000000000000000000000000005f82015250565b5f611e8a601283611d18565b9150611e9582611e56565b601282019050919050565b7f22696e707574223a2022000000000000000000000000000000000000000000005f82015250565b5f611ed4600a83611d18565b9150611edf82611ea0565b600a82019050919050565b7f2276616c7565223a202230222c000000000000000000000000000000000000005f82015250565b5f611f1e600d83611d18565b9150611f2982611eea565b600d82019050919050565b7f226163636573735f6c697374223a205b5d7d00000000000000000000000000005f82015250565b5f611f68601283611d18565b9150611f7382611f34565b601282019050919050565b5f611f8882611d70565b9150611f9382611dba565b9150611f9f8286611ddc565b9150611faa82611e34565b9150611fb582611e7e565b9150611fc18285611ddc565b9150611fcc82611e34565b9150611fd782611ec8565b9150611fe38284611ddc565b9150611fee82611e34565b9150611ff982611f12565b915061200482611f5c565b9150819050949350505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f6120488261110d565b91505f820361205a576120596115d3565b5b60018203905091905056fea164736f6c6343000819000a",
}

// CounterABI is the input ABI used to generate the binding from.
// Deprecated: Use CounterMetaData.ABI instead.
var CounterABI = CounterMetaData.ABI

// CounterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CounterMetaData.Bin instead.
var CounterBin = CounterMetaData.Bin

// DeployCounter deploys a new Ethereum contract, binding an instance of Counter to it.
func DeployCounter(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Counter, error) {
	parsed, err := CounterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CounterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Counter{CounterCaller: CounterCaller{contract: contract}, CounterTransactor: CounterTransactor{contract: contract}, CounterFilterer: CounterFilterer{contract: contract}}, nil
}

// Counter is an auto generated Go binding around an Ethereum contract.
type Counter struct {
	CounterCaller     // Read-only binding to the contract
	CounterTransactor // Write-only binding to the contract
	CounterFilterer   // Log filterer for contract events
}

// CounterCaller is an auto generated read-only Go binding around an Ethereum contract.
type CounterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CounterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CounterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CounterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CounterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CounterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CounterSession struct {
	Contract     *Counter          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CounterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CounterCallerSession struct {
	Contract *CounterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// CounterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CounterTransactorSession struct {
	Contract     *CounterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// CounterRaw is an auto generated low-level Go binding around an Ethereum contract.
type CounterRaw struct {
	Contract *Counter // Generic contract binding to access the raw methods on
}

// CounterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CounterCallerRaw struct {
	Contract *CounterCaller // Generic read-only contract binding to access the raw methods on
}

// CounterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CounterTransactorRaw struct {
	Contract *CounterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCounter creates a new instance of Counter, bound to a specific deployed contract.
func NewCounter(address common.Address, backend bind.ContractBackend) (*Counter, error) {
	contract, err := bindCounter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Counter{CounterCaller: CounterCaller{contract: contract}, CounterTransactor: CounterTransactor{contract: contract}, CounterFilterer: CounterFilterer{contract: contract}}, nil
}

// NewCounterCaller creates a new read-only instance of Counter, bound to a specific deployed contract.
func NewCounterCaller(address common.Address, caller bind.ContractCaller) (*CounterCaller, error) {
	contract, err := bindCounter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CounterCaller{contract: contract}, nil
}

// NewCounterTransactor creates a new write-only instance of Counter, bound to a specific deployed contract.
func NewCounterTransactor(address common.Address, transactor bind.ContractTransactor) (*CounterTransactor, error) {
	contract, err := bindCounter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CounterTransactor{contract: contract}, nil
}

// NewCounterFilterer creates a new log filterer instance of Counter, bound to a specific deployed contract.
func NewCounterFilterer(address common.Address, filterer bind.ContractFilterer) (*CounterFilterer, error) {
	contract, err := bindCounter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CounterFilterer{contract: contract}, nil
}

// bindCounter binds a generic wrapper to an already deployed contract.
func bindCounter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CounterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Counter *CounterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Counter.Contract.CounterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Counter *CounterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Counter.Contract.CounterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Counter *CounterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Counter.Contract.CounterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Counter *CounterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Counter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Counter *CounterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Counter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Counter *CounterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Counter.Contract.contract.Transact(opts, method, params...)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_Counter *CounterCaller) Count(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Counter.contract.Call(opts, &out, "count")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_Counter *CounterSession) Count() (*big.Int, error) {
	return _Counter.Contract.Count(&_Counter.CallOpts)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_Counter *CounterCallerSession) Count() (*big.Int, error) {
	return _Counter.Contract.Count(&_Counter.CallOpts)
}

// GetBlockhash is a free data retrieval call binding the contract method 0xac7fde5f.
//
// Solidity: function get_blockhash(uint64 n) view returns(bytes32)
func (_Counter *CounterCaller) GetBlockhash(opts *bind.CallOpts, n uint64) ([32]byte, error) {
	var out []interface{}
	err := _Counter.contract.Call(opts, &out, "get_blockhash", n)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetBlockhash is a free data retrieval call binding the contract method 0xac7fde5f.
//
// Solidity: function get_blockhash(uint64 n) view returns(bytes32)
func (_Counter *CounterSession) GetBlockhash(n uint64) ([32]byte, error) {
	return _Counter.Contract.GetBlockhash(&_Counter.CallOpts, n)
}

// GetBlockhash is a free data retrieval call binding the contract method 0xac7fde5f.
//
// Solidity: function get_blockhash(uint64 n) view returns(bytes32)
func (_Counter *CounterCallerSession) GetBlockhash(n uint64) ([32]byte, error) {
	return _Counter.Contract.GetBlockhash(&_Counter.CallOpts, n)
}

// QueryCosmos is a free data retrieval call binding the contract method 0xcad23554.
//
// Solidity: function query_cosmos(string path, string req) view returns(string result)
func (_Counter *CounterCaller) QueryCosmos(opts *bind.CallOpts, path string, req string) (string, error) {
	var out []interface{}
	err := _Counter.contract.Call(opts, &out, "query_cosmos", path, req)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// QueryCosmos is a free data retrieval call binding the contract method 0xcad23554.
//
// Solidity: function query_cosmos(string path, string req) view returns(string result)
func (_Counter *CounterSession) QueryCosmos(path string, req string) (string, error) {
	return _Counter.Contract.QueryCosmos(&_Counter.CallOpts, path, req)
}

// QueryCosmos is a free data retrieval call binding the contract method 0xcad23554.
//
// Solidity: function query_cosmos(string path, string req) view returns(string result)
func (_Counter *CounterCallerSession) QueryCosmos(path string, req string) (string, error) {
	return _Counter.Contract.QueryCosmos(&_Counter.CallOpts, path, req)
}

// Callback is a paid mutator transaction binding the contract method 0xc31925a7.
//
// Solidity: function callback(uint64 callback_id, bool success) returns()
func (_Counter *CounterTransactor) Callback(opts *bind.TransactOpts, callback_id uint64, success bool) (*types.Transaction, error) {
	return _Counter.contract.Transact(opts, "callback", callback_id, success)
}

// Callback is a paid mutator transaction binding the contract method 0xc31925a7.
//
// Solidity: function callback(uint64 callback_id, bool success) returns()
func (_Counter *CounterSession) Callback(callback_id uint64, success bool) (*types.Transaction, error) {
	return _Counter.Contract.Callback(&_Counter.TransactOpts, callback_id, success)
}

// Callback is a paid mutator transaction binding the contract method 0xc31925a7.
//
// Solidity: function callback(uint64 callback_id, bool success) returns()
func (_Counter *CounterTransactorSession) Callback(callback_id uint64, success bool) (*types.Transaction, error) {
	return _Counter.Contract.Callback(&_Counter.TransactOpts, callback_id, success)
}

// DisableExecuteCosmos is a paid mutator transaction binding the contract method 0x95eb8912.
//
// Solidity: function disable_execute_cosmos(string exec_msg, uint64 gas_limit) returns()
func (_Counter *CounterTransactor) DisableExecuteCosmos(opts *bind.TransactOpts, exec_msg string, gas_limit uint64) (*types.Transaction, error) {
	return _Counter.contract.Transact(opts, "disable_execute_cosmos", exec_msg, gas_limit)
}

// DisableExecuteCosmos is a paid mutator transaction binding the contract method 0x95eb8912.
//
// Solidity: function disable_execute_cosmos(string exec_msg, uint64 gas_limit) returns()
func (_Counter *CounterSession) DisableExecuteCosmos(exec_msg string, gas_limit uint64) (*types.Transaction, error) {
	return _Counter.Contract.DisableExecuteCosmos(&_Counter.TransactOpts, exec_msg, gas_limit)
}

// DisableExecuteCosmos is a paid mutator transaction binding the contract method 0x95eb8912.
//
// Solidity: function disable_execute_cosmos(string exec_msg, uint64 gas_limit) returns()
func (_Counter *CounterTransactorSession) DisableExecuteCosmos(exec_msg string, gas_limit uint64) (*types.Transaction, error) {
	return _Counter.Contract.DisableExecuteCosmos(&_Counter.TransactOpts, exec_msg, gas_limit)
}

// ExecuteCosmos is a paid mutator transaction binding the contract method 0x7aa4d9cc.
//
// Solidity: function execute_cosmos(string exec_msg, uint64 gas_limit, bool call_revert) returns()
func (_Counter *CounterTransactor) ExecuteCosmos(opts *bind.TransactOpts, exec_msg string, gas_limit uint64, call_revert bool) (*types.Transaction, error) {
	return _Counter.contract.Transact(opts, "execute_cosmos", exec_msg, gas_limit, call_revert)
}

// ExecuteCosmos is a paid mutator transaction binding the contract method 0x7aa4d9cc.
//
// Solidity: function execute_cosmos(string exec_msg, uint64 gas_limit, bool call_revert) returns()
func (_Counter *CounterSession) ExecuteCosmos(exec_msg string, gas_limit uint64, call_revert bool) (*types.Transaction, error) {
	return _Counter.Contract.ExecuteCosmos(&_Counter.TransactOpts, exec_msg, gas_limit, call_revert)
}

// ExecuteCosmos is a paid mutator transaction binding the contract method 0x7aa4d9cc.
//
// Solidity: function execute_cosmos(string exec_msg, uint64 gas_limit, bool call_revert) returns()
func (_Counter *CounterTransactorSession) ExecuteCosmos(exec_msg string, gas_limit uint64, call_revert bool) (*types.Transaction, error) {
	return _Counter.Contract.ExecuteCosmos(&_Counter.TransactOpts, exec_msg, gas_limit, call_revert)
}

// ExecuteCosmosWithOptions is a paid mutator transaction binding the contract method 0x52dadc5a.
//
// Solidity: function execute_cosmos_with_options(string exec_msg, uint64 gas_limit, bool allow_failure, uint64 callback_id) returns()
func (_Counter *CounterTransactor) ExecuteCosmosWithOptions(opts *bind.TransactOpts, exec_msg string, gas_limit uint64, allow_failure bool, callback_id uint64) (*types.Transaction, error) {
	return _Counter.contract.Transact(opts, "execute_cosmos_with_options", exec_msg, gas_limit, allow_failure, callback_id)
}

// ExecuteCosmosWithOptions is a paid mutator transaction binding the contract method 0x52dadc5a.
//
// Solidity: function execute_cosmos_with_options(string exec_msg, uint64 gas_limit, bool allow_failure, uint64 callback_id) returns()
func (_Counter *CounterSession) ExecuteCosmosWithOptions(exec_msg string, gas_limit uint64, allow_failure bool, callback_id uint64) (*types.Transaction, error) {
	return _Counter.Contract.ExecuteCosmosWithOptions(&_Counter.TransactOpts, exec_msg, gas_limit, allow_failure, callback_id)
}

// ExecuteCosmosWithOptions is a paid mutator transaction binding the contract method 0x52dadc5a.
//
// Solidity: function execute_cosmos_with_options(string exec_msg, uint64 gas_limit, bool allow_failure, uint64 callback_id) returns()
func (_Counter *CounterTransactorSession) ExecuteCosmosWithOptions(exec_msg string, gas_limit uint64, allow_failure bool, callback_id uint64) (*types.Transaction, error) {
	return _Counter.Contract.ExecuteCosmosWithOptions(&_Counter.TransactOpts, exec_msg, gas_limit, allow_failure, callback_id)
}

// IbcAck is a paid mutator transaction binding the contract method 0x0d4f1f9d.
//
// Solidity: function ibc_ack(uint64 callback_id, bool success) returns()
func (_Counter *CounterTransactor) IbcAck(opts *bind.TransactOpts, callback_id uint64, success bool) (*types.Transaction, error) {
	return _Counter.contract.Transact(opts, "ibc_ack", callback_id, success)
}

// IbcAck is a paid mutator transaction binding the contract method 0x0d4f1f9d.
//
// Solidity: function ibc_ack(uint64 callback_id, bool success) returns()
func (_Counter *CounterSession) IbcAck(callback_id uint64, success bool) (*types.Transaction, error) {
	return _Counter.Contract.IbcAck(&_Counter.TransactOpts, callback_id, success)
}

// IbcAck is a paid mutator transaction binding the contract method 0x0d4f1f9d.
//
// Solidity: function ibc_ack(uint64 callback_id, bool success) returns()
func (_Counter *CounterTransactorSession) IbcAck(callback_id uint64, success bool) (*types.Transaction, error) {
	return _Counter.Contract.IbcAck(&_Counter.TransactOpts, callback_id, success)
}

// IbcTimeout is a paid mutator transaction binding the contract method 0x31a503f0.
//
// Solidity: function ibc_timeout(uint64 callback_id) returns()
func (_Counter *CounterTransactor) IbcTimeout(opts *bind.TransactOpts, callback_id uint64) (*types.Transaction, error) {
	return _Counter.contract.Transact(opts, "ibc_timeout", callback_id)
}

// IbcTimeout is a paid mutator transaction binding the contract method 0x31a503f0.
//
// Solidity: function ibc_timeout(uint64 callback_id) returns()
func (_Counter *CounterSession) IbcTimeout(callback_id uint64) (*types.Transaction, error) {
	return _Counter.Contract.IbcTimeout(&_Counter.TransactOpts, callback_id)
}

// IbcTimeout is a paid mutator transaction binding the contract method 0x31a503f0.
//
// Solidity: function ibc_timeout(uint64 callback_id) returns()
func (_Counter *CounterTransactorSession) IbcTimeout(callback_id uint64) (*types.Transaction, error) {
	return _Counter.Contract.IbcTimeout(&_Counter.TransactOpts, callback_id)
}

// Increase is a paid mutator transaction binding the contract method 0xe8927fbc.
//
// Solidity: function increase() payable returns()
func (_Counter *CounterTransactor) Increase(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Counter.contract.Transact(opts, "increase")
}

// Increase is a paid mutator transaction binding the contract method 0xe8927fbc.
//
// Solidity: function increase() payable returns()
func (_Counter *CounterSession) Increase() (*types.Transaction, error) {
	return _Counter.Contract.Increase(&_Counter.TransactOpts)
}

// Increase is a paid mutator transaction binding the contract method 0xe8927fbc.
//
// Solidity: function increase() payable returns()
func (_Counter *CounterTransactorSession) Increase() (*types.Transaction, error) {
	return _Counter.Contract.Increase(&_Counter.TransactOpts)
}

// IncreaseForFuzz is a paid mutator transaction binding the contract method 0x2607baf8.
//
// Solidity: function increase_for_fuzz(uint64 num) returns()
func (_Counter *CounterTransactor) IncreaseForFuzz(opts *bind.TransactOpts, num uint64) (*types.Transaction, error) {
	return _Counter.contract.Transact(opts, "increase_for_fuzz", num)
}

// IncreaseForFuzz is a paid mutator transaction binding the contract method 0x2607baf8.
//
// Solidity: function increase_for_fuzz(uint64 num) returns()
func (_Counter *CounterSession) IncreaseForFuzz(num uint64) (*types.Transaction, error) {
	return _Counter.Contract.IncreaseForFuzz(&_Counter.TransactOpts, num)
}

// IncreaseForFuzz is a paid mutator transaction binding the contract method 0x2607baf8.
//
// Solidity: function increase_for_fuzz(uint64 num) returns()
func (_Counter *CounterTransactorSession) IncreaseForFuzz(num uint64) (*types.Transaction, error) {
	return _Counter.Contract.IncreaseForFuzz(&_Counter.TransactOpts, num)
}

// NestedRecursiveRevert is a paid mutator transaction binding the contract method 0xdf3f7250.
//
// Solidity: function nested_recursive_revert(uint64 n) returns()
func (_Counter *CounterTransactor) NestedRecursiveRevert(opts *bind.TransactOpts, n uint64) (*types.Transaction, error) {
	return _Counter.contract.Transact(opts, "nested_recursive_revert", n)
}

// NestedRecursiveRevert is a paid mutator transaction binding the contract method 0xdf3f7250.
//
// Solidity: function nested_recursive_revert(uint64 n) returns()
func (_Counter *CounterSession) NestedRecursiveRevert(n uint64) (*types.Transaction, error) {
	return _Counter.Contract.NestedRecursiveRevert(&_Counter.TransactOpts, n)
}

// NestedRecursiveRevert is a paid mutator transaction binding the contract method 0xdf3f7250.
//
// Solidity: function nested_recursive_revert(uint64 n) returns()
func (_Counter *CounterTransactorSession) NestedRecursiveRevert(n uint64) (*types.Transaction, error) {
	return _Counter.Contract.NestedRecursiveRevert(&_Counter.TransactOpts, n)
}

// Recursive is a paid mutator transaction binding the contract method 0x61936895.
//
// Solidity: function recursive(uint64 n) returns()
func (_Counter *CounterTransactor) Recursive(opts *bind.TransactOpts, n uint64) (*types.Transaction, error) {
	return _Counter.contract.Transact(opts, "recursive", n)
}

// Recursive is a paid mutator transaction binding the contract method 0x61936895.
//
// Solidity: function recursive(uint64 n) returns()
func (_Counter *CounterSession) Recursive(n uint64) (*types.Transaction, error) {
	return _Counter.Contract.Recursive(&_Counter.TransactOpts, n)
}

// Recursive is a paid mutator transaction binding the contract method 0x61936895.
//
// Solidity: function recursive(uint64 n) returns()
func (_Counter *CounterTransactorSession) Recursive(n uint64) (*types.Transaction, error) {
	return _Counter.Contract.Recursive(&_Counter.TransactOpts, n)
}

// RecursiveRevert is a paid mutator transaction binding the contract method 0x5b133d02.
//
// Solidity: function recursive_revert(uint64 n) returns()
func (_Counter *CounterTransactor) RecursiveRevert(opts *bind.TransactOpts, n uint64) (*types.Transaction, error) {
	return _Counter.contract.Transact(opts, "recursive_revert", n)
}

// RecursiveRevert is a paid mutator transaction binding the contract method 0x5b133d02.
//
// Solidity: function recursive_revert(uint64 n) returns()
func (_Counter *CounterSession) RecursiveRevert(n uint64) (*types.Transaction, error) {
	return _Counter.Contract.RecursiveRevert(&_Counter.TransactOpts, n)
}

// RecursiveRevert is a paid mutator transaction binding the contract method 0x5b133d02.
//
// Solidity: function recursive_revert(uint64 n) returns()
func (_Counter *CounterTransactorSession) RecursiveRevert(n uint64) (*types.Transaction, error) {
	return _Counter.Contract.RecursiveRevert(&_Counter.TransactOpts, n)
}

// CounterCallbackReceivedIterator is returned from FilterCallbackReceived and is used to iterate over the raw logs and unpacked data for CallbackReceived events raised by the Counter contract.
type CounterCallbackReceivedIterator struct {
	Event *CounterCallbackReceived // Event containing the contract specifics and raw log

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
func (it *CounterCallbackReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CounterCallbackReceived)
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
		it.Event = new(CounterCallbackReceived)
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
func (it *CounterCallbackReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CounterCallbackReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CounterCallbackReceived represents a CallbackReceived event raised by the Counter contract.
type CounterCallbackReceived struct {
	CallbackId uint64
	Success    bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCallbackReceived is a free log retrieval operation binding the contract event 0xa019c7431cdfd7ba63501ffa1ba7d8f2a028e447653a5af5a96077e5038e0339.
//
// Solidity: event callback_received(uint64 callback_id, bool success)
func (_Counter *CounterFilterer) FilterCallbackReceived(opts *bind.FilterOpts) (*CounterCallbackReceivedIterator, error) {

	logs, sub, err := _Counter.contract.FilterLogs(opts, "callback_received")
	if err != nil {
		return nil, err
	}
	return &CounterCallbackReceivedIterator{contract: _Counter.contract, event: "callback_received", logs: logs, sub: sub}, nil
}

// WatchCallbackReceived is a free log subscription operation binding the contract event 0xa019c7431cdfd7ba63501ffa1ba7d8f2a028e447653a5af5a96077e5038e0339.
//
// Solidity: event callback_received(uint64 callback_id, bool success)
func (_Counter *CounterFilterer) WatchCallbackReceived(opts *bind.WatchOpts, sink chan<- *CounterCallbackReceived) (event.Subscription, error) {

	logs, sub, err := _Counter.contract.WatchLogs(opts, "callback_received")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CounterCallbackReceived)
				if err := _Counter.contract.UnpackLog(event, "callback_received", log); err != nil {
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

// ParseCallbackReceived is a log parse operation binding the contract event 0xa019c7431cdfd7ba63501ffa1ba7d8f2a028e447653a5af5a96077e5038e0339.
//
// Solidity: event callback_received(uint64 callback_id, bool success)
func (_Counter *CounterFilterer) ParseCallbackReceived(log types.Log) (*CounterCallbackReceived, error) {
	event := new(CounterCallbackReceived)
	if err := _Counter.contract.UnpackLog(event, "callback_received", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CounterIncreasedIterator is returned from FilterIncreased and is used to iterate over the raw logs and unpacked data for Increased events raised by the Counter contract.
type CounterIncreasedIterator struct {
	Event *CounterIncreased // Event containing the contract specifics and raw log

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
func (it *CounterIncreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CounterIncreased)
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
		it.Event = new(CounterIncreased)
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
func (it *CounterIncreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CounterIncreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CounterIncreased represents a Increased event raised by the Counter contract.
type CounterIncreased struct {
	OldCount *big.Int
	NewCount *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterIncreased is a free log retrieval operation binding the contract event 0x61996fe196f72cb598c483e896a1221263a28bb630480aa89495f737d4a8e3df.
//
// Solidity: event increased(uint256 oldCount, uint256 newCount)
func (_Counter *CounterFilterer) FilterIncreased(opts *bind.FilterOpts) (*CounterIncreasedIterator, error) {

	logs, sub, err := _Counter.contract.FilterLogs(opts, "increased")
	if err != nil {
		return nil, err
	}
	return &CounterIncreasedIterator{contract: _Counter.contract, event: "increased", logs: logs, sub: sub}, nil
}

// WatchIncreased is a free log subscription operation binding the contract event 0x61996fe196f72cb598c483e896a1221263a28bb630480aa89495f737d4a8e3df.
//
// Solidity: event increased(uint256 oldCount, uint256 newCount)
func (_Counter *CounterFilterer) WatchIncreased(opts *bind.WatchOpts, sink chan<- *CounterIncreased) (event.Subscription, error) {

	logs, sub, err := _Counter.contract.WatchLogs(opts, "increased")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CounterIncreased)
				if err := _Counter.contract.UnpackLog(event, "increased", log); err != nil {
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

// ParseIncreased is a log parse operation binding the contract event 0x61996fe196f72cb598c483e896a1221263a28bb630480aa89495f737d4a8e3df.
//
// Solidity: event increased(uint256 oldCount, uint256 newCount)
func (_Counter *CounterFilterer) ParseIncreased(log types.Log) (*CounterIncreased, error) {
	event := new(CounterIncreased)
	if err := _Counter.contract.UnpackLog(event, "increased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CounterRecursiveCalledIterator is returned from FilterRecursiveCalled and is used to iterate over the raw logs and unpacked data for RecursiveCalled events raised by the Counter contract.
type CounterRecursiveCalledIterator struct {
	Event *CounterRecursiveCalled // Event containing the contract specifics and raw log

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
func (it *CounterRecursiveCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CounterRecursiveCalled)
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
		it.Event = new(CounterRecursiveCalled)
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
func (it *CounterRecursiveCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CounterRecursiveCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CounterRecursiveCalled represents a RecursiveCalled event raised by the Counter contract.
type CounterRecursiveCalled struct {
	N   uint64
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRecursiveCalled is a free log retrieval operation binding the contract event 0x4475bbd12ef452e28d39c4cb77494d85136c2d89ca1354b52188d4aaa8f4ba51.
//
// Solidity: event recursive_called(uint64 n)
func (_Counter *CounterFilterer) FilterRecursiveCalled(opts *bind.FilterOpts) (*CounterRecursiveCalledIterator, error) {

	logs, sub, err := _Counter.contract.FilterLogs(opts, "recursive_called")
	if err != nil {
		return nil, err
	}
	return &CounterRecursiveCalledIterator{contract: _Counter.contract, event: "recursive_called", logs: logs, sub: sub}, nil
}

// WatchRecursiveCalled is a free log subscription operation binding the contract event 0x4475bbd12ef452e28d39c4cb77494d85136c2d89ca1354b52188d4aaa8f4ba51.
//
// Solidity: event recursive_called(uint64 n)
func (_Counter *CounterFilterer) WatchRecursiveCalled(opts *bind.WatchOpts, sink chan<- *CounterRecursiveCalled) (event.Subscription, error) {

	logs, sub, err := _Counter.contract.WatchLogs(opts, "recursive_called")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CounterRecursiveCalled)
				if err := _Counter.contract.UnpackLog(event, "recursive_called", log); err != nil {
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

// ParseRecursiveCalled is a log parse operation binding the contract event 0x4475bbd12ef452e28d39c4cb77494d85136c2d89ca1354b52188d4aaa8f4ba51.
//
// Solidity: event recursive_called(uint64 n)
func (_Counter *CounterFilterer) ParseRecursiveCalled(log types.Log) (*CounterRecursiveCalled, error) {
	event := new(CounterRecursiveCalled)
	if err := _Counter.contract.UnpackLog(event, "recursive_called", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
