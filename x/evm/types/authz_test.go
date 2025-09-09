package types

import (
	"testing"
	"time"

	"cosmossdk.io/log"
	"cosmossdk.io/math"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestNewCallAuthorization(t *testing.T) {
	contracts := []string{
		"0x1234567890123456789012345678901234567890",
		"0xabcdef1234567890abcdef1234567890abcdef12",
	}

	auth := NewCallAuthorization(contracts)
	require.Equal(t, contracts, auth.Contracts)
}

func TestCallAuthorization_MsgTypeURL(t *testing.T) {
	auth := CallAuthorization{}
	expectedURL := sdk.MsgTypeURL(&MsgCall{})
	require.Equal(t, expectedURL, auth.MsgTypeURL())
}

func TestCallAuthorization_ValidateBasic(t *testing.T) {
	// Use proper checksum addresses
	validContract1 := common.HexToAddress("0x1234567890123456789012345678901234567890").Hex()

	tests := []struct {
		name        string
		contracts   []string
		expectError bool
		errorMsg    string
	}{
		{
			name:        "valid contracts",
			contracts:   []string{validContract1},
			expectError: false,
		},
		{
			name:        "empty contracts",
			contracts:   []string{},
			expectError: false,
		},
		{
			name:        "invalid contract address",
			contracts:   []string{"invalid_address"},
			expectError: true,
			errorMsg:    "invalid contract address",
		},
		{
			name:        "non-checksum address",
			contracts:   []string{"0xabcdef1234567890abcdef1234567890abcdef12"},
			expectError: true,
			errorMsg:    "address must be in ethereum checksum hex format",
		},
		{
			name:        "mixed valid and invalid",
			contracts:   []string{validContract1, "invalid"},
			expectError: true,
			errorMsg:    "invalid contract address",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			auth := CallAuthorization{Contracts: tt.contracts}
			err := auth.ValidateBasic()

			if tt.expectError {
				require.Error(t, err)
				if tt.errorMsg != "" {
					require.Contains(t, err.Error(), tt.errorMsg)
				}
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestCallAuthorization_Accept(t *testing.T) {
	// Setup valid contracts for testing using checksum addresses
	validContract1 := common.HexToAddress("0x1234567890123456789012345678901234567890").Hex()
	validContract2 := common.HexToAddress("0xABCDEF1234567890ABCDEF1234567890ABCDEF12").Hex()

	auth := CallAuthorization{
		Contracts: []string{validContract1, validContract2},
	}

	tests := []struct {
		name         string
		msg          sdk.Msg
		expectError  bool
		expectAccept bool
		errorMsg     string
	}{
		{
			name: "valid MsgCall with authorized contract",
			msg: &MsgCall{
				Sender:       "init1test",
				ContractAddr: validContract1,
				Input:        "0x",
				Value:        math.ZeroInt(),
			},
			expectError:  false,
			expectAccept: true,
		},
		{
			name: "valid MsgCall with second authorized contract",
			msg: &MsgCall{
				Sender:       "init1test",
				ContractAddr: validContract2,
				Input:        "0x",
				Value:        math.ZeroInt(),
			},
			expectError:  false,
			expectAccept: true,
		},
		{
			name: "valid MsgCall with unauthorized contract",
			msg: &MsgCall{
				Sender:       "init1test",
				ContractAddr: "0x9999999999999999999999999999999999999999",
				Input:        "0x",
				Value:        math.ZeroInt(),
			},
			expectError:  true,
			expectAccept: false,
			errorMsg:     "unauthorized",
		},
		{
			name: "MsgCall with invalid contract address",
			msg: &MsgCall{
				Sender:       "init1test",
				ContractAddr: "invalid_address",
				Input:        "0x",
				Value:        math.ZeroInt(),
			},
			expectError:  true,
			expectAccept: false,
			errorMsg:     "invalid contract address",
		},
		{
			name:         "wrong message type",
			msg:          &MsgCreate{},
			expectError:  true,
			expectAccept: false,
			errorMsg:     "unknown msg type",
		},
		{
			name:         "nil message",
			msg:          nil,
			expectError:  true,
			expectAccept: false,
			errorMsg:     "unknown msg type",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := sdk.NewContext(nil, tmproto.Header{
				Height: 1,
				Time:   time.Date(2020, time.April, 22, 12, 0, 0, 0, time.UTC),
			}, false, log.NewNopLogger())
			response, err := auth.Accept(ctx, tt.msg)

			if tt.expectError {
				require.Error(t, err)
				if tt.errorMsg != "" {
					require.Contains(t, err.Error(), tt.errorMsg)
				}
			} else {
				require.NoError(t, err)
			}

			if tt.expectAccept {
				require.True(t, response.Accept)
			} else {
				require.False(t, response.Accept)
			}
		})
	}
}

func TestCallAuthorization_Accept_GasConsumption(t *testing.T) {
	// Test that gas is consumed during contract address iteration
	validContract := common.HexToAddress("0x1234567890123456789012345678901234567890").Hex()
	auth := CallAuthorization{
		Contracts: []string{validContract},
	}

	msg := &MsgCall{
		Sender:       "init1test",
		ContractAddr: validContract,
		Input:        "0x",
		Value:        math.ZeroInt(),
	}

	// Create a context with a gas meter
	ctx := sdk.NewContext(nil, tmproto.Header{
		Height: 1,
		Time:   time.Date(2020, time.April, 22, 12, 0, 0, 0, time.UTC),
	}, false, log.NewNopLogger())
	initialGas := ctx.GasMeter().GasConsumed()

	_, err := auth.Accept(ctx, msg)
	require.NoError(t, err)

	// Check that gas was consumed
	finalGas := ctx.GasMeter().GasConsumed()
	require.Greater(t, finalGas, initialGas)
}

func TestCallAuthorization_InterfaceCompliance(t *testing.T) {
	// Test that CallAuthorization implements the authz.Authorization interface
	var _ authz.Authorization = &CallAuthorization{}
}
func TestCallAuthorization_EmptyContracts(t *testing.T) {
	// Test that empty contracts list allows all contract calls
	auth := CallAuthorization{
		Contracts: []string{},
	}

	testCases := []struct {
		name         string
		contractAddr string
	}{
		{
			name:         "random contract address",
			contractAddr: common.HexToAddress("0x1234567890123456789012345678901234567890").Hex(),
		},
		{
			name:         "another contract address",
			contractAddr: common.HexToAddress("0xABCDEF1234567890ABCDEF1234567890ABCDEF12").Hex(),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			msg := &MsgCall{
				Sender:       "init1test",
				ContractAddr: tc.contractAddr,
				Input:        "0x",
				Value:        math.ZeroInt(),
			}

			ctx := sdk.NewContext(nil, tmproto.Header{
				Height: 1,
				Time:   time.Date(2020, time.April, 22, 12, 0, 0, 0, time.UTC),
			}, false, log.NewNopLogger())

			response, err := auth.Accept(ctx, msg)
			require.NoError(t, err)
			require.True(t, response.Accept, "Empty contracts list should authorize all contracts")
		})
	}
}

func TestCallAuthorization_WithMultipleContracts(t *testing.T) {
	// Test authorization with multiple contracts using checksum addresses
	contracts := []string{
		common.HexToAddress("0x1234567890123456789012345678901234567890").Hex(),
		common.HexToAddress("0xABCDEF1234567890ABCDEF1234567890ABCDEF12").Hex(),
		common.HexToAddress("0x9999999999999999999999999999999999999999").Hex(),
	}

	auth := CallAuthorization{Contracts: contracts}

	// Test each contract
	for i, contract := range contracts {
		msg := &MsgCall{
			Sender:       "init1test",
			ContractAddr: contract,
			Input:        "0x",
			Value:        math.ZeroInt(),
		}

		ctx := sdk.NewContext(nil, tmproto.Header{
			Height: 1,
			Time:   time.Date(2020, time.April, 22, 12, 0, 0, 0, time.UTC),
		}, false, log.NewNopLogger())
		response, err := auth.Accept(ctx, msg)

		require.NoError(t, err)
		require.True(t, response.Accept, "Contract %d should be authorized", i)
	}

	// Test unauthorized contract
	unauthorizedMsg := &MsgCall{
		Sender:       "init1test",
		ContractAddr: "0x1111111111111111111111111111111111111111",
		Input:        "0x",
		Value:        math.ZeroInt(),
	}

	ctx := sdk.NewContext(nil, tmproto.Header{
		Height: 1,
		Time:   time.Date(2020, time.April, 22, 12, 0, 0, 0, time.UTC),
	}, false, log.NewNopLogger())
	response, err := auth.Accept(ctx, unauthorizedMsg)

	require.Error(t, err)
	require.False(t, response.Accept)
	require.Contains(t, err.Error(), "unauthorized")
}
