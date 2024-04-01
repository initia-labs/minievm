package keeper

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"math/big"

	"github.com/holiman/uint256"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/tracing"
	coretypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/crypto"

	sdk "github.com/cosmos/cosmos-sdk/types"

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

// createEVM creates a new EVM instance.
func (k Keeper) createEVM(ctx context.Context, caller common.Address, tracer *tracing.Hooks) (context.Context, *vm.EVM, error) {
	extraEIPs, err := k.ExtraEIPs(ctx)
	if err != nil {
		return ctx, nil, err
	}

	blockContext := k.buildBlockContext(ctx)
	txContext := k.buildTxContext(ctx, caller)
	stateDB, err := k.newStateDB(ctx)
	if err != nil {
		return ctx, nil, err
	}

	vmConfig := vm.Config{
		Tracer:              tracer,
		ExtraEips:           extraEIPs,
		ContractCreatedHook: k.contractCreatedHook(ctx),
	}

	// set cosmos messages to context
	ctx = sdk.UnwrapSDKContext(ctx).WithValue(types.CONTEXT_KEY_COSMOS_MESSAGES, &[]sdk.Msg{})
	evm := vm.NewEVMWithPrecompiles(
		blockContext,
		txContext,
		stateDB,
		types.DefaultChainConfig(),
		vmConfig,
		k.precompiles.toMap(ctx),
	)

	if tracer != nil {
		// register vm context to tracer
		tracer.OnTxStart(evm.GetVMContext(), nil, caller)
	}

	return ctx, evm, nil
}

// contractCreatedHook returns a callback function that is called when a contract is created.
//
// It converts a normal account to a contract account if the account is empty and create
// creates a contract account if the account does not exist.
func (k Keeper) contractCreatedHook(ctx context.Context) vm.ContractCreatedHook {
	return func(contractAddr common.Address) error {
		if k.accountKeeper.HasAccount(ctx, sdk.AccAddress(contractAddr.Bytes())) {
			account := k.accountKeeper.GetAccount(ctx, sdk.AccAddress(contractAddr.Bytes()))

			// check the account is empty or not
			if !types.IsEmptyAccount(account) {
				return types.ErrAddressAlreadyExists.Wrap(contractAddr.String())
			}

			// convert base account to contract account only if this account is empty
			contractAccount := types.NewContractAccountWithAddress(contractAddr.Bytes())
			contractAccount.AccountNumber = account.GetAccountNumber()
			k.accountKeeper.SetAccount(ctx, contractAccount)
		} else {
			// create contract account
			contractAccount := types.NewContractAccountWithAddress(contractAddr.Bytes())
			contractAccount.AccountNumber = k.accountKeeper.NextAccountNumber(ctx)
			k.accountKeeper.SetAccount(ctx, contractAccount)
		}

		return nil
	}
}

// EVMStaticCall executes an EVM call with the given input data in static mode.
func (k Keeper) EVMStaticCall(ctx context.Context, caller common.Address, contractAddr common.Address, inputBz []byte) ([]byte, error) {
	return k.EVMStaticCallWithTracer(ctx, caller, contractAddr, inputBz, nil)
}

// EVMStaticCallWithTracer executes an EVM call with the given input data and tracer in static mode.
func (k Keeper) EVMStaticCallWithTracer(ctx context.Context, caller common.Address, contractAddr common.Address, inputBz []byte, tracer *tracing.Hooks) ([]byte, error) {
	ctx, evm, err := k.createEVM(ctx, caller, tracer)
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

// EVMCall executes an EVM call with the given input data.
func (k Keeper) EVMCall(ctx context.Context, caller common.Address, contractAddr common.Address, inputBz []byte) ([]byte, types.Logs, error) {
	return k.EVMCallWithTracer(ctx, caller, contractAddr, inputBz, nil)
}

// EVMCallWithTracer executes an EVM call with the given input data and tracer.
func (k Keeper) EVMCallWithTracer(ctx context.Context, caller common.Address, contractAddr common.Address, inputBz []byte, tracer *tracing.Hooks) ([]byte, types.Logs, error) {
	ctx, evm, err := k.createEVM(ctx, caller, tracer)
	if err != nil {
		return nil, nil, err
	}

	// check the contract is empty or not
	if !types.IsPrecompileAddress(contractAddr) && evm.StateDB.GetCodeSize(contractAddr) == 0 {
		return nil, nil, types.ErrEmptyContractAddress.Wrap(contractAddr.String())
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

	retHex := hexutil.Encode(retBz)
	logs := types.NewLogs(stateDB.Logs())

	// emit action events
	sdkCtx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeCall,
		sdk.NewAttribute(types.AttributeKeyContract, contractAddr.Hex()),
		sdk.NewAttribute(types.AttributeKeyRet, retHex),
	))

	// emit logs events
	attrs := make([]sdk.Attribute, len(logs))
	for i, log := range logs {
		jsonBz, err := json.Marshal(log)
		if err != nil {
			return nil, nil, types.ErrFailedToEncodeLogs.Wrap(err.Error())
		}

		attrs[i] = sdk.NewAttribute(types.AttributeKeyLog, string(jsonBz))
	}
	sdkCtx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeLogs,
		attrs...,
	))

	// handle cosmos messages
	messages := sdkCtx.Value(types.CONTEXT_KEY_COSMOS_MESSAGES).(*[]sdk.Msg)
	if err := k.dispatchMessages(sdkCtx, *messages); err != nil {
		return nil, nil, err
	}

	return retBz, logs, nil
}

// EVMCreate creates a new contract with the given code.
func (k Keeper) EVMCreate(ctx context.Context, caller common.Address, codeBz []byte) ([]byte, common.Address, error) {
	return k.EVMCreateWithTracer(ctx, caller, codeBz, nil, nil)
}

// EVMCreate creates a new contract with the given code.
func (k Keeper) EVMCreate2(ctx context.Context, caller common.Address, codeBz []byte, salt uint64) ([]byte, common.Address, error) {
	return k.EVMCreateWithTracer(ctx, caller, codeBz, &salt, nil)
}

// EVMCreateWithTracer creates a new contract with the given code and tracer.
// if salt is nil, it will create a contract with the CREATE opcode.
// if salt is not nil, it will create a contract with the CREATE2 opcode.
func (k Keeper) EVMCreateWithTracer(ctx context.Context, caller common.Address, codeBz []byte, salt *uint64, tracer *tracing.Hooks) (retBz []byte, contractAddr common.Address, err error) {
	ctx, evm, err := k.createEVM(ctx, caller, tracer)
	if err != nil {
		return nil, common.Address{}, err
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	gasBalance := k.computeGasLimit(sdkCtx)

	var gasRemaining uint64
	if salt == nil {
		retBz, contractAddr, gasRemaining, err = evm.Create(
			vm.AccountRef(caller),
			codeBz,
			gasBalance,
			uint256.NewInt(0),
		)
	} else {
		retBz, contractAddr, gasRemaining, err = evm.Create2(
			vm.AccountRef(caller),
			codeBz,
			gasBalance,
			uint256.NewInt(0),
			uint256.NewInt(*salt),
		)
	}

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

	retHex := hexutil.Encode(retBz)
	logs := types.NewLogs(stateDB.Logs())

	// emit action events
	sdkCtx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeCreate,
		sdk.NewAttribute(types.AttributeKeyContract, contractAddr.Hex()),
		sdk.NewAttribute(types.AttributeKeyRet, retHex),
	))

	// emit logs events
	attrs := make([]sdk.Attribute, len(logs))
	for i, log := range logs {
		jsonBz, err := json.Marshal(log)
		if err != nil {
			return nil, common.Address{}, types.ErrFailedToEncodeLogs.Wrap(err.Error())
		}

		attrs[i] = sdk.NewAttribute(types.AttributeKeyLog, string(jsonBz))
	}
	sdkCtx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeLogs,
		attrs...,
	))

	// handle cosmos messages
	messages := sdkCtx.Value(types.CONTEXT_KEY_COSMOS_MESSAGES).(*[]sdk.Msg)
	if err := k.dispatchMessages(sdkCtx, *messages); err != nil {
		return nil, common.Address{}, err
	}

	return retBz, contractAddr, nil
}

// nextContractAddress returns the next contract address which will be created by the given caller
// in CREATE opcode.
func (k Keeper) nextContractAddress(ctx context.Context, caller common.Address) (common.Address, error) {
	stateDB, err := k.newStateDB(ctx)
	if err != nil {
		return common.Address{}, err
	}

	return crypto.CreateAddress(caller, stateDB.GetNonce(caller)), nil
}

// dispatchMessages run the given cosmos msgs and emit events
func (k Keeper) dispatchMessages(ctx context.Context, msgs []sdk.Msg) error {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	for _, msg := range msgs {

		// validate msg
		if msg, ok := msg.(sdk.HasValidateBasic); ok {
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
		}

		// find the handler
		handler := k.msgRouter.Handler(msg)
		if handler == nil {
			return types.ErrNotSupportedCosmosMessage
		}

		//  and execute it
		res, err := handler(sdkCtx, msg)
		if err != nil {
			fmt.Println(err)
			return err
		}

		// emit events
		sdkCtx.EventManager().EmitEvents(res.GetEvents())
	}

	return nil
}
