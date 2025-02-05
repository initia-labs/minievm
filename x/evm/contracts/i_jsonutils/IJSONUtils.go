// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package i_jsonutils

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

// IJSONUtilsJSONElement is an auto generated low-level Go binding around an user-defined struct.
type IJSONUtilsJSONElement struct {
	Key   string
	Value []byte
}

// IJSONUtilsJSONObject is an auto generated low-level Go binding around an user-defined struct.
type IJSONUtilsJSONObject struct {
	Elements []IJSONUtilsJSONElement
}

// IJsonutilsMetaData contains all meta data concerning the IJsonutils contract.
var IJsonutilsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dst_json\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"src_json\",\"type\":\"string\"}],\"name\":\"merge_json\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"json\",\"type\":\"string\"}],\"name\":\"stringify_json\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"json_bytes\",\"type\":\"bytes\"}],\"name\":\"unmarshal_iso_to_unix\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"json_bytes\",\"type\":\"bytes\"}],\"name\":\"unmarshal_to_array\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"json_bytes\",\"type\":\"bytes\"}],\"name\":\"unmarshal_to_bool\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"json_bytes\",\"type\":\"bytes\"}],\"name\":\"unmarshal_to_object\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"internalType\":\"structIJSONUtils.JSONElement[]\",\"name\":\"elements\",\"type\":\"tuple[]\"}],\"internalType\":\"structIJSONUtils.JSONObject\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"json_bytes\",\"type\":\"bytes\"}],\"name\":\"unmarshal_to_string\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"json_bytes\",\"type\":\"bytes\"}],\"name\":\"unmarshal_to_uint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IJsonutilsABI is the input ABI used to generate the binding from.
// Deprecated: Use IJsonutilsMetaData.ABI instead.
var IJsonutilsABI = IJsonutilsMetaData.ABI

// IJsonutils is an auto generated Go binding around an Ethereum contract.
type IJsonutils struct {
	IJsonutilsCaller     // Read-only binding to the contract
	IJsonutilsTransactor // Write-only binding to the contract
	IJsonutilsFilterer   // Log filterer for contract events
}

// IJsonutilsCaller is an auto generated read-only Go binding around an Ethereum contract.
type IJsonutilsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IJsonutilsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IJsonutilsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IJsonutilsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IJsonutilsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IJsonutilsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IJsonutilsSession struct {
	Contract     *IJsonutils       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IJsonutilsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IJsonutilsCallerSession struct {
	Contract *IJsonutilsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// IJsonutilsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IJsonutilsTransactorSession struct {
	Contract     *IJsonutilsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// IJsonutilsRaw is an auto generated low-level Go binding around an Ethereum contract.
type IJsonutilsRaw struct {
	Contract *IJsonutils // Generic contract binding to access the raw methods on
}

// IJsonutilsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IJsonutilsCallerRaw struct {
	Contract *IJsonutilsCaller // Generic read-only contract binding to access the raw methods on
}

// IJsonutilsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IJsonutilsTransactorRaw struct {
	Contract *IJsonutilsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIJsonutils creates a new instance of IJsonutils, bound to a specific deployed contract.
func NewIJsonutils(address common.Address, backend bind.ContractBackend) (*IJsonutils, error) {
	contract, err := bindIJsonutils(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IJsonutils{IJsonutilsCaller: IJsonutilsCaller{contract: contract}, IJsonutilsTransactor: IJsonutilsTransactor{contract: contract}, IJsonutilsFilterer: IJsonutilsFilterer{contract: contract}}, nil
}

// NewIJsonutilsCaller creates a new read-only instance of IJsonutils, bound to a specific deployed contract.
func NewIJsonutilsCaller(address common.Address, caller bind.ContractCaller) (*IJsonutilsCaller, error) {
	contract, err := bindIJsonutils(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IJsonutilsCaller{contract: contract}, nil
}

// NewIJsonutilsTransactor creates a new write-only instance of IJsonutils, bound to a specific deployed contract.
func NewIJsonutilsTransactor(address common.Address, transactor bind.ContractTransactor) (*IJsonutilsTransactor, error) {
	contract, err := bindIJsonutils(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IJsonutilsTransactor{contract: contract}, nil
}

// NewIJsonutilsFilterer creates a new log filterer instance of IJsonutils, bound to a specific deployed contract.
func NewIJsonutilsFilterer(address common.Address, filterer bind.ContractFilterer) (*IJsonutilsFilterer, error) {
	contract, err := bindIJsonutils(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IJsonutilsFilterer{contract: contract}, nil
}

// bindIJsonutils binds a generic wrapper to an already deployed contract.
func bindIJsonutils(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IJsonutilsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IJsonutils *IJsonutilsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IJsonutils.Contract.IJsonutilsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IJsonutils *IJsonutilsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IJsonutils.Contract.IJsonutilsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IJsonutils *IJsonutilsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IJsonutils.Contract.IJsonutilsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IJsonutils *IJsonutilsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IJsonutils.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IJsonutils *IJsonutilsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IJsonutils.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IJsonutils *IJsonutilsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IJsonutils.Contract.contract.Transact(opts, method, params...)
}

// MergeJson is a free data retrieval call binding the contract method 0x5cc855e3.
//
// Solidity: function merge_json(string dst_json, string src_json) view returns(string)
func (_IJsonutils *IJsonutilsCaller) MergeJson(opts *bind.CallOpts, dst_json string, src_json string) (string, error) {
	var out []interface{}
	err := _IJsonutils.contract.Call(opts, &out, "merge_json", dst_json, src_json)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// MergeJson is a free data retrieval call binding the contract method 0x5cc855e3.
//
// Solidity: function merge_json(string dst_json, string src_json) view returns(string)
func (_IJsonutils *IJsonutilsSession) MergeJson(dst_json string, src_json string) (string, error) {
	return _IJsonutils.Contract.MergeJson(&_IJsonutils.CallOpts, dst_json, src_json)
}

// MergeJson is a free data retrieval call binding the contract method 0x5cc855e3.
//
// Solidity: function merge_json(string dst_json, string src_json) view returns(string)
func (_IJsonutils *IJsonutilsCallerSession) MergeJson(dst_json string, src_json string) (string, error) {
	return _IJsonutils.Contract.MergeJson(&_IJsonutils.CallOpts, dst_json, src_json)
}

// StringifyJson is a free data retrieval call binding the contract method 0x8d5c8817.
//
// Solidity: function stringify_json(string json) view returns(string)
func (_IJsonutils *IJsonutilsCaller) StringifyJson(opts *bind.CallOpts, json string) (string, error) {
	var out []interface{}
	err := _IJsonutils.contract.Call(opts, &out, "stringify_json", json)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// StringifyJson is a free data retrieval call binding the contract method 0x8d5c8817.
//
// Solidity: function stringify_json(string json) view returns(string)
func (_IJsonutils *IJsonutilsSession) StringifyJson(json string) (string, error) {
	return _IJsonutils.Contract.StringifyJson(&_IJsonutils.CallOpts, json)
}

// StringifyJson is a free data retrieval call binding the contract method 0x8d5c8817.
//
// Solidity: function stringify_json(string json) view returns(string)
func (_IJsonutils *IJsonutilsCallerSession) StringifyJson(json string) (string, error) {
	return _IJsonutils.Contract.StringifyJson(&_IJsonutils.CallOpts, json)
}

// UnmarshalIsoToUnix is a free data retrieval call binding the contract method 0x5922f631.
//
// Solidity: function unmarshal_iso_to_unix(bytes json_bytes) view returns(uint256)
func (_IJsonutils *IJsonutilsCaller) UnmarshalIsoToUnix(opts *bind.CallOpts, json_bytes []byte) (*big.Int, error) {
	var out []interface{}
	err := _IJsonutils.contract.Call(opts, &out, "unmarshal_iso_to_unix", json_bytes)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnmarshalIsoToUnix is a free data retrieval call binding the contract method 0x5922f631.
//
// Solidity: function unmarshal_iso_to_unix(bytes json_bytes) view returns(uint256)
func (_IJsonutils *IJsonutilsSession) UnmarshalIsoToUnix(json_bytes []byte) (*big.Int, error) {
	return _IJsonutils.Contract.UnmarshalIsoToUnix(&_IJsonutils.CallOpts, json_bytes)
}

// UnmarshalIsoToUnix is a free data retrieval call binding the contract method 0x5922f631.
//
// Solidity: function unmarshal_iso_to_unix(bytes json_bytes) view returns(uint256)
func (_IJsonutils *IJsonutilsCallerSession) UnmarshalIsoToUnix(json_bytes []byte) (*big.Int, error) {
	return _IJsonutils.Contract.UnmarshalIsoToUnix(&_IJsonutils.CallOpts, json_bytes)
}

// UnmarshalToArray is a free data retrieval call binding the contract method 0x7e8fa9cd.
//
// Solidity: function unmarshal_to_array(bytes json_bytes) view returns(bytes[])
func (_IJsonutils *IJsonutilsCaller) UnmarshalToArray(opts *bind.CallOpts, json_bytes []byte) ([][]byte, error) {
	var out []interface{}
	err := _IJsonutils.contract.Call(opts, &out, "unmarshal_to_array", json_bytes)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

// UnmarshalToArray is a free data retrieval call binding the contract method 0x7e8fa9cd.
//
// Solidity: function unmarshal_to_array(bytes json_bytes) view returns(bytes[])
func (_IJsonutils *IJsonutilsSession) UnmarshalToArray(json_bytes []byte) ([][]byte, error) {
	return _IJsonutils.Contract.UnmarshalToArray(&_IJsonutils.CallOpts, json_bytes)
}

// UnmarshalToArray is a free data retrieval call binding the contract method 0x7e8fa9cd.
//
// Solidity: function unmarshal_to_array(bytes json_bytes) view returns(bytes[])
func (_IJsonutils *IJsonutilsCallerSession) UnmarshalToArray(json_bytes []byte) ([][]byte, error) {
	return _IJsonutils.Contract.UnmarshalToArray(&_IJsonutils.CallOpts, json_bytes)
}

// UnmarshalToBool is a free data retrieval call binding the contract method 0x75d58a65.
//
// Solidity: function unmarshal_to_bool(bytes json_bytes) view returns(bool)
func (_IJsonutils *IJsonutilsCaller) UnmarshalToBool(opts *bind.CallOpts, json_bytes []byte) (bool, error) {
	var out []interface{}
	err := _IJsonutils.contract.Call(opts, &out, "unmarshal_to_bool", json_bytes)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UnmarshalToBool is a free data retrieval call binding the contract method 0x75d58a65.
//
// Solidity: function unmarshal_to_bool(bytes json_bytes) view returns(bool)
func (_IJsonutils *IJsonutilsSession) UnmarshalToBool(json_bytes []byte) (bool, error) {
	return _IJsonutils.Contract.UnmarshalToBool(&_IJsonutils.CallOpts, json_bytes)
}

// UnmarshalToBool is a free data retrieval call binding the contract method 0x75d58a65.
//
// Solidity: function unmarshal_to_bool(bytes json_bytes) view returns(bool)
func (_IJsonutils *IJsonutilsCallerSession) UnmarshalToBool(json_bytes []byte) (bool, error) {
	return _IJsonutils.Contract.UnmarshalToBool(&_IJsonutils.CallOpts, json_bytes)
}

// UnmarshalToObject is a free data retrieval call binding the contract method 0x48ad3c3a.
//
// Solidity: function unmarshal_to_object(bytes json_bytes) view returns(((string,bytes)[]))
func (_IJsonutils *IJsonutilsCaller) UnmarshalToObject(opts *bind.CallOpts, json_bytes []byte) (IJSONUtilsJSONObject, error) {
	var out []interface{}
	err := _IJsonutils.contract.Call(opts, &out, "unmarshal_to_object", json_bytes)

	if err != nil {
		return *new(IJSONUtilsJSONObject), err
	}

	out0 := *abi.ConvertType(out[0], new(IJSONUtilsJSONObject)).(*IJSONUtilsJSONObject)

	return out0, err

}

// UnmarshalToObject is a free data retrieval call binding the contract method 0x48ad3c3a.
//
// Solidity: function unmarshal_to_object(bytes json_bytes) view returns(((string,bytes)[]))
func (_IJsonutils *IJsonutilsSession) UnmarshalToObject(json_bytes []byte) (IJSONUtilsJSONObject, error) {
	return _IJsonutils.Contract.UnmarshalToObject(&_IJsonutils.CallOpts, json_bytes)
}

// UnmarshalToObject is a free data retrieval call binding the contract method 0x48ad3c3a.
//
// Solidity: function unmarshal_to_object(bytes json_bytes) view returns(((string,bytes)[]))
func (_IJsonutils *IJsonutilsCallerSession) UnmarshalToObject(json_bytes []byte) (IJSONUtilsJSONObject, error) {
	return _IJsonutils.Contract.UnmarshalToObject(&_IJsonutils.CallOpts, json_bytes)
}

// UnmarshalToString is a free data retrieval call binding the contract method 0x532478ad.
//
// Solidity: function unmarshal_to_string(bytes json_bytes) view returns(string)
func (_IJsonutils *IJsonutilsCaller) UnmarshalToString(opts *bind.CallOpts, json_bytes []byte) (string, error) {
	var out []interface{}
	err := _IJsonutils.contract.Call(opts, &out, "unmarshal_to_string", json_bytes)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UnmarshalToString is a free data retrieval call binding the contract method 0x532478ad.
//
// Solidity: function unmarshal_to_string(bytes json_bytes) view returns(string)
func (_IJsonutils *IJsonutilsSession) UnmarshalToString(json_bytes []byte) (string, error) {
	return _IJsonutils.Contract.UnmarshalToString(&_IJsonutils.CallOpts, json_bytes)
}

// UnmarshalToString is a free data retrieval call binding the contract method 0x532478ad.
//
// Solidity: function unmarshal_to_string(bytes json_bytes) view returns(string)
func (_IJsonutils *IJsonutilsCallerSession) UnmarshalToString(json_bytes []byte) (string, error) {
	return _IJsonutils.Contract.UnmarshalToString(&_IJsonutils.CallOpts, json_bytes)
}

// UnmarshalToUint is a free data retrieval call binding the contract method 0x85989f68.
//
// Solidity: function unmarshal_to_uint(bytes json_bytes) view returns(uint256)
func (_IJsonutils *IJsonutilsCaller) UnmarshalToUint(opts *bind.CallOpts, json_bytes []byte) (*big.Int, error) {
	var out []interface{}
	err := _IJsonutils.contract.Call(opts, &out, "unmarshal_to_uint", json_bytes)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnmarshalToUint is a free data retrieval call binding the contract method 0x85989f68.
//
// Solidity: function unmarshal_to_uint(bytes json_bytes) view returns(uint256)
func (_IJsonutils *IJsonutilsSession) UnmarshalToUint(json_bytes []byte) (*big.Int, error) {
	return _IJsonutils.Contract.UnmarshalToUint(&_IJsonutils.CallOpts, json_bytes)
}

// UnmarshalToUint is a free data retrieval call binding the contract method 0x85989f68.
//
// Solidity: function unmarshal_to_uint(bytes json_bytes) view returns(uint256)
func (_IJsonutils *IJsonutilsCallerSession) UnmarshalToUint(json_bytes []byte) (*big.Int, error) {
	return _IJsonutils.Contract.UnmarshalToUint(&_IJsonutils.CallOpts, json_bytes)
}
