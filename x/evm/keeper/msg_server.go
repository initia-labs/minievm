package keeper

import (
	"context"

	"github.com/ethereum/go-ethereum/common"

	sdk "github.com/cosmos/cosmos-sdk/types"
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

	codeBz := msg.Code

	// argument validation
	if len(sender) != common.AddressLength {
		return nil, types.ErrInvalidAddressLength
	}
	if len(codeBz) == 0 {
		return nil, sdkerrors.ErrInvalidRequest.Wrap("empty code bytes")
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
	retBz, contractAddr, err := ms.EVMCreate(ctx, sender, codeBz)
	if err != nil {
		return nil, types.ErrEVMCallFailed.Wrap(err.Error())
	}

	// emit action events
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	sdkCtx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeCall,
		sdk.NewAttribute(types.AttributeKeyContract, contractAddr.Hex()),
		sdk.NewAttribute(types.AttributeKeyRet, common.Bytes2Hex(retBz)),
	))

	return &types.MsgCreateResponse{
		Result:       retBz,
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

	inputBz := msg.Input

	// argument validation
	if len(sender) != common.AddressLength {
		return nil, types.ErrInvalidAddressLength
	}
	if len(contractAddr) != common.AddressLength {
		return nil, types.ErrInvalidAddressLength
	}

	retBz, logs, err := ms.EVMCall(ctx, sender, contractAddr, inputBz)
	if err != nil {
		return nil, types.ErrEVMCreateFailed.Wrap(err.Error())
	}

	// emit action events
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	sdkCtx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeCall,
		sdk.NewAttribute(types.AttributeKeyContract, contractAddr.Hex()),
		sdk.NewAttribute(types.AttributeKeyRet, common.Bytes2Hex(retBz)),
	))

	// emit logs events
	for _, log := range logs {
		dataInHex := common.Bytes2Hex(log.Data)
		for _, topic := range log.Topics {
			sdkCtx.EventManager().EmitEvent(sdk.NewEvent(
				types.EventTypeLog,
				sdk.NewAttribute(types.AttributeKeyTopic, topic),
				sdk.NewAttribute(types.AttributeKeyRet, dataInHex),
			))
		}
	}

	return &types.MsgCallResponse{Result: retBz, Logs: logs}, nil
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
