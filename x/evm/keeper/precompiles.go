package keeper

import (
	"math/big"
	"slices"

	"github.com/ethereum/go-ethereum/core/vm"

	sdk "github.com/cosmos/cosmos-sdk/types"

	cosmosprecompile "github.com/initia-labs/minievm/x/evm/precompiles/cosmos"
	erc20registryprecompile "github.com/initia-labs/minievm/x/evm/precompiles/erc20_registry"
	"github.com/initia-labs/minievm/x/evm/types"
)

// loadPrecompiles loads the precompiled contracts.
func (k *Keeper) loadPrecompiles() error {
	erc20RegistryPrecompile, err := erc20registryprecompile.NewERC20RegistryPrecompile(k.erc20StoresKeeper)
	if err != nil {
		return err
	}

	cosmosPrecompile, err := cosmosprecompile.NewCosmosPrecompile(
		k.cdc,
		k.ac,
		k.accountKeeper,
		k.bankKeeper,
		k,
		k.grpcRouter,
		k.queryCosmosWhitelist,
	)
	if err != nil {
		return err
	}

	// prepare precompiles; always use latest chain config
	// to load all precompiles.
	chainConfig := types.DefaultChainConfig(sdk.Context{})
	rules := chainConfig.Rules(big.NewInt(1), true, 1)

	precompiles := vm.ActivePrecompiledContracts(rules)
	precompiles[types.CosmosPrecompileAddress] = cosmosPrecompile
	precompiles[types.ERC20RegistryPrecompileAddress] = erc20RegistryPrecompile
	k.precompiles = precompiles

	precompileAddrs := slices.Clone(vm.ActivePrecompiles(rules))
	precompileAddrs = append(precompileAddrs, types.CosmosPrecompileAddress, types.ERC20RegistryPrecompileAddress)
	k.precompileAddrs = precompileAddrs

	return nil
}

func (k *Keeper) Precompiles(stateDB types.StateDB) vm.PrecompiledContracts {
	k.precompiles[types.CosmosPrecompileAddress].(types.SetStateDB).SetStateDB(stateDB)
	k.precompiles[types.ERC20RegistryPrecompileAddress].(types.SetStateDB).SetStateDB(stateDB)
	return k.precompiles
}
