package types

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/require"
)

func TestNewLog(t *testing.T) {
	t.Run("create_log_from_eth_log", func(t *testing.T) {
		// Create a test Ethereum log
		address := common.HexToAddress("0x1234567890123456789012345678901234567890")
		topics := []common.Hash{
			common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000001"),
			common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000002"),
		}
		data := []byte{0x01, 0x02, 0x03, 0x04}

		ethLog := &types.Log{
			Address: address,
			Topics:  topics,
			Data:    data,
		}

		// Convert to internal log
		log := NewLog(ethLog)

		// Verify the conversion
		require.Equal(t, address.Hex(), log.Address)
		require.Len(t, log.Topics, 2)
		require.Equal(t, topics[0].Hex(), log.Topics[0])
		require.Equal(t, topics[1].Hex(), log.Topics[1])
		require.Equal(t, "0x01020304", log.Data)
	})

	t.Run("create_log_with_empty_topics", func(t *testing.T) {
		address := common.HexToAddress("0x1234567890123456789012345678901234567890")
		ethLog := &types.Log{
			Address: address,
			Topics:  []common.Hash{},
			Data:    []byte{},
		}

		log := NewLog(ethLog)

		require.Equal(t, address.Hex(), log.Address)
		require.Empty(t, log.Topics)
		require.Equal(t, "0x", log.Data)
	})

	t.Run("create_log_with_nil_data", func(t *testing.T) {
		address := common.HexToAddress("0x1234567890123456789012345678901234567890")
		ethLog := &types.Log{
			Address: address,
			Topics:  []common.Hash{},
			Data:    nil,
		}

		log := NewLog(ethLog)

		require.Equal(t, address.Hex(), log.Address)
		require.Empty(t, log.Topics)
		require.Equal(t, "0x", log.Data)
	})
}

func TestNewLogs(t *testing.T) {
	t.Run("create_logs_from_eth_logs", func(t *testing.T) {
		// Create test Ethereum logs
		ethLogs := []*types.Log{
			{
				Address: common.HexToAddress("0x1234567890123456789012345678901234567890"),
				Topics:  []common.Hash{common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000001")},
				Data:    []byte{0x01, 0x02},
			},
			{
				Address: common.HexToAddress("0xabcdefabcdefabcdefabcdefabcdefabcdefabcd"),
				Topics:  []common.Hash{common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000002")},
				Data:    []byte{0x03, 0x04},
			},
		}

		// Convert to internal logs
		logs := NewLogs(ethLogs)

		// Verify the conversion
		require.Len(t, logs, 2)
		require.Equal(t, ethLogs[0].Address.Hex(), logs[0].Address)
		require.Equal(t, ethLogs[1].Address.Hex(), logs[1].Address)
		require.Equal(t, "0x0102", logs[0].Data)
		require.Equal(t, "0x0304", logs[1].Data)
	})

	t.Run("create_logs_from_empty_slice", func(t *testing.T) {
		ethLogs := []*types.Log{}
		logs := NewLogs(ethLogs)
		require.Empty(t, logs)
	})

	t.Run("create_logs_from_nil_slice", func(t *testing.T) {
		var ethLogs []*types.Log
		logs := NewLogs(ethLogs)
		require.Empty(t, logs)
	})
}

func TestLog_ToEthLog(t *testing.T) {
	t.Run("convert_log_to_eth_log", func(t *testing.T) {
		// Create an internal log
		log := Log{
			Address: "0x1234567890123456789012345678901234567890",
			Topics: []string{
				"0x0000000000000000000000000000000000000000000000000000000000000001",
				"0x0000000000000000000000000000000000000000000000000000000000000002",
			},
			Data: "0x01020304",
		}

		// Convert to Ethereum log
		ethLog := log.ToEthLog()

		// Verify the conversion
		require.Equal(t, common.HexToAddress("0x1234567890123456789012345678901234567890"), ethLog.Address)
		require.Len(t, ethLog.Topics, 2)
		require.Equal(t, common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000001"), ethLog.Topics[0])
		require.Equal(t, common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000002"), ethLog.Topics[1])
		require.Equal(t, []byte{0x01, 0x02, 0x03, 0x04}, ethLog.Data)
	})

	t.Run("convert_log_with_empty_topics", func(t *testing.T) {
		log := Log{
			Address: "0x1234567890123456789012345678901234567890",
			Topics:  []string{},
			Data:    "0x",
		}

		ethLog := log.ToEthLog()

		require.Equal(t, common.HexToAddress("0x1234567890123456789012345678901234567890"), ethLog.Address)
		require.Empty(t, ethLog.Topics)
		require.Empty(t, ethLog.Data)
	})

	t.Run("convert_log_with_empty_data", func(t *testing.T) {
		log := Log{
			Address: "0x1234567890123456789012345678901234567890",
			Topics:  []string{},
			Data:    "0x",
		}

		ethLog := log.ToEthLog()

		require.Equal(t, common.HexToAddress("0x1234567890123456789012345678901234567890"), ethLog.Address)
		require.Empty(t, ethLog.Topics)
		require.Empty(t, ethLog.Data)
	})
}

func TestLogs_ToEthLogs(t *testing.T) {
	t.Run("convert_logs_to_eth_logs", func(t *testing.T) {
		// Create internal logs
		logs := Logs{
			{
				Address: "0x1234567890123456789012345678901234567890",
				Topics:  []string{"0x0000000000000000000000000000000000000000000000000000000000000001"},
				Data:    "0x0102",
			},
			{
				Address: "0xabcdefabcdefabcdefabcdefabcdefabcdefabcd",
				Topics:  []string{"0x0000000000000000000000000000000000000000000000000000000000000002"},
				Data:    "0x0304",
			},
		}

		// Convert to Ethereum logs
		ethLogs := logs.ToEthLogs()

		// Verify the conversion
		require.Len(t, ethLogs, 2)
		require.Equal(t, common.HexToAddress("0x1234567890123456789012345678901234567890"), ethLogs[0].Address)
		require.Equal(t, common.HexToAddress("0xabcdefabcdefabcdefabcdefabcdefabcdefabcd"), ethLogs[1].Address)
		require.Equal(t, []byte{0x01, 0x02}, ethLogs[0].Data)
		require.Equal(t, []byte{0x03, 0x04}, ethLogs[1].Data)
	})

	t.Run("convert_empty_logs", func(t *testing.T) {
		logs := Logs{}
		ethLogs := logs.ToEthLogs()
		require.Empty(t, ethLogs)
	})
}

func TestLogConversionRoundTrip(t *testing.T) {
	t.Run("eth_log_to_internal_to_eth_round_trip", func(t *testing.T) {
		// Create original Ethereum log
		originalEthLog := &types.Log{
			Address: common.HexToAddress("0x1234567890123456789012345678901234567890"),
			Topics: []common.Hash{
				common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000001"),
				common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000002"),
			},
			Data: []byte{0x01, 0x02, 0x03, 0x04},
		}

		// Convert to internal log
		internalLog := NewLog(originalEthLog)

		// Convert back to Ethereum log
		convertedEthLog := internalLog.ToEthLog()

		// Verify round trip
		require.Equal(t, originalEthLog.Address, convertedEthLog.Address)
		require.Equal(t, originalEthLog.Topics, convertedEthLog.Topics)
		require.Equal(t, originalEthLog.Data, convertedEthLog.Data)
	})

	t.Run("internal_log_to_eth_to_internal_round_trip", func(t *testing.T) {
		// Create original internal log
		originalInternalLog := Log{
			Address: "0x1234567890123456789012345678901234567890",
			Topics: []string{
				"0x0000000000000000000000000000000000000000000000000000000000000001",
				"0x0000000000000000000000000000000000000000000000000000000000000002",
			},
			Data: "0x01020304",
		}

		// Convert to Ethereum log
		ethLog := originalInternalLog.ToEthLog()

		// Convert back to internal log
		convertedInternalLog := NewLog(ethLog)

		// Verify round trip
		require.Equal(t, originalInternalLog.Address, convertedInternalLog.Address)
		require.Equal(t, originalInternalLog.Topics, convertedInternalLog.Topics)
		require.Equal(t, originalInternalLog.Data, convertedInternalLog.Data)
	})
}

func TestLogsConversionRoundTrip(t *testing.T) {
	t.Run("eth_logs_to_internal_to_eth_round_trip", func(t *testing.T) {
		// Create original Ethereum logs
		originalEthLogs := []*types.Log{
			{
				Address: common.HexToAddress("0x1234567890123456789012345678901234567890"),
				Topics:  []common.Hash{common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000001")},
				Data:    []byte{0x01, 0x02},
			},
			{
				Address: common.HexToAddress("0xabcdefabcdefabcdefabcdefabcdefabcdefabcd"),
				Topics:  []common.Hash{common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000002")},
				Data:    []byte{0x03, 0x04},
			},
		}

		// Convert to internal logs
		internalLogs := NewLogs(originalEthLogs)

		// Convert back to Ethereum logs
		convertedEthLogs := internalLogs.ToEthLogs()

		// Verify round trip
		require.Len(t, convertedEthLogs, len(originalEthLogs))
		for i := range originalEthLogs {
			require.Equal(t, originalEthLogs[i].Address, convertedEthLogs[i].Address)
			require.Equal(t, originalEthLogs[i].Topics, convertedEthLogs[i].Topics)
			require.Equal(t, originalEthLogs[i].Data, convertedEthLogs[i].Data)
		}
	})
}
