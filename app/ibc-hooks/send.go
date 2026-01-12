package evm_hooks

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"

	capabilitytypes "github.com/cosmos/ibc-go/modules/capability/types"
	transfertypes "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"
	clienttypes "github.com/cosmos/ibc-go/v8/modules/core/02-client/types"

	ibchooks "github.com/initia-labs/initia/x/ibc-hooks"
	ibchookstypes "github.com/initia-labs/initia/x/ibc-hooks/types"
	nfttransfertypes "github.com/initia-labs/initia/x/ibc/nft-transfer/types"

	evmtypes "github.com/initia-labs/minievm/x/evm/types"
)

func (h EVMHooks) sendIcs20Packet(
	ctx sdk.Context,
	im ibchooks.ICS4Middleware,
	chanCap *capabilitytypes.Capability,
	sourcePort string,
	sourceChannel string,
	timeoutHeight clienttypes.Height,
	timeoutTimestamp uint64,
	ics20Data transfertypes.FungibleTokenPacketData,
) (uint64, error) {
	return h.handleSendPacket(ctx, im, chanCap, sourcePort, sourceChannel, timeoutHeight, timeoutTimestamp, ibchookstypes.ICSData{
		ICS20Data: &ics20Data,
	})
}

func (h EVMHooks) sendIcs721Packet(
	ctx sdk.Context,
	im ibchooks.ICS4Middleware,
	chanCap *capabilitytypes.Capability,
	sourcePort string,
	sourceChannel string,
	timeoutHeight clienttypes.Height,
	timeoutTimestamp uint64,
	ics721Data nfttransfertypes.NonFungibleTokenPacketData,
) (uint64, error) {
	return h.handleSendPacket(ctx, im, chanCap, sourcePort, sourceChannel, timeoutHeight, timeoutTimestamp, ibchookstypes.ICSData{
		ICS721Data: &ics721Data,
	})
}

func (h EVMHooks) handleSendPacket(
	ctx sdk.Context,
	im ibchooks.ICS4Middleware,
	chanCap *capabilitytypes.Capability,
	sourcePort string,
	sourceChannel string,
	timeoutHeight clienttypes.Height,
	timeoutTimestamp uint64,
	icsData ibchookstypes.ICSData,
) (uint64, error) {
	hookData, isEVMRouted, err := parseHookData(icsData.GetMemo())
	if err != nil {
		return 0, err
	}
	if !isEVMRouted || hookData == nil || hookData.AsyncCallback == nil {
		return im.ICS4Wrapper.SendPacket(ctx, chanCap, sourcePort, sourceChannel, timeoutHeight, timeoutTimestamp, icsData.GetBytes())
	}

	asyncCallback := hookData.AsyncCallback
	if _, err := evmtypes.ContractAddressFromString(h.ac, asyncCallback.ContractAddress); err != nil {
		return 0, err
	}

	var memoMap map[string]any
	// ignore error, it is already checked in parseHookData
	_ = json.Unmarshal([]byte(icsData.GetMemo()), &memoMap)
	if hookData.Message == nil {
		delete(memoMap, evmHookMemoKey)
	} else {
		hookData.AsyncCallback = nil
		bz, err := json.Marshal(hookData)
		if err != nil {
			return 0, err
		}
		memoMap[evmHookMemoKey] = json.RawMessage(bz)
	}
	bz, err := json.Marshal(memoMap)
	if err != nil {
		return 0, err
	}
	icsData.SetMemo(string(bz))

	sequence, err := im.ICS4Wrapper.SendPacket(ctx, chanCap, sourcePort, sourceChannel, timeoutHeight, timeoutTimestamp, icsData.GetBytes())
	if err != nil {
		return sequence, err
	}

	asyncCallbackBz, err := json.Marshal(asyncCallback)
	if err != nil {
		return sequence, err
	}
	if err := im.HooksKeeper.SetAsyncCallback(ctx, sourcePort, sourceChannel, sequence, asyncCallbackBz); err != nil {
		return sequence, err
	}

	return sequence, nil
}
