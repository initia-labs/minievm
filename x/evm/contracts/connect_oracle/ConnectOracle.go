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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"pair_id\",\"type\":\"string\"}],\"name\":\"get_price\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"decimal\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"}],\"internalType\":\"structConnectOracle.Price\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"pair_ids\",\"type\":\"string[]\"}],\"name\":\"get_prices\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"decimal\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"}],\"internalType\":\"structConnectOracle.Price[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f80fd5b50611a3b8061001c5f395ff3fe608060405234801561000f575f80fd5b5060043610610034575f3560e01c80636330ac09146100385780639d83565314610068575b5f80fd5b610052600480360381019061004d9190610dd2565b610098565b60405161005f9190610ecc565b60405180910390f35b610082600480360381019061007d9190610fc7565b610269565b60405161008f919061112f565b60405180910390f35b6100a0610c2d565b5f6040518060600160405280602181526020016119ec6021913990505f60f173ffffffffffffffffffffffffffffffffffffffff1663cad235548360f373ffffffffffffffffffffffffffffffffffffffff16638d5c8817886040518263ffffffff1660e01b815260040161011591906111af565b5f60405180830381865afa15801561012f573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f82011682018060405250810190610157919061123d565b6040516020016101679190611352565b6040516020818303038152906040526040518363ffffffff1660e01b815260040161019392919061137e565b5f60405180830381865afa1580156101ad573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f820116820180604052508101906101d5919061123d565b90505f60f373ffffffffffffffffffffffffffffffffffffffff166348ad3c3a836040518263ffffffff1660e01b81526004016102129190611405565b5f60405180830381865afa15801561022c573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f820116820180604052508101906102549190611683565b905061025f8161058d565b9350505050919050565b60605f604051806060016040528060228152602001611a0d6022913990505f61029184610a4b565b6040516020016102a19190611714565b60405160208183030381529060405290505f60f173ffffffffffffffffffffffffffffffffffffffff1663cad2355484846040518363ffffffff1660e01b81526004016102ef92919061137e565b5f60405180830381865afa158015610309573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f82011682018060405250810190610331919061123d565b90505f60f373ffffffffffffffffffffffffffffffffffffffff166348ad3c3a836040518263ffffffff1660e01b815260040161036e9190611405565b5f60405180830381865afa158015610388573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f820116820180604052508101906103b09190611683565b90505f60f373ffffffffffffffffffffffffffffffffffffffff16637e8fa9cd835f01515f815181106103e6576103e5611740565b5b6020026020010151602001516040518263ffffffff1660e01b815260040161040e9190611405565b5f60405180830381865afa158015610428573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f82011682018060405250810190610450919061184b565b90505f815167ffffffffffffffff81111561046e5761046d610cae565b5b6040519080825280602002602001820160405280156104a757816020015b610494610c2d565b81526020019060019003908161048c5790505b5090505f5b825181101561057e5761055360f373ffffffffffffffffffffffffffffffffffffffff166348ad3c3a8584815181106104e8576104e7611740565b5b60200260200101516040518263ffffffff1660e01b815260040161050c9190611405565b5f60405180830381865afa158015610526573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f8201168201806040525081019061054e9190611683565b61058d565b82828151811061056657610565611740565b5b602002602001018190525080806001019150506104ac565b50809650505050505050919050565b610595610c2d565b5f60f373ffffffffffffffffffffffffffffffffffffffff166348ad3c3a845f01516003815181106105ca576105c9611740565b5b6020026020010151602001516040518263ffffffff1660e01b81526004016105f29190611405565b5f60405180830381865afa15801561060c573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f820116820180604052508101906106349190611683565b90505f60f373ffffffffffffffffffffffffffffffffffffffff166385989f68835f015160028151811061066b5761066a611740565b5b6020026020010151602001516040518263ffffffff1660e01b81526004016106939190611405565b602060405180830381865afa1580156106ae573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906106d291906118bc565b90505f60f373ffffffffffffffffffffffffffffffffffffffff16635922f631845f015160018151811061070957610708611740565b5b6020026020010151602001516040518263ffffffff1660e01b81526004016107319190611405565b602060405180830381865afa15801561074c573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061077091906118bc565b90505f60f373ffffffffffffffffffffffffffffffffffffffff166385989f68855f01515f815181106107a6576107a5611740565b5b6020026020010151602001516040518263ffffffff1660e01b81526004016107ce9190611405565b602060405180830381865afa1580156107e9573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061080d91906118bc565b90505f60f373ffffffffffffffffffffffffffffffffffffffff166385989f68885f015160028151811061084457610843611740565b5b6020026020010151602001516040518263ffffffff1660e01b815260040161086c9190611405565b602060405180830381865afa158015610887573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906108ab91906118bc565b90505f60f373ffffffffffffffffffffffffffffffffffffffff166385989f68895f01515f815181106108e1576108e0611740565b5b6020026020010151602001516040518263ffffffff1660e01b81526004016109099190611405565b602060405180830381865afa158015610924573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061094891906118bc565b90505f60f373ffffffffffffffffffffffffffffffffffffffff166385989f688a5f015160018151811061097f5761097e611740565b5b6020026020010151602001516040518263ffffffff1660e01b81526004016109a79190611405565b602060405180830381865afa1580156109c2573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906109e691906118bc565b90506040518060c001604052808781526020018681526020018567ffffffffffffffff1681526020018467ffffffffffffffff1681526020018367ffffffffffffffff1681526020018267ffffffffffffffff16815250975050505050505050919050565b60605f825190505f60f373ffffffffffffffffffffffffffffffffffffffff16638d5c8817855f81518110610a8357610a82611740565b5b60200260200101516040518263ffffffff1660e01b8152600401610aa791906111af565b5f60405180830381865afa158015610ac1573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f82011682018060405250810190610ae9919061123d565b604051602001610af9919061190d565b60405160208183030381529060405290505f600190505b82811015610c005781604051602001610b299190611958565b60405160208183030381529060405291508160f373ffffffffffffffffffffffffffffffffffffffff16638d5c8817878481518110610b6b57610b6a611740565b5b60200260200101516040518263ffffffff1660e01b8152600401610b8f91906111af565b5f60405180830381865afa158015610ba9573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f82011682018060405250810190610bd1919061123d565b604051602001610be292919061197d565b60405160208183030381529060405291508080600101915050610b10565b5080604051602001610c1291906119c6565b60405160208183030381529060405290508092505050919050565b6040518060c001604052805f81526020015f81526020015f67ffffffffffffffff1681526020015f67ffffffffffffffff1681526020015f67ffffffffffffffff1681526020015f67ffffffffffffffff1681525090565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610ce482610c9e565b810181811067ffffffffffffffff82111715610d0357610d02610cae565b5b80604052505050565b5f610d15610c85565b9050610d218282610cdb565b919050565b5f67ffffffffffffffff821115610d4057610d3f610cae565b5b610d4982610c9e565b9050602081019050919050565b828183375f83830152505050565b5f610d76610d7184610d26565b610d0c565b905082815260208101848484011115610d9257610d91610c9a565b5b610d9d848285610d56565b509392505050565b5f82601f830112610db957610db8610c96565b5b8135610dc9848260208601610d64565b91505092915050565b5f60208284031215610de757610de6610c8e565b5b5f82013567ffffffffffffffff811115610e0457610e03610c92565b5b610e1084828501610da5565b91505092915050565b5f819050919050565b610e2b81610e19565b82525050565b5f67ffffffffffffffff82169050919050565b610e4d81610e31565b82525050565b60c082015f820151610e675f850182610e22565b506020820151610e7a6020850182610e22565b506040820151610e8d6040850182610e44565b506060820151610ea06060850182610e44565b506080820151610eb36080850182610e44565b5060a0820151610ec660a0850182610e44565b50505050565b5f60c082019050610edf5f830184610e53565b92915050565b5f67ffffffffffffffff821115610eff57610efe610cae565b5b602082029050602081019050919050565b5f80fd5b5f610f26610f2184610ee5565b610d0c565b90508083825260208201905060208402830185811115610f4957610f48610f10565b5b835b81811015610f9057803567ffffffffffffffff811115610f6e57610f6d610c96565b5b808601610f7b8982610da5565b85526020850194505050602081019050610f4b565b5050509392505050565b5f82601f830112610fae57610fad610c96565b5b8135610fbe848260208601610f14565b91505092915050565b5f60208284031215610fdc57610fdb610c8e565b5b5f82013567ffffffffffffffff811115610ff957610ff8610c92565b5b61100584828501610f9a565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b60c082015f82015161104b5f850182610e22565b50602082015161105e6020850182610e22565b5060408201516110716040850182610e44565b5060608201516110846060850182610e44565b5060808201516110976080850182610e44565b5060a08201516110aa60a0850182610e44565b50505050565b5f6110bb8383611037565b60c08301905092915050565b5f602082019050919050565b5f6110dd8261100e565b6110e78185611018565b93506110f283611028565b805f5b8381101561112257815161110988826110b0565b9750611114836110c7565b9250506001810190506110f5565b5085935050505092915050565b5f6020820190508181035f83015261114781846110d3565b905092915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f6111818261114f565b61118b8185611159565b935061119b818560208601611169565b6111a481610c9e565b840191505092915050565b5f6020820190508181035f8301526111c78184611177565b905092915050565b5f6111e16111dc84610d26565b610d0c565b9050828152602081018484840111156111fd576111fc610c9a565b5b611208848285611169565b509392505050565b5f82601f83011261122457611223610c96565b5b81516112348482602086016111cf565b91505092915050565b5f6020828403121561125257611251610c8e565b5b5f82015167ffffffffffffffff81111561126f5761126e610c92565b5b61127b84828501611210565b91505092915050565b5f81905092915050565b7f7b2263757272656e63795f70616972223a2000000000000000000000000000005f82015250565b5f6112c2601283611284565b91506112cd8261128e565b601282019050919050565b5f6112e28261114f565b6112ec8185611284565b93506112fc818560208601611169565b80840191505092915050565b7f7d000000000000000000000000000000000000000000000000000000000000005f82015250565b5f61133c600183611284565b915061134782611308565b600182019050919050565b5f61135c826112b6565b915061136882846112d8565b915061137382611330565b915081905092915050565b5f6040820190508181035f8301526113968185611177565b905081810360208301526113aa8184611177565b90509392505050565b5f81519050919050565b5f82825260208201905092915050565b5f6113d7826113b3565b6113e181856113bd565b93506113f1818560208601611169565b6113fa81610c9e565b840191505092915050565b5f6020820190508181035f83015261141d81846113cd565b905092915050565b5f80fd5b5f80fd5b5f67ffffffffffffffff82111561144757611446610cae565b5b602082029050602081019050919050565b5f67ffffffffffffffff82111561147257611471610cae565b5b61147b82610c9e565b9050602081019050919050565b5f61149a61149584611458565b610d0c565b9050828152602081018484840111156114b6576114b5610c9a565b5b6114c1848285611169565b509392505050565b5f82601f8301126114dd576114dc610c96565b5b81516114ed848260208601611488565b91505092915050565b5f6040828403121561150b5761150a611425565b5b6115156040610d0c565b90505f82015167ffffffffffffffff81111561153457611533611429565b5b61154084828501611210565b5f83015250602082015167ffffffffffffffff81111561156357611562611429565b5b61156f848285016114c9565b60208301525092915050565b5f61158d6115888461142d565b610d0c565b905080838252602082019050602084028301858111156115b0576115af610f10565b5b835b818110156115f757805167ffffffffffffffff8111156115d5576115d4610c96565b5b8086016115e289826114f6565b855260208501945050506020810190506115b2565b5050509392505050565b5f82601f83011261161557611614610c96565b5b815161162584826020860161157b565b91505092915050565b5f6020828403121561164357611642611425565b5b61164d6020610d0c565b90505f82015167ffffffffffffffff81111561166c5761166b611429565b5b61167884828501611601565b5f8301525092915050565b5f6020828403121561169857611697610c8e565b5b5f82015167ffffffffffffffff8111156116b5576116b4610c92565b5b6116c18482850161162e565b91505092915050565b7f7b2263757272656e63795f706169725f696473223a00000000000000000000005f82015250565b5f6116fe601583611284565b9150611709826116ca565b601582019050919050565b5f61171e826116f2565b915061172a82846112d8565b915061173582611330565b915081905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f67ffffffffffffffff82111561178757611786610cae565b5b602082029050602081019050919050565b5f6117aa6117a58461176d565b610d0c565b905080838252602082019050602084028301858111156117cd576117cc610f10565b5b835b8181101561181457805167ffffffffffffffff8111156117f2576117f1610c96565b5b8086016117ff89826114c9565b855260208501945050506020810190506117cf565b5050509392505050565b5f82601f83011261183257611831610c96565b5b8151611842848260208601611798565b91505092915050565b5f602082840312156118605761185f610c8e565b5b5f82015167ffffffffffffffff81111561187d5761187c610c92565b5b6118898482850161181e565b91505092915050565b61189b81610e19565b81146118a5575f80fd5b50565b5f815190506118b681611892565b92915050565b5f602082840312156118d1576118d0610c8e565b5b5f6118de848285016118a8565b91505092915050565b7f5b00000000000000000000000000000000000000000000000000000000000000815250565b5f611917826118e7565b60018201915061192782846112d8565b915081905092915050565b7f2c00000000000000000000000000000000000000000000000000000000000000815250565b5f61196382846112d8565b915061196e82611932565b60018201915081905092915050565b5f61198882856112d8565b915061199482846112d8565b91508190509392505050565b7f5d00000000000000000000000000000000000000000000000000000000000000815250565b5f6119d182846112d8565b91506119dc826119a0565b6001820191508190509291505056fe2f636f6e6e6563742e6f7261636c652e76322e51756572792f47657450726963652f636f6e6e6563742e6f7261636c652e76322e51756572792f476574507269636573a164736f6c6343000819000a",
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

// GetPrice is a free data retrieval call binding the contract method 0x6330ac09.
//
// Solidity: function get_price(string pair_id) view returns((uint256,uint256,uint64,uint64,uint64,uint64))
func (_ConnectOracle *ConnectOracleCaller) GetPrice(opts *bind.CallOpts, pair_id string) (ConnectOraclePrice, error) {
	var out []interface{}
	err := _ConnectOracle.contract.Call(opts, &out, "get_price", pair_id)

	if err != nil {
		return *new(ConnectOraclePrice), err
	}

	out0 := *abi.ConvertType(out[0], new(ConnectOraclePrice)).(*ConnectOraclePrice)

	return out0, err

}

// GetPrice is a free data retrieval call binding the contract method 0x6330ac09.
//
// Solidity: function get_price(string pair_id) view returns((uint256,uint256,uint64,uint64,uint64,uint64))
func (_ConnectOracle *ConnectOracleSession) GetPrice(pair_id string) (ConnectOraclePrice, error) {
	return _ConnectOracle.Contract.GetPrice(&_ConnectOracle.CallOpts, pair_id)
}

// GetPrice is a free data retrieval call binding the contract method 0x6330ac09.
//
// Solidity: function get_price(string pair_id) view returns((uint256,uint256,uint64,uint64,uint64,uint64))
func (_ConnectOracle *ConnectOracleCallerSession) GetPrice(pair_id string) (ConnectOraclePrice, error) {
	return _ConnectOracle.Contract.GetPrice(&_ConnectOracle.CallOpts, pair_id)
}

// GetPrices is a free data retrieval call binding the contract method 0x9d835653.
//
// Solidity: function get_prices(string[] pair_ids) view returns((uint256,uint256,uint64,uint64,uint64,uint64)[])
func (_ConnectOracle *ConnectOracleCaller) GetPrices(opts *bind.CallOpts, pair_ids []string) ([]ConnectOraclePrice, error) {
	var out []interface{}
	err := _ConnectOracle.contract.Call(opts, &out, "get_prices", pair_ids)

	if err != nil {
		return *new([]ConnectOraclePrice), err
	}

	out0 := *abi.ConvertType(out[0], new([]ConnectOraclePrice)).(*[]ConnectOraclePrice)

	return out0, err

}

// GetPrices is a free data retrieval call binding the contract method 0x9d835653.
//
// Solidity: function get_prices(string[] pair_ids) view returns((uint256,uint256,uint64,uint64,uint64,uint64)[])
func (_ConnectOracle *ConnectOracleSession) GetPrices(pair_ids []string) ([]ConnectOraclePrice, error) {
	return _ConnectOracle.Contract.GetPrices(&_ConnectOracle.CallOpts, pair_ids)
}

// GetPrices is a free data retrieval call binding the contract method 0x9d835653.
//
// Solidity: function get_prices(string[] pair_ids) view returns((uint256,uint256,uint64,uint64,uint64,uint64)[])
func (_ConnectOracle *ConnectOracleCallerSession) GetPrices(pair_ids []string) ([]ConnectOraclePrice, error) {
	return _ConnectOracle.Contract.GetPrices(&_ConnectOracle.CallOpts, pair_ids)
}
