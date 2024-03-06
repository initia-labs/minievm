package keeper

import (
	"context"
	"math"
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	coretypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/holiman/uint256"

	"github.com/initia-labs/minievm/x/evm/types"
)

func (k Keeper) computeGasLimit(sdkCtx sdk.Context) uint64 {
	gasLimit := sdkCtx.GasMeter().GasRemaining()
	if sdkCtx.ExecMode() == sdk.ExecModeSimulate {
		gasLimit = k.config.ContractSimulationGasLimit
	} else if sdkCtx.GasMeter().Limit() == 0 {
		// infinite gas meter
		gasLimit = math.MaxUint64
	}

	return gasLimit
}

func (k Keeper) buildBlockContext(ctx context.Context) vm.BlockContext {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	return vm.BlockContext{
		GasLimit:    k.computeGasLimit(sdkCtx),
		BlockNumber: big.NewInt(sdkCtx.BlockHeight()),
		Time:        uint64(sdkCtx.BlockTime().Unix()),
		CanTransfer: func(sd vm.StateDB, a common.Address, i *uint256.Int) bool { return true },
		Transfer:    func(sd vm.StateDB, a1, a2 common.Address, i *uint256.Int) {},
		GetHash:     func(u uint64) common.Hash { return common.Hash{} },
		// unused fields
		Coinbase:    common.Address{},
		Difficulty:  nil,
		BaseFee:     nil,
		BlobBaseFee: nil,
		// put header hash to bypass isMerge check in evm
		Random: (*common.Hash)(sdkCtx.HeaderHash()),
	}

}

func (k Keeper) buildTxContext(ctx context.Context, sender sdk.AccAddress) vm.TxContext {
	return vm.TxContext{
		Origin:     common.BytesToAddress(sender.Bytes()),
		BlobFeeCap: nil,
		BlobHashes: nil,
		GasPrice:   nil,
	}
}

func (k Keeper) createEVM(ctx context.Context, sender sdk.AccAddress) (*vm.EVM, error) {
	extraEIPs, err := k.ExtraEIPs(ctx)
	if err != nil {
		return nil, err
	}

	blockContext := k.buildBlockContext(ctx)
	txContext := k.buildTxContext(ctx, sender)
	stateDB, err := k.newStateDB(ctx)
	if err != nil {
		return nil, err
	}

	return vm.NewEVM(blockContext, txContext, stateDB, types.DefaultChainConfig(), vm.Config{ExtraEips: extraEIPs}), nil
}

func (k Keeper) EVMCall(ctx context.Context, sender sdk.AccAddress, contractAddr common.Address, inputBz []byte) ([]byte, types.Logs, error) {
	evm, err := k.createEVM(ctx, sender)
	if err != nil {
		return nil, nil, err
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	gasBalance := k.computeGasLimit(sdkCtx)

	retBz, gasRemaining, err := evm.Call(
		vm.AccountRef(common.BytesToAddress(sender)),
		contractAddr,
		inputBz,
		gasBalance,
		uint256.NewInt(0),
	)

	// London enforced
	gasUsed := types.CalGasUsed(gasBalance, gasRemaining, evm.StateDB.GetRefund())
	sdkCtx.GasMeter().ConsumeGas(gasUsed, "EVM gas consumption")
	if err != nil {
		return nil, nil, err
	}

	// commit state transition
	stateDB := evm.StateDB.(*state.StateDB)
	stateRoot, err := stateDB.Commit(evm.Context.BlockNumber.Uint64(), true)
	if err != nil {
		return nil, nil, err
	}

	// commit trie db
	if stateRoot != coretypes.EmptyRootHash {
		err := stateDB.Database().TrieDB().Commit(stateRoot, false)
		if err != nil {
			return nil, nil, err
		}
	}

	// update state root
	if err := k.VMRoot.Set(ctx, stateRoot[:]); err != nil {
		return nil, nil, err
	}

	logs := types.NewLogs(stateDB.Logs())
	return retBz, logs, nil
}

func (k Keeper) EVMCreate(ctx context.Context, sender sdk.AccAddress, codeBz []byte) ([]byte, common.Address, error) {
	evm, err := k.createEVM(ctx, sender)
	if err != nil {
		return nil, common.Address{}, err
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	gasBalance := k.computeGasLimit(sdkCtx)
	retBz, contractAddr, gasRemaining, err := evm.Create(
		vm.AccountRef(common.BytesToAddress(sender)),
		codeBz,
		gasBalance,
		uint256.NewInt(0),
	)

	// London enforced
	gasUsed := types.CalGasUsed(gasBalance, gasRemaining, evm.StateDB.GetRefund())
	sdkCtx.GasMeter().ConsumeGas(gasUsed, "EVM gas consumption")
	if err != nil {
		return nil, common.Address{}, err
	}

	// commit state transition
	stateDB := evm.StateDB.(*state.StateDB)
	stateRoot, err := stateDB.Commit(evm.Context.BlockNumber.Uint64(), true)
	if err != nil {
		return nil, common.Address{}, err
	}

	// commit trie db
	if stateRoot != coretypes.EmptyRootHash {
		err := stateDB.Database().TrieDB().Commit(stateRoot, false)
		if err != nil {
			return nil, common.Address{}, err
		}
	}

	// update state root
	if err := k.VMRoot.Set(ctx, stateRoot[:]); err != nil {
		return nil, common.Address{}, err
	}

	return retBz, contractAddr, nil
}
