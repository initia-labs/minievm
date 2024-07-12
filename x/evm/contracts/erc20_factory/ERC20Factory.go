// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package erc20_factory

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

// Erc20FactoryMetaData contains all meta data concerning the Erc20Factory contract.
var Erc20FactoryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"erc20\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC20Created\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"name\":\"createERC20\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b5061259e8061001d5f395ff3fe608060405234801562000010575f80fd5b50600436106200002c575f3560e01c806306ef1a861462000030575b5f80fd5b6200004e6004803603810190620000489190620003a5565b62000066565b6040516200005d91906200047f565b60405180910390f35b5f808484846040516200007990620001f3565b62000087939291906200052f565b604051809103905ff080158015620000a1573d5f803e3d5ffd5b50905060f273ffffffffffffffffffffffffffffffffffffffff1663d126274a826040518263ffffffff1660e01b8152600401620000e091906200047f565b6020604051808303815f875af1158015620000fd573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190620001239190620005b2565b508073ffffffffffffffffffffffffffffffffffffffff1663f2fde38b336040518263ffffffff1660e01b81526004016200015f91906200047f565b5f604051808303815f87803b15801562000177575f80fd5b505af11580156200018a573d5f803e3d5ffd5b505050503373ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f85e892981b234101136bc30081e0a5c44345bebc0940193230c20a43b279e2d160405160405180910390a3809150509392505050565b611f8680620005e383390190565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b62000262826200021a565b810181811067ffffffffffffffff821117156200028457620002836200022a565b5b80604052505050565b5f6200029862000201565b9050620002a6828262000257565b919050565b5f67ffffffffffffffff821115620002c857620002c76200022a565b5b620002d3826200021a565b9050602081019050919050565b828183375f83830152505050565b5f62000304620002fe84620002ab565b6200028d565b90508281526020810184848401111562000323576200032262000216565b5b62000330848285620002e0565b509392505050565b5f82601f8301126200034f576200034e62000212565b5b813562000361848260208601620002ee565b91505092915050565b5f60ff82169050919050565b62000381816200036a565b81146200038c575f80fd5b50565b5f813590506200039f8162000376565b92915050565b5f805f60608486031215620003bf57620003be6200020a565b5b5f84013567ffffffffffffffff811115620003df57620003de6200020e565b5b620003ed8682870162000338565b935050602084013567ffffffffffffffff8111156200041157620004106200020e565b5b6200041f8682870162000338565b925050604062000432868287016200038f565b9150509250925092565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f62000467826200043c565b9050919050565b62000479816200045b565b82525050565b5f602082019050620004945f8301846200046e565b92915050565b5f81519050919050565b5f82825260208201905092915050565b5f5b83811015620004d3578082015181840152602081019050620004b6565b5f8484015250505050565b5f620004ea826200049a565b620004f68185620004a4565b935062000508818560208601620004b4565b62000513816200021a565b840191505092915050565b62000529816200036a565b82525050565b5f6060820190508181035f830152620005498186620004de565b905081810360208301526200055f8185620004de565b90506200057060408301846200051e565b949350505050565b5f8115159050919050565b6200058e8162000578565b811462000599575f80fd5b50565b5f81519050620005ac8162000583565b92915050565b5f60208284031215620005ca57620005c96200020a565b5b5f620005d9848285016200059c565b9150509291505056fe608060405234801562000010575f80fd5b5060405162001f8638038062001f8683398181016040528101906200003691906200027c565b335f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555082600390816200008691906200054a565b5081600490816200009891906200054a565b508060055f6101000a81548160ff021916908360ff1602179055505050506200062e565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6200011d82620000d5565b810181811067ffffffffffffffff821117156200013f576200013e620000e5565b5b80604052505050565b5f62000153620000bc565b905062000161828262000112565b919050565b5f67ffffffffffffffff821115620001835762000182620000e5565b5b6200018e82620000d5565b9050602081019050919050565b5f5b83811015620001ba5780820151818401526020810190506200019d565b5f8484015250505050565b5f620001db620001d58462000166565b62000148565b905082815260208101848484011115620001fa57620001f9620000d1565b5b620002078482856200019b565b509392505050565b5f82601f830112620002265762000225620000cd565b5b815162000238848260208601620001c5565b91505092915050565b5f60ff82169050919050565b620002588162000241565b811462000263575f80fd5b50565b5f8151905062000276816200024d565b92915050565b5f805f60608486031215620002965762000295620000c5565b5b5f84015167ffffffffffffffff811115620002b657620002b5620000c9565b5b620002c4868287016200020f565b935050602084015167ffffffffffffffff811115620002e857620002e7620000c9565b5b620002f6868287016200020f565b9250506040620003098682870162000266565b9150509250925092565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806200036257607f821691505b6020821081036200037857620003776200031d565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302620003dc7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826200039f565b620003e886836200039f565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f620004326200042c620004268462000400565b62000409565b62000400565b9050919050565b5f819050919050565b6200044d8362000412565b620004656200045c8262000439565b848454620003ab565b825550505050565b5f90565b6200047b6200046d565b6200048881848462000442565b505050565b5b81811015620004af57620004a35f8262000471565b6001810190506200048e565b5050565b601f821115620004fe57620004c8816200037e565b620004d38462000390565b81016020851015620004e3578190505b620004fb620004f28562000390565b8301826200048d565b50505b505050565b5f82821c905092915050565b5f620005205f198460080262000503565b1980831691505092915050565b5f6200053a83836200050f565b9150826002028217905092915050565b620005558262000313565b67ffffffffffffffff811115620005715762000570620000e5565b5b6200057d82546200034a565b6200058a828285620004b3565b5f60209050601f831160018114620005c0575f8415620005ab578287015190505b620005b785826200052d565b86555062000626565b601f198416620005d0866200037e565b5f5b82811015620005f957848901518255600182019150602085019450602081019050620005d2565b8683101562000619578489015162000615601f8916826200050f565b8355505b6001600288020188555050505b505050505050565b61194a806200063c5f395ff3fe608060405234801561000f575f80fd5b5060043610610109575f3560e01c806340c10f19116100a05780639dc29fac1161006f5780639dc29fac146102b7578063a9059cbb146102d3578063dd62ed3e14610303578063f2fde38b14610333578063fe1195ec1461034f57610109565b806340c10f191461022f57806370a082311461024b5780638da5cb5b1461027b57806395d89b411461029957610109565b80631988513b116100dc5780631988513b146101a957806323b872dd146101c55780632d688ca8146101f5578063313ce5671461021157610109565b806301ffc9a71461010d57806306fdde031461013d578063095ea7b31461015b57806318160ddd1461018b575b5f80fd5b610127600480360381019061012291906112f9565b61036b565b604051610134919061133e565b60405180910390f35b6101456103e4565b60405161015291906113e1565b60405180910390f35b6101756004803603810190610170919061148e565b610470565b604051610182919061133e565b60405180910390f35b61019361055d565b6040516101a091906114db565b60405180910390f35b6101c360048036038101906101be91906114f4565b610563565b005b6101df60048036038101906101da91906114f4565b6105e2565b6040516101ec919061133e565b60405180910390f35b61020f600480360381019061020a919061148e565b610742565b005b6102196107bf565b604051610226919061155f565b60405180910390f35b6102496004803603810190610244919061148e565b6107d1565b005b61026560048036038101906102609190611578565b6108f0565b60405161027291906114db565b60405180910390f35b610283610905565b60405161029091906115b2565b60405180910390f35b6102a1610928565b6040516102ae91906113e1565b60405180910390f35b6102d160048036038101906102cc919061148e565b6109b4565b005b6102ed60048036038101906102e8919061148e565b610ad3565b6040516102fa919061133e565b60405180910390f35b61031d600480360381019061031891906115cb565b610ba4565b60405161032a91906114db565b60405180910390f35b61034d60048036038101906103489190611578565b610bc4565b005b6103696004803603810190610364919061148e565b610d0c565b005b5f7f8da6da19000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614806103dd57506103dc82610d89565b5b9050919050565b600380546103f190611636565b80601f016020809104026020016040519081016040528092919081815260200182805461041d90611636565b80156104685780601f1061043f57610100808354040283529160200191610468565b820191905f5260205f20905b81548152906001019060200180831161044b57829003601f168201915b505050505081565b5f8160025f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9258460405161054b91906114db565b60405180910390a36001905092915050565b60065481565b600173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146105d2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105c9906116b0565b60405180910390fd5b6105dd838383610df2565b505050565b5f8260f173ffffffffffffffffffffffffffffffffffffffff1663f2af9ac9826040518263ffffffff1660e01b815260040161061e91906115b2565b602060405180830381865afa158015610639573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061065d91906116f8565b1561069d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161069490611793565b60405180910390fd5b8260025f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825461072491906117de565b92505081905550610736858585610df2565b60019150509392505050565b600173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146107b1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107a8906116b0565b60405180910390fd5b6107bb8282610ffd565b5050565b60055f9054906101000a900460ff1681565b8160f173ffffffffffffffffffffffffffffffffffffffff1663f2af9ac9826040518263ffffffff1660e01b815260040161080c91906115b2565b602060405180830381865afa158015610827573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061084b91906116f8565b1561088b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108829061185b565b60405180910390fd5b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146108e1575f80fd5b6108eb8383610ffd565b505050565b6001602052805f5260405f205f915090505481565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6004805461093590611636565b80601f016020809104026020016040519081016040528092919081815260200182805461096190611636565b80156109ac5780601f10610983576101008083540402835291602001916109ac565b820191905f5260205f20905b81548152906001019060200180831161098f57829003601f168201915b505050505081565b8160f173ffffffffffffffffffffffffffffffffffffffff166360dc402f826040518263ffffffff1660e01b81526004016109ef91906115b2565b602060405180830381865afa158015610a0a573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610a2e91906116f8565b15610a6e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a65906118c3565b60405180910390fd5b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610ac4575f80fd5b610ace83836111cc565b505050565b5f8260f173ffffffffffffffffffffffffffffffffffffffff1663f2af9ac9826040518263ffffffff1660e01b8152600401610b0f91906115b2565b602060405180830381865afa158015610b2a573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610b4e91906116f8565b15610b8e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b8590611793565b60405180910390fd5b610b99338585610df2565b600191505092915050565b6002602052815f5260405f20602052805f5260405f205f91509150505481565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610c1a575f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610c51575f80fd5b8073ffffffffffffffffffffffffffffffffffffffff165f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3805f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610d7b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d72906116b0565b60405180910390fd5b610d8582826111cc565b5050565b5f7f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b8160f273ffffffffffffffffffffffffffffffffffffffff16634e25ab64826040518263ffffffff1660e01b8152600401610e2d91906115b2565b602060405180830381865afa158015610e48573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610e6c91906116f8565b610eec5760f273ffffffffffffffffffffffffffffffffffffffff1663ceeae52a826040518263ffffffff1660e01b8152600401610eaa91906115b2565b6020604051808303815f875af1158015610ec6573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610eea91906116f8565b505b8160015f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254610f3891906117de565b925050819055508160015f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254610f8b91906118e1565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051610fef91906114db565b60405180910390a350505050565b8160f273ffffffffffffffffffffffffffffffffffffffff16634e25ab64826040518263ffffffff1660e01b815260040161103891906115b2565b602060405180830381865afa158015611053573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061107791906116f8565b6110f75760f273ffffffffffffffffffffffffffffffffffffffff1663ceeae52a826040518263ffffffff1660e01b81526004016110b591906115b2565b6020604051808303815f875af11580156110d1573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906110f591906116f8565b505b8160015f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825461114391906118e1565b925050819055508160065f82825461115b91906118e1565b925050819055508273ffffffffffffffffffffffffffffffffffffffff165f73ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516111bf91906114db565b60405180910390a3505050565b8060015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825461121891906117de565b925050819055508060065f82825461123091906117de565b925050819055505f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8360405161129491906114db565b60405180910390a35050565b5f80fd5b5f7fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b6112d8816112a4565b81146112e2575f80fd5b50565b5f813590506112f3816112cf565b92915050565b5f6020828403121561130e5761130d6112a0565b5b5f61131b848285016112e5565b91505092915050565b5f8115159050919050565b61133881611324565b82525050565b5f6020820190506113515f83018461132f565b92915050565b5f81519050919050565b5f82825260208201905092915050565b5f5b8381101561138e578082015181840152602081019050611373565b5f8484015250505050565b5f601f19601f8301169050919050565b5f6113b382611357565b6113bd8185611361565b93506113cd818560208601611371565b6113d681611399565b840191505092915050565b5f6020820190508181035f8301526113f981846113a9565b905092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61142a82611401565b9050919050565b61143a81611420565b8114611444575f80fd5b50565b5f8135905061145581611431565b92915050565b5f819050919050565b61146d8161145b565b8114611477575f80fd5b50565b5f8135905061148881611464565b92915050565b5f80604083850312156114a4576114a36112a0565b5b5f6114b185828601611447565b92505060206114c28582860161147a565b9150509250929050565b6114d58161145b565b82525050565b5f6020820190506114ee5f8301846114cc565b92915050565b5f805f6060848603121561150b5761150a6112a0565b5b5f61151886828701611447565b935050602061152986828701611447565b925050604061153a8682870161147a565b9150509250925092565b5f60ff82169050919050565b61155981611544565b82525050565b5f6020820190506115725f830184611550565b92915050565b5f6020828403121561158d5761158c6112a0565b5b5f61159a84828501611447565b91505092915050565b6115ac81611420565b82525050565b5f6020820190506115c55f8301846115a3565b92915050565b5f80604083850312156115e1576115e06112a0565b5b5f6115ee85828601611447565b92505060206115ff85828601611447565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061164d57607f821691505b6020821081036116605761165f611609565b5b50919050565b7f45524332303a2063616c6c6572206973206e6f742074686520636861696e00005f82015250565b5f61169a601e83611361565b91506116a582611666565b602082019050919050565b5f6020820190508181035f8301526116c78161168e565b9050919050565b6116d781611324565b81146116e1575f80fd5b50565b5f815190506116f2816116ce565b92915050565b5f6020828403121561170d5761170c6112a0565b5b5f61171a848285016116e4565b91505092915050565b7f45524332303a207472616e7366657220746f20626c6f636b65642061646472655f8201527f7373000000000000000000000000000000000000000000000000000000000000602082015250565b5f61177d602283611361565b915061178882611723565b604082019050919050565b5f6020820190508181035f8301526117aa81611771565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6117e88261145b565b91506117f38361145b565b925082820390508181111561180b5761180a6117b1565b5b92915050565b7f45524332303a206d696e7420746f20626c6f636b6564206164647265737300005f82015250565b5f611845601e83611361565b915061185082611811565b602082019050919050565b5f6020820190508181035f83015261187281611839565b9050919050565b7f45524332303a206275726e2066726f6d206d6f64756c652061646472657373005f82015250565b5f6118ad601f83611361565b91506118b882611879565b602082019050919050565b5f6020820190508181035f8301526118da816118a1565b9050919050565b5f6118eb8261145b565b91506118f68361145b565b925082820190508082111561190e5761190d6117b1565b5b9291505056fea264697066735822122092a15d9b61c6a6de22122c9c81d375bc99e3984e299e75efa4ffe4653b8ae00d64736f6c63430008180033a26469706673582212206b99826e9ebc2324919b3d174bd14ff76ea5b93ccaabee698691b0cfa3c3064964736f6c63430008180033",
}

// Erc20FactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use Erc20FactoryMetaData.ABI instead.
var Erc20FactoryABI = Erc20FactoryMetaData.ABI

// Erc20FactoryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Erc20FactoryMetaData.Bin instead.
var Erc20FactoryBin = Erc20FactoryMetaData.Bin

// DeployErc20Factory deploys a new Ethereum contract, binding an instance of Erc20Factory to it.
func DeployErc20Factory(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Erc20Factory, error) {
	parsed, err := Erc20FactoryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Erc20FactoryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Erc20Factory{Erc20FactoryCaller: Erc20FactoryCaller{contract: contract}, Erc20FactoryTransactor: Erc20FactoryTransactor{contract: contract}, Erc20FactoryFilterer: Erc20FactoryFilterer{contract: contract}}, nil
}

// Erc20Factory is an auto generated Go binding around an Ethereum contract.
type Erc20Factory struct {
	Erc20FactoryCaller     // Read-only binding to the contract
	Erc20FactoryTransactor // Write-only binding to the contract
	Erc20FactoryFilterer   // Log filterer for contract events
}

// Erc20FactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type Erc20FactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20FactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Erc20FactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20FactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Erc20FactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20FactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Erc20FactorySession struct {
	Contract     *Erc20Factory     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Erc20FactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Erc20FactoryCallerSession struct {
	Contract *Erc20FactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// Erc20FactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Erc20FactoryTransactorSession struct {
	Contract     *Erc20FactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// Erc20FactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type Erc20FactoryRaw struct {
	Contract *Erc20Factory // Generic contract binding to access the raw methods on
}

// Erc20FactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Erc20FactoryCallerRaw struct {
	Contract *Erc20FactoryCaller // Generic read-only contract binding to access the raw methods on
}

// Erc20FactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Erc20FactoryTransactorRaw struct {
	Contract *Erc20FactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewErc20Factory creates a new instance of Erc20Factory, bound to a specific deployed contract.
func NewErc20Factory(address common.Address, backend bind.ContractBackend) (*Erc20Factory, error) {
	contract, err := bindErc20Factory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Erc20Factory{Erc20FactoryCaller: Erc20FactoryCaller{contract: contract}, Erc20FactoryTransactor: Erc20FactoryTransactor{contract: contract}, Erc20FactoryFilterer: Erc20FactoryFilterer{contract: contract}}, nil
}

// NewErc20FactoryCaller creates a new read-only instance of Erc20Factory, bound to a specific deployed contract.
func NewErc20FactoryCaller(address common.Address, caller bind.ContractCaller) (*Erc20FactoryCaller, error) {
	contract, err := bindErc20Factory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Erc20FactoryCaller{contract: contract}, nil
}

// NewErc20FactoryTransactor creates a new write-only instance of Erc20Factory, bound to a specific deployed contract.
func NewErc20FactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*Erc20FactoryTransactor, error) {
	contract, err := bindErc20Factory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Erc20FactoryTransactor{contract: contract}, nil
}

// NewErc20FactoryFilterer creates a new log filterer instance of Erc20Factory, bound to a specific deployed contract.
func NewErc20FactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*Erc20FactoryFilterer, error) {
	contract, err := bindErc20Factory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Erc20FactoryFilterer{contract: contract}, nil
}

// bindErc20Factory binds a generic wrapper to an already deployed contract.
func bindErc20Factory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Erc20FactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc20Factory *Erc20FactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc20Factory.Contract.Erc20FactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc20Factory *Erc20FactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc20Factory.Contract.Erc20FactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc20Factory *Erc20FactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc20Factory.Contract.Erc20FactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc20Factory *Erc20FactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc20Factory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc20Factory *Erc20FactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc20Factory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc20Factory *Erc20FactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc20Factory.Contract.contract.Transact(opts, method, params...)
}

// CreateERC20 is a paid mutator transaction binding the contract method 0x06ef1a86.
//
// Solidity: function createERC20(string name, string symbol, uint8 decimals) returns(address)
func (_Erc20Factory *Erc20FactoryTransactor) CreateERC20(opts *bind.TransactOpts, name string, symbol string, decimals uint8) (*types.Transaction, error) {
	return _Erc20Factory.contract.Transact(opts, "createERC20", name, symbol, decimals)
}

// CreateERC20 is a paid mutator transaction binding the contract method 0x06ef1a86.
//
// Solidity: function createERC20(string name, string symbol, uint8 decimals) returns(address)
func (_Erc20Factory *Erc20FactorySession) CreateERC20(name string, symbol string, decimals uint8) (*types.Transaction, error) {
	return _Erc20Factory.Contract.CreateERC20(&_Erc20Factory.TransactOpts, name, symbol, decimals)
}

// CreateERC20 is a paid mutator transaction binding the contract method 0x06ef1a86.
//
// Solidity: function createERC20(string name, string symbol, uint8 decimals) returns(address)
func (_Erc20Factory *Erc20FactoryTransactorSession) CreateERC20(name string, symbol string, decimals uint8) (*types.Transaction, error) {
	return _Erc20Factory.Contract.CreateERC20(&_Erc20Factory.TransactOpts, name, symbol, decimals)
}

// Erc20FactoryERC20CreatedIterator is returned from FilterERC20Created and is used to iterate over the raw logs and unpacked data for ERC20Created events raised by the Erc20Factory contract.
type Erc20FactoryERC20CreatedIterator struct {
	Event *Erc20FactoryERC20Created // Event containing the contract specifics and raw log

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
func (it *Erc20FactoryERC20CreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc20FactoryERC20Created)
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
		it.Event = new(Erc20FactoryERC20Created)
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
func (it *Erc20FactoryERC20CreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc20FactoryERC20CreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc20FactoryERC20Created represents a ERC20Created event raised by the Erc20Factory contract.
type Erc20FactoryERC20Created struct {
	Erc20 common.Address
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterERC20Created is a free log retrieval operation binding the contract event 0x85e892981b234101136bc30081e0a5c44345bebc0940193230c20a43b279e2d1.
//
// Solidity: event ERC20Created(address indexed erc20, address indexed owner)
func (_Erc20Factory *Erc20FactoryFilterer) FilterERC20Created(opts *bind.FilterOpts, erc20 []common.Address, owner []common.Address) (*Erc20FactoryERC20CreatedIterator, error) {

	var erc20Rule []interface{}
	for _, erc20Item := range erc20 {
		erc20Rule = append(erc20Rule, erc20Item)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Erc20Factory.contract.FilterLogs(opts, "ERC20Created", erc20Rule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &Erc20FactoryERC20CreatedIterator{contract: _Erc20Factory.contract, event: "ERC20Created", logs: logs, sub: sub}, nil
}

// WatchERC20Created is a free log subscription operation binding the contract event 0x85e892981b234101136bc30081e0a5c44345bebc0940193230c20a43b279e2d1.
//
// Solidity: event ERC20Created(address indexed erc20, address indexed owner)
func (_Erc20Factory *Erc20FactoryFilterer) WatchERC20Created(opts *bind.WatchOpts, sink chan<- *Erc20FactoryERC20Created, erc20 []common.Address, owner []common.Address) (event.Subscription, error) {

	var erc20Rule []interface{}
	for _, erc20Item := range erc20 {
		erc20Rule = append(erc20Rule, erc20Item)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Erc20Factory.contract.WatchLogs(opts, "ERC20Created", erc20Rule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc20FactoryERC20Created)
				if err := _Erc20Factory.contract.UnpackLog(event, "ERC20Created", log); err != nil {
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

// ParseERC20Created is a log parse operation binding the contract event 0x85e892981b234101136bc30081e0a5c44345bebc0940193230c20a43b279e2d1.
//
// Solidity: event ERC20Created(address indexed erc20, address indexed owner)
func (_Erc20Factory *Erc20FactoryFilterer) ParseERC20Created(log types.Log) (*Erc20FactoryERC20Created, error) {
	event := new(Erc20FactoryERC20Created)
	if err := _Erc20Factory.contract.UnpackLog(event, "ERC20Created", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
