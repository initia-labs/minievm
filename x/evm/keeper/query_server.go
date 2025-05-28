package keeper

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/holiman/uint256"

	"cosmossdk.io/collections"
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/tracing"
	coretypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
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
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}

		caller = common.BytesToAddress(senderBz)
	}

	contractAddr := common.Address{}
	if req.ContractAddr != "" {
		contractAddr, err = types.ContractAddressFromString(qs.ac, req.ContractAddr)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}

	inputBz, err := hexutil.Decode(req.Input)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	value, overflow := uint256.FromBig(req.Value.BigInt())
	if overflow {
		return nil, status.Error(codes.InvalidArgument, "value is out of range")
	}

	list := types.ConvertCosmosAccessListToEth(req.AccessList)

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

	// set tracer to context
	timeoutRevert := false
	if tracer != nil {
		evmPointer := new(*vm.EVM)
		deadlineCtx, cancel := context.WithTimeout(sdkCtx, qs.config.TracerTimeout)
		go func() {
			<-deadlineCtx.Done()
			if errors.Is(deadlineCtx.Err(), context.DeadlineExceeded) {
				// Stop evm execution. Note cancellation is not necessarily immediate.
				if *evmPointer != nil {
					(*evmPointer).Cancel()
				}
				timeoutRevert = true
			}
		}()
		defer cancel()

		// create evm to create tracing
		_, evm, _, err := qs.CreateEVM(sdkCtx, caller)
		if err != nil {
			return nil, err
		}

		tracing := types.NewTracing(evm, tracer)
		sdkCtx = sdkCtx.WithValue(types.CONTEXT_KEY_TRACING, tracing)
		sdkCtx = sdkCtx.WithValue(types.CONTEXT_KEY_TRACE_EVM, evmPointer)

		// execute OnTxStart and dummy OnEnter
		gasLimit := qs.computeGasLimit(sdkCtx)
		if tracer.OnTxStart != nil {
			tracer.OnTxStart(tracing.VMContext(), types.TracingTx(gasLimit), caller)
		}
		if tracer.OnEnter != nil {
			tracer.OnEnter(0, byte(vm.CALL), types.NullAddress, types.NullAddress, []byte{}, gasLimit, nil)
		}
	}

	var retBz []byte
	var logs []types.Log
	if contractAddr == (common.Address{}) {
		// if contract address is not provided, then it's a contract creation
		retBz, _, logs, err = qs.EVMCreate(sdkCtx, caller, inputBz, value, list)
	} else {
		retBz, logs, err = qs.EVMCall(sdkCtx, caller, contractAddr, inputBz, value, list)
	}

	gasUsed := sdkCtx.GasMeter().GasConsumedToLimit()

	// execute dummy OnExit and OnTxEnd
	if tracer != nil {
		if tracer.OnExit != nil {
			if revertErr, ok := err.(*types.RevertError); ok {
				tracer.OnExit(0, revertErr.Ret(), gasUsed, vm.ErrExecutionReverted, true)
			} else {
				tracer.OnExit(0, nil, gasUsed, err, false)
			}
		}
		if tracer.OnTxEnd != nil {
			tracer.OnTxEnd(&coretypes.Receipt{GasUsed: gasUsed}, err)
		}
	}

	if timeoutRevert {
		return &types.QueryCallResponse{
			Error:       "execution timeout",
			UsedGas:     gasUsed,
			TraceOutput: tracerOutput.String(),
		}, nil
	} else if err != nil {
		return &types.QueryCallResponse{
			Error:       err.Error(),
			UsedGas:     gasUsed,
			TraceOutput: tracerOutput.String(),
		}, nil
	}

	return &types.QueryCallResponse{
		Response:    hexutil.Encode(retBz),
		UsedGas:     gasUsed,
		Logs:        logs,
		TraceOutput: tracerOutput.String(),
	}, nil
}

// Code implements types.QueryServer.
func (qs *queryServerImpl) Code(ctx context.Context, req *types.QueryCodeRequest) (*types.QueryCodeResponse, error) {
	if len(req.ContractAddr) == 0 {
		return nil, status.Error(codes.InvalidArgument, "empty contract address")
	}

	stateDB, err := qs.NewStateDB(ctx, nil, types.Fee{})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	contractAddr, err := types.ContractAddressFromString(qs.ac, req.ContractAddr)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	codeBz := stateDB.GetCode(contractAddr)
	return &types.QueryCodeResponse{
		Code: hexutil.Encode(codeBz),
	}, nil
}

// State implements types.QueryServer.
func (qs *queryServerImpl) State(ctx context.Context, req *types.QueryStateRequest) (*types.QueryStateResponse, error) {
	if len(req.ContractAddr) == 0 {
		return nil, status.Error(codes.InvalidArgument, "empty contract address")
	}

	stateDB, err := qs.NewStateDB(ctx, nil, types.Fee{})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	contractAddr, err := types.ContractAddressFromString(qs.ac, req.ContractAddr)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	keyBz, err := hexutil.Decode(req.Key)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	state := stateDB.GetState(contractAddr, common.BytesToHash(keyBz))
	return &types.QueryStateResponse{
		Value: state.Hex(),
	}, nil
}

// ERC20Factory implements types.QueryServer.
func (qs *queryServerImpl) ERC20Factory(ctx context.Context, req *types.QueryERC20FactoryRequest) (*types.QueryERC20FactoryResponse, error) {
	factoryAddr, err := qs.Keeper.GetERC20FactoryAddr(ctx)
	if err != nil && errors.Is(err, collections.ErrNotFound) {
		return nil, status.Error(codes.NotFound, err.Error())
	} else if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryERC20FactoryResponse{
		Address: factoryAddr.Hex(),
	}, nil
}

// ERC20Wrapper implements types.QueryServer.
func (qs *queryServerImpl) ERC20Wrapper(ctx context.Context, req *types.QueryERC20WrapperRequest) (*types.QueryERC20WrapperResponse, error) {
	wrapper, err := qs.Keeper.GetERC20WrapperAddr(ctx)
	if err != nil && errors.Is(err, collections.ErrNotFound) {
		return nil, status.Error(codes.NotFound, err.Error())
	} else if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryERC20WrapperResponse{
		Address: wrapper.Hex(),
	}, nil
}

// ConnectOracle implements types.QueryServer.
func (qs *queryServerImpl) ConnectOracle(ctx context.Context, req *types.QueryConnectOracleRequest) (*types.QueryConnectOracleResponse, error) {
	oracle, err := qs.Keeper.GetConnectOracleAddr(ctx)
	if err != nil && errors.Is(err, collections.ErrNotFound) {
		return nil, status.Error(codes.NotFound, err.Error())
	} else if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryConnectOracleResponse{
		Address: oracle.Hex(),
	}, nil
}

// ContractAddrByDenom implements types.QueryServer.
func (qs *queryServerImpl) ContractAddrByDenom(ctx context.Context, req *types.QueryContractAddrByDenomRequest) (*types.QueryContractAddrByDenomResponse, error) {
	if len(req.Denom) == 0 {
		return nil, status.Error(codes.InvalidArgument, "empty denom")
	}

	contractAddr, err := types.DenomToContractAddr(ctx, qs, req.Denom)
	if err != nil && errors.Is(err, collections.ErrNotFound) {
		return nil, status.Error(codes.NotFound, err.Error())
	} else if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryContractAddrByDenomResponse{
		Address: contractAddr.Hex(),
	}, nil
}

// Denom implements types.QueryServer.
func (qs *queryServerImpl) Denom(ctx context.Context, req *types.QueryDenomRequest) (*types.QueryDenomResponse, error) {
	if len(req.ContractAddr) == 0 {
		return nil, status.Error(codes.InvalidArgument, "empty contract address")
	}

	addr, err := types.ContractAddressFromString(qs.ac, req.ContractAddr)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	denom, err := types.ContractAddrToDenom(ctx, qs, addr)
	if err != nil && errors.Is(err, collections.ErrNotFound) {
		return nil, status.Error(codes.NotFound, err.Error())
	} else if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryDenomResponse{Denom: denom}, nil
}

// ERC721ClassIdByContractAddr implements types.QueryServer.
func (qs *queryServerImpl) ERC721ClassIdByContractAddr(ctx context.Context, req *types.QueryERC721ClassIdByContractAddrRequest) (*types.QueryERC721ClassIdByContractAddrResponse, error) {
	if len(req.ContractAddr) == 0 {
		return nil, status.Error(codes.InvalidArgument, "empty contract address")
	}

	contractAddr, err := types.ContractAddressFromString(qs.ac, req.ContractAddr)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	classId, err := qs.Keeper.GetClassIdByContractAddr(ctx, contractAddr)
	if err != nil && errors.Is(err, collections.ErrNotFound) {
		return nil, status.Error(codes.NotFound, err.Error())
	} else if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryERC721ClassIdByContractAddrResponse{ClassId: classId}, nil
}

// ERC721OriginTokenInfos implements types.QueryServer.
func (qs *queryServerImpl) ERC721OriginTokenInfos(ctx context.Context, req *types.QueryERC721OriginTokenInfosRequest) (*types.QueryERC721OriginTokenInfosResponse, error) {
	if len(req.ClassId) == 0 {
		return nil, status.Error(codes.InvalidArgument, "empty class id")
	}

	if len(req.TokenIds) == 0 {
		return nil, status.Error(codes.InvalidArgument, "empty token ids")
	}

	tokenOriginIds, tokenUris, err := qs.Keeper.GetOriginTokenInfos(ctx, req.ClassId, req.TokenIds)
	if err != nil && errors.Is(err, collections.ErrNotFound) {
		return nil, status.Error(codes.NotFound, err.Error())
	} else if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	tokenInfos := make([]*types.ERC721OriginTokenInfo, len(tokenOriginIds))
	for i, tokenOriginId := range tokenOriginIds {
		tokenInfos[i] = &types.ERC721OriginTokenInfo{
			TokenOriginId: tokenOriginId,
			TokenUri:      tokenUris[i],
		}
	}

	return &types.QueryERC721OriginTokenInfosResponse{TokenInfos: tokenInfos}, nil
}

// ERC721ClassInfo implements types.QueryServer.
func (qs *queryServerImpl) ERC721ClassInfo(ctx context.Context, req *types.QueryERC721ClassInfoRequest) (*types.QueryERC721ClassInfoResponse, error) {
	if len(req.ClassId) == 0 {
		return nil, status.Error(codes.InvalidArgument, "empty class id")
	}

	className, classUri, classDescs, err := qs.erc721Keeper.GetClassInfo(ctx, req.ClassId)
	if err != nil && errors.Is(err, collections.ErrNotFound) {
		return nil, status.Error(codes.NotFound, err.Error())
	} else if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryERC721ClassInfoResponse{
		ClassInfo: &types.ERC721ClassInfo{
			ClassId:    req.ClassId,
			ClassName:  className,
			ClassUri:   classUri,
			ClassDescs: classDescs,
		},
	}, nil
}

// ERC721ClassInfos implements types.QueryServer.
func (qs *queryServerImpl) ERC721ClassInfos(ctx context.Context, req *types.QueryERC721ClassInfosRequest) (*types.QueryERC721ClassInfosResponse, error) {
	classInfos, pageRes, err := query.CollectionPaginate(ctx, qs.Keeper.ERC721ContractAddrsByClassId, req.Pagination, func(classId string, contractAddr []byte) (types.ERC721ClassInfo, error) {
		className, classUri, classDescs, err := qs.erc721Keeper.GetClassInfo(ctx, classId)
		if err != nil {
			return types.ERC721ClassInfo{}, err
		}

		return types.ERC721ClassInfo{
			ClassId:    classId,
			ClassName:  className,
			ClassUri:   classUri,
			ClassDescs: classDescs,
		}, nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryERC721ClassInfosResponse{
		ClassInfos: classInfos,
		Pagination: pageRes,
	}, nil
}

// Params implements types.QueryServer.
func (qs *queryServerImpl) Params(ctx context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	params, err := qs.Keeper.Params.Get(ctx)
	if err != nil && errors.Is(err, collections.ErrNotFound) {
		return nil, status.Error(codes.NotFound, err.Error())
	} else if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryParamsResponse{Params: params}, nil
}
