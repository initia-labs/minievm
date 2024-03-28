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

// ExecuteCosmosMessageArguments is the arguments for the execute_cosmos_message method.
type ExecuteCosmosMessageArguments struct {
	Msg string `abi:"msg"`
}

const (
	TO_COSMOS_ADDRESS_GAS      storetypes.Gas = 200
	TO_EVM_ADDRESS_GAS         storetypes.Gas = 200
	EXECUTE_COSMOS_MESSAGE_GAS storetypes.Gas = 200
	GAS_PER_BYTE               storetypes.Gas = 1
)
