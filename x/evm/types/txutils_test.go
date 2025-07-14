package types

import (
	"encoding/json"
	"math/big"
	"testing"

	"cosmossdk.io/math"
	"github.com/ethereum/go-ethereum/common"
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
		params            *Params
		sender            common.Address
		gasLimit          uint64
		gasFeeCap         *big.Int
		expectedGasLimit  uint64
		expectedGasFeeCap *big.Int
	}{
		{
			name:              "no enforcement - return original values",
			params:            &Params{GasEnforcement: nil},
			sender:            common.HexToAddress("0x1234567890123456789012345678901234567890"),
			gasLimit:          100,
			gasFeeCap:         big.NewInt(100),
			expectedGasLimit:  100,
			expectedGasFeeCap: big.NewInt(100),
		},
		{
			name: "unlimited gas sender - return original values",
			params: &Params{
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
			params: &Params{
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
			params: &Params{
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
			params: &Params{
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
			params: &Params{
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
			params: &Params{
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
			params: &Params{
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
			params: &Params{
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
			actualGasLimit, actualGasFeeCap := getActualGasMetadata(tt.params, tt.sender, tt.gasLimit, tt.gasFeeCap)

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
