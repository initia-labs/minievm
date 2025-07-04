package types

import (
	"context"
	"crypto/sha256"
	"errors"
	"math/big"
	"testing"

	"cosmossdk.io/collections"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

type mockERC721ClassIdKeeper struct {
	addr  common.Address
	class string
	err   error
}

func (m mockERC721ClassIdKeeper) GetContractAddrByClassId(_ context.Context, classId string) (common.Address, error) {
	if m.err != nil {
		return common.Address{}, m.err
	}
	return m.addr, nil
}

func (m mockERC721ClassIdKeeper) GetClassIdByContractAddr(_ context.Context, addr common.Address) (string, error) {
	if m.err != nil {
		return "", m.err
	}
	return m.class, nil
}

func TestTokenIdToBigInt(t *testing.T) {
	// IBC classId: should hash tokenId
	classId := IBCPrefix + "someibcclass"
	tokenId := "token123"
	bi, ok := TokenIdToBigInt(classId, tokenId)
	require.True(t, ok)
	hash := sha256.Sum256([]byte(tokenId))
	require.Equal(t, new(big.Int).SetBytes(hash[:]), bi)

	// Non-IBC classId: should parse as decimal
	classId = "otherclass"
	tokenId = "123456"
	bi, ok = TokenIdToBigInt(classId, tokenId)
	require.True(t, ok)
	require.Equal(t, big.NewInt(123456), bi)

	// Non-IBC classId: invalid decimal
	classId = "otherclass"
	tokenId = "notanumber"
	bi, ok = TokenIdToBigInt(classId, tokenId)
	require.False(t, ok)
	require.Nil(t, bi)
}

func TestContractAddressFromClassId(t *testing.T) {
	ctx := context.Background()
	keeper := mockERC721ClassIdKeeper{addr: common.HexToAddress("0x1234567890123456789012345678901234567890")}

	// EVM classId, valid address
	classId := EVMPrefix + "1234567890123456789012345678901234567890"
	addr, err := ContractAddressFromClassId(ctx, keeper, classId)
	require.NoError(t, err)
	require.Equal(t, common.HexToAddress("0x1234567890123456789012345678901234567890"), addr)

	// EVM classId, invalid address
	classId = EVMPrefix + "notanaddress"
	addr, err = ContractAddressFromClassId(ctx, keeper, classId)
	require.Error(t, err)
	require.Equal(t, NullAddress, addr)

	// Non-EVM classId, keeper returns address
	classId = "otherclass"
	addr, err = ContractAddressFromClassId(ctx, keeper, classId)
	require.NoError(t, err)
	require.Equal(t, keeper.addr, addr)

	// Non-EVM classId, keeper returns error
	keeperErr := mockERC721ClassIdKeeper{err: errors.New("fail")}
	_, err = ContractAddressFromClassId(ctx, keeperErr, classId)
	require.Error(t, err)
}

func TestClassIdFromCollectionAddress(t *testing.T) {
	ctx := context.Background()
	keeper := mockERC721ClassIdKeeper{class: "myclass"}
	addr := common.HexToAddress("0x1234567890123456789012345678901234567890")

	// Keeper returns classId
	classId, err := ClassIdFromCollectionAddress(ctx, keeper, addr)
	require.NoError(t, err)
	require.Equal(t, "myclass", classId)

	// Keeper returns collections.ErrNotFound
	keeperNF := mockERC721ClassIdKeeper{err: collections.ErrNotFound}
	classId, err = ClassIdFromCollectionAddress(ctx, keeperNF, addr)
	require.NoError(t, err)
	require.Equal(t, EVMPrefix+"1234567890123456789012345678901234567890", classId)

	// Keeper returns other error
	keeperErr := mockERC721ClassIdKeeper{err: errors.New("fail")}
	_, err = ClassIdFromCollectionAddress(ctx, keeperErr, addr)
	require.Error(t, err)
}

func TestIsEVMClassId(t *testing.T) {
	require.True(t, IsEVMClassId(EVMPrefix+"something"))
	require.False(t, IsEVMClassId("ibc/something"))
	require.False(t, IsEVMClassId("other"))
}
