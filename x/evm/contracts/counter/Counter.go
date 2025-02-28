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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"StringsInsufficientHexLength\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"callback_id\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"name\":\"callback_received\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newCount\",\"type\":\"uint256\"}],\"name\":\"increased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"n\",\"type\":\"uint64\"}],\"name\":\"recursive_called\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"callback_id\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"name\":\"callback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"count\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"exec_msg\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"gas_limit\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"call_revert\",\"type\":\"bool\"}],\"name\":\"execute_cosmos\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"exec_msg\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"gas_limit\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"allow_failure\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"callback_id\",\"type\":\"uint64\"}],\"name\":\"execute_cosmos_with_options\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"n\",\"type\":\"uint64\"}],\"name\":\"get_blockhash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"callback_id\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"name\":\"ibc_ack\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"callback_id\",\"type\":\"uint64\"}],\"name\":\"ibc_timeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"increase\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"num\",\"type\":\"uint64\"}],\"name\":\"increase_for_fuzz\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"n\",\"type\":\"uint64\"}],\"name\":\"nested_recursive_revert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"path\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"req\",\"type\":\"string\"}],\"name\":\"query_cosmos\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"result\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"n\",\"type\":\"uint64\"}],\"name\":\"recursive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"n\",\"type\":\"uint64\"}],\"name\":\"recursive_revert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052611ee6806100115f395ff3fe6080604052600436106100c1575f3560e01c8063619368951161007e578063c31925a711610058578063c31925a714610243578063cad235541461026b578063df3f7250146102a7578063e8927fbc146102cf576100c1565b806361936895146101b75780637aa4d9cc146101df578063ac7fde5f14610207576100c1565b806306661abd146100c55780630d4f1f9d146100ef5780632607baf81461011757806331a503f01461013f57806352dadc5a146101675780635b133d021461018f575b5f80fd5b3480156100d0575f80fd5b506100d96102d9565b6040516100e69190610ff3565b60405180910390f35b3480156100fa575f80fd5b506101156004803603810190610110919061108f565b6102de565b005b348015610122575f80fd5b5061013d600480360381019061013891906110cd565b610325565b005b34801561014a575f80fd5b50610165600480360381019061016091906110cd565b610358565b005b348015610172575f80fd5b5061018d60048036038101906101889190611234565b61037c565b005b34801561019a575f80fd5b506101b560048036038101906101b091906110cd565b610421565b005b3480156101c2575f80fd5b506101dd60048036038101906101d891906110cd565b6104d0565b005b3480156101ea575f80fd5b50610205600480360381019061020091906112b4565b6106f5565b005b348015610212575f80fd5b5061022d600480360381019061022891906110cd565b6107b8565b60405161023a9190611338565b60405180910390f35b34801561024e575f80fd5b506102696004803603810190610264919061108f565b6107cc565b005b348015610276575f80fd5b50610291600480360381019061028c9190611351565b610809565b60405161029e9190611427565b60405180910390f35b3480156102b2575f80fd5b506102cd60048036038101906102c891906110cd565b610892565b005b6102d7610981565b005b5f5481565b801561030a578167ffffffffffffffff165f808282546102fe9190611474565b92505081905550610321565b5f8081548092919061031b906114a7565b91905055505b5050565b5f8167ffffffffffffffff1603156103555761033f610981565b61035460018261034f91906114ee565b610325565b5b50565b8067ffffffffffffffff165f808282546103729190611474565b9250508190555050565b60f173ffffffffffffffffffffffffffffffffffffffff1663f1ed795d8585604051806040016040528087151581526020018667ffffffffffffffff168152506040518463ffffffff1660e01b81526004016103da93929190611583565b6020604051808303815f875af11580156103f6573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061041a91906115d3565b5050505050565b7f4475bbd12ef452e28d39c4cb77494d85136c2d89ca1354b52188d4aaa8f4ba518160405161045091906115fe565b60405180910390a15f8167ffffffffffffffff1603156104cd573073ffffffffffffffffffffffffffffffffffffffff1663df3f7250826040518263ffffffff1660e01b81526004016104a391906115fe565b5f604051808303815f87803b1580156104ba575f80fd5b505af19250505080156104cb575060015b505b50565b7f4475bbd12ef452e28d39c4cb77494d85136c2d89ca1354b52188d4aaa8f4ba51816040516104ff91906115fe565b60405180910390a15f8167ffffffffffffffff1603156106f25760f173ffffffffffffffffffffffffffffffffffffffff166356c657a561053f836109e0565b8361271061054d9190611617565b61753061055a9190611653565b67ffffffffffffffff16600180866105729190611653565b600261057e91906117bd565b6105889190611807565b8567ffffffffffffffff1661059d919061183a565b6105a7919061183a565b6040518363ffffffff1660e01b81526004016105c492919061187b565b6020604051808303815f875af11580156105e0573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061060491906115d3565b5060f173ffffffffffffffffffffffffffffffffffffffff166356c657a561062b836109e0565b836127106106399190611617565b6175306106469190611653565b67ffffffffffffffff166001808661065e9190611653565b600261066a91906117bd565b6106749190611807565b8567ffffffffffffffff16610689919061183a565b610693919061183a565b6040518363ffffffff1660e01b81526004016106b092919061187b565b6020604051808303815f875af11580156106cc573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906106f091906115d3565b505b50565b60f173ffffffffffffffffffffffffffffffffffffffff166356c657a584846040518363ffffffff1660e01b815260040161073192919061187b565b6020604051808303815f875af115801561074d573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061077191906115d3565b5080156107b3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107aa90611919565b60405180910390fd5b505050565b5f8167ffffffffffffffff16409050919050565b7fa019c7431cdfd7ba63501ffa1ba7d8f2a028e447653a5af5a96077e5038e033982826040516107fd929190611946565b60405180910390a15050565b606060f173ffffffffffffffffffffffffffffffffffffffff1663cad2355484846040518363ffffffff1660e01b815260040161084792919061196d565b5f604051808303815f875af1158015610862573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f8201168201806040525081019061088a9190611a10565b905092915050565b60f173ffffffffffffffffffffffffffffffffffffffff166356c657a56108b8836109e0565b836127106108c69190611617565b6175306108d39190611653565b67ffffffffffffffff16600180866108eb9190611653565b60026108f791906117bd565b6109019190611807565b8567ffffffffffffffff16610916919061183a565b610920919061183a565b6040518363ffffffff1660e01b815260040161093d92919061187b565b6020604051808303815f875af1158015610959573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061097d91906115d3565b5f80fd5b5f80815480929190610992906114a7565b91905055507f61996fe196f72cb598c483e896a1221263a28bb630480aa89495f737d4a8e3df60015f546109c69190611807565b5f546040516109d6929190611a57565b60405180910390a1565b606060f173ffffffffffffffffffffffffffffffffffffffff16636af32a55306040518263ffffffff1660e01b8152600401610a1c9190611abd565b5f604051808303815f875af1158015610a37573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f82011682018060405250810190610a5f9190611a10565b610a6830610aec565b610ac4636193689560e01b600186610a8091906114ee565b604051602001610a9091906115fe565b604051602081830303815290604052604051602001610ab0929190611b65565b604051602081830303815290604052610b19565b604051602001610ad693929190611df2565b6040516020818303038152906040529050919050565b6060610b128273ffffffffffffffffffffffffffffffffffffffff16601460ff16610d9d565b9050919050565b60605f6002808451610b2b919061183a565b610b359190611474565b67ffffffffffffffff811115610b4e57610b4d611110565b5b6040519080825280601f01601f191660200182016040528015610b805781602001600182028036833780820191505090505b5090507f3000000000000000000000000000000000000000000000000000000000000000815f81518110610bb757610bb6611e85565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a9053507f780000000000000000000000000000000000000000000000000000000000000081600181518110610c1a57610c19611e85565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a9053505f5b8351811015610d93575f848281518110610c6757610c66611e85565b5b602001015160f81c60f81b60f81c90507f303132333435363738396162636465660000000000000000000000000000000060048260ff16901c60ff1660108110610cb457610cb3611e85565b5b1a60f81b8360028085020181518110610cd057610ccf611e85565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a9053507f3031323334353637383961626364656600000000000000000000000000000000600f821660ff1660108110610d3757610d36611e85565b5b1a60f81b836002600160028602010181518110610d5757610d56611e85565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a905350508080600101915050610c4a565b5080915050919050565b60605f8390505f6002846002610db3919061183a565b610dbd9190611474565b67ffffffffffffffff811115610dd657610dd5611110565b5b6040519080825280601f01601f191660200182016040528015610e085781602001600182028036833780820191505090505b5090507f3000000000000000000000000000000000000000000000000000000000000000815f81518110610e3f57610e3e611e85565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a9053507f780000000000000000000000000000000000000000000000000000000000000081600181518110610ea257610ea1611e85565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a9053505f6001856002610ee0919061183a565b610eea9190611474565b90505b6001811115610f89577f3031323334353637383961626364656600000000000000000000000000000000600f841660108110610f2c57610f2b611e85565b5b1a60f81b828281518110610f4357610f42611e85565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a905350600483901c925080610f8290611eb2565b9050610eed565b505f8214610fd05784846040517fe22e27eb000000000000000000000000000000000000000000000000000000008152600401610fc7929190611a57565b60405180910390fd5b809250505092915050565b5f819050919050565b610fed81610fdb565b82525050565b5f6020820190506110065f830184610fe4565b92915050565b5f604051905090565b5f80fd5b5f80fd5b5f67ffffffffffffffff82169050919050565b6110398161101d565b8114611043575f80fd5b50565b5f8135905061105481611030565b92915050565b5f8115159050919050565b61106e8161105a565b8114611078575f80fd5b50565b5f8135905061108981611065565b92915050565b5f80604083850312156110a5576110a4611015565b5b5f6110b285828601611046565b92505060206110c38582860161107b565b9150509250929050565b5f602082840312156110e2576110e1611015565b5b5f6110ef84828501611046565b91505092915050565b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b61114682611100565b810181811067ffffffffffffffff8211171561116557611164611110565b5b80604052505050565b5f61117761100c565b9050611183828261113d565b919050565b5f67ffffffffffffffff8211156111a2576111a1611110565b5b6111ab82611100565b9050602081019050919050565b828183375f83830152505050565b5f6111d86111d384611188565b61116e565b9050828152602081018484840111156111f4576111f36110fc565b5b6111ff8482856111b8565b509392505050565b5f82601f83011261121b5761121a6110f8565b5b813561122b8482602086016111c6565b91505092915050565b5f805f806080858703121561124c5761124b611015565b5b5f85013567ffffffffffffffff81111561126957611268611019565b5b61127587828801611207565b945050602061128687828801611046565b93505060406112978782880161107b565b92505060606112a887828801611046565b91505092959194509250565b5f805f606084860312156112cb576112ca611015565b5b5f84013567ffffffffffffffff8111156112e8576112e7611019565b5b6112f486828701611207565b935050602061130586828701611046565b92505060406113168682870161107b565b9150509250925092565b5f819050919050565b61133281611320565b82525050565b5f60208201905061134b5f830184611329565b92915050565b5f806040838503121561136757611366611015565b5b5f83013567ffffffffffffffff81111561138457611383611019565b5b61139085828601611207565b925050602083013567ffffffffffffffff8111156113b1576113b0611019565b5b6113bd85828601611207565b9150509250929050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f6113f9826113c7565b61140381856113d1565b93506114138185602086016113e1565b61141c81611100565b840191505092915050565b5f6020820190508181035f83015261143f81846113ef565b905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61147e82610fdb565b915061148983610fdb565b92508282019050808211156114a1576114a0611447565b5b92915050565b5f6114b182610fdb565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036114e3576114e2611447565b5b600182019050919050565b5f6114f88261101d565b91506115038361101d565b9250828203905067ffffffffffffffff81111561152357611522611447565b5b92915050565b6115328161101d565b82525050565b6115418161105a565b82525050565b6115508161101d565b82525050565b604082015f82015161156a5f850182611538565b50602082015161157d6020850182611547565b50505050565b5f6080820190508181035f83015261159b81866113ef565b90506115aa6020830185611529565b6115b76040830184611556565b949350505050565b5f815190506115cd81611065565b92915050565b5f602082840312156115e8576115e7611015565b5b5f6115f5848285016115bf565b91505092915050565b5f6020820190506116115f830184611529565b92915050565b5f6116218261101d565b915061162c8361101d565b925082820261163a8161101d565b915080821461164c5761164b611447565b5b5092915050565b5f61165d8261101d565b91506116688361101d565b9250828201905067ffffffffffffffff81111561168857611687611447565b5b92915050565b5f8160011c9050919050565b5f808291508390505b60018511156116e3578086048111156116bf576116be611447565b5b60018516156116ce5780820291505b80810290506116dc8561168e565b94506116a3565b94509492505050565b5f826116fb57600190506117b6565b81611708575f90506117b6565b816001811461171e576002811461172857611757565b60019150506117b6565b60ff84111561173a57611739611447565b5b8360020a91508482111561175157611750611447565b5b506117b6565b5060208310610133831016604e8410600b841016171561178c5782820a90508381111561178757611786611447565b5b6117b6565b611799848484600161169a565b925090508184048111156117b0576117af611447565b5b81810290505b9392505050565b5f6117c782610fdb565b91506117d28361101d565b92506117ff7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84846116ec565b905092915050565b5f61181182610fdb565b915061181c83610fdb565b925082820390508181111561183457611833611447565b5b92915050565b5f61184482610fdb565b915061184f83610fdb565b925082820261185d81610fdb565b9150828204841483151761187457611873611447565b5b5092915050565b5f6040820190508181035f83015261189381856113ef565b90506118a26020830184611529565b9392505050565b7f72657665727420726561736f6e2064756d6d792076616c756520666f722074655f8201527f7374000000000000000000000000000000000000000000000000000000000000602082015250565b5f6119036022836113d1565b915061190e826118a9565b604082019050919050565b5f6020820190508181035f830152611930816118f7565b9050919050565b6119408161105a565b82525050565b5f6040820190506119595f830185611529565b6119666020830184611937565b9392505050565b5f6040820190508181035f83015261198581856113ef565b9050818103602083015261199981846113ef565b90509392505050565b5f6119b46119af84611188565b61116e565b9050828152602081018484840111156119d0576119cf6110fc565b5b6119db8482856113e1565b509392505050565b5f82601f8301126119f7576119f66110f8565b5b8151611a078482602086016119a2565b91505092915050565b5f60208284031215611a2557611a24611015565b5b5f82015167ffffffffffffffff811115611a4257611a41611019565b5b611a4e848285016119e3565b91505092915050565b5f604082019050611a6a5f830185610fe4565b611a776020830184610fe4565b9392505050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f611aa782611a7e565b9050919050565b611ab781611a9d565b82525050565b5f602082019050611ad05f830184611aae565b92915050565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b5f819050919050565b611b1b611b1682611ad6565b611b01565b82525050565b5f81519050919050565b5f81905092915050565b5f611b3f82611b21565b611b498185611b2b565b9350611b598185602086016113e1565b80840191505092915050565b5f611b708285611b0a565b600482019150611b808284611b35565b91508190509392505050565b5f81905092915050565b7f7b224074797065223a20222f6d696e6965766d2e65766d2e76312e4d736743615f8201527f6c6c222c00000000000000000000000000000000000000000000000000000000602082015250565b5f611bf0602483611b8c565b9150611bfb82611b96565b602482019050919050565b7f2273656e646572223a20220000000000000000000000000000000000000000005f82015250565b5f611c3a600b83611b8c565b9150611c4582611c06565b600b82019050919050565b5f611c5a826113c7565b611c648185611b8c565b9350611c748185602086016113e1565b80840191505092915050565b7f222c0000000000000000000000000000000000000000000000000000000000005f82015250565b5f611cb4600283611b8c565b9150611cbf82611c80565b600282019050919050565b7f22636f6e74726163745f61646472223a202200000000000000000000000000005f82015250565b5f611cfe601283611b8c565b9150611d0982611cca565b601282019050919050565b7f22696e707574223a2022000000000000000000000000000000000000000000005f82015250565b5f611d48600a83611b8c565b9150611d5382611d14565b600a82019050919050565b7f2276616c7565223a202230222c000000000000000000000000000000000000005f82015250565b5f611d92600d83611b8c565b9150611d9d82611d5e565b600d82019050919050565b7f226163636573735f6c697374223a205b5d7d00000000000000000000000000005f82015250565b5f611ddc601283611b8c565b9150611de782611da8565b601282019050919050565b5f611dfc82611be4565b9150611e0782611c2e565b9150611e138286611c50565b9150611e1e82611ca8565b9150611e2982611cf2565b9150611e358285611c50565b9150611e4082611ca8565b9150611e4b82611d3c565b9150611e578284611c50565b9150611e6282611ca8565b9150611e6d82611d86565b9150611e7882611dd0565b9150819050949350505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f611ebc82610fdb565b91505f8203611ece57611ecd611447565b5b60018203905091905056fea164736f6c6343000819000a",
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

// QueryCosmos is a paid mutator transaction binding the contract method 0xcad23554.
//
// Solidity: function query_cosmos(string path, string req) returns(string result)
func (_Counter *CounterTransactor) QueryCosmos(opts *bind.TransactOpts, path string, req string) (*types.Transaction, error) {
	return _Counter.contract.Transact(opts, "query_cosmos", path, req)
}

// QueryCosmos is a paid mutator transaction binding the contract method 0xcad23554.
//
// Solidity: function query_cosmos(string path, string req) returns(string result)
func (_Counter *CounterSession) QueryCosmos(path string, req string) (*types.Transaction, error) {
	return _Counter.Contract.QueryCosmos(&_Counter.TransactOpts, path, req)
}

// QueryCosmos is a paid mutator transaction binding the contract method 0xcad23554.
//
// Solidity: function query_cosmos(string path, string req) returns(string result)
func (_Counter *CounterTransactorSession) QueryCosmos(path string, req string) (*types.Transaction, error) {
	return _Counter.Contract.QueryCosmos(&_Counter.TransactOpts, path, req)
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
