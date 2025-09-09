package types

import (
	"fmt"
	"testing"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/stretchr/testify/require"
)

type authMockCodec struct{}

func (m authMockCodec) StringToBytes(s string) ([]byte, error) {
	if s == "bad" {
		return nil, fmt.Errorf("mock error")
	}
	return []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f, 0x10, 0x11, 0x12, 0x13, 0x14}, nil
}

func (m authMockCodec) BytesToString(b []byte) (string, error) {
	if len(b) == 0 {
		return "", fmt.Errorf("mock error")
	}
	return "mock_address", nil
}

func TestNewContractAccountWithAddress(t *testing.T) {
	addr := sdk.AccAddress("test_address_12345678901234567890")

	account := NewContractAccountWithAddress(addr)
	require.NotNil(t, account)
	require.Equal(t, addr, account.GetAddress())
	require.Empty(t, account.CodeHash)
	require.Equal(t, uint64(0), account.GetSequence())
	require.Nil(t, account.GetPubKey())
}

func TestContractAccount_SetPubKey(t *testing.T) {
	addr := sdk.AccAddress("test_address_12345678901234567890")
	account := NewContractAccountWithAddress(addr)

	pubKey := secp256k1.GenPrivKey().PubKey()
	err := account.SetPubKey(pubKey)
	require.Error(t, err)
	require.Contains(t, err.Error(), "not supported for contract accounts")
}

func TestNewShorthandAccountWithAddress(t *testing.T) {
	ac := authMockCodec{}
	addr := sdk.AccAddress("test_address_12345678901234567890")

	account, err := NewShorthandAccountWithAddress(ac, addr)
	require.NoError(t, err)
	require.NotNil(t, account)
	require.Equal(t, "mock_address", account.OriginalAddress)
	require.Equal(t, uint64(0), account.GetSequence())
	require.Nil(t, account.GetPubKey())
}

func TestNewShorthandAccountWithAddress_Error(t *testing.T) {
	ac := authMockCodec{}
	addr := sdk.AccAddress("") // empty address will cause error in mockCodec

	account, err := NewShorthandAccountWithAddress(ac, addr)
	require.Error(t, err)
	require.Nil(t, account)
	require.Contains(t, err.Error(), "mock error")
}

func TestShorthandAccount_SetPubKey(t *testing.T) {
	ac := authMockCodec{}
	addr := sdk.AccAddress("test_address_12345678901234567890")
	account, err := NewShorthandAccountWithAddress(ac, addr)
	require.NoError(t, err)

	pubKey := secp256k1.GenPrivKey().PubKey()
	err = account.SetPubKey(pubKey)
	require.Error(t, err)
	require.Contains(t, err.Error(), "not supported for shorthand accounts")
}

func TestShorthandAccount_GetOriginalAddress(t *testing.T) {
	ac := authMockCodec{}
	addr := sdk.AccAddress("test_address_12345678901234567890")
	account, err := NewShorthandAccountWithAddress(ac, addr)
	require.NoError(t, err)

	originalAddr, err := account.GetOriginalAddress(ac)
	require.NoError(t, err)
	require.Equal(t, sdk.AccAddress([]byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f, 0x10, 0x11, 0x12, 0x13, 0x14}), originalAddr)
}

func TestShorthandAccount_GetOriginalAddress_Error(t *testing.T) {
	ac := authMockCodec{}
	addr := sdk.AccAddress("test_address_12345678901234567890")
	account, err := NewShorthandAccountWithAddress(ac, addr)
	require.NoError(t, err)

	// Create a bad codec that will fail
	badCodec := badAuthMockCodec{}
	originalAddr, err := account.GetOriginalAddress(badCodec)
	require.Error(t, err)
	require.Nil(t, originalAddr)
	require.Contains(t, err.Error(), "mock error")
}

type badAuthMockCodec struct{}

func (m badAuthMockCodec) StringToBytes(s string) ([]byte, error) {
	return nil, fmt.Errorf("mock error")
}

func (m badAuthMockCodec) BytesToString(b []byte) (string, error) {
	return "", fmt.Errorf("mock error")
}

func TestIsEmptyAccount(t *testing.T) {
	t.Run("empty_base_account", func(t *testing.T) {
		addr := sdk.AccAddress("test_address_12345678901234567890")
		account := authtypes.NewBaseAccountWithAddress(addr)
		require.True(t, IsEmptyAccount(account))
	})

	t.Run("non_empty_base_account", func(t *testing.T) {
		addr := sdk.AccAddress("test_address_12345678901234567890")
		account := authtypes.NewBaseAccountWithAddress(addr)
		account.SetSequence(1)
		require.False(t, IsEmptyAccount(account))
	})

	t.Run("account_with_pubkey", func(t *testing.T) {
		addr := sdk.AccAddress("test_address_12345678901234567890")
		account := authtypes.NewBaseAccountWithAddress(addr)
		pubKey := secp256k1.GenPrivKey().PubKey()
		account.SetPubKey(pubKey)
		require.False(t, IsEmptyAccount(account))
	})

	t.Run("contract_account", func(t *testing.T) {
		addr := sdk.AccAddress("test_address_12345678901234567890")
		account := NewContractAccountWithAddress(addr)
		require.False(t, IsEmptyAccount(account))
	})

	t.Run("shorthand_account", func(t *testing.T) {
		ac := authMockCodec{}
		addr := sdk.AccAddress("test_address_12345678901234567890")
		account, err := NewShorthandAccountWithAddress(ac, addr)
		require.NoError(t, err)
		require.False(t, IsEmptyAccount(account))
	})

	t.Run("module_account", func(t *testing.T) {
		// Create a mock module account
		moduleAccount := &mockModuleAccount{
			BaseAccount: authtypes.NewBaseAccountWithAddress(sdk.AccAddress("module_address_12345678901234567890")),
		}
		require.False(t, IsEmptyAccount(moduleAccount))
	})
}

type mockModuleAccount struct {
	*authtypes.BaseAccount
}

func (m *mockModuleAccount) GetName() string {
	return "mock_module"
}

func (m *mockModuleAccount) GetPermissions() []string {
	return []string{"mock_permission"}
}

func (m *mockModuleAccount) HasPermission(permission string) bool {
	return permission == "mock_permission"
}

func TestAccountAddressConversion(t *testing.T) {
	ac := authMockCodec{}
	originalAddr := sdk.AccAddress("test_address_12345678901234567890")

	// Test the full cycle: create shorthand account and get original address back
	shorthandAccount, err := NewShorthandAccountWithAddress(ac, originalAddr)
	require.NoError(t, err)

	retrievedAddr, err := shorthandAccount.GetOriginalAddress(ac)
	require.NoError(t, err)
	require.Equal(t, sdk.AccAddress([]byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f, 0x10, 0x11, 0x12, 0x13, 0x14}), retrievedAddr)
}
