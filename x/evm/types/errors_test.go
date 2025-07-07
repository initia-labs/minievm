package types

import (
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/stretchr/testify/require"
)

func TestRevertError(t *testing.T) {
	t.Run("NewRevertError", func(t *testing.T) {
		revertData := []byte{0x01, 0x02, 0x03}
		err := NewRevertError(revertData)

		require.NotNil(t, err)
		revertErr, ok := err.(*RevertError)
		require.True(t, ok)
		require.Equal(t, ErrReverted, revertErr.ModError)
		require.Equal(t, revertData, revertErr.ret)
	})

	t.Run("RevertError_Ret", func(t *testing.T) {
		revertData := []byte{0x04, 0x05, 0x06}
		revertErr := &RevertError{
			ModError: ErrReverted,
			ret:      revertData,
		}

		require.Equal(t, revertData, revertErr.Ret())
	})

	t.Run("RevertError_Error_WithValidRevertData", func(t *testing.T) {
		// Create a valid revert reason
		revertData := NewRevertReason(ErrInvalidValue)

		revertErr := &RevertError{
			ModError: ErrReverted,
			ret:      revertData,
		}

		errorMsg := revertErr.Error()
		require.Contains(t, errorMsg, "revert:")
		require.Contains(t, errorMsg, "Invalid value")
	})

	t.Run("RevertError_Error_WithInvalidRevertData", func(t *testing.T) {
		// Use invalid revert data that can't be unpacked
		invalidRevertData := []byte{0x01, 0x02, 0x03}

		revertErr := &RevertError{
			ModError: ErrReverted,
			ret:      invalidRevertData,
		}

		errorMsg := revertErr.Error()
		require.Equal(t, ErrReverted.Error(), errorMsg)
	})

	t.Run("RevertError_Error_WithEmptyRevertData", func(t *testing.T) {
		revertErr := &RevertError{
			ModError: ErrReverted,
			ret:      []byte{},
		}

		errorMsg := revertErr.Error()
		require.Equal(t, ErrReverted.Error(), errorMsg)
	})
}

func TestNewRevertReason(t *testing.T) {
	t.Run("NewRevertReason_WithValidError", func(t *testing.T) {
		revertData := NewRevertReason(ErrInvalidValue)
		require.NotNil(t, revertData)
		require.NotEmpty(t, revertData)
		require.Equal(t, revertSelector, revertData[:4])
	})

	t.Run("NewRevertReason_WithCustomError", func(t *testing.T) {
		customErr := ErrEVMCallFailed
		revertData := NewRevertReason(customErr)
		require.NotNil(t, revertData)
		require.NotEmpty(t, revertData)
		require.Equal(t, revertSelector, revertData[:4])
	})

	t.Run("NewRevertReason_UnpackRoundTrip", func(t *testing.T) {
		revertData := NewRevertReason(ErrInvalidValue)

		// Try to unpack the revert data
		unpackedReason, err := abi.UnpackRevert(revertData)
		require.NoError(t, err)
		require.Contains(t, unpackedReason, "Invalid value")
	})
}

func TestRevertSelectorAndABI(t *testing.T) {
	t.Run("RevertSelector_IsCorrect", func(t *testing.T) {
		// The revert selector should be the first 4 bytes of the keccak256 hash of "Error(string)"
		expectedSelector := []byte{0x08, 0xc3, 0x79, 0xa0} // keccak256("Error(string)")[:4]
		require.Equal(t, expectedSelector, revertSelector)
	})

	t.Run("RevertABIArguments_IsInitialized", func(t *testing.T) {
		require.NotNil(t, revertABIArguments)
		require.Len(t, revertABIArguments, 1)
		require.Equal(t, "string", revertABIArguments[0].Type.String())
	})
}

func TestModErrorType(t *testing.T) {
	// Test that ModError is properly defined as an alias
	var modErr ModError
	require.Nil(t, modErr) // Should be nil by default

	// Test assignment
	modErr = ErrInvalidValue
	require.NotNil(t, modErr)
	require.Equal(t, ErrInvalidValue, modErr)
}
