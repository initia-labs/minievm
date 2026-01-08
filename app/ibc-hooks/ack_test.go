package evm_hooks_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"

	transfertypes "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"
	channeltypes "github.com/cosmos/ibc-go/v8/modules/core/04-channel/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/holiman/uint256"
	"github.com/stretchr/testify/require"

	nfttransfertypes "github.com/initia-labs/initia/x/ibc/nft-transfer/types"
	evmhooks "github.com/initia-labs/minievm/app/ibc-hooks"
	"github.com/initia-labs/minievm/x/evm/contracts/counter"
)

func Test_onAckIcs20Packet_noMemo(t *testing.T) {
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

	ackBz, err := json.Marshal(channeltypes.NewResultAcknowledgement([]byte{byte(1)}))
	require.NoError(t, err)

	err = input.IBCHooksMiddleware.OnAcknowledgementPacket(ctx, channeltypes.Packet{
		Data: dataBz,
	}, ackBz, addr)
	require.NoError(t, err)
}

func Test_onAckIcs20Packet_memo(t *testing.T) {
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

	successAckBz := channeltypes.NewResultAcknowledgement([]byte{byte(1)}).Acknowledgement()
	failedAckBz := channeltypes.NewErrorAcknowledgement(errors.New("failed")).Acknowledgement()
	callbackBz, err := json.Marshal(evmhooks.AsyncCallback{
		Id:              99,
		ContractAddress: contractAddr.Hex(),
	})
	require.NoError(t, err)
	require.NoError(t, input.IBCHooksKeeper.SetAsyncCallback(ctx, sourcePort, sourceChannel, sequence, callbackBz))

	// hook should not be called to due to acl
	err = input.IBCHooksMiddleware.OnAcknowledgementPacket(ctx, channeltypes.Packet{
		Data:          dataBz,
		SourcePort:    sourcePort,
		SourceChannel: sourceChannel,
		Sequence:      sequence,
	}, successAckBz, addr)
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

	// success with success ack
	err = input.IBCHooksMiddleware.OnAcknowledgementPacket(ctx, channeltypes.Packet{
		Data:          dataBz,
		SourcePort:    sourcePort,
		SourceChannel: sourceChannel,
		Sequence:      sequence,
	}, successAckBz, addr)
	require.NoError(t, err)

	// check the contract state; increased by 99 if ack is success
	queryRes, logs, err = input.EVMKeeper.EVMCall(ctx, evmAddr, contractAddr, queryInputBz, nil, nil, nil)
	require.NoError(t, err)
	require.Equal(t, uint256.NewInt(99).Bytes32(), [32]byte(queryRes))
	require.Empty(t, logs)

	// success with failed ack
	require.NoError(t, input.IBCHooksKeeper.SetAsyncCallback(ctx, sourcePort, sourceChannel, sequence, callbackBz))
	err = input.IBCHooksMiddleware.OnAcknowledgementPacket(ctx, channeltypes.Packet{
		Data:          dataBz,
		SourcePort:    sourcePort,
		SourceChannel: sourceChannel,
		Sequence:      sequence,
	}, failedAckBz, addr)
	require.NoError(t, err)

	// check the contract state; increased by 1 if ack is failed
	queryRes, logs, err = input.EVMKeeper.EVMCall(ctx, evmAddr, contractAddr, queryInputBz, nil, nil, nil)
	require.NoError(t, err)
	require.Equal(t, uint256.NewInt(100).Bytes32(), [32]byte(queryRes))
	require.Empty(t, logs)
}

func Test_OnAckPacket_ICS721(t *testing.T) {
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

	ackBz, err := json.Marshal(channeltypes.NewResultAcknowledgement([]byte{byte(1)}))
	require.NoError(t, err)

	err = input.IBCHooksMiddleware.OnAcknowledgementPacket(ctx, channeltypes.Packet{
		Data: dataBz,
	}, ackBz, addr)
	require.NoError(t, err)
}

func Test_onAckPacket_memo_ICS721(t *testing.T) {
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

	data := nfttransfertypes.NonFungibleTokenPacketData{
		ClassId:   "classId",
		ClassUri:  "classUri",
		ClassData: "classData",
		TokenIds:  []string{"tokenId"},
		TokenUris: []string{"tokenUri"},
		TokenData: []string{"tokenData"},
		Sender:    addr.String(),
		Receiver:  "0x1::Counter::increase",
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

	successAckBz := channeltypes.NewResultAcknowledgement([]byte{byte(1)}).Acknowledgement()
	failedAckBz := channeltypes.NewErrorAcknowledgement(errors.New("failed")).Acknowledgement()
	callbackBz, err := json.Marshal(evmhooks.AsyncCallback{
		Id:              99,
		ContractAddress: contractAddr.Hex(),
	})
	require.NoError(t, err)
	require.NoError(t, input.IBCHooksKeeper.SetAsyncCallback(ctx, sourcePort, sourceChannel, sequence, callbackBz))

	// hook should not be called to due to acl
	err = input.IBCHooksMiddleware.OnAcknowledgementPacket(ctx, channeltypes.Packet{
		Data:          dataBz,
		SourcePort:    sourcePort,
		SourceChannel: sourceChannel,
		Sequence:      sequence,
	}, successAckBz, addr)
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

	// success with success ack
	err = input.IBCHooksMiddleware.OnAcknowledgementPacket(ctx, channeltypes.Packet{
		Data:          dataBz,
		SourcePort:    sourcePort,
		SourceChannel: sourceChannel,
		Sequence:      sequence,
	}, successAckBz, addr)
	require.NoError(t, err)

	// check the contract state; increased by 99 if ack is success
	queryRes, logs, err = input.EVMKeeper.EVMCall(ctx, evmAddr, contractAddr, queryInputBz, nil, nil, nil)
	require.NoError(t, err)
	require.Equal(t, uint256.NewInt(99).Bytes32(), [32]byte(queryRes))
	require.Empty(t, logs)

	// success with failed ack
	require.NoError(t, input.IBCHooksKeeper.SetAsyncCallback(ctx, sourcePort, sourceChannel, sequence, callbackBz))
	err = input.IBCHooksMiddleware.OnAcknowledgementPacket(ctx, channeltypes.Packet{
		Data:          dataBz,
		SourcePort:    sourcePort,
		SourceChannel: sourceChannel,
		Sequence:      sequence,
	}, failedAckBz, addr)
	require.NoError(t, err)

	// check the contract state; increased by 1 if ack is failed
	queryRes, logs, err = input.EVMKeeper.EVMCall(ctx, evmAddr, contractAddr, queryInputBz, nil, nil, nil)
	require.NoError(t, err)
	require.Equal(t, uint256.NewInt(100).Bytes32(), [32]byte(queryRes))
	require.Empty(t, logs)
}
