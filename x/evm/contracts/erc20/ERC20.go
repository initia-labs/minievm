// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package erc20

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

// Erc20MetaData contains all meta data concerning the Erc20 contract.
var Erc20MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"_decimals\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"sudoBurn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"sudoMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"sudoTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801562000010575f80fd5b5060405162001f8638038062001f8683398181016040528101906200003691906200027c565b335f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555082600390816200008691906200054a565b5081600490816200009891906200054a565b508060055f6101000a81548160ff021916908360ff1602179055505050506200062e565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6200011d82620000d5565b810181811067ffffffffffffffff821117156200013f576200013e620000e5565b5b80604052505050565b5f62000153620000bc565b905062000161828262000112565b919050565b5f67ffffffffffffffff821115620001835762000182620000e5565b5b6200018e82620000d5565b9050602081019050919050565b5f5b83811015620001ba5780820151818401526020810190506200019d565b5f8484015250505050565b5f620001db620001d58462000166565b62000148565b905082815260208101848484011115620001fa57620001f9620000d1565b5b620002078482856200019b565b509392505050565b5f82601f830112620002265762000225620000cd565b5b815162000238848260208601620001c5565b91505092915050565b5f60ff82169050919050565b620002588162000241565b811462000263575f80fd5b50565b5f8151905062000276816200024d565b92915050565b5f805f60608486031215620002965762000295620000c5565b5b5f84015167ffffffffffffffff811115620002b657620002b5620000c9565b5b620002c4868287016200020f565b935050602084015167ffffffffffffffff811115620002e857620002e7620000c9565b5b620002f6868287016200020f565b9250506040620003098682870162000266565b9150509250925092565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806200036257607f821691505b6020821081036200037857620003776200031d565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302620003dc7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826200039f565b620003e886836200039f565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f620004326200042c620004268462000400565b62000409565b62000400565b9050919050565b5f819050919050565b6200044d8362000412565b620004656200045c8262000439565b848454620003ab565b825550505050565b5f90565b6200047b6200046d565b6200048881848462000442565b505050565b5b81811015620004af57620004a35f8262000471565b6001810190506200048e565b5050565b601f821115620004fe57620004c8816200037e565b620004d38462000390565b81016020851015620004e3578190505b620004fb620004f28562000390565b8301826200048d565b50505b505050565b5f82821c905092915050565b5f620005205f198460080262000503565b1980831691505092915050565b5f6200053a83836200050f565b9150826002028217905092915050565b620005558262000313565b67ffffffffffffffff811115620005715762000570620000e5565b5b6200057d82546200034a565b6200058a828285620004b3565b5f60209050601f831160018114620005c0575f8415620005ab578287015190505b620005b785826200052d565b86555062000626565b601f198416620005d0866200037e565b5f5b82811015620005f957848901518255600182019150602085019450602081019050620005d2565b8683101562000619578489015162000615601f8916826200050f565b8355505b6001600288020188555050505b505050505050565b61194a806200063c5f395ff3fe608060405234801561000f575f80fd5b5060043610610109575f3560e01c806340c10f19116100a05780639dc29fac1161006f5780639dc29fac146102b7578063a9059cbb146102d3578063dd62ed3e14610303578063f2fde38b14610333578063fe1195ec1461034f57610109565b806340c10f191461022f57806370a082311461024b5780638da5cb5b1461027b57806395d89b411461029957610109565b80631988513b116100dc5780631988513b146101a957806323b872dd146101c55780632d688ca8146101f5578063313ce5671461021157610109565b806301ffc9a71461010d57806306fdde031461013d578063095ea7b31461015b57806318160ddd1461018b575b5f80fd5b610127600480360381019061012291906112f9565b61036b565b604051610134919061133e565b60405180910390f35b6101456103e4565b60405161015291906113e1565b60405180910390f35b6101756004803603810190610170919061148e565b610470565b604051610182919061133e565b60405180910390f35b61019361055d565b6040516101a091906114db565b60405180910390f35b6101c360048036038101906101be91906114f4565b610563565b005b6101df60048036038101906101da91906114f4565b6105e2565b6040516101ec919061133e565b60405180910390f35b61020f600480360381019061020a919061148e565b610742565b005b6102196107bf565b604051610226919061155f565b60405180910390f35b6102496004803603810190610244919061148e565b6107d1565b005b61026560048036038101906102609190611578565b6108f0565b60405161027291906114db565b60405180910390f35b610283610905565b60405161029091906115b2565b60405180910390f35b6102a1610928565b6040516102ae91906113e1565b60405180910390f35b6102d160048036038101906102cc919061148e565b6109b4565b005b6102ed60048036038101906102e8919061148e565b610ad3565b6040516102fa919061133e565b60405180910390f35b61031d600480360381019061031891906115cb565b610ba4565b60405161032a91906114db565b60405180910390f35b61034d60048036038101906103489190611578565b610bc4565b005b6103696004803603810190610364919061148e565b610d0c565b005b5f7f8da6da19000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614806103dd57506103dc82610d89565b5b9050919050565b600380546103f190611636565b80601f016020809104026020016040519081016040528092919081815260200182805461041d90611636565b80156104685780601f1061043f57610100808354040283529160200191610468565b820191905f5260205f20905b81548152906001019060200180831161044b57829003601f168201915b505050505081565b5f8160025f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9258460405161054b91906114db565b60405180910390a36001905092915050565b60065481565b600173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146105d2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105c9906116b0565b60405180910390fd5b6105dd838383610df2565b505050565b5f8260f173ffffffffffffffffffffffffffffffffffffffff1663f2af9ac9826040518263ffffffff1660e01b815260040161061e91906115b2565b602060405180830381865afa158015610639573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061065d91906116f8565b1561069d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161069490611793565b60405180910390fd5b8260025f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825461072491906117de565b92505081905550610736858585610df2565b60019150509392505050565b600173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146107b1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107a8906116b0565b60405180910390fd5b6107bb8282610ffd565b5050565b60055f9054906101000a900460ff1681565b8160f173ffffffffffffffffffffffffffffffffffffffff1663f2af9ac9826040518263ffffffff1660e01b815260040161080c91906115b2565b602060405180830381865afa158015610827573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061084b91906116f8565b1561088b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108829061185b565b60405180910390fd5b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146108e1575f80fd5b6108eb8383610ffd565b505050565b6001602052805f5260405f205f915090505481565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6004805461093590611636565b80601f016020809104026020016040519081016040528092919081815260200182805461096190611636565b80156109ac5780601f10610983576101008083540402835291602001916109ac565b820191905f5260205f20905b81548152906001019060200180831161098f57829003601f168201915b505050505081565b8160f173ffffffffffffffffffffffffffffffffffffffff166360dc402f826040518263ffffffff1660e01b81526004016109ef91906115b2565b602060405180830381865afa158015610a0a573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610a2e91906116f8565b15610a6e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a65906118c3565b60405180910390fd5b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610ac4575f80fd5b610ace83836111cc565b505050565b5f8260f173ffffffffffffffffffffffffffffffffffffffff1663f2af9ac9826040518263ffffffff1660e01b8152600401610b0f91906115b2565b602060405180830381865afa158015610b2a573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610b4e91906116f8565b15610b8e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b8590611793565b60405180910390fd5b610b99338585610df2565b600191505092915050565b6002602052815f5260405f20602052805f5260405f205f91509150505481565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610c1a575f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610c51575f80fd5b8073ffffffffffffffffffffffffffffffffffffffff165f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3805f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610d7b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d72906116b0565b60405180910390fd5b610d8582826111cc565b5050565b5f7f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b8160f273ffffffffffffffffffffffffffffffffffffffff16634e25ab64826040518263ffffffff1660e01b8152600401610e2d91906115b2565b602060405180830381865afa158015610e48573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610e6c91906116f8565b610eec5760f273ffffffffffffffffffffffffffffffffffffffff1663ceeae52a826040518263ffffffff1660e01b8152600401610eaa91906115b2565b6020604051808303815f875af1158015610ec6573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610eea91906116f8565b505b8160015f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254610f3891906117de565b925050819055508160015f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254610f8b91906118e1565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051610fef91906114db565b60405180910390a350505050565b8160f273ffffffffffffffffffffffffffffffffffffffff16634e25ab64826040518263ffffffff1660e01b815260040161103891906115b2565b602060405180830381865afa158015611053573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061107791906116f8565b6110f75760f273ffffffffffffffffffffffffffffffffffffffff1663ceeae52a826040518263ffffffff1660e01b81526004016110b591906115b2565b6020604051808303815f875af11580156110d1573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906110f591906116f8565b505b8160015f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825461114391906118e1565b925050819055508160065f82825461115b91906118e1565b925050819055508273ffffffffffffffffffffffffffffffffffffffff165f73ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516111bf91906114db565b60405180910390a3505050565b8060015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825461121891906117de565b925050819055508060065f82825461123091906117de565b925050819055505f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8360405161129491906114db565b60405180910390a35050565b5f80fd5b5f7fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b6112d8816112a4565b81146112e2575f80fd5b50565b5f813590506112f3816112cf565b92915050565b5f6020828403121561130e5761130d6112a0565b5b5f61131b848285016112e5565b91505092915050565b5f8115159050919050565b61133881611324565b82525050565b5f6020820190506113515f83018461132f565b92915050565b5f81519050919050565b5f82825260208201905092915050565b5f5b8381101561138e578082015181840152602081019050611373565b5f8484015250505050565b5f601f19601f8301169050919050565b5f6113b382611357565b6113bd8185611361565b93506113cd818560208601611371565b6113d681611399565b840191505092915050565b5f6020820190508181035f8301526113f981846113a9565b905092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61142a82611401565b9050919050565b61143a81611420565b8114611444575f80fd5b50565b5f8135905061145581611431565b92915050565b5f819050919050565b61146d8161145b565b8114611477575f80fd5b50565b5f8135905061148881611464565b92915050565b5f80604083850312156114a4576114a36112a0565b5b5f6114b185828601611447565b92505060206114c28582860161147a565b9150509250929050565b6114d58161145b565b82525050565b5f6020820190506114ee5f8301846114cc565b92915050565b5f805f6060848603121561150b5761150a6112a0565b5b5f61151886828701611447565b935050602061152986828701611447565b925050604061153a8682870161147a565b9150509250925092565b5f60ff82169050919050565b61155981611544565b82525050565b5f6020820190506115725f830184611550565b92915050565b5f6020828403121561158d5761158c6112a0565b5b5f61159a84828501611447565b91505092915050565b6115ac81611420565b82525050565b5f6020820190506115c55f8301846115a3565b92915050565b5f80604083850312156115e1576115e06112a0565b5b5f6115ee85828601611447565b92505060206115ff85828601611447565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061164d57607f821691505b6020821081036116605761165f611609565b5b50919050565b7f45524332303a2063616c6c6572206973206e6f742074686520636861696e00005f82015250565b5f61169a601e83611361565b91506116a582611666565b602082019050919050565b5f6020820190508181035f8301526116c78161168e565b9050919050565b6116d781611324565b81146116e1575f80fd5b50565b5f815190506116f2816116ce565b92915050565b5f6020828403121561170d5761170c6112a0565b5b5f61171a848285016116e4565b91505092915050565b7f45524332303a207472616e7366657220746f20626c6f636b65642061646472655f8201527f7373000000000000000000000000000000000000000000000000000000000000602082015250565b5f61177d602283611361565b915061178882611723565b604082019050919050565b5f6020820190508181035f8301526117aa81611771565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6117e88261145b565b91506117f38361145b565b925082820390508181111561180b5761180a6117b1565b5b92915050565b7f45524332303a206d696e7420746f20626c6f636b6564206164647265737300005f82015250565b5f611845601e83611361565b915061185082611811565b602082019050919050565b5f6020820190508181035f83015261187281611839565b9050919050565b7f45524332303a206275726e2066726f6d206d6f64756c652061646472657373005f82015250565b5f6118ad601f83611361565b91506118b882611879565b602082019050919050565b5f6020820190508181035f8301526118da816118a1565b9050919050565b5f6118eb8261145b565b91506118f68361145b565b925082820190508082111561190e5761190d6117b1565b5b9291505056fea264697066735822122092a15d9b61c6a6de22122c9c81d375bc99e3984e299e75efa4ffe4653b8ae00d64736f6c63430008180033",
}

// Erc20ABI is the input ABI used to generate the binding from.
// Deprecated: Use Erc20MetaData.ABI instead.
var Erc20ABI = Erc20MetaData.ABI

// Erc20Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Erc20MetaData.Bin instead.
var Erc20Bin = Erc20MetaData.Bin

// DeployErc20 deploys a new Ethereum contract, binding an instance of Erc20 to it.
func DeployErc20(auth *bind.TransactOpts, backend bind.ContractBackend, _name string, _symbol string, _decimals uint8) (common.Address, *types.Transaction, *Erc20, error) {
	parsed, err := Erc20MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Erc20Bin), backend, _name, _symbol, _decimals)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Erc20{Erc20Caller: Erc20Caller{contract: contract}, Erc20Transactor: Erc20Transactor{contract: contract}, Erc20Filterer: Erc20Filterer{contract: contract}}, nil
}

// Erc20 is an auto generated Go binding around an Ethereum contract.
type Erc20 struct {
	Erc20Caller     // Read-only binding to the contract
	Erc20Transactor // Write-only binding to the contract
	Erc20Filterer   // Log filterer for contract events
}

// Erc20Caller is an auto generated read-only Go binding around an Ethereum contract.
type Erc20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Erc20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Erc20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Erc20Session struct {
	Contract     *Erc20            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Erc20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Erc20CallerSession struct {
	Contract *Erc20Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// Erc20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Erc20TransactorSession struct {
	Contract     *Erc20Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Erc20Raw is an auto generated low-level Go binding around an Ethereum contract.
type Erc20Raw struct {
	Contract *Erc20 // Generic contract binding to access the raw methods on
}

// Erc20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Erc20CallerRaw struct {
	Contract *Erc20Caller // Generic read-only contract binding to access the raw methods on
}

// Erc20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Erc20TransactorRaw struct {
	Contract *Erc20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewErc20 creates a new instance of Erc20, bound to a specific deployed contract.
func NewErc20(address common.Address, backend bind.ContractBackend) (*Erc20, error) {
	contract, err := bindErc20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Erc20{Erc20Caller: Erc20Caller{contract: contract}, Erc20Transactor: Erc20Transactor{contract: contract}, Erc20Filterer: Erc20Filterer{contract: contract}}, nil
}

// NewErc20Caller creates a new read-only instance of Erc20, bound to a specific deployed contract.
func NewErc20Caller(address common.Address, caller bind.ContractCaller) (*Erc20Caller, error) {
	contract, err := bindErc20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Erc20Caller{contract: contract}, nil
}

// NewErc20Transactor creates a new write-only instance of Erc20, bound to a specific deployed contract.
func NewErc20Transactor(address common.Address, transactor bind.ContractTransactor) (*Erc20Transactor, error) {
	contract, err := bindErc20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Erc20Transactor{contract: contract}, nil
}

// NewErc20Filterer creates a new log filterer instance of Erc20, bound to a specific deployed contract.
func NewErc20Filterer(address common.Address, filterer bind.ContractFilterer) (*Erc20Filterer, error) {
	contract, err := bindErc20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Erc20Filterer{contract: contract}, nil
}

// bindErc20 binds a generic wrapper to an already deployed contract.
func bindErc20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Erc20MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc20 *Erc20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc20.Contract.Erc20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc20 *Erc20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc20.Contract.Erc20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc20 *Erc20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc20.Contract.Erc20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc20 *Erc20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc20 *Erc20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc20 *Erc20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Erc20 *Erc20Caller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Erc20 *Erc20Session) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Erc20.Contract.Allowance(&_Erc20.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Erc20 *Erc20CallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Erc20.Contract.Allowance(&_Erc20.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Erc20 *Erc20Caller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Erc20 *Erc20Session) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Erc20.Contract.BalanceOf(&_Erc20.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Erc20 *Erc20CallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Erc20.Contract.BalanceOf(&_Erc20.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Erc20 *Erc20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Erc20 *Erc20Session) Decimals() (uint8, error) {
	return _Erc20.Contract.Decimals(&_Erc20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Erc20 *Erc20CallerSession) Decimals() (uint8, error) {
	return _Erc20.Contract.Decimals(&_Erc20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Erc20 *Erc20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Erc20 *Erc20Session) Name() (string, error) {
	return _Erc20.Contract.Name(&_Erc20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Erc20 *Erc20CallerSession) Name() (string, error) {
	return _Erc20.Contract.Name(&_Erc20.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Erc20 *Erc20Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Erc20 *Erc20Session) Owner() (common.Address, error) {
	return _Erc20.Contract.Owner(&_Erc20.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Erc20 *Erc20CallerSession) Owner() (common.Address, error) {
	return _Erc20.Contract.Owner(&_Erc20.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Erc20 *Erc20Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Erc20 *Erc20Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Erc20.Contract.SupportsInterface(&_Erc20.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Erc20 *Erc20CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Erc20.Contract.SupportsInterface(&_Erc20.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Erc20 *Erc20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Erc20 *Erc20Session) Symbol() (string, error) {
	return _Erc20.Contract.Symbol(&_Erc20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Erc20 *Erc20CallerSession) Symbol() (string, error) {
	return _Erc20.Contract.Symbol(&_Erc20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Erc20 *Erc20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Erc20 *Erc20Session) TotalSupply() (*big.Int, error) {
	return _Erc20.Contract.TotalSupply(&_Erc20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Erc20 *Erc20CallerSession) TotalSupply() (*big.Int, error) {
	return _Erc20.Contract.TotalSupply(&_Erc20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Erc20 *Erc20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Erc20 *Erc20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.Approve(&_Erc20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Erc20 *Erc20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.Approve(&_Erc20.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address from, uint256 amount) returns()
func (_Erc20 *Erc20Transactor) Burn(opts *bind.TransactOpts, from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "burn", from, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address from, uint256 amount) returns()
func (_Erc20 *Erc20Session) Burn(from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.Burn(&_Erc20.TransactOpts, from, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address from, uint256 amount) returns()
func (_Erc20 *Erc20TransactorSession) Burn(from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.Burn(&_Erc20.TransactOpts, from, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_Erc20 *Erc20Transactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "mint", to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_Erc20 *Erc20Session) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.Mint(&_Erc20.TransactOpts, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_Erc20 *Erc20TransactorSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.Mint(&_Erc20.TransactOpts, to, amount)
}

// SudoBurn is a paid mutator transaction binding the contract method 0xfe1195ec.
//
// Solidity: function sudoBurn(address from, uint256 amount) returns()
func (_Erc20 *Erc20Transactor) SudoBurn(opts *bind.TransactOpts, from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "sudoBurn", from, amount)
}

// SudoBurn is a paid mutator transaction binding the contract method 0xfe1195ec.
//
// Solidity: function sudoBurn(address from, uint256 amount) returns()
func (_Erc20 *Erc20Session) SudoBurn(from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.SudoBurn(&_Erc20.TransactOpts, from, amount)
}

// SudoBurn is a paid mutator transaction binding the contract method 0xfe1195ec.
//
// Solidity: function sudoBurn(address from, uint256 amount) returns()
func (_Erc20 *Erc20TransactorSession) SudoBurn(from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.SudoBurn(&_Erc20.TransactOpts, from, amount)
}

// SudoMint is a paid mutator transaction binding the contract method 0x2d688ca8.
//
// Solidity: function sudoMint(address to, uint256 amount) returns()
func (_Erc20 *Erc20Transactor) SudoMint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "sudoMint", to, amount)
}

// SudoMint is a paid mutator transaction binding the contract method 0x2d688ca8.
//
// Solidity: function sudoMint(address to, uint256 amount) returns()
func (_Erc20 *Erc20Session) SudoMint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.SudoMint(&_Erc20.TransactOpts, to, amount)
}

// SudoMint is a paid mutator transaction binding the contract method 0x2d688ca8.
//
// Solidity: function sudoMint(address to, uint256 amount) returns()
func (_Erc20 *Erc20TransactorSession) SudoMint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.SudoMint(&_Erc20.TransactOpts, to, amount)
}

// SudoTransfer is a paid mutator transaction binding the contract method 0x1988513b.
//
// Solidity: function sudoTransfer(address sender, address recipient, uint256 amount) returns()
func (_Erc20 *Erc20Transactor) SudoTransfer(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "sudoTransfer", sender, recipient, amount)
}

// SudoTransfer is a paid mutator transaction binding the contract method 0x1988513b.
//
// Solidity: function sudoTransfer(address sender, address recipient, uint256 amount) returns()
func (_Erc20 *Erc20Session) SudoTransfer(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.SudoTransfer(&_Erc20.TransactOpts, sender, recipient, amount)
}

// SudoTransfer is a paid mutator transaction binding the contract method 0x1988513b.
//
// Solidity: function sudoTransfer(address sender, address recipient, uint256 amount) returns()
func (_Erc20 *Erc20TransactorSession) SudoTransfer(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.SudoTransfer(&_Erc20.TransactOpts, sender, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Erc20 *Erc20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Erc20 *Erc20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.Transfer(&_Erc20.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Erc20 *Erc20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.Transfer(&_Erc20.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Erc20 *Erc20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Erc20 *Erc20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.TransferFrom(&_Erc20.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Erc20 *Erc20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.TransferFrom(&_Erc20.TransactOpts, sender, recipient, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Erc20 *Erc20Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Erc20 *Erc20Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Erc20.Contract.TransferOwnership(&_Erc20.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Erc20 *Erc20TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Erc20.Contract.TransferOwnership(&_Erc20.TransactOpts, newOwner)
}

// Erc20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Erc20 contract.
type Erc20ApprovalIterator struct {
	Event *Erc20Approval // Event containing the contract specifics and raw log

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
func (it *Erc20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc20Approval)
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
		it.Event = new(Erc20Approval)
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
func (it *Erc20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc20Approval represents a Approval event raised by the Erc20 contract.
type Erc20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Erc20 *Erc20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*Erc20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Erc20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &Erc20ApprovalIterator{contract: _Erc20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Erc20 *Erc20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *Erc20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Erc20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc20Approval)
				if err := _Erc20.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Erc20 *Erc20Filterer) ParseApproval(log types.Log) (*Erc20Approval, error) {
	event := new(Erc20Approval)
	if err := _Erc20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc20OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Erc20 contract.
type Erc20OwnershipTransferredIterator struct {
	Event *Erc20OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *Erc20OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc20OwnershipTransferred)
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
		it.Event = new(Erc20OwnershipTransferred)
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
func (it *Erc20OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc20OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc20OwnershipTransferred represents a OwnershipTransferred event raised by the Erc20 contract.
type Erc20OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Erc20 *Erc20Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*Erc20OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Erc20.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &Erc20OwnershipTransferredIterator{contract: _Erc20.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Erc20 *Erc20Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *Erc20OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Erc20.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc20OwnershipTransferred)
				if err := _Erc20.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Erc20 *Erc20Filterer) ParseOwnershipTransferred(log types.Log) (*Erc20OwnershipTransferred, error) {
	event := new(Erc20OwnershipTransferred)
	if err := _Erc20.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Erc20 contract.
type Erc20TransferIterator struct {
	Event *Erc20Transfer // Event containing the contract specifics and raw log

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
func (it *Erc20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc20Transfer)
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
		it.Event = new(Erc20Transfer)
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
func (it *Erc20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc20Transfer represents a Transfer event raised by the Erc20 contract.
type Erc20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Erc20 *Erc20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*Erc20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Erc20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Erc20TransferIterator{contract: _Erc20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Erc20 *Erc20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *Erc20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Erc20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc20Transfer)
				if err := _Erc20.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Erc20 *Erc20Filterer) ParseTransfer(log types.Log) (*Erc20Transfer, error) {
	event := new(Erc20Transfer)
	if err := _Erc20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
