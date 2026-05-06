package checktx

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	abci "github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type feeDenomGetterFunc func(context.Context) (string, error)

func (f feeDenomGetterFunc) GetFeeDenom(ctx context.Context) (string, error) {
	return f(ctx)
}

type panicContextGetter struct{}

func (panicContextGetter) GetContextForCheckTx([]byte) sdk.Context {
	panic("boom")
}

func TestUpdateFeeDenomCacheLoadsAtHeightZero(t *testing.T) {
	loads := 0
	w := &CheckTxWrapper{
		fdg: feeDenomGetterFunc(func(context.Context) (string, error) {
			loads++
			return sdk.DefaultBondDenom, nil
		}),
	}

	denom, err := w.updateFeeDenomCache(sdk.Context{}.WithBlockHeight(0))
	require.NoError(t, err)
	require.Equal(t, sdk.DefaultBondDenom, denom)
	require.Equal(t, 1, loads)

	denom, err = w.updateFeeDenomCache(sdk.Context{}.WithBlockHeight(0))
	require.NoError(t, err)
	require.Equal(t, sdk.DefaultBondDenom, denom)
	require.Equal(t, 1, loads)
}

func TestCheckTxRecoversPanic(t *testing.T) {
	w := &CheckTxWrapper{
		cg: panicContextGetter{},
	}

	res, err := w.CheckTx()(&abci.RequestCheckTx{})
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, sdkerrors.ErrPanic.ABCICode(), res.Code)
}
