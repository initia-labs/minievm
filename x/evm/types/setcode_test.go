package types

import (
	"math/big"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	coretypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/holiman/uint256"
	"github.com/stretchr/testify/require"
)

func TestConvertEthSetCodeAuthorizationsToCosmos(t *testing.T) {
	tests := []struct {
		name     string
		input    []coretypes.SetCodeAuthorization
		expected []SetCodeAuthorization
	}{
		{
			name:     "empty list",
			input:    []coretypes.SetCodeAuthorization{},
			expected: nil,
		},
		{
			name: "single authorization",
			input: []coretypes.SetCodeAuthorization{
				{
					ChainID: *uint256.NewInt(1),
					Address: common.HexToAddress("0x1234567890123456789012345678901234567890"),
					Nonce:   42,
					V:       27,
					R:       *uint256.MustFromBig(big.NewInt(12345)),
					S:       *uint256.MustFromBig(big.NewInt(67890)),
				},
			},
			expected: []SetCodeAuthorization{
				{
					ChainId:   "1",
					Address:   "0x1234567890123456789012345678901234567890",
					Nonce:     42,
					Signature: createSignatureBytes(uint256.MustFromBig(big.NewInt(12345)), uint256.MustFromBig(big.NewInt(67890)), 27),
				},
			},
		},
		{
			name: "multiple authorizations",
			input: []coretypes.SetCodeAuthorization{
				{
					ChainID: *uint256.NewInt(1),
					Address: common.HexToAddress("0x1111111111111111111111111111111111111111"),
					Nonce:   1,
					V:       27,
					R:       *uint256.MustFromBig(big.NewInt(111)),
					S:       *uint256.MustFromBig(big.NewInt(222)),
				},
				{
					ChainID: *uint256.NewInt(1),
					Address: common.HexToAddress("0x2222222222222222222222222222222222222222"),
					Nonce:   2,
					V:       28,
					R:       *uint256.MustFromBig(big.NewInt(333)),
					S:       *uint256.MustFromBig(big.NewInt(444)),
				},
			},
			expected: []SetCodeAuthorization{
				{
					ChainId:   "1",
					Address:   "0x1111111111111111111111111111111111111111",
					Nonce:     1,
					Signature: createSignatureBytes(uint256.MustFromBig(big.NewInt(111)), uint256.MustFromBig(big.NewInt(222)), 27),
				},
				{
					ChainId:   "1",
					Address:   "0x2222222222222222222222222222222222222222",
					Nonce:     2,
					Signature: createSignatureBytes(uint256.MustFromBig(big.NewInt(333)), uint256.MustFromBig(big.NewInt(444)), 28),
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ConvertEthSetCodeAuthorizationsToCosmos(tt.input)
			require.Equal(t, tt.expected, result)
		})
	}
}

func TestConvertCosmosSetCodeAuthorizationsToEth(t *testing.T) {
	tests := []struct {
		name        string
		input       []SetCodeAuthorization
		expected    []coretypes.SetCodeAuthorization
		expectError bool
	}{
		{
			name:        "empty list",
			input:       []SetCodeAuthorization{},
			expected:    nil,
			expectError: false,
		},
		{
			name: "single authorization",
			input: []SetCodeAuthorization{
				{
					ChainId:   "1",
					Address:   "0x1234567890123456789012345678901234567890",
					Nonce:     42,
					Signature: createSignatureBytes(uint256.MustFromBig(big.NewInt(12345)), uint256.MustFromBig(big.NewInt(67890)), 27),
				},
			},
			expected: []coretypes.SetCodeAuthorization{
				{
					ChainID: *uint256.NewInt(1),
					Address: common.HexToAddress("0x1234567890123456789012345678901234567890"),
					Nonce:   42,
					V:       27,
					R:       *uint256.MustFromBig(big.NewInt(12345)),
					S:       *uint256.MustFromBig(big.NewInt(67890)),
				},
			},
			expectError: false,
		},
		{
			name: "invalid signature length",
			input: []SetCodeAuthorization{
				{
					ChainId:   "1",
					Address:   "0x1234567890123456789012345678901234567890",
					Nonce:     42,
					Signature: []byte{1, 2, 3}, // Invalid length
				},
			},
			expected:    nil,
			expectError: true,
		},
		{
			name: "invalid address format",
			input: []SetCodeAuthorization{
				{
					ChainId:   "1",
					Address:   "invalid_address",
					Nonce:     42,
					Signature: make([]byte, 65), // Valid length but invalid content
				},
			},
			expected:    nil,
			expectError: true,
		},
		{
			name: "invalid chain ID",
			input: []SetCodeAuthorization{
				{
					ChainId:   "invalid_chain_id",
					Address:   "0x1234567890123456789012345678901234567890",
					Nonce:     42,
					Signature: make([]byte, 65),
				},
			},
			expected:    nil,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ConvertCosmosSetCodeAuthorizationsToEth(tt.input)
			if tt.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.expected, result)
			}
		})
	}
}

func TestSetCodeAuthorizationRoundTrip(t *testing.T) {
	// Test that converting from Ethereum to Cosmos and back preserves the data
	original := []coretypes.SetCodeAuthorization{
		{
			ChainID: *uint256.NewInt(1),
			Address: common.HexToAddress("0x1234567890123456789012345678901234567890"),
			Nonce:   42,
			V:       27,
			R:       *uint256.MustFromBig(big.NewInt(12345)),
			S:       *uint256.MustFromBig(big.NewInt(67890)),
		},
		{
			ChainID: *uint256.NewInt(1337),
			Address: common.HexToAddress("0xabcdefabcdefabcdefabcdefabcdefabcdefabcd"),
			Nonce:   100,
			V:       28,
			R:       *uint256.MustFromBig(big.NewInt(999999)),
			S:       *uint256.MustFromBig(big.NewInt(888888)),
		},
	}

	// Convert to Cosmos format
	cosmosAuths := ConvertEthSetCodeAuthorizationsToCosmos(original)
	require.Len(t, cosmosAuths, 2)

	// Convert back to Ethereum format
	ethAuths, err := ConvertCosmosSetCodeAuthorizationsToEth(cosmosAuths)
	require.NoError(t, err)
	require.Len(t, ethAuths, 2)

	// Verify the data matches
	for i, originalAuth := range original {
		convertedAuth := ethAuths[i]
		require.Equal(t, originalAuth.ChainID, convertedAuth.ChainID)
		require.Equal(t, originalAuth.Address, convertedAuth.Address)
		require.Equal(t, originalAuth.Nonce, convertedAuth.Nonce)
		require.Equal(t, originalAuth.V, convertedAuth.V)
		require.Equal(t, originalAuth.R, convertedAuth.R)
		require.Equal(t, originalAuth.S, convertedAuth.S)
	}
}

func TestValidateAuthorization(t *testing.T) {
	// Create valid signature values
	validR, _ := new(big.Int).SetString("1234567890123456789012345678901234567890123456789012345678901", 10)
	validS, _ := new(big.Int).SetString("1234567890123456789012345678901234567890123456789012345678901", 10)

	tests := []struct {
		name        string
		auth        coretypes.SetCodeAuthorization
		chainID     string
		expectError bool
		errorType   error
	}{
		{
			name: "valid authorization with matching chain ID",
			auth: coretypes.SetCodeAuthorization{
				ChainID: *uint256.MustFromBig(ConvertCosmosChainIDToEthereumChainID("minievm-1")),
				Address: common.HexToAddress("0x1234567890123456789012345678901234567890"),
				Nonce:   42,
				V:       1, // Valid v value (0 or 1)
				R:       *uint256.MustFromBig(validR),
				S:       *uint256.MustFromBig(validS),
			},
			chainID:     "minievm-1",
			expectError: false,
		},
		{
			name: "authorization with matching chain ID (signature validation will fail)",
			auth: coretypes.SetCodeAuthorization{
				ChainID: *uint256.MustFromBig(ConvertCosmosChainIDToEthereumChainID("minievm-1")),
				Address: common.HexToAddress("0x1234567890123456789012345678901234567890"),
				Nonce:   42,
				V:       27,
				R:       *uint256.MustFromBig(big.NewInt(12345)),
				S:       *uint256.MustFromBig(big.NewInt(67890)),
			},
			chainID:     "minievm-1",
			expectError: true, // Signature validation will fail
		},
		{
			name: "authorization with zero chain ID (signature validation will fail)",
			auth: coretypes.SetCodeAuthorization{
				ChainID: *uint256.NewInt(0),
				Address: common.HexToAddress("0x1234567890123456789012345678901234567890"),
				Nonce:   42,
				V:       27,
				R:       *uint256.MustFromBig(big.NewInt(12345)),
				S:       *uint256.MustFromBig(big.NewInt(67890)),
			},
			chainID:     "minievm-1",
			expectError: true, // Signature validation will fail
		},
		{
			name: "invalid authorization with wrong chain ID",
			auth: coretypes.SetCodeAuthorization{
				ChainID: *uint256.MustFromBig(ConvertCosmosChainIDToEthereumChainID("minievm-999")),
				Address: common.HexToAddress("0x1234567890123456789012345678901234567890"),
				Nonce:   42,
				V:       27,
				R:       *uint256.MustFromBig(big.NewInt(12345)),
				S:       *uint256.MustFromBig(big.NewInt(67890)),
			},
			chainID:     "minievm-1",
			expectError: true,
			errorType:   core.ErrAuthorizationWrongChainID,
		},
		{
			name: "invalid authorization with nonce overflow",
			auth: coretypes.SetCodeAuthorization{
				ChainID: *uint256.NewInt(0), // Use zero chain ID to avoid chain ID validation error
				Address: common.HexToAddress("0x1234567890123456789012345678901234567890"),
				Nonce:   ^uint64(0), // Max uint64
				V:       27,
				R:       *uint256.MustFromBig(big.NewInt(12345)),
				S:       *uint256.MustFromBig(big.NewInt(67890)),
			},
			chainID:     "minievm-1",
			expectError: true,
			errorType:   core.ErrAuthorizationNonceOverflow,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a mock context with the chain ID
			ctx := createMockContext(tt.chainID)

			err := ValidateAuthorization(ctx, tt.auth)
			if tt.expectError {
				require.Error(t, err)
				if tt.errorType != nil {
					require.ErrorIs(t, err, tt.errorType)
				}
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestSetCodeAuthorizationSignatureHandling(t *testing.T) {
	t.Run("signature byte manipulation", func(t *testing.T) {
		r := uint256.MustFromBig(big.NewInt(12345))
		s := uint256.MustFromBig(big.NewInt(67890))
		v := uint8(27)

		// Create signature bytes
		sigBytes := createSignatureBytes(r, s, v)
		require.Len(t, sigBytes, 65)

		// Verify we can extract the components back
		extractedR := sigBytes[:32]
		extractedS := sigBytes[32:64]
		extractedV := sigBytes[64]

		rBytes := r.Bytes32()
		sBytes := s.Bytes32()
		require.Equal(t, rBytes[:], extractedR)
		require.Equal(t, sBytes[:], extractedS)
		require.Equal(t, v, extractedV)
	})

	t.Run("edge case signature values", func(t *testing.T) {
		// Test with zero values
		r := uint256.NewInt(0)
		s := uint256.NewInt(0)
		v := uint8(0)

		sigBytes := createSignatureBytes(r, s, v)
		require.Len(t, sigBytes, 65)

		// Test with maximum values
		r = uint256.MustFromBig(new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 256), big.NewInt(1))) // Max uint256
		s = uint256.MustFromBig(new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 256), big.NewInt(1))) // Max uint256
		v = uint8(255)

		sigBytes = createSignatureBytes(r, s, v)
		require.Len(t, sigBytes, 65)
	})
}

func TestSetCodeAuthorizationWithDifferentChainIDs(t *testing.T) {
	chainIDs := []string{
		"minievm-1",
		"minievm-1337",
		"minievm-999999",
	}

	for _, chainID := range chainIDs {
		t.Run("chain_id_"+chainID, func(t *testing.T) {
			ctx := createMockContext(chainID)

			auth := coretypes.SetCodeAuthorization{
				ChainID: *uint256.NewInt(0), // Zero chain ID should be valid for any chain
				Address: common.HexToAddress("0x1234567890123456789012345678901234567890"),
				Nonce:   42,
				V:       27,
				R:       *uint256.MustFromBig(big.NewInt(12345)),
				S:       *uint256.MustFromBig(big.NewInt(67890)),
			}

			err := ValidateAuthorization(ctx, auth)
			// The signature validation will fail because we're using invalid signature values
			// This is expected behavior - the test validates that signature validation is working
			require.Error(t, err)
			require.Contains(t, err.Error(), "invalid transaction v, r, s values")
		})
	}
}

// Helper function to create signature bytes from r, s, v components
func createSignatureBytes(r, s *uint256.Int, v uint8) []byte {
	rBytes := r.Bytes32()
	sBytes := s.Bytes32()
	sigBytes := make([]byte, 65)
	copy(sigBytes[:32], rBytes[:])
	copy(sigBytes[32:64], sBytes[:])
	sigBytes[64] = v
	return sigBytes
}

// Helper function to create a mock context for testing
func createMockContext(chainID string) sdk.Context {
	// This is a simplified mock context creation
	// In a real test environment, you might want to use a more sophisticated mock
	return sdk.Context{}.WithChainID(chainID)
}
