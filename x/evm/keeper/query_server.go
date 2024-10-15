package keeper

import (
	"context"
	"fmt"
	"strings"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/holiman/uint256"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/tracing"
	"github.com/ethereum/go-ethereum/eth/tracers/logger"

	"github.com/initia-labs/minievm/x/evm/types"
)

type queryServerImpl struct {
	*Keeper
}

func NewQueryServer(k *Keeper) types.QueryServer {
	return &queryServerImpl{k}
}

// Call implements types.QueryServer.
func (qs *queryServerImpl) Call(ctx context.Context, req *types.QueryCallRequest) (res *types.QueryCallResponse, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errorsmod.Wrap(types.ErrEVMCallFailed, fmt.Sprintf("vm panic: %v", r))
		}
	}()

	sdkCtx := sdk.UnwrapSDKContext(ctx)

	caller := common.Address{}
	if req.Sender != "" {
		senderBz, err := qs.ac.StringToBytes(req.Sender)
		if err != nil {
			return nil, err
		}

		caller = common.BytesToAddress(senderBz)
	}

	contractAddr := common.Address{}
	if req.ContractAddr != "" {
		contractAddr, err = types.ContractAddressFromString(qs.ac, req.ContractAddr)
		if err != nil {
			return nil, err
		}
	}

	inputBz, err := hexutil.Decode(req.Input)
	if err != nil {
		return nil, err
	}

	value, overflow := uint256.FromBig(req.Value.BigInt())
	if overflow {
		return nil, types.ErrInvalidValue.Wrap("value is out of range")
	}

	var tracer *tracing.Hooks
	tracerOutput := new(strings.Builder)
	if req.TraceOptions != nil {
		tracer = logger.NewJSONLogger(&logger.Config{
			EnableMemory:     req.TraceOptions.WithMemory,
			DisableStack:     !req.TraceOptions.WithStack,
			DisableStorage:   !req.TraceOptions.WithStorage,
			EnableReturnData: req.TraceOptions.WithReturnData,
		}, tracerOutput)
	}

	// use cache context to rollback writes
	sdkCtx, _ = sdkCtx.CacheContext()

	var retBz []byte
	var logs []types.Log
	if contractAddr == (common.Address{}) {
		// if contract address is not provided, then it's a contract creation
		retBz, _, logs, err = qs.EVMCreateWithTracer(sdkCtx, caller, inputBz, value, nil, nil, tracer)
	} else {
		retBz, logs, err = qs.EVMCallWithTracer(sdkCtx, caller, contractAddr, inputBz, value, nil, tracer)

	}

	if err != nil {
		return &types.QueryCallResponse{
			Error:       err.Error(),
			UsedGas:     sdkCtx.GasMeter().GasConsumedToLimit(),
			TraceOutput: tracerOutput.String(),
		}, nil
	}

	return &types.QueryCallResponse{
		Response:    hexutil.Encode(retBz),
		UsedGas:     sdkCtx.GasMeter().GasConsumedToLimit(),
		Logs:        logs,
		TraceOutput: tracerOutput.String(),
	}, nil
}

// Code implements types.QueryServer.
func (qs *queryServerImpl) Code(ctx context.Context, req *types.QueryCodeRequest) (*types.QueryCodeResponse, error) {
	stateDB, err := qs.NewStateDB(ctx, nil, types.Fee{})
	if err != nil {
		return nil, err
	}

	contractAddr, err := types.ContractAddressFromString(qs.ac, req.ContractAddr)
	if err != nil {
		return nil, err
	}

	codeBz := stateDB.GetCode(contractAddr)
	return &types.QueryCodeResponse{
		Code: hexutil.Encode(codeBz),
	}, nil
}

// State implements types.QueryServer.
func (qs *queryServerImpl) State(ctx context.Context, req *types.QueryStateRequest) (*types.QueryStateResponse, error) {
	stateDB, err := qs.NewStateDB(ctx, nil, types.Fee{})
	if err != nil {
		return nil, err
	}

	contractAddr, err := types.ContractAddressFromString(qs.ac, req.ContractAddr)
	if err != nil {
		return nil, err
	}

	keyBz, err := hexutil.Decode(req.Key)
	if err != nil {
		return nil, err
	}

	state := stateDB.GetState(contractAddr, common.BytesToHash(keyBz))
	return &types.QueryStateResponse{
		Value: state.Hex(),
	}, nil
}

// ERC20Factory implements types.QueryServer.
func (qs *queryServerImpl) ERC20Factory(ctx context.Context, req *types.QueryERC20FactoryRequest) (*types.QueryERC20FactoryResponse, error) {
	factoryAddr, err := qs.Keeper.GetERC20FactoryAddr(ctx)
	if err != nil {
		return nil, err
	}

	return &types.QueryERC20FactoryResponse{
		Address: factoryAddr.Hex(),
	}, nil
}

// ContractAddrByDenom implements types.QueryServer.
func (qs *queryServerImpl) ContractAddrByDenom(ctx context.Context, req *types.QueryContractAddrByDenomRequest) (*types.QueryContractAddrByDenomResponse, error) {
	contractAddr, err := types.DenomToContractAddr(ctx, qs, req.Denom)
	if err != nil {
		return nil, err
	}

	return &types.QueryContractAddrByDenomResponse{
		Address: contractAddr.Hex(),
	}, nil
}

// Denom implements types.QueryServer.
func (qs *queryServerImpl) Denom(ctx context.Context, req *types.QueryDenomRequest) (*types.QueryDenomResponse, error) {
	addr, err := types.ContractAddressFromString(qs.ac, req.ContractAddr)
	if err != nil {
		return nil, err
	}

	denom, err := types.ContractAddrToDenom(ctx, qs, addr)
	if err != nil {
		return nil, err
	}

	return &types.QueryDenomResponse{Denom: denom}, nil
}

// Params implements types.QueryServer.
func (qs *queryServerImpl) Params(ctx context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	params, err := qs.Keeper.Params.Get(ctx)
	if err != nil {
		return nil, err
	}

	return &types.QueryParamsResponse{Params: params}, nil
}
