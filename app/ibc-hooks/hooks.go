package evm_hooks

import (
	"cosmossdk.io/core/address"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	capabilitytypes "github.com/cosmos/ibc-go/modules/capability/types"
	clienttypes "github.com/cosmos/ibc-go/v8/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v8/modules/core/04-channel/types"
	ibcexported "github.com/cosmos/ibc-go/v8/modules/core/exported"
	"github.com/ethereum/go-ethereum/accounts/abi"

	ibchooks "github.com/initia-labs/initia/x/ibc-hooks"
	"github.com/initia-labs/minievm/x/evm/contracts/i_ibc_async_callback"
	evmkeeper "github.com/initia-labs/minievm/x/evm/keeper"
	evmtypes "github.com/initia-labs/minievm/x/evm/types"
)

var (
	_ ibchooks.OnRecvPacketOverrideHooks            = EVMHooks{}
	_ ibchooks.OnAcknowledgementPacketOverrideHooks = EVMHooks{}
	_ ibchooks.OnTimeoutPacketOverrideHooks         = EVMHooks{}
	_ ibchooks.SendPacketOverrideHooks              = EVMHooks{}
)

type EVMHooks struct {
	codec            codec.Codec
	ac               address.Codec
	evmKeeper        *evmkeeper.Keeper
	opchildKeeper    OPChildKeeper
	asyncCallbackABI *abi.ABI
}

func NewEVMHooks(codec codec.Codec, ac address.Codec, evmKeeper *evmkeeper.Keeper, opchildKeeper OPChildKeeper) *EVMHooks {
	abi, err := i_ibc_async_callback.IIbcAsyncCallbackMetaData.GetAbi()
	if err != nil {
		panic(err)
	}

	return &EVMHooks{
		codec:            codec,
		ac:               ac,
		evmKeeper:        evmKeeper,
		opchildKeeper:    opchildKeeper,
		asyncCallbackABI: abi,
	}
}

func (h EVMHooks) OnRecvPacketOverride(im ibchooks.IBCMiddleware, ctx sdk.Context, packet channeltypes.Packet, relayer sdk.AccAddress) ibcexported.Acknowledgement {
	if isIcs20, ics20Data := isIcs20Packet(packet.GetData()); isIcs20 {
		return h.onRecvIcs20Packet(ctx, im, packet, relayer, ics20Data)
	}

	if isIcs721, ics721Data := isIcs721Packet(packet.GetData()); isIcs721 {
		return h.onRecvIcs721Packet(ctx, im, packet, relayer, ics721Data)
	}

	return im.App.OnRecvPacket(ctx, packet, relayer)
}

func (h EVMHooks) SendPacketOverride(im ibchooks.ICS4Middleware, ctx sdk.Context, chanCap *capabilitytypes.Capability, sourcePort string, sourceChannel string, timeoutHeight clienttypes.Height, timeoutTimestamp uint64, data []byte) (uint64, error) {
	if isIcs20, ics20Data := isIcs20Packet(data); isIcs20 {
		return h.sendIcs20Packet(ctx, im, chanCap, sourcePort, sourceChannel, timeoutHeight, timeoutTimestamp, ics20Data)
	}

	if isIcs721, ics721Data := isIcs721Packet(data); isIcs721 {
		return h.sendIcs721Packet(ctx, im, chanCap, sourcePort, sourceChannel, timeoutHeight, timeoutTimestamp, ics721Data)
	}

	return im.ICS4Wrapper.SendPacket(ctx, chanCap, sourcePort, sourceChannel, timeoutHeight, timeoutTimestamp, data)
}

func (h EVMHooks) OnAcknowledgementPacketOverride(im ibchooks.IBCMiddleware, ctx sdk.Context, packet channeltypes.Packet, acknowledgement []byte, relayer sdk.AccAddress) error {
	if isIcs20, ics20Data := isIcs20Packet(packet.GetData()); isIcs20 {
		return h.onAckIcs20Packet(ctx, im, packet, acknowledgement, relayer, ics20Data)
	}

	if isIcs721, ics721Data := isIcs721Packet(packet.GetData()); isIcs721 {
		return h.onAckIcs721Packet(ctx, im, packet, acknowledgement, relayer, ics721Data)
	}

	return im.App.OnAcknowledgementPacket(ctx, packet, acknowledgement, relayer)
}

func (h EVMHooks) OnTimeoutPacketOverride(im ibchooks.IBCMiddleware, ctx sdk.Context, packet channeltypes.Packet, relayer sdk.AccAddress) error {
	if isIcs20, ics20Data := isIcs20Packet(packet.GetData()); isIcs20 {
		return h.onTimeoutIcs20Packet(ctx, im, packet, relayer, ics20Data)
	}

	if isIcs721, ics721Data := isIcs721Packet(packet.GetData()); isIcs721 {
		return h.onTimeoutIcs721Packet(ctx, im, packet, relayer, ics721Data)
	}

	return im.App.OnTimeoutPacket(ctx, packet, relayer)
}

func (h EVMHooks) checkACL(im ibchooks.IBCMiddleware, ctx sdk.Context, addrStr string) (bool, error) {
	addr, err := evmtypes.ContractAddressFromString(h.ac, addrStr)
	if err != nil {
		return false, err
	}

	return im.HooksKeeper.GetAllowed(ctx, addr.Bytes())
}
