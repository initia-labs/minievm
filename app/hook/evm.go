package hook

import (
	"context"
	"encoding/json"
	"strings"

	"cosmossdk.io/core/address"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	evmkeeper "github.com/initia-labs/minievm/x/evm/keeper"
	evmtypes "github.com/initia-labs/minievm/x/evm/types"
)

// bridge hook implementation for evm
type EVMBridgeHook struct {
	ac        address.Codec
	evmKeeper *evmkeeper.Keeper
}

func NewEVMBridgeHook(ac address.Codec, evmKeeper *evmkeeper.Keeper) EVMBridgeHook {
	return EVMBridgeHook{ac, evmKeeper}
}

func (mbh EVMBridgeHook) Hook(ctx context.Context, sender sdk.AccAddress, msgBytes []byte) error {
	var msg evmtypes.MsgCall
	decoder := json.NewDecoder(strings.NewReader(string(msgBytes)))
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&msg)
	if err != nil {
		return err
	}

	senderAddr, err := mbh.ac.StringToBytes(msg.Sender)
	if err != nil {
		return err
	} else if !sender.Equals(sdk.AccAddress(senderAddr)) {
		return sdkerrors.ErrUnauthorized
	}

	ms := evmkeeper.NewMsgServerImpl(mbh.evmKeeper)
	_, err = ms.Call(ctx, &msg)

	return err
}
