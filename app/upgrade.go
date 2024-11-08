package app

import (
	"context"
	"strings"

	"github.com/ethereum/go-ethereum/core/vm"

	upgradetypes "cosmossdk.io/x/upgrade/types"
	"github.com/cosmos/cosmos-sdk/types/module"

	opchildtypes "github.com/initia-labs/OPinit/x/opchild/types"
)

const upgradeName = "0.6.4"

// RegisterUpgradeHandlers returns upgrade handlers
func (app *MinitiaApp) RegisterUpgradeHandlers(cfg module.Configurator) {
	app.UpgradeKeeper.SetUpgradeHandler(
		upgradeName,
		func(ctx context.Context, _ upgradetypes.Plan, versionMap module.VersionMap) (module.VersionMap, error) {
			// deploy and store erc20 factory contract address
			if err := app.EVMKeeper.DeployERC20Factory(ctx); err != nil &&
				// ignore contract address collision error (contract already deployed)
				!strings.Contains(err.Error(), vm.ErrContractAddressCollision.Error()) {
				return nil, err
			}

			// deploy and store erc20 wrapper contract address
			if err := app.EVMKeeper.DeployERC20Wrapper(ctx); err != nil &&
				// ignore contract address collision error (contract already deployed)
				!strings.Contains(err.Error(), vm.ErrContractAddressCollision.Error()) {
				return nil, err
			}

			// opchild params update
			params, err := app.OPChildKeeper.GetParams(ctx)
			if err != nil {
				return nil, err
			}

			// set non-zero default values for new params
			params.HookMaxGas = opchildtypes.DefaultHookMaxGas

			err = app.OPChildKeeper.SetParams(ctx, params)
			if err != nil {
				return nil, err
			}

			return versionMap, nil
		},
	)
}
