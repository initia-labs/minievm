package cosmosprecompile

import (
	storetypes "cosmossdk.io/store/types"

	"github.com/ethereum/go-ethereum/common"
)

type ToCosmosAddressArguments struct {
	EVMAddress common.Address `abi:"evm_address"`
}

type ToEVMAddressArguments struct {
	CosmosAddress string `abi:"cosmos_address"`
}

const (
	TO_COSMOS_ADDRESS_GAS storetypes.Gas = 1000
	TO_EVM_ADDRESS_GAS    storetypes.Gas = 1000
)
