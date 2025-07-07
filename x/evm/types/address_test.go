package types

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

type mockCodec struct{}

func (m mockCodec) StringToBytes(s string) ([]byte, error) {
	if s == "bad" {
		return nil, errMock
	}
	return []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f, 0x10, 0x11, 0x12, 0x13, 0x14}, nil
}

func (m mockCodec) BytesToString(b []byte) (string, error) {
	return "mock", nil
}

var errMock = &mockError{"mock error"}

type mockError struct{ msg string }

func (e *mockError) Error() string { return e.msg }

func TestAddressConstants(t *testing.T) {
	require.Equal(t, common.HexToAddress("0x0"), NullAddress)
	require.Equal(t, common.HexToAddress("0x1"), StdAddress)
	require.Equal(t, uint64(1), ERC20FactorySalt)
	require.Equal(t, uint64(2), ERC20WrapperSalt)
	require.Equal(t, uint64(3), ConnectOracleSalt)
	require.Equal(t, common.HexToAddress("0xf1"), CosmosPrecompileAddress)
	require.Equal(t, common.HexToAddress("0xf2"), ERC20RegistryPrecompileAddress)
	require.Equal(t, common.HexToAddress("0xf3"), JSONUtilsPrecompileAddress)
}

func TestContractAddressFromString_Hex(t *testing.T) {
	ac := mockCodec{}
	addr, err := ContractAddressFromString(ac, "0x1234567890123456789012345678901234567890")
	require.NoError(t, err)
	require.Equal(t, common.HexToAddress("0x1234567890123456789012345678901234567890"), addr)
}

func TestContractAddressFromString_NonHex(t *testing.T) {
	ac := mockCodec{}
	addr, err := ContractAddressFromString(ac, "nothex")
	require.NoError(t, err)
	require.Equal(t, common.BytesToAddress([]byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f, 0x10, 0x11, 0x12, 0x13, 0x14}), addr)
}

func TestContractAddressFromString_Error(t *testing.T) {
	ac := mockCodec{}
	addr, err := ContractAddressFromString(ac, "bad")
	require.Error(t, err)
	require.Equal(t, common.Address{}, addr)
}

func TestContractAddressFromString_LengthEnforcement(t *testing.T) {
	ac := mockCodec{}

	// Valid 20 bytes hex address
	addr, err := ContractAddressFromString(ac, "0x1234567890123456789012345678901234567890")
	require.NoError(t, err)
	require.Equal(t, common.HexToAddress("0x1234567890123456789012345678901234567890"), addr)

	// Valid 20 bytes non-hex address (mockCodec returns 20 bytes)
	addr, err = ContractAddressFromString(ac, "mock")
	require.NoError(t, err)
	require.Equal(t, common.BytesToAddress([]byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f, 0x10, 0x11, 0x12, 0x13, 0x14}), addr)

	// Invalid non-hex address (mockCodec returns less than 20 bytes)
	shortCodec := shortCodec{}
	_, err = ContractAddressFromString(shortCodec, "short")
	require.Error(t, err)
	require.ErrorIs(t, err, ErrInvalidAddressLength)
}

type shortCodec struct{}

func (m shortCodec) StringToBytes(s string) ([]byte, error) {
	return []byte{0x01, 0x02}, nil // only 2 bytes
}

func (m shortCodec) BytesToString(b []byte) (string, error) {
	return "short", nil
}
