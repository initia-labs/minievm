// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package connect_oracle

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

// ConnectOraclePrice is an auto generated low-level Go binding around an user-defined struct.
type ConnectOraclePrice struct {
	Price     *big.Int
	Timestamp *big.Int
	Height    uint64
	Nonce     uint64
	Decimal   uint64
	Id        uint64
}

// ConnectOracleMetaData contains all meta data concerning the ConnectOracle contract.
var ConnectOracleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"pair_id\",\"type\":\"string\"}],\"name\":\"get_price\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"decimal\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"}],\"internalType\":\"structConnectOracle.Price\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"pair_ids\",\"type\":\"string[]\"}],\"name\":\"get_prices\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"decimal\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"}],\"internalType\":\"structConnectOracle.Price[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f80fd5b50611a3d8061001c5f395ff3fe608060405234801561000f575f80fd5b5060043610610034575f3560e01c80636330ac09146100385780639d83565314610068575b5f80fd5b610052600480360381019061004d9190610dd4565b610098565b60405161005f9190610ece565b60405180910390f35b610082600480360381019061007d9190610fc9565b61026a565b60405161008f9190611131565b60405180910390f35b6100a0610c2f565b5f6040518060600160405280602181526020016119ee6021913990505f60f173ffffffffffffffffffffffffffffffffffffffff1663cad235548360f373ffffffffffffffffffffffffffffffffffffffff16638d5c8817886040518263ffffffff1660e01b815260040161011591906111b1565b5f60405180830381865afa15801561012f573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f82011682018060405250810190610157919061123f565b6040516020016101679190611354565b6040516020818303038152906040526040518363ffffffff1660e01b8152600401610193929190611380565b5f604051808303815f875af11580156101ae573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f820116820180604052508101906101d6919061123f565b90505f60f373ffffffffffffffffffffffffffffffffffffffff166348ad3c3a836040518263ffffffff1660e01b81526004016102139190611407565b5f60405180830381865afa15801561022d573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f820116820180604052508101906102559190611685565b90506102608161058f565b9350505050919050565b60605f604051806060016040528060228152602001611a0f6022913990505f61029284610a4d565b6040516020016102a29190611716565b60405160208183030381529060405290505f60f173ffffffffffffffffffffffffffffffffffffffff1663cad2355484846040518363ffffffff1660e01b81526004016102f0929190611380565b5f604051808303815f875af115801561030b573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f82011682018060405250810190610333919061123f565b90505f60f373ffffffffffffffffffffffffffffffffffffffff166348ad3c3a836040518263ffffffff1660e01b81526004016103709190611407565b5f60405180830381865afa15801561038a573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f820116820180604052508101906103b29190611685565b90505f60f373ffffffffffffffffffffffffffffffffffffffff16637e8fa9cd835f01515f815181106103e8576103e7611742565b5b6020026020010151602001516040518263ffffffff1660e01b81526004016104109190611407565b5f60405180830381865afa15801561042a573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f82011682018060405250810190610452919061184d565b90505f815167ffffffffffffffff8111156104705761046f610cb0565b5b6040519080825280602002602001820160405280156104a957816020015b610496610c2f565b81526020019060019003908161048e5790505b5090505f5b82518110156105805761055560f373ffffffffffffffffffffffffffffffffffffffff166348ad3c3a8584815181106104ea576104e9611742565b5b60200260200101516040518263ffffffff1660e01b815260040161050e9190611407565b5f60405180830381865afa158015610528573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f820116820180604052508101906105509190611685565b61058f565b82828151811061056857610567611742565b5b602002602001018190525080806001019150506104ae565b50809650505050505050919050565b610597610c2f565b5f60f373ffffffffffffffffffffffffffffffffffffffff166348ad3c3a845f01516003815181106105cc576105cb611742565b5b6020026020010151602001516040518263ffffffff1660e01b81526004016105f49190611407565b5f60405180830381865afa15801561060e573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f820116820180604052508101906106369190611685565b90505f60f373ffffffffffffffffffffffffffffffffffffffff166385989f68835f015160028151811061066d5761066c611742565b5b6020026020010151602001516040518263ffffffff1660e01b81526004016106959190611407565b602060405180830381865afa1580156106b0573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906106d491906118be565b90505f60f373ffffffffffffffffffffffffffffffffffffffff16635922f631845f015160018151811061070b5761070a611742565b5b6020026020010151602001516040518263ffffffff1660e01b81526004016107339190611407565b602060405180830381865afa15801561074e573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061077291906118be565b90505f60f373ffffffffffffffffffffffffffffffffffffffff166385989f68855f01515f815181106107a8576107a7611742565b5b6020026020010151602001516040518263ffffffff1660e01b81526004016107d09190611407565b602060405180830381865afa1580156107eb573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061080f91906118be565b90505f60f373ffffffffffffffffffffffffffffffffffffffff166385989f68885f015160028151811061084657610845611742565b5b6020026020010151602001516040518263ffffffff1660e01b815260040161086e9190611407565b602060405180830381865afa158015610889573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906108ad91906118be565b90505f60f373ffffffffffffffffffffffffffffffffffffffff166385989f68895f01515f815181106108e3576108e2611742565b5b6020026020010151602001516040518263ffffffff1660e01b815260040161090b9190611407565b602060405180830381865afa158015610926573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061094a91906118be565b90505f60f373ffffffffffffffffffffffffffffffffffffffff166385989f688a5f015160018151811061098157610980611742565b5b6020026020010151602001516040518263ffffffff1660e01b81526004016109a99190611407565b602060405180830381865afa1580156109c4573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906109e891906118be565b90506040518060c001604052808781526020018681526020018567ffffffffffffffff1681526020018467ffffffffffffffff1681526020018367ffffffffffffffff1681526020018267ffffffffffffffff16815250975050505050505050919050565b60605f825190505f60f373ffffffffffffffffffffffffffffffffffffffff16638d5c8817855f81518110610a8557610a84611742565b5b60200260200101516040518263ffffffff1660e01b8152600401610aa991906111b1565b5f60405180830381865afa158015610ac3573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f82011682018060405250810190610aeb919061123f565b604051602001610afb919061190f565b60405160208183030381529060405290505f600190505b82811015610c025781604051602001610b2b919061195a565b60405160208183030381529060405291508160f373ffffffffffffffffffffffffffffffffffffffff16638d5c8817878481518110610b6d57610b6c611742565b5b60200260200101516040518263ffffffff1660e01b8152600401610b9191906111b1565b5f60405180830381865afa158015610bab573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f82011682018060405250810190610bd3919061123f565b604051602001610be492919061197f565b60405160208183030381529060405291508080600101915050610b12565b5080604051602001610c1491906119c8565b60405160208183030381529060405290508092505050919050565b6040518060c001604052805f81526020015f81526020015f67ffffffffffffffff1681526020015f67ffffffffffffffff1681526020015f67ffffffffffffffff1681526020015f67ffffffffffffffff1681525090565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610ce682610ca0565b810181811067ffffffffffffffff82111715610d0557610d04610cb0565b5b80604052505050565b5f610d17610c87565b9050610d238282610cdd565b919050565b5f67ffffffffffffffff821115610d4257610d41610cb0565b5b610d4b82610ca0565b9050602081019050919050565b828183375f83830152505050565b5f610d78610d7384610d28565b610d0e565b905082815260208101848484011115610d9457610d93610c9c565b5b610d9f848285610d58565b509392505050565b5f82601f830112610dbb57610dba610c98565b5b8135610dcb848260208601610d66565b91505092915050565b5f60208284031215610de957610de8610c90565b5b5f82013567ffffffffffffffff811115610e0657610e05610c94565b5b610e1284828501610da7565b91505092915050565b5f819050919050565b610e2d81610e1b565b82525050565b5f67ffffffffffffffff82169050919050565b610e4f81610e33565b82525050565b60c082015f820151610e695f850182610e24565b506020820151610e7c6020850182610e24565b506040820151610e8f6040850182610e46565b506060820151610ea26060850182610e46565b506080820151610eb56080850182610e46565b5060a0820151610ec860a0850182610e46565b50505050565b5f60c082019050610ee15f830184610e55565b92915050565b5f67ffffffffffffffff821115610f0157610f00610cb0565b5b602082029050602081019050919050565b5f80fd5b5f610f28610f2384610ee7565b610d0e565b90508083825260208201905060208402830185811115610f4b57610f4a610f12565b5b835b81811015610f9257803567ffffffffffffffff811115610f7057610f6f610c98565b5b808601610f7d8982610da7565b85526020850194505050602081019050610f4d565b5050509392505050565b5f82601f830112610fb057610faf610c98565b5b8135610fc0848260208601610f16565b91505092915050565b5f60208284031215610fde57610fdd610c90565b5b5f82013567ffffffffffffffff811115610ffb57610ffa610c94565b5b61100784828501610f9c565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b60c082015f82015161104d5f850182610e24565b5060208201516110606020850182610e24565b5060408201516110736040850182610e46565b5060608201516110866060850182610e46565b5060808201516110996080850182610e46565b5060a08201516110ac60a0850182610e46565b50505050565b5f6110bd8383611039565b60c08301905092915050565b5f602082019050919050565b5f6110df82611010565b6110e9818561101a565b93506110f48361102a565b805f5b8381101561112457815161110b88826110b2565b9750611116836110c9565b9250506001810190506110f7565b5085935050505092915050565b5f6020820190508181035f83015261114981846110d5565b905092915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f61118382611151565b61118d818561115b565b935061119d81856020860161116b565b6111a681610ca0565b840191505092915050565b5f6020820190508181035f8301526111c98184611179565b905092915050565b5f6111e36111de84610d28565b610d0e565b9050828152602081018484840111156111ff576111fe610c9c565b5b61120a84828561116b565b509392505050565b5f82601f83011261122657611225610c98565b5b81516112368482602086016111d1565b91505092915050565b5f6020828403121561125457611253610c90565b5b5f82015167ffffffffffffffff81111561127157611270610c94565b5b61127d84828501611212565b91505092915050565b5f81905092915050565b7f7b2263757272656e63795f70616972223a2000000000000000000000000000005f82015250565b5f6112c4601283611286565b91506112cf82611290565b601282019050919050565b5f6112e482611151565b6112ee8185611286565b93506112fe81856020860161116b565b80840191505092915050565b7f7d000000000000000000000000000000000000000000000000000000000000005f82015250565b5f61133e600183611286565b91506113498261130a565b600182019050919050565b5f61135e826112b8565b915061136a82846112da565b915061137582611332565b915081905092915050565b5f6040820190508181035f8301526113988185611179565b905081810360208301526113ac8184611179565b90509392505050565b5f81519050919050565b5f82825260208201905092915050565b5f6113d9826113b5565b6113e381856113bf565b93506113f381856020860161116b565b6113fc81610ca0565b840191505092915050565b5f6020820190508181035f83015261141f81846113cf565b905092915050565b5f80fd5b5f80fd5b5f67ffffffffffffffff82111561144957611448610cb0565b5b602082029050602081019050919050565b5f67ffffffffffffffff82111561147457611473610cb0565b5b61147d82610ca0565b9050602081019050919050565b5f61149c6114978461145a565b610d0e565b9050828152602081018484840111156114b8576114b7610c9c565b5b6114c384828561116b565b509392505050565b5f82601f8301126114df576114de610c98565b5b81516114ef84826020860161148a565b91505092915050565b5f6040828403121561150d5761150c611427565b5b6115176040610d0e565b90505f82015167ffffffffffffffff8111156115365761153561142b565b5b61154284828501611212565b5f83015250602082015167ffffffffffffffff8111156115655761156461142b565b5b611571848285016114cb565b60208301525092915050565b5f61158f61158a8461142f565b610d0e565b905080838252602082019050602084028301858111156115b2576115b1610f12565b5b835b818110156115f957805167ffffffffffffffff8111156115d7576115d6610c98565b5b8086016115e489826114f8565b855260208501945050506020810190506115b4565b5050509392505050565b5f82601f83011261161757611616610c98565b5b815161162784826020860161157d565b91505092915050565b5f6020828403121561164557611644611427565b5b61164f6020610d0e565b90505f82015167ffffffffffffffff81111561166e5761166d61142b565b5b61167a84828501611603565b5f8301525092915050565b5f6020828403121561169a57611699610c90565b5b5f82015167ffffffffffffffff8111156116b7576116b6610c94565b5b6116c384828501611630565b91505092915050565b7f7b2263757272656e63795f706169725f696473223a00000000000000000000005f82015250565b5f611700601583611286565b915061170b826116cc565b601582019050919050565b5f611720826116f4565b915061172c82846112da565b915061173782611332565b915081905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f67ffffffffffffffff82111561178957611788610cb0565b5b602082029050602081019050919050565b5f6117ac6117a78461176f565b610d0e565b905080838252602082019050602084028301858111156117cf576117ce610f12565b5b835b8181101561181657805167ffffffffffffffff8111156117f4576117f3610c98565b5b80860161180189826114cb565b855260208501945050506020810190506117d1565b5050509392505050565b5f82601f83011261183457611833610c98565b5b815161184484826020860161179a565b91505092915050565b5f6020828403121561186257611861610c90565b5b5f82015167ffffffffffffffff81111561187f5761187e610c94565b5b61188b84828501611820565b91505092915050565b61189d81610e1b565b81146118a7575f80fd5b50565b5f815190506118b881611894565b92915050565b5f602082840312156118d3576118d2610c90565b5b5f6118e0848285016118aa565b91505092915050565b7f5b00000000000000000000000000000000000000000000000000000000000000815250565b5f611919826118e9565b60018201915061192982846112da565b915081905092915050565b7f2c00000000000000000000000000000000000000000000000000000000000000815250565b5f61196582846112da565b915061197082611934565b60018201915081905092915050565b5f61198a82856112da565b915061199682846112da565b91508190509392505050565b7f5d00000000000000000000000000000000000000000000000000000000000000815250565b5f6119d382846112da565b91506119de826119a2565b6001820191508190509291505056fe2f636f6e6e6563742e6f7261636c652e76322e51756572792f47657450726963652f636f6e6e6563742e6f7261636c652e76322e51756572792f476574507269636573a164736f6c6343000819000a",
}

// ConnectOracleABI is the input ABI used to generate the binding from.
// Deprecated: Use ConnectOracleMetaData.ABI instead.
var ConnectOracleABI = ConnectOracleMetaData.ABI

// ConnectOracleBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ConnectOracleMetaData.Bin instead.
var ConnectOracleBin = ConnectOracleMetaData.Bin

// DeployConnectOracle deploys a new Ethereum contract, binding an instance of ConnectOracle to it.
func DeployConnectOracle(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ConnectOracle, error) {
	parsed, err := ConnectOracleMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ConnectOracleBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ConnectOracle{ConnectOracleCaller: ConnectOracleCaller{contract: contract}, ConnectOracleTransactor: ConnectOracleTransactor{contract: contract}, ConnectOracleFilterer: ConnectOracleFilterer{contract: contract}}, nil
}

// ConnectOracle is an auto generated Go binding around an Ethereum contract.
type ConnectOracle struct {
	ConnectOracleCaller     // Read-only binding to the contract
	ConnectOracleTransactor // Write-only binding to the contract
	ConnectOracleFilterer   // Log filterer for contract events
}

// ConnectOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type ConnectOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConnectOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ConnectOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConnectOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ConnectOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConnectOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ConnectOracleSession struct {
	Contract     *ConnectOracle    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ConnectOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ConnectOracleCallerSession struct {
	Contract *ConnectOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ConnectOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ConnectOracleTransactorSession struct {
	Contract     *ConnectOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ConnectOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type ConnectOracleRaw struct {
	Contract *ConnectOracle // Generic contract binding to access the raw methods on
}

// ConnectOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ConnectOracleCallerRaw struct {
	Contract *ConnectOracleCaller // Generic read-only contract binding to access the raw methods on
}

// ConnectOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ConnectOracleTransactorRaw struct {
	Contract *ConnectOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewConnectOracle creates a new instance of ConnectOracle, bound to a specific deployed contract.
func NewConnectOracle(address common.Address, backend bind.ContractBackend) (*ConnectOracle, error) {
	contract, err := bindConnectOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ConnectOracle{ConnectOracleCaller: ConnectOracleCaller{contract: contract}, ConnectOracleTransactor: ConnectOracleTransactor{contract: contract}, ConnectOracleFilterer: ConnectOracleFilterer{contract: contract}}, nil
}

// NewConnectOracleCaller creates a new read-only instance of ConnectOracle, bound to a specific deployed contract.
func NewConnectOracleCaller(address common.Address, caller bind.ContractCaller) (*ConnectOracleCaller, error) {
	contract, err := bindConnectOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ConnectOracleCaller{contract: contract}, nil
}

// NewConnectOracleTransactor creates a new write-only instance of ConnectOracle, bound to a specific deployed contract.
func NewConnectOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*ConnectOracleTransactor, error) {
	contract, err := bindConnectOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ConnectOracleTransactor{contract: contract}, nil
}

// NewConnectOracleFilterer creates a new log filterer instance of ConnectOracle, bound to a specific deployed contract.
func NewConnectOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*ConnectOracleFilterer, error) {
	contract, err := bindConnectOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ConnectOracleFilterer{contract: contract}, nil
}

// bindConnectOracle binds a generic wrapper to an already deployed contract.
func bindConnectOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ConnectOracleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ConnectOracle *ConnectOracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ConnectOracle.Contract.ConnectOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ConnectOracle *ConnectOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConnectOracle.Contract.ConnectOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ConnectOracle *ConnectOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ConnectOracle.Contract.ConnectOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ConnectOracle *ConnectOracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ConnectOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ConnectOracle *ConnectOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConnectOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ConnectOracle *ConnectOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ConnectOracle.Contract.contract.Transact(opts, method, params...)
}

// GetPrice is a paid mutator transaction binding the contract method 0x6330ac09.
//
// Solidity: function get_price(string pair_id) returns((uint256,uint256,uint64,uint64,uint64,uint64))
func (_ConnectOracle *ConnectOracleTransactor) GetPrice(opts *bind.TransactOpts, pair_id string) (*types.Transaction, error) {
	return _ConnectOracle.contract.Transact(opts, "get_price", pair_id)
}

// GetPrice is a paid mutator transaction binding the contract method 0x6330ac09.
//
// Solidity: function get_price(string pair_id) returns((uint256,uint256,uint64,uint64,uint64,uint64))
func (_ConnectOracle *ConnectOracleSession) GetPrice(pair_id string) (*types.Transaction, error) {
	return _ConnectOracle.Contract.GetPrice(&_ConnectOracle.TransactOpts, pair_id)
}

// GetPrice is a paid mutator transaction binding the contract method 0x6330ac09.
//
// Solidity: function get_price(string pair_id) returns((uint256,uint256,uint64,uint64,uint64,uint64))
func (_ConnectOracle *ConnectOracleTransactorSession) GetPrice(pair_id string) (*types.Transaction, error) {
	return _ConnectOracle.Contract.GetPrice(&_ConnectOracle.TransactOpts, pair_id)
}

// GetPrices is a paid mutator transaction binding the contract method 0x9d835653.
//
// Solidity: function get_prices(string[] pair_ids) returns((uint256,uint256,uint64,uint64,uint64,uint64)[])
func (_ConnectOracle *ConnectOracleTransactor) GetPrices(opts *bind.TransactOpts, pair_ids []string) (*types.Transaction, error) {
	return _ConnectOracle.contract.Transact(opts, "get_prices", pair_ids)
}

// GetPrices is a paid mutator transaction binding the contract method 0x9d835653.
//
// Solidity: function get_prices(string[] pair_ids) returns((uint256,uint256,uint64,uint64,uint64,uint64)[])
func (_ConnectOracle *ConnectOracleSession) GetPrices(pair_ids []string) (*types.Transaction, error) {
	return _ConnectOracle.Contract.GetPrices(&_ConnectOracle.TransactOpts, pair_ids)
}

// GetPrices is a paid mutator transaction binding the contract method 0x9d835653.
//
// Solidity: function get_prices(string[] pair_ids) returns((uint256,uint256,uint64,uint64,uint64,uint64)[])
func (_ConnectOracle *ConnectOracleTransactorSession) GetPrices(pair_ids []string) (*types.Transaction, error) {
	return _ConnectOracle.Contract.GetPrices(&_ConnectOracle.TransactOpts, pair_ids)
}
