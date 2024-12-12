package keeper

import (
	"context"
	"errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/tracing"
	coretypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/holiman/uint256"

	"cosmossdk.io/collections"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	evmante "github.com/initia-labs/minievm/x/evm/ante"
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

	// handle cosmos<>evm different sequence increment logic
	err = ms.handleSequenceIncremented(ctx, sender, true)
	if err != nil {
		return nil, err
	}

	// argument validation
	caller, codeBz, value, accessList, err := ms.validateArguments(ctx, sender, msg.Code, msg.Value, msg.AccessList, true)
	if err != nil {
		return nil, err
	}

	// check the sender is allowed publisher
	err = ms.assertAllowedPublishers(ctx, msg.Sender)
	if err != nil {
		return nil, err
	}

	// deploy a contract
	retBz, contractAddr, logs, err := ms.EVMCreate(ctx, caller, codeBz, value, accessList)
	if err != nil {
		return nil, types.ErrEVMCallFailed.Wrap(err.Error())
	}

	return &types.MsgCreateResponse{
		Result:       hexutil.Encode(retBz),
		ContractAddr: contractAddr.Hex(),
		Logs:         logs,
	}, nil
}

// Create2 implements types.MsgServer.
func (ms *msgServerImpl) Create2(ctx context.Context, msg *types.MsgCreate2) (*types.MsgCreate2Response, error) {
	sender, err := ms.ac.StringToBytes(msg.Sender)
	if err != nil {
		return nil, err
	}

	// handle cosmos<>evm different sequence increment logic
	err = ms.handleSequenceIncremented(ctx, sender, true)
	if err != nil {
		return nil, err
	}

	// argument validation
	caller, codeBz, value, accessList, err := ms.validateArguments(ctx, sender, msg.Code, msg.Value, msg.AccessList, true)
	if err != nil {
		return nil, err
	}

	// check the sender is allowed publisher
	err = ms.assertAllowedPublishers(ctx, msg.Sender)
	if err != nil {
		return nil, err
	}

	// deploy a contract
	retBz, contractAddr, logs, err := ms.EVMCreate2(ctx, caller, codeBz, value, msg.Salt, accessList)
	if err != nil {
		return nil, types.ErrEVMCallFailed.Wrap(err.Error())
	}

	return &types.MsgCreate2Response{
		Result:       hexutil.Encode(retBz),
		ContractAddr: contractAddr.Hex(),
		Logs:         logs,
	}, nil
}

// Call implements types.MsgServer.
func (ms *msgServerImpl) Call(ctx context.Context, msg *types.MsgCall) (*types.MsgCallResponse, error) {
	sender, err := ms.ac.StringToBytes(msg.Sender)
	if err != nil {
		return nil, err
	}

	// handle cosmos<>evm different sequence increment logic
	err = ms.handleSequenceIncremented(ctx, sender, false)
	if err != nil {
		return nil, err
	}

	contractAddr, err := types.ContractAddressFromString(ms.ac, msg.ContractAddr)
	if err != nil {
		return nil, err
	}

	// argument validation
	caller, inputBz, value, accessList, err := ms.validateArguments(ctx, sender, msg.Input, msg.Value, msg.AccessList, false)
	if err != nil {
		return nil, err
	}

	retBz, logs, err := ms.EVMCall(ctx, caller, contractAddr, inputBz, value, accessList)
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

	// test fee denom has sudoMint and sudoBurn
	// for stateDB AddBalance and SubBalance.
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	cacheCtx, _ := sdkCtx.CacheContext()
	err := ms.testFeeDenom(cacheCtx, msg.Params)
	if err != nil {
		return nil, err
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

// testFeeDenom tests if the fee denom has sudoMint and sudoBurn.
func (ms *msgServerImpl) testFeeDenom(ctx context.Context, params types.Params) (err error) {
	_, err = types.DenomToContractAddr(ctx, ms, params.FeeDenom)
	if err != nil && errors.Is(err, collections.ErrNotFound) {
		return types.ErrInvalidFeeDenom.Wrap("fee denom is not found")
	} else if err != nil {
		return err
	}

	err = ms.Params.Set(ctx, params)
	if err != nil {
		return err
	}

	_, evm, err := ms.CreateEVM(ctx, types.StdAddress, nil)
	if err != nil {
		return err
	}

	defer func() {
		if r := recover(); r != nil {
			err = types.ErrInvalidFeeDenom.Wrap("failed to conduct sudoMint and sudoBurn")
		}
	}()

	evm.StateDB.AddBalance(types.StdAddress, uint256.NewInt(1), tracing.BalanceChangeUnspecified)
	evm.StateDB.SubBalance(types.StdAddress, uint256.NewInt(1), tracing.BalanceChangeUnspecified)

	return nil
}

// In the Cosmos SDK, the sequence number is incremented in the ante handler.
// In the EVM, the sequence number is incremented during the execution of create and create2 messages.
//
// If the sequence number is already incremented in the ante handler and the message is create, decrement the sequence number to prevent double incrementing.
// If the sequence number is not incremented in the ante handler and the message is call, increment the sequence number to ensure proper sequencing.
func (k *msgServerImpl) handleSequenceIncremented(ctx context.Context, sender sdk.AccAddress, isCreate bool) error {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	if sdkCtx.Value(evmante.ContextKeySequenceIncremented) == nil {
		return nil
	}

	incremented := sdkCtx.Value(evmante.ContextKeySequenceIncremented).(*bool)
	if isCreate && *incremented {
		// if the sequence is already incremented, decrement it to prevent double incrementing the sequence number at create.
		acc := k.accountKeeper.GetAccount(ctx, sender)
		if err := acc.SetSequence(acc.GetSequence() - 1); err != nil {
			return err
		}

		k.accountKeeper.SetAccount(ctx, acc)
	} else if !isCreate && !*incremented {
		// if the sequence is not incremented and the message is call, increment the sequence number.
		acc := k.accountKeeper.GetAccount(ctx, sender)
		if err := acc.SetSequence(acc.GetSequence() + 1); err != nil {
			return err
		}

		k.accountKeeper.SetAccount(ctx, acc)
	}

	// set the flag to false
	*incremented = false

	return nil
}

// validateArguments validates the arguments of create, create2, and call messages.
func (ms *msgServerImpl) validateArguments(
	ctx context.Context, sender []byte, data string,
	value math.Int, accessList []types.AccessTuple, isCreate bool,
) (common.Address, []byte, *uint256.Int, coretypes.AccessList, error) {
	caller, err := ms.convertToEVMAddress(ctx, sender, true)
	if err != nil {
		return common.Address{}, nil, nil, nil, err
	}
	if isCreate && len(data) == 0 {
		return common.Address{}, nil, nil, nil, sdkerrors.ErrInvalidRequest.Wrap("empty code bytes")
	}
	dataBz, err := hexutil.Decode(data)
	if err != nil {
		return common.Address{}, nil, nil, nil, types.ErrInvalidHexString.Wrap(err.Error())
	}
	val, overflow := uint256.FromBig(value.BigInt())
	if overflow {
		return common.Address{}, nil, nil, nil, types.ErrInvalidValue.Wrap("value is out of range")
	}

	return caller, dataBz, val, types.ConvertCosmosAccessListToEth(accessList), nil
}

// assertAllowedPublishers asserts the sender is allowed to deploy a contract.
func (ms *msgServerImpl) assertAllowedPublishers(ctx context.Context, sender string) error {
	params, err := ms.Params.Get(ctx)
	if err != nil {
		return err
	}

	// assert deploy authorization
	if len(params.AllowedPublishers) != 0 {
		allowed := false
		for _, publisher := range params.AllowedPublishers {
			if sender == publisher {
				allowed = true

				break
			}
		}

		if !allowed {
			return sdkerrors.ErrUnauthorized.Wrapf("`%s` is not allowed to deploy a contract", sender)
		}
	}

	return nil
}
