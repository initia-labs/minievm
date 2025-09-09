package types

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/stretchr/testify/require"
)

func TestDefaultQueryCosmosWhitelist(t *testing.T) {
	whitelist := DefaultQueryCosmosWhitelist()

	// Test that whitelist is not empty
	require.NotEmpty(t, whitelist)

	// Test specific expected entries
	expectedPaths := []string{
		"/connect.oracle.v2.Query/GetPrices",
		"/connect.oracle.v2.Query/GetPrice",
		"/connect.oracle.v2.Query/GetAllCurrencyPairs",
	}

	for _, path := range expectedPaths {
		protoSet, exists := whitelist[path]
		require.True(t, exists, "Expected path %s to exist in whitelist", path)
		require.NotNil(t, protoSet.Request, "Request should not be nil for path %s", path)
		require.NotNil(t, protoSet.Response, "Response should not be nil for path %s", path)
	}

	// Test that all entries have valid proto messages
	for path, protoSet := range whitelist {
		require.NotNil(t, protoSet.Request, "Request should not be nil for path %s", path)
		require.NotNil(t, protoSet.Response, "Response should not be nil for path %s", path)
	}
}

func TestProtoSet(t *testing.T) {
	t.Run("create_proto_set", func(t *testing.T) {
		// Create a mock proto message for testing
		mockRequest := &MsgCall{}
		mockResponse := &MsgCallResponse{}

		protoSet := ProtoSet{
			Request:  mockRequest,
			Response: mockResponse,
		}

		require.NotNil(t, protoSet.Request)
		require.NotNil(t, protoSet.Response)
	})

	t.Run("empty_proto_set", func(t *testing.T) {
		protoSet := ProtoSet{}

		require.Nil(t, protoSet.Request)
		require.Nil(t, protoSet.Response)
	})
}

func TestQueryCosmosWhitelist(t *testing.T) {
	t.Run("create_whitelist", func(t *testing.T) {
		whitelist := make(QueryCosmosWhitelist)
		require.Empty(t, whitelist)

		// Add an entry
		whitelist["/test/path"] = ProtoSet{
			Request:  &MsgCall{},
			Response: &MsgCallResponse{},
		}

		require.Len(t, whitelist, 1)
		require.Contains(t, whitelist, "/test/path")
	})

	t.Run("whitelist_operations", func(t *testing.T) {
		whitelist := make(QueryCosmosWhitelist)

		// Test adding entries
		whitelist["/path1"] = ProtoSet{Request: &MsgCall{}, Response: &MsgCallResponse{}}
		whitelist["/path2"] = ProtoSet{Request: &MsgCreate{}, Response: &MsgCreateResponse{}}

		require.Len(t, whitelist, 2)

		// Test checking existence
		require.True(t, whitelist["/path1"].Request != nil)
		require.True(t, whitelist["/path2"].Request != nil)
		require.Nil(t, whitelist["/nonexistent"].Request)
	})
}

func TestConvertProtoToJSON(t *testing.T) {
	// Create a codec for testing
	interfaceRegistry := cdctypes.NewInterfaceRegistry()
	RegisterInterfaces(interfaceRegistry)
	cdc := codec.NewProtoCodec(interfaceRegistry)

	t.Run("convert_valid_proto_to_json", func(t *testing.T) {
		// Create a test proto message
		request := &MsgCall{
			Sender:       "0x1234567890123456789012345678901234567890",
			ContractAddr: "0xabcdefabcdefabcdefabcdefabcdefabcdefabcd",
			Input:        "0x12345678",
		}

		// Marshal the proto message to bytes
		protoBytes, err := cdc.Marshal(request)
		require.NoError(t, err)

		// Convert to JSON
		jsonBytes, err := ConvertProtoToJSON(cdc, &MsgCall{}, protoBytes)
		require.NoError(t, err)
		require.NotEmpty(t, jsonBytes)

		// Verify it's valid JSON
		require.Contains(t, string(jsonBytes), "{")
		require.Contains(t, string(jsonBytes), "}")
	})

	t.Run("convert_invalid_proto_to_json", func(t *testing.T) {
		// Try to convert invalid bytes
		invalidBytes := []byte("invalid proto data")

		_, err := ConvertProtoToJSON(cdc, &MsgCall{}, invalidBytes)
		require.Error(t, err)
	})

	t.Run("convert_with_nil_codec", func(t *testing.T) {
		request := &MsgCall{}
		protoBytes, err := cdc.Marshal(request)
		require.NoError(t, err)

		// This should panic or return an error
		require.Panics(t, func() {
			ConvertProtoToJSON(nil, &MsgCall{}, protoBytes)
		})
	})
}

func TestConvertJSONToProto(t *testing.T) {
	// Create a codec for testing
	interfaceRegistry := cdctypes.NewInterfaceRegistry()
	RegisterInterfaces(interfaceRegistry)
	cdc := codec.NewProtoCodec(interfaceRegistry)

	t.Run("convert_valid_json_to_proto", func(t *testing.T) {
		// Create valid JSON for a proto message
		jsonData := []byte(`{"sender":"0x1234567890123456789012345678901234567890","contract_addr":"0xabcdefabcdefabcdefabcdefabcdefabcdefabcd","input":"0x12345678"}`)

		// Convert to proto
		protoBytes, err := ConvertJSONToProto(cdc, &MsgCall{}, jsonData)
		require.NoError(t, err)
		require.NotEmpty(t, protoBytes)
	})

	t.Run("convert_invalid_json_to_proto", func(t *testing.T) {
		// Try to convert invalid JSON
		invalidJSON := []byte(`{invalid json}`)

		_, err := ConvertJSONToProto(cdc, &MsgCall{}, invalidJSON)
		require.Error(t, err)
	})

	t.Run("convert_with_nil_codec", func(t *testing.T) {
		jsonData := []byte(`{}`)

		// This should panic or return an error
		require.Panics(t, func() {
			ConvertJSONToProto(nil, &MsgCall{}, jsonData)
		})
	})
}

func TestConversionRoundTrip(t *testing.T) {
	// Create a codec for testing
	interfaceRegistry := cdctypes.NewInterfaceRegistry()
	RegisterInterfaces(interfaceRegistry)
	cdc := codec.NewProtoCodec(interfaceRegistry)

	t.Run("proto_to_json_to_proto_round_trip", func(t *testing.T) {
		// Start with a proto message
		originalRequest := &MsgCall{
			Sender:       "0x1234567890123456789012345678901234567890",
			ContractAddr: "0xabcdefabcdefabcdefabcdefabcdefabcdefabcd",
			Input:        "0x12345678",
		}

		// Marshal to proto bytes
		protoBytes, err := cdc.Marshal(originalRequest)
		require.NoError(t, err)

		// Convert to JSON
		jsonBytes, err := ConvertProtoToJSON(cdc, &MsgCall{}, protoBytes)
		require.NoError(t, err)

		// Convert back to proto
		convertedProtoBytes, err := ConvertJSONToProto(cdc, &MsgCall{}, jsonBytes)
		require.NoError(t, err)

		// Verify the round trip
		require.Equal(t, protoBytes, convertedProtoBytes)
	})
}

func TestWhitelistIntegration(t *testing.T) {
	t.Run("whitelist_with_default_entries", func(t *testing.T) {
		whitelist := DefaultQueryCosmosWhitelist()

		// Test that all entries can be used for conversion
		for path, protoSet := range whitelist {
			require.NotNil(t, protoSet.Request, "Request should not be nil for path %s", path)
			require.NotNil(t, protoSet.Response, "Response should not be nil for path %s", path)

			// Test that the proto messages can be marshaled
			interfaceRegistry := cdctypes.NewInterfaceRegistry()
			RegisterInterfaces(interfaceRegistry)
			cdc := codec.NewProtoCodec(interfaceRegistry)

			// Test request marshaling
			_, err := cdc.Marshal(protoSet.Request)
			require.NoError(t, err, "Failed to marshal request for path %s", path)

			// Test response marshaling
			_, err = cdc.Marshal(protoSet.Response)
			require.NoError(t, err, "Failed to marshal response for path %s", path)
		}
	})
}
