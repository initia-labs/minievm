package types

import (
	"context"

	corestoretypes "cosmossdk.io/core/store"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/initia-labs/minievm/x/evm/types"
)

func MigrateParams(ctx context.Context, cdc codec.Codec, storeService corestoretypes.KVStoreService) error {
	kvStore := storeService.OpenKVStore(ctx)

	legacyParamsBz, err := kvStore.Get(types.ParamsKey)
	if err != nil {
		return err
	}

	var legacyParams Params
	err = cdc.Unmarshal(legacyParamsBz, &legacyParams)
	if err != nil {
		return err
	}

	params := types.Params{
		ExtraEIPs:           legacyParams.ExtraEIPs,
		AllowedPublishers:   legacyParams.AllowedPublishers,
		AllowCustomERC20:    legacyParams.AllowCustomERC20,
		AllowedCustomERC20s: legacyParams.AllowedCustomERC20s,
		FeeDenom:            legacyParams.FeeDenom,
		GasRefundRatio:      types.DefaultParams().GasRefundRatio,
	}

	paramsBz := cdc.MustMarshal(&params)
	err = kvStore.Set(types.ParamsKey, paramsBz)
	if err != nil {
		return err
	}

	return nil
}
