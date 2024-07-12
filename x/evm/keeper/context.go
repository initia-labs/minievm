package keeper

import (
	"context"
	"encoding/json"
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

type callableEVM interface {
	Call(vm.ContractRef, common.Address, []byte, uint64, *uint256.Int) ([]byte, uint64, error)
	StaticCall(vm.ContractRef, common.Address, []byte, uint64) ([]byte, uint64, error)
}

func (k Keeper) buildBlockContext(ctx context.Context, evm callableEVM) (vm.BlockContext, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	headerHash := sdkCtx.HeaderHash()
	if len(headerHash) == 0 {
		headerHash = make([]byte, 32)
	}

	var contractAddr common.Address
	if !k.initializing {
		params, err := k.Params.Get(ctx)
		if err != nil {
			return vm.BlockContext{}, err
		}

		contractAddr, err = types.DenomToContractAddr(ctx, k, params.FeeDenom)
		if err != nil {
			return vm.BlockContext{}, err
		}
	}

	// TODO: should we charge gas for CanTransfer and Transfer?
	//
	// In order to charge gas, we need to fork the EVM and add gas charge
	// logic to the CanTransfer and Transfer functions.
	//
	return vm.BlockContext{
		GasLimit:    k.computeGasLimit(sdkCtx),
		BlockNumber: big.NewInt(sdkCtx.BlockHeight()),
		Time:        uint64(sdkCtx.BlockTime().Unix()),
		CanTransfer: func(sd vm.StateDB, a common.Address, i *uint256.Int) bool {
			if i == nil || i.IsZero() {
				return true
			}

			inputBz, err := k.erc20Keeper.GetERC20ABI().Pack("balanceOf", a)
			if err != nil {
				return false
			}

			retBz, _, err := evm.StaticCall(vm.AccountRef(types.NullAddress), contractAddr, inputBz, 100000)
			if err != nil {
				k.Logger(ctx).Warn("failed to check balance", "error", err)
				return false
			}

			res, err := k.erc20Keeper.GetERC20ABI().Unpack("balanceOf", retBz)
			if err != nil {
				return false
			}

			balance, ok := res[0].(*big.Int)
			if !ok {
				return false
			}

			return i.CmpBig(balance) <= 0
		},
		Transfer: func(sd vm.StateDB, a1, a2 common.Address, i *uint256.Int) {
			if i == nil || i.IsZero() {
				return
			}

			inputBz, err := k.erc20Keeper.GetERC20ABI().Pack("transfer", a2, i.ToBig())
			if err != nil {
				panic(err)
			}

			_, _, err = evm.Call(vm.AccountRef(a1), contractAddr, inputBz, 100000, uint256.NewInt(0))
			if err != nil {
				k.Logger(ctx).Warn("failed to transfer token", "error", err)
				panic(err)
			}
		},
		GetHash: func(u uint64) common.Hash { return common.Hash{} },
		// unused fields
		Coinbase:    common.Address{},
		Difficulty:  nil,
		BaseFee:     nil,
		BlobBaseFee: nil,
		// put header hash to bypass isMerge check in evm
		Random: (*common.Hash)(headerHash),
	}, nil
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

	evm := &vm.EVM{}
	blockContext, err := k.buildBlockContext(ctx, evm)
	if err != nil {
		return ctx, nil, err
	}

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
	*evm = *vm.NewEVMWithPrecompiles(
		blockContext,
		txContext,
		stateDB,
		types.DefaultChainConfig(ctx),
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
func (k Keeper) EVMCall(ctx context.Context, caller common.Address, contractAddr common.Address, inputBz []byte, value *uint256.Int) ([]byte, types.Logs, error) {
	return k.EVMCallWithTracer(ctx, caller, contractAddr, inputBz, value, nil)
}

// EVMCallWithTracer executes an EVM call with the given input data and tracer.
func (k Keeper) EVMCallWithTracer(ctx context.Context, caller common.Address, contractAddr common.Address, inputBz []byte, value *uint256.Int, tracer *tracing.Hooks) ([]byte, types.Logs, error) {
	ctx, evm, err := k.createEVM(ctx, caller, tracer)
	if err != nil {
		return nil, nil, err
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	gasBalance := k.computeGasLimit(sdkCtx)
	if value == nil {
		value = uint256.NewInt(0)
	}

	retBz, gasRemaining, err := evm.Call(
		vm.AccountRef(caller),
		contractAddr,
		inputBz,
		gasBalance,
		value,
	)

	// London enforced
	gasUsed := types.CalGasUsed(gasBalance, gasRemaining, evm.StateDB.GetRefund())
	sdkCtx.GasMeter().ConsumeGas(gasUsed, "EVM gas consumption")
	if err != nil {
		if err == vm.ErrExecutionReverted {
			err = types.NewRevertError(common.CopyBytes(retBz))
		}

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
		types.EventTypeEVM,
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
func (k Keeper) EVMCreate(ctx context.Context, caller common.Address, codeBz []byte, value *uint256.Int) ([]byte, common.Address, types.Logs, error) {
	return k.EVMCreateWithTracer(ctx, caller, codeBz, value, nil, nil)
}

// EVMCreate creates a new contract with the given code.
func (k Keeper) EVMCreate2(ctx context.Context, caller common.Address, codeBz []byte, value *uint256.Int, salt uint64) ([]byte, common.Address, types.Logs, error) {
	return k.EVMCreateWithTracer(ctx, caller, codeBz, value, &salt, nil)
}

// EVMCreateWithTracer creates a new contract with the given code and tracer.
// if salt is nil, it will create a contract with the CREATE opcode.
// if salt is not nil, it will create a contract with the CREATE2 opcode.
func (k Keeper) EVMCreateWithTracer(ctx context.Context, caller common.Address, codeBz []byte, value *uint256.Int, salt *uint64, tracer *tracing.Hooks) (retBz []byte, contractAddr common.Address, logs types.Logs, err error) {
	ctx, evm, err := k.createEVM(ctx, caller, tracer)
	if err != nil {
		return nil, common.Address{}, nil, err
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	gasBalance := k.computeGasLimit(sdkCtx)
	if value == nil {
		value = uint256.NewInt(0)
	}

	var gasRemaining uint64
	if salt == nil {
		retBz, contractAddr, gasRemaining, err = evm.Create(
			vm.AccountRef(caller),
			codeBz,
			gasBalance,
			value,
		)
	} else {
		retBz, contractAddr, gasRemaining, err = evm.Create2(
			vm.AccountRef(caller),
			codeBz,
			gasBalance,
			value,
			uint256.NewInt(*salt),
		)
	}

	// London enforced
	gasUsed := types.CalGasUsed(gasBalance, gasRemaining, evm.StateDB.GetRefund())
	sdkCtx.GasMeter().ConsumeGas(gasUsed, "EVM gas consumption")
	if err != nil {
		if err == vm.ErrExecutionReverted {
			err = types.NewRevertError(common.CopyBytes(retBz))
		}

		return nil, common.Address{}, nil, types.ErrEVMCreateFailed.Wrap(err.Error())
	}

	// commit state transition
	stateDB := evm.StateDB.(*state.StateDB)
	stateRoot, err := stateDB.Commit(evm.Context.BlockNumber.Uint64(), true)
	if err != nil {
		return nil, common.Address{}, nil, err
	}

	// commit trie db
	if stateRoot != coretypes.EmptyRootHash {
		err := stateDB.Database().TrieDB().Commit(stateRoot, false)
		if err != nil {
			return nil, common.Address{}, nil, err
		}
	}

	// update state root
	if err := k.VMRoot.Set(ctx, stateRoot[:]); err != nil {
		return nil, common.Address{}, nil, err
	}

	retHex := hexutil.Encode(retBz)
	logs = types.NewLogs(stateDB.Logs())

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
			return nil, common.Address{}, nil, types.ErrFailedToEncodeLogs.Wrap(err.Error())
		}

		attrs[i] = sdk.NewAttribute(types.AttributeKeyLog, string(jsonBz))
	}
	sdkCtx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeEVM,
		attrs...,
	))

	// handle cosmos messages
	messages := sdkCtx.Value(types.CONTEXT_KEY_COSMOS_MESSAGES).(*[]sdk.Msg)
	if err := k.dispatchMessages(sdkCtx, *messages); err != nil {
		return nil, common.Address{}, nil, err
	}

	return retBz, contractAddr, logs, nil
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
			return err
		}

		// emit events
		sdkCtx.EventManager().EmitEvents(res.GetEvents())
	}

	return nil
}
