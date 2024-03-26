package keeper

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"

	erc20registry "github.com/initia-labs/minievm/x/evm/precompiles/erc20_registry"
	"github.com/initia-labs/minievm/x/evm/types"
)

// precompile is a precompiled contract.
type precompile struct {
	addr     common.Address
	contract vm.PrecompiledContract
}

// loadPrecompiles loads the precompiled contracts.
func (k *Keeper) loadPrecompiles() error {
	erc20Registry, err := erc20registry.NewERC20Registry(k.erc20StoresKeeper)
	if err != nil {
		return err
	}

	k.precompiles = precompiles{
		{
			addr:     common.BytesToAddress([]byte{0xf1}),
			contract: erc20Registry,
		},
	}

	return nil
}

// precompiles is a list of precompiled contracts.
type precompiles []precompile

// toMap converts the precompiles to a map.
func (ps precompiles) toMap(ctx context.Context) map[common.Address]vm.PrecompiledContract {
	m := make(map[common.Address]vm.PrecompiledContract)
	for _, p := range ps {
		m[p.addr] = p.contract.(types.WithContext).WithContext(ctx)
	}

	return m
}
