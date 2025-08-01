package keeper_test

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/holiman/uint256"

	sdkmath "cosmossdk.io/math"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/initia-labs/minievm/x/evm/contracts/counter"
	"github.com/initia-labs/minievm/x/evm/keeper"
	"github.com/initia-labs/minievm/x/evm/types"

	"github.com/stretchr/testify/require"
)

func Test_MsgServer_Create(t *testing.T) {
	ctx, input := createDefaultTestInput(t)
	_, _, addr := keyPubAddr()
	evmAddr := common.BytesToAddress(addr.Bytes())

	msgServer := keeper.NewMsgServerImpl(&input.EVMKeeper)
	res, err := msgServer.Create(ctx, &types.MsgCreate{
		Sender: addr.String(),
		Code:   counter.CounterBin,
	})
	require.NoError(t, err)
	require.NotEmpty(t, res.Result)
	require.True(t, common.IsHexAddress(res.ContractAddr))

	// check generated contract address
	expectedContractAddr := crypto.CreateAddress(evmAddr, 0)
	require.Equal(t, expectedContractAddr, common.HexToAddress(res.ContractAddr))

	// update params to set allowed publishers
	params := types.DefaultParams()
	params.AllowedPublishers = []string{addr.String()}
	params.NormalizeAddresses(input.AccountKeeper.AddressCodec())
	err = input.EVMKeeper.Params.Set(ctx, params)
	require.NoError(t, err)

	// allowed
	res, err = msgServer.Create(ctx, &types.MsgCreate{
		Sender: addr.String(),
		Code:   counter.CounterBin,
	})
	require.NoError(t, err)
	require.NotEmpty(t, res.Result)
	require.True(t, common.IsHexAddress(res.ContractAddr))

	// not allowed
	_, _, addr = keyPubAddr()
	_, err = msgServer.Create(ctx, &types.MsgCreate{
		Sender: addr.String(),
		Code:   counter.CounterBin,
	})
	require.ErrorIs(t, err, sdkerrors.ErrUnauthorized)
}

func Test_MsgServer_Create2(t *testing.T) {
	ctx, input := createDefaultTestInput(t)
	_, _, addr := keyPubAddr()
	evmAddr := common.BytesToAddress(addr.Bytes())

	msgServer := keeper.NewMsgServerImpl(&input.EVMKeeper)
	res, err := msgServer.Create2(ctx, &types.MsgCreate2{
		Sender: addr.String(),
		Code:   counter.CounterBin,
		Salt:   sdkmath.NewInt(1),
	})
	require.NoError(t, err)
	require.NotEmpty(t, res.Result)

	// check generated contract address
	salt := uint256.NewInt(1)
	expectedContractAddr := crypto.CreateAddress2(evmAddr, salt.Bytes32(), crypto.Keccak256Hash(hexutil.MustDecode(counter.CounterBin)).Bytes())
	require.Equal(t, expectedContractAddr, common.HexToAddress(res.ContractAddr))

	// negative salt
	_, err = msgServer.Create2(ctx, &types.MsgCreate2{
		Sender: addr.String(),
		Code:   counter.CounterBin,
		Salt:   sdkmath.NewInt(-1),
	})
	require.ErrorIs(t, err, types.ErrInvalidSalt)
}

func Test_MsgServer_Call(t *testing.T) {
	ctx, input := createDefaultTestInput(t)
	_, _, addr := keyPubAddr()
	caller := common.BytesToAddress(addr.Bytes())

	counterBz, err := hexutil.Decode(counter.CounterBin)
	require.NoError(t, err)

	retBz, contractAddr, _, err := input.EVMKeeper.EVMCreate(ctx, caller, counterBz, nil, nil)
	require.NoError(t, err)
	require.NotEmpty(t, retBz)
	require.Len(t, contractAddr, 20)

	parsed, err := counter.CounterMetaData.GetAbi()
	require.NoError(t, err)

	queryInputBz, err := parsed.Pack("count")
	require.NoError(t, err)

	queryRes, logs, err := input.EVMKeeper.EVMCall(ctx, caller, contractAddr, queryInputBz, nil, nil)
	require.NoError(t, err)
	require.Equal(t, uint256.NewInt(0).Bytes32(), [32]byte(queryRes))
	require.Empty(t, logs)

	inputBz, err := parsed.Pack("increase")
	require.NoError(t, err)

	msgServer := keeper.NewMsgServerImpl(&input.EVMKeeper)
	res, err := msgServer.Call(ctx, &types.MsgCall{
		Sender:       addr.String(),
		ContractAddr: contractAddr.Hex(),
		Input:        hexutil.Encode(inputBz),
	})
	require.NoError(t, err)
	require.Equal(t, "0x", res.Result)
	require.NotEmpty(t, res.Logs)

	queryRes, logs, err = input.EVMKeeper.EVMCall(ctx, caller, contractAddr, queryInputBz, nil, nil)
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

	// cannot change fee denom
	params = types.DefaultParams()
	params.FeeDenom = "otherdenom"
	_, err = msgServer.UpdateParams(ctx, &types.MsgUpdateParams{
		Authority: input.EVMKeeper.GetAuthority(),
		Params:    params,
	})
	require.ErrorContains(t, err, "not found")

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
	// normalize addresses to compare results
	err = params.NormalizeAddresses(input.AccountKeeper.AddressCodec())
	require.NoError(t, err)
	require.Equal(t, params, resParams)

	// deploy custom erc20 contract
	evmAddr := common.BytesToAddress(addr.Bytes())
	fooContractAddr := deployCustomERC20(t, ctx, input, evmAddr, "foo", true)
	fooDenom, err := types.ContractAddrToDenom(ctx, &input.EVMKeeper, fooContractAddr)
	require.NoError(t, err)
	require.Equal(t, "evm/"+fooContractAddr.Hex()[2:], fooDenom)

	// cannot change fee denom which does not support sudoMint and sudoBurn
	params = types.DefaultParams()
	params.FeeDenom = fooDenom
	_, err = msgServer.UpdateParams(ctx, &types.MsgUpdateParams{
		Authority: input.EVMKeeper.GetAuthority(),
		Params:    params,
	})
	require.ErrorContains(t, err, "sudoMint and sudoBurn")
}

func Test_MsgServer_NonceIncrement_Call(t *testing.T) {
	ctx, input := createDefaultTestInput(t)
	_, _, addr := keyPubAddr()
	caller := common.BytesToAddress(addr.Bytes())

	counterBz, err := hexutil.Decode(counter.CounterBin)
	require.NoError(t, err)

	retBz, contractAddr, _, err := input.EVMKeeper.EVMCreate(ctx, caller, counterBz, nil, nil)
	require.NoError(t, err)
	require.NotEmpty(t, retBz)
	require.Len(t, contractAddr, 20)

	parsed, err := counter.CounterMetaData.GetAbi()
	require.NoError(t, err)

	// increment sequence
	incremented := true
	ctx = ctx.WithValue(types.CONTEXT_KEY_SEQUENCE_INCREMENTED, &incremented)
	acc := input.AccountKeeper.GetAccount(ctx, addr)
	seq := acc.GetSequence() + 1
	acc.SetSequence(seq)
	input.AccountKeeper.SetAccount(ctx, acc)

	inputBz, err := parsed.Pack("increase")
	require.NoError(t, err)

	// should not increment sequence
	msgServer := keeper.NewMsgServerImpl(&input.EVMKeeper)
	res, err := msgServer.Call(ctx, &types.MsgCall{
		Sender:       addr.String(),
		ContractAddr: contractAddr.Hex(),
		Input:        hexutil.Encode(inputBz),
	})
	require.NoError(t, err)
	require.Equal(t, "0x", res.Result)
	require.NotEmpty(t, res.Logs)

	acc = input.AccountKeeper.GetAccount(ctx, addr)
	require.Equal(t, seq, acc.GetSequence())

	// call again should increment sequence
	res, err = msgServer.Call(ctx, &types.MsgCall{
		Sender:       addr.String(),
		ContractAddr: contractAddr.Hex(),
		Input:        hexutil.Encode(inputBz),
	})
	require.NoError(t, err)
	require.Equal(t, "0x", res.Result)
	require.NotEmpty(t, res.Logs)

	acc = input.AccountKeeper.GetAccount(ctx, addr)
	require.Equal(t, seq+1, acc.GetSequence())

	// create should increment sequence
	_, err = msgServer.Create(ctx, &types.MsgCreate{
		Sender: addr.String(),
		Code:   counter.CounterBin,
	})
	require.NoError(t, err)

	acc = input.AccountKeeper.GetAccount(ctx, addr)
	require.Equal(t, seq+2, acc.GetSequence())
}

func Test_MsgServer_NonceIncrement_Create(t *testing.T) {
	ctx, input := createDefaultTestInput(t)
	_, _, addr := keyPubAddr()
	caller := common.BytesToAddress(addr.Bytes())

	counterBz, err := hexutil.Decode(counter.CounterBin)
	require.NoError(t, err)

	retBz, contractAddr, _, err := input.EVMKeeper.EVMCreate(ctx, caller, counterBz, nil, nil)
	require.NoError(t, err)
	require.NotEmpty(t, retBz)
	require.Len(t, contractAddr, 20)

	parsed, err := counter.CounterMetaData.GetAbi()
	require.NoError(t, err)

	// increment sequence
	incremented := true
	ctx = ctx.WithValue(types.CONTEXT_KEY_SEQUENCE_INCREMENTED, &incremented)
	acc := input.AccountKeeper.GetAccount(ctx, addr)
	seq := acc.GetSequence() + 1
	acc.SetSequence(seq)
	input.AccountKeeper.SetAccount(ctx, acc)

	inputBz, err := parsed.Pack("increase")
	require.NoError(t, err)

	// should not increment sequence
	msgServer := keeper.NewMsgServerImpl(&input.EVMKeeper)
	_, err = msgServer.Create(ctx, &types.MsgCreate{
		Sender: addr.String(),
		Code:   counter.CounterBin,
	})
	require.NoError(t, err)

	acc = input.AccountKeeper.GetAccount(ctx, addr)
	require.Equal(t, seq, acc.GetSequence())

	// call again should increment sequence
	_, err = msgServer.Call(ctx, &types.MsgCall{
		Sender:       addr.String(),
		ContractAddr: contractAddr.Hex(),
		Input:        hexutil.Encode(inputBz),
	})
	require.NoError(t, err)

	acc = input.AccountKeeper.GetAccount(ctx, addr)
	require.Equal(t, seq+1, acc.GetSequence())

	// create should increment sequence
	_, err = msgServer.Create(ctx, &types.MsgCreate{
		Sender: addr.String(),
		Code:   counter.CounterBin,
	})
	require.NoError(t, err)

	acc = input.AccountKeeper.GetAccount(ctx, addr)
	require.Equal(t, seq+2, acc.GetSequence())
}
