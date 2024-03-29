package keeper_test

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/initia-labs/minievm/x/evm/contracts/erc20"
	"github.com/initia-labs/minievm/x/evm/keeper"
	"github.com/initia-labs/minievm/x/evm/types"
)

func Test_Query_Call(t *testing.T) {
	ctx, input := createDefaultTestInput(t)
	_, _, addr := keyPubAddr()
	evmAddr := common.BytesToAddress(addr.Bytes())

	// deploy erc20 contract
	fooContractAddr := deployERC20(t, ctx, input, evmAddr, "foo")
	fooDenom, err := types.ContractAddrToDenom(ctx, &input.EVMKeeper, fooContractAddr)
	require.NoError(t, err)
	require.Equal(t, "evm/"+fooContractAddr.Hex()[2:], fooDenom)

	// mint erc20
	mintERC20(t, ctx, input, evmAddr, evmAddr, sdk.NewCoin(fooDenom, math.NewInt(100)))

	abi, err := erc20.Erc20MetaData.GetAbi()
	require.NoError(t, err)
	inputBz, err := abi.Pack("balanceOf", evmAddr)
	require.NoError(t, err)

	qs := keeper.NewQueryServer(&input.EVMKeeper)
	res, err := qs.Call(ctx, &types.QueryCallRequest{
		Sender:       addr.String(),
		ContractAddr: fooContractAddr.String(),
		Input:        hexutil.Encode(inputBz),
		WithTrace:    true,
	})
	require.NoError(t, err)

	outputBz := hexutil.MustDecode(res.Response)
	ret, err := abi.Unpack("balanceOf", outputBz)
	require.NoError(t, err)

	require.Equal(t, uint64(100), ret[0].(*big.Int).Uint64())
}
