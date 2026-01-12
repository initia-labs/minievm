package evm_hooks

import (
	"errors"

	"cosmossdk.io/collections"
	sdk "github.com/cosmos/cosmos-sdk/types"
	transfertypes "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"
	channeltypes "github.com/cosmos/ibc-go/v8/modules/core/04-channel/types"
	"github.com/ethereum/go-ethereum/common/hexutil"

	ibchooks "github.com/initia-labs/initia/x/ibc-hooks"
	"github.com/initia-labs/initia/x/ibc-hooks/types"
	nfttransfertypes "github.com/initia-labs/initia/x/ibc/nft-transfer/types"

	evmtypes "github.com/initia-labs/minievm/x/evm/types"
)

func (h EVMHooks) onTimeoutIcs20Packet(
	ctx sdk.Context,
	im ibchooks.IBCMiddleware,
	packet channeltypes.Packet,
	relayer sdk.AccAddress,
	data transfertypes.FungibleTokenPacketData,
) error {
	return h.handleOnTimeout(ctx, im, packet, relayer, data.Sender)
}

func (h EVMHooks) onTimeoutIcs721Packet(
	ctx sdk.Context,
	im ibchooks.IBCMiddleware,
	packet channeltypes.Packet,
	relayer sdk.AccAddress,
	data nfttransfertypes.NonFungibleTokenPacketData,
) error {
	return h.handleOnTimeout(ctx, im, packet, relayer, data.Sender)
}

func (h EVMHooks) handleOnTimeout(
	ctx sdk.Context,
	im ibchooks.IBCMiddleware,
	packet channeltypes.Packet,
	relayer sdk.AccAddress,
	sender string,
) error {
	if err := im.App.OnTimeoutPacket(ctx, packet, relayer); err != nil {
		return err
	}

	// if no async callback, return early
	bz, err := im.HooksKeeper.GetAsyncCallback(ctx, packet.GetSourcePort(), packet.GetSourceChannel(), packet.GetSequence())
	if err != nil && errors.Is(err, collections.ErrNotFound) {
		return nil
	} else if err != nil {
		h.evmKeeper.Logger(ctx).Error("failed to get async callback", "error", err)
		ctx.EventManager().EmitEvent(sdk.NewEvent(
			types.EventTypeHookFailed,
			sdk.NewAttribute(types.AttributeKeyReason, "failed to get async callback"),
			sdk.NewAttribute(types.AttributeKeyError, err.Error()),
		))

		return nil
	}

	// ignore error on removal; it should not happen
	_ = im.HooksKeeper.RemoveAsyncCallback(ctx, packet.GetSourcePort(), packet.GetSourceChannel(), packet.GetSequence())

	var callback AsyncCallback
	if err := callback.UnmarshalJSON(bz); err != nil {
		h.evmKeeper.Logger(ctx).Error("failed to unmarshal async callback", "error", err)
		ctx.EventManager().EmitEvent(sdk.NewEvent(
			types.EventTypeHookFailed,
			sdk.NewAttribute(types.AttributeKeyReason, "failed to unmarshal async callback"),
			sdk.NewAttribute(types.AttributeKeyError, err.Error()),
		))
		return nil
	}

	// create a new cache context to ignore errors during
	// the execution of the callback
	cacheCtx, write := ctx.CacheContext()

	if allowed, err := h.checkACL(im, cacheCtx, callback.ContractAddress); err != nil {
		h.evmKeeper.Logger(cacheCtx).Error("failed to check ACL", "error", err)
		ctx.EventManager().EmitEvent(sdk.NewEvent(
			types.EventTypeHookFailed,
			sdk.NewAttribute(types.AttributeKeyReason, "failed to check ACL"),
			sdk.NewAttribute(types.AttributeKeyError, err.Error()),
		))

		return nil
	} else if !allowed {
		h.evmKeeper.Logger(cacheCtx).Error("failed to check ACL", "reason", "not allowed")
		ctx.EventManager().EmitEvent(sdk.NewEvent(
			types.EventTypeHookFailed,
			sdk.NewAttribute(types.AttributeKeyReason, "failed to check ACL"),
			sdk.NewAttribute(types.AttributeKeyError, "not allowed"),
		))

		return nil
	}

	inputBz, err := h.asyncCallbackABI.Pack(functionNameTimeout, callback.Id)
	if err != nil {
		h.evmKeeper.Logger(cacheCtx).Error("failed to pack input", "error", err)
		return nil
	}

	_, err = h.execMsg(cacheCtx, &evmtypes.MsgCall{
		Sender:       sender,
		ContractAddr: callback.ContractAddress,
		Input:        hexutil.Encode(inputBz),
	})
	if err != nil {
		h.evmKeeper.Logger(cacheCtx).Error("failed to execute callback", "error", err)
		ctx.EventManager().EmitEvent(sdk.NewEvent(
			types.EventTypeHookFailed,
			sdk.NewAttribute(types.AttributeKeyReason, "failed to execute callback"),
			sdk.NewAttribute(types.AttributeKeyError, err.Error()),
		))

		return nil
	}

	// write the cache context only if the callback execution was successful
	write()

	return nil
}
