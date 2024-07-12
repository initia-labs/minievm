package evm_hooks_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/holiman/uint256"
	"github.com/stretchr/testify/require"

	transfertypes "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"
	channeltypes "github.com/cosmos/ibc-go/v8/modules/core/04-channel/types"

	nfttransfertypes "github.com/initia-labs/initia/x/ibc/nft-transfer/types"
	"github.com/initia-labs/minievm/x/evm/contracts/counter"
)

func Test_onReceiveIcs20Packet_noMemo(t *testing.T) {
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

	ack := input.IBCHooksMiddleware.OnRecvPacket(ctx, channeltypes.Packet{
		Data: dataBz,
	}, addr)

	require.True(t, ack.Success())
}

func Test_onReceiveIcs20Packet_memo(t *testing.T) {
	ctx, input := createDefaultTestInput(t)
	_, _, addr := keyPubAddr()
	evmAddr := common.BytesToAddress(addr.Bytes())

	codeBz, err := hexutil.Decode(counter.CounterBin)
	require.NoError(t, err)

	_, contractAddr, _, err := input.EVMKeeper.EVMCreate(ctx, evmAddr, codeBz, nil)
	require.NoError(t, err)

	abi, err := counter.CounterMetaData.GetAbi()
	require.NoError(t, err)

	inputBz, err := abi.Pack("increase")
	require.NoError(t, err)

	data := transfertypes.FungibleTokenPacketData{
		Denom:    "foo",
		Amount:   "10000",
		Sender:   addr.String(),
		Receiver: contractAddr.Hex(),
		Memo: fmt.Sprintf(`{
			"evm": {
				"message": {
					"contract_addr": "%s",
					"input": "%s"
				}
			}
		}`, contractAddr.Hex(), hexutil.Encode(inputBz)),
	}

	dataBz, err := json.Marshal(&data)
	require.NoError(t, err)

	// failed to due to acl
	ack := input.IBCHooksMiddleware.OnRecvPacket(ctx, channeltypes.Packet{
		Data: dataBz,
	}, addr)
	require.False(t, ack.Success())

	// set acl
	require.NoError(t, input.IBCHooksKeeper.SetAllowed(ctx, contractAddr[:], true))

	// success
	ack = input.IBCHooksMiddleware.OnRecvPacket(ctx, channeltypes.Packet{
		Data: dataBz,
	}, addr)
	require.True(t, ack.Success())

	queryInputBz, err := abi.Pack("count")
	require.NoError(t, err)

	// check the contract state
	queryRes, logs, err := input.EVMKeeper.EVMCall(ctx, evmAddr, contractAddr, queryInputBz, nil)
	require.NoError(t, err)
	require.Equal(t, uint256.NewInt(1).Bytes32(), [32]byte(queryRes))
	require.Empty(t, logs)
}

func Test_OnReceivePacket_ICS721(t *testing.T) {
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

	ack := input.IBCHooksMiddleware.OnRecvPacket(ctx, channeltypes.Packet{
		Data: dataBz,
	}, addr)

	require.True(t, ack.Success())
}

func Test_onReceiveIcs20Packet_memo_ICS721(t *testing.T) {
	ctx, input := createDefaultTestInput(t)
	_, _, addr := keyPubAddr()
	evmAddr := common.BytesToAddress(addr.Bytes())

	codeBz, err := hexutil.Decode(counter.CounterBin)
	require.NoError(t, err)

	_, contractAddr, _, err := input.EVMKeeper.EVMCreate(ctx, evmAddr, codeBz, nil)
	require.NoError(t, err)

	abi, err := counter.CounterMetaData.GetAbi()
	require.NoError(t, err)

	inputBz, err := abi.Pack("increase")
	require.NoError(t, err)

	data := transfertypes.FungibleTokenPacketData{
		Denom:    "foo",
		Amount:   "10000",
		Sender:   addr.String(),
		Receiver: contractAddr.Hex(),
		Memo: fmt.Sprintf(`{
			"evm": {
				"message": {
					"contract_addr": "%s",
					"input": "%s"
				}
			}
		}`, contractAddr.Hex(), hexutil.Encode(inputBz)),
	}

	dataBz, err := json.Marshal(&data)
	require.NoError(t, err)

	// failed to due to acl
	ack := input.IBCHooksMiddleware.OnRecvPacket(ctx, channeltypes.Packet{
		Data: dataBz,
	}, addr)
	require.False(t, ack.Success())

	// set acl
	require.NoError(t, input.IBCHooksKeeper.SetAllowed(ctx, contractAddr[:], true))

	// success
	ack = input.IBCHooksMiddleware.OnRecvPacket(ctx, channeltypes.Packet{
		Data: dataBz,
	}, addr)
	require.True(t, ack.Success())

	queryInputBz, err := abi.Pack("count")
	require.NoError(t, err)

	// check the contract state
	queryRes, logs, err := input.EVMKeeper.EVMCall(ctx, evmAddr, contractAddr, queryInputBz, nil)
	require.NoError(t, err)
	require.Equal(t, uint256.NewInt(1).Bytes32(), [32]byte(queryRes))
	require.Empty(t, logs)
}
