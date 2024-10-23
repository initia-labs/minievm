package keeper

import (
	"context"
	"encoding/json"
	"math/big"

	"github.com/holiman/uint256"

	storetypes "cosmossdk.io/store/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/tracing"
	coretype "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/trie/utils"

	sdk "github.com/cosmos/cosmos-sdk/types"

	evmstate "github.com/initia-labs/minievm/x/evm/state"
	"github.com/initia-labs/minievm/x/evm/types"
)

func (k Keeper) NewStateDB(ctx context.Context, evm callableEVM, fee types.Fee) (*evmstate.StateDB, error) {
	return evmstate.NewStateDB(
		// delegate gas meter to the EVM
		sdk.UnwrapSDKContext(ctx).WithGasMeter(storetypes.NewInfiniteGasMeter()), k.Logger(ctx),
		k.accountKeeper, k.VMStore, k.TransientVMStore, k.TransientCreated,
		k.TransientSelfDestruct, k.TransientLogs, k.TransientLogSize,
		k.TransientAccessList, k.TransientRefund, k.execIndex,
		evm, k.ERC20Keeper().GetERC20ABI(), fee.Contract(),
	)
}

func (k Keeper) computeGasLimit(sdkCtx sdk.Context) uint64 {
	gasLimit := sdkCtx.GasMeter().Limit() - sdkCtx.GasMeter().GasConsumedToLimit()
	if sdkCtx.ExecMode() == sdk.ExecModeSimulate {
		gasLimit = k.config.ContractSimulationGasLimit
	}

	return gasLimit
}

type callableEVM interface {
	Call(vm.ContractRef, common.Address, []byte, uint64, *uint256.Int) ([]byte, uint64, error)
	StaticCall(vm.ContractRef, common.Address, []byte, uint64) ([]byte, uint64, error)
}

func (k Keeper) buildBlockContext(ctx context.Context, evm callableEVM, fee types.Fee) (vm.BlockContext, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	headerHash := sdkCtx.HeaderHash()
	if len(headerHash) == 0 {
		headerHash = make([]byte, 32)
	}

	baseFee, err := k.baseFee(ctx, fee)
	if err != nil {
		return vm.BlockContext{}, err
	}

	return vm.BlockContext{
		BaseFee:     baseFee,
		GasLimit:    k.computeGasLimit(sdkCtx),
		BlockNumber: big.NewInt(sdkCtx.BlockHeight()),
		Time:        uint64(sdkCtx.BlockTime().Unix()),
		CanTransfer: func(sd vm.StateDB, a common.Address, i *uint256.Int) bool {
			if i == nil || i.IsZero() {
				return true
			}

			// if the contract is not found, return false
			if (fee.Contract() == common.Address{}) {
				return false
			}

			inputBz, err := k.erc20Keeper.GetERC20ABI().Pack("balanceOf", a)
			if err != nil {
				return false
			}

			retBz, _, err := evm.StaticCall(vm.AccountRef(types.NullAddress), fee.Contract(), inputBz, 100000)
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

			_, _, err = evm.Call(vm.AccountRef(a1), fee.Contract(), inputBz, 100000, uint256.NewInt(0))
			if err != nil {
				k.Logger(ctx).Warn("failed to transfer token", "error", err)
				panic(err)
			}
		},
		GetHash: func(n uint64) common.Hash {
			bz, err := k.EVMBlockHashes.Get(sdkCtx, n)
			if err != nil {
				return common.Hash{}
			}

			return common.BytesToHash(bz)
		},
		// put header hash to bypass isMerge check in evm
		Random: (*common.Hash)(headerHash),
		// unused fields
		Coinbase:    common.Address{},
		Difficulty:  big.NewInt(0),
		BlobBaseFee: big.NewInt(0),
	}, nil
}

func (k Keeper) buildTxContext(ctx context.Context, caller common.Address, fee types.Fee) (vm.TxContext, error) {
	gasPrice, err := k.extractGasPriceFromContext(ctx, fee)
	if err != nil {
		return vm.TxContext{}, err
	}

	return vm.TxContext{
		Origin:       caller,
		GasPrice:     gasPrice,
		AccessEvents: state.NewAccessEvents(utils.NewPointCache(4096)),
		// unused fields
		BlobFeeCap: big.NewInt(0),
		BlobHashes: []common.Hash{},
	}, nil
}

// createEVM creates a new EVM instance.
func (k Keeper) CreateEVM(ctx context.Context, caller common.Address, tracer *tracing.Hooks) (context.Context, *vm.EVM, error) {
	extraEIPs, err := k.ExtraEIPs(ctx)
	if err != nil {
		return ctx, nil, err
	}

	fee, err := k.LoadFee(ctx)
	if err != nil {
		return ctx, nil, err
	}

	evm := &vm.EVM{}
	blockContext, err := k.buildBlockContext(ctx, evm, fee)
	if err != nil {
		return ctx, nil, err
	}
	txContext, err := k.buildTxContext(ctx, caller, fee)
	if err != nil {
		return ctx, nil, err
	}
	stateDB, err := k.NewStateDB(ctx, evm, fee)
	if err != nil {
		return ctx, nil, err
	}

	vmConfig := vm.Config{
		Tracer:    tracer,
		ExtraEips: extraEIPs,
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

// EVMStaticCall executes an EVM call with the given input data in static mode.
func (k Keeper) EVMStaticCall(ctx context.Context, caller common.Address, contractAddr common.Address, inputBz []byte, accessList coretype.AccessList) ([]byte, error) {
	return k.EVMStaticCallWithTracer(ctx, caller, contractAddr, inputBz, accessList, nil)
}

// EVMStaticCallWithTracer executes an EVM call with the given input data and tracer in static mode.
func (k Keeper) EVMStaticCallWithTracer(ctx context.Context, caller common.Address, contractAddr common.Address, inputBz []byte, accessList coretype.AccessList, tracer *tracing.Hooks) ([]byte, error) {
	ctx, evm, err := k.CreateEVM(ctx, caller, tracer)
	if err != nil {
		return nil, err
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	gasBalance := k.computeGasLimit(sdkCtx)
	rules := evm.ChainConfig().Rules(evm.Context.BlockNumber, evm.Context.Random != nil, evm.Context.Time)
	evm.StateDB.Prepare(rules, caller, types.NullAddress, &contractAddr, append(vm.ActivePrecompiles(rules), k.precompiles.toAddrs()...), accessList)

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
func (k Keeper) EVMCall(ctx context.Context, caller common.Address, contractAddr common.Address, inputBz []byte, value *uint256.Int, accessList coretype.AccessList) ([]byte, types.Logs, error) {
	return k.EVMCallWithTracer(ctx, caller, contractAddr, inputBz, value, accessList, nil)
}

// EVMCallWithTracer executes an EVM call with the given input data and tracer.
func (k Keeper) EVMCallWithTracer(ctx context.Context, caller common.Address, contractAddr common.Address, inputBz []byte, value *uint256.Int, accessList coretype.AccessList, tracer *tracing.Hooks) ([]byte, types.Logs, error) {
	ctx, evm, err := k.CreateEVM(ctx, caller, tracer)
	if err != nil {
		return nil, nil, err
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	gasBalance := k.computeGasLimit(sdkCtx)
	if value == nil {
		value = uint256.NewInt(0)
	}

	rules := evm.ChainConfig().Rules(evm.Context.BlockNumber, evm.Context.Random != nil, evm.Context.Time)
	evm.StateDB.Prepare(rules, caller, types.NullAddress, &contractAddr, append(vm.ActivePrecompiles(rules), k.precompiles.toAddrs()...), accessList)

	retBz, gasRemaining, err := evm.Call(
		vm.AccountRef(caller),
		contractAddr,
		inputBz,
		gasBalance,
		value,
	)

	// evm sometimes return 0 gasRemaining, but it's not an out of gas error.
	if gasRemaining == 0 && err != nil && err != vm.ErrOutOfGas {
		return nil, nil, types.ErrEVMCreateFailed.Wrap(err.Error())
	}

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
	stateDB := evm.StateDB.(*evmstate.StateDB)
	if err := stateDB.Commit(); err != nil {
		return nil, nil, err
	}

	retHex := hexutil.Encode(retBz)
	logs := stateDB.Logs()

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
func (k Keeper) EVMCreate(ctx context.Context, caller common.Address, codeBz []byte, value *uint256.Int, accessList coretype.AccessList) ([]byte, common.Address, types.Logs, error) {
	return k.EVMCreateWithTracer(ctx, caller, codeBz, value, nil, accessList, nil)
}

// EVMCreate creates a new contract with the given code.
func (k Keeper) EVMCreate2(ctx context.Context, caller common.Address, codeBz []byte, value *uint256.Int, salt uint64, accessList coretype.AccessList) ([]byte, common.Address, types.Logs, error) {
	return k.EVMCreateWithTracer(ctx, caller, codeBz, value, &salt, accessList, nil)
}

// EVMCreateWithTracer creates a new contract with the given code and tracer.
// if salt is nil, it will create a contract with the CREATE opcode.
// if salt is not nil, it will create a contract with the CREATE2 opcode.
func (k Keeper) EVMCreateWithTracer(ctx context.Context, caller common.Address, codeBz []byte, value *uint256.Int, salt *uint64, accessList coretype.AccessList, tracer *tracing.Hooks) (retBz []byte, contractAddr common.Address, logs types.Logs, err error) {
	ctx, evm, err := k.CreateEVM(ctx, caller, tracer)
	if err != nil {
		return nil, common.Address{}, nil, err
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	gasBalance := k.computeGasLimit(sdkCtx)
	if value == nil {
		value = uint256.NewInt(0)
	}

	rules := evm.ChainConfig().Rules(evm.Context.BlockNumber, evm.Context.Random != nil, evm.Context.Time)
	evm.StateDB.Prepare(rules, caller, types.NullAddress, nil, append(vm.ActivePrecompiles(rules), k.precompiles.toAddrs()...), accessList)

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

	// evm sometimes return 0 gasRemaining, but it's not an out of gas error.
	if gasRemaining == 0 && err != nil && err != vm.ErrOutOfGas {
		return nil, common.Address{}, nil, types.ErrEVMCreateFailed.Wrap(err.Error())
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
	stateDB := evm.StateDB.(*evmstate.StateDB)
	err = stateDB.Commit()
	if err != nil {
		return nil, common.Address{}, nil, err
	}

	retHex := hexutil.Encode(retBz)
	logs = stateDB.Logs()

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

// NextContractAddress returns the next contract address which will be created by the given caller
// in CREATE opcode.
func (k Keeper) NextContractAddress(ctx context.Context, caller common.Address) (common.Address, error) {
	stateDB, err := k.NewStateDB(ctx, nil, types.Fee{})
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
