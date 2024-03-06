package keeper_test

import (
	"encoding/hex"
	"strings"
	"testing"

	"github.com/holiman/uint256"
	"github.com/initia-labs/minievm/x/evm/contracts/counter"
	"github.com/stretchr/testify/require"
)

func Test_Create(t *testing.T) {
	ctx, input := createDefaultTestInput(t)
	_, _, addr := keyPubAddr()

	counterBz, err := hex.DecodeString(strings.TrimPrefix(counter.CounterBin, "0x"))
	require.NoError(t, err)

	retBz, contractAddr, err := input.EVMKeeper.EVMCreate(ctx, addr, counterBz)
	require.NoError(t, err)
	require.NotEmpty(t, retBz)
	require.Len(t, contractAddr, 20)
}

func Test_Call(t *testing.T) {
	ctx, input := createDefaultTestInput(t)
	_, _, addr := keyPubAddr()

	counterBz, err := hex.DecodeString(strings.TrimPrefix(counter.CounterBin, "0x"))
	require.NoError(t, err)

	retBz, contractAddr, err := input.EVMKeeper.EVMCreate(ctx, addr, counterBz)
	require.NoError(t, err)
	require.NotEmpty(t, retBz)
	require.Len(t, contractAddr, 20)

	parsed, err := counter.CounterMetaData.GetAbi()
	require.NoError(t, err)

	queryInputBz, err := parsed.Pack("count")
	require.NoError(t, err)

	queryRes, logs, err := input.EVMKeeper.EVMCall(ctx, addr, contractAddr, queryInputBz)
	require.NoError(t, err)
	require.Equal(t, uint256.NewInt(0).Bytes32(), [32]byte(queryRes))
	require.Empty(t, logs)

	inputBz, err := parsed.Pack("increase")
	require.NoError(t, err)

	res, logs, err := input.EVMKeeper.EVMCall(ctx, addr, contractAddr, inputBz)
	require.NoError(t, err)
	require.Empty(t, res)
	require.NotEmpty(t, logs)

	queryRes, logs, err = input.EVMKeeper.EVMCall(ctx, addr, contractAddr, queryInputBz)
	require.NoError(t, err)
	require.Equal(t, uint256.NewInt(1).Bytes32(), [32]byte(queryRes))
	require.Empty(t, logs)
}
