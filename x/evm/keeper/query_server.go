package keeper

import (
	"context"
	"encoding/hex"
	"fmt"
	"strings"

	errorsmod "cosmossdk.io/errors"
	storetypes "cosmossdk.io/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ethereum/go-ethereum/common"
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
	sdkCtx = sdkCtx.WithGasMeter(storetypes.NewGasMeter(qs.config.ContractQueryGasLimit))

	sender, err := qs.ac.StringToBytes(req.Sender)
	if err != nil {
		return nil, err
	}

	contractAddr, err := types.ContractAddressFromString(qs.ac, req.ContractAddr)
	if err != nil {
		return nil, err
	}

	inputBz, err := hex.DecodeString(req.Input)
	if err != nil {
		return nil, err
	}

	var tracer *tracing.Hooks
	tracerOutput := new(strings.Builder)
	if req.WithTrace {
		tracer = logger.NewJSONLogger(&logger.Config{
			EnableMemory:     qs.config.TracingEnableMemory,
			DisableStack:     !qs.config.TracingEnableStack,
			DisableStorage:   !qs.config.TracingEnableStorage,
			EnableReturnData: qs.config.TracingEnableReturnData,
		}, tracerOutput)
	}

	// use cache context to rollback writes
	sdkCtx, _ = sdkCtx.CacheContext()
	caller := common.BytesToAddress(sender)
	retBz, logs, err := qs.EVMCallWithTracer(sdkCtx, caller, contractAddr, inputBz, tracer)
	if err != nil {
		return nil, err
	}

	return &types.QueryCallResponse{
		Response:    common.Bytes2Hex(retBz),
		UsedGas:     sdkCtx.GasMeter().GasConsumedToLimit(),
		Logs:        logs,
		TraceOutput: tracerOutput.String(),
	}, nil
}

// Code implements types.QueryServer.
func (qs *queryServerImpl) Code(ctx context.Context, req *types.QueryCodeRequest) (*types.QueryCodeResponse, error) {
	stateDB, err := qs.newStateDB(ctx)
	if err != nil {
		return nil, err
	}

	contractAddr, err := types.ContractAddressFromString(qs.ac, req.ContractAddr)
	if err != nil {
		return nil, err
	}

	codeBz := stateDB.GetCode(common.Address(contractAddr.Bytes()))
	return &types.QueryCodeResponse{
		Code: codeBz,
	}, nil
}

// State implements types.QueryServer.
func (qs *queryServerImpl) State(ctx context.Context, req *types.QueryStateRequest) (*types.QueryStateResponse, error) {
	stateDB, err := qs.newStateDB(ctx)
	if err != nil {
		return nil, err
	}

	contractAddr, err := types.ContractAddressFromString(qs.ac, req.ContractAddr)
	if err != nil {
		return nil, err
	}

	state := stateDB.GetState(common.Address(contractAddr.Bytes()), common.HexToHash(req.Key))
	return &types.QueryStateResponse{
		Value: state.Hex(),
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
