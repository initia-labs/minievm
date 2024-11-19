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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"StringsInsufficientHexLength\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"callback_id\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"name\":\"callback_received\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newCount\",\"type\":\"uint256\"}],\"name\":\"increased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"n\",\"type\":\"uint64\"}],\"name\":\"recursive_called\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"callback_id\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"name\":\"callback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"count\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"exec_msg\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"call_revert\",\"type\":\"bool\"}],\"name\":\"execute_cosmos\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"exec_msg\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"allow_failure\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"callback_id\",\"type\":\"uint64\"}],\"name\":\"execute_cosmos_with_options\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"n\",\"type\":\"uint64\"}],\"name\":\"get_blockhash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"callback_id\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"name\":\"ibc_ack\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"callback_id\",\"type\":\"uint64\"}],\"name\":\"ibc_timeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"increase\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"num\",\"type\":\"uint64\"}],\"name\":\"increase_for_fuzz\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"path\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"req\",\"type\":\"string\"}],\"name\":\"query_cosmos\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"result\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"n\",\"type\":\"uint64\"}],\"name\":\"recursive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040526119d5806100115f395ff3fe60806040526004361061009b575f3560e01c8063619368951161006357806361936895146101695780637876da7514610191578063ac7fde5f146101b9578063c31925a7146101f5578063cad235541461021d578063e8927fbc146102595761009b565b806306661abd1461009f5780630d4f1f9d146100c957806324c68fce146100f15780632607baf81461011957806331a503f014610141575b5f80fd5b3480156100aa575f80fd5b506100b3610263565b6040516100c09190610d0b565b60405180910390f35b3480156100d4575f80fd5b506100ef60048036038101906100ea9190610da7565b610268565b005b3480156100fc575f80fd5b5061011760048036038101906101129190610f21565b6102af565b005b348015610124575f80fd5b5061013f600480360381019061013a9190610f7b565b61036f565b005b34801561014c575f80fd5b5061016760048036038101906101629190610f7b565b6103a6565b005b348015610174575f80fd5b5061018f600480360381019061018a9190610f7b565b6103ca565b005b34801561019c575f80fd5b506101b760048036038101906101b29190610fa6565b61051d565b005b3480156101c4575f80fd5b506101df60048036038101906101da9190610f7b565b6105bf565b6040516101ec919061102a565b60405180910390f35b348015610200575f80fd5b5061021b60048036038101906102169190610da7565b6105d3565b005b348015610228575f80fd5b50610243600480360381019061023e9190611043565b610610565b6040516102509190611119565b60405180910390f35b610261610699565b005b5f5481565b8015610294578167ffffffffffffffff165f808282546102889190611166565b925050819055506102ab565b5f808154809291906102a590611199565b91905055505b5050565b60f173ffffffffffffffffffffffffffffffffffffffff1663d46f64e6836040518263ffffffff1660e01b81526004016102e99190611119565b6020604051808303815f875af1158015610305573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061032991906111f4565b50801561036b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103629061128f565b60405180910390fd5b5050565b5f5b8167ffffffffffffffff168167ffffffffffffffff1610156103a257610395610699565b8080600101915050610371565b5050565b8067ffffffffffffffff165f808282546103c09190611166565b9250508190555050565b7f4475bbd12ef452e28d39c4cb77494d85136c2d89ca1354b52188d4aaa8f4ba51816040516103f991906112bc565b60405180910390a15f8167ffffffffffffffff16031561051a5760f173ffffffffffffffffffffffffffffffffffffffff1663d46f64e6610439836106f8565b6040518263ffffffff1660e01b81526004016104559190611119565b6020604051808303815f875af1158015610471573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061049591906111f4565b5060f173ffffffffffffffffffffffffffffffffffffffff1663d46f64e66104bc836106f8565b6040518263ffffffff1660e01b81526004016104d89190611119565b6020604051808303815f875af11580156104f4573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061051891906111f4565b505b50565b60f173ffffffffffffffffffffffffffffffffffffffff16636c4f6bd584604051806040016040528086151581526020018567ffffffffffffffff168152506040518363ffffffff1660e01b8152600401610579929190611320565b6020604051808303815f875af1158015610595573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906105b991906111f4565b50505050565b5f8167ffffffffffffffff16409050919050565b7fa019c7431cdfd7ba63501ffa1ba7d8f2a028e447653a5af5a96077e5038e0339828260405161060492919061135d565b60405180910390a15050565b606060f173ffffffffffffffffffffffffffffffffffffffff1663cad2355484846040518363ffffffff1660e01b815260040161064e929190611384565b5f604051808303815f875af1158015610669573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f820116820180604052508101906106919190611427565b905092915050565b5f808154809291906106aa90611199565b91905055507f61996fe196f72cb598c483e896a1221263a28bb630480aa89495f737d4a8e3df60015f546106de919061146e565b5f546040516106ee9291906114a1565b60405180910390a1565b606060f173ffffffffffffffffffffffffffffffffffffffff16636af32a55306040518263ffffffff1660e01b81526004016107349190611507565b5f604051808303815f875af115801561074f573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f820116820180604052508101906107779190611427565b61078030610804565b6107dc636193689560e01b6001866107989190611520565b6040516020016107a891906112bc565b6040516020818303038152906040526040516020016107c89291906115ea565b604051602081830303815290604052610831565b6040516020016107ee93929190611877565b6040516020818303038152906040529050919050565b606061082a8273ffffffffffffffffffffffffffffffffffffffff16601460ff16610ab5565b9050919050565b60605f6002808451610843919061190a565b61084d9190611166565b67ffffffffffffffff81111561086657610865610dfd565b5b6040519080825280601f01601f1916602001820160405280156108985781602001600182028036833780820191505090505b5090507f3000000000000000000000000000000000000000000000000000000000000000815f815181106108cf576108ce61194b565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a9053507f7800000000000000000000000000000000000000000000000000000000000000816001815181106109325761093161194b565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a9053505f5b8351811015610aab575f84828151811061097f5761097e61194b565b5b602001015160f81c60f81b60f81c90507f303132333435363738396162636465660000000000000000000000000000000060048260ff16901c60ff16601081106109cc576109cb61194b565b5b1a60f81b83600280850201815181106109e8576109e761194b565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a9053507f3031323334353637383961626364656600000000000000000000000000000000600f821660ff1660108110610a4f57610a4e61194b565b5b1a60f81b836002600160028602010181518110610a6f57610a6e61194b565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a905350508080600101915050610962565b5080915050919050565b60605f8390505f6002846002610acb919061190a565b610ad59190611166565b67ffffffffffffffff811115610aee57610aed610dfd565b5b6040519080825280601f01601f191660200182016040528015610b205781602001600182028036833780820191505090505b5090507f3000000000000000000000000000000000000000000000000000000000000000815f81518110610b5757610b5661194b565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a9053507f780000000000000000000000000000000000000000000000000000000000000081600181518110610bba57610bb961194b565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a9053505f6001856002610bf8919061190a565b610c029190611166565b90505b6001811115610ca1577f3031323334353637383961626364656600000000000000000000000000000000600f841660108110610c4457610c4361194b565b5b1a60f81b828281518110610c5b57610c5a61194b565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a905350600483901c925080610c9a90611978565b9050610c05565b505f8214610ce85784846040517fe22e27eb000000000000000000000000000000000000000000000000000000008152600401610cdf9291906114a1565b60405180910390fd5b809250505092915050565b5f819050919050565b610d0581610cf3565b82525050565b5f602082019050610d1e5f830184610cfc565b92915050565b5f604051905090565b5f80fd5b5f80fd5b5f67ffffffffffffffff82169050919050565b610d5181610d35565b8114610d5b575f80fd5b50565b5f81359050610d6c81610d48565b92915050565b5f8115159050919050565b610d8681610d72565b8114610d90575f80fd5b50565b5f81359050610da181610d7d565b92915050565b5f8060408385031215610dbd57610dbc610d2d565b5b5f610dca85828601610d5e565b9250506020610ddb85828601610d93565b9150509250929050565b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610e3382610ded565b810181811067ffffffffffffffff82111715610e5257610e51610dfd565b5b80604052505050565b5f610e64610d24565b9050610e708282610e2a565b919050565b5f67ffffffffffffffff821115610e8f57610e8e610dfd565b5b610e9882610ded565b9050602081019050919050565b828183375f83830152505050565b5f610ec5610ec084610e75565b610e5b565b905082815260208101848484011115610ee157610ee0610de9565b5b610eec848285610ea5565b509392505050565b5f82601f830112610f0857610f07610de5565b5b8135610f18848260208601610eb3565b91505092915050565b5f8060408385031215610f3757610f36610d2d565b5b5f83013567ffffffffffffffff811115610f5457610f53610d31565b5b610f6085828601610ef4565b9250506020610f7185828601610d93565b9150509250929050565b5f60208284031215610f9057610f8f610d2d565b5b5f610f9d84828501610d5e565b91505092915050565b5f805f60608486031215610fbd57610fbc610d2d565b5b5f84013567ffffffffffffffff811115610fda57610fd9610d31565b5b610fe686828701610ef4565b9350506020610ff786828701610d93565b925050604061100886828701610d5e565b9150509250925092565b5f819050919050565b61102481611012565b82525050565b5f60208201905061103d5f83018461101b565b92915050565b5f806040838503121561105957611058610d2d565b5b5f83013567ffffffffffffffff81111561107657611075610d31565b5b61108285828601610ef4565b925050602083013567ffffffffffffffff8111156110a3576110a2610d31565b5b6110af85828601610ef4565b9150509250929050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f6110eb826110b9565b6110f581856110c3565b93506111058185602086016110d3565b61110e81610ded565b840191505092915050565b5f6020820190508181035f83015261113181846110e1565b905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61117082610cf3565b915061117b83610cf3565b925082820190508082111561119357611192611139565b5b92915050565b5f6111a382610cf3565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036111d5576111d4611139565b5b600182019050919050565b5f815190506111ee81610d7d565b92915050565b5f6020828403121561120957611208610d2d565b5b5f611216848285016111e0565b91505092915050565b7f72657665727420726561736f6e2064756d6d792076616c756520666f722074655f8201527f7374000000000000000000000000000000000000000000000000000000000000602082015250565b5f6112796022836110c3565b91506112848261121f565b604082019050919050565b5f6020820190508181035f8301526112a68161126d565b9050919050565b6112b681610d35565b82525050565b5f6020820190506112cf5f8301846112ad565b92915050565b6112de81610d72565b82525050565b6112ed81610d35565b82525050565b604082015f8201516113075f8501826112d5565b50602082015161131a60208501826112e4565b50505050565b5f6060820190508181035f83015261133881856110e1565b905061134760208301846112f3565b9392505050565b61135781610d72565b82525050565b5f6040820190506113705f8301856112ad565b61137d602083018461134e565b9392505050565b5f6040820190508181035f83015261139c81856110e1565b905081810360208301526113b081846110e1565b90509392505050565b5f6113cb6113c684610e75565b610e5b565b9050828152602081018484840111156113e7576113e6610de9565b5b6113f28482856110d3565b509392505050565b5f82601f83011261140e5761140d610de5565b5b815161141e8482602086016113b9565b91505092915050565b5f6020828403121561143c5761143b610d2d565b5b5f82015167ffffffffffffffff81111561145957611458610d31565b5b611465848285016113fa565b91505092915050565b5f61147882610cf3565b915061148383610cf3565b925082820390508181111561149b5761149a611139565b5b92915050565b5f6040820190506114b45f830185610cfc565b6114c16020830184610cfc565b9392505050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6114f1826114c8565b9050919050565b611501816114e7565b82525050565b5f60208201905061151a5f8301846114f8565b92915050565b5f61152a82610d35565b915061153583610d35565b9250828203905067ffffffffffffffff81111561155557611554611139565b5b92915050565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b5f819050919050565b6115a061159b8261155b565b611586565b82525050565b5f81519050919050565b5f81905092915050565b5f6115c4826115a6565b6115ce81856115b0565b93506115de8185602086016110d3565b80840191505092915050565b5f6115f5828561158f565b60048201915061160582846115ba565b91508190509392505050565b5f81905092915050565b7f7b224074797065223a20222f6d696e6965766d2e65766d2e76312e4d736743615f8201527f6c6c222c00000000000000000000000000000000000000000000000000000000602082015250565b5f611675602483611611565b91506116808261161b565b602482019050919050565b7f2273656e646572223a20220000000000000000000000000000000000000000005f82015250565b5f6116bf600b83611611565b91506116ca8261168b565b600b82019050919050565b5f6116df826110b9565b6116e98185611611565b93506116f98185602086016110d3565b80840191505092915050565b7f222c0000000000000000000000000000000000000000000000000000000000005f82015250565b5f611739600283611611565b915061174482611705565b600282019050919050565b7f22636f6e74726163745f61646472223a202200000000000000000000000000005f82015250565b5f611783601283611611565b915061178e8261174f565b601282019050919050565b7f22696e707574223a2022000000000000000000000000000000000000000000005f82015250565b5f6117cd600a83611611565b91506117d882611799565b600a82019050919050565b7f2276616c7565223a202230222c000000000000000000000000000000000000005f82015250565b5f611817600d83611611565b9150611822826117e3565b600d82019050919050565b7f226163636573735f6c697374223a205b5d7d00000000000000000000000000005f82015250565b5f611861601283611611565b915061186c8261182d565b601282019050919050565b5f61188182611669565b915061188c826116b3565b915061189882866116d5565b91506118a38261172d565b91506118ae82611777565b91506118ba82856116d5565b91506118c58261172d565b91506118d0826117c1565b91506118dc82846116d5565b91506118e78261172d565b91506118f28261180b565b91506118fd82611855565b9150819050949350505050565b5f61191482610cf3565b915061191f83610cf3565b925082820261192d81610cf3565b9150828204841483151761194457611943611139565b5b5092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f61198282610cf3565b91505f820361199457611993611139565b5b60018203905091905056fea26469706673582212206d5cf4b512edb5d7b357c68848de57b6e0c91fa149dd41462cae53686cb7e44d64736f6c63430008190033",
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

// ExecuteCosmos is a paid mutator transaction binding the contract method 0x24c68fce.
//
// Solidity: function execute_cosmos(string exec_msg, bool call_revert) returns()
func (_Counter *CounterTransactor) ExecuteCosmos(opts *bind.TransactOpts, exec_msg string, call_revert bool) (*types.Transaction, error) {
	return _Counter.contract.Transact(opts, "execute_cosmos", exec_msg, call_revert)
}

// ExecuteCosmos is a paid mutator transaction binding the contract method 0x24c68fce.
//
// Solidity: function execute_cosmos(string exec_msg, bool call_revert) returns()
func (_Counter *CounterSession) ExecuteCosmos(exec_msg string, call_revert bool) (*types.Transaction, error) {
	return _Counter.Contract.ExecuteCosmos(&_Counter.TransactOpts, exec_msg, call_revert)
}

// ExecuteCosmos is a paid mutator transaction binding the contract method 0x24c68fce.
//
// Solidity: function execute_cosmos(string exec_msg, bool call_revert) returns()
func (_Counter *CounterTransactorSession) ExecuteCosmos(exec_msg string, call_revert bool) (*types.Transaction, error) {
	return _Counter.Contract.ExecuteCosmos(&_Counter.TransactOpts, exec_msg, call_revert)
}

// ExecuteCosmosWithOptions is a paid mutator transaction binding the contract method 0x7876da75.
//
// Solidity: function execute_cosmos_with_options(string exec_msg, bool allow_failure, uint64 callback_id) returns()
func (_Counter *CounterTransactor) ExecuteCosmosWithOptions(opts *bind.TransactOpts, exec_msg string, allow_failure bool, callback_id uint64) (*types.Transaction, error) {
	return _Counter.contract.Transact(opts, "execute_cosmos_with_options", exec_msg, allow_failure, callback_id)
}

// ExecuteCosmosWithOptions is a paid mutator transaction binding the contract method 0x7876da75.
//
// Solidity: function execute_cosmos_with_options(string exec_msg, bool allow_failure, uint64 callback_id) returns()
func (_Counter *CounterSession) ExecuteCosmosWithOptions(exec_msg string, allow_failure bool, callback_id uint64) (*types.Transaction, error) {
	return _Counter.Contract.ExecuteCosmosWithOptions(&_Counter.TransactOpts, exec_msg, allow_failure, callback_id)
}

// ExecuteCosmosWithOptions is a paid mutator transaction binding the contract method 0x7876da75.
//
// Solidity: function execute_cosmos_with_options(string exec_msg, bool allow_failure, uint64 callback_id) returns()
func (_Counter *CounterTransactorSession) ExecuteCosmosWithOptions(exec_msg string, allow_failure bool, callback_id uint64) (*types.Transaction, error) {
	return _Counter.Contract.ExecuteCosmosWithOptions(&_Counter.TransactOpts, exec_msg, allow_failure, callback_id)
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
