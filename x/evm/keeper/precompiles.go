package keeper

import (
	"slices"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/params"

	cosmosprecompile "github.com/initia-labs/minievm/x/evm/precompiles/cosmos"
	erc20registryprecompile "github.com/initia-labs/minievm/x/evm/precompiles/erc20_registry"
	"github.com/initia-labs/minievm/x/evm/precompiles/jsonutils"
	"github.com/initia-labs/minievm/x/evm/types"
)

// precompiles returns the precompiled contracts for the EVM.
func (k *Keeper) precompiles(rules params.Rules, stateDB types.StateDB) (vm.PrecompiledContracts, error) {
	erc20RegistryPrecompile, err := erc20registryprecompile.NewERC20RegistryPrecompile(stateDB, k.erc20StoresKeeper)
	if err != nil {
		return nil, err
	}

	cosmosPrecompile, err := cosmosprecompile.NewCosmosPrecompile(
		stateDB,
		k.cdc,
		k.ac,
		k.accountKeeper,
		k.bankKeeper,
		k,
		k.grpcRouter,
		k.queryCosmosWhitelist,
		k.authority,
	)
	if err != nil {
		return nil, err
	}

	jsonutilsPrecompile, err := jsonutils.NewJSONUtilsPrecompile(stateDB)
	if err != nil {
		return nil, err
	}

	// clone the active precompiles and add the new precompiles
	precompiles := vm.ActivePrecompiledContracts(rules)
	precompiles[types.CosmosPrecompileAddress] = cosmosPrecompile
	precompiles[types.ERC20RegistryPrecompileAddress] = erc20RegistryPrecompile
	precompiles[types.JSONUtilsPrecompileAddress] = jsonutilsPrecompile

	return precompiles, nil
}

// PrecompileAddrs returns the precompile addresses for the EVM.
func (k *Keeper) precompileAddrs(rules params.Rules) []common.Address {
	addrs := append(slices.Clone(vm.ActivePrecompiles(rules)), types.CosmosPrecompileAddress, types.ERC20RegistryPrecompileAddress)
	return addrs
}
