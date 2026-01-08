package evm_hooks

import (
	"encoding/json"
	"fmt"
	"strings"

	"cosmossdk.io/errors"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"

	transfertypes "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"
	channeltypes "github.com/cosmos/ibc-go/v8/modules/core/04-channel/types"

	nfttransfertypes "github.com/initia-labs/initia/x/ibc/nft-transfer/types"

	evmtypes "github.com/initia-labs/minievm/x/evm/types"
)

const senderPrefix = "ibc-evm-hook-intermediary"

// deriveIntermediateSender compute intermediate sender address
// Bech32(Hash(Hash("ibc-hook-intermediary") + channelID/sender)[12:])
//
// @dev: use 20bytes of address as intermediate sender address
//
// TODO - make this as module account to check address collision
func DeriveIntermediateSender(channel, originalSender string) string {
	senderStr := fmt.Sprintf("%s/%s", channel, originalSender)
	senderAddr := sdk.AccAddress(address.Hash(senderPrefix, []byte(senderStr)))
	return senderAddr.String()
}

func isIcs20Packet(packetData []byte) (isIcs20 bool, ics20data transfertypes.FungibleTokenPacketData) {
	var data transfertypes.FungibleTokenPacketData
	decoder := json.NewDecoder(strings.NewReader(string(packetData)))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&data); err != nil {
		return false, data
	}
	return true, data
}

func isIcs721Packet(packetData []byte) (isIcs721 bool, ics721data nfttransfertypes.NonFungibleTokenPacketData) {
	if data, err := nfttransfertypes.DecodePacketData(packetData); err != nil {
		return false, data
	} else {
		return true, data
	}
}

func parseHookData(memo string) (*HookData, bool, error) {
	if len(memo) == 0 {
		return nil, false, nil
	}

	var memoMap map[string]json.RawMessage
	if err := json.Unmarshal([]byte(memo), &memoMap); err != nil {
		return nil, false, nil
	}

	raw, ok := memoMap[evmHookMemoKey]
	if !ok {
		return nil, false, nil
	}

	var hookData HookData
	decoder := json.NewDecoder(strings.NewReader(string(raw)))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&hookData); err != nil {
		return nil, true, errors.Wrap(channeltypes.ErrInvalidPacket, err.Error())
	}

	return &hookData, true, nil
}

func validateReceiver(msg *evmtypes.MsgCall, receiver string) error {
	if receiver != msg.ContractAddr {
		return errors.Wrapf(channeltypes.ErrInvalidPacket, "receiver is not properly set")
	}

	return nil
}

// newEmitErrorAcknowledgement creates a new error acknowledgement after having emitted an event with the
// details of the error.
func newEmitErrorAcknowledgement(err error) channeltypes.Acknowledgement {
	return channeltypes.Acknowledgement{
		Response: &channeltypes.Acknowledgement_Error{
			Error: fmt.Sprintf("ibc evm hook error: %s", err.Error()),
		},
	}
}

// isAckError checks an IBC acknowledgement to see if it's an error.
func isAckError(appCodec codec.Codec, acknowledgement []byte) bool {
	var ack channeltypes.Acknowledgement
	if err := appCodec.UnmarshalJSON(acknowledgement, &ack); err == nil && !ack.Success() {
		return true
	}

	return false
}

func LocalDenom(packet channeltypes.Packet, denom string) string {
	if transfertypes.ReceiverChainIsSource(packet.GetSourcePort(), packet.GetSourceChannel(), denom) {
		voucherPrefix := transfertypes.GetDenomPrefix(packet.GetSourcePort(), packet.GetSourceChannel())
		unprefixedDenom := denom[len(voucherPrefix):]

		// coin denomination used in sending from the escrow address
		denom := unprefixedDenom

		// The denomination used to send the coins is either the native denom or the hash of the path
		// if the denomination is not native.
		denomTrace := transfertypes.ParseDenomTrace(unprefixedDenom)
		if !denomTrace.IsNativeDenom() {
			denom = denomTrace.IBCDenom()
		}

		return denom
	}

	// since SendPacket did not prefix the denomination, we must prefix denomination here
	sourcePrefix := transfertypes.GetDenomPrefix(packet.GetDestPort(), packet.GetDestChannel())
	// NOTE: sourcePrefix contains the trailing "/"
	prefixedDenom := sourcePrefix + denom

	// construct the denomination trace from the full raw denomination
	denomTrace := transfertypes.ParseDenomTrace(prefixedDenom)

	voucherDenom := denomTrace.IBCDenom()
	return voucherDenom
}

func LocalClassId(packet channeltypes.Packet, classId string) string {
	if nfttransfertypes.ReceiverChainIsSource(packet.GetSourcePort(), packet.GetSourceChannel(), classId) {
		// remove prefix added by sender chain
		voucherPrefix := nfttransfertypes.GetClassIdPrefix(packet.GetSourcePort(), packet.GetSourceChannel())
		unprefixedClassId := classId[len(voucherPrefix):]

		// token class id used in sending from the escrow address
		classId := unprefixedClassId

		// The class id used to send the coins is either the native classId or the hash of the path
		// if the class id is not native.
		classTrace := nfttransfertypes.ParseClassTrace(unprefixedClassId)
		if classTrace.Path != "" {
			classId = classTrace.IBCClassId()
		}

		return classId
	}

	// since SendPacket did not prefix the class id, we must prefix class id here
	sourcePrefix := nfttransfertypes.GetClassIdPrefix(packet.GetDestPort(), packet.GetDestChannel())
	// NOTE: sourcePrefix contains the trailing "/"
	prefixedClassId := sourcePrefix + classId

	// construct the class id trace from the full raw class id
	classTrace := nfttransfertypes.ParseClassTrace(prefixedClassId)

	voucherClassId := classTrace.IBCClassId()
	return voucherClassId
}
