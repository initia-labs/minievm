package keeper

import (
	"context"
	"encoding/hex"
	"strings"

	"github.com/ethereum/go-ethereum/common"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/initia-labs/minievm/x/evm/types"
)

type msgServerImpl struct {
	*Keeper
}

func NewMsgServerImpl(k *Keeper) types.MsgServer {
	return &msgServerImpl{k}
}

// Create implements types.MsgServer.
func (ms *msgServerImpl) Create(ctx context.Context, msg *types.MsgCreate) (*types.MsgCreateResponse, error) {
	sender, err := ms.ac.StringToBytes(msg.Sender)
	if err != nil {
		return nil, err
	}

	// argument validation
	caller, err := ms.convertToEVMAddress(ctx, sender)
	if err != nil {
		return nil, err
	}
	if len(msg.Code) == 0 {
		return nil, sdkerrors.ErrInvalidRequest.Wrap("empty code bytes")
	}
	codeBz, err := hex.DecodeString(strings.TrimPrefix(msg.Code, "0x"))
	if err != nil {
		return nil, types.ErrInvalidHexString.Wrap(err.Error())
	}

	// check the sender is allowed publisher
	params, err := ms.Params.Get(ctx)
	if err != nil {
		return nil, err
	}

	// assert deploy authorization
	if len(params.AllowedPublishers) != 0 {
		allowed := false
		for _, publisher := range params.AllowedPublishers {
			if msg.Sender == publisher {
				allowed = true

				break
			}
		}

		if !allowed {
			return nil, sdkerrors.ErrUnauthorized.Wrapf("`%s` is not allowed to deploy a contract", msg.Sender)
		}
	}

	// deploy a contract
	retBz, contractAddr, err := ms.EVMCreate(ctx, caller, codeBz)
	if err != nil {
		return nil, types.ErrEVMCallFailed.Wrap(err.Error())
	}

	retHex := common.Bytes2Hex(retBz)
	return &types.MsgCreateResponse{
		Result:       retHex,
		ContractAddr: contractAddr.Hex(),
	}, nil
}

// Call implements types.MsgServer.
func (ms *msgServerImpl) Call(ctx context.Context, msg *types.MsgCall) (*types.MsgCallResponse, error) {
	sender, err := ms.ac.StringToBytes(msg.Sender)
	if err != nil {
		return nil, err
	}

	contractAddr, err := types.ContractAddressFromString(ms.ac, msg.ContractAddr)
	if err != nil {
		return nil, err
	}

	// argument validation
	caller, err := ms.convertToEVMAddress(ctx, sender)
	if err != nil {
		return nil, err
	}
	inputBz, err := hex.DecodeString(strings.TrimPrefix(msg.Input, "0x"))
	if err != nil {
		return nil, types.ErrInvalidHexString.Wrap(err.Error())
	}

	retBz, logs, err := ms.EVMCall(ctx, caller, contractAddr, inputBz)
	if err != nil {
		return nil, types.ErrEVMCreateFailed.Wrap(err.Error())
	}

	retHex := common.Bytes2Hex(retBz)
	return &types.MsgCallResponse{Result: retHex, Logs: logs}, nil
}

// UpdateParams implements types.MsgServer.
func (ms *msgServerImpl) UpdateParams(ctx context.Context, msg *types.MsgUpdateParams) (*types.MsgUpdateParamsResponse, error) {
	// assert permission
	if ms.authority != msg.Authority {
		return nil, govtypes.ErrInvalidSigner.Wrapf("invalid authority; expected %s, got %s", ms.authority, msg.Authority)
	}

	// validate params
	if err := msg.Params.Validate(ms.ac); err != nil {
		return nil, err
	}

	// update params
	if err := ms.Params.Set(ctx, msg.Params); err != nil {
		return nil, err
	}

	return &types.MsgUpdateParamsResponse{}, nil
}
