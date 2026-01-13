package evm_hooks_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"

	"cosmossdk.io/collections"
	sdk "github.com/cosmos/cosmos-sdk/types"
	transfertypes "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"
	channeltypes "github.com/cosmos/ibc-go/v8/modules/core/04-channel/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/holiman/uint256"
	"github.com/stretchr/testify/require"

	ibchookstypes "github.com/initia-labs/initia/x/ibc-hooks/types"
	nfttransfertypes "github.com/initia-labs/initia/x/ibc/nft-transfer/types"
	evmhooks "github.com/initia-labs/minievm/app/ibc-hooks"
	"github.com/initia-labs/minievm/x/evm/contracts/counter"
)

func Test_onTimeoutIcs20Packet_noMemo(t *testing.T) {
	ctx, input := createDefaultTestInput(t)
	_, _, addr := keyPubAddr()
	_, _, addr2 := keyPubAddr()

	data := transfertypes.FungibleTokenPacketData{
		Denom:    "foo",
		Amount:   "10000",
		Sender:   addr.String(),
		Receiver: addr2.String(),
		Memo:     "",
	}

	dataBz, err := json.Marshal(&data)
	require.NoError(t, err)

	err = input.IBCHooksMiddleware.OnTimeoutPacket(ctx, channeltypes.Packet{
		Data: dataBz,
	}, addr)
	require.NoError(t, err)
}

func Test_onTimeoutPacket_acl_not_allowed(t *testing.T) {
	ctx, input := createDefaultTestInput(t)
	_, _, addr := keyPubAddr()

	sourcePort := "transfer"
	sourceChannel := "channel-timeout-acl"
	sequence := uint64(1)
	contractAddr := common.BytesToAddress(addr.Bytes()).Hex()

	callbackBz, err := json.Marshal(evmhooks.AsyncCallback{
		Id:              sequence,
		ContractAddress: contractAddr,
	})
	require.NoError(t, err)
	require.NoError(t, input.IBCHooksKeeper.SetAsyncCallback(ctx, sourcePort, sourceChannel, sequence, callbackBz))

	ctx = ctx.WithEventManager(sdk.NewEventManager())

	data := transfertypes.FungibleTokenPacketData{
		Denom:    "foo",
		Amount:   "1",
		Sender:   addr.String(),
		Receiver: addr.String(),
		Memo:     "",
	}

	dataBz, err := json.Marshal(&data)
	require.NoError(t, err)

	err = input.IBCHooksMiddleware.OnTimeoutPacket(ctx, channeltypes.Packet{
		Data:          dataBz,
		SourcePort:    sourcePort,
		SourceChannel: sourceChannel,
		Sequence:      sequence,
	}, addr)
	require.NoError(t, err)

	_, err = input.IBCHooksKeeper.GetAsyncCallback(ctx, sourcePort, sourceChannel, sequence)
	require.Error(t, err)
	require.True(t, errors.Is(err, collections.ErrNotFound))

	events := ctx.EventManager().Events()
	require.Len(t, events, 1)
	event := events[0]
	require.Equal(t, ibchookstypes.EventTypeHookFailed, event.Type)
	require.Len(t, event.Attributes, 2)
	require.Equal(t, ibchookstypes.AttributeKeyReason, string(event.Attributes[0].Key))
	require.Equal(t, "failed to check ACL", string(event.Attributes[0].Value))
	require.Equal(t, ibchookstypes.AttributeKeyError, string(event.Attributes[1].Key))
	require.Equal(t, "not allowed", string(event.Attributes[1].Value))
}

func Test_onTimeoutIcs20Packet_memo(t *testing.T) {
	ctx, input := createDefaultTestInput(t)
	_, _, addr := keyPubAddr()
	evmAddr := common.BytesToAddress(addr.Bytes())
	sourcePort := "transfer"
	sourceChannel := "channel-0"
	sequence := uint64(99)

	codeBz, err := hexutil.Decode(counter.CounterBin)
	require.NoError(t, err)

	_, contractAddr, _, err := input.EVMKeeper.EVMCreate(ctx, evmAddr, codeBz, nil, nil)
	require.NoError(t, err)

	abi, err := counter.CounterMetaData.GetAbi()
	require.NoError(t, err)

	data := transfertypes.FungibleTokenPacketData{
		Denom:    "foo",
		Amount:   "10000",
		Sender:   addr.String(),
		Receiver: contractAddr.Hex(),
		Memo: fmt.Sprintf(`{
			"evm": {
				"async_callback": {
					"id": 99,
					"contract_address": "%s"
				}
			}
		}`, contractAddr.Hex()),
	}

	dataBz, err := json.Marshal(&data)
	require.NoError(t, err)
	callbackBz, err := json.Marshal(evmhooks.AsyncCallback{
		Id:              99,
		ContractAddress: contractAddr.Hex(),
	})
	require.NoError(t, err)
	require.NoError(t, input.IBCHooksKeeper.SetAsyncCallback(ctx, sourcePort, sourceChannel, sequence, callbackBz))

	// hook should not be called to due to acl
	err = input.IBCHooksMiddleware.OnTimeoutPacket(ctx, channeltypes.Packet{
		Data:          dataBz,
		SourcePort:    sourcePort,
		SourceChannel: sourceChannel,
		Sequence:      sequence,
	}, addr)
	require.NoError(t, err)

	// check the contract state
	queryInputBz, err := abi.Pack("count")
	require.NoError(t, err)
	queryRes, logs, err := input.EVMKeeper.EVMCall(ctx, evmAddr, contractAddr, queryInputBz, nil, nil, nil)
	require.NoError(t, err)
	require.Equal(t, uint256.NewInt(0).Bytes32(), [32]byte(queryRes))
	require.Empty(t, logs)

	// set acl
	require.NoError(t, input.IBCHooksKeeper.SetAllowed(ctx, contractAddr[:], true))
	require.NoError(t, input.IBCHooksKeeper.SetAsyncCallback(ctx, sourcePort, sourceChannel, sequence, callbackBz))

	// success
	err = input.IBCHooksMiddleware.OnTimeoutPacket(ctx, channeltypes.Packet{
		Data:          dataBz,
		SourcePort:    sourcePort,
		SourceChannel: sourceChannel,
		Sequence:      sequence,
	}, addr)
	require.NoError(t, err)

	// check the contract state; increased by 99
	queryRes, logs, err = input.EVMKeeper.EVMCall(ctx, evmAddr, contractAddr, queryInputBz, nil, nil, nil)
	require.NoError(t, err)
	require.Equal(t, uint256.NewInt(99).Bytes32(), [32]byte(queryRes))
	require.Empty(t, logs)
}

func Test_OnTimeoutPacket_ICS721(t *testing.T) {
	ctx, input := createDefaultTestInput(t)
	_, _, addr := keyPubAddr()
	_, _, addr2 := keyPubAddr()

	data := nfttransfertypes.NonFungibleTokenPacketData{
		ClassId:   "classId",
		ClassUri:  "classUri",
		ClassData: "classData",
		TokenIds:  []string{"tokenId"},
		TokenUris: []string{"tokenUri"},
		TokenData: []string{"tokenData"},
		Sender:    addr.String(),
		Receiver:  addr2.String(),
		Memo:      "",
	}

	dataBz, err := json.Marshal(&data)
	require.NoError(t, err)

	err = input.IBCHooksMiddleware.OnTimeoutPacket(ctx, channeltypes.Packet{
		Data: dataBz,
	}, addr)
	require.NoError(t, err)
}

func Test_onTimeoutPacket_memo_ICS721(t *testing.T) {
	ctx, input := createDefaultTestInput(t)
	_, _, addr := keyPubAddr()
	evmAddr := common.BytesToAddress(addr.Bytes())
	sourcePort := "nft-transfer"
	sourceChannel := "channel-1"
	sequence := uint64(99)

	codeBz, err := hexutil.Decode(counter.CounterBin)
	require.NoError(t, err)

	_, contractAddr, _, err := input.EVMKeeper.EVMCreate(ctx, evmAddr, codeBz, nil, nil)
	require.NoError(t, err)

	abi, err := counter.CounterMetaData.GetAbi()
	require.NoError(t, err)

	data := transfertypes.FungibleTokenPacketData{
		Denom:    "foo",
		Amount:   "10000",
		Sender:   addr.String(),
		Receiver: contractAddr.Hex(),
		Memo: fmt.Sprintf(`{
			"evm": {
				"async_callback": {
					"id": 99,
					"contract_address": "%s"
				}
			}
		}`, contractAddr.Hex()),
	}

	dataBz, err := json.Marshal(&data)
	require.NoError(t, err)
	callbackBz, err := json.Marshal(evmhooks.AsyncCallback{
		Id:              99,
		ContractAddress: contractAddr.Hex(),
	})
	require.NoError(t, err)
	require.NoError(t, input.IBCHooksKeeper.SetAsyncCallback(ctx, sourcePort, sourceChannel, sequence, callbackBz))

	// hook should not be called to due to acl
	err = input.IBCHooksMiddleware.OnTimeoutPacket(ctx, channeltypes.Packet{
		Data:          dataBz,
		SourcePort:    sourcePort,
		SourceChannel: sourceChannel,
		Sequence:      sequence,
	}, addr)
	require.NoError(t, err)

	// check the contract state
	queryInputBz, err := abi.Pack("count")
	require.NoError(t, err)
	queryRes, logs, err := input.EVMKeeper.EVMCall(ctx, evmAddr, contractAddr, queryInputBz, nil, nil, nil)
	require.NoError(t, err)
	require.Equal(t, uint256.NewInt(0).Bytes32(), [32]byte(queryRes))
	require.Empty(t, logs)

	// set acl
	require.NoError(t, input.IBCHooksKeeper.SetAllowed(ctx, contractAddr[:], true))
	require.NoError(t, input.IBCHooksKeeper.SetAsyncCallback(ctx, sourcePort, sourceChannel, sequence, callbackBz))

	// success
	err = input.IBCHooksMiddleware.OnTimeoutPacket(ctx, channeltypes.Packet{
		Data:          dataBz,
		SourcePort:    sourcePort,
		SourceChannel: sourceChannel,
		Sequence:      sequence,
	}, addr)
	require.NoError(t, err)

	// check the contract state; increased by 99
	queryRes, logs, err = input.EVMKeeper.EVMCall(ctx, evmAddr, contractAddr, queryInputBz, nil, nil, nil)
	require.NoError(t, err)
	require.Equal(t, uint256.NewInt(99).Bytes32(), [32]byte(queryRes))
	require.Empty(t, logs)
}
