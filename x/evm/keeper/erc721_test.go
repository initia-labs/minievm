package keeper_test

import (
	"encoding/base64"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/initia-labs/minievm/x/evm/keeper"
	"github.com/initia-labs/minievm/x/evm/types"
)

func Test_CreateCollection(t *testing.T) {
	ctx, input := createDefaultTestInput(t)
	evmKeeper := input.EVMKeeper
	erc721Keeper, err := keeper.NewERC721Keeper(&evmKeeper)
	require.NoError(t, err)

	classId := "test-class-id"
	classUri := "test-class-uri"

	contractAddr := crypto.CreateAddress(types.StdAddress, 0)

	err = erc721Keeper.CreateOrUpdateClass(ctx, classId, classUri, "")
	require.NoError(t, err)

	_classId, err := evmKeeper.ERC721ClassIdsByContractAddr.Get(ctx, contractAddr.Bytes())
	require.NoError(t, err)
	require.Equal(t, classId, _classId)

	_contractAddr, err := evmKeeper.ERC721ContractAddrsByClassId.Get(ctx, classId)
	require.NoError(t, err)
	require.Equal(t, contractAddr, common.BytesToAddress(_contractAddr))

	_classUri, _classData, err := erc721Keeper.GetClassInfo(ctx, classId)
	require.NoError(t, err)
	require.Equal(t, classUri, _classUri)
	classDesc := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("{\"name\":\"%s\"}", classId)))
	require.Equal(t, classDesc, _classData)
}

func Test_CreateNFT(t *testing.T) {
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
	evmAddr := common.BytesToAddress(addr.Bytes())

	tokenIds := []string{"test-token-id", "token-idasdfasdf", "2198372123"}
	tokenUris := []string{"test-token-uri", "", "23123"}
	tokenDatas := []string{"test-token-data", "", "1239827194812"}

	err = erc721Keeper.Mints(ctx, addr, classId, tokenIds, tokenUris, tokenDatas)
	require.NoError(t, err)

	_tokenUris, _tokenDatas, err := erc721Keeper.GetTokenInfos(ctx, classId, tokenIds)
	require.NoError(t, err)

	require.Equal(t, tokenUris, _tokenUris)
	// not store tokendata
	require.Equal(t, []string{"", "", ""}, _tokenDatas)

	for _, tokenId := range tokenIds {
		owner, err := erc721Keeper.OwnerOf(ctx, tokenId, classId)
		require.NoError(t, err)
		require.Equal(t, evmAddr, owner)
	}

	balance, err := erc721Keeper.BalanceOf(ctx, addr, classId)
	require.NoError(t, err)
	require.Equal(t, int64(3), balance.Int64())

	classId = "test-class-id"
	err = erc721Keeper.CreateOrUpdateClass(ctx, classId, classUri, classData)
	require.NoError(t, err)

	tokenIds = []string{"askdjfhl1khjlk12j312"}
	err = erc721Keeper.Mints(ctx, addr, classId, tokenIds, tokenUris, tokenDatas)
	require.Error(t, err)

	tokenIds = []string{"2918379128738237"}
	err = erc721Keeper.Mints(ctx, addr, classId, tokenIds, tokenUris, tokenDatas)
	require.NoError(t, err)
}

func Test_BurnNFTs(t *testing.T) {
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

	err = erc721Keeper.Burns(ctx, addr, classId, tokenIds)
	require.NoError(t, err)

	for _, tokenId := range tokenIds {
		_, err := erc721Keeper.OwnerOf(ctx, tokenId, classId)
		require.Error(t, err)
	}

	balance, err := erc721Keeper.BalanceOf(ctx, addr, classId)
	require.NoError(t, err)
	require.Equal(t, int64(0), balance.Int64())
}

func Test_TransferNFTs(t *testing.T) {
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

	_, _, sender := keyPubAddr()
	_, _, receiver := keyPubAddr()
	receiverAddr := common.BytesToAddress(receiver.Bytes())

	tokenIds := []string{"test-token-id", "token-idasdfasdf", "2198372123"}
	tokenUris := []string{"test-token-uri", "", "23123"}
	tokenDatas := []string{"test-token-data", "", "1239827194812"}

	err = erc721Keeper.Mints(ctx, sender, classId, tokenIds, tokenUris, tokenDatas)
	require.NoError(t, err)

	err = erc721Keeper.Transfers(ctx, sender, receiver, classId, tokenIds)
	require.NoError(t, err)

	for _, tokenId := range tokenIds {
		owner, err := erc721Keeper.OwnerOf(ctx, tokenId, classId)
		require.NoError(t, err)
		require.Equal(t, receiverAddr, owner)
	}
}
