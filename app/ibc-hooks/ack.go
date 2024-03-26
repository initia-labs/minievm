package evm_hooks

import (
	"encoding/hex"

	sdk "github.com/cosmos/cosmos-sdk/types"
	transfertypes "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"
	channeltypes "github.com/cosmos/ibc-go/v8/modules/core/04-channel/types"

	ibchooks "github.com/initia-labs/initia/x/ibc-hooks"
	evmtypes "github.com/initia-labs/minievm/x/evm/types"
)

func (h EVMHooks) onAckIcs20Packet(
	ctx sdk.Context,
	im ibchooks.IBCMiddleware,
	packet channeltypes.Packet,
	acknowledgement []byte,
	relayer sdk.AccAddress,
	data transfertypes.FungibleTokenPacketData,
) error {
	if err := im.App.OnAcknowledgementPacket(ctx, packet, acknowledgement, relayer); err != nil {
		return err
	}

	isEVMRouted, hookData, err := validateAndParseMemo(data.GetMemo())
	if !isEVMRouted || hookData.AsyncCallback == nil {
		return nil
	} else if err != nil {
		return err
	}

	callback := hookData.AsyncCallback
	if allowed, err := h.checkACL(im, ctx, callback.ContractAddress); err != nil {
		return err
	} else if !allowed {
		return nil
	}

	inputBz, err := h.asyncCallbackABI.Pack(functionNameAck, callback.Id, !isAckError(acknowledgement))
	if err != nil {
		return err
	}

	_, err = h.execMsg(ctx, &evmtypes.MsgCall{
		Sender:       data.Sender,
		ContractAddr: callback.ContractAddress,
		Input:        hex.EncodeToString(inputBz),
	})
	if err != nil {
		return err
	}

	return nil
}

// func (h EVMHooks) onAckIcs721Packet(
// 	ctx sdk.Context,
// 	im ibchooks.IBCMiddleware,
// 	packet channeltypes.Packet,
// 	acknowledgement []byte,
// 	relayer sdk.AccAddress,
// 	data nfttransfertypes.NonFungibleTokenPacketData,
// ) error {
// 	if err := im.App.OnAcknowledgementPacket(ctx, packet, acknowledgement, relayer); err != nil {
// 		return err
// 	}

// 	isEVMRouted, hookData, err := validateAndParseMemo(data.GetMemo())
// 	if !isEVMRouted || hookData.AsyncCallback == nil {
// 		return nil
// 	} else if err != nil {
// 		return err
// 	}

// 	callback := hookData.AsyncCallback
// 	if allowed, err := h.checkACL(im, ctx, callback.ContractAddress); err != nil {
// 		return err
// 	} else if !allowed {
// 		return nil
// 	}

// 	inputBz, err := h.asyncCallbackABI.Pack(functionNameAck, callback.Id, !isAckError(acknowledgement))
// 	if err != nil {
// 		return err
// 	}

// 	_, err = h.execMsg(ctx, &evmtypes.MsgCall{
// 		Sender:       data.Sender,
// 		ContractAddr: callback.ContractAddress,
// 		Input:        hex.EncodeToString(inputBz),
// 	})
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
