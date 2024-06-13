package keeper

import (
	"context"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/holiman/uint256"

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
	caller, err := ms.convertToEVMAddress(ctx, sender, true)
	if err != nil {
		return nil, err
	}
	if len(msg.Code) == 0 {
		return nil, sdkerrors.ErrInvalidRequest.Wrap("empty code bytes")
	}
	codeBz, err := hexutil.Decode(msg.Code)
	if err != nil {
		return nil, types.ErrInvalidHexString.Wrap(err.Error())
	}
	value, overflow := uint256.FromBig(msg.Value.BigInt())
	if overflow {
		return nil, types.ErrInvalidValue.Wrap("value is out of range")
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
	retBz, contractAddr, _, err := ms.EVMCreate(ctx, caller, codeBz, value)
	if err != nil {
		return nil, types.ErrEVMCallFailed.Wrap(err.Error())
	}

	return &types.MsgCreateResponse{
		Result:       hexutil.Encode(retBz),
		ContractAddr: contractAddr.Hex(),
	}, nil
}

// Create2 implements types.MsgServer.
func (ms *msgServerImpl) Create2(ctx context.Context, msg *types.MsgCreate2) (*types.MsgCreate2Response, error) {
	sender, err := ms.ac.StringToBytes(msg.Sender)
	if err != nil {
		return nil, err
	}

	// argument validation
	caller, err := ms.convertToEVMAddress(ctx, sender, true)
	if err != nil {
		return nil, err
	}
	if len(msg.Code) == 0 {
		return nil, sdkerrors.ErrInvalidRequest.Wrap("empty code bytes")
	}
	codeBz, err := hexutil.Decode(msg.Code)
	if err != nil {
		return nil, types.ErrInvalidHexString.Wrap(err.Error())
	}
	value, overflow := uint256.FromBig(msg.Value.BigInt())
	if overflow {
		return nil, types.ErrInvalidValue.Wrap("value is out of range")
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
	retBz, contractAddr, _, err := ms.EVMCreate2(ctx, caller, codeBz, value, msg.Salt)
	if err != nil {
		return nil, types.ErrEVMCallFailed.Wrap(err.Error())
	}

	return &types.MsgCreate2Response{
		Result:       hexutil.Encode(retBz),
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
	caller, err := ms.convertToEVMAddress(ctx, sender, true)
	if err != nil {
		return nil, err
	}
	inputBz, err := hexutil.Decode(msg.Input)
	if err != nil {
		return nil, types.ErrInvalidHexString.Wrap(err.Error())
	}
	value, overflow := uint256.FromBig(msg.Value.BigInt())
	if overflow {
		return nil, types.ErrInvalidValue.Wrap("value is out of range")
	}

	retBz, logs, err := ms.EVMCall(ctx, caller, contractAddr, inputBz, value)
	if err != nil {
		return nil, types.ErrEVMCallFailed.Wrap(err.Error())
	}

	return &types.MsgCallResponse{Result: hexutil.Encode(retBz), Logs: logs}, nil
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
