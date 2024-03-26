package evm_hooks_test

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/holiman/uint256"
	"github.com/stretchr/testify/require"

	transfertypes "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"
	channeltypes "github.com/cosmos/ibc-go/v8/modules/core/04-channel/types"

	"github.com/initia-labs/minievm/x/evm/contracts/counter"
)

func Test_onReceiveIcs20Packet_noMemo(t *testing.T) {
	ctx, input := createDefaultTestInput(t)
	_, _, addr := keyPubAddr()
	_, _, addr2 := keyPubAddr()

	// don't have any memo
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

	codeBz, err := hex.DecodeString(strings.TrimPrefix(counter.CounterBin, "0x"))
	require.NoError(t, err)

	_, contractAddr, err := input.EVMKeeper.EVMCreate(ctx, evmAddr, codeBz)
	require.NoError(t, err)

	abi, err := counter.CounterMetaData.GetAbi()
	require.NoError(t, err)

	inputBz, err := abi.Pack("increase")
	require.NoError(t, err)

	// don't have any memo
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
		}`, contractAddr.Hex(), hex.EncodeToString(inputBz)),
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
	queryRes, logs, err := input.EVMKeeper.EVMCall(ctx, evmAddr, contractAddr, queryInputBz)
	require.NoError(t, err)
	require.Equal(t, uint256.NewInt(1).Bytes32(), [32]byte(queryRes))
	require.Empty(t, logs)
}
