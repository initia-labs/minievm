package app

import (
	"context"

	upgradetypes "cosmossdk.io/x/upgrade/types"
	"github.com/cosmos/cosmos-sdk/types/module"

	v057 "github.com/initia-labs/minievm/x/evm/migrations/v057"
)

const upgradeName = "0.5.7"

// RegisterUpgradeHandlers returns upgrade handlers
func (app *MinitiaApp) RegisterUpgradeHandlers(cfg module.Configurator) {
	app.UpgradeKeeper.SetUpgradeHandler(
		upgradeName,
		func(ctx context.Context, _ upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
			// update evm params
			err := v057.MigrateParams(ctx, app.AppCodec(), app.EVMKeeper.StoreService())
			if err != nil {
				return nil, err
			}

			// deploy and store erc20 wrapper contract address
			err = app.EVMKeeper.DeployERC20Wrapper(ctx)
			if err != nil {
				return nil, err
			}

			return vm, nil
		},
	)
}
