package evm_hooks

import (
	"encoding/json"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	transfertypes "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"
	channeltypes "github.com/cosmos/ibc-go/v8/modules/core/04-channel/types"
	ibcexported "github.com/cosmos/ibc-go/v8/modules/core/exported"

	ibchooks "github.com/initia-labs/initia/x/ibc-hooks"

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
		return im.App.OnRecvPacket(ctx, packet, relayer)
	}

	// Validate whether the receiver is correctly specified or not.
	if err := validateReceiver(msg, data.Receiver); err != nil {
		return newEmitErrorAcknowledgement(err)
	}

	// Calculate the receiver / contract caller based on the packet's channel and sender
	intermediateSender := deriveIntermediateSender(packet.GetDestChannel(), data.GetSender())

	// The funds sent on this packet need to be transferred to the intermediary account for the sender.
	// For this, we override the ICS20 packet's Receiver (essentially hijacking the funds to this new address)
	// and execute the underlying OnRecvPacket() call (which should eventually land on the transfer app's
	// relay.go and send the funds to the intermediary account.
	//
	// If that succeeds, we make the contract call
	data.Receiver = intermediateSender
	bz, err := json.Marshal(data)
	if err != nil {
		return newEmitErrorAcknowledgement(err)
	}
	packet.Data = bz

	ack := im.App.OnRecvPacket(ctx, packet, relayer)
	if !ack.Success() {
		return ack
	}

	fmt.Println("SIBONG")
	msg.Sender = intermediateSender
	_, err = h.execMsg(ctx, msg)
	if err != nil {
		return newEmitErrorAcknowledgement(err)
	}

	return ack
}

// func (h EVMHooks) onRecvIcs721Packet(
// 	ctx sdk.Context,
// 	im ibchooks.IBCMiddleware,
// 	packet channeltypes.Packet,
// 	relayer sdk.AccAddress,
// 	data nfttransfertypes.NonFungibleTokenPacketData,
// ) ibcexported.Acknowledgement {
// 	isEVMRouted, hookData, err := validateAndParseMemo(data.GetMemo())
// 	if !isEVMRouted || hookData.Message == nil {
// 		return im.App.OnRecvPacket(ctx, packet, relayer)
// 	} else if err != nil {
// 		return newEmitErrorAcknowledgement(err)
// 	}

// 	msg := hookData.Message
// 	if allowed, err := h.checkACL(im, ctx, msg.ContractAddr); err != nil {
// 		return newEmitErrorAcknowledgement(err)
// 	} else if !allowed {
// 		return im.App.OnRecvPacket(ctx, packet, relayer)
// 	}

// 	// Validate whether the receiver is correctly specified or not.
// 	if err := validateReceiver(msg, data.Receiver); err != nil {
// 		return newEmitErrorAcknowledgement(err)
// 	}

// 	// Calculate the receiver / contract caller based on the packet's channel and sender
// 	intermediateSender := deriveIntermediateSender(packet.GetDestChannel(), data.GetSender())

// 	// The funds sent on this packet need to be transferred to the intermediary account for the sender.
// 	// For this, we override the ICS721 packet's Receiver (essentially hijacking the funds to this new address)
// 	// and execute the underlying OnRecvPacket() call (which should eventually land on the transfer app's
// 	// relay.go and send the funds to the intermediary account.
// 	//
// 	// If that succeeds, we make the contract call
// 	data.Receiver = intermediateSender
// 	bz, err := json.Marshal(data)
// 	if err != nil {
// 		return newEmitErrorAcknowledgement(err)
// 	}
// 	packet.Data = bz

// 	ack := im.App.OnRecvPacket(ctx, packet, relayer)
// 	if !ack.Success() {
// 		return ack
// 	}

// 	msg.Sender = intermediateSender
// 	_, err = h.execMsg(ctx, msg)
// 	if err != nil {
// 		return newEmitErrorAcknowledgement(err)
// 	}

// 	return ack
// }

func (im EVMHooks) execMsg(ctx sdk.Context, msg *evmtypes.MsgCall) (*evmtypes.MsgCallResponse, error) {
	evmMsgServer := evmkeeper.NewMsgServerImpl(im.evmKeeper)
	return evmMsgServer.Call(ctx, msg)
}
