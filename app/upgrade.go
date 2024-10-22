package app

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	upgradetypes "cosmossdk.io/x/upgrade/types"
	"github.com/cosmos/cosmos-sdk/types/module"

	evmtypes "github.com/initia-labs/minievm/x/evm/types"
)

const upgradeName = "0.5.7"

// RegisterUpgradeHandlers returns upgrade handlers
func (app *MinitiaApp) RegisterUpgradeHandlers(cfg module.Configurator) {
	app.UpgradeKeeper.SetUpgradeHandler(
		upgradeName,
		func(ctx context.Context, _ upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
			// update evm params
			params, err := app.EVMKeeper.Params.Get(ctx)
			if err != nil {
				return nil, errorsmod.Wrap(err, "failed to get evm params")
			}
			params.GasRefundRatio = evmtypes.DefaultParams().GasRefundRatio
			params.NumRetainBlockHashes = 0
			err = app.EVMKeeper.Params.Set(ctx, params)
			if err != nil {
				return nil, errorsmod.Wrap(err, "failed to set evm params")
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
