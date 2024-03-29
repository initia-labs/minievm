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

const (
	TO_COSMOS_ADDRESS_GAS storetypes.Gas = 200
	TO_EVM_ADDRESS_GAS    storetypes.Gas = 200
	EXECUTE_COSMOS_GAS    storetypes.Gas = 200
	QUERY_COSMOS_GAS      storetypes.Gas = 200
	GAS_PER_BYTE          storetypes.Gas = 1
)
