package v1_1_9

import (
	"context"

	upgradetypes "cosmossdk.io/x/upgrade/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/initia-labs/minievm/app/upgrades"
	"github.com/initia-labs/minievm/app/upgrades/contracts/erc20_factory"
	"github.com/initia-labs/minievm/app/upgrades/contracts/erc20_wrapper"
)

const upgradeName = "v1.1.9"

// RegisterUpgradeHandlers registers the upgrade handlers for the app.
func RegisterUpgradeHandlers(app upgrades.MinitiaApp) {
	app.GetUpgradeKeeper().SetUpgradeHandler(
		upgradeName,
		func(ctx context.Context, _ upgradetypes.Plan, versionMap module.VersionMap) (module.VersionMap, error) {

			// 1. update erc20 wrapper contract
			wrapperAddr, err := app.GetEVMKeeper().GetERC20WrapperAddr(ctx)
			if err != nil {
				return nil, err
			}
			wrapperRuntimeCode, err := hexutil.Decode(erc20_wrapper.Erc20WrapperBin)
			if err != nil {
				return nil, err
			}
			wrapperCodeHash := upgrades.CodeHash(wrapperRuntimeCode)
			if err := upgrades.ReplaceCodeAndCodeHash(ctx, app, wrapperAddr.Bytes(), wrapperRuntimeCode, wrapperCodeHash); err != nil {
				return nil, err
			}

			// 2. update erc20 factory contract
			factoryAddr, err := app.GetEVMKeeper().GetERC20FactoryAddr(ctx)
			if err != nil {
				return nil, err
			}
			factoryRuntimeCode, err := hexutil.Decode(erc20_factory.Erc20FactoryBin)
			if err != nil {
				return nil, err
			}
			factoryCodeHash := upgrades.CodeHash(factoryRuntimeCode)
			if err := upgrades.ReplaceCodeAndCodeHash(ctx, app, factoryAddr.Bytes(), factoryRuntimeCode, factoryCodeHash); err != nil {
				return nil, err
			}

			// 3. run migrations
			return app.GetModuleManager().RunMigrations(ctx, app.GetConfigurator(), versionMap)
		},
	)
}
