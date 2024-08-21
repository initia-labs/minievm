package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	evmtypes "github.com/initia-labs/minievm/x/evm/types"
)

func Test_GetFeeDenom(t *testing.T) {
	ctx, input := createDefaultTestInput(t)

	denom, err := input.EVMKeeper.GetFeeDenom(ctx)
	require.NoError(t, err)
	require.Equal(t, evmtypes.DefaultParams().FeeDenom, denom)

	err = input.EVMKeeper.Params.Set(ctx, evmtypes.Params{
		FeeDenom: "eth",
	})
	require.NoError(t, err)

	denom, err = input.EVMKeeper.GetFeeDenom(ctx)
	require.NoError(t, err)
	require.Equal(t, "eth", denom)
}
