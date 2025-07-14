package types

import (
	"context"
	"errors"
	"testing"

	"cosmossdk.io/collections"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

type mockERC20DenomKeeper struct {
	addr  common.Address
	denom string
	err   error
}

func (m mockERC20DenomKeeper) GetContractAddrByDenom(_ context.Context, denom string) (common.Address, error) {
	if m.err != nil {
		return common.Address{}, m.err
	}
	return m.addr, nil
}

func (m mockERC20DenomKeeper) GetDenomByContractAddr(_ context.Context, addr common.Address) (string, error) {
	if m.err != nil {
		return "", m.err
	}
	return m.denom, nil
}

func TestDenomToContractAddr(t *testing.T) {
	ctx := context.Background()
	keeper := mockERC20DenomKeeper{addr: common.HexToAddress("0x1234567890123456789012345678901234567890")}

	// ERC20 denom, valid hex
	addr, err := DenomToContractAddr(ctx, keeper, "evm/1234567890123456789012345678901234567890")
	require.NoError(t, err)
	require.Equal(t, keeper.addr, addr)

	// ERC20 denom, invalid hex
	_, err = DenomToContractAddr(ctx, keeper, "evm/invalidhex")
	require.Error(t, err)

	// Non-ERC20 denom, fallback to keeper
	keeper2 := mockERC20DenomKeeper{addr: common.HexToAddress("0xabcdefabcdefabcdefabcdefabcdefabcdefabcd")}
	addr, err = DenomToContractAddr(ctx, keeper2, "foo/bar")
	require.NoError(t, err)
	require.Equal(t, keeper2.addr, addr)
}

func TestContractAddrToDenom(t *testing.T) {
	ctx := context.Background()
	keeper := mockERC20DenomKeeper{denom: "evm/1234567890123456789012345678901234567890"}
	addr := common.HexToAddress("0x1234567890123456789012345678901234567890")

	// Keeper returns denom
	denom, err := ContractAddrToDenom(ctx, keeper, addr)
	require.NoError(t, err)
	require.Equal(t, keeper.denom, denom)

	// Keeper returns collections.ErrNotFound
	keeperNF := mockERC20DenomKeeper{err: collections.ErrNotFound}
	denom, err = ContractAddrToDenom(ctx, keeperNF, addr)
	require.NoError(t, err)
	require.True(t, IsERC20Denom(denom))

	// Keeper returns other error
	keeperErr := mockERC20DenomKeeper{err: errors.New("other error")}
	_, err = ContractAddrToDenom(ctx, keeperErr, addr)
	require.Error(t, err)
}

func TestIsERC20Denom(t *testing.T) {
	require.True(t, IsERC20Denom("evm/123"))
	require.False(t, IsERC20Denom("foo/123"))
}
