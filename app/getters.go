package app

import (
	storetypes "cosmossdk.io/store/types"
	upgradekeeper "cosmossdk.io/x/upgrade/keeper"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"

	capabilitykeeper "github.com/cosmos/ibc-go/modules/capability/keeper"
	icacontrollerkeeper "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/controller/keeper"
	ibckeeper "github.com/cosmos/ibc-go/v8/modules/core/keeper"

	ibctestingtypes "github.com/initia-labs/initia/x/ibc/testing/types"
	icaauthkeeper "github.com/initia-labs/initia/x/intertx/keeper"
	evmindexer "github.com/initia-labs/minievm/indexer"
	evmkeeper "github.com/initia-labs/minievm/x/evm/keeper"
)

// GetBaseApp returns the base app for the app.
func (app *MinitiaApp) GetBaseApp() *baseapp.BaseApp {
	return app.BaseApp
}

// GetAccountKeeper returns the account keeper for the app.
func (app *MinitiaApp) GetAccountKeeper() *authkeeper.AccountKeeper {
	return app.AccountKeeper
}

// GetEVMKeeper returns the evm keeper for the app.
func (app *MinitiaApp) GetEVMKeeper() *evmkeeper.Keeper {
	return app.EVMKeeper
}

// GetUpgradeKeeper returns the upgrade keeper for the app.
func (app *MinitiaApp) GetUpgradeKeeper() *upgradekeeper.Keeper {
	return app.UpgradeKeeper
}

// GetStakingKeeper returns the staking keeper for the app.
// It returns opchild instead of original staking keeper.
func (app *MinitiaApp) GetStakingKeeper() ibctestingtypes.StakingKeeper {
	return app.OPChildKeeper
}

// GetIBCKeeper returns the ibc keeper for the app.
func (app *MinitiaApp) GetIBCKeeper() *ibckeeper.Keeper {
	return app.IBCKeeper
}

// GetICAControllerKeeper returns the ica controller keeper for the app.
func (app *MinitiaApp) GetICAControllerKeeper() *icacontrollerkeeper.Keeper {
	return app.ICAControllerKeeper
}

// GetICAAuthKeeper returns the ica auth keeper for the app.
func (app *MinitiaApp) GetICAAuthKeeper() *icaauthkeeper.Keeper {
	return app.ICAAuthKeeper
}

// GetScopedIBCKeeper returns the scoped ibc keeper for the app.
func (app *MinitiaApp) GetScopedIBCKeeper() capabilitykeeper.ScopedKeeper {
	return app.ScopedIBCKeeper
}

// TxConfig returns the tx config for the app.
func (app *MinitiaApp) TxConfig() client.TxConfig {
	return app.txConfig
}

// GetConfigurator returns the configurator for the app.
func (app *MinitiaApp) GetConfigurator() module.Configurator {
	return app.configurator
}

// GetModuleManager returns the module manager for the app.
func (app *MinitiaApp) GetModuleManager() *module.Manager {
	return app.ModuleManager
}

// IndexerKeeper returns the evm indexer
func (app *MinitiaApp) EVMIndexer() evmindexer.EVMIndexer {
	return app.evmIndexer
}

// CheckStateContextGetter returns a function that returns a new Context for state checking.
func (app *MinitiaApp) CheckStateContextGetter() func() sdk.Context {
	return func() sdk.Context {
		return app.GetContextForCheckTx(nil)
	}
}

func (app *MinitiaApp) GetQueryMultiStore() storetypes.MultiStore {
	return app.qms
}
