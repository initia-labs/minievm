package keeper

import (
	"context"
	"encoding/hex"
	"math"
	"math/big"
	"strings"

	"github.com/holiman/uint256"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	coretypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/crypto"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/initia-labs/minievm/x/evm/contracts/factory"
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
	headerHash := sdkCtx.HeaderHash()
	if len(headerHash) == 0 {
		headerHash = make([]byte, 32)
	}

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
		Random: (*common.Hash)(headerHash),
	}

}

func (k Keeper) buildTxContext(_ context.Context, caller common.Address) vm.TxContext {
	return vm.TxContext{
		Origin:     caller,
		BlobFeeCap: nil,
		BlobHashes: nil,
		GasPrice:   nil,
	}
}

func (k Keeper) createEVM(ctx context.Context, caller common.Address) (*vm.EVM, error) {
	extraEIPs, err := k.ExtraEIPs(ctx)
	if err != nil {
		return nil, err
	}

	blockContext := k.buildBlockContext(ctx)
	txContext := k.buildTxContext(ctx, caller)
	stateDB, err := k.newStateDB(ctx)
	if err != nil {
		return nil, err
	}

	return vm.NewEVMWithPrecompiles(
		blockContext,
		txContext,
		stateDB,
		types.DefaultChainConfig(),
		vm.Config{ExtraEips: extraEIPs},
		// vm.Config{ExtraEips: extraEIPs,
		// 	Tracer: logger.NewJSONLogger(&logger.Config{
		// 		EnableMemory:     false,
		// 		DisableStack:     true,
		// 		DisableStorage:   true,
		// 		EnableReturnData: true,
		// 	}, os.Stderr),
		// },
		k.precompiles.toMap(ctx),
	), nil
}

func (k Keeper) EVMStaticCall(ctx context.Context, caller common.Address, contractAddr common.Address, inputBz []byte) ([]byte, error) {
	evm, err := k.createEVM(ctx, caller)
	if err != nil {
		return nil, err
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	gasBalance := k.computeGasLimit(sdkCtx)

	retBz, gasRemaining, err := evm.StaticCall(
		vm.AccountRef(caller),
		contractAddr,
		inputBz,
		gasBalance,
	)

	// London enforced
	gasUsed := types.CalGasUsed(gasBalance, gasRemaining, evm.StateDB.GetRefund())
	sdkCtx.GasMeter().ConsumeGas(gasUsed, "EVM gas consumption")
	if err != nil {
		return nil, types.ErrEVMCallFailed.Wrap(err.Error())
	}

	return retBz, nil
}

func (k Keeper) EVMCall(ctx context.Context, caller common.Address, contractAddr common.Address, inputBz []byte) ([]byte, types.Logs, error) {
	evm, err := k.createEVM(ctx, caller)
	if err != nil {
		return nil, nil, err
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	gasBalance := k.computeGasLimit(sdkCtx)

	retBz, gasRemaining, err := evm.Call(
		vm.AccountRef(caller),
		contractAddr,
		inputBz,
		gasBalance,
		uint256.NewInt(0),
	)

	// London enforced
	gasUsed := types.CalGasUsed(gasBalance, gasRemaining, evm.StateDB.GetRefund())
	sdkCtx.GasMeter().ConsumeGas(gasUsed, "EVM gas consumption")
	if err != nil {
		return nil, nil, types.ErrEVMCallFailed.Wrap(err.Error())
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

	retHex := common.Bytes2Hex(retBz)
	logs := types.NewLogs(stateDB.Logs())

	// emit action events
	sdkCtx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeCall,
		sdk.NewAttribute(types.AttributeKeyContract, contractAddr.Hex()),
		sdk.NewAttribute(types.AttributeKeyRet, retHex),
	))

	// emit logs events
	for _, log := range logs {
		for _, topic := range log.Topics {
			sdkCtx.EventManager().EmitEvent(sdk.NewEvent(
				types.EventTypeLog,
				sdk.NewAttribute(types.AttributeKeyTopic, topic),
				sdk.NewAttribute(types.AttributeKeyRet, log.Data),
			))
		}
	}

	return retBz, logs, nil
}

func (k Keeper) EVMCreate(ctx context.Context, caller common.Address, codeBz []byte) ([]byte, common.Address, error) {
	evm, err := k.createEVM(ctx, caller)
	if err != nil {
		return nil, common.Address{}, err
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	gasBalance := k.computeGasLimit(sdkCtx)
	retBz, contractAddr, gasRemaining, err := evm.Create(
		vm.AccountRef(caller),
		codeBz,
		gasBalance,
		uint256.NewInt(0),
	)

	// London enforced
	gasUsed := types.CalGasUsed(gasBalance, gasRemaining, evm.StateDB.GetRefund())
	sdkCtx.GasMeter().ConsumeGas(gasUsed, "EVM gas consumption")
	if err != nil {
		return nil, common.Address{}, types.ErrEVMCreateFailed.Wrap(err.Error())
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

	retHex := common.Bytes2Hex(retBz)

	// emit action events
	sdkCtx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeCall,
		sdk.NewAttribute(types.AttributeKeyContract, contractAddr.Hex()),
		sdk.NewAttribute(types.AttributeKeyRet, retHex),
	))

	return retBz, contractAddr, nil
}

func (k Keeper) Initialize(ctx context.Context) error {
	bin, err := hex.DecodeString(strings.TrimPrefix(factory.FactoryBin, "0x"))
	if err != nil {
		return err
	}

	_, _, err = k.EVMCreate(ctx, types.StdAddress, bin)
	if err != nil {
		return err
	}

	return nil
}

func (k Keeper) NextContractAddress(ctx context.Context, caller common.Address) (common.Address, error) {
	stateDB, err := k.newStateDB(ctx)
	if err != nil {
		return common.Address{}, err
	}

	return crypto.CreateAddress(caller, stateDB.GetNonce(caller)), nil
}
