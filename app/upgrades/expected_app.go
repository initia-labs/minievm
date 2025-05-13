package upgrades

import (
	upgradekeeper "cosmossdk.io/x/upgrade/keeper"
	"github.com/cosmos/cosmos-sdk/types/module"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"

	evmkeeper "github.com/initia-labs/minievm/x/evm/keeper"
)

type MinitiaApp interface {
	GetAccountKeeper() *authkeeper.AccountKeeper
	GetEVMKeeper() *evmkeeper.Keeper
	GetUpgradeKeeper() *upgradekeeper.Keeper

	GetConfigurator() module.Configurator
	GetModuleManager() *module.Manager
}
