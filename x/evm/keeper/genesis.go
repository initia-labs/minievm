package keeper

import (
	"context"

	"github.com/initia-labs/minievm/x/evm/types"
)

func (k Keeper) InitGenesis(ctx context.Context, genState types.GenesisState) error {
	if err := k.Params.Set(ctx, genState.Params); err != nil {
		return err
	}

	if err := k.VMRoot.Set(ctx, genState.StateRoot); err != nil {
		return err
	}

	for _, kv := range genState.KeyValues {
		if err := k.VMStore.Set(ctx, kv.Key, kv.Value); err != nil {
			return err
		}
	}

	return nil
}

func (k Keeper) ExportGenesis(ctx context.Context) *types.GenesisState {
	params, err := k.Params.Get(ctx)
	if err != nil {
		panic(err)
	}

	stateRoot, err := k.VMRoot.Get(ctx)
	if err != nil {
		panic(err)
	}

	kvs := []types.GenesisKeyValue{}
	k.VMStore.Walk(ctx, nil, func(key, value []byte) (stop bool, err error) {
		kvs = append(kvs, types.GenesisKeyValue{Key: key, Value: value})
		return false, nil
	})

	return &types.GenesisState{
		Params:    params,
		StateRoot: stateRoot,
		KeyValues: kvs,
	}
}
