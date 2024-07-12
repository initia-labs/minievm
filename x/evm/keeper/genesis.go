package keeper

import (
	"bytes"
	"context"

	"cosmossdk.io/collections"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	coretypes "github.com/ethereum/go-ethereum/core/types"

	"github.com/initia-labs/minievm/x/evm/contracts/erc20_factory"
	"github.com/initia-labs/minievm/x/evm/types"
)

func (k Keeper) Initialize(ctx context.Context) error {
	code, err := hexutil.Decode(erc20_factory.Erc20FactoryBin)
	if err != nil {
		return err
	}

	k.initializing = true
	_, _, _, err = k.EVMCreate2(ctx, types.StdAddress, code, nil, types.ERC20FactorySalt)
	if err != nil {
		return err
	}

	return nil
}

func (k Keeper) InitGenesis(ctx context.Context, genState types.GenesisState) error {
	if err := k.Params.Set(ctx, genState.Params); err != nil {
		return err
	}

	// if the state root is empty, initialize the state
	if common.BytesToHash(genState.StateRoot) == coretypes.EmptyRootHash {
		if err := k.Initialize(ctx); err != nil {
			return err
		}
	} else if err := k.VMRoot.Set(ctx, genState.StateRoot); err != nil {
		return err
	}

	for _, kv := range genState.KeyValues {
		if err := k.VMStore.Set(ctx, kv.Key, kv.Value); err != nil {
			return err
		}
	}

	for _, stores := range genState.Erc20Stores {
		for _, store := range stores.Stores {
			if err := k.ERC20Stores.Set(ctx, collections.Join(stores.Address, store)); err != nil {
				return err
			}
		}
	}

	for _, denomAddress := range genState.DenomAddresses {
		if err := k.ERC20ContractAddrsByDenom.Set(ctx, denomAddress.Denom, denomAddress.ContractAddress); err != nil {
			return err
		}

		if err := k.ERC20DenomsByContractAddr.Set(ctx, denomAddress.ContractAddress, denomAddress.Denom); err != nil {
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

	var stores *types.GenesisERC20Stores
	erc20Stores := []types.GenesisERC20Stores{}
	k.ERC20Stores.Walk(ctx, nil, func(key collections.Pair[[]byte, []byte]) (stop bool, err error) {
		if stores == nil || !bytes.Equal(stores.Address, key.K1()) {
			erc20Stores = append(erc20Stores, types.GenesisERC20Stores{
				Address: key.K1(),
				Stores:  [][]byte{key.K2()},
			})

			stores = &erc20Stores[len(erc20Stores)-1]
			return false, nil
		}

		stores.Stores = append(stores.Stores, key.K2())
		return false, nil
	})

	denomAddresses := []types.GenesisDenomAddress{}
	k.ERC20ContractAddrsByDenom.Walk(ctx, nil, func(denom string, contractAddr []byte) (stop bool, err error) {
		denomAddresses = append(denomAddresses, types.GenesisDenomAddress{
			Denom:           denom,
			ContractAddress: contractAddr,
		})

		return false, nil
	})

	return &types.GenesisState{
		Params:         params,
		StateRoot:      stateRoot,
		KeyValues:      kvs,
		Erc20Stores:    erc20Stores,
		DenomAddresses: denomAddresses,
	}
}
