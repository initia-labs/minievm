package keeper_test

import (
	"encoding/hex"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/holiman/uint256"
	"github.com/initia-labs/minievm/x/evm/contracts/counter"
	"github.com/initia-labs/minievm/x/evm/keeper"
	"github.com/initia-labs/minievm/x/evm/types"
	"github.com/stretchr/testify/require"
)

func Test_MsgServer_Create(t *testing.T) {
	ctx, input := createDefaultTestInput(t)
	_, _, addr := keyPubAddr()

	counterBz, err := hex.DecodeString(strings.TrimPrefix(counter.CounterBin, "0x"))
	require.NoError(t, err)

	msgServer := keeper.NewMsgServerImpl(&input.EVMKeeper)
	res, err := msgServer.Create(ctx, &types.MsgCreate{
		Sender: addr.String(),
		Code:   counterBz,
	})
	require.NoError(t, err)
	require.NotEmpty(t, res.Result)
	require.True(t, common.IsHexAddress(res.ContractAddr))

	// update params to set allowed publishers
	params := types.DefaultParams()
	params.AllowedPublishers = []string{addr.String()}
	err = input.EVMKeeper.Params.Set(ctx, params)
	require.NoError(t, err)

	// allowed
	res, err = msgServer.Create(ctx, &types.MsgCreate{
		Sender: addr.String(),
		Code:   counterBz,
	})
	require.NoError(t, err)
	require.NotEmpty(t, res.Result)
	require.True(t, common.IsHexAddress(res.ContractAddr))

	// not allowed
	_, _, addr = keyPubAddr()
	_, err = msgServer.Create(ctx, &types.MsgCreate{
		Sender: addr.String(),
		Code:   counterBz,
	})
	require.Error(t, err)
}

func Test_MsgServer_Call(t *testing.T) {
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

	msgServer := keeper.NewMsgServerImpl(&input.EVMKeeper)
	res, err := msgServer.Call(ctx, &types.MsgCall{
		Sender:       addr.String(),
		ContractAddr: contractAddr.Hex(),
		Input:        inputBz,
	})
	require.NoError(t, err)
	require.Empty(t, res.Result)
	require.NotEmpty(t, res.Logs)

	queryRes, logs, err = input.EVMKeeper.EVMCall(ctx, addr, contractAddr, queryInputBz)
	require.NoError(t, err)
	require.Equal(t, uint256.NewInt(1).Bytes32(), [32]byte(queryRes))
	require.Empty(t, logs)
}

func Test_MsgServer_UpdateParams(t *testing.T) {
	ctx, input := createDefaultTestInput(t)
	msgServer := keeper.NewMsgServerImpl(&input.EVMKeeper)

	// unauthorized
	_, err := msgServer.UpdateParams(ctx, &types.MsgUpdateParams{
		Authority: "unauthorized",
		Params:    types.DefaultParams(),
	})
	require.Error(t, err)

	// invalid params
	params := types.DefaultParams()
	params.AllowedPublishers = []string{"invalid addr"}
	_, err = msgServer.UpdateParams(ctx, &types.MsgUpdateParams{
		Authority: input.EVMKeeper.GetAuthority(),
		Params:    params,
	})
	require.Error(t, err)

	// valid
	_, _, addr := keyPubAddr()
	params = types.DefaultParams()
	params.AllowedPublishers = []string{addr.String()}
	_, err = msgServer.UpdateParams(ctx, &types.MsgUpdateParams{
		Authority: input.EVMKeeper.GetAuthority(),
		Params:    params,
	})
	require.NoError(t, err)
	resParams, err := input.EVMKeeper.Params.Get(ctx)
	require.NoError(t, err)
	require.Equal(t, params, resParams)
}
