package types

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestModuleConstants(t *testing.T) {
	// Test module name constants
	require.Equal(t, "evm", ModuleName)
	require.Equal(t, ModuleName, StoreKey)
	require.Equal(t, "transient_"+ModuleName, TStoreKey)
	require.Equal(t, ModuleName, QuerierRoute)
	require.Equal(t, ModuleName, RouterKey)
}

func TestStorePrefixes(t *testing.T) {
	// Test VM store prefix
	require.Equal(t, []byte{0x21}, VMStorePrefix)

	// Test ERC20 prefixes
	require.Equal(t, []byte{0x31}, ERC20sPrefix)
	require.Equal(t, []byte{0x32}, ERC20StoresPrefix)
	require.Equal(t, []byte{0x33}, ERC20DenomsByContractAddrPrefix)
	require.Equal(t, []byte{0x34}, ERC20ContractAddrsByDenomPrefix)

	// Test ERC721 prefixes
	require.Equal(t, []byte{0x41}, ERC721ClassURIPrefix)
	require.Equal(t, []byte{0x42}, ERC721ClassIdsByContractAddrPrefix)
	require.Equal(t, []byte{0x43}, ERC721ContractAddrsByClassIdPrefix)

	// Test EVM block hash prefix
	require.Equal(t, []byte{0x71}, EVMBlockHashPrefix)

	// Test parameter and address keys
	require.Equal(t, []byte{0x51}, ParamsKey)
	require.Equal(t, []byte{0x61}, ERC20FactoryAddrKey)
	require.Equal(t, []byte{0x62}, ERC20WrapperAddrKey)
	require.Equal(t, []byte{0x63}, ConnectOracleAddrKey)
}

func TestContextKeys(t *testing.T) {
	// Test that context keys are properly defined
	require.Equal(t, ContextKey(0), CONTEXT_KEY_EXECUTE_REQUESTS)
	require.Equal(t, ContextKey(1), CONTEXT_KEY_PARENT_EXECUTE_REQUESTS)
	require.Equal(t, ContextKey(2), CONTEXT_KEY_RECURSIVE_DEPTH)
	require.Equal(t, ContextKey(3), CONTEXT_KEY_LOAD_DECIMALS)
	require.Equal(t, ContextKey(4), CONTEXT_KEY_TRACING)
	require.Equal(t, ContextKey(5), CONTEXT_KEY_ETH_TX)
	require.Equal(t, ContextKey(6), CONTEXT_KEY_ETH_TX_SENDER)
	require.Equal(t, ContextKey(7), CONTEXT_KEY_GAS_PRICES)
	require.Equal(t, ContextKey(8), CONTEXT_KEY_SEQUENCE_INCREMENTED)
	require.Equal(t, ContextKey(9), CONTEXT_KEY_TRACE_EVM)
}

func TestStorePrefixesUniqueness(t *testing.T) {
	// Test that all store prefixes are unique
	prefixes := [][]byte{
		VMStorePrefix,
		ERC20sPrefix,
		ERC20StoresPrefix,
		ERC20DenomsByContractAddrPrefix,
		ERC20ContractAddrsByDenomPrefix,
		ERC721ClassURIPrefix,
		ERC721ClassIdsByContractAddrPrefix,
		ERC721ContractAddrsByClassIdPrefix,
		EVMBlockHashPrefix,
		ParamsKey,
		ERC20FactoryAddrKey,
		ERC20WrapperAddrKey,
		ConnectOracleAddrKey,
	}

	seen := make(map[string]bool)
	for _, prefix := range prefixes {
		prefixStr := string(prefix)
		require.False(t, seen[prefixStr], "Duplicate prefix: %v", prefix)
		seen[prefixStr] = true
	}
}

func TestContextKeysUniqueness(t *testing.T) {
	// Test that all context keys are unique
	contextKeys := []ContextKey{
		CONTEXT_KEY_EXECUTE_REQUESTS,
		CONTEXT_KEY_PARENT_EXECUTE_REQUESTS,
		CONTEXT_KEY_RECURSIVE_DEPTH,
		CONTEXT_KEY_LOAD_DECIMALS,
		CONTEXT_KEY_TRACING,
		CONTEXT_KEY_ETH_TX,
		CONTEXT_KEY_ETH_TX_SENDER,
		CONTEXT_KEY_GAS_PRICES,
		CONTEXT_KEY_SEQUENCE_INCREMENTED,
		CONTEXT_KEY_TRACE_EVM,
	}

	seen := make(map[ContextKey]bool)
	for _, key := range contextKeys {
		require.False(t, seen[key], "Duplicate context key: %v", key)
		seen[key] = true
	}
}

func TestKeyStructure(t *testing.T) {
	// Test that keys have the expected structure
	t.Run("vm_store_prefix", func(t *testing.T) {
		require.Len(t, VMStorePrefix, 1)
		require.Equal(t, byte(0x21), VMStorePrefix[0])
	})

	t.Run("erc20_prefixes", func(t *testing.T) {
		erc20Prefixes := [][]byte{
			ERC20sPrefix,
			ERC20StoresPrefix,
			ERC20DenomsByContractAddrPrefix,
			ERC20ContractAddrsByDenomPrefix,
		}
		for _, prefix := range erc20Prefixes {
			require.Len(t, prefix, 1)
			require.GreaterOrEqual(t, prefix[0], byte(0x31))
			require.LessOrEqual(t, prefix[0], byte(0x34))
		}
	})

	t.Run("erc721_prefixes", func(t *testing.T) {
		erc721Prefixes := [][]byte{
			ERC721ClassURIPrefix,
			ERC721ClassIdsByContractAddrPrefix,
			ERC721ContractAddrsByClassIdPrefix,
		}
		for _, prefix := range erc721Prefixes {
			require.Len(t, prefix, 1)
			require.GreaterOrEqual(t, prefix[0], byte(0x41))
			require.LessOrEqual(t, prefix[0], byte(0x43))
		}
	})

	t.Run("parameter_keys", func(t *testing.T) {
		paramKeys := [][]byte{
			ParamsKey,
			ERC20FactoryAddrKey,
			ERC20WrapperAddrKey,
			ConnectOracleAddrKey,
		}
		for _, key := range paramKeys {
			require.Len(t, key, 1)
			require.GreaterOrEqual(t, key[0], byte(0x51))
			require.LessOrEqual(t, key[0], byte(0x63))
		}
	})
}

func TestContextKeyType(t *testing.T) {
	// Test that ContextKey is properly typed
	var key ContextKey
	require.IsType(t, ContextKey(0), key)

	// Test that context keys can be used as map keys
	contextMap := make(map[ContextKey]string)
	contextMap[CONTEXT_KEY_EXECUTE_REQUESTS] = "execute_requests"
	contextMap[CONTEXT_KEY_TRACING] = "tracing"

	require.Equal(t, "execute_requests", contextMap[CONTEXT_KEY_EXECUTE_REQUESTS])
	require.Equal(t, "tracing", contextMap[CONTEXT_KEY_TRACING])
}
