package evm_hooks

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	transfertypes "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"
	channeltypes "github.com/cosmos/ibc-go/v8/modules/core/04-channel/types"
	"github.com/ethereum/go-ethereum/common/hexutil"

	ibchooks "github.com/initia-labs/initia/x/ibc-hooks"

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
		return err
	}

	callback := hookData.AsyncCallback
	if allowed, err := h.checkACL(im, ctx, callback.ContractAddress); err != nil {
		return err
	} else if !allowed {
		// just return nil here to avoid packet stuck due to hook acl.
		return nil
	}

	inputBz, err := h.asyncCallbackABI.Pack(functionNameTimeout, callback.Id)
	if err != nil {
		return err
	}

	_, err = h.execMsg(ctx, &evmtypes.MsgCall{
		Sender:       data.Sender,
		ContractAddr: callback.ContractAddress,
		Input:        hexutil.Encode(inputBz),
	})
	if err != nil {
		return err
	}

	return nil
}

// func (h EVMHooks) onTimeoutIcs721Packet(
// 	ctx sdk.Context,
// 	im ibchooks.IBCMiddleware,
// 	packet channeltypes.Packet,
// 	relayer sdk.AccAddress,
// 	data nfttransfertypes.NonFungibleTokenPacketData,
// ) error {
// 	if err := im.App.OnTimeoutPacket(ctx, packet, relayer); err != nil {
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

// 	inputBz, err := h.asyncCallbackABI.Pack(functionNameTimeout, callback.Id)
// 	if err != nil {
// 		return err
// 	}

// 	_, err = h.execMsg(ctx, &evmtypes.MsgCall{
// 		Sender:       data.Sender,
// 		ContractAddr: callback.ContractAddress,
// 		Input:        hexutil.Encode(inputBz),
// 	})
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
