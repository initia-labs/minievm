package v1_2_21

import (
	"context"
	"slices"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/initia-labs/minievm/app/upgrades"
)

func run(ctx context.Context, app upgrades.MinitiaApp) error {
	denoms := app.GetForwardingKeeper().GetAllowedDenoms(ctx)
	if !slices.Contains(denoms, "*") {
		sdk.UnwrapSDKContext(ctx).Logger().Info("forwarding allowed_denoms left as is", "denoms", denoms)
		return nil
	}

	sdk.UnwrapSDKContext(ctx).Logger().Info("forwarding allowed_denoms cleared", "was", denoms)
	return app.GetForwardingKeeper().AllowedDenoms.Clear(ctx, nil)
}
