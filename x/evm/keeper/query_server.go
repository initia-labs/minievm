package keeper

import (
	"context"

	storetypes "cosmossdk.io/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ethereum/go-ethereum/common"

	"github.com/initia-labs/minievm/x/evm/types"
)

type queryServerImpl struct {
	*Keeper
}

func NewQueryServer(k *Keeper) types.QueryServer {
	return &queryServerImpl{k}
}

// Call implements types.QueryServer.
func (qs *queryServerImpl) Call(ctx context.Context, req *types.QueryCallRequest) (*types.QueryCallResponse, error) {
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

	retBz, logs, err := qs.EVMCall(sdkCtx, sender, contractAddr, req.Input)
	if err != nil {
		return nil, err
	}

	return &types.QueryCallResponse{
		Response: retBz,
		UsedGas:  sdkCtx.GasMeter().GasConsumedToLimit(),
		Logs:     logs,
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

	// argument validation
	if len(contractAddr) != common.AddressLength {
		return nil, types.ErrInvalidAddressLength
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

	// argument validation
	if len(contractAddr) != common.AddressLength {
		return nil, types.ErrInvalidAddressLength
	}

	state := stateDB.GetState(common.Address(contractAddr.Bytes()), common.HexToHash(req.Key))
	return &types.QueryStateResponse{
		Value: state.Hex(),
	}, nil
}

// Params implements types.QueryServer.
func (qs *queryServerImpl) Params(ctx context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	params, err := qs.Keeper.Params.Get(ctx)
	if err != nil {
		return nil, err
	}

	return &types.QueryParamsResponse{Params: params}, nil
}
