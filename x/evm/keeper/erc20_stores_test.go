package keeper_test

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/holiman/uint256"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/initia-labs/minievm/x/evm/contracts/initia_erc20"
	"github.com/initia-labs/minievm/x/evm/types"
)

func deployCustomERC20(t *testing.T, ctx sdk.Context, input TestKeepers, caller common.Address, denom string, success bool) common.Address {
	abi, err := initia_erc20.InitiaErc20MetaData.GetAbi()
	require.NoError(t, err)

	bin, err := hexutil.Decode(initia_erc20.InitiaErc20MetaData.Bin)
	require.NoError(t, err)

	inputBz, err := abi.Pack("", denom, denom, uint8(6))
	require.NoError(t, err)

	_, contractAddr, _, err := input.EVMKeeper.EVMCreate(ctx, caller, append(bin, inputBz...), uint256.NewInt(0), nil)
	if success {
		require.NoError(t, err)
	} else {
		require.Error(t, err)
	}

	return contractAddr
}

func Test_CanDeployCustomERC20(t *testing.T) {
	ctx, input := createDefaultTestInput(t)

	_, _, addr := keyPubAddr()
	evmAddr := common.BytesToAddress(addr.Bytes())

	_, _, addr2 := keyPubAddr()
	evmAddr2 := common.BytesToAddress(addr2.Bytes())

	params, err := input.EVMKeeper.Params.Get(ctx)
	require.NoError(t, err)

	// allow custom erc20
	params.AllowCustomERC20 = true
	err = input.EVMKeeper.Params.Set(ctx, params)
	require.NoError(t, err)

	// deploy custom erc20 contract
	fooContractAddr := deployCustomERC20(t, ctx, input, evmAddr, "foo", true)
	fooDenom, err := types.ContractAddrToDenom(ctx, &input.EVMKeeper, fooContractAddr)
	require.NoError(t, err)
	require.Equal(t, "evm/"+fooContractAddr.Hex()[2:], fooDenom)

	// limit allowed custom erc20s
	estimatedContractAddr := crypto.CreateAddress(evmAddr2, 0)
	params.AllowedCustomERC20s = []string{"foo", estimatedContractAddr.Hex()}
	err = input.EVMKeeper.Params.Set(ctx, params)
	require.NoError(t, err)

	// should failed to deploy new custom contract from addr
	deployCustomERC20(t, ctx, input, evmAddr, "foo2", false)

	// should success to deploy new custom contract from addr2
	deployCustomERC20(t, ctx, input, evmAddr2, "foo2", true)
}
