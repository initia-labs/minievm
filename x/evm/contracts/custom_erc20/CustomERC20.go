// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package custom_erc20

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

// CustomErc20MetaData contains all meta data concerning the CustomErc20 contract.
var CustomErc20MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"_decimals\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"sudoTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801562000010575f80fd5b5060405162001f1b38038062001f1b8339818101604052810190620000369190620002ef565b335f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060f273ffffffffffffffffffffffffffffffffffffffff16635e6c57596040518163ffffffff1660e01b81526004016020604051808303815f875af1158015620000c1573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190620000e79190620003c0565b508260039081620000f9919062000627565b5081600490816200010b919062000627565b508060055f6101000a81548160ff021916908360ff1602179055505050506200070b565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b620001908262000148565b810181811067ffffffffffffffff82111715620001b257620001b162000158565b5b80604052505050565b5f620001c66200012f565b9050620001d4828262000185565b919050565b5f67ffffffffffffffff821115620001f657620001f562000158565b5b620002018262000148565b9050602081019050919050565b5f5b838110156200022d57808201518184015260208101905062000210565b5f8484015250505050565b5f6200024e6200024884620001d9565b620001bb565b9050828152602081018484840111156200026d576200026c62000144565b5b6200027a8482856200020e565b509392505050565b5f82601f83011262000299576200029862000140565b5b8151620002ab84826020860162000238565b91505092915050565b5f60ff82169050919050565b620002cb81620002b4565b8114620002d6575f80fd5b50565b5f81519050620002e981620002c0565b92915050565b5f805f6060848603121562000309576200030862000138565b5b5f84015167ffffffffffffffff8111156200032957620003286200013c565b5b620003378682870162000282565b935050602084015167ffffffffffffffff8111156200035b576200035a6200013c565b5b620003698682870162000282565b92505060406200037c86828701620002d9565b9150509250925092565b5f8115159050919050565b6200039c8162000386565b8114620003a7575f80fd5b50565b5f81519050620003ba8162000391565b92915050565b5f60208284031215620003d857620003d762000138565b5b5f620003e784828501620003aa565b91505092915050565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806200043f57607f821691505b602082108103620004555762000454620003fa565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302620004b97fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826200047c565b620004c586836200047c565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f6200050f620005096200050384620004dd565b620004e6565b620004dd565b9050919050565b5f819050919050565b6200052a83620004ef565b62000542620005398262000516565b84845462000488565b825550505050565b5f90565b620005586200054a565b620005658184846200051f565b505050565b5b818110156200058c57620005805f826200054e565b6001810190506200056b565b5050565b601f821115620005db57620005a5816200045b565b620005b0846200046d565b81016020851015620005c0578190505b620005d8620005cf856200046d565b8301826200056a565b50505b505050565b5f82821c905092915050565b5f620005fd5f1984600802620005e0565b1980831691505092915050565b5f620006178383620005ec565b9150826002028217905092915050565b6200063282620003f0565b67ffffffffffffffff8111156200064e576200064d62000158565b5b6200065a825462000427565b6200066782828562000590565b5f60209050601f8311600181146200069d575f841562000688578287015190505b6200069485826200060a565b86555062000703565b601f198416620006ad866200045b565b5f5b82811015620006d657848901518255600182019150602085019450602081019050620006af565b86831015620006f65784890151620006f2601f891682620005ec565b8355505b6001600288020188555050505b505050505050565b61180280620007195f395ff3fe608060405234801561000f575f80fd5b50600436106100f3575f3560e01c806340c10f19116100955780639dc29fac116100645780639dc29fac14610285578063a9059cbb146102a1578063dd62ed3e146102d1578063f2fde38b14610301576100f3565b806340c10f19146101fd57806370a08231146102195780638da5cb5b1461024957806395d89b4114610267576100f3565b806318160ddd116100d157806318160ddd146101755780631988513b1461019357806323b872dd146101af578063313ce567146101df576100f3565b806301ffc9a7146100f757806306fdde0314610127578063095ea7b314610145575b5f80fd5b610111600480360381019061010c91906111b1565b61031d565b60405161011e91906111f6565b60405180910390f35b61012f610396565b60405161013c9190611299565b60405180910390f35b61015f600480360381019061015a9190611346565b610422565b60405161016c91906111f6565b60405180910390f35b61017d61050f565b60405161018a9190611393565b60405180910390f35b6101ad60048036038101906101a891906113ac565b610515565b005b6101c960048036038101906101c491906113ac565b610594565b6040516101d691906111f6565b60405180910390f35b6101e76106f4565b6040516101f49190611417565b60405180910390f35b61021760048036038101906102129190611346565b610706565b005b610233600480360381019061022e9190611430565b610825565b6040516102409190611393565b60405180910390f35b61025161083a565b60405161025e919061146a565b60405180910390f35b61026f61085d565b60405161027c9190611299565b60405180910390f35b61029f600480360381019061029a9190611346565b6108e9565b005b6102bb60048036038101906102b69190611346565b610a08565b6040516102c891906111f6565b60405180910390f35b6102eb60048036038101906102e69190611483565b610ad9565b6040516102f89190611393565b60405180910390f35b61031b60048036038101906103169190611430565b610af9565b005b5f7f8da6da19000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916148061038f575061038e82610c41565b5b9050919050565b600380546103a3906114ee565b80601f01602080910402602001604051908101604052809291908181526020018280546103cf906114ee565b801561041a5780601f106103f15761010080835404028352916020019161041a565b820191905f5260205f20905b8154815290600101906020018083116103fd57829003601f168201915b505050505081565b5f8160025f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040516104fd9190611393565b60405180910390a36001905092915050565b60065481565b600173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610584576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161057b90611568565b60405180910390fd5b61058f838383610caa565b505050565b5f8260f173ffffffffffffffffffffffffffffffffffffffff1663f2af9ac9826040518263ffffffff1660e01b81526004016105d0919061146a565b602060405180830381865afa1580156105eb573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061060f91906115b0565b1561064f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106469061164b565b60405180910390fd5b8260025f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546106d69190611696565b925050819055506106e8858585610caa565b60019150509392505050565b60055f9054906101000a900460ff1681565b8160f173ffffffffffffffffffffffffffffffffffffffff1663f2af9ac9826040518263ffffffff1660e01b8152600401610741919061146a565b602060405180830381865afa15801561075c573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061078091906115b0565b156107c0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107b790611713565b60405180910390fd5b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610816575f80fd5b6108208383610eb5565b505050565b6001602052805f5260405f205f915090505481565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6004805461086a906114ee565b80601f0160208091040260200160405190810160405280929190818152602001828054610896906114ee565b80156108e15780601f106108b8576101008083540402835291602001916108e1565b820191905f5260205f20905b8154815290600101906020018083116108c457829003601f168201915b505050505081565b8160f173ffffffffffffffffffffffffffffffffffffffff166360dc402f826040518263ffffffff1660e01b8152600401610924919061146a565b602060405180830381865afa15801561093f573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061096391906115b0565b156109a3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161099a9061177b565b60405180910390fd5b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146109f9575f80fd5b610a038383611084565b505050565b5f8260f173ffffffffffffffffffffffffffffffffffffffff1663f2af9ac9826040518263ffffffff1660e01b8152600401610a44919061146a565b602060405180830381865afa158015610a5f573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610a8391906115b0565b15610ac3576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610aba9061164b565b60405180910390fd5b610ace338585610caa565b600191505092915050565b6002602052815f5260405f20602052805f5260405f205f91509150505481565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610b4f575f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610b86575f80fd5b8073ffffffffffffffffffffffffffffffffffffffff165f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3805f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b5f7f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b8160f273ffffffffffffffffffffffffffffffffffffffff16634e25ab64826040518263ffffffff1660e01b8152600401610ce5919061146a565b602060405180830381865afa158015610d00573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610d2491906115b0565b610da45760f273ffffffffffffffffffffffffffffffffffffffff1663ceeae52a826040518263ffffffff1660e01b8152600401610d62919061146a565b6020604051808303815f875af1158015610d7e573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610da291906115b0565b505b8160015f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254610df09190611696565b925050819055508160015f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254610e439190611799565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051610ea79190611393565b60405180910390a350505050565b8160f273ffffffffffffffffffffffffffffffffffffffff16634e25ab64826040518263ffffffff1660e01b8152600401610ef0919061146a565b602060405180830381865afa158015610f0b573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610f2f91906115b0565b610faf5760f273ffffffffffffffffffffffffffffffffffffffff1663ceeae52a826040518263ffffffff1660e01b8152600401610f6d919061146a565b6020604051808303815f875af1158015610f89573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610fad91906115b0565b505b8160015f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254610ffb9190611799565b925050819055508160065f8282546110139190611799565b925050819055508273ffffffffffffffffffffffffffffffffffffffff165f73ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516110779190611393565b60405180910390a3505050565b8060015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546110d09190611696565b925050819055508060065f8282546110e89190611696565b925050819055505f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8360405161114c9190611393565b60405180910390a35050565b5f80fd5b5f7fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b6111908161115c565b811461119a575f80fd5b50565b5f813590506111ab81611187565b92915050565b5f602082840312156111c6576111c5611158565b5b5f6111d38482850161119d565b91505092915050565b5f8115159050919050565b6111f0816111dc565b82525050565b5f6020820190506112095f8301846111e7565b92915050565b5f81519050919050565b5f82825260208201905092915050565b5f5b8381101561124657808201518184015260208101905061122b565b5f8484015250505050565b5f601f19601f8301169050919050565b5f61126b8261120f565b6112758185611219565b9350611285818560208601611229565b61128e81611251565b840191505092915050565b5f6020820190508181035f8301526112b18184611261565b905092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6112e2826112b9565b9050919050565b6112f2816112d8565b81146112fc575f80fd5b50565b5f8135905061130d816112e9565b92915050565b5f819050919050565b61132581611313565b811461132f575f80fd5b50565b5f813590506113408161131c565b92915050565b5f806040838503121561135c5761135b611158565b5b5f611369858286016112ff565b925050602061137a85828601611332565b9150509250929050565b61138d81611313565b82525050565b5f6020820190506113a65f830184611384565b92915050565b5f805f606084860312156113c3576113c2611158565b5b5f6113d0868287016112ff565b93505060206113e1868287016112ff565b92505060406113f286828701611332565b9150509250925092565b5f60ff82169050919050565b611411816113fc565b82525050565b5f60208201905061142a5f830184611408565b92915050565b5f6020828403121561144557611444611158565b5b5f611452848285016112ff565b91505092915050565b611464816112d8565b82525050565b5f60208201905061147d5f83018461145b565b92915050565b5f806040838503121561149957611498611158565b5b5f6114a6858286016112ff565b92505060206114b7858286016112ff565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061150557607f821691505b602082108103611518576115176114c1565b5b50919050565b7f45524332303a2063616c6c6572206973206e6f742074686520636861696e00005f82015250565b5f611552601e83611219565b915061155d8261151e565b602082019050919050565b5f6020820190508181035f83015261157f81611546565b9050919050565b61158f816111dc565b8114611599575f80fd5b50565b5f815190506115aa81611586565b92915050565b5f602082840312156115c5576115c4611158565b5b5f6115d28482850161159c565b91505092915050565b7f45524332303a207472616e7366657220746f20626c6f636b65642061646472655f8201527f7373000000000000000000000000000000000000000000000000000000000000602082015250565b5f611635602283611219565b9150611640826115db565b604082019050919050565b5f6020820190508181035f83015261166281611629565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6116a082611313565b91506116ab83611313565b92508282039050818111156116c3576116c2611669565b5b92915050565b7f45524332303a206d696e7420746f20626c6f636b6564206164647265737300005f82015250565b5f6116fd601e83611219565b9150611708826116c9565b602082019050919050565b5f6020820190508181035f83015261172a816116f1565b9050919050565b7f45524332303a206275726e2066726f6d206d6f64756c652061646472657373005f82015250565b5f611765601f83611219565b915061177082611731565b602082019050919050565b5f6020820190508181035f83015261179281611759565b9050919050565b5f6117a382611313565b91506117ae83611313565b92508282019050808211156117c6576117c5611669565b5b9291505056fea26469706673582212207cb456b1f272c3058457ae14b90cc057df7f6b452a44e5fce79e903c67692df164736f6c63430008180033",
}

// CustomErc20ABI is the input ABI used to generate the binding from.
// Deprecated: Use CustomErc20MetaData.ABI instead.
var CustomErc20ABI = CustomErc20MetaData.ABI

// CustomErc20Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CustomErc20MetaData.Bin instead.
var CustomErc20Bin = CustomErc20MetaData.Bin

// DeployCustomErc20 deploys a new Ethereum contract, binding an instance of CustomErc20 to it.
func DeployCustomErc20(auth *bind.TransactOpts, backend bind.ContractBackend, _name string, _symbol string, _decimals uint8) (common.Address, *types.Transaction, *CustomErc20, error) {
	parsed, err := CustomErc20MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CustomErc20Bin), backend, _name, _symbol, _decimals)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CustomErc20{CustomErc20Caller: CustomErc20Caller{contract: contract}, CustomErc20Transactor: CustomErc20Transactor{contract: contract}, CustomErc20Filterer: CustomErc20Filterer{contract: contract}}, nil
}

// CustomErc20 is an auto generated Go binding around an Ethereum contract.
type CustomErc20 struct {
	CustomErc20Caller     // Read-only binding to the contract
	CustomErc20Transactor // Write-only binding to the contract
	CustomErc20Filterer   // Log filterer for contract events
}

// CustomErc20Caller is an auto generated read-only Go binding around an Ethereum contract.
type CustomErc20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CustomErc20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type CustomErc20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CustomErc20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CustomErc20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CustomErc20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CustomErc20Session struct {
	Contract     *CustomErc20      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CustomErc20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CustomErc20CallerSession struct {
	Contract *CustomErc20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// CustomErc20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CustomErc20TransactorSession struct {
	Contract     *CustomErc20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// CustomErc20Raw is an auto generated low-level Go binding around an Ethereum contract.
type CustomErc20Raw struct {
	Contract *CustomErc20 // Generic contract binding to access the raw methods on
}

// CustomErc20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CustomErc20CallerRaw struct {
	Contract *CustomErc20Caller // Generic read-only contract binding to access the raw methods on
}

// CustomErc20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CustomErc20TransactorRaw struct {
	Contract *CustomErc20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewCustomErc20 creates a new instance of CustomErc20, bound to a specific deployed contract.
func NewCustomErc20(address common.Address, backend bind.ContractBackend) (*CustomErc20, error) {
	contract, err := bindCustomErc20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CustomErc20{CustomErc20Caller: CustomErc20Caller{contract: contract}, CustomErc20Transactor: CustomErc20Transactor{contract: contract}, CustomErc20Filterer: CustomErc20Filterer{contract: contract}}, nil
}

// NewCustomErc20Caller creates a new read-only instance of CustomErc20, bound to a specific deployed contract.
func NewCustomErc20Caller(address common.Address, caller bind.ContractCaller) (*CustomErc20Caller, error) {
	contract, err := bindCustomErc20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CustomErc20Caller{contract: contract}, nil
}

// NewCustomErc20Transactor creates a new write-only instance of CustomErc20, bound to a specific deployed contract.
func NewCustomErc20Transactor(address common.Address, transactor bind.ContractTransactor) (*CustomErc20Transactor, error) {
	contract, err := bindCustomErc20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CustomErc20Transactor{contract: contract}, nil
}

// NewCustomErc20Filterer creates a new log filterer instance of CustomErc20, bound to a specific deployed contract.
func NewCustomErc20Filterer(address common.Address, filterer bind.ContractFilterer) (*CustomErc20Filterer, error) {
	contract, err := bindCustomErc20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CustomErc20Filterer{contract: contract}, nil
}

// bindCustomErc20 binds a generic wrapper to an already deployed contract.
func bindCustomErc20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CustomErc20MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CustomErc20 *CustomErc20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CustomErc20.Contract.CustomErc20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CustomErc20 *CustomErc20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CustomErc20.Contract.CustomErc20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CustomErc20 *CustomErc20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CustomErc20.Contract.CustomErc20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CustomErc20 *CustomErc20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CustomErc20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CustomErc20 *CustomErc20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CustomErc20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CustomErc20 *CustomErc20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CustomErc20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_CustomErc20 *CustomErc20Caller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CustomErc20.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_CustomErc20 *CustomErc20Session) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _CustomErc20.Contract.Allowance(&_CustomErc20.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_CustomErc20 *CustomErc20CallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _CustomErc20.Contract.Allowance(&_CustomErc20.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_CustomErc20 *CustomErc20Caller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CustomErc20.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_CustomErc20 *CustomErc20Session) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _CustomErc20.Contract.BalanceOf(&_CustomErc20.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_CustomErc20 *CustomErc20CallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _CustomErc20.Contract.BalanceOf(&_CustomErc20.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_CustomErc20 *CustomErc20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _CustomErc20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_CustomErc20 *CustomErc20Session) Decimals() (uint8, error) {
	return _CustomErc20.Contract.Decimals(&_CustomErc20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_CustomErc20 *CustomErc20CallerSession) Decimals() (uint8, error) {
	return _CustomErc20.Contract.Decimals(&_CustomErc20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CustomErc20 *CustomErc20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CustomErc20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CustomErc20 *CustomErc20Session) Name() (string, error) {
	return _CustomErc20.Contract.Name(&_CustomErc20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CustomErc20 *CustomErc20CallerSession) Name() (string, error) {
	return _CustomErc20.Contract.Name(&_CustomErc20.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CustomErc20 *CustomErc20Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CustomErc20.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CustomErc20 *CustomErc20Session) Owner() (common.Address, error) {
	return _CustomErc20.Contract.Owner(&_CustomErc20.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CustomErc20 *CustomErc20CallerSession) Owner() (common.Address, error) {
	return _CustomErc20.Contract.Owner(&_CustomErc20.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CustomErc20 *CustomErc20Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _CustomErc20.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CustomErc20 *CustomErc20Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CustomErc20.Contract.SupportsInterface(&_CustomErc20.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CustomErc20 *CustomErc20CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CustomErc20.Contract.SupportsInterface(&_CustomErc20.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CustomErc20 *CustomErc20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CustomErc20.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CustomErc20 *CustomErc20Session) Symbol() (string, error) {
	return _CustomErc20.Contract.Symbol(&_CustomErc20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CustomErc20 *CustomErc20CallerSession) Symbol() (string, error) {
	return _CustomErc20.Contract.Symbol(&_CustomErc20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CustomErc20 *CustomErc20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CustomErc20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CustomErc20 *CustomErc20Session) TotalSupply() (*big.Int, error) {
	return _CustomErc20.Contract.TotalSupply(&_CustomErc20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CustomErc20 *CustomErc20CallerSession) TotalSupply() (*big.Int, error) {
	return _CustomErc20.Contract.TotalSupply(&_CustomErc20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_CustomErc20 *CustomErc20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CustomErc20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_CustomErc20 *CustomErc20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CustomErc20.Contract.Approve(&_CustomErc20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_CustomErc20 *CustomErc20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CustomErc20.Contract.Approve(&_CustomErc20.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address from, uint256 amount) returns()
func (_CustomErc20 *CustomErc20Transactor) Burn(opts *bind.TransactOpts, from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CustomErc20.contract.Transact(opts, "burn", from, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address from, uint256 amount) returns()
func (_CustomErc20 *CustomErc20Session) Burn(from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CustomErc20.Contract.Burn(&_CustomErc20.TransactOpts, from, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address from, uint256 amount) returns()
func (_CustomErc20 *CustomErc20TransactorSession) Burn(from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CustomErc20.Contract.Burn(&_CustomErc20.TransactOpts, from, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_CustomErc20 *CustomErc20Transactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CustomErc20.contract.Transact(opts, "mint", to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_CustomErc20 *CustomErc20Session) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CustomErc20.Contract.Mint(&_CustomErc20.TransactOpts, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_CustomErc20 *CustomErc20TransactorSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CustomErc20.Contract.Mint(&_CustomErc20.TransactOpts, to, amount)
}

// SudoTransfer is a paid mutator transaction binding the contract method 0x1988513b.
//
// Solidity: function sudoTransfer(address sender, address recipient, uint256 amount) returns()
func (_CustomErc20 *CustomErc20Transactor) SudoTransfer(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CustomErc20.contract.Transact(opts, "sudoTransfer", sender, recipient, amount)
}

// SudoTransfer is a paid mutator transaction binding the contract method 0x1988513b.
//
// Solidity: function sudoTransfer(address sender, address recipient, uint256 amount) returns()
func (_CustomErc20 *CustomErc20Session) SudoTransfer(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CustomErc20.Contract.SudoTransfer(&_CustomErc20.TransactOpts, sender, recipient, amount)
}

// SudoTransfer is a paid mutator transaction binding the contract method 0x1988513b.
//
// Solidity: function sudoTransfer(address sender, address recipient, uint256 amount) returns()
func (_CustomErc20 *CustomErc20TransactorSession) SudoTransfer(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CustomErc20.Contract.SudoTransfer(&_CustomErc20.TransactOpts, sender, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_CustomErc20 *CustomErc20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CustomErc20.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_CustomErc20 *CustomErc20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CustomErc20.Contract.Transfer(&_CustomErc20.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_CustomErc20 *CustomErc20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CustomErc20.Contract.Transfer(&_CustomErc20.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_CustomErc20 *CustomErc20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CustomErc20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_CustomErc20 *CustomErc20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CustomErc20.Contract.TransferFrom(&_CustomErc20.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_CustomErc20 *CustomErc20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CustomErc20.Contract.TransferFrom(&_CustomErc20.TransactOpts, sender, recipient, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CustomErc20 *CustomErc20Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _CustomErc20.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CustomErc20 *CustomErc20Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CustomErc20.Contract.TransferOwnership(&_CustomErc20.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CustomErc20 *CustomErc20TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CustomErc20.Contract.TransferOwnership(&_CustomErc20.TransactOpts, newOwner)
}

// CustomErc20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the CustomErc20 contract.
type CustomErc20ApprovalIterator struct {
	Event *CustomErc20Approval // Event containing the contract specifics and raw log

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
func (it *CustomErc20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CustomErc20Approval)
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
		it.Event = new(CustomErc20Approval)
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
func (it *CustomErc20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CustomErc20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CustomErc20Approval represents a Approval event raised by the CustomErc20 contract.
type CustomErc20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_CustomErc20 *CustomErc20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*CustomErc20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _CustomErc20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &CustomErc20ApprovalIterator{contract: _CustomErc20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_CustomErc20 *CustomErc20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *CustomErc20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _CustomErc20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CustomErc20Approval)
				if err := _CustomErc20.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_CustomErc20 *CustomErc20Filterer) ParseApproval(log types.Log) (*CustomErc20Approval, error) {
	event := new(CustomErc20Approval)
	if err := _CustomErc20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CustomErc20OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the CustomErc20 contract.
type CustomErc20OwnershipTransferredIterator struct {
	Event *CustomErc20OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CustomErc20OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CustomErc20OwnershipTransferred)
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
		it.Event = new(CustomErc20OwnershipTransferred)
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
func (it *CustomErc20OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CustomErc20OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CustomErc20OwnershipTransferred represents a OwnershipTransferred event raised by the CustomErc20 contract.
type CustomErc20OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CustomErc20 *CustomErc20Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CustomErc20OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CustomErc20.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CustomErc20OwnershipTransferredIterator{contract: _CustomErc20.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CustomErc20 *CustomErc20Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CustomErc20OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CustomErc20.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CustomErc20OwnershipTransferred)
				if err := _CustomErc20.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_CustomErc20 *CustomErc20Filterer) ParseOwnershipTransferred(log types.Log) (*CustomErc20OwnershipTransferred, error) {
	event := new(CustomErc20OwnershipTransferred)
	if err := _CustomErc20.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CustomErc20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the CustomErc20 contract.
type CustomErc20TransferIterator struct {
	Event *CustomErc20Transfer // Event containing the contract specifics and raw log

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
func (it *CustomErc20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CustomErc20Transfer)
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
		it.Event = new(CustomErc20Transfer)
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
func (it *CustomErc20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CustomErc20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CustomErc20Transfer represents a Transfer event raised by the CustomErc20 contract.
type CustomErc20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_CustomErc20 *CustomErc20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CustomErc20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CustomErc20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CustomErc20TransferIterator{contract: _CustomErc20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_CustomErc20 *CustomErc20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *CustomErc20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CustomErc20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CustomErc20Transfer)
				if err := _CustomErc20.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_CustomErc20 *CustomErc20Filterer) ParseTransfer(log types.Log) (*CustomErc20Transfer, error) {
	event := new(CustomErc20Transfer)
	if err := _CustomErc20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
