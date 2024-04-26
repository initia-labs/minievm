package erc721registryprecompile

import (
	storetypes "cosmossdk.io/store/types"

	"github.com/ethereum/go-ethereum/common"
)

type RegisterStoreArguments struct {
	Account common.Address `abi:"account"`
}

type IsStoreRegisteredArguments struct {
	Account common.Address `abi:"account"`
}

const (
	REGISTER_GAS            storetypes.Gas = 200
	REGISTER_STORE_GAS      storetypes.Gas = 200
	IS_STORE_REGISTERED_GAS storetypes.Gas = 200
	GAS_PER_BYTE            storetypes.Gas = 1
)
