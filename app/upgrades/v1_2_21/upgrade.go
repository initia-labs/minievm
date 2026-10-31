package v1_2_21

import (
	"context"
	"fmt"
	"time"

	upgradetypes "cosmossdk.io/x/upgrade/types"

	abci "github.com/cometbft/cometbft/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"

	"github.com/initia-labs/minievm/app/upgrades"
)

const UpgradeName = "v1.2.21"

var (
	RestartChains = map[string]bool{
		"cabal-1":   true,
		"yominet-1": true,
	}
	RestartAfterRFC3339 = "2026-11-01T00:00:00Z"
	RestartAfter        = mustParseTime(RestartAfterRFC3339)
)

func mustParseTime(s string) time.Time {
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		panic(err)
	}
	return t
}

// RegisterUpgradeHandlers returns upgrade handlers
func RegisterUpgradeHandlers(app upgrades.MinitiaApp) {
	app.GetUpgradeKeeper().SetUpgradeHandler(
		UpgradeName,
		func(ctx context.Context, _ upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
			if err := run(ctx, app); err != nil {
				return nil, err
			}

			return vm, nil
		},
	)
}

// WrapPreBlocker applies the upgrade once, in the first block at or after RestartAfter on a
// restart chain, before the module pre-blockers run. ApplyUpgrade records the done height and
// clears the plan, so the x/upgrade pre-blocker that follows sees nothing to do.
func WrapPreBlocker(app upgrades.MinitiaApp, next sdk.PreBlocker) sdk.PreBlocker {
	return func(ctx sdk.Context, req *abci.RequestFinalizeBlock) (*sdk.ResponsePreBlock, error) {
		if RestartChains[ctx.ChainID()] && !ctx.BlockTime().Before(RestartAfter) {
			done, err := app.GetUpgradeKeeper().GetDoneHeight(ctx, UpgradeName)
			if err != nil {
				return nil, err
			}

			if done == 0 {
				if !app.GetUpgradeKeeper().HasHandler(UpgradeName) {
					return nil, fmt.Errorf("upgrade handler %s is not registered", UpgradeName)
				}

				ctx.Logger().Info("applying restart upgrade", "name", UpgradeName, "height", ctx.BlockHeight(), "time", ctx.BlockTime())
				if err := app.GetUpgradeKeeper().ApplyUpgrade(ctx, upgradetypes.Plan{
					Name:   UpgradeName,
					Height: ctx.BlockHeight(),
					Info:   "rollup restart after the 2026-10-23 interwoven-1 halt",
				}); err != nil {
					return nil, err
				}
			}
		}

		return next(ctx, req)
	}
}
