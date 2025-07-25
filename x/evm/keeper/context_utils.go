package keeper

import (
	"context"
	"errors"
	"math/big"

	"cosmossdk.io/collections"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ethereum/go-ethereum/common"

	"github.com/initia-labs/minievm/x/evm/types"
)

func (k Keeper) LoadFee(ctx context.Context, params types.Params) (types.Fee, error) {
	feeContract, err := types.DenomToContractAddr(ctx, k, params.FeeDenom)
	if err != nil && !errors.Is(err, collections.ErrNotFound) {
		return types.Fee{}, err
	}

	decimals := uint8(0)
	if (feeContract != common.Address{} &&
		// erc20Keeper.Decimals is also calling LoadFee, so we need to check this call is not recursive
		sdk.UnwrapSDKContext(ctx).Value(types.CONTEXT_KEY_LOAD_DECIMALS) == nil) {
		decimals = k.erc20Keeper.Decimals(ctx, feeContract)
	}

	return types.NewFee(params.FeeDenom, feeContract, decimals), nil
}

func (k Keeper) extractGasPriceFromContext(ctx context.Context, fee types.Fee) (*big.Int, error) {
	if (fee.Contract() == common.Address{}) {
		return big.NewInt(0), nil
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	value := sdkCtx.Value(types.CONTEXT_KEY_GAS_PRICES)
	if value == nil {
		return big.NewInt(0), nil
	}

	gasPrices := value.(sdk.DecCoins)
	gasPriceDec := gasPrices.AmountOf(fee.Denom())
	if !gasPriceDec.IsPositive() {
		return big.NewInt(0), nil
	}

	// multiply by 1e9 to prevent decimal drops
	gasPrice := gasPriceDec.
		MulTruncate(math.LegacyNewDec(1e9)).
		TruncateInt().BigInt()

	return types.ToEthersUnit(fee.Decimals()+9, gasPrice), nil
}

func (k Keeper) baseFee(ctx context.Context, fee types.Fee) (*big.Int, error) {
	gasPriceDec, err := k.gasPriceKeeper.GasPrice(ctx, fee.Denom())
	if err != nil {
		return nil, err
	}

	// multiply by 1e9 to prevent decimal drops
	gasPrice := gasPriceDec.
		MulTruncate(math.LegacyNewDec(1e9)).
		TruncateInt().BigInt()

	return types.ToEthersUnit(fee.Decimals()+9, gasPrice), nil
}

func (k Keeper) BaseFee(ctx context.Context) (*big.Int, error) {
	params, err := k.Params.Get(ctx)
	if err != nil {
		return nil, err
	}

	fee, err := k.LoadFee(ctx, params)
	if err != nil {
		return nil, err
	}

	return k.baseFee(ctx, fee)
}
