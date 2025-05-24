package keeper

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"math/big"

	"github.com/holiman/uint256"

	storetypes "cosmossdk.io/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/state"
	coretype "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/trie/utils"

	evmstate "github.com/initia-labs/minievm/x/evm/state"
	"github.com/initia-labs/minievm/x/evm/types"
)

func (k Keeper) NewStateDB(ctx context.Context, evm *vm.EVM, fee types.Fee) (*evmstate.StateDB, error) {
	return evmstate.NewStateDB(
		// delegate gas meter to the EVM
		sdk.UnwrapSDKContext(ctx).WithGasMeter(storetypes.NewInfiniteGasMeter()), k.cdc, k.Logger(ctx),
		k.accountKeeper, k.VMStore, evm, k.ERC20Keeper().GetERC20ABI(), fee.Contract(),
	)
}

func (k Keeper) chargeIntrinsicGas(gasBalance uint64, isContractCreation bool, data []byte, list coretype.AccessList, rules params.Rules) (uint64, error) {
	intrinsicGas, err := core.IntrinsicGas(data, list, isContractCreation, rules.IsHomestead, rules.IsIstanbul, rules.IsShanghai)
	if err != nil {
		return 0, err
	}
	if gasBalance < intrinsicGas {
		return 0, fmt.Errorf("%w: have %d, want %d", core.ErrIntrinsicGas, gasBalance, intrinsicGas)
	}
	return gasBalance - intrinsicGas, nil
}

func (k Keeper) computeGasLimit(sdkCtx sdk.Context) uint64 {
	gasLimit := sdkCtx.GasMeter().Limit() - sdkCtx.GasMeter().GasConsumedToLimit()
	if sdkCtx.ExecMode() == sdk.ExecModeSimulate {
		gasLimit = k.config.ContractSimulationGasLimit
	}

	return gasLimit
}

func (k Keeper) buildDefaultBlockContext(ctx context.Context) (vm.BlockContext, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	headerHash := sdkCtx.HeaderHash()
	if len(headerHash) == 0 {
		headerHash = make([]byte, 32)
	}

	return vm.BlockContext{
		BlockNumber: big.NewInt(sdkCtx.BlockHeight()),
		Time:        uint64(sdkCtx.BlockTime().Unix()),
		Random:      (*common.Hash)(headerHash),
	}, nil
}

func (k Keeper) buildBlockContext(ctx context.Context, defaultBlockCtx vm.BlockContext, evm *vm.EVM, fee types.Fee) (vm.BlockContext, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	baseFee, err := k.baseFee(ctx, fee)
	if err != nil {
		return vm.BlockContext{}, err
	}

	return vm.BlockContext{
		BlockNumber: defaultBlockCtx.BlockNumber,
		Time:        defaultBlockCtx.Time,
		Random:      defaultBlockCtx.Random,
		BaseFee:     baseFee,
		GasLimit:    types.BlockGasLimit(sdkCtx),
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

			// increase depth manually because it is not executed by interpreter
			evm.IncreaseDepth()
			defer func() { evm.DecreaseDepth() }()

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

			// increase depth manually because it is not executed by interpreter
			evm.IncreaseDepth()
			defer func() { evm.DecreaseDepth() }()

			_, _, err = evm.Call(vm.AccountRef(a1), fee.Contract(), inputBz, 100000, uint256.NewInt(0))
			if err != nil {
				k.Logger(ctx).Warn("failed to transfer token", "error", err)
				panic(err)
			}
		},
		GetHash: func(n uint64) common.Hash {
			// use snapshot context to get block hash
			ctx := evm.StateDB.(types.StateDB).Context()
			bz, err := k.EVMBlockHashes.Get(ctx, n)
			if err != nil {
				return common.Hash{}
			}

			return common.BytesToHash(bz)
		},
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

// CreateEVM creates a new EVM instance with the given caller address and optional tracing.
// It returns:
// - context: The context used for EVM execution
// - evm: A new EVM instance configured with chain rules and state
// - cleanup: A function to rollback the tracer's stateDB changes (only needed when using tracing)
// - error: Any error that occurred during EVM creation
func (k Keeper) CreateEVM(ctx context.Context, caller common.Address, tracing *types.Tracing) (context.Context, *vm.EVM, func(), error) {
	params, err := k.Params.Get(ctx)
	if err != nil {
		return ctx, nil, func() {}, err
	}

	extraEIPs := params.ToExtraEIPs()
	fee, err := k.LoadFee(ctx, params)
	if err != nil {
		return ctx, nil, func() {}, err
	}

	// prepare SDK context for EVM execution
	ctx, err = prepareSDKContext(sdk.UnwrapSDKContext(ctx))
	if err != nil {
		return ctx, nil, func() {}, err
	}

	chainConfig := types.DefaultChainConfig(ctx)
	vmConfig := vm.Config{ExtraEips: extraEIPs, NumRetainBlockHashes: &params.NumRetainBlockHashes}

	// use default block context for chain rules in EVM creation
	defaultBlockContext, err := k.buildDefaultBlockContext(ctx)
	if err != nil {
		return ctx, nil, func() {}, err
	}

	txContext, err := k.buildTxContext(ctx, caller, fee)
	if err != nil {
		return ctx, nil, func() {}, err
	}

	// NOTE: need to check if the EVM is correctly initialized with empty context and stateDB
	evm := vm.NewEVM(
		defaultBlockContext,
		txContext,
		nil,
		chainConfig,
		vmConfig,
	)
	// customize EVM contexts and stateDB and precompiles
	evm.Context, err = k.buildBlockContext(ctx, defaultBlockContext, evm, fee)
	if err != nil {
		return ctx, nil, func() {}, err
	}
	evm.StateDB, err = k.NewStateDB(ctx, evm, fee)
	if err != nil {
		return ctx, nil, func() {}, err
	}

	rules := chainConfig.Rules(evm.Context.BlockNumber, evm.Context.Random != nil, evm.Context.Time)
	precompiles, err := k.precompiles(rules, evm.StateDB.(types.StateDB))
	if err != nil {
		return ctx, nil, func() {}, err
	}
	evm.SetPrecompiles(precompiles)

	// prepare tracing
	cleanup, err := prepareTracing(ctx, evm, tracing)
	if err != nil {
		return ctx, nil, func() {}, err
	}

	return ctx, evm, cleanup, nil
}

// prepareSDKContext prepares the SDK context for EVM execution.
// 1. sets the cosmos messages to context
// 2. checks the recursive depth and increments it (the maximum depth is 8)
func prepareSDKContext(ctx sdk.Context) (sdk.Context, error) {
	// set cosmos messages to context
	ctx = ctx.WithValue(types.CONTEXT_KEY_EXECUTE_REQUESTS, &[]types.ExecuteRequest{})

	depth := 1
	if val := ctx.Value(types.CONTEXT_KEY_RECURSIVE_DEPTH); val != nil {
		depth = val.(int) + 1
		if depth > types.MAX_RECURSIVE_DEPTH {
			return ctx, types.ErrExceedMaxRecursiveDepth
		}
	}

	// set recursive depth to context
	return ctx.WithValue(types.CONTEXT_KEY_RECURSIVE_DEPTH, depth), nil
}

// prepareTracing prepares the tracing for the EVM execution.
// 1. sets the tracer to the EVM and the stateDB to the tracing VMContext.
// 2. returns a cleanup function to rollback the tracing.
func prepareTracing(ctx context.Context, evm *vm.EVM, tracing *types.Tracing) (func(), error) {
	if tracing == nil || tracing.Tracer() == nil {
		return func() {}, nil
	}

	evm.Config.Tracer = tracing.Tracer()
	evm.StateDB.(types.StateDB).SetTracer(tracing.Tracer())

	originalStateDB := tracing.VMContext().StateDB
	tracing.VMContext().StateDB = evm.StateDB

	return func() {
		tracing.VMContext().StateDB = originalStateDB
	}, nil
}

// EVMStaticCall executes an EVM call with the given input data in static mode.
func (k Keeper) EVMStaticCall(ctx context.Context, caller common.Address, contractAddr common.Address, inputBz []byte, accessList coretype.AccessList) ([]byte, error) {
	var tracing *types.Tracing
	if sdkCtx := sdk.UnwrapSDKContext(ctx); sdkCtx.Value(types.CONTEXT_KEY_TRACING) != nil {
		tracing = sdkCtx.Value(types.CONTEXT_KEY_TRACING).(*types.Tracing)
	}

	return k.evmStaticCall(ctx, caller, contractAddr, inputBz, accessList, tracing)
}

// evmStaticCall executes an EVM call with the given input data and tracer in static mode.
func (k Keeper) evmStaticCall(ctx context.Context, caller common.Address, contractAddr common.Address, inputBz []byte, accessList coretype.AccessList, tracing *types.Tracing) ([]byte, error) {
	ctx, evm, cleanup, err := k.CreateEVM(ctx, caller, tracing)
	if err != nil {
		return nil, err
	}

	defer cleanup()

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	gasBalance := k.computeGasLimit(sdkCtx)
	rules := evm.ChainConfig().Rules(evm.Context.BlockNumber, evm.Context.Random != nil, evm.Context.Time)
	gasRemaining, err := k.chargeIntrinsicGas(gasBalance, false, inputBz, accessList, rules)
	if err != nil {
		return nil, err
	}

	if rules.IsEIP4762 {
		evm.AccessEvents.AddTxOrigin(caller)
	}
	evm.StateDB.Prepare(rules, caller, types.NullAddress, &contractAddr, k.precompileAddrs(rules), accessList)

	retBz, gasRemaining, err := evm.StaticCall(
		vm.AccountRef(caller),
		contractAddr,
		inputBz,
		gasRemaining,
	)

	// London enforced
	gasUsed := types.CalGasUsed(gasBalance, gasRemaining, evm.StateDB.GetRefund())
	consumeGas(sdkCtx, gasUsed, gasRemaining, "EVM gas consumption")
	if err != nil {
		return nil, types.ErrEVMCallFailed.Wrap(err.Error())
	}

	return retBz, nil
}

// EVMCall executes an EVM call with the given input data.
func (k Keeper) EVMCall(ctx context.Context, caller common.Address, contractAddr common.Address, inputBz []byte, value *uint256.Int, accessList coretype.AccessList) ([]byte, types.Logs, error) {
	var tracing *types.Tracing
	if sdkCtx := sdk.UnwrapSDKContext(ctx); sdkCtx.Value(types.CONTEXT_KEY_TRACING) != nil {
		tracing = sdkCtx.Value(types.CONTEXT_KEY_TRACING).(*types.Tracing)
	}

	return k.evmCall(ctx, caller, contractAddr, inputBz, value, accessList, tracing)
}

// evmCall executes an EVM call with the given input data and tracer.
func (k Keeper) evmCall(ctx context.Context, caller common.Address, contractAddr common.Address, inputBz []byte, value *uint256.Int, accessList coretype.AccessList, tracing *types.Tracing) ([]byte, types.Logs, error) {
	ctx, evm, cleanup, err := k.CreateEVM(ctx, caller, tracing)
	if err != nil {
		return nil, nil, err
	}

	defer cleanup()

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	gasBalance := k.computeGasLimit(sdkCtx)
	if value == nil {
		value = uint256.NewInt(0)
	}

	rules := evm.ChainConfig().Rules(evm.Context.BlockNumber, evm.Context.Random != nil, evm.Context.Time)
	gasRemaining, err := k.chargeIntrinsicGas(gasBalance, false, inputBz, accessList, rules)
	if err != nil {
		return nil, nil, err
	}

	if rules.IsEIP4762 {
		evm.AccessEvents.AddTxOrigin(caller)
		evm.AccessEvents.AddTxDestination(contractAddr, value.Sign() != 0)
	}

	evm.StateDB.Prepare(rules, caller, types.NullAddress, &contractAddr, k.precompileAddrs(rules), accessList)

	retBz, gasRemaining, err := evm.Call(
		vm.AccountRef(caller),
		contractAddr,
		inputBz,
		gasRemaining,
		value,
	)

	// evm sometimes return 0 gasRemaining, but it's not an out of gas error.
	switch sdkCtx.ExecMode() {
	case sdk.ExecModeSimulate, sdk.ExecModeReCheck, sdk.ExecModeCheck:
		// return exact error instead of out of gas error
		if gasRemaining == 0 && err != nil && err != vm.ErrOutOfGas {
			if err == vm.ErrExecutionReverted {
				return nil, nil, types.NewRevertError(common.CopyBytes(retBz))
			}

			return nil, nil, types.ErrEVMCallFailed.Wrap(err.Error())
		}
	default:
	}

	// London enforced
	gasUsed := types.CalGasUsed(gasBalance, gasRemaining, evm.StateDB.GetRefund())
	consumeGas(sdkCtx, gasUsed, gasRemaining, "EVM gas consumption")
	if err != nil {
		if err == vm.ErrExecutionReverted {
			return nil, nil, types.NewRevertError(common.CopyBytes(retBz))
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

	// handle cosmos execute requests
	requests := sdkCtx.Value(types.CONTEXT_KEY_EXECUTE_REQUESTS).(*[]types.ExecuteRequest)
	if dispatchLogs, err := k.dispatchMessages(sdkCtx, *requests); err != nil {
		return nil, nil, err
	} else {
		logs = append(logs, dispatchLogs...)
	}

	return retBz, logs, nil
}

// EVMCreate creates a new contract with the given code.
func (k Keeper) EVMCreate(ctx context.Context, caller common.Address, codeBz []byte, value *uint256.Int, accessList coretype.AccessList) ([]byte, common.Address, types.Logs, error) {
	var tracing *types.Tracing
	if sdkCtx := sdk.UnwrapSDKContext(ctx); sdkCtx.Value(types.CONTEXT_KEY_TRACING) != nil {
		tracing = sdkCtx.Value(types.CONTEXT_KEY_TRACING).(*types.Tracing)
	}

	return k.evmCreate(ctx, caller, codeBz, value, nil, accessList, tracing)
}

// EVMCreate2 creates a new contract with the given code.
func (k Keeper) EVMCreate2(ctx context.Context, caller common.Address, codeBz []byte, value *uint256.Int, salt *uint256.Int, accessList coretype.AccessList) ([]byte, common.Address, types.Logs, error) {
	var tracing *types.Tracing
	if sdkCtx := sdk.UnwrapSDKContext(ctx); sdkCtx.Value(types.CONTEXT_KEY_TRACING) != nil {
		tracing = sdkCtx.Value(types.CONTEXT_KEY_TRACING).(*types.Tracing)
	}

	return k.evmCreate(ctx, caller, codeBz, value, salt, accessList, tracing)
}

// evmCreate creates a new contract with the given code and tracer.
// if salt is nil, it will create a contract with the CREATE opcode.
// if salt is not nil, it will create a contract with the CREATE2 opcode.
func (k Keeper) evmCreate(ctx context.Context, caller common.Address, codeBz []byte, value *uint256.Int, salt *uint256.Int, accessList coretype.AccessList, tracing *types.Tracing) (retBz []byte, contractAddr common.Address, logs types.Logs, err error) {
	ctx, evm, cleanup, err := k.CreateEVM(ctx, caller, tracing)
	if err != nil {
		return nil, common.Address{}, nil, err
	}

	defer cleanup()

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	gasBalance := k.computeGasLimit(sdkCtx)
	if value == nil {
		value = uint256.NewInt(0)
	}

	rules := evm.ChainConfig().Rules(evm.Context.BlockNumber, evm.Context.Random != nil, evm.Context.Time)

	gasRemaining, err := k.chargeIntrinsicGas(gasBalance, true, codeBz, accessList, rules)
	if err != nil {
		return nil, common.Address{}, nil, err
	}

	if rules.IsEIP4762 {
		evm.AccessEvents.AddTxOrigin(caller)
	}

	evm.StateDB.Prepare(rules, caller, types.NullAddress, nil, k.precompileAddrs(rules), accessList)
	if salt == nil {
		retBz, contractAddr, gasRemaining, err = evm.Create(
			vm.AccountRef(caller),
			codeBz,
			gasRemaining,
			value,
		)
	} else {
		retBz, contractAddr, gasRemaining, err = evm.Create2(
			vm.AccountRef(caller),
			codeBz,
			gasRemaining,
			value,
			salt,
		)
	}

	// evm sometimes return 0 gasRemaining, but it's not an out of gas error.
	switch sdkCtx.ExecMode() {
	case sdk.ExecModeSimulate, sdk.ExecModeReCheck, sdk.ExecModeCheck:
		// return exact error instead of out of gas error
		if gasRemaining == 0 && err != nil && err != vm.ErrOutOfGas {
			if err == vm.ErrExecutionReverted {
				return nil, common.Address{}, nil, types.NewRevertError(common.CopyBytes(retBz))
			}

			return nil, common.Address{}, nil, types.ErrEVMCreateFailed.Wrap(err.Error())
		}
	default:
	}

	// London enforced
	gasUsed := types.CalGasUsed(gasBalance, gasRemaining, evm.StateDB.GetRefund())
	consumeGas(sdkCtx, gasUsed, gasRemaining, "EVM gas consumption")
	if err != nil {
		if err == vm.ErrExecutionReverted {
			return nil, common.Address{}, nil, types.NewRevertError(common.CopyBytes(retBz))
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

	// handle cosmos execute requests
	requests := sdkCtx.Value(types.CONTEXT_KEY_EXECUTE_REQUESTS).(*[]types.ExecuteRequest)
	if dispatchLogs, err := k.dispatchMessages(sdkCtx, *requests); err != nil {
		return nil, common.Address{}, nil, err
	} else {
		logs = append(logs, dispatchLogs...)
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
func (k Keeper) dispatchMessages(ctx context.Context, requests []types.ExecuteRequest) (types.Logs, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	var logs types.Logs
	for _, request := range requests {
		callLogs, err := k.dispatchMessage(sdkCtx, request)
		if err != nil {
			return nil, err
		}

		logs = append(logs, callLogs...)
	}

	return logs, nil
}

func (k Keeper) dispatchMessage(parentCtx sdk.Context, request types.ExecuteRequest) (logs types.Logs, err error) {
	msg := request.Msg
	caller := request.Caller

	allowFailure := request.AllowFailure
	callbackId := request.CallbackId

	gasLimit := request.GasLimit

	ctx, commit := parentCtx.CacheContext()
	ctx = ctx.WithGasMeter(storetypes.NewGasMeter(gasLimit))
	defer func() {
		if r := recover(); r != nil {
			switch r.(type) {
			case storetypes.ErrorOutOfGas:
				// propagate out of gas error
				panic(r)
			default:
				err = fmt.Errorf("panic: %v", r)
			}
		}

		success := err == nil

		// create submsg event
		event := sdk.NewEvent(
			types.EventTypeSubmsg,
			sdk.NewAttribute(types.AttributeKeySuccess, fmt.Sprintf("%v", success)),
		)
		// refund remaining gas
		parentCtx.GasMeter().RefundGas(ctx.GasMeter().GasRemaining(), "refund gas from submsg")
		if !success {
			// return error if failed and not allowed to fail
			if !allowFailure {
				return
			}

			// emit failed reason event if failed and allowed to fail
			event = event.AppendAttributes(sdk.NewAttribute(types.AttributeKeyReason, err.Error()))
		} else {
			// commit if success
			commit()

		}

		// reset error because it's allowed to fail
		err = nil

		// emit submessage event
		parentCtx.EventManager().EmitEvent(event)

		// if callback exists, execute it with parent context because it's already committed
		if callbackId > 0 {
			var inputBz []byte
			inputBz, err = k.cosmosCallbackABI.Pack("callback", callbackId, success)
			if err != nil {
				return
			}

			var callbackLogs types.Logs
			_, callbackLogs, err = k.EVMCall(parentCtx, caller.Address(), caller.Address(), inputBz, nil, nil)
			if err != nil {
				return
			}

			logs = append(logs, callbackLogs...)
		}
	}()

	// find the handler
	handler := k.msgRouter.Handler(msg)
	if handler == nil {
		err = types.ErrNotSupportedCosmosMessage
		return
	}

	// and execute it
	res, err := handler(ctx, msg)
	if err != nil {
		return
	}

	// emit events
	ctx.EventManager().EmitEvents(res.GetEvents())

	// extract logs
	dispatchLogs, err := types.ExtractLogsFromResponse(res.Data, sdk.MsgTypeURL(msg))
	if err != nil {
		return
	}

	// append logs
	logs = append(logs, dispatchLogs...)

	return
}

// consumeGas consumes gas
func consumeGas(ctx sdk.Context, gasUsed, gasRemaining uint64, description string) {
	// evm sometimes return 0 gasRemaining, but it's not an out of gas error.
	// cosmos use infinite gas meter at simulation and block operations.
	//
	// to prevent uint64 overflow, we don't consume gas when gas meter is infinite
	// and gasRemaining is 0.
	if ctx.GasMeter().Limit() == math.MaxUint64 && gasRemaining == 0 {
		return
	}

	ctx.GasMeter().ConsumeGas(gasUsed, description)
}
