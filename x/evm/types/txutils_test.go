package types

import (
	"encoding/json"
	"math/big"
	"testing"

	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/std"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/signing"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	coretypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/holiman/uint256"
	"github.com/stretchr/testify/require"
)

func Test_txMetadata_EncodingDecoding(t *testing.T) {
	meta := txMetadata{
		Type:      2,
		GasFeeCap: big.NewInt(0),
		GasTipCap: big.NewInt(0),
		GasLimit:  100,
	}

	bz, err := json.Marshal(meta)
	require.NoError(t, err)

	var meta2 txMetadata
	err = json.Unmarshal(bz, &meta2)
	require.NoError(t, err)
	require.Equal(t, meta, meta2)

	require.True(t, meta2.GasFeeCap.Uint64() == 0)
	require.True(t, meta2.GasTipCap.Uint64() == 0)
}

func Test_getActualGasMetadata(t *testing.T) {
	tests := []struct {
		name              string
		params            Params
		sender            common.Address
		gasLimit          uint64
		gasFeeCap         *big.Int
		expectedGasLimit  uint64
		expectedGasFeeCap *big.Int
	}{
		{
			name:              "no enforcement - return original values",
			params:            Params{GasEnforcement: nil},
			sender:            common.HexToAddress("0x1234567890123456789012345678901234567890"),
			gasLimit:          100,
			gasFeeCap:         big.NewInt(100),
			expectedGasLimit:  100,
			expectedGasFeeCap: big.NewInt(100),
		},
		{
			name: "unlimited gas sender - return original values",
			params: Params{
				GasEnforcement: &GasEnforcement{
					MaxGasLimit:         50,
					MaxGasFeeCap:        math.NewInt(50),
					UnlimitedGasSenders: []string{"0x1234567890123456789012345678901234567890"},
				},
			},
			sender:            common.HexToAddress("0x1234567890123456789012345678901234567890"),
			gasLimit:          100,
			gasFeeCap:         big.NewInt(100),
			expectedGasLimit:  100,
			expectedGasFeeCap: big.NewInt(100),
		},
		{
			name: "gas limit capped - below max",
			params: Params{
				GasEnforcement: &GasEnforcement{
					MaxGasLimit:         50,
					MaxGasFeeCap:        math.NewInt(100),
					UnlimitedGasSenders: []string{},
				},
			},
			sender:            common.HexToAddress("0x1234567890123456789012345678901234567890"),
			gasLimit:          30,
			gasFeeCap:         big.NewInt(100),
			expectedGasLimit:  30,
			expectedGasFeeCap: big.NewInt(100),
		},
		{
			name: "gas limit capped - above max",
			params: Params{
				GasEnforcement: &GasEnforcement{
					MaxGasLimit:         50,
					MaxGasFeeCap:        math.NewInt(100),
					UnlimitedGasSenders: []string{},
				},
			},
			sender:            common.HexToAddress("0x1234567890123456789012345678901234567890"),
			gasLimit:          100,
			gasFeeCap:         big.NewInt(100),
			expectedGasLimit:  50,
			expectedGasFeeCap: big.NewInt(100),
		},
		{
			name: "gas fee cap enforced - below max",
			params: Params{
				GasEnforcement: &GasEnforcement{
					MaxGasLimit:         100,
					MaxGasFeeCap:        math.NewInt(100),
					UnlimitedGasSenders: []string{},
				},
			},
			sender:            common.HexToAddress("0x1234567890123456789012345678901234567890"),
			gasLimit:          50,
			gasFeeCap:         big.NewInt(50),
			expectedGasLimit:  50,
			expectedGasFeeCap: big.NewInt(50),
		},
		{
			name: "gas fee cap enforced - above max",
			params: Params{
				GasEnforcement: &GasEnforcement{
					MaxGasLimit:         100,
					MaxGasFeeCap:        math.NewInt(50),
					UnlimitedGasSenders: []string{},
				},
			},
			sender:            common.HexToAddress("0x1234567890123456789012345678901234567890"),
			gasLimit:          50,
			gasFeeCap:         big.NewInt(100),
			expectedGasLimit:  50,
			expectedGasFeeCap: big.NewInt(50),
		},
		{
			name: "both gas limit and fee cap enforced",
			params: Params{
				GasEnforcement: &GasEnforcement{
					MaxGasLimit:         50,
					MaxGasFeeCap:        math.NewInt(50),
					UnlimitedGasSenders: []string{},
				},
			},
			sender:            common.HexToAddress("0x1234567890123456789012345678901234567890"),
			gasLimit:          100,
			gasFeeCap:         big.NewInt(100),
			expectedGasLimit:  50,
			expectedGasFeeCap: big.NewInt(50),
		},
		{
			name: "zero gas fee cap - no enforcement",
			params: Params{
				GasEnforcement: &GasEnforcement{
					MaxGasLimit:         50,
					MaxGasFeeCap:        math.ZeroInt(),
					UnlimitedGasSenders: []string{},
				},
			},
			sender:            common.HexToAddress("0x1234567890123456789012345678901234567890"),
			gasLimit:          100,
			gasFeeCap:         big.NewInt(100),
			expectedGasLimit:  50,
			expectedGasFeeCap: big.NewInt(100),
		},
		{
			name: "zero max gas limit - no limit enforcement",
			params: Params{
				GasEnforcement: &GasEnforcement{
					MaxGasLimit:         0,
					MaxGasFeeCap:        math.NewInt(50),
					UnlimitedGasSenders: []string{},
				},
			},
			sender:            common.HexToAddress("0x1234567890123456789012345678901234567890"),
			gasLimit:          100,
			gasFeeCap:         big.NewInt(100),
			expectedGasLimit:  100,
			expectedGasFeeCap: big.NewInt(50),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualGasLimit, actualGasFeeCap := applyGasEnforcement(tt.params, tt.sender, tt.gasLimit, tt.gasFeeCap)

			require.Equal(t, tt.expectedGasLimit, actualGasLimit)

			if tt.expectedGasFeeCap == nil {
				require.Nil(t, actualGasFeeCap)
			} else {
				require.NotNil(t, actualGasFeeCap)
				require.Equal(t, tt.expectedGasFeeCap.Cmp(actualGasFeeCap), 0)
			}
		})
	}
}

func Test_computeGasFeeAmount(t *testing.T) {
	t.Run("zero gas fee cap", func(t *testing.T) {
		amt := computeGasFeeAmount(big.NewInt(0), 100, 18)
		require.Equal(t, big.NewInt(0), amt)
	})

	t.Run("nonzero gas fee cap and gas", func(t *testing.T) {
		amt := computeGasFeeAmount(big.NewInt(1e18), 2, 18)
		// (1e18 * 2) = 2e18, FromEthersUnit(18, 2e18) = 2e18, +1 = 2000000000000000001
		require.Equal(t, big.NewInt(2000000000000000001), amt)
	})

	t.Run("rounding up", func(t *testing.T) {
		amt := computeGasFeeAmount(big.NewInt(3), 2, 0)
		// (3*2) = 6, FromEthersUnit(0, 6) = 0, +1 = 1
		require.Equal(t, big.NewInt(1), amt)
	})
}

func Test_ConvertCosmosAccessListToEth_and_ConvertEthAccessListToCosmos(t *testing.T) {
	cosmosList := []AccessTuple{
		{
			Address:     "0x1234567890123456789012345678901234567890",
			StorageKeys: []string{"0x1", "0x2"},
		},
		{
			Address:     "0xabcdefabcdefabcdefabcdefabcdefabcdefabcd",
			StorageKeys: []string{"0x3"},
		},
	}

	ethList := ConvertCosmosAccessListToEth(cosmosList)
	require.Len(t, ethList, 2)
	require.Equal(t, common.HexToAddress("0x1234567890123456789012345678901234567890"), ethList[0].Address)
	require.Equal(t, common.HexToHash("0x1"), ethList[0].StorageKeys[0])
	require.Equal(t, common.HexToHash("0x2"), ethList[0].StorageKeys[1])
	require.Equal(t, common.HexToAddress("0xabcdefabcdefabcdefabcdefabcdefabcdefabcd"), ethList[1].Address)
	require.Equal(t, common.HexToHash("0x3"), ethList[1].StorageKeys[0])

	cosmosList2 := ConvertEthAccessListToCosmos(ethList)
	// The round-trip will produce 32-byte padded storage keys and checksum addresses
	require.Equal(t, "0x1234567890123456789012345678901234567890", cosmosList2[0].Address)
	require.Equal(t, "0x0000000000000000000000000000000000000000000000000000000000000001", cosmosList2[0].StorageKeys[0])
	require.Equal(t, "0x0000000000000000000000000000000000000000000000000000000000000002", cosmosList2[0].StorageKeys[1])
	require.Equal(t, "0xABcdEFABcdEFabcdEfAbCdefabcdeFABcDEFabCD", cosmosList2[1].Address)
	require.Equal(t, "0x0000000000000000000000000000000000000000000000000000000000000003", cosmosList2[1].StorageKeys[0])

	// Test empty input
	require.Nil(t, ConvertCosmosAccessListToEth(nil))
	require.Nil(t, ConvertEthAccessListToCosmos(nil))
}

// Helper function to create a mock address codec
type mockAddressCodec struct{}

func (m *mockAddressCodec) StringToBytes(addrStr string) ([]byte, error) {
	return common.HexToAddress(addrStr).Bytes(), nil
}

func (m *mockAddressCodec) BytesToString(bz []byte) (string, error) {
	return common.BytesToAddress(bz).Hex(), nil
}

// Helper function to create a mock lazy args getter
func createMockLazyArgsGetter(params Params, feeDecimals uint8) LazyArgsGetterForConvertEthereumTxToCosmosTx {
	return func() (Params, uint8, error) {
		return params, feeDecimals, nil
	}
}

// Helper function to create a test Ethereum transaction with real signature
func createTestEthTx(txType uint8, to *common.Address, value *big.Int, data []byte, gasLimit uint64, gasPrice *big.Int, accessList coretypes.AccessList, authList []coretypes.SetCodeAuthorization) *coretypes.Transaction {
	var txData coretypes.TxData

	ethChainID := big.NewInt(3068811972085126) // Correct chain ID for minievm-1

	switch txType {
	case coretypes.LegacyTxType:
		txData = &coretypes.LegacyTx{
			Nonce:    0,
			GasPrice: gasPrice,
			Gas:      gasLimit,
			To:       to,
			Value:    value,
			Data:     data,
		}
	case coretypes.AccessListTxType:
		txData = &coretypes.AccessListTx{
			ChainID:    ethChainID,
			Nonce:      0,
			GasPrice:   gasPrice,
			Gas:        gasLimit,
			To:         to,
			Value:      value,
			Data:       data,
			AccessList: accessList,
		}
	case coretypes.DynamicFeeTxType:
		txData = &coretypes.DynamicFeeTx{
			ChainID:    ethChainID,
			Nonce:      0,
			GasTipCap:  gasPrice,
			GasFeeCap:  gasPrice,
			Gas:        gasLimit,
			To:         to,
			Value:      value,
			Data:       data,
			AccessList: accessList,
		}
	case coretypes.SetCodeTxType:
		txData = &coretypes.SetCodeTx{
			ChainID:    uint256.NewInt(ethChainID.Uint64()),
			Nonce:      0,
			GasTipCap:  uint256.MustFromBig(gasPrice),
			GasFeeCap:  uint256.MustFromBig(gasPrice),
			Gas:        gasLimit,
			To:         *to,
			Value:      uint256.MustFromBig(value),
			Data:       data,
			AccessList: accessList,
			AuthList:   authList,
		}
	default:
		// For unsupported transaction types, return a legacy transaction
		// This is used for testing error cases
		txData = &coretypes.LegacyTx{
			Nonce:    0,
			GasPrice: gasPrice,
			Gas:      gasLimit,
			To:       to,
			Value:    value,
			Data:     data,
		}
	}

	ethTx := coretypes.NewTx(txData)

	// Generate a real private key and sign the transaction
	privKey, _ := crypto.GenerateKey()
	signer := coretypes.LatestSignerForChainID(ethChainID)
	signedTx, _ := coretypes.SignTx(ethTx, signer, privKey)
	return signedTx
}

func TestTransactionConversion_RoundTrip(t *testing.T) {
	// Create test parameters
	params := Params{
		FeeDenom: "stake",
		GasEnforcement: &GasEnforcement{
			MaxGasFeeCap: math.NewIntFromBigInt(big.NewInt(1000)),
			MaxGasLimit:  1000000,
		},
	}
	feeDecimals := uint8(18)

	// Create mock codec
	cdc := codec.NewLegacyAmino()
	std.RegisterLegacyAminoCodec(cdc)
	interfaceRegistry := types.NewInterfaceRegistry()
	std.RegisterInterfaces(interfaceRegistry)
	cdcWithInterfaces := codec.NewProtoCodec(interfaceRegistry)

	// Create mock address codec
	ac := &mockAddressCodec{}

	// Create lazy args getter
	lazyArgsGetter := createMockLazyArgsGetter(params, feeDecimals)

	tests := []struct {
		name        string
		chainID     string
		ethTx       *coretypes.Transaction
		expectError bool
		errorMsg    string
	}{
		{
			name:    "legacy transaction - call",
			chainID: "minievm-1",
			ethTx: createTestEthTx(
				coretypes.LegacyTxType,
				&[]common.Address{common.HexToAddress("0x1234567890123456789012345678901234567890")}[0],
				big.NewInt(1000),
				[]byte{0x12, 0x34, 0x56},
				21000,
				big.NewInt(20000000000), // 20 gwei
				coretypes.AccessList{},  // empty access list
				nil,                     // no auth list for legacy
			),
			expectError: false,
		},
		{
			name:    "legacy transaction - create",
			chainID: "minievm-1",
			ethTx: createTestEthTx(
				coretypes.LegacyTxType,
				nil, // nil to indicates contract creation
				big.NewInt(0),
				[]byte{0x60, 0x60, 0x60, 0x60, 0x60, 0x60, 0x60, 0x60}, // simple contract bytecode
				100000,
				big.NewInt(20000000000),
				coretypes.AccessList{}, // empty access list
				nil,                    // no auth list for legacy
			),
			expectError: false,
		},
		{
			name:    "access list transaction",
			chainID: "minievm-1",
			ethTx: createTestEthTx(
				coretypes.AccessListTxType,
				&[]common.Address{common.HexToAddress("0x1234567890123456789012345678901234567890")}[0],
				big.NewInt(1000),
				[]byte{0x12, 0x34, 0x56},
				21000,
				big.NewInt(20000000000),
				coretypes.AccessList{ // actual access list items
					{
						Address: common.HexToAddress("0xabcdefabcdefabcdefabcdefabcdefabcdefabcd"),
						StorageKeys: []common.Hash{
							common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000001"),
							common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000002"),
						},
					},
					{
						Address: common.HexToAddress("0x9876543210987654321098765432109876543210"),
						StorageKeys: []common.Hash{
							common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000003"),
						},
					},
				},
				nil, // no auth list for access list tx
			),
			expectError: false,
		},
		{
			name:    "dynamic fee transaction",
			chainID: "minievm-1",
			ethTx: createTestEthTx(
				coretypes.DynamicFeeTxType,
				&[]common.Address{common.HexToAddress("0x1234567890123456789012345678901234567890")}[0],
				big.NewInt(1000),
				[]byte{0x12, 0x34, 0x56},
				21000,
				big.NewInt(20000000000),
				coretypes.AccessList{ // actual access list items
					{
						Address: common.HexToAddress("0xfedcba9876543210fedcba9876543210fedcba98"),
						StorageKeys: []common.Hash{
							common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000004"),
							common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000005"),
							common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000006"),
						},
					},
				},
				nil, // no auth list for dynamic fee tx
			),
			expectError: false,
		},
		{
			name:    "set code transaction",
			chainID: "minievm-1",
			ethTx: createTestEthTx(
				coretypes.SetCodeTxType,
				&[]common.Address{common.HexToAddress("0x1234567890123456789012345678901234567890")}[0],
				big.NewInt(1000),
				[]byte{0x12, 0x34, 0x56},
				21000,
				big.NewInt(20000000000),
				coretypes.AccessList{}, // empty access list
				[]coretypes.SetCodeAuthorization{ // simple auth list for testing
					{
						ChainID: *uint256.NewInt(3068811972085126), // minievm-1 chain ID
						Address: common.HexToAddress("0x1234567890123456789012345678901234567890"),
						Nonce:   42,
						V:       0,
						R:       *uint256.NewInt(1),
						S:       *uint256.NewInt(1),
					},
				},
			),
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ConvertEthereumTxToCosmosTx(
				tt.chainID,
				ac,
				cdcWithInterfaces,
				tt.ethTx,
				lazyArgsGetter,
			)

			if tt.expectError {
				require.Error(t, err)
				if tt.errorMsg != "" {
					require.Contains(t, err.Error(), tt.errorMsg)
				}
			} else {
				require.NoError(t, err)
				require.NotNil(t, result)

				// Verify the converted transaction has the expected structure
				msgs := result.GetMsgs()
				require.Len(t, msgs, 1)

				// Check if it's a call or create message
				if tt.ethTx.To() == nil {
					// Should be a create message
					createMsg, ok := msgs[0].(*MsgCreate)
					require.True(t, ok)
					require.Equal(t, hexutil.Encode(tt.ethTx.Data()), createMsg.Code)
					require.Equal(t, math.NewIntFromBigInt(tt.ethTx.Value()), createMsg.Value)
				} else {
					// Should be a call message
					callMsg, ok := msgs[0].(*MsgCall)
					require.True(t, ok)
					require.Equal(t, tt.ethTx.To().String(), callMsg.ContractAddr)
					require.Equal(t, hexutil.Encode(tt.ethTx.Data()), callMsg.Input)
					require.Equal(t, math.NewIntFromBigInt(tt.ethTx.Value()), callMsg.Value)
				}

				// Verify fee and gas limit
				authTx := result.(signing.Tx)
				fee := authTx.GetFee()
				require.Len(t, fee, 1)
				require.Equal(t, params.FeeDenom, fee[0].Denom)
				require.True(t, fee[0].Amount.IsPositive())

				// Verify memo contains metadata
				memo := authTx.GetMemo()
				require.NotEmpty(t, memo)

				var metadata txMetadata
				err = json.Unmarshal([]byte(memo), &metadata)
				require.NoError(t, err)
				require.Equal(t, tt.ethTx.Type(), metadata.Type)
				require.Equal(t, tt.ethTx.Gas(), metadata.GasLimit)

				// Test round-trip conversion: Cosmos → Ethereum
				ethTx2, senderAddr, err := ConvertCosmosTxToEthereumTx(
					true, // allowLegacy
					tt.chainID,
					ac,
					result,
					func() (Params, uint8, error) {
						return params, feeDecimals, nil
					},
				)
				require.NoError(t, err)
				require.NotNil(t, ethTx2)
				require.NotNil(t, senderAddr)

				// Verify the round-trip transaction matches the original
				equalEthTransaction(t, tt.ethTx, ethTx2)
			}
		})
	}
}

func equalEthTransaction(t *testing.T, expected, actual *coretypes.Transaction) {
	require.Equal(t, expected.ChainId(), actual.ChainId())
	require.Equal(t, expected.Nonce(), actual.Nonce())
	require.Equal(t, expected.GasTipCap(), actual.GasTipCap())
	require.Equal(t, expected.GasFeeCap(), actual.GasFeeCap())
	require.Equal(t, expected.Gas(), actual.Gas())
	require.Equal(t, expected.To(), actual.To())
	require.Equal(t, expected.Data(), actual.Data())
	require.Equal(t, expected.Value(), actual.Value())
	require.Equal(t, expected.Type(), actual.Type())
	require.Equal(t, expected.AccessList(), actual.AccessList())
	require.Equal(t, expected.SetCodeAuthorizations(), actual.SetCodeAuthorizations())
}

func TestTransactionConversion_ErrorCases(t *testing.T) {
	params := Params{FeeDenom: "stake"}
	feeDecimals := uint8(18)

	// Create proper codec
	cdc := codec.NewLegacyAmino()
	std.RegisterLegacyAminoCodec(cdc)
	interfaceRegistry := types.NewInterfaceRegistry()
	std.RegisterInterfaces(interfaceRegistry)
	cdcWithInterfaces := codec.NewProtoCodec(interfaceRegistry)

	ac := &mockAddressCodec{}

	// Test with invalid lazy args getter
	t.Run("lazy args getter error", func(t *testing.T) {
		invalidLazyArgsGetter := func() (Params, uint8, error) {
			return params, feeDecimals, sdkerrors.ErrInvalidRequest.Wrap("test error")
		}

		ethTx := createTestEthTx(
			coretypes.LegacyTxType,
			&[]common.Address{common.HexToAddress("0x1234567890123456789012345678901234567890")}[0],
			big.NewInt(1000),
			[]byte{0x12, 0x34, 0x56},
			21000,
			big.NewInt(20000000000),
			coretypes.AccessList{}, // empty access list
			nil,                    // no auth list for legacy
		)

		_, err := ConvertEthereumTxToCosmosTx(
			"minievm-1",
			ac,
			cdcWithInterfaces,
			ethTx,
			invalidLazyArgsGetter,
		)

		require.Error(t, err)
		require.Contains(t, err.Error(), "test error")
	})

	// Test with invalid sender recovery
	t.Run("invalid sender recovery", func(t *testing.T) {
		lazyArgsGetter := createMockLazyArgsGetter(params, feeDecimals)

		// Create a transaction with invalid signature that will fail sender recovery
		invalidTx := &coretypes.LegacyTx{
			Nonce:    0,
			GasPrice: big.NewInt(20000000000),
			Gas:      21000,
			To:       &[]common.Address{common.HexToAddress("0x1234567890123456789012345678901234567890")}[0],
			Value:    big.NewInt(1000),
			Data:     []byte{0x12, 0x34, 0x56},
			V:        big.NewInt(999), // Invalid V value
			R:        big.NewInt(999), // Invalid R value
			S:        big.NewInt(999), // Invalid S value
		}
		ethTx := coretypes.NewTx(invalidTx)

		_, err := ConvertEthereumTxToCosmosTx(
			"minievm-1",
			ac,
			cdcWithInterfaces,
			ethTx,
			lazyArgsGetter,
		)

		require.Error(t, err)
		require.Contains(t, err.Error(), "invalid chain id for signer")
	})
}

// TestTransactionConversion_ComprehensiveScenarios tests comprehensive scenarios with round-trip conversion similar to keeper tests
func TestTransactionConversion_ComprehensiveScenarios(t *testing.T) {
	// Create test parameters with gas enforcement
	params := Params{
		FeeDenom: "stake",
		GasEnforcement: &GasEnforcement{
			MaxGasFeeCap: math.NewIntFromBigInt(big.NewInt(1000000000000000000)), // 1 ETH in wei
			MaxGasLimit:  500000,                                                 // Half of test gas limit
		},
	}
	feeDecimals := uint8(18)

	// Create mock codec
	cdc := codec.NewLegacyAmino()
	std.RegisterLegacyAminoCodec(cdc)
	interfaceRegistry := types.NewInterfaceRegistry()
	std.RegisterInterfaces(interfaceRegistry)
	cdcWithInterfaces := codec.NewProtoCodec(interfaceRegistry)

	// Create mock address codec
	ac := &mockAddressCodec{}

	// Create lazy args getter
	lazyArgsGetter := createMockLazyArgsGetter(params, feeDecimals)

	t.Run("dynamic fee transaction with access list", func(t *testing.T) {
		// Create a dynamic fee transaction similar to the keeper test
		gasLimit := uint64(1000000)
		gasFeeCap := big.NewInt(20000000000)     // 20 gwei
		gasTipCap := big.NewInt(1000000000)      // 1 gwei
		value := big.NewInt(1000000000000000000) // 1 ETH

		// Create access list
		accessList := coretypes.AccessList{
			coretypes.AccessTuple{
				Address: common.HexToAddress("0x1234567890123456789012345678901234567890"),
				StorageKeys: []common.Hash{
					common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
					common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000001"),
					common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000002"),
				},
			},
		}

		// Create dynamic fee transaction
		dynTx := &coretypes.DynamicFeeTx{
			ChainID:    big.NewInt(3068811972085126), // Correct chain ID for minievm-1
			Nonce:      100,
			GasTipCap:  gasTipCap,
			GasFeeCap:  gasFeeCap,
			Gas:        gasLimit,
			To:         &[]common.Address{common.HexToAddress("0x1234567890123456789012345678901234567890")}[0],
			Data:       []byte{0x12, 0x34, 0x56, 0x78},
			Value:      value,
			AccessList: accessList,
		}

		ethTx := coretypes.NewTx(dynTx)

		// Generate a real private key and sign the transaction
		privKey, err := crypto.GenerateKey()
		require.NoError(t, err)

		ethChainID := big.NewInt(3068811972085126)
		signer := coretypes.LatestSignerForChainID(ethChainID)
		signedTx, err := coretypes.SignTx(ethTx, signer, privKey)
		require.NoError(t, err)

		// Test conversion with real signature
		result, err := ConvertEthereumTxToCosmosTx(
			"minievm-1",
			ac,
			cdcWithInterfaces,
			signedTx,
			lazyArgsGetter,
		)

		// Should succeed with real signature
		require.NoError(t, err)
		require.NotNil(t, result)

		// Verify the converted transaction structure
		msgs := result.GetMsgs()
		require.Len(t, msgs, 1)

		callMsg, ok := msgs[0].(*MsgCall)
		require.True(t, ok)
		require.Equal(t, "0x1234567890123456789012345678901234567890", callMsg.ContractAddr)
		require.Equal(t, hexutil.Encode([]byte{0x12, 0x34, 0x56, 0x78}), callMsg.Input)
		require.Equal(t, math.NewIntFromBigInt(value), callMsg.Value)

		// Verify access list conversion
		require.Len(t, callMsg.AccessList, 1)
		require.Equal(t, "0x1234567890123456789012345678901234567890", callMsg.AccessList[0].Address)
		require.Len(t, callMsg.AccessList[0].StorageKeys, 3)

		// Verify fee and gas limit
		authTx := result.(signing.Tx)
		fee := authTx.GetFee()
		require.Len(t, fee, 1)
		require.Equal(t, params.FeeDenom, fee[0].Denom)
		require.True(t, fee[0].Amount.IsPositive())

		// Verify memo contains metadata
		memo := authTx.GetMemo()
		require.NotEmpty(t, memo)

		var metadata txMetadata
		err = json.Unmarshal([]byte(memo), &metadata)
		require.NoError(t, err)
		require.Equal(t, signedTx.Type(), metadata.Type)
		require.Equal(t, signedTx.Gas(), metadata.GasLimit)

		// Test round-trip conversion: Cosmos → Ethereum
		ethTx2, senderAddr, err := ConvertCosmosTxToEthereumTx(
			true, // allowLegacy
			"minievm-1",
			ac,
			result,
			func() (Params, uint8, error) {
				return params, feeDecimals, nil
			},
		)
		require.NoError(t, err)
		require.NotNil(t, ethTx2)
		require.NotNil(t, senderAddr)

		// Verify the round-trip transaction matches the original
		equalEthTransaction(t, signedTx, ethTx2)
	})

	t.Run("access list transaction with empty access list", func(t *testing.T) {
		// Create access list transaction with empty access list
		gasLimit := uint64(1000000)
		gasPrice := big.NewInt(20000000000)      // 20 gwei
		value := big.NewInt(1000000000000000000) // 1 ETH

		accessTx := &coretypes.AccessListTx{
			ChainID:    big.NewInt(3068811972085126), // Correct chain ID for minievm-1
			Nonce:      100,
			GasPrice:   gasPrice,
			Gas:        gasLimit,
			To:         &[]common.Address{common.HexToAddress("0x1234567890123456789012345678901234567890")}[0],
			Data:       []byte{0x12, 0x34, 0x56, 0x78},
			Value:      value,
			AccessList: coretypes.AccessList{}, // Empty access list
		}

		ethTx := coretypes.NewTx(accessTx)

		// Generate a real private key and sign the transaction
		privKey, err := crypto.GenerateKey()
		require.NoError(t, err)

		ethChainID := big.NewInt(3068811972085126)
		signer := coretypes.LatestSignerForChainID(ethChainID)
		signedTx, err := coretypes.SignTx(ethTx, signer, privKey)
		require.NoError(t, err)

		// Test conversion with real signature
		result, err := ConvertEthereumTxToCosmosTx(
			"minievm-1",
			ac,
			cdcWithInterfaces,
			signedTx,
			lazyArgsGetter,
		)

		// Should succeed with real signature
		require.NoError(t, err)
		require.NotNil(t, result)

		// Verify the converted transaction structure
		msgs := result.GetMsgs()
		require.Len(t, msgs, 1)

		callMsg, ok := msgs[0].(*MsgCall)
		require.True(t, ok)
		require.Equal(t, "0x1234567890123456789012345678901234567890", callMsg.ContractAddr)
		require.Equal(t, hexutil.Encode([]byte{0x12, 0x34, 0x56, 0x78}), callMsg.Input)
		require.Equal(t, math.NewIntFromBigInt(value), callMsg.Value)

		// Verify empty access list
		require.Len(t, callMsg.AccessList, 0)

		// Test round-trip conversion: Cosmos → Ethereum
		ethTx2, senderAddr, err := ConvertCosmosTxToEthereumTx(
			true, // allowLegacy
			"minievm-1",
			ac,
			result,
			func() (Params, uint8, error) {
				return params, feeDecimals, nil
			},
		)
		require.NoError(t, err)
		require.NotNil(t, ethTx2)
		require.NotNil(t, senderAddr)

		// Verify the round-trip transaction matches the original
		equalEthTransaction(t, signedTx, ethTx2)
	})

	t.Run("legacy transaction with gas enforcement", func(t *testing.T) {
		// Create legacy transaction
		gasLimit := uint64(1000000)
		gasPrice := big.NewInt(20000000000)      // 20 gwei
		value := big.NewInt(1000000000000000000) // 1 ETH

		signedTx := createTestEthTx(
			coretypes.LegacyTxType,
			&[]common.Address{common.HexToAddress("0x1234567890123456789012345678901234567890")}[0],
			value,
			[]byte{0x12, 0x34, 0x56, 0x78},
			gasLimit,
			gasPrice,
			coretypes.AccessList{}, // empty access list
			nil,                    // no auth list for legacy
		)

		// Test conversion with real signature
		result, err := ConvertEthereumTxToCosmosTx(
			"minievm-1",
			ac,
			cdcWithInterfaces,
			signedTx,
			lazyArgsGetter,
		)

		// Should succeed with real signature
		require.NoError(t, err)
		require.NotNil(t, result)

		// Verify the converted transaction structure
		msgs := result.GetMsgs()
		require.Len(t, msgs, 1)

		callMsg, ok := msgs[0].(*MsgCall)
		require.True(t, ok)
		require.Equal(t, "0x1234567890123456789012345678901234567890", callMsg.ContractAddr)
		require.Equal(t, hexutil.Encode([]byte{0x12, 0x34, 0x56, 0x78}), callMsg.Input)
		require.Equal(t, math.NewIntFromBigInt(value), callMsg.Value)

		// Verify fee and gas limit
		authTx := result.(signing.Tx)
		fee := authTx.GetFee()
		require.Len(t, fee, 1)
		require.Equal(t, params.FeeDenom, fee[0].Denom)
		require.True(t, fee[0].Amount.IsPositive())

		// Test round-trip conversion: Cosmos → Ethereum
		ethTx2, senderAddr, err := ConvertCosmosTxToEthereumTx(
			true, // allowLegacy
			"minievm-1",
			ac,
			result,
			func() (Params, uint8, error) {
				return params, feeDecimals, nil
			},
		)
		require.NoError(t, err)
		require.NotNil(t, ethTx2)
		require.NotNil(t, senderAddr)

		// Verify the round-trip transaction matches the original
		equalEthTransaction(t, signedTx, ethTx2)
	})

	t.Run("set code transaction with authorization list", func(t *testing.T) {
		// Create set code transaction with authorization list
		gasLimit := uint64(1000000)
		gasFeeCap := big.NewInt(20000000000)     // 20 gwei
		gasTipCap := big.NewInt(1000000000)      // 1 gwei
		value := big.NewInt(1000000000000000000) // 1 ETH

		// Create authorization list with real signature
		authList := []coretypes.SetCodeAuthorization{
			{
				ChainID: *uint256.NewInt(3068811972085126), // Correct chain ID for minievm-1
				Address: common.HexToAddress("0x1234567890123456789012345678901234567890"),
				Nonce:   42,
				V:       0,                  // Will be set by signing
				R:       *uint256.NewInt(0), // Will be set by signing
				S:       *uint256.NewInt(0), // Will be set by signing
			},
		}

		setCodeTx := &coretypes.SetCodeTx{
			ChainID:    uint256.NewInt(3068811972085126), // Correct chain ID for minievm-1
			Nonce:      100,
			GasTipCap:  uint256.MustFromBig(gasTipCap),
			GasFeeCap:  uint256.MustFromBig(gasFeeCap),
			Gas:        gasLimit,
			To:         common.HexToAddress("0x1234567890123456789012345678901234567890"),
			Data:       []byte{0x12, 0x34, 0x56, 0x78},
			Value:      uint256.MustFromBig(value),
			AccessList: coretypes.AccessList{},
			AuthList:   authList,
			V:          uint256.NewInt(0), // Will be set by signing
			R:          uint256.NewInt(0), // Will be set by signing
			S:          uint256.NewInt(0), // Will be set by signing
		}

		ethTx := coretypes.NewTx(setCodeTx)

		// Generate a real private key and sign the transaction
		privKey, err := crypto.GenerateKey()
		require.NoError(t, err)

		ethChainID := big.NewInt(3068811972085126)
		signer := coretypes.LatestSignerForChainID(ethChainID)
		signedTx, err := coretypes.SignTx(ethTx, signer, privKey)
		require.NoError(t, err)

		// Test conversion with real signature
		result, err := ConvertEthereumTxToCosmosTx(
			"minievm-1",
			ac,
			cdcWithInterfaces,
			signedTx,
			lazyArgsGetter,
		)

		// Should succeed with real signature
		require.NoError(t, err)
		require.NotNil(t, result)

		// Verify the converted transaction structure
		msgs := result.GetMsgs()
		require.Len(t, msgs, 1)

		callMsg, ok := msgs[0].(*MsgCall)
		require.True(t, ok)
		require.Equal(t, "0x1234567890123456789012345678901234567890", callMsg.ContractAddr)
		require.Equal(t, hexutil.Encode([]byte{0x12, 0x34, 0x56, 0x78}), callMsg.Input)
		require.Equal(t, math.NewIntFromBigInt(value), callMsg.Value)

		// Verify fee and gas limit
		authTx := result.(signing.Tx)
		fee := authTx.GetFee()
		require.Len(t, fee, 1)
		require.Equal(t, params.FeeDenom, fee[0].Denom)
		require.True(t, fee[0].Amount.IsPositive())

		// Test round-trip conversion: Cosmos → Ethereum
		ethTx2, senderAddr, err := ConvertCosmosTxToEthereumTx(
			true, // allowLegacy
			"minievm-1",
			ac,
			result,
			func() (Params, uint8, error) {
				return params, feeDecimals, nil
			},
		)
		require.NoError(t, err)
		require.NotNil(t, ethTx2)
		require.NotNil(t, senderAddr)

		// Verify the round-trip transaction matches the original
		equalEthTransaction(t, signedTx, ethTx2)
	})

	t.Run("contract creation transaction", func(t *testing.T) {
		// Create contract creation transaction (To is nil)
		gasLimit := uint64(1000000)
		gasPrice := big.NewInt(20000000000) // 20 gwei
		value := big.NewInt(0)              // No value for contract creation

		// Contract bytecode
		contractCode := []byte{
			0x60, 0x60, 0x60, 0x60, 0x60, 0x60, 0x60, 0x60, // PUSH1 0x60 (8 times)
			0x60, 0x00, 0x52, 0x60, 0x20, 0x60, 0x00, 0xf3, // MSTORE, RETURN
		}

		signedTx := createTestEthTx(
			coretypes.LegacyTxType,
			nil, // Contract creation
			value,
			contractCode,
			gasLimit,
			gasPrice,
			coretypes.AccessList{}, // empty access list
			nil,                    // no auth list for legacy
		)

		// Test conversion with real signature
		result, err := ConvertEthereumTxToCosmosTx(
			"minievm-1",
			ac,
			cdcWithInterfaces,
			signedTx,
			lazyArgsGetter,
		)

		// Should succeed with real signature
		require.NoError(t, err)
		require.NotNil(t, result)

		// Verify the converted transaction structure
		msgs := result.GetMsgs()
		require.Len(t, msgs, 1)

		deployMsg, ok := msgs[0].(*MsgCreate)
		require.True(t, ok)
		require.Equal(t, hexutil.Encode(contractCode), deployMsg.Code)
		require.Equal(t, math.NewIntFromBigInt(value), deployMsg.Value)

		// Verify fee and gas limit
		authTx := result.(signing.Tx)
		fee := authTx.GetFee()
		require.Len(t, fee, 1)
		require.Equal(t, params.FeeDenom, fee[0].Denom)
		require.True(t, fee[0].Amount.IsPositive())

		// Test round-trip conversion: Cosmos → Ethereum
		ethTx2, senderAddr, err := ConvertCosmosTxToEthereumTx(
			true, // allowLegacy
			"minievm-1",
			ac,
			result,
			func() (Params, uint8, error) {
				return params, feeDecimals, nil
			},
		)
		require.NoError(t, err)
		require.NotNil(t, ethTx2)
		require.NotNil(t, senderAddr)

		// Verify the round-trip transaction matches the original
		equalEthTransaction(t, signedTx, ethTx2)
	})

	t.Run("transaction with zero gas price", func(t *testing.T) {
		// Create transaction with zero gas price
		gasLimit := uint64(1000000)
		gasPrice := big.NewInt(0)                // Zero gas price
		value := big.NewInt(1000000000000000000) // 1 ETH

		signedTx := createTestEthTx(
			coretypes.LegacyTxType,
			&[]common.Address{common.HexToAddress("0x1234567890123456789012345678901234567890")}[0],
			value,
			[]byte{0x12, 0x34, 0x56, 0x78},
			gasLimit,
			gasPrice,
			coretypes.AccessList{}, // empty access list
			nil,                    // no auth list for legacy
		)

		// Test conversion with real signature
		result, err := ConvertEthereumTxToCosmosTx(
			"minievm-1",
			ac,
			cdcWithInterfaces,
			signedTx,
			lazyArgsGetter,
		)

		// Should succeed with real signature
		require.NoError(t, err)
		require.NotNil(t, result)

		// Verify the converted transaction structure
		msgs := result.GetMsgs()
		require.Len(t, msgs, 1)

		callMsg, ok := msgs[0].(*MsgCall)
		require.True(t, ok)
		require.Equal(t, "0x1234567890123456789012345678901234567890", callMsg.ContractAddr)
		require.Equal(t, hexutil.Encode([]byte{0x12, 0x34, 0x56, 0x78}), callMsg.Input)
		require.Equal(t, math.NewIntFromBigInt(value), callMsg.Value)

		// Verify fee is zero due to zero gas price
		authTx := result.(signing.Tx)
		fee := authTx.GetFee()
		// When gas price is zero, fee might be empty or contain zero amount
		if len(fee) == 0 {
			// Fee array is empty when gas price is zero
			require.Len(t, fee, 0)
		} else {
			require.Len(t, fee, 1)
			require.Equal(t, params.FeeDenom, fee[0].Denom)
			require.True(t, fee[0].Amount.IsZero())
		}

		// Test round-trip conversion: Cosmos → Ethereum
		ethTx2, senderAddr, err := ConvertCosmosTxToEthereumTx(
			true, // allowLegacy
			"minievm-1",
			ac,
			result,
			func() (Params, uint8, error) {
				return params, feeDecimals, nil
			},
		)
		require.NoError(t, err)
		require.NotNil(t, ethTx2)
		require.NotNil(t, senderAddr)

		// Verify the round-trip transaction matches the original
		equalEthTransaction(t, signedTx, ethTx2)
	})

	t.Run("transaction with very large gas limit", func(t *testing.T) {
		// Create transaction with gas limit exceeding max
		gasLimit := uint64(2000000)              // Exceeds MaxGasLimit of 1000000
		gasPrice := big.NewInt(20000000000)      // 20 gwei
		value := big.NewInt(1000000000000000000) // 1 ETH

		signedTx := createTestEthTx(
			coretypes.LegacyTxType,
			&[]common.Address{common.HexToAddress("0x1234567890123456789012345678901234567890")}[0],
			value,
			[]byte{0x12, 0x34, 0x56, 0x78},
			gasLimit,
			gasPrice,
			coretypes.AccessList{}, // empty access list
			nil,                    // no auth list for legacy
		)

		// Test conversion with real signature
		result, err := ConvertEthereumTxToCosmosTx(
			"minievm-1",
			ac,
			cdcWithInterfaces,
			signedTx,
			lazyArgsGetter,
		)

		// Should succeed with real signature (gas enforcement happens in keeper layer)
		require.NoError(t, err)
		require.NotNil(t, result)

		// Verify the converted transaction structure
		msgs := result.GetMsgs()
		require.Len(t, msgs, 1)

		callMsg, ok := msgs[0].(*MsgCall)
		require.True(t, ok)
		require.Equal(t, "0x1234567890123456789012345678901234567890", callMsg.ContractAddr)
		require.Equal(t, hexutil.Encode([]byte{0x12, 0x34, 0x56, 0x78}), callMsg.Input)
		require.Equal(t, math.NewIntFromBigInt(value), callMsg.Value)

		// Verify fee calculation with large gas limit
		authTx := result.(signing.Tx)
		fee := authTx.GetFee()
		require.Len(t, fee, 1)
		require.Equal(t, params.FeeDenom, fee[0].Denom)
		require.True(t, fee[0].Amount.IsPositive())

		// Test round-trip conversion: Cosmos → Ethereum
		ethTx2, senderAddr, err := ConvertCosmosTxToEthereumTx(
			true, // allowLegacy
			"minievm-1",
			ac,
			result,
			func() (Params, uint8, error) {
				return params, feeDecimals, nil
			},
		)
		require.NoError(t, err)
		require.NotNil(t, ethTx2)
		require.NotNil(t, senderAddr)

		// Verify the round-trip transaction matches the original
		equalEthTransaction(t, signedTx, ethTx2)
	})
}
