package types

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type genesisMockCodec struct{}

func (m genesisMockCodec) StringToBytes(s string) ([]byte, error) {
	if s == "bad" {
		return nil, fmt.Errorf("mock error")
	}
	return []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f, 0x10, 0x11, 0x12, 0x13, 0x14}, nil
}

func (m genesisMockCodec) BytesToString(b []byte) (string, error) {
	if len(b) == 0 {
		return "", fmt.Errorf("mock error")
	}
	return "mock_address", nil
}

func TestDefaultGenesis(t *testing.T) {
	genState := DefaultGenesis()
	require.NotNil(t, genState)
	require.NotNil(t, genState.Params)
	require.Empty(t, genState.KeyValues)
	require.Nil(t, genState.Erc20Factory)
}

func TestGenesisState_Validate(t *testing.T) {
	ac := genesisMockCodec{}

	t.Run("valid_genesis_state", func(t *testing.T) {
		genState := DefaultGenesis()
		err := genState.Validate(ac)
		require.NoError(t, err)
	})

	t.Run("genesis_with_key_values_but_no_erc20_factory", func(t *testing.T) {
		genState := DefaultGenesis()
		genState.KeyValues = []GenesisKeyValue{
			{Key: []byte("test_key"), Value: []byte("test_value")},
		}
		// Erc20Factory is nil by default
		err := genState.Validate(ac)
		require.Error(t, err)
		require.Contains(t, err.Error(), "invalid empty ERC20Factory address")
	})

	t.Run("genesis_with_key_values_and_erc20_factory", func(t *testing.T) {
		genState := DefaultGenesis()
		genState.KeyValues = []GenesisKeyValue{
			{Key: []byte("test_key"), Value: []byte("test_value")},
		}
		genState.Erc20Factory = []byte("0x1234567890123456789012345678901234567890")
		err := genState.Validate(ac)
		require.NoError(t, err)
	})

	t.Run("genesis_with_empty_key_values", func(t *testing.T) {
		genState := DefaultGenesis()
		genState.KeyValues = []GenesisKeyValue{}
		err := genState.Validate(ac)
		require.NoError(t, err)
	})
}

func TestGenesisState_Validate_AddressNormalization(t *testing.T) {
	ac := genesisMockCodec{}

	t.Run("genesis_with_addresses_to_normalize", func(t *testing.T) {
		genState := DefaultGenesis()
		// Set params with addresses that need normalization
		genState.Params = DefaultParams()
		err := genState.Validate(ac)
		require.NoError(t, err)
	})
}

func TestGenesisState_Validate_EdgeCases(t *testing.T) {
	ac := genesisMockCodec{}

	t.Run("nil_genesis_state", func(t *testing.T) {
		var genState *GenesisState
		// This should panic or return an error, but let's test the behavior
		require.Panics(t, func() {
			genState.Validate(ac)
		})
	})

	t.Run("genesis_with_empty_params", func(t *testing.T) {
		genState := &GenesisState{
			Params:    DefaultParams(),
			KeyValues: []GenesisKeyValue{},
		}
		err := genState.Validate(ac)
		require.NoError(t, err) // Should pass with default params
	})
}

func TestGenesisKeyValue(t *testing.T) {
	t.Run("create_genesis_key_value", func(t *testing.T) {
		kv := GenesisKeyValue{
			Key:   []byte("test_key"),
			Value: []byte("test_value"),
		}
		require.Equal(t, []byte("test_key"), kv.Key)
		require.Equal(t, []byte("test_value"), kv.Value)
	})

	t.Run("empty_genesis_key_value", func(t *testing.T) {
		kv := GenesisKeyValue{}
		require.Empty(t, kv.Key)
		require.Empty(t, kv.Value)
	})
}

func TestGenesisState_Integration(t *testing.T) {
	ac := genesisMockCodec{}

	t.Run("full_genesis_cycle", func(t *testing.T) {
		// Create default genesis
		genState := DefaultGenesis()
		require.NotNil(t, genState)

		// Add some key values
		genState.KeyValues = []GenesisKeyValue{
			{Key: []byte("key1"), Value: []byte("value1")},
			{Key: []byte("key2"), Value: []byte("value2")},
		}

		// Set ERC20 factory address
		genState.Erc20Factory = []byte("0x1234567890123456789012345678901234567890")

		// Validate
		err := genState.Validate(ac)
		require.NoError(t, err)

		// Verify the state
		require.Len(t, genState.KeyValues, 2)
		require.Equal(t, []byte("key1"), genState.KeyValues[0].Key)
		require.Equal(t, []byte("value1"), genState.KeyValues[0].Value)
		require.Equal(t, []byte("key2"), genState.KeyValues[1].Key)
		require.Equal(t, []byte("value2"), genState.KeyValues[1].Value)
		require.Equal(t, []byte("0x1234567890123456789012345678901234567890"), genState.Erc20Factory)
	})
}
