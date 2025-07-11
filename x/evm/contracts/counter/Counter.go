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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"StringsInsufficientHexLength\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"callback_id\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"name\":\"callback_received\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newCount\",\"type\":\"uint256\"}],\"name\":\"increased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"n\",\"type\":\"uint64\"}],\"name\":\"recursive_called\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"callback_id\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"name\":\"callback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"count\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"exec_msg\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"gas_limit\",\"type\":\"uint64\"}],\"name\":\"disable_and_execute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"test_addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"exec_msg\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"gas_limit\",\"type\":\"uint64\"}],\"name\":\"disable_and_execute_in_child\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"test_addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"exec_msg\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"gas_limit\",\"type\":\"uint64\"}],\"name\":\"disable_and_execute_in_parent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"exec_msg\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"gas_limit\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"call_revert\",\"type\":\"bool\"}],\"name\":\"execute_cosmos\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"exec_msg\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"gas_limit\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"allow_failure\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"callback_id\",\"type\":\"uint64\"}],\"name\":\"execute_cosmos_with_options\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"n\",\"type\":\"uint64\"}],\"name\":\"get_blockhash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"callback_id\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"name\":\"ibc_ack\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"callback_id\",\"type\":\"uint64\"}],\"name\":\"ibc_timeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"increase\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"num\",\"type\":\"uint64\"}],\"name\":\"increase_for_fuzz\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"loop\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"n\",\"type\":\"uint64\"}],\"name\":\"nested_recursive_revert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"path\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"req\",\"type\":\"string\"}],\"name\":\"query_cosmos\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"result\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"n\",\"type\":\"uint64\"}],\"name\":\"recursive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"n\",\"type\":\"uint64\"}],\"name\":\"recursive_revert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040526123ad806100115f395ff3fe6080604052600436106100fd575f3560e01c80637aa4d9cc11610094578063c31925a711610063578063c31925a7146102e5578063cad235541461030d578063df3f725014610349578063e8927fbc14610371578063fbb2c5dd1461037b576100fd565b80637aa4d9cc14610243578063a92100cb1461026b578063ac7fde5f14610281578063bb78714d146102bd576100fd565b806352dadc5a116100d057806352dadc5a146101a357806353a38e8f146101cb5780635b133d02146101f3578063619368951461021b576100fd565b806306661abd146101015780630d4f1f9d1461012b5780632607baf81461015357806331a503f01461017b575b5f80fd5b34801561010c575f80fd5b506101156103a3565b60405161012291906113ca565b60405180910390f35b348015610136575f80fd5b50610151600480360381019061014c9190611466565b6103a8565b005b34801561015e575f80fd5b50610179600480360381019061017491906114a4565b6103ef565b005b348015610186575f80fd5b506101a1600480360381019061019c91906114a4565b610422565b005b3480156101ae575f80fd5b506101c960048036038101906101c4919061160b565b610446565b005b3480156101d6575f80fd5b506101f160048036038101906101ec91906116e5565b6104eb565b005b3480156101fe575f80fd5b50610219600480360381019061021491906114a4565b6105c8565b005b348015610226575f80fd5b50610241600480360381019061023c91906114a4565b610677565b005b34801561024e575f80fd5b5061026960048036038101906102649190611751565b61089c565b005b348015610276575f80fd5b5061027f61095f565b005b34801561028c575f80fd5b506102a760048036038101906102a291906114a4565b610976565b6040516102b491906117d5565b60405180910390f35b3480156102c8575f80fd5b506102e360048036038101906102de91906117ee565b61098a565b005b3480156102f0575f80fd5b5061030b60048036038101906103069190611466565b610a7b565b005b348015610318575f80fd5b50610333600480360381019061032e9190611848565b610b05565b604051610340919061191e565b60405180910390f35b348015610354575f80fd5b5061036f600480360381019061036a91906114a4565b610b8d565b005b610379610c7c565b005b348015610386575f80fd5b506103a1600480360381019061039c91906116e5565b610cdb565b005b5f5481565b80156103d4578167ffffffffffffffff165f808282546103c8919061196b565b925050819055506103eb565b5f808154809291906103e59061199e565b91905055505b5050565b5f8167ffffffffffffffff16031561041f57610409610c7c565b61041e60018261041991906119e5565b6103ef565b5b50565b8067ffffffffffffffff165f8082825461043c919061196b565b9250508190555050565b60f173ffffffffffffffffffffffffffffffffffffffff1663f1ed795d8585604051806040016040528087151581526020018667ffffffffffffffff168152506040518463ffffffff1660e01b81526004016104a493929190611a7a565b6020604051808303815f875af11580156104c0573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906104e49190611aca565b5050505050565b60f173ffffffffffffffffffffffffffffffffffffffff16638c1370cd6040518163ffffffff1660e01b81526004016020604051808303815f875af1158015610536573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061055a9190611aca565b508273ffffffffffffffffffffffffffffffffffffffff166356c657a583836040518363ffffffff1660e01b8152600401610596929190611af5565b5f604051808303815f87803b1580156105ad575f80fd5b505af11580156105bf573d5f803e3d5ffd5b50505050505050565b7f4475bbd12ef452e28d39c4cb77494d85136c2d89ca1354b52188d4aaa8f4ba51816040516105f79190611b23565b60405180910390a15f8167ffffffffffffffff160315610674573073ffffffffffffffffffffffffffffffffffffffff1663df3f7250826040518263ffffffff1660e01b815260040161064a9190611b23565b5f604051808303815f87803b158015610661575f80fd5b505af1925050508015610672575060015b505b50565b7f4475bbd12ef452e28d39c4cb77494d85136c2d89ca1354b52188d4aaa8f4ba51816040516106a69190611b23565b60405180910390a15f8167ffffffffffffffff1603156108995760f173ffffffffffffffffffffffffffffffffffffffff166356c657a56106e683610db8565b836127106106f49190611b3c565b6175306107019190611b78565b67ffffffffffffffff16600180866107199190611b78565b60026107259190611ce2565b61072f9190611d2c565b8567ffffffffffffffff166107449190611d5f565b61074e9190611d5f565b6040518363ffffffff1660e01b815260040161076b929190611af5565b6020604051808303815f875af1158015610787573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906107ab9190611aca565b5060f173ffffffffffffffffffffffffffffffffffffffff166356c657a56107d283610db8565b836127106107e09190611b3c565b6175306107ed9190611b78565b67ffffffffffffffff16600180866108059190611b78565b60026108119190611ce2565b61081b9190611d2c565b8567ffffffffffffffff166108309190611d5f565b61083a9190611d5f565b6040518363ffffffff1660e01b8152600401610857929190611af5565b6020604051808303815f875af1158015610873573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906108979190611aca565b505b50565b60f173ffffffffffffffffffffffffffffffffffffffff166356c657a584846040518363ffffffff1660e01b81526004016108d8929190611af5565b6020604051808303815f875af11580156108f4573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906109189190611aca565b50801561095a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161095190611e10565b60405180910390fd5b505050565b5b6001156109745761096f610c7c565b610960565b565b5f8167ffffffffffffffff16409050919050565b60f173ffffffffffffffffffffffffffffffffffffffff16638c1370cd6040518163ffffffff1660e01b81526004016020604051808303815f875af11580156109d5573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906109f99190611aca565b5060f173ffffffffffffffffffffffffffffffffffffffff166356c657a583836040518363ffffffff1660e01b8152600401610a36929190611af5565b6020604051808303815f875af1158015610a52573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610a769190611aca565b505050565b7fa019c7431cdfd7ba63501ffa1ba7d8f2a028e447653a5af5a96077e5038e03398282604051610aac929190611e3d565b60405180910390a160078267ffffffffffffffff1603610b01576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610af890611e10565b60405180910390fd5b5050565b606060f173ffffffffffffffffffffffffffffffffffffffff1663cad2355484846040518363ffffffff1660e01b8152600401610b43929190611e64565b5f60405180830381865afa158015610b5d573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f82011682018060405250810190610b859190611f07565b905092915050565b60f173ffffffffffffffffffffffffffffffffffffffff166356c657a5610bb383610db8565b83612710610bc19190611b3c565b617530610bce9190611b78565b67ffffffffffffffff1660018086610be69190611b78565b6002610bf29190611ce2565b610bfc9190611d2c565b8567ffffffffffffffff16610c119190611d5f565b610c1b9190611d5f565b6040518363ffffffff1660e01b8152600401610c38929190611af5565b6020604051808303815f875af1158015610c54573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610c789190611aca565b5f80fd5b5f80815480929190610c8d9061199e565b91905055507f61996fe196f72cb598c483e896a1221263a28bb630480aa89495f737d4a8e3df60015f54610cc19190611d2c565b5f54604051610cd1929190611f4e565b60405180910390a1565b8273ffffffffffffffffffffffffffffffffffffffff16632f2770db6040518163ffffffff1660e01b81526004015f604051808303815f87803b158015610d20575f80fd5b505af1158015610d32573d5f803e3d5ffd5b5050505060f173ffffffffffffffffffffffffffffffffffffffff166356c657a583836040518363ffffffff1660e01b8152600401610d72929190611af5565b6020604051808303815f875af1158015610d8e573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610db29190611aca565b50505050565b606060f173ffffffffffffffffffffffffffffffffffffffff16636af32a55306040518263ffffffff1660e01b8152600401610df49190611f84565b5f60405180830381865afa158015610e0e573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f82011682018060405250810190610e369190611f07565b610e3f30610ec3565b610e9b636193689560e01b600186610e5791906119e5565b604051602001610e679190611b23565b604051602081830303815290604052604051602001610e8792919061202c565b604051602081830303815290604052610ef0565b604051602001610ead939291906122b9565b6040516020818303038152906040529050919050565b6060610ee98273ffffffffffffffffffffffffffffffffffffffff16601460ff16611174565b9050919050565b60605f6002808451610f029190611d5f565b610f0c919061196b565b67ffffffffffffffff811115610f2557610f246114e7565b5b6040519080825280601f01601f191660200182016040528015610f575781602001600182028036833780820191505090505b5090507f3000000000000000000000000000000000000000000000000000000000000000815f81518110610f8e57610f8d61234c565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a9053507f780000000000000000000000000000000000000000000000000000000000000081600181518110610ff157610ff061234c565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a9053505f5b835181101561116a575f84828151811061103e5761103d61234c565b5b602001015160f81c60f81b60f81c90507f303132333435363738396162636465660000000000000000000000000000000060048260ff16901c60ff166010811061108b5761108a61234c565b5b1a60f81b83600280850201815181106110a7576110a661234c565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a9053507f3031323334353637383961626364656600000000000000000000000000000000600f821660ff166010811061110e5761110d61234c565b5b1a60f81b83600260016002860201018151811061112e5761112d61234c565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a905350508080600101915050611021565b5080915050919050565b60605f8390505f600284600261118a9190611d5f565b611194919061196b565b67ffffffffffffffff8111156111ad576111ac6114e7565b5b6040519080825280601f01601f1916602001820160405280156111df5781602001600182028036833780820191505090505b5090507f3000000000000000000000000000000000000000000000000000000000000000815f815181106112165761121561234c565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a9053507f7800000000000000000000000000000000000000000000000000000000000000816001815181106112795761127861234c565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a9053505f60018560026112b79190611d5f565b6112c1919061196b565b90505b6001811115611360577f3031323334353637383961626364656600000000000000000000000000000000600f8416601081106113035761130261234c565b5b1a60f81b82828151811061131a5761131961234c565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a905350600483901c92508061135990612379565b90506112c4565b505f82146113a75784846040517fe22e27eb00000000000000000000000000000000000000000000000000000000815260040161139e929190611f4e565b60405180910390fd5b809250505092915050565b5f819050919050565b6113c4816113b2565b82525050565b5f6020820190506113dd5f8301846113bb565b92915050565b5f604051905090565b5f80fd5b5f80fd5b5f67ffffffffffffffff82169050919050565b611410816113f4565b811461141a575f80fd5b50565b5f8135905061142b81611407565b92915050565b5f8115159050919050565b61144581611431565b811461144f575f80fd5b50565b5f813590506114608161143c565b92915050565b5f806040838503121561147c5761147b6113ec565b5b5f6114898582860161141d565b925050602061149a85828601611452565b9150509250929050565b5f602082840312156114b9576114b86113ec565b5b5f6114c68482850161141d565b91505092915050565b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b61151d826114d7565b810181811067ffffffffffffffff8211171561153c5761153b6114e7565b5b80604052505050565b5f61154e6113e3565b905061155a8282611514565b919050565b5f67ffffffffffffffff821115611579576115786114e7565b5b611582826114d7565b9050602081019050919050565b828183375f83830152505050565b5f6115af6115aa8461155f565b611545565b9050828152602081018484840111156115cb576115ca6114d3565b5b6115d684828561158f565b509392505050565b5f82601f8301126115f2576115f16114cf565b5b813561160284826020860161159d565b91505092915050565b5f805f8060808587031215611623576116226113ec565b5b5f85013567ffffffffffffffff8111156116405761163f6113f0565b5b61164c878288016115de565b945050602061165d8782880161141d565b935050604061166e87828801611452565b925050606061167f8782880161141d565b91505092959194509250565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6116b48261168b565b9050919050565b6116c4816116aa565b81146116ce575f80fd5b50565b5f813590506116df816116bb565b92915050565b5f805f606084860312156116fc576116fb6113ec565b5b5f611709868287016116d1565b935050602084013567ffffffffffffffff81111561172a576117296113f0565b5b611736868287016115de565b92505060406117478682870161141d565b9150509250925092565b5f805f60608486031215611768576117676113ec565b5b5f84013567ffffffffffffffff811115611785576117846113f0565b5b611791868287016115de565b93505060206117a28682870161141d565b92505060406117b386828701611452565b9150509250925092565b5f819050919050565b6117cf816117bd565b82525050565b5f6020820190506117e85f8301846117c6565b92915050565b5f8060408385031215611804576118036113ec565b5b5f83013567ffffffffffffffff811115611821576118206113f0565b5b61182d858286016115de565b925050602061183e8582860161141d565b9150509250929050565b5f806040838503121561185e5761185d6113ec565b5b5f83013567ffffffffffffffff81111561187b5761187a6113f0565b5b611887858286016115de565b925050602083013567ffffffffffffffff8111156118a8576118a76113f0565b5b6118b4858286016115de565b9150509250929050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f6118f0826118be565b6118fa81856118c8565b935061190a8185602086016118d8565b611913816114d7565b840191505092915050565b5f6020820190508181035f83015261193681846118e6565b905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f611975826113b2565b9150611980836113b2565b92508282019050808211156119985761199761193e565b5b92915050565b5f6119a8826113b2565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036119da576119d961193e565b5b600182019050919050565b5f6119ef826113f4565b91506119fa836113f4565b9250828203905067ffffffffffffffff811115611a1a57611a1961193e565b5b92915050565b611a29816113f4565b82525050565b611a3881611431565b82525050565b611a47816113f4565b82525050565b604082015f820151611a615f850182611a2f565b506020820151611a746020850182611a3e565b50505050565b5f6080820190508181035f830152611a9281866118e6565b9050611aa16020830185611a20565b611aae6040830184611a4d565b949350505050565b5f81519050611ac48161143c565b92915050565b5f60208284031215611adf57611ade6113ec565b5b5f611aec84828501611ab6565b91505092915050565b5f6040820190508181035f830152611b0d81856118e6565b9050611b1c6020830184611a20565b9392505050565b5f602082019050611b365f830184611a20565b92915050565b5f611b46826113f4565b9150611b51836113f4565b9250828202611b5f816113f4565b9150808214611b7157611b7061193e565b5b5092915050565b5f611b82826113f4565b9150611b8d836113f4565b9250828201905067ffffffffffffffff811115611bad57611bac61193e565b5b92915050565b5f8160011c9050919050565b5f808291508390505b6001851115611c0857808604811115611be457611be361193e565b5b6001851615611bf35780820291505b8081029050611c0185611bb3565b9450611bc8565b94509492505050565b5f82611c205760019050611cdb565b81611c2d575f9050611cdb565b8160018114611c435760028114611c4d57611c7c565b6001915050611cdb565b60ff841115611c5f57611c5e61193e565b5b8360020a915084821115611c7657611c7561193e565b5b50611cdb565b5060208310610133831016604e8410600b8410161715611cb15782820a905083811115611cac57611cab61193e565b5b611cdb565b611cbe8484846001611bbf565b92509050818404811115611cd557611cd461193e565b5b81810290505b9392505050565b5f611cec826113b2565b9150611cf7836113f4565b9250611d247fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8484611c11565b905092915050565b5f611d36826113b2565b9150611d41836113b2565b9250828203905081811115611d5957611d5861193e565b5b92915050565b5f611d69826113b2565b9150611d74836113b2565b9250828202611d82816113b2565b91508282048414831517611d9957611d9861193e565b5b5092915050565b7f72657665727420726561736f6e2064756d6d792076616c756520666f722074655f8201527f7374000000000000000000000000000000000000000000000000000000000000602082015250565b5f611dfa6022836118c8565b9150611e0582611da0565b604082019050919050565b5f6020820190508181035f830152611e2781611dee565b9050919050565b611e3781611431565b82525050565b5f604082019050611e505f830185611a20565b611e5d6020830184611e2e565b9392505050565b5f6040820190508181035f830152611e7c81856118e6565b90508181036020830152611e9081846118e6565b90509392505050565b5f611eab611ea68461155f565b611545565b905082815260208101848484011115611ec757611ec66114d3565b5b611ed28482856118d8565b509392505050565b5f82601f830112611eee57611eed6114cf565b5b8151611efe848260208601611e99565b91505092915050565b5f60208284031215611f1c57611f1b6113ec565b5b5f82015167ffffffffffffffff811115611f3957611f386113f0565b5b611f4584828501611eda565b91505092915050565b5f604082019050611f615f8301856113bb565b611f6e60208301846113bb565b9392505050565b611f7e816116aa565b82525050565b5f602082019050611f975f830184611f75565b92915050565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b5f819050919050565b611fe2611fdd82611f9d565b611fc8565b82525050565b5f81519050919050565b5f81905092915050565b5f61200682611fe8565b6120108185611ff2565b93506120208185602086016118d8565b80840191505092915050565b5f6120378285611fd1565b6004820191506120478284611ffc565b91508190509392505050565b5f81905092915050565b7f7b224074797065223a20222f6d696e6965766d2e65766d2e76312e4d736743615f8201527f6c6c222c00000000000000000000000000000000000000000000000000000000602082015250565b5f6120b7602483612053565b91506120c28261205d565b602482019050919050565b7f2273656e646572223a20220000000000000000000000000000000000000000005f82015250565b5f612101600b83612053565b915061210c826120cd565b600b82019050919050565b5f612121826118be565b61212b8185612053565b935061213b8185602086016118d8565b80840191505092915050565b7f222c0000000000000000000000000000000000000000000000000000000000005f82015250565b5f61217b600283612053565b915061218682612147565b600282019050919050565b7f22636f6e74726163745f61646472223a202200000000000000000000000000005f82015250565b5f6121c5601283612053565b91506121d082612191565b601282019050919050565b7f22696e707574223a2022000000000000000000000000000000000000000000005f82015250565b5f61220f600a83612053565b915061221a826121db565b600a82019050919050565b7f2276616c7565223a202230222c000000000000000000000000000000000000005f82015250565b5f612259600d83612053565b915061226482612225565b600d82019050919050565b7f226163636573735f6c697374223a205b5d7d00000000000000000000000000005f82015250565b5f6122a3601283612053565b91506122ae8261226f565b601282019050919050565b5f6122c3826120ab565b91506122ce826120f5565b91506122da8286612117565b91506122e58261216f565b91506122f0826121b9565b91506122fc8285612117565b91506123078261216f565b915061231282612203565b915061231e8284612117565b91506123298261216f565b91506123348261224d565b915061233f82612297565b9150819050949350505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f612383826113b2565b91505f82036123955761239461193e565b5b60018203905091905056fea164736f6c6343000819000a",
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

// DisableAndExecute is a paid mutator transaction binding the contract method 0xbb78714d.
//
// Solidity: function disable_and_execute(string exec_msg, uint64 gas_limit) returns()
func (_Counter *CounterTransactor) DisableAndExecute(opts *bind.TransactOpts, exec_msg string, gas_limit uint64) (*types.Transaction, error) {
	return _Counter.contract.Transact(opts, "disable_and_execute", exec_msg, gas_limit)
}

// DisableAndExecute is a paid mutator transaction binding the contract method 0xbb78714d.
//
// Solidity: function disable_and_execute(string exec_msg, uint64 gas_limit) returns()
func (_Counter *CounterSession) DisableAndExecute(exec_msg string, gas_limit uint64) (*types.Transaction, error) {
	return _Counter.Contract.DisableAndExecute(&_Counter.TransactOpts, exec_msg, gas_limit)
}

// DisableAndExecute is a paid mutator transaction binding the contract method 0xbb78714d.
//
// Solidity: function disable_and_execute(string exec_msg, uint64 gas_limit) returns()
func (_Counter *CounterTransactorSession) DisableAndExecute(exec_msg string, gas_limit uint64) (*types.Transaction, error) {
	return _Counter.Contract.DisableAndExecute(&_Counter.TransactOpts, exec_msg, gas_limit)
}

// DisableAndExecuteInChild is a paid mutator transaction binding the contract method 0x53a38e8f.
//
// Solidity: function disable_and_execute_in_child(address test_addr, string exec_msg, uint64 gas_limit) returns()
func (_Counter *CounterTransactor) DisableAndExecuteInChild(opts *bind.TransactOpts, test_addr common.Address, exec_msg string, gas_limit uint64) (*types.Transaction, error) {
	return _Counter.contract.Transact(opts, "disable_and_execute_in_child", test_addr, exec_msg, gas_limit)
}

// DisableAndExecuteInChild is a paid mutator transaction binding the contract method 0x53a38e8f.
//
// Solidity: function disable_and_execute_in_child(address test_addr, string exec_msg, uint64 gas_limit) returns()
func (_Counter *CounterSession) DisableAndExecuteInChild(test_addr common.Address, exec_msg string, gas_limit uint64) (*types.Transaction, error) {
	return _Counter.Contract.DisableAndExecuteInChild(&_Counter.TransactOpts, test_addr, exec_msg, gas_limit)
}

// DisableAndExecuteInChild is a paid mutator transaction binding the contract method 0x53a38e8f.
//
// Solidity: function disable_and_execute_in_child(address test_addr, string exec_msg, uint64 gas_limit) returns()
func (_Counter *CounterTransactorSession) DisableAndExecuteInChild(test_addr common.Address, exec_msg string, gas_limit uint64) (*types.Transaction, error) {
	return _Counter.Contract.DisableAndExecuteInChild(&_Counter.TransactOpts, test_addr, exec_msg, gas_limit)
}

// DisableAndExecuteInParent is a paid mutator transaction binding the contract method 0xfbb2c5dd.
//
// Solidity: function disable_and_execute_in_parent(address test_addr, string exec_msg, uint64 gas_limit) returns()
func (_Counter *CounterTransactor) DisableAndExecuteInParent(opts *bind.TransactOpts, test_addr common.Address, exec_msg string, gas_limit uint64) (*types.Transaction, error) {
	return _Counter.contract.Transact(opts, "disable_and_execute_in_parent", test_addr, exec_msg, gas_limit)
}

// DisableAndExecuteInParent is a paid mutator transaction binding the contract method 0xfbb2c5dd.
//
// Solidity: function disable_and_execute_in_parent(address test_addr, string exec_msg, uint64 gas_limit) returns()
func (_Counter *CounterSession) DisableAndExecuteInParent(test_addr common.Address, exec_msg string, gas_limit uint64) (*types.Transaction, error) {
	return _Counter.Contract.DisableAndExecuteInParent(&_Counter.TransactOpts, test_addr, exec_msg, gas_limit)
}

// DisableAndExecuteInParent is a paid mutator transaction binding the contract method 0xfbb2c5dd.
//
// Solidity: function disable_and_execute_in_parent(address test_addr, string exec_msg, uint64 gas_limit) returns()
func (_Counter *CounterTransactorSession) DisableAndExecuteInParent(test_addr common.Address, exec_msg string, gas_limit uint64) (*types.Transaction, error) {
	return _Counter.Contract.DisableAndExecuteInParent(&_Counter.TransactOpts, test_addr, exec_msg, gas_limit)
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

// Loop is a paid mutator transaction binding the contract method 0xa92100cb.
//
// Solidity: function loop() returns()
func (_Counter *CounterTransactor) Loop(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Counter.contract.Transact(opts, "loop")
}

// Loop is a paid mutator transaction binding the contract method 0xa92100cb.
//
// Solidity: function loop() returns()
func (_Counter *CounterSession) Loop() (*types.Transaction, error) {
	return _Counter.Contract.Loop(&_Counter.TransactOpts)
}

// Loop is a paid mutator transaction binding the contract method 0xa92100cb.
//
// Solidity: function loop() returns()
func (_Counter *CounterTransactorSession) Loop() (*types.Transaction, error) {
	return _Counter.Contract.Loop(&_Counter.TransactOpts)
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
