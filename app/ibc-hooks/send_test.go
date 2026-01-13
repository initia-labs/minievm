package evm_hooks_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"

	"cosmossdk.io/collections"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	transfertypes "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"
	clienttypes "github.com/cosmos/ibc-go/v8/modules/core/02-client/types"

	nfttransfertypes "github.com/initia-labs/initia/x/ibc/nft-transfer/types"
	evmhooks "github.com/initia-labs/minievm/app/ibc-hooks"
)

func Test_SendPacket_asyncCallback_only(t *testing.T) {
	ctx, input := createDefaultTestInput(t)
	_, _, addr := keyPubAddr()
	_, _, addr2 := keyPubAddr()

	input.MockIBCMiddleware.setSequence(42)

	contractAddr := common.BytesToAddress(addr.Bytes()).Hex()
	data := transfertypes.FungibleTokenPacketData{
		Denom:    "foo",
		Amount:   "10000",
		Sender:   addr.String(),
		Receiver: addr2.String(),
		Memo:     fmt.Sprintf(`{"evm":{"async_callback":{"id":99,"contract_address":"%s"}},"key":"value"}`, contractAddr),
	}
	dataBz, err := json.Marshal(&data)
	require.NoError(t, err)

	seq, err := input.IBCHooksMiddleware.ICS4Middleware.SendPacket(ctx, nil, "transfer", "channel-0", clienttypes.ZeroHeight(), 0, dataBz)
	require.NoError(t, err)
	require.Equal(t, uint64(42), seq)

	var gotData transfertypes.FungibleTokenPacketData
	require.NoError(t, json.Unmarshal(input.MockIBCMiddleware.lastData, &gotData))
	require.Equal(t, `{"key":"value"}`, gotData.Memo)

	callbackBz, err := input.IBCHooksKeeper.GetAsyncCallback(ctx, "transfer", "channel-0", seq)
	require.NoError(t, err)
	expectedCallbackBz, err := json.Marshal(evmhooks.AsyncCallback{
		Id:              99,
		ContractAddress: contractAddr,
	})
	require.NoError(t, err)
	require.Equal(t, expectedCallbackBz, callbackBz)
}

func Test_SendPacket_asyncCallback_invalid_contract_addr(t *testing.T) {
	ctx, input := createDefaultTestInput(t)

	input.MockIBCMiddleware.setSequence(13)
	input.MockIBCMiddleware.lastData = []byte("initial")

	data := transfertypes.FungibleTokenPacketData{
		Denom:    "foo",
		Amount:   "10000",
		Sender:   "sender",
		Receiver: "receiver",
		Memo:     `{"evm":{"async_callback":{"id":1,"contract_address":"invalid-address"}}}`,
	}
	dataBz, err := json.Marshal(&data)
	require.NoError(t, err)

	seq, err := input.IBCHooksMiddleware.ICS4Middleware.SendPacket(ctx, nil, "transfer", "channel-invalid", clienttypes.ZeroHeight(), 0, dataBz)
	require.Error(t, err)
	require.Zero(t, seq)

	require.Equal(t, []byte("initial"), input.MockIBCMiddleware.lastData)

	_, err = input.IBCHooksKeeper.GetAsyncCallback(ctx, "transfer", "channel-invalid", seq)
	require.Error(t, err)
	require.True(t, errors.Is(err, collections.ErrNotFound))
}

func Test_SendPacket_asyncCallback_with_message(t *testing.T) {
	ctx, input := createDefaultTestInput(t)
	_, _, addr := keyPubAddr()
	_, _, addr2 := keyPubAddr()

	input.MockIBCMiddleware.setSequence(7)

	contractAddr := common.BytesToAddress(addr.Bytes()).Hex()
	memo := fmt.Sprintf(`{
		"evm": {
			"message": {
				"sender": "%s",
				"contract_addr": "%s",
				"input": "0x010203"
			},
			"async_callback": {
				"id": 99,
				"contract_address": "%s"
			}
		},
		"key": "value"
	}`, addr.String(), contractAddr, contractAddr)
	data := transfertypes.FungibleTokenPacketData{
		Denom:    "foo",
		Amount:   "10000",
		Sender:   addr.String(),
		Receiver: addr2.String(),
		Memo:     memo,
	}
	dataBz, err := json.Marshal(&data)
	require.NoError(t, err)

	seq, err := input.IBCHooksMiddleware.ICS4Middleware.SendPacket(ctx, nil, "transfer", "channel-1", clienttypes.ZeroHeight(), 0, dataBz)
	require.NoError(t, err)
	require.Equal(t, uint64(7), seq)

	var gotData transfertypes.FungibleTokenPacketData
	require.NoError(t, json.Unmarshal(input.MockIBCMiddleware.lastData, &gotData))

	var memoMap map[string]any
	require.NoError(t, json.Unmarshal([]byte(gotData.Memo), &memoMap))
	evmMap, ok := memoMap["evm"].(map[string]any)
	require.True(t, ok)
	_, hasAsync := evmMap["async_callback"]
	require.False(t, hasAsync)
	_, hasMessage := evmMap["message"]
	require.True(t, hasMessage)
	keyValue, ok := memoMap["key"].(string)
	require.True(t, ok)
	require.Equal(t, "value", keyValue)

	callbackBz, err := input.IBCHooksKeeper.GetAsyncCallback(ctx, "transfer", "channel-1", seq)
	require.NoError(t, err)
	expectedCallbackBz, err := json.Marshal(evmhooks.AsyncCallback{
		Id:              99,
		ContractAddress: contractAddr,
	})
	require.NoError(t, err)
	require.Equal(t, expectedCallbackBz, callbackBz)
}

func Test_SendPacket_asyncCallback_ics721(t *testing.T) {
	ctx, input := createDefaultTestInput(t)
	_, _, addr := keyPubAddr()

	input.MockIBCMiddleware.setSequence(11)

	contractAddr := common.BytesToAddress(addr.Bytes()).Hex()
	data := nfttransfertypes.NonFungibleTokenPacketData{
		ClassId:   "classId",
		ClassUri:  "classUri",
		ClassData: "classData",
		TokenIds:  []string{"tokenId"},
		TokenUris: []string{"tokenUri"},
		TokenData: []string{"tokenData"},
		Sender:    addr.String(),
		Receiver:  addr.String(),
		Memo:      fmt.Sprintf(`{"evm":{"async_callback":{"id":99,"contract_address":"%s"}}}`, contractAddr),
	}
	dataBz := data.GetBytes()

	seq, err := input.IBCHooksMiddleware.ICS4Middleware.SendPacket(ctx, nil, "nft-transfer", "channel-2", clienttypes.ZeroHeight(), 0, dataBz)
	require.NoError(t, err)
	require.Equal(t, uint64(11), seq)

	gotData, err := nfttransfertypes.DecodePacketData(input.MockIBCMiddleware.lastData)
	require.NoError(t, err)
	require.Equal(t, "{}", gotData.Memo)

	callbackBz, err := input.IBCHooksKeeper.GetAsyncCallback(ctx, "nft-transfer", "channel-2", seq)
	require.NoError(t, err)
	expectedCallbackBz, err := json.Marshal(evmhooks.AsyncCallback{
		Id:              99,
		ContractAddress: contractAddr,
	})
	require.NoError(t, err)
	require.Equal(t, expectedCallbackBz, callbackBz)
}

func Test_SendPacket_not_routed_passthrough(t *testing.T) {
	ctx, input := createDefaultTestInput(t)
	_, _, addr := keyPubAddr()
	_, _, addr2 := keyPubAddr()

	input.MockIBCMiddleware.setSequence(5)

	data := transfertypes.FungibleTokenPacketData{
		Denom:    "foo",
		Amount:   "10000",
		Sender:   addr.String(),
		Receiver: addr2.String(),
		Memo:     "not-json",
	}
	dataBz, err := json.Marshal(&data)
	require.NoError(t, err)

	seq, err := input.IBCHooksMiddleware.ICS4Middleware.SendPacket(ctx, nil, "transfer", "channel-9", clienttypes.ZeroHeight(), 0, dataBz)
	require.NoError(t, err)
	require.Equal(t, uint64(5), seq)

	var sent transfertypes.FungibleTokenPacketData
	require.NoError(t, json.Unmarshal(input.MockIBCMiddleware.lastData, &sent))
	require.Equal(t, data.Memo, sent.Memo)

	_, err = input.IBCHooksKeeper.GetAsyncCallback(ctx, "transfer", "channel-9", seq)
	require.Error(t, err)
	require.True(t, errors.Is(err, collections.ErrNotFound))
}
