package app

import (
	"context"

	upgradetypes "cosmossdk.io/x/upgrade/types"
	"github.com/cosmos/cosmos-sdk/types/module"
)

const upgradeName = "0.5.14"

// RegisterUpgradeHandlers returns upgrade handlers
func (app *MinitiaApp) RegisterUpgradeHandlers(cfg module.Configurator) {
	app.UpgradeKeeper.SetUpgradeHandler(
		upgradeName,
		func(ctx context.Context, _ upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
			// deploy and store erc20 wrapper contract address
			err := app.EVMKeeper.DeployERC20Wrapper(ctx)
			if err != nil {
				return nil, err
			}

			return vm, nil
		},
	)
}
