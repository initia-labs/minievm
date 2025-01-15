package keeper_test

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"

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
	mintERC20(t, ctx, input, evmAddr, evmAddr, sdk.NewCoin(fooDenom, math.NewInt(100)), false)

	abi, err := erc20.Erc20MetaData.GetAbi()
	require.NoError(t, err)
	inputBz, err := abi.Pack("balanceOf", evmAddr)
	require.NoError(t, err)

	qs := keeper.NewQueryServer(&input.EVMKeeper)
	res, err := qs.Call(ctx, &types.QueryCallRequest{
		Sender:       addr.String(),
		ContractAddr: fooContractAddr.String(),
		Input:        hexutil.Encode(inputBz),
		TraceOptions: &types.TraceOptions{},
	})
	require.NoError(t, err)

	outputBz := hexutil.MustDecode(res.Response)
	ret, err := abi.Unpack("balanceOf", outputBz)
	require.NoError(t, err)

	require.Equal(t, uint64(100), ret[0].(*big.Int).Uint64())
}

func Test_Query_ERC20Factory(t *testing.T) {
	ctx, input := createDefaultTestInput(t)

	factoryAddr, err := input.EVMKeeper.GetERC20FactoryAddr(ctx)
	require.NoError(t, err)

	qs := keeper.NewQueryServer(&input.EVMKeeper)
	res, err := qs.ERC20Factory(ctx, &types.QueryERC20FactoryRequest{})
	require.NoError(t, err)

	require.Equal(t, factoryAddr.Hex(), res.Address)
}

func Test_Query_ERC20Wrapper(t *testing.T) {
	ctx, input := createDefaultTestInput(t)

	wrapperAddr, err := input.EVMKeeper.GetERC20WrapperAddr(ctx)
	require.NoError(t, err)

	qs := keeper.NewQueryServer(&input.EVMKeeper)
	res, err := qs.ERC20Wrapper(ctx, &types.QueryERC20WrapperRequest{})
	require.NoError(t, err)

	require.Equal(t, wrapperAddr.Hex(), res.Address)
}

func Test_Query_ERC721ClassIdByContractAddr(t *testing.T) {
	ctx, input := createDefaultTestInput(t)

	evmKeeper := input.EVMKeeper
	erc721Keeper, err := keeper.NewERC721Keeper(&evmKeeper)
	require.NoError(t, err)

	classId := "test-class-id"
	classUri := "test-class-uri"

	contractAddr := crypto.CreateAddress(types.StdAddress, 2)

	err = erc721Keeper.CreateOrUpdateClass(ctx, classId, classUri, "")
	require.NoError(t, err)

	qs := keeper.NewQueryServer(&input.EVMKeeper)
	res, err := qs.ERC721ClassIdByContractAddr(ctx, &types.QueryERC721ClassIdByContractAddrRequest{
		ContractAddr: contractAddr.Hex(),
	})
	require.NoError(t, err)

	require.Equal(t, classId, res.ClassId)
}

func Test_Query_ERC721OriginTokenInfos(t *testing.T) {
	ctx, input := createDefaultTestInput(t)

	evmKeeper := input.EVMKeeper
	ierc721Keeper, err := keeper.NewERC721Keeper(&evmKeeper)
	require.NoError(t, err)
	erc721Keeper := ierc721Keeper.(*keeper.ERC721Keeper)

	classId := "ibc/test-class-id"
	classUri := "test-class-uri"
	classData := "test-class-data"

	err = erc721Keeper.CreateOrUpdateClass(ctx, classId, classUri, classData)
	require.NoError(t, err)

	_, _, addr := keyPubAddr()

	tokenIds := []string{"test-token-id", "token-idasdfasdf", "2198372123"}
	tokenUris := []string{"test-token-uri", "", "23123"}
	tokenDatas := []string{"test-token-data", "", "1239827194812"}

	err = erc721Keeper.Mints(ctx, addr, classId, tokenIds, tokenUris, tokenDatas)
	require.NoError(t, err)

	evmTokenIds := make([]string, 0)
	for i := range tokenIds {
		tokenId, ok := types.TokenIdToBigInt(classId, tokenIds[i])
		require.True(t, ok)
		evmTokenIds = append(evmTokenIds, tokenId.String())
	}

	qs := keeper.NewQueryServer(&input.EVMKeeper)
	res, err := qs.ERC721OriginTokenInfos(ctx, &types.QueryERC721OriginTokenInfosRequest{
		ClassId:  classId,
		TokenIds: evmTokenIds,
	})
	require.NoError(t, err)

	actualTokenIds := make([]string, 0)
	actualTokenUris := make([]string, 0)
	for i := range res.TokenInfos {
		actualTokenIds = append(actualTokenIds, res.TokenInfos[i].TokenOriginId)
		actualTokenUris = append(actualTokenUris, res.TokenInfos[i].TokenUri)
	}
	require.Equal(t, tokenIds, actualTokenIds)
	require.Equal(t, tokenUris, actualTokenUris)
}
