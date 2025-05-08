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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"StringsInsufficientHexLength\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"callback_id\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"name\":\"callback_received\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newCount\",\"type\":\"uint256\"}],\"name\":\"increased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"n\",\"type\":\"uint64\"}],\"name\":\"recursive_called\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"callback_id\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"name\":\"callback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"count\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"exec_msg\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"gas_limit\",\"type\":\"uint64\"}],\"name\":\"disable_and_execute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"test_addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"exec_msg\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"gas_limit\",\"type\":\"uint64\"}],\"name\":\"disable_and_execute_in_child\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"test_addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"exec_msg\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"gas_limit\",\"type\":\"uint64\"}],\"name\":\"disable_and_execute_in_parent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"exec_msg\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"gas_limit\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"call_revert\",\"type\":\"bool\"}],\"name\":\"execute_cosmos\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"exec_msg\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"gas_limit\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"allow_failure\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"callback_id\",\"type\":\"uint64\"}],\"name\":\"execute_cosmos_with_options\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"n\",\"type\":\"uint64\"}],\"name\":\"get_blockhash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"callback_id\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"name\":\"ibc_ack\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"callback_id\",\"type\":\"uint64\"}],\"name\":\"ibc_timeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"increase\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"num\",\"type\":\"uint64\"}],\"name\":\"increase_for_fuzz\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"n\",\"type\":\"uint64\"}],\"name\":\"nested_recursive_revert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"path\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"req\",\"type\":\"string\"}],\"name\":\"query_cosmos\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"result\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"n\",\"type\":\"uint64\"}],\"name\":\"recursive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"n\",\"type\":\"uint64\"}],\"name\":\"recursive_revert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052612375806100115f395ff3fe6080604052600436106100f2575f3560e01c80637aa4d9cc11610089578063cad2355411610058578063cad23554146102ec578063df3f725014610328578063e8927fbc14610350578063fbb2c5dd1461035a576100f2565b80637aa4d9cc14610238578063ac7fde5f14610260578063bb78714d1461029c578063c31925a7146102c4576100f2565b806352dadc5a116100c557806352dadc5a1461019857806353a38e8f146101c05780635b133d02146101e85780636193689514610210576100f2565b806306661abd146100f65780630d4f1f9d146101205780632607baf81461014857806331a503f014610170575b5f80fd5b348015610101575f80fd5b5061010a610382565b6040516101179190611392565b60405180910390f35b34801561012b575f80fd5b506101466004803603810190610141919061142e565b610387565b005b348015610153575f80fd5b5061016e6004803603810190610169919061146c565b6103ce565b005b34801561017b575f80fd5b506101966004803603810190610191919061146c565b610401565b005b3480156101a3575f80fd5b506101be60048036038101906101b991906115d3565b610425565b005b3480156101cb575f80fd5b506101e660048036038101906101e191906116ad565b6104ca565b005b3480156101f3575f80fd5b5061020e6004803603810190610209919061146c565b6105a7565b005b34801561021b575f80fd5b506102366004803603810190610231919061146c565b610656565b005b348015610243575f80fd5b5061025e60048036038101906102599190611719565b61087b565b005b34801561026b575f80fd5b506102866004803603810190610281919061146c565b61093e565b604051610293919061179d565b60405180910390f35b3480156102a7575f80fd5b506102c260048036038101906102bd91906117b6565b610952565b005b3480156102cf575f80fd5b506102ea60048036038101906102e5919061142e565b610a43565b005b3480156102f7575f80fd5b50610312600480360381019061030d9190611810565b610acd565b60405161031f91906118e6565b60405180910390f35b348015610333575f80fd5b5061034e6004803603810190610349919061146c565b610b55565b005b610358610c44565b005b348015610365575f80fd5b50610380600480360381019061037b91906116ad565b610ca3565b005b5f5481565b80156103b3578167ffffffffffffffff165f808282546103a79190611933565b925050819055506103ca565b5f808154809291906103c490611966565b91905055505b5050565b5f8167ffffffffffffffff1603156103fe576103e8610c44565b6103fd6001826103f891906119ad565b6103ce565b5b50565b8067ffffffffffffffff165f8082825461041b9190611933565b9250508190555050565b60f173ffffffffffffffffffffffffffffffffffffffff1663f1ed795d8585604051806040016040528087151581526020018667ffffffffffffffff168152506040518463ffffffff1660e01b815260040161048393929190611a42565b6020604051808303815f875af115801561049f573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906104c39190611a92565b5050505050565b60f173ffffffffffffffffffffffffffffffffffffffff16638c1370cd6040518163ffffffff1660e01b81526004016020604051808303815f875af1158015610515573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906105399190611a92565b508273ffffffffffffffffffffffffffffffffffffffff166356c657a583836040518363ffffffff1660e01b8152600401610575929190611abd565b5f604051808303815f87803b15801561058c575f80fd5b505af115801561059e573d5f803e3d5ffd5b50505050505050565b7f4475bbd12ef452e28d39c4cb77494d85136c2d89ca1354b52188d4aaa8f4ba51816040516105d69190611aeb565b60405180910390a15f8167ffffffffffffffff160315610653573073ffffffffffffffffffffffffffffffffffffffff1663df3f7250826040518263ffffffff1660e01b81526004016106299190611aeb565b5f604051808303815f87803b158015610640575f80fd5b505af1925050508015610651575060015b505b50565b7f4475bbd12ef452e28d39c4cb77494d85136c2d89ca1354b52188d4aaa8f4ba51816040516106859190611aeb565b60405180910390a15f8167ffffffffffffffff1603156108785760f173ffffffffffffffffffffffffffffffffffffffff166356c657a56106c583610d80565b836127106106d39190611b04565b6175306106e09190611b40565b67ffffffffffffffff16600180866106f89190611b40565b60026107049190611caa565b61070e9190611cf4565b8567ffffffffffffffff166107239190611d27565b61072d9190611d27565b6040518363ffffffff1660e01b815260040161074a929190611abd565b6020604051808303815f875af1158015610766573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061078a9190611a92565b5060f173ffffffffffffffffffffffffffffffffffffffff166356c657a56107b183610d80565b836127106107bf9190611b04565b6175306107cc9190611b40565b67ffffffffffffffff16600180866107e49190611b40565b60026107f09190611caa565b6107fa9190611cf4565b8567ffffffffffffffff1661080f9190611d27565b6108199190611d27565b6040518363ffffffff1660e01b8152600401610836929190611abd565b6020604051808303815f875af1158015610852573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906108769190611a92565b505b50565b60f173ffffffffffffffffffffffffffffffffffffffff166356c657a584846040518363ffffffff1660e01b81526004016108b7929190611abd565b6020604051808303815f875af11580156108d3573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906108f79190611a92565b508015610939576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161093090611dd8565b60405180910390fd5b505050565b5f8167ffffffffffffffff16409050919050565b60f173ffffffffffffffffffffffffffffffffffffffff16638c1370cd6040518163ffffffff1660e01b81526004016020604051808303815f875af115801561099d573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906109c19190611a92565b5060f173ffffffffffffffffffffffffffffffffffffffff166356c657a583836040518363ffffffff1660e01b81526004016109fe929190611abd565b6020604051808303815f875af1158015610a1a573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610a3e9190611a92565b505050565b7fa019c7431cdfd7ba63501ffa1ba7d8f2a028e447653a5af5a96077e5038e03398282604051610a74929190611e05565b60405180910390a160078267ffffffffffffffff1603610ac9576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ac090611dd8565b60405180910390fd5b5050565b606060f173ffffffffffffffffffffffffffffffffffffffff1663cad2355484846040518363ffffffff1660e01b8152600401610b0b929190611e2c565b5f60405180830381865afa158015610b25573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f82011682018060405250810190610b4d9190611ecf565b905092915050565b60f173ffffffffffffffffffffffffffffffffffffffff166356c657a5610b7b83610d80565b83612710610b899190611b04565b617530610b969190611b40565b67ffffffffffffffff1660018086610bae9190611b40565b6002610bba9190611caa565b610bc49190611cf4565b8567ffffffffffffffff16610bd99190611d27565b610be39190611d27565b6040518363ffffffff1660e01b8152600401610c00929190611abd565b6020604051808303815f875af1158015610c1c573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610c409190611a92565b5f80fd5b5f80815480929190610c5590611966565b91905055507f61996fe196f72cb598c483e896a1221263a28bb630480aa89495f737d4a8e3df60015f54610c899190611cf4565b5f54604051610c99929190611f16565b60405180910390a1565b8273ffffffffffffffffffffffffffffffffffffffff16632f2770db6040518163ffffffff1660e01b81526004015f604051808303815f87803b158015610ce8575f80fd5b505af1158015610cfa573d5f803e3d5ffd5b5050505060f173ffffffffffffffffffffffffffffffffffffffff166356c657a583836040518363ffffffff1660e01b8152600401610d3a929190611abd565b6020604051808303815f875af1158015610d56573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610d7a9190611a92565b50505050565b606060f173ffffffffffffffffffffffffffffffffffffffff16636af32a55306040518263ffffffff1660e01b8152600401610dbc9190611f4c565b5f60405180830381865afa158015610dd6573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f82011682018060405250810190610dfe9190611ecf565b610e0730610e8b565b610e63636193689560e01b600186610e1f91906119ad565b604051602001610e2f9190611aeb565b604051602081830303815290604052604051602001610e4f929190611ff4565b604051602081830303815290604052610eb8565b604051602001610e7593929190612281565b6040516020818303038152906040529050919050565b6060610eb18273ffffffffffffffffffffffffffffffffffffffff16601460ff1661113c565b9050919050565b60605f6002808451610eca9190611d27565b610ed49190611933565b67ffffffffffffffff811115610eed57610eec6114af565b5b6040519080825280601f01601f191660200182016040528015610f1f5781602001600182028036833780820191505090505b5090507f3000000000000000000000000000000000000000000000000000000000000000815f81518110610f5657610f55612314565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a9053507f780000000000000000000000000000000000000000000000000000000000000081600181518110610fb957610fb8612314565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a9053505f5b8351811015611132575f84828151811061100657611005612314565b5b602001015160f81c60f81b60f81c90507f303132333435363738396162636465660000000000000000000000000000000060048260ff16901c60ff166010811061105357611052612314565b5b1a60f81b836002808502018151811061106f5761106e612314565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a9053507f3031323334353637383961626364656600000000000000000000000000000000600f821660ff16601081106110d6576110d5612314565b5b1a60f81b8360026001600286020101815181106110f6576110f5612314565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a905350508080600101915050610fe9565b5080915050919050565b60605f8390505f60028460026111529190611d27565b61115c9190611933565b67ffffffffffffffff811115611175576111746114af565b5b6040519080825280601f01601f1916602001820160405280156111a75781602001600182028036833780820191505090505b5090507f3000000000000000000000000000000000000000000000000000000000000000815f815181106111de576111dd612314565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a9053507f78000000000000000000000000000000000000000000000000000000000000008160018151811061124157611240612314565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a9053505f600185600261127f9190611d27565b6112899190611933565b90505b6001811115611328577f3031323334353637383961626364656600000000000000000000000000000000600f8416601081106112cb576112ca612314565b5b1a60f81b8282815181106112e2576112e1612314565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a905350600483901c92508061132190612341565b905061128c565b505f821461136f5784846040517fe22e27eb000000000000000000000000000000000000000000000000000000008152600401611366929190611f16565b60405180910390fd5b809250505092915050565b5f819050919050565b61138c8161137a565b82525050565b5f6020820190506113a55f830184611383565b92915050565b5f604051905090565b5f80fd5b5f80fd5b5f67ffffffffffffffff82169050919050565b6113d8816113bc565b81146113e2575f80fd5b50565b5f813590506113f3816113cf565b92915050565b5f8115159050919050565b61140d816113f9565b8114611417575f80fd5b50565b5f8135905061142881611404565b92915050565b5f8060408385031215611444576114436113b4565b5b5f611451858286016113e5565b92505060206114628582860161141a565b9150509250929050565b5f60208284031215611481576114806113b4565b5b5f61148e848285016113e5565b91505092915050565b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6114e58261149f565b810181811067ffffffffffffffff82111715611504576115036114af565b5b80604052505050565b5f6115166113ab565b905061152282826114dc565b919050565b5f67ffffffffffffffff821115611541576115406114af565b5b61154a8261149f565b9050602081019050919050565b828183375f83830152505050565b5f61157761157284611527565b61150d565b9050828152602081018484840111156115935761159261149b565b5b61159e848285611557565b509392505050565b5f82601f8301126115ba576115b9611497565b5b81356115ca848260208601611565565b91505092915050565b5f805f80608085870312156115eb576115ea6113b4565b5b5f85013567ffffffffffffffff811115611608576116076113b8565b5b611614878288016115a6565b9450506020611625878288016113e5565b93505060406116368782880161141a565b9250506060611647878288016113e5565b91505092959194509250565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61167c82611653565b9050919050565b61168c81611672565b8114611696575f80fd5b50565b5f813590506116a781611683565b92915050565b5f805f606084860312156116c4576116c36113b4565b5b5f6116d186828701611699565b935050602084013567ffffffffffffffff8111156116f2576116f16113b8565b5b6116fe868287016115a6565b925050604061170f868287016113e5565b9150509250925092565b5f805f606084860312156117305761172f6113b4565b5b5f84013567ffffffffffffffff81111561174d5761174c6113b8565b5b611759868287016115a6565b935050602061176a868287016113e5565b925050604061177b8682870161141a565b9150509250925092565b5f819050919050565b61179781611785565b82525050565b5f6020820190506117b05f83018461178e565b92915050565b5f80604083850312156117cc576117cb6113b4565b5b5f83013567ffffffffffffffff8111156117e9576117e86113b8565b5b6117f5858286016115a6565b9250506020611806858286016113e5565b9150509250929050565b5f8060408385031215611826576118256113b4565b5b5f83013567ffffffffffffffff811115611843576118426113b8565b5b61184f858286016115a6565b925050602083013567ffffffffffffffff8111156118705761186f6113b8565b5b61187c858286016115a6565b9150509250929050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f6118b882611886565b6118c28185611890565b93506118d28185602086016118a0565b6118db8161149f565b840191505092915050565b5f6020820190508181035f8301526118fe81846118ae565b905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61193d8261137a565b91506119488361137a565b92508282019050808211156119605761195f611906565b5b92915050565b5f6119708261137a565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036119a2576119a1611906565b5b600182019050919050565b5f6119b7826113bc565b91506119c2836113bc565b9250828203905067ffffffffffffffff8111156119e2576119e1611906565b5b92915050565b6119f1816113bc565b82525050565b611a00816113f9565b82525050565b611a0f816113bc565b82525050565b604082015f820151611a295f8501826119f7565b506020820151611a3c6020850182611a06565b50505050565b5f6080820190508181035f830152611a5a81866118ae565b9050611a6960208301856119e8565b611a766040830184611a15565b949350505050565b5f81519050611a8c81611404565b92915050565b5f60208284031215611aa757611aa66113b4565b5b5f611ab484828501611a7e565b91505092915050565b5f6040820190508181035f830152611ad581856118ae565b9050611ae460208301846119e8565b9392505050565b5f602082019050611afe5f8301846119e8565b92915050565b5f611b0e826113bc565b9150611b19836113bc565b9250828202611b27816113bc565b9150808214611b3957611b38611906565b5b5092915050565b5f611b4a826113bc565b9150611b55836113bc565b9250828201905067ffffffffffffffff811115611b7557611b74611906565b5b92915050565b5f8160011c9050919050565b5f808291508390505b6001851115611bd057808604811115611bac57611bab611906565b5b6001851615611bbb5780820291505b8081029050611bc985611b7b565b9450611b90565b94509492505050565b5f82611be85760019050611ca3565b81611bf5575f9050611ca3565b8160018114611c0b5760028114611c1557611c44565b6001915050611ca3565b60ff841115611c2757611c26611906565b5b8360020a915084821115611c3e57611c3d611906565b5b50611ca3565b5060208310610133831016604e8410600b8410161715611c795782820a905083811115611c7457611c73611906565b5b611ca3565b611c868484846001611b87565b92509050818404811115611c9d57611c9c611906565b5b81810290505b9392505050565b5f611cb48261137a565b9150611cbf836113bc565b9250611cec7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8484611bd9565b905092915050565b5f611cfe8261137a565b9150611d098361137a565b9250828203905081811115611d2157611d20611906565b5b92915050565b5f611d318261137a565b9150611d3c8361137a565b9250828202611d4a8161137a565b91508282048414831517611d6157611d60611906565b5b5092915050565b7f72657665727420726561736f6e2064756d6d792076616c756520666f722074655f8201527f7374000000000000000000000000000000000000000000000000000000000000602082015250565b5f611dc2602283611890565b9150611dcd82611d68565b604082019050919050565b5f6020820190508181035f830152611def81611db6565b9050919050565b611dff816113f9565b82525050565b5f604082019050611e185f8301856119e8565b611e256020830184611df6565b9392505050565b5f6040820190508181035f830152611e4481856118ae565b90508181036020830152611e5881846118ae565b90509392505050565b5f611e73611e6e84611527565b61150d565b905082815260208101848484011115611e8f57611e8e61149b565b5b611e9a8482856118a0565b509392505050565b5f82601f830112611eb657611eb5611497565b5b8151611ec6848260208601611e61565b91505092915050565b5f60208284031215611ee457611ee36113b4565b5b5f82015167ffffffffffffffff811115611f0157611f006113b8565b5b611f0d84828501611ea2565b91505092915050565b5f604082019050611f295f830185611383565b611f366020830184611383565b9392505050565b611f4681611672565b82525050565b5f602082019050611f5f5f830184611f3d565b92915050565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b5f819050919050565b611faa611fa582611f65565b611f90565b82525050565b5f81519050919050565b5f81905092915050565b5f611fce82611fb0565b611fd88185611fba565b9350611fe88185602086016118a0565b80840191505092915050565b5f611fff8285611f99565b60048201915061200f8284611fc4565b91508190509392505050565b5f81905092915050565b7f7b224074797065223a20222f6d696e6965766d2e65766d2e76312e4d736743615f8201527f6c6c222c00000000000000000000000000000000000000000000000000000000602082015250565b5f61207f60248361201b565b915061208a82612025565b602482019050919050565b7f2273656e646572223a20220000000000000000000000000000000000000000005f82015250565b5f6120c9600b8361201b565b91506120d482612095565b600b82019050919050565b5f6120e982611886565b6120f3818561201b565b93506121038185602086016118a0565b80840191505092915050565b7f222c0000000000000000000000000000000000000000000000000000000000005f82015250565b5f61214360028361201b565b915061214e8261210f565b600282019050919050565b7f22636f6e74726163745f61646472223a202200000000000000000000000000005f82015250565b5f61218d60128361201b565b915061219882612159565b601282019050919050565b7f22696e707574223a2022000000000000000000000000000000000000000000005f82015250565b5f6121d7600a8361201b565b91506121e2826121a3565b600a82019050919050565b7f2276616c7565223a202230222c000000000000000000000000000000000000005f82015250565b5f612221600d8361201b565b915061222c826121ed565b600d82019050919050565b7f226163636573735f6c697374223a205b5d7d00000000000000000000000000005f82015250565b5f61226b60128361201b565b915061227682612237565b601282019050919050565b5f61228b82612073565b9150612296826120bd565b91506122a282866120df565b91506122ad82612137565b91506122b882612181565b91506122c482856120df565b91506122cf82612137565b91506122da826121cb565b91506122e682846120df565b91506122f182612137565b91506122fc82612215565b91506123078261225f565b9150819050949350505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f61234b8261137a565b91505f820361235d5761235c611906565b5b60018203905091905056fea164736f6c6343000819000a",
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
