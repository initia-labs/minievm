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
	Bin: "0x6080604052348015600e575f80fd5b5061185b8061001c5f395ff3fe608060405234801561000f575f80fd5b5060043610610034575f3560e01c80636330ac09146100385780639d83565314610068575b5f80fd5b610052600480360381019061004d9190610d93565b610098565b60405161005f9190610e8d565b60405180910390f35b610082600480360381019061007d9190610f88565b610308565b60405161008f91906110f0565b60405180910390f35b6100a0610bee565b5f60405180606001604052806021815260200161180c6021913990505f600367ffffffffffffffff8111156100d8576100d7610c6f565b5b60405190808252806020026020018201604052801561010b57816020015b60608152602001906001900390816100f65790505b5090506040518060400160405280601381526020017f7b2263757272656e63795f70616972223a202200000000000000000000000000815250815f8151811061015757610156611110565b5b6020026020010181905250838160018151811061017757610176611110565b5b60200260200101819052506040518060400160405280600281526020017f227d000000000000000000000000000000000000000000000000000000000000815250816002815181106101cc576101cb611110565b5b60200260200101819052505f6101f08260405180602001604052805f815250610682565b90505f60f173ffffffffffffffffffffffffffffffffffffffff1663cad2355485846040518363ffffffff1660e01b815260040161022f92919061119d565b5f604051808303815f875af115801561024a573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f820116820180604052508101906102729190611240565b90505f60f373ffffffffffffffffffffffffffffffffffffffff166348ad3c3a836040518263ffffffff1660e01b81526004016102af91906112d9565b5f60405180830381865afa1580156102c9573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f820116820180604052508101906102f19190611557565b90506102fc81610730565b95505050505050919050565b60605f60405180606001604052806022815260200161182d6022913990505f610366846040518060400160405280600381526020017f222c220000000000000000000000000000000000000000000000000000000000815250610682565b60405160200161037691906115fe565b6040516020818303038152906040526040516020016103959190611649565b60405160208183030381529060405290505f60f173ffffffffffffffffffffffffffffffffffffffff1663cad2355484846040518363ffffffff1660e01b81526004016103e392919061119d565b5f604051808303815f875af11580156103fe573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f820116820180604052508101906104269190611240565b90505f60f373ffffffffffffffffffffffffffffffffffffffff166348ad3c3a836040518263ffffffff1660e01b815260040161046391906112d9565b5f60405180830381865afa15801561047d573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f820116820180604052508101906104a59190611557565b90505f60f373ffffffffffffffffffffffffffffffffffffffff16637e8fa9cd835f01515f815181106104db576104da611110565b5b6020026020010151602001516040518263ffffffff1660e01b815260040161050391906112d9565b5f60405180830381865afa15801561051d573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f82011682018060405250810190610545919061174c565b90505f815167ffffffffffffffff81111561056357610562610c6f565b5b60405190808252806020026020018201604052801561059c57816020015b610589610bee565b8152602001906001900390816105815790505b5090505f5b82518110156106735761064860f373ffffffffffffffffffffffffffffffffffffffff166348ad3c3a8584815181106105dd576105dc611110565b5b60200260200101516040518263ffffffff1660e01b815260040161060191906112d9565b5f60405180830381865afa15801561061b573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f820116820180604052508101906106439190611557565b610730565b82828151811061065b5761065a611110565b5b602002602001018190525080806001019150506105a1565b50809650505050505050919050565b60605f835190505f845f8151811061069d5761069c611110565b5b602002602001015190505f600190505b828110156107245781856040516020016106c8929190611793565b6040516020818303038152906040529150818682815181106106ed576106ec611110565b5b6020026020010151604051602001610706929190611793565b604051602081830303815290604052915080806001019150506106ad565b50809250505092915050565b610738610bee565b5f60f373ffffffffffffffffffffffffffffffffffffffff166348ad3c3a845f015160038151811061076d5761076c611110565b5b6020026020010151602001516040518263ffffffff1660e01b815260040161079591906112d9565b5f60405180830381865afa1580156107af573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f820116820180604052508101906107d79190611557565b90505f60f373ffffffffffffffffffffffffffffffffffffffff166385989f68835f015160028151811061080e5761080d611110565b5b6020026020010151602001516040518263ffffffff1660e01b815260040161083691906112d9565b602060405180830381865afa158015610851573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061087591906117e0565b90505f60f373ffffffffffffffffffffffffffffffffffffffff16635922f631845f01516001815181106108ac576108ab611110565b5b6020026020010151602001516040518263ffffffff1660e01b81526004016108d491906112d9565b602060405180830381865afa1580156108ef573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061091391906117e0565b90505f60f373ffffffffffffffffffffffffffffffffffffffff166385989f68855f01515f8151811061094957610948611110565b5b6020026020010151602001516040518263ffffffff1660e01b815260040161097191906112d9565b602060405180830381865afa15801561098c573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906109b091906117e0565b90505f60f373ffffffffffffffffffffffffffffffffffffffff166385989f68885f01516002815181106109e7576109e6611110565b5b6020026020010151602001516040518263ffffffff1660e01b8152600401610a0f91906112d9565b602060405180830381865afa158015610a2a573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610a4e91906117e0565b90505f60f373ffffffffffffffffffffffffffffffffffffffff166385989f68895f01515f81518110610a8457610a83611110565b5b6020026020010151602001516040518263ffffffff1660e01b8152600401610aac91906112d9565b602060405180830381865afa158015610ac7573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610aeb91906117e0565b90505f60f373ffffffffffffffffffffffffffffffffffffffff166385989f688a5f0151600181518110610b2257610b21611110565b5b6020026020010151602001516040518263ffffffff1660e01b8152600401610b4a91906112d9565b602060405180830381865afa158015610b65573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610b8991906117e0565b90506040518060c001604052808781526020018681526020018567ffffffffffffffff1681526020018467ffffffffffffffff1681526020018367ffffffffffffffff1681526020018267ffffffffffffffff16815250975050505050505050919050565b6040518060c001604052805f81526020015f81526020015f67ffffffffffffffff1681526020015f67ffffffffffffffff1681526020015f67ffffffffffffffff1681526020015f67ffffffffffffffff1681525090565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610ca582610c5f565b810181811067ffffffffffffffff82111715610cc457610cc3610c6f565b5b80604052505050565b5f610cd6610c46565b9050610ce28282610c9c565b919050565b5f67ffffffffffffffff821115610d0157610d00610c6f565b5b610d0a82610c5f565b9050602081019050919050565b828183375f83830152505050565b5f610d37610d3284610ce7565b610ccd565b905082815260208101848484011115610d5357610d52610c5b565b5b610d5e848285610d17565b509392505050565b5f82601f830112610d7a57610d79610c57565b5b8135610d8a848260208601610d25565b91505092915050565b5f60208284031215610da857610da7610c4f565b5b5f82013567ffffffffffffffff811115610dc557610dc4610c53565b5b610dd184828501610d66565b91505092915050565b5f819050919050565b610dec81610dda565b82525050565b5f67ffffffffffffffff82169050919050565b610e0e81610df2565b82525050565b60c082015f820151610e285f850182610de3565b506020820151610e3b6020850182610de3565b506040820151610e4e6040850182610e05565b506060820151610e616060850182610e05565b506080820151610e746080850182610e05565b5060a0820151610e8760a0850182610e05565b50505050565b5f60c082019050610ea05f830184610e14565b92915050565b5f67ffffffffffffffff821115610ec057610ebf610c6f565b5b602082029050602081019050919050565b5f80fd5b5f610ee7610ee284610ea6565b610ccd565b90508083825260208201905060208402830185811115610f0a57610f09610ed1565b5b835b81811015610f5157803567ffffffffffffffff811115610f2f57610f2e610c57565b5b808601610f3c8982610d66565b85526020850194505050602081019050610f0c565b5050509392505050565b5f82601f830112610f6f57610f6e610c57565b5b8135610f7f848260208601610ed5565b91505092915050565b5f60208284031215610f9d57610f9c610c4f565b5b5f82013567ffffffffffffffff811115610fba57610fb9610c53565b5b610fc684828501610f5b565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b60c082015f82015161100c5f850182610de3565b50602082015161101f6020850182610de3565b5060408201516110326040850182610e05565b5060608201516110456060850182610e05565b5060808201516110586080850182610e05565b5060a082015161106b60a0850182610e05565b50505050565b5f61107c8383610ff8565b60c08301905092915050565b5f602082019050919050565b5f61109e82610fcf565b6110a88185610fd9565b93506110b383610fe9565b805f5b838110156110e35781516110ca8882611071565b97506110d583611088565b9250506001810190506110b6565b5085935050505092915050565b5f6020820190508181035f8301526111088184611094565b905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f61116f8261113d565b6111798185611147565b9350611189818560208601611157565b61119281610c5f565b840191505092915050565b5f6040820190508181035f8301526111b58185611165565b905081810360208301526111c98184611165565b90509392505050565b5f6111e46111df84610ce7565b610ccd565b905082815260208101848484011115611200576111ff610c5b565b5b61120b848285611157565b509392505050565b5f82601f83011261122757611226610c57565b5b81516112378482602086016111d2565b91505092915050565b5f6020828403121561125557611254610c4f565b5b5f82015167ffffffffffffffff81111561127257611271610c53565b5b61127e84828501611213565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f6112ab82611287565b6112b58185611291565b93506112c5818560208601611157565b6112ce81610c5f565b840191505092915050565b5f6020820190508181035f8301526112f181846112a1565b905092915050565b5f80fd5b5f80fd5b5f67ffffffffffffffff82111561131b5761131a610c6f565b5b602082029050602081019050919050565b5f67ffffffffffffffff82111561134657611345610c6f565b5b61134f82610c5f565b9050602081019050919050565b5f61136e6113698461132c565b610ccd565b90508281526020810184848401111561138a57611389610c5b565b5b611395848285611157565b509392505050565b5f82601f8301126113b1576113b0610c57565b5b81516113c184826020860161135c565b91505092915050565b5f604082840312156113df576113de6112f9565b5b6113e96040610ccd565b90505f82015167ffffffffffffffff811115611408576114076112fd565b5b61141484828501611213565b5f83015250602082015167ffffffffffffffff811115611437576114366112fd565b5b6114438482850161139d565b60208301525092915050565b5f61146161145c84611301565b610ccd565b9050808382526020820190506020840283018581111561148457611483610ed1565b5b835b818110156114cb57805167ffffffffffffffff8111156114a9576114a8610c57565b5b8086016114b689826113ca565b85526020850194505050602081019050611486565b5050509392505050565b5f82601f8301126114e9576114e8610c57565b5b81516114f984826020860161144f565b91505092915050565b5f60208284031215611517576115166112f9565b5b6115216020610ccd565b90505f82015167ffffffffffffffff8111156115405761153f6112fd565b5b61154c848285016114d5565b5f8301525092915050565b5f6020828403121561156c5761156b610c4f565b5b5f82015167ffffffffffffffff81111561158957611588610c53565b5b61159584828501611502565b91505092915050565b7f7b2263757272656e63795f706169725f696473223a5b22000000000000000000815250565b5f81905092915050565b5f6115d88261113d565b6115e281856115c4565b93506115f2818560208601611157565b80840191505092915050565b5f6116088261159e565b60178201915061161882846115ce565b915081905092915050565b7f225d7d0000000000000000000000000000000000000000000000000000000000815250565b5f61165482846115ce565b915061165f82611623565b60038201915081905092915050565b5f67ffffffffffffffff82111561168857611687610c6f565b5b602082029050602081019050919050565b5f6116ab6116a68461166e565b610ccd565b905080838252602082019050602084028301858111156116ce576116cd610ed1565b5b835b8181101561171557805167ffffffffffffffff8111156116f3576116f2610c57565b5b808601611700898261139d565b855260208501945050506020810190506116d0565b5050509392505050565b5f82601f83011261173357611732610c57565b5b8151611743848260208601611699565b91505092915050565b5f6020828403121561176157611760610c4f565b5b5f82015167ffffffffffffffff81111561177e5761177d610c53565b5b61178a8482850161171f565b91505092915050565b5f61179e82856115ce565b91506117aa82846115ce565b91508190509392505050565b6117bf81610dda565b81146117c9575f80fd5b50565b5f815190506117da816117b6565b92915050565b5f602082840312156117f5576117f4610c4f565b5b5f611802848285016117cc565b9150509291505056fe2f636f6e6e6563742e6f7261636c652e76322e51756572792f47657450726963652f636f6e6e6563742e6f7261636c652e76322e51756572792f476574507269636573a164736f6c6343000819000a",
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
