package keeper

import (
	"context"
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"

	"github.com/ethereum/go-ethereum/common"
	coretypes "github.com/ethereum/go-ethereum/core/types"

	"github.com/initia-labs/minievm/x/evm/types"
)

const SignMode_SIGN_MODE_ETHEREUM = signing.SignMode(9999)

type TxUtils struct {
	*Keeper
}

func NewTxUtils(k *Keeper) *TxUtils {
	return &TxUtils{
		Keeper: k,
	}
}

func computeGasFeeAmount(gasFeeCap *big.Int, gas uint64, decimals uint8) *big.Int {
	if gasFeeCap.Cmp(big.NewInt(0)) == 0 {
		return big.NewInt(0)
	}

	gasFeeCap = new(big.Int).Mul(gasFeeCap, new(big.Int).SetUint64(gas))
	gasFeeAmount := types.FromEthersUnit(decimals, gasFeeCap)

	// add 1 to the gas fee amount to avoid rounding errors
	return new(big.Int).Add(gasFeeAmount, big.NewInt(1))
}

// ConvertEthereumTxToCosmosTx converts an Ethereum transaction to a Cosmos SDK transaction.
func (u *TxUtils) ConvertEthereumTxToCosmosTx(ctx context.Context, ethTx *coretypes.Transaction) (sdk.Tx, error) {
	return types.ConvertEthereumTxToCosmosTx(
		sdk.UnwrapSDKContext(ctx).ChainID(), u.ac, u.cdc, ethTx,
		func() (types.Params, uint8, error) {
			params, err := u.Params.Get(ctx)
			if err != nil {
				return types.Params{}, 0, err
			}
			feeDecimals, err := u.ERC20Keeper().GetDecimals(ctx, params.FeeDenom)
			if err != nil {
				return types.Params{}, 0, err
			}
			return params, feeDecimals, nil
		},
	)
}

// ConvertCosmosTxToEthereumTx converts a Cosmos SDK transaction to an Ethereum transaction.
// It returns nil if the transaction is not an EVM transaction.
func (u *TxUtils) ConvertCosmosTxToEthereumTx(ctx context.Context, sdkTx sdk.Tx) (*coretypes.Transaction, *common.Address, error) {
	return types.ConvertCosmosTxToEthereumTx(
		sdk.UnwrapSDKContext(ctx).ChainID(), u.ac, sdkTx,
		func() (types.Params, uint8, error) {
			params, err := u.Params.Get(ctx)
			if err != nil {
				return types.Params{}, 0, err
			}
			decimals, err := u.ERC20Keeper().GetDecimals(ctx, params.FeeDenom)
			if err != nil {
				return types.Params{}, 0, err
			}
			return params, decimals, nil
		},
	)
}

// IsEthereumTx checks current context has ethereum tx
// This is used to check if the transaction is an ethereum transaction.
func (u *TxUtils) IsEthereumTx(ctx context.Context) bool {
	return sdk.UnwrapSDKContext(ctx).Value(types.CONTEXT_KEY_ETH_TX) != nil
}
