package evm_hooks_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/holiman/uint256"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
	transfertypes "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"
	channeltypes "github.com/cosmos/ibc-go/v8/modules/core/04-channel/types"

	nfttransfertypes "github.com/initia-labs/initia/x/ibc/nft-transfer/types"
	evm_hooks "github.com/initia-labs/minievm/app/ibc-hooks"
	"github.com/initia-labs/minievm/x/evm/contracts/counter"
	evmtypes "github.com/initia-labs/minievm/x/evm/types"
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

	_, contractAddr, _, err := input.EVMKeeper.EVMCreate(ctx, evmAddr, codeBz, nil, nil)
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

	pk := channeltypes.Packet{
		Data:               dataBz,
		DestinationPort:    "transfer-1",
		DestinationChannel: "channel-1",
	}

	// mint for approval test
	localDenom := evm_hooks.LocalDenom(pk, data.Denom)
	intermediateSender := sdk.MustAccAddressFromBech32(evm_hooks.DeriveIntermediateSender(pk.DestinationChannel, data.Sender))
	input.Faucet.Fund(ctx, intermediateSender, sdk.NewInt64Coin(localDenom, 1000000000))

	// failed to due to acl
	ack := input.IBCHooksMiddleware.OnRecvPacket(ctx, pk, addr)
	require.False(t, ack.Success())

	// set acl
	require.NoError(t, input.IBCHooksKeeper.SetAllowed(ctx, contractAddr[:], true))

	// success
	ack = input.IBCHooksMiddleware.OnRecvPacket(ctx, pk, addr)
	require.True(t, ack.Success())

	queryInputBz, err := abi.Pack("count")
	require.NoError(t, err)

	// check the contract state
	queryRes, err := input.EVMKeeper.EVMStaticCall(ctx, evmAddr, contractAddr, queryInputBz, nil)
	require.NoError(t, err)
	require.Equal(t, uint256.NewInt(1).Bytes32(), [32]byte(queryRes))

	// check allowance
	erc20Addr, err := evmtypes.DenomToContractAddr(ctx, input.EVMKeeper, localDenom)
	require.NoError(t, err)
	queryInputBz, err = input.EVMKeeper.ERC20Keeper().GetERC20ABI().Pack("allowance", common.BytesToAddress(intermediateSender.Bytes()), contractAddr)
	require.NoError(t, err)
	queryRes, err = input.EVMKeeper.EVMStaticCall(ctx, evmtypes.StdAddress, erc20Addr, queryInputBz, nil)
	require.NoError(t, err)
	require.Equal(t, uint256.NewInt(10000).Bytes32(), [32]byte(queryRes))
}

func Test_onReceiveIcs20Packet_memo_migrated(t *testing.T) {
	ctx, input := createDefaultTestInput(t)
	_, _, addr := keyPubAddr()
	evmAddr := common.BytesToAddress(addr.Bytes())

	codeBz, err := hexutil.Decode(counter.CounterBin)
	require.NoError(t, err)

	_, contractAddr, _, err := input.EVMKeeper.EVMCreate(ctx, evmAddr, codeBz, nil, nil)
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

	pk := channeltypes.Packet{
		Data:               dataBz,
		DestinationPort:    "transfer-1",
		DestinationChannel: "channel-1",
	}

	// mint for approval test
	localDenom := evm_hooks.LocalDenom(pk, data.Denom)
	// set the denom migration in OPChildKeeper
	l2Denom := "l2/" + localDenom
	input.OPChildKeeper.IBCToL2DenomMap[localDenom] = l2Denom
	intermediateSender := sdk.MustAccAddressFromBech32(evm_hooks.DeriveIntermediateSender(pk.DestinationChannel, data.Sender))
	input.Faucet.Fund(ctx, intermediateSender, sdk.NewInt64Coin(l2Denom, 1000000000))

	// failed to due to acl
	ack := input.IBCHooksMiddleware.OnRecvPacket(ctx, pk, addr)
	require.False(t, ack.Success())

	// set acl
	require.NoError(t, input.IBCHooksKeeper.SetAllowed(ctx, contractAddr[:], true))

	// success
	ack = input.IBCHooksMiddleware.OnRecvPacket(ctx, pk, addr)
	require.True(t, ack.Success())

	queryInputBz, err := abi.Pack("count")
	require.NoError(t, err)

	// check the contract state
	queryRes, err := input.EVMKeeper.EVMStaticCall(ctx, evmAddr, contractAddr, queryInputBz, nil)
	require.NoError(t, err)
	require.Equal(t, uint256.NewInt(1).Bytes32(), [32]byte(queryRes))

	// check allowance
	erc20Addr, err := evmtypes.DenomToContractAddr(ctx, input.EVMKeeper, l2Denom)
	require.NoError(t, err)
	queryInputBz, err = input.EVMKeeper.ERC20Keeper().GetERC20ABI().Pack("allowance", common.BytesToAddress(intermediateSender.Bytes()), contractAddr)
	require.NoError(t, err)
	queryRes, err = input.EVMKeeper.EVMStaticCall(ctx, evmtypes.StdAddress, erc20Addr, queryInputBz, nil)
	require.NoError(t, err)
	require.Equal(t, uint256.NewInt(10000).Bytes32(), [32]byte(queryRes))
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

func Test_onReceivePacket_memo_ICS721(t *testing.T) {
	ctx, input := createDefaultTestInput(t)
	_, _, addr := keyPubAddr()

	evmAddr := common.BytesToAddress(addr.Bytes())

	codeBz, err := hexutil.Decode(counter.CounterBin)
	require.NoError(t, err)

	_, contractAddr, _, err := input.EVMKeeper.EVMCreate(ctx, evmAddr, codeBz, nil, nil)
	require.NoError(t, err)

	abi, err := counter.CounterMetaData.GetAbi()
	require.NoError(t, err)

	inputBz, err := abi.Pack("increase")
	require.NoError(t, err)

	data := nfttransfertypes.NonFungibleTokenPacketData{
		ClassId:   "classId",
		ClassUri:  "classUri",
		ClassData: "classData",
		TokenIds:  []string{"tokenId"},
		TokenUris: []string{"tokenUri"},
		TokenData: []string{"tokenData"},
		Sender:    addr.String(),
		Receiver:  contractAddr.Hex(),
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

	pk := channeltypes.Packet{
		Data:               dataBz,
		DestinationPort:    "nfttransfer-1",
		DestinationChannel: "channel-1",
	}

	// mint for approval test
	localClassId := evm_hooks.LocalClassId(pk, data.ClassId)
	intermediateSender := sdk.MustAccAddressFromBech32(evm_hooks.DeriveIntermediateSender(pk.DestinationChannel, data.Sender))
	err = input.EVMKeeper.ERC721Keeper().CreateOrUpdateClass(ctx, localClassId, data.ClassUri, data.ClassData)
	require.NoError(t, err)
	err = input.EVMKeeper.ERC721Keeper().Mints(
		ctx,
		intermediateSender,
		localClassId,
		[]string{"tokenId"},
		[]string{"tokenUri"},
		[]string{"tokenData"},
	)
	require.NoError(t, err)

	// failed to due to acl
	ack := input.IBCHooksMiddleware.OnRecvPacket(ctx, pk, addr)
	require.False(t, ack.Success())

	// set acl
	require.NoError(t, input.IBCHooksKeeper.SetAllowed(ctx, contractAddr[:], true))

	// success
	ack = input.IBCHooksMiddleware.OnRecvPacket(ctx, pk, addr)
	require.True(t, ack.Success())

	queryInputBz, err := abi.Pack("count")
	require.NoError(t, err)

	// check the contract state
	queryRes, logs, err := input.EVMKeeper.EVMCall(ctx, evmAddr, contractAddr, queryInputBz, nil, nil, nil)
	require.NoError(t, err)
	require.Equal(t, uint256.NewInt(1).Bytes32(), [32]byte(queryRes))
	require.Empty(t, logs)

	// check allowance
	tokenId, ok := evmtypes.TokenIdToBigInt(localClassId, data.TokenIds[0])
	require.True(t, ok)
	erc721Addr, err := input.EVMKeeper.GetContractAddrByClassId(ctx, localClassId)
	require.NoError(t, err)
	queryInputBz, err = input.EVMKeeper.ERC721Keeper().GetERC721ABI().Pack("getApproved", tokenId)
	require.NoError(t, err)
	queryRes, err = input.EVMKeeper.EVMStaticCall(ctx, evmtypes.StdAddress, erc721Addr, queryInputBz, nil)
	require.NoError(t, err)
	require.Equal(t, contractAddr.Bytes(), common.HexToAddress(hexutil.Encode(queryRes)).Bytes())
}
