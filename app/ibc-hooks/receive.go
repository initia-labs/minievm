package evm_hooks

import (
	"fmt"
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	transfertypes "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"
	channeltypes "github.com/cosmos/ibc-go/v8/modules/core/04-channel/types"
	ibcexported "github.com/cosmos/ibc-go/v8/modules/core/exported"

	ibchooks "github.com/initia-labs/initia/x/ibc-hooks"
	nfttransfertypes "github.com/initia-labs/initia/x/ibc/nft-transfer/types"

	evmkeeper "github.com/initia-labs/minievm/x/evm/keeper"
	evmtypes "github.com/initia-labs/minievm/x/evm/types"
)

func (h EVMHooks) onRecvIcs20Packet(
	ctx sdk.Context,
	im ibchooks.IBCMiddleware,
	packet channeltypes.Packet,
	relayer sdk.AccAddress,
	data transfertypes.FungibleTokenPacketData,
) ibcexported.Acknowledgement {
	isEVMRouted, hookData, err := validateAndParseMemo(data.GetMemo())
	if !isEVMRouted || hookData.Message == nil {
		return im.App.OnRecvPacket(ctx, packet, relayer)
	} else if err != nil {
		return newEmitErrorAcknowledgement(err)
	}

	msg := hookData.Message
	if allowed, err := h.checkACL(im, ctx, msg.ContractAddr); err != nil {
		return newEmitErrorAcknowledgement(err)
	} else if !allowed {
		return newEmitErrorAcknowledgement(fmt.Errorf("contract `%s` not allowed to be used in ibchooks", msg.ContractAddr))
	}

	// Validate whether the receiver is correctly specified or not.
	if err := validateReceiver(msg, data.Receiver); err != nil {
		return newEmitErrorAcknowledgement(err)
	}

	// Calculate the receiver / contract caller based on the packet's channel and sender
	intermediateSender := DeriveIntermediateSender(packet.GetDestChannel(), data.GetSender())

	// The funds sent on this packet need to be transferred to the intermediary account for the sender.
	// For this, we override the ICS20 packet's Receiver (essentially hijacking the funds to this new address)
	// and execute the underlying OnRecvPacket() call (which should eventually land on the transfer app's
	// relay.go and send the funds to the intermediary account.
	//
	// If that succeeds, we make the contract call
	data.Receiver = intermediateSender
	packet.Data = data.GetBytes()

	ack := im.App.OnRecvPacket(ctx, packet, relayer)
	if !ack.Success() {
		return ack
	}

	msg.Sender = intermediateSender
	localDenom := LocalDenom(packet, data.Denom)
	_, err = h.approveERC20(ctx, intermediateSender, common.HexToAddress(msg.ContractAddr), localDenom, data.Amount)
	if err != nil {
		return newEmitErrorAcknowledgement(err)
	}
	_, err = h.execMsg(ctx, msg)
	if err != nil {
		return newEmitErrorAcknowledgement(err)
	}

	return ack
}

func (h EVMHooks) approveERC20(ctx sdk.Context, intermediateSender string, contractAddr common.Address, denom, amount string) (*evmtypes.MsgCallResponse, error) {
	amt, ok := new(big.Int).SetString(amount, 10)
	if !ok {
		return nil, fmt.Errorf("failed to parse amount %s", amount)
	}

	erc20ABI := h.evmKeeper.ERC20Keeper().GetERC20ABI()
	inputBz, err := erc20ABI.Pack("approve", contractAddr, amt)
	if err != nil {
		return nil, err
	}

	erc20Addr, err := h.evmKeeper.GetContractAddrByDenom(ctx, denom)
	if err != nil {
		return nil, err
	}

	msg := &evmtypes.MsgCall{
		Sender:       intermediateSender,
		ContractAddr: erc20Addr.Hex(),
		Input:        hexutil.Encode(inputBz),
	}

	evmMsgServer := evmkeeper.NewMsgServerImpl(h.evmKeeper)
	return evmMsgServer.Call(ctx, msg)
}

func (h EVMHooks) onRecvIcs721Packet(
	ctx sdk.Context,
	im ibchooks.IBCMiddleware,
	packet channeltypes.Packet,
	relayer sdk.AccAddress,
	data nfttransfertypes.NonFungibleTokenPacketData,
) ibcexported.Acknowledgement {
	isEVMRouted, hookData, err := validateAndParseMemo(data.GetMemo())
	if !isEVMRouted || hookData.Message == nil {
		return im.App.OnRecvPacket(ctx, packet, relayer)
	} else if err != nil {
		return newEmitErrorAcknowledgement(err)
	}

	msg := hookData.Message
	if allowed, err := h.checkACL(im, ctx, msg.ContractAddr); err != nil {
		return newEmitErrorAcknowledgement(err)
	} else if !allowed {
		return newEmitErrorAcknowledgement(fmt.Errorf("contract `%s` is not allowed to be used in ibchooks", msg.ContractAddr))
	}

	// Validate whether the receiver is correctly specified or not.
	if err := validateReceiver(msg, data.Receiver); err != nil {
		return newEmitErrorAcknowledgement(err)
	}

	// Calculate the receiver / contract caller based on the packet's channel and sender
	intermediateSender := DeriveIntermediateSender(packet.GetDestChannel(), data.GetSender())

	// The funds sent on this packet need to be transferred to the intermediary account for the sender.
	// For this, we override the ICS721 packet's Receiver (essentially hijacking the funds to this new address)
	// and execute the underlying OnRecvPacket() call (which should eventually land on the transfer app's
	// relay.go and send the funds to the intermediary account.
	//
	// If that succeeds, we make the contract call
	data.Receiver = intermediateSender
	packet.Data = data.GetBytes(packet.SourcePort)

	ack := im.App.OnRecvPacket(ctx, packet, relayer)
	if !ack.Success() {
		return ack
	}

	// approve the transfer of the NFT to the contract
	msg.Sender = intermediateSender
	localClassId := LocalClassId(packet, data.ClassId)
	for _, tokenId := range data.TokenIds {
		_, err = h.approveERC721(ctx, intermediateSender, common.HexToAddress(msg.ContractAddr), localClassId, tokenId)
		if err != nil {
			return newEmitErrorAcknowledgement(err)
		}
	}
	_, err = h.execMsg(ctx, msg)
	if err != nil {
		return newEmitErrorAcknowledgement(err)
	}

	return ack
}

func (h EVMHooks) execMsg(ctx sdk.Context, msg *evmtypes.MsgCall) (*evmtypes.MsgCallResponse, error) {
	evmMsgServer := evmkeeper.NewMsgServerImpl(h.evmKeeper)
	return evmMsgServer.Call(ctx, msg)
}

func (h EVMHooks) approveERC721(ctx sdk.Context, intermediateSender string, contractAddr common.Address, classId, tokenId string) (*evmtypes.MsgCallResponse, error) {
	tid, ok := evmtypes.TokenIdToBigInt(classId, tokenId)
	if !ok {
		return nil, evmtypes.ErrInvalidTokenId
	}

	erc721ABI := h.evmKeeper.ERC721Keeper().GetERC721ABI()
	inputBz, err := erc721ABI.Pack("approve", contractAddr, tid)
	if err != nil {
		return nil, err
	}

	erc721Addr, err := h.evmKeeper.GetContractAddrByClassId(ctx, classId)
	if err != nil {
		return nil, err
	}

	msg := &evmtypes.MsgCall{
		Sender:       intermediateSender,
		ContractAddr: erc721Addr.Hex(),
		Input:        hexutil.Encode(inputBz),
	}

	evmMsgServer := evmkeeper.NewMsgServerImpl(h.evmKeeper)
	return evmMsgServer.Call(ctx, msg)
}
