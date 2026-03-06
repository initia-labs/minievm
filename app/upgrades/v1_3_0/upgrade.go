package v1_3_0

import (
	"context"

	storetypes "cosmossdk.io/store/types"
	upgradetypes "cosmossdk.io/x/upgrade/types"

	tmos "github.com/cometbft/cometbft/libs/os"
	"github.com/cosmos/cosmos-sdk/types/module"

	"github.com/initia-labs/minievm/app/upgrades"
)

const upgradeName = "v1.3.0"

// RegisterUpgradeHandlers registers the upgrade handlers for the app.
func RegisterUpgradeHandlers(app upgrades.MinitiaApp) {
	// apply store upgrade only if this upgrade is scheduled at a height
	if upgradeInfo, err := app.GetUpgradeKeeper().ReadUpgradeInfoFromDisk(); err == nil {
		if upgradeInfo.Name == upgradeName && !app.GetUpgradeKeeper().IsSkipHeight(upgradeInfo.Height) {
			storeUpgrades := storetypes.StoreUpgrades{
				Deleted: []string{"auction"},
			}

			app.SetStoreLoader(upgradetypes.UpgradeStoreLoader(upgradeInfo.Height, &storeUpgrades))
		}
	} else {
		tmos.Exit(err.Error())
	}

	app.GetUpgradeKeeper().SetUpgradeHandler(
		upgradeName,
		func(ctx context.Context, _ upgradetypes.Plan, versionMap module.VersionMap) (module.VersionMap, error) {
			if err := upgrades.BindOPInitPort(ctx, app); err != nil {
				return nil, err
			}

			if err := upgrades.NormalizeEVMParams(ctx, app); err != nil {
				return nil, err
			}

			if err := upgrades.UpdateERC20WrapperContract(ctx, app); err != nil {
				return nil, err
			}

			if err := upgrades.UpdateERC20FactoryContract(ctx, app); err != nil {
				return nil, err
			}

			return app.GetModuleManager().RunMigrations(ctx, app.GetConfigurator(), versionMap)
		},
	)
}
