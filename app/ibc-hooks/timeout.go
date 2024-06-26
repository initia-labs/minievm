package evm_hooks

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	transfertypes "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"
	channeltypes "github.com/cosmos/ibc-go/v8/modules/core/04-channel/types"
	"github.com/ethereum/go-ethereum/common/hexutil"

	ibchooks "github.com/initia-labs/initia/x/ibc-hooks"
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
	if err := im.App.OnTimeoutPacket(ctx, packet, relayer); err != nil {
		return err
	}

	isEVMRouted, hookData, err := validateAndParseMemo(data.GetMemo())
	if !isEVMRouted || hookData.AsyncCallback == nil {
		return nil
	} else if err != nil {
		h.evmKeeper.Logger(ctx).Error("failed to parse memo", "error", err)
		return nil
	}

	// create a new cache context to ignore errors during
	// the execution of the callback
	cacheCtx, write := ctx.CacheContext()

	callback := hookData.AsyncCallback
	if allowed, err := h.checkACL(im, cacheCtx, callback.ContractAddress); err != nil {
		h.evmKeeper.Logger(cacheCtx).Error("failed to check ACL", "error", err)
		return nil
	} else if !allowed {
		h.evmKeeper.Logger(cacheCtx).Error("failed to check ACL", "not allowed")
		return nil
	}

	inputBz, err := h.asyncCallbackABI.Pack(functionNameTimeout, callback.Id)
	if err != nil {
		h.evmKeeper.Logger(cacheCtx).Error("failed to pack input", "error", err)
		return nil
	}

	_, err = h.execMsg(cacheCtx, &evmtypes.MsgCall{
		Sender:       data.Sender,
		ContractAddr: callback.ContractAddress,
		Input:        hexutil.Encode(inputBz),
	})
	if err != nil {
		h.evmKeeper.Logger(cacheCtx).Error("failed to execute callback", "error", err)
		return nil
	}

	// write the cache context only if the callback execution was successful
	write()

	return nil
}

func (h EVMHooks) onTimeoutIcs721Packet(
	ctx sdk.Context,
	im ibchooks.IBCMiddleware,
	packet channeltypes.Packet,
	relayer sdk.AccAddress,
	data nfttransfertypes.NonFungibleTokenPacketData,
) error {
	if err := im.App.OnTimeoutPacket(ctx, packet, relayer); err != nil {
		return err
	}

	isEVMRouted, hookData, err := validateAndParseMemo(data.GetMemo())
	if !isEVMRouted || hookData.AsyncCallback == nil {
		return nil
	} else if err != nil {
		h.evmKeeper.Logger(ctx).Error("failed to parse memo", "error", err)
		return nil
	}

	// create a new cache context to ignore errors during
	// the execution of the callback
	cacheCtx, write := ctx.CacheContext()

	callback := hookData.AsyncCallback
	if allowed, err := h.checkACL(im, cacheCtx, callback.ContractAddress); err != nil {
		h.evmKeeper.Logger(cacheCtx).Error("failed to check ACL", "error", err)
		return nil
	} else if !allowed {
		h.evmKeeper.Logger(cacheCtx).Error("failed to check ACL", "not allowed")
		return nil
	}

	inputBz, err := h.asyncCallbackABI.Pack(functionNameTimeout, callback.Id)
	if err != nil {
		h.evmKeeper.Logger(cacheCtx).Error("failed to pack input", "error", err)
		return nil
	}

	_, err = h.execMsg(cacheCtx, &evmtypes.MsgCall{
		Sender:       data.Sender,
		ContractAddr: callback.ContractAddress,
		Input:        hexutil.Encode(inputBz),
	})
	if err != nil {
		h.evmKeeper.Logger(cacheCtx).Error("failed to execute callback", "error", err)
		return nil
	}

	// write the cache context only if the callback execution was successful
	write()

	return nil
}
