package cosmosprecompile

import (
	storetypes "cosmossdk.io/store/types"

	"github.com/ethereum/go-ethereum/common"
)

// ToCosmosAddressArguments is the arguments for the to_cosmos_address method.
type ToCosmosAddressArguments struct {
	EVMAddress common.Address `abi:"evm_address"`
}

// ToEVMAddressArguments is the arguments for the to_evm_address method.
type ToEVMAddressArguments struct {
	CosmosAddress string `abi:"cosmos_address"`
}

// ExecuteCosmosArguments is the arguments for the execute_cosmos method.
type ExecuteCosmosArguments struct {
	Msg string `abi:"msg"`
}

// QueryCosmosArguments is the arguments for the query_cosmos method.
type QueryCosmosArguments struct {
	Path string `abi:"path"`
	Req  string `abi:"req"`
}

// ToDenomArguments is the arguments for the to_denom method.
type ToDenomArguments struct {
	ERC20Address common.Address `abi:"erc20_address"`
}

// ToERC20Arguments is the arguments for the to_erc20 method.
type ToERC20Arguments struct {
	Denom string `abi:"denom"`
}

const (
	TO_COSMOS_ADDRESS_GAS storetypes.Gas = 200
	TO_EVM_ADDRESS_GAS    storetypes.Gas = 200
	TO_DENOM_GAS          storetypes.Gas = 100
	TO_ERC20_GAS          storetypes.Gas = 100
	EXECUTE_COSMOS_GAS    storetypes.Gas = 200
	QUERY_COSMOS_GAS      storetypes.Gas = 200
	GAS_PER_BYTE          storetypes.Gas = 1
)

const (
	METHOD_TO_COSMOS_ADDRESS = "to_cosmos_address"
	METHOD_TO_EVM_ADDRESS    = "to_evm_address"
	METHOD_EXECUTE_COSMOS    = "execute_cosmos"
	METHOD_QUERY_COSMOS      = "query_cosmos"
	METHOD_TO_DENOM          = "to_denom"
	METHOD_TO_ERC20          = "to_erc20"
)
