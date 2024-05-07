package erc20registryprecompile

import (
	storetypes "cosmossdk.io/store/types"

	"github.com/ethereum/go-ethereum/common"
)

type RegisterERC20FromFactoryArguments struct {
	ERC20 common.Address `abi:"erc20"`
}

type RegisterStoreArguments struct {
	Account common.Address `abi:"account"`
}

type IsStoreRegisteredArguments struct {
	Account common.Address `abi:"account"`
}

const (
	REGISTER_GAS              storetypes.Gas = 200
	REGISTER_FROM_FACTORY_GAS storetypes.Gas = 200
	REGISTER_STORE_GAS        storetypes.Gas = 200
	IS_STORE_REGISTERED_GAS   storetypes.Gas = 200
	GAS_PER_BYTE              storetypes.Gas = 1
)
