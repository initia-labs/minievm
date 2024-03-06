package hook

import (
	"context"
	"encoding/json"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"

	evmkeeper "github.com/initia-labs/minievm/x/evm/keeper"
	evmtypes "github.com/initia-labs/minievm/x/evm/types"
)

// bridge hook implementation for evm
type EVMBridgeHook struct {
	evmKeeper *evmkeeper.Keeper
}

func NewEVMBridgeHook(evmKeeper *evmkeeper.Keeper) EVMBridgeHook {
	return EVMBridgeHook{evmKeeper}
}

func (mbh EVMBridgeHook) Hook(ctx context.Context, sender sdk.AccAddress, msgBytes []byte) error {
	msg := evmtypes.MsgCall{}
	decoder := json.NewDecoder(strings.NewReader(string(msgBytes)))
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&msg)
	if err != nil {
		return err
	}

	ms := evmkeeper.NewMsgServerImpl(mbh.evmKeeper)
	_, err = ms.Call(ctx, &msg)

	return err
}
