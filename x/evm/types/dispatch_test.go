package types

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/gogoproto/proto"
	"github.com/stretchr/testify/require"
)

func TestExtractLogsFromResponse(t *testing.T) {
	t.Run("extract_logs_from_msg_call_response", func(t *testing.T) {
		logs := Logs{{Address: "0x1", Data: "0x6c6f6764617461"}} // "logdata" in hex
		respCall := &MsgCallResponse{Logs: logs}
		data, err := proto.Marshal(respCall)
		require.NoError(t, err)

		gotLogs, err := ExtractLogsFromResponse(data, sdk.MsgTypeURL(&MsgCall{}))
		require.NoError(t, err)
		require.Equal(t, logs, gotLogs)
	})

	t.Run("extract_logs_from_msg_create_response", func(t *testing.T) {
		logs := Logs{{Address: "0x2", Data: "0x637265617465"}} // "create" in hex
		respCreate := &MsgCreateResponse{Logs: logs}
		data, err := proto.Marshal(respCreate)
		require.NoError(t, err)

		gotLogs, err := ExtractLogsFromResponse(data, sdk.MsgTypeURL(&MsgCreate{}))
		require.NoError(t, err)
		require.Equal(t, logs, gotLogs)
	})

	t.Run("extract_logs_from_msg_create2_response", func(t *testing.T) {
		logs := Logs{{Address: "0x3", Data: "0x63726561746532"}} // "create2" in hex
		respCreate2 := &MsgCreate2Response{Logs: logs}
		data, err := proto.Marshal(respCreate2)
		require.NoError(t, err)

		gotLogs, err := ExtractLogsFromResponse(data, sdk.MsgTypeURL(&MsgCreate2{}))
		require.NoError(t, err)
		require.Equal(t, logs, gotLogs)
	})

	t.Run("extract_logs_from_empty_response", func(t *testing.T) {
		// Test with empty logs
		emptyLogs := Logs{}
		respCall := &MsgCallResponse{Logs: emptyLogs}
		data, err := proto.Marshal(respCall)
		require.NoError(t, err)

		gotLogs, err := ExtractLogsFromResponse(data, sdk.MsgTypeURL(&MsgCall{}))
		require.NoError(t, err)
		require.Nil(t, gotLogs)
	})

	t.Run("extract_logs_from_multiple_logs", func(t *testing.T) {
		// Test with multiple logs
		logs := Logs{
			{Address: "0x1", Data: "0x6c6f6731"},
			{Address: "0x2", Data: "0x6c6f6732"},
			{Address: "0x3", Data: "0x6c6f6733"},
		}
		respCall := &MsgCallResponse{Logs: logs}
		data, err := proto.Marshal(respCall)
		require.NoError(t, err)

		gotLogs, err := ExtractLogsFromResponse(data, sdk.MsgTypeURL(&MsgCall{}))
		require.NoError(t, err)
		require.Equal(t, logs, gotLogs)
		require.Len(t, gotLogs, 3)
	})

	t.Run("extract_logs_from_unknown_type", func(t *testing.T) {
		// Test with unknown message type
		logs := Logs{{Address: "0x1", Data: "0x6c6f6764617461"}}
		respCall := &MsgCallResponse{Logs: logs}
		data, err := proto.Marshal(respCall)
		require.NoError(t, err)

		gotLogs, err := ExtractLogsFromResponse(data, "unknown/type")
		require.NoError(t, err)
		require.Nil(t, gotLogs)
	})

	t.Run("extract_logs_from_invalid_data", func(t *testing.T) {
		// Test with invalid protobuf data
		gotLogs, err := ExtractLogsFromResponse([]byte("invalid protobuf data"), sdk.MsgTypeURL(&MsgCall{}))
		require.Error(t, err)
		require.Nil(t, gotLogs)
	})

	t.Run("extract_logs_from_nil_data", func(t *testing.T) {
		// Test with nil data
		gotLogs, err := ExtractLogsFromResponse(nil, sdk.MsgTypeURL(&MsgCall{}))
		require.NoError(t, err)
		require.Nil(t, gotLogs)
	})

	t.Run("extract_logs_from_empty_data", func(t *testing.T) {
		// Test with empty data
		gotLogs, err := ExtractLogsFromResponse([]byte{}, sdk.MsgTypeURL(&MsgCall{}))
		require.NoError(t, err)
		require.Nil(t, gotLogs)
	})

	t.Run("extract_logs_with_wrong_type_for_data", func(t *testing.T) {
		// Test with data that doesn't match the message type
		logs := Logs{{Address: "0x1", Data: "0x6c6f6764617461"}}
		respCall := &MsgCallResponse{Logs: logs}
		data, err := proto.Marshal(respCall)
		require.NoError(t, err)

		// Use MsgCreate type URL but with MsgCall data
		gotLogs, err := ExtractLogsFromResponse(data, sdk.MsgTypeURL(&MsgCreate{}))
		require.NoError(t, err)
		require.Nil(t, gotLogs)
	})
}
