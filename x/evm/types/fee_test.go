package types

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestNewFee(t *testing.T) {
	denom := "testdenom"
	contract := common.HexToAddress("0x1234567890123456789012345678901234567890")
	decimals := uint8(18)

	fee := NewFee(denom, contract, decimals)
	require.Equal(t, denom, fee.Denom())
	require.Equal(t, contract, fee.Contract())
	require.Equal(t, decimals, fee.Decimals())
	require.True(t, fee.HasContract())
}

func TestFee_WithNullAddress(t *testing.T) {
	denom := "testdenom"
	contract := common.Address{} // null address
	decimals := uint8(6)

	fee := NewFee(denom, contract, decimals)
	require.Equal(t, denom, fee.Denom())
	require.Equal(t, contract, fee.Contract())
	require.Equal(t, decimals, fee.Decimals())
	require.False(t, fee.HasContract())
}

func TestFee_WithEmptyDenom(t *testing.T) {
	denom := ""
	contract := common.HexToAddress("0x1234567890123456789012345678901234567890")
	decimals := uint8(8)

	fee := NewFee(denom, contract, decimals)
	require.Equal(t, denom, fee.Denom())
	require.Equal(t, contract, fee.Contract())
	require.Equal(t, decimals, fee.Decimals())
	require.True(t, fee.HasContract())
}

func TestFee_WithZeroDecimals(t *testing.T) {
	denom := "testdenom"
	contract := common.HexToAddress("0x1234567890123456789012345678901234567890")
	decimals := uint8(0)

	fee := NewFee(denom, contract, decimals)
	require.Equal(t, denom, fee.Denom())
	require.Equal(t, contract, fee.Contract())
	require.Equal(t, decimals, fee.Decimals())
	require.True(t, fee.HasContract())
}

func TestFee_WithMaxDecimals(t *testing.T) {
	denom := "testdenom"
	contract := common.HexToAddress("0x1234567890123456789012345678901234567890")
	decimals := uint8(255) // max uint8 value

	fee := NewFee(denom, contract, decimals)
	require.Equal(t, denom, fee.Denom())
	require.Equal(t, contract, fee.Contract())
	require.Equal(t, decimals, fee.Decimals())
	require.True(t, fee.HasContract())
}
