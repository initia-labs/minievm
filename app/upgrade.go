package app

import (
	"context"
	"strings"

	"github.com/ethereum/go-ethereum/core/vm"

	upgradetypes "cosmossdk.io/x/upgrade/types"
	"github.com/cosmos/cosmos-sdk/types/module"
)

const upgradeName = "0.6.0"

// RegisterUpgradeHandlers returns upgrade handlers
func (app *MinitiaApp) RegisterUpgradeHandlers(cfg module.Configurator) {
	app.UpgradeKeeper.SetUpgradeHandler(
		upgradeName,
		func(ctx context.Context, _ upgradetypes.Plan, versionMap module.VersionMap) (module.VersionMap, error) {
			// deploy and store erc20 wrapper contract address
			if err := app.EVMKeeper.DeployERC20Wrapper(ctx); err != nil &&
				// ignore contract address collision error (contract already deployed)
				!strings.Contains(err.Error(), vm.ErrContractAddressCollision.Error()) {
				return nil, err
			}

			return versionMap, nil
		},
	)
}
