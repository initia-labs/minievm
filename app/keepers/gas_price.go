package keepers

import (
	"context"
	"errors"

	"cosmossdk.io/collections"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type opchildKeeperForGasPriceKeeper interface {
	MinGasPrices(ctx context.Context) (sdk.DecCoins, error)
}

type GasPriceKeeper struct {
	opck opchildKeeperForGasPriceKeeper
}

func newGasPriceKeeper(opck opchildKeeperForGasPriceKeeper) GasPriceKeeper {
	return GasPriceKeeper{
		opck: opck,
	}
}

func (k GasPriceKeeper) GasPrice(ctx context.Context, denom string) (math.LegacyDec, error) {
	gasPrices, err := k.opck.MinGasPrices(ctx)
	if err != nil && errors.Is(err, collections.ErrNotFound) {
		// allow this case due to init_genesis ordering
		return math.LegacyZeroDec(), nil
	} else if err != nil {
		return math.LegacyZeroDec(), err
	}

	return gasPrices.AmountOf(denom), nil
}
