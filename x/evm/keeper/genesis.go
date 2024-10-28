package keeper

import (
	"bytes"
	"context"

	"cosmossdk.io/collections"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/initia-labs/minievm/x/evm/contracts/erc20_factory"
	"github.com/initia-labs/minievm/x/evm/contracts/erc20_wrapper"
	"github.com/initia-labs/minievm/x/evm/types"
)

// Initialize initializes the EVM state at genesis by performing the following steps:
// 1. Deploy and store the ERC20 factory contract.
// 2. Deploy fee denom ERC20 coins at genesis bootstrapping with 18 decimals.
// 3. Deploy and store the wrapper ERC20 factory contract for IBC transfers.
func (k Keeper) Initialize(ctx context.Context) error {
	return k.InitializeWithDecimals(ctx, types.EtherDecimals)
}

// InitializeWithDecimals initializes the EVM state at genesis with the given decimals
func (k Keeper) InitializeWithDecimals(ctx context.Context, decimals uint8) error {
	// 1. Deploy and store the ERC20 factory contract.
	err := k.DeployERC20Factory(ctx)
	if err != nil {
		return err
	}

	params, err := k.Params.Get(ctx)
	if err != nil {
		return err
	}

	// 2. Deploy fee denom ERC20 coins at genesis bootstrapping with decimals.
	err = k.erc20Keeper.CreateERC20(ctx, params.FeeDenom, decimals)
	if err != nil {
		return err
	}

	// 3. Deploy and store the ERC20 wrapper factory contract for IBC transfers.
	err = k.DeployERC20Wrapper(ctx)
	if err != nil {
		return err
	}

	return nil
}

// DeployERC20Factory deploys the ERC20 factory contract and stores the address in the keeper.
func (k Keeper) DeployERC20Factory(ctx context.Context) error {
	factoryCode, err := hexutil.Decode(erc20_factory.Erc20FactoryBin)
	if err != nil {
		return err
	}

	_, factoryAddr, _, err := k.EVMCreate2(ctx, types.StdAddress, factoryCode, nil, types.ERC20FactorySalt, nil)
	if err != nil {
		return err
	}

	return k.ERC20FactoryAddr.Set(ctx, factoryAddr.Bytes())
}

// DeployERC20Wrapper deploys the ERC20 wrapper contract and stores the address in the keeper.
func (k Keeper) DeployERC20Wrapper(ctx context.Context) error {
	factoryAddr, err := k.ERC20FactoryAddr.Get(ctx)
	if err != nil {
		return err
	}

	wrapperCode, err := hexutil.Decode(erc20_wrapper.Erc20WrapperBin)
	if err != nil {
		return err
	}
	abi, _ := erc20_wrapper.Erc20WrapperMetaData.GetAbi()
	wrapperConstructorArg, err := abi.Constructor.Inputs.Pack(common.BytesToAddress(factoryAddr))
	if err != nil {
		return err
	}

	_, wrapperAddr, _, err := k.EVMCreate2(ctx, types.StdAddress, append(wrapperCode, wrapperConstructorArg...), nil, types.ERC20WrapperSalt, nil)
	if err != nil {
		return err
	}
	if err = k.ERC20WrapperAddr.Set(ctx, wrapperAddr.Bytes()); err != nil {
		return err
	}

	// whitelist the wrapper contract for IBC hook
	if err = k.ibcHookKeeper.SetAllowed(ctx, wrapperAddr.Bytes(), true); err != nil {
		return err
	}

	return nil
}

func (k Keeper) InitGenesis(ctx context.Context, genState types.GenesisState) error {
	if err := k.Params.Set(ctx, genState.Params); err != nil {
		return err
	}

	// if the state root is empty, initialize the state
	if len(genState.KeyValues) == 0 {
		if err := k.Initialize(ctx); err != nil {
			return err
		}
	} else {
		if err := k.ERC20FactoryAddr.Set(ctx, genState.Erc20Factory); err != nil {
			return err
		}
		if err := k.ERC20WrapperAddr.Set(ctx, genState.Erc20Wrapper); err != nil {
			return err
		}
	}

	for _, kv := range genState.KeyValues {
		if err := k.VMStore.Set(ctx, kv.Key, kv.Value); err != nil {
			return err
		}
	}

	for _, erc20 := range genState.ERC20s {
		if err := k.ERC20s.Set(ctx, erc20); err != nil {
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

	for _, denomTrace := range genState.DenomTraces {
		if err := k.ERC20ContractAddrsByDenom.Set(ctx, denomTrace.Denom, denomTrace.ContractAddress); err != nil {
			return err
		}

		if err := k.ERC20DenomsByContractAddr.Set(ctx, denomTrace.ContractAddress, denomTrace.Denom); err != nil {
			return err
		}
	}

	for _, classTrace := range genState.ClassTraces {
		if err := k.ERC721ContractAddrsByClassId.Set(ctx, classTrace.ClassId, classTrace.ContractAddress); err != nil {
			return err
		}

		if err := k.ERC721ClassIdsByContractAddr.Set(ctx, classTrace.ContractAddress, classTrace.ClassId); err != nil {
			return err
		}

		if err := k.ERC721ClassURIs.Set(ctx, classTrace.ContractAddress, classTrace.Uri); err != nil {
			return err
		}
	}

	for _, blockHash := range genState.EVMBlockHashes {
		if err := k.EVMBlockHashes.Set(ctx, blockHash.Height, blockHash.Hash); err != nil {
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

	kvs := []types.GenesisKeyValue{}
	err = k.VMStore.Walk(ctx, nil, func(key, value []byte) (stop bool, err error) {
		kvs = append(kvs, types.GenesisKeyValue{Key: key, Value: value})
		return false, nil
	})
	if err != nil {
		panic(err)
	}

	erc20s := [][]byte{}
	err = k.ERC20s.Walk(ctx, nil, func(erc20 []byte) (stop bool, err error) {
		erc20s = append(erc20s, erc20)
		return false, nil
	})
	if err != nil {
		panic(err)
	}

	var stores *types.GenesisERC20Stores
	erc20Stores := []types.GenesisERC20Stores{}
	err = k.ERC20Stores.Walk(ctx, nil, func(key collections.Pair[[]byte, []byte]) (stop bool, err error) {
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
	if err != nil {
		panic(err)
	}

	denomTraces := []types.GenesisDenomTrace{}
	err = k.ERC20ContractAddrsByDenom.Walk(ctx, nil, func(denom string, contractAddr []byte) (stop bool, err error) {
		denomTraces = append(denomTraces, types.GenesisDenomTrace{
			Denom:           denom,
			ContractAddress: contractAddr,
		})

		return false, nil
	})
	if err != nil {
		panic(err)
	}

	classTraces := []types.GenesisClassTrace{}
	err = k.ERC721ContractAddrsByClassId.Walk(ctx, nil, func(classId string, contractAddr []byte) (stop bool, err error) {
		uri, err := k.ERC721ClassURIs.Get(ctx, contractAddr)
		if err != nil {
			panic(err)
		}

		classTraces = append(classTraces, types.GenesisClassTrace{
			ClassId:         classId,
			ContractAddress: contractAddr,
			Uri:             uri,
		})

		return false, nil
	})
	if err != nil {
		panic(err)
	}

	evmBlockHashes := []types.GenesisEVMBlockHash{}
	err = k.EVMBlockHashes.Walk(ctx, nil, func(height uint64, hash []byte) (stop bool, err error) {
		evmBlockHashes = append(evmBlockHashes, types.GenesisEVMBlockHash{
			Height: height,
			Hash:   hash,
		})

		return false, nil
	})
	if err != nil {
		panic(err)
	}

	factoryAddr, err := k.ERC20FactoryAddr.Get(ctx)
	if err != nil {
		panic(err)
	}

	wrapperAddr, err := k.ERC20WrapperAddr.Get(ctx)
	if err != nil {
		panic(err)
	}

	return &types.GenesisState{
		Params:         params,
		KeyValues:      kvs,
		ERC20s:         erc20s,
		Erc20Stores:    erc20Stores,
		DenomTraces:    denomTraces,
		ClassTraces:    classTraces,
		Erc20Factory:   factoryAddr,
		Erc20Wrapper:   wrapperAddr,
		EVMBlockHashes: evmBlockHashes,
	}
}
