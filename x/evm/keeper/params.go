package keeper

import (
	"context"

	"cosmossdk.io/math"
)

func (k Keeper) GetFeeDenom(ctx context.Context) (string, error) {
	params, err := k.Params.Get(ctx)
	if err != nil {
		return "", err
	}

	return params.FeeDenom, nil
}

func (k Keeper) GasRefundRatio(ctx context.Context) (math.LegacyDec, error) {
	params, err := k.Params.Get(ctx)
	if err != nil {
		return math.LegacyZeroDec(), err
	}

	return params.GasRefundRatio, nil
}

func (k Keeper) NumRetainBlockHashes(ctx context.Context) (uint64, error) {
	params, err := k.Params.Get(ctx)
	if err != nil {
		return 0, err
	}

	return params.NumRetainBlockHashes, nil
}
