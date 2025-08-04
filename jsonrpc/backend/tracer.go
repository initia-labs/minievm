package backend

import (
	"context"
	"fmt"

	"github.com/pkg/errors"

	"cosmossdk.io/collections"
	storetypes "cosmossdk.io/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	coretypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/eth/tracers"
	"github.com/ethereum/go-ethereum/eth/tracers/logger"
	_ "github.com/ethereum/go-ethereum/eth/tracers/native"
	"github.com/ethereum/go-ethereum/rpc"

	rpctypes "github.com/initia-labs/minievm/jsonrpc/types"
	"github.com/initia-labs/minievm/x/evm/state"
	evmtypes "github.com/initia-labs/minievm/x/evm/types"
)

// TraceBlockByNumber configures a new tracer according to the provided configuration, and
// executes all the transactions contained within. The return value will be one item
// per transaction, dependent on the requested tracer.
func (b *JSONRPCBackend) TraceBlockByNumber(ethBlockNum rpc.BlockNumber, config *tracers.TraceConfig) ([]*rpctypes.TxTraceResult, error) {
	blockNumber, err := b.resolveBlockNr(ethBlockNum)
	if err != nil {
		return nil, err
	} else if blockNumber < 2 {
		return nil, errors.New("genesis is not traceable")
	}

	ctx, err := b.getQueryCtxWithHeight(blockNumber - 1)
	if err != nil {
		return nil, err
	}

	header, err := b.GetHeaderByNumber(ethBlockNum)
	if err != nil {
		return nil, err
	} else if header == nil {
		return nil, fmt.Errorf("block #%d not found", ethBlockNum)
	}

	bn := int64(blockNumber)
	cosmosBlock, err := b.clientCtx.Client.Block(ctx, &bn)
	if err != nil {
		return nil, err
	}

	var (
		cosmosTxs = cosmosBlock.Block.Data.Txs
		results   = make([]*rpctypes.TxTraceResult, 0, len(cosmosTxs))
	)

	idx := 0
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	for _, cosmosTxBytes := range cosmosTxs {
		cosmosTxHash := cosmosTxBytes.Hash()
		cosmosTx, err := b.app.TxDecode(cosmosTxBytes)
		if err != nil {
			return nil, err
		}

		// If the tx is not evm tx or cosmos tx which is not indexed, we need to run it without tracer
		// and continue to the next tx
		txHash, err := b.TxHashByCosmosTxHash(cosmosTxHash)
		if err != nil {
			return nil, err
		}
		if txHash == (common.Hash{}) {
			_ = b.runTxWithTracer(sdkCtx, cosmosTx, nil)
			continue
		}

		// Generate the next state snapshot fast without tracing
		txctx := &tracers.Context{
			BlockHash:   header.Hash(),
			BlockNumber: header.Number,
			TxIndex:     idx,
			TxHash:      txHash,
		}
		res, err := b.traceTx(sdkCtx, cosmosTx, txctx, config)
		results = append(results, &rpctypes.TxTraceResult{TxHash: txHash, Result: res})
		if err != nil {
			results[len(results)-1].Error = err.Error()
		}

		idx++
	}

	return results, nil
}

// TraceBlockByHash configures a new tracer according to the provided configuration, and
// executes all the transactions contained within. The return value will be one item
// per transaction, dependent on the requested tracer.
func (b *JSONRPCBackend) TraceBlockByHash(hash common.Hash, config *tracers.TraceConfig) ([]*rpctypes.TxTraceResult, error) {
	blockNumber, err := b.resolveBlockNrOrHash(rpc.BlockNumberOrHash{BlockHash: &hash})
	if err != nil && errors.Is(err, collections.ErrNotFound) {
		return nil, nil
	} else if err != nil {
		b.logger.Error("failed to get block number by hash", "err", err)
		return nil, err
	}
	return b.TraceBlockByNumber(rpc.BlockNumber(blockNumber), config)
}

func (b *JSONRPCBackend) TraceTransaction(hash common.Hash, config *tracers.TraceConfig) (any, error) {
	// check if the tx is indexed
	tx, err := b.app.EVMIndexer().TxByHash(b.ctx, hash)
	if err != nil {
		return nil, err
	} else if tx == nil {
		return nil, errors.New("transaction not found")
	} else if tx.BlockNumber == nil || tx.TransactionIndex == nil || tx.BlockHash == nil {
		return nil, errors.New("transaction is not indexed")
	}

	// check if the block is traceable
	blockNumber := tx.BlockNumber.ToInt().Uint64()
	if blockNumber < 2 {
		return nil, errors.New("genesis is not traceable")
	}
	ctx, err := b.getQueryCtxWithHeight(blockNumber - 1)
	if err != nil {
		return nil, err
	}

	// load the block
	bn := int64(blockNumber)
	cosmosBlock, err := b.clientCtx.Client.Block(ctx, &bn)
	if err != nil {
		return nil, err
	}

	// iterate over the txs in the block
	cosmosTxs := cosmosBlock.Block.Data.Txs
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	for _, cosmosTxBytes := range cosmosTxs {
		cosmosTxHash := cosmosTxBytes.Hash()
		cosmosTx, err := b.app.TxDecode(cosmosTxBytes)
		if err != nil {
			return nil, err
		}

		// until the target tx is found, run the tx without tracing
		txHash, err := b.TxHashByCosmosTxHash(cosmosTxHash)
		if err != nil {
			return nil, err
		}
		if txHash != hash {
			_ = b.runTxWithTracer(sdkCtx, cosmosTx, nil)
			continue
		}

		// execute the target tx
		txctx := &tracers.Context{
			BlockHash:   *tx.BlockHash,
			BlockNumber: tx.BlockNumber.ToInt(),
			TxIndex:     int(*tx.TransactionIndex),
			TxHash:      txHash,
		}
		return b.traceTx(sdkCtx, cosmosTx, txctx, config)
	}

	// return an error if the transaction is not found in the block
	return nil, errors.New("transaction not found in the block")
}

func (b *JSONRPCBackend) StorageRangeAt(blockNrOrHash rpc.BlockNumberOrHash, txIndex int, contractAddress common.Address, keyStart hexutil.Bytes, maxResult int) (rpctypes.StorageRangeResult, error) {
	blockNumber, err := b.resolveBlockNrOrHash(blockNrOrHash)
	if err != nil {
		return rpctypes.StorageRangeResult{}, err
	} else if blockNumber < 2 {
		return rpctypes.StorageRangeResult{}, errors.New("genesis is not traceable")
	}

	// check if the transaction is indexed
	tx, err := b.GetTransactionByBlockNumberAndIndex(rpc.BlockNumber(blockNumber), hexutil.Uint(txIndex))
	if err != nil {
		return rpctypes.StorageRangeResult{}, err
	} else if tx == nil {
		return rpctypes.StorageRangeResult{}, errors.New("transaction not found in the block")
	}

	// load the state snapshot
	ctx, err := b.getQueryCtxWithHeight(blockNumber - 1)
	if err != nil {
		return rpctypes.StorageRangeResult{}, err
	}

	// load the block
	bn := int64(blockNumber)
	cosmosBlock, err := b.clientCtx.Client.Block(ctx, &bn)
	if err != nil {
		return rpctypes.StorageRangeResult{}, err
	}

	// iterate over the txs in the block
	cosmosTxs := cosmosBlock.Block.Data.Txs
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	found := false
	for _, cosmosTxBytes := range cosmosTxs {
		cosmosTxHash := cosmosTxBytes.Hash()
		cosmosTx, err := b.app.TxDecode(cosmosTxBytes)
		if err != nil {
			return rpctypes.StorageRangeResult{}, err
		}

		// run the tx without tracing until the target tx is found
		txHash, err := b.TxHashByCosmosTxHash(cosmosTxHash)
		if err != nil {
			return rpctypes.StorageRangeResult{}, err
		}
		if txHash != tx.Hash {
			_ = b.runTxWithTracer(sdkCtx, cosmosTx, nil)
			continue
		}

		found = true
		break
	}

	// return an error if the transaction is not found in the block
	if !found {
		return rpctypes.StorageRangeResult{}, errors.New("transaction not found in the block")
	}

	// extract the slot from the key
	extractSlot := func(key []byte) (slot common.Hash) {
		prefixLen := 20 + len(state.StateKeyPrefix)
		copy(slot[:], key[prefixLen:])
		return slot
	}

	result := rpctypes.StorageRangeResult{Storage: rpctypes.StorageMap{}}
	prefix := append(contractAddress.Bytes(), state.StateKeyPrefix...)
	startKey := append(prefix, keyStart...)
	iter, err := b.app.EVMKeeper.VMStore.Iterate(ctx, new(collections.Range[[]byte]).Prefix(prefix).StartInclusive(startKey))
	if err != nil {
		return rpctypes.StorageRangeResult{}, err
	}

	for i := 0; i < maxResult && iter.Valid(); i++ {
		keyValue, err := iter.KeyValue()
		if err != nil {
			return rpctypes.StorageRangeResult{}, err
		}
		key := keyValue.Key
		slot := extractSlot(key)
		content := keyValue.Value
		e := rpctypes.StorageEntry{Key: &slot, Value: common.BytesToHash(content)}
		result.Storage[slot] = e
		iter.Next()
	}

	if iter.Valid() {
		nextKey, err := iter.Key()
		if err != nil {
			result.NextKey = nil
		} else {
			nextSlot := extractSlot(nextKey)
			result.NextKey = &nextSlot
		}
	}

	return result, nil
}

// traceTx configures a new tracer according to the provided configuration, and
// executes the given message in the provided environment. The return value will
// be tracer dependent.
func (b *JSONRPCBackend) traceTx(
	sdkCtx sdk.Context,
	cosmosTx sdk.Tx,
	txctx *tracers.Context,
	config *tracers.TraceConfig,
) (any, error) {
	var (
		tracer *tracers.Tracer
		err    error
	)
	if config == nil {
		config = &tracers.TraceConfig{}
	}

	// Default tracer is the struct logger
	if config.Tracer == nil {
		logger := logger.NewStructLogger(config.Config)
		tracer = &tracers.Tracer{
			Hooks:     logger.Hooks(),
			GetResult: logger.GetResult,
			Stop:      logger.Stop,
		}
	} else {
		tracer, err = tracers.DefaultDirectory.New(*config.Tracer, txctx, config.TracerConfig)
		if err != nil {
			return nil, err
		}
	}

	execErr := b.runTxWithTracer(sdkCtx, cosmosTx, tracer)
	result, err := tracer.GetResult()
	if err != nil {
		if execErr != nil {
			return nil, errors.Wrap(err, execErr.Error())
		}

		return nil, err
	}

	return result, nil
}

func (b *JSONRPCBackend) runTxWithTracer(
	sdkCtx sdk.Context,
	cosmosTx sdk.Tx,
	tracer *tracers.Tracer,
) (err error) {
	feeTx := cosmosTx.(sdk.FeeTx)
	gasLimit := feeTx.GetGas()
	sdkCtx = sdkCtx.WithGasMeter(storetypes.NewGasMeter(gasLimit)).WithExecMode(sdk.ExecModeFinalize)

	// setup tracing
	// execute OnTxStart and dummy OnEnter
	if tracer != nil {
		evmPointer := new(*vm.EVM)
		deadlineCtx, cancel := context.WithTimeout(sdkCtx, b.cfg.TracerTimeout)
		go func() {
			<-deadlineCtx.Done()
			if errors.Is(deadlineCtx.Err(), context.DeadlineExceeded) {
				tracer.Stop(errors.New("execution timeout"))
				// Stop evm execution. Note cancellation is not necessarily immediate.
				if *evmPointer != nil {
					(*evmPointer).Cancel()
				}
			}
		}()
		defer cancel()

		feePayer := common.BytesToAddress(feeTx.FeePayer())
		_, evm, _, err := b.app.EVMKeeper.CreateEVM(sdkCtx, feePayer)
		if err != nil {
			return err
		}

		tracing := evmtypes.NewTracing(evm, tracer.Hooks)
		sdkCtx = sdkCtx.WithValue(evmtypes.CONTEXT_KEY_TRACING, tracing)
		sdkCtx = sdkCtx.WithValue(evmtypes.CONTEXT_KEY_TRACE_EVM, evmPointer)

		if tracer.OnTxStart != nil {
			tracer.OnTxStart(tracing.VMContext(), evmtypes.TracingTx(gasLimit), feePayer)
		}
		if tracer.OnEnter != nil {
			tracer.OnEnter(0, byte(vm.CALL), evmtypes.NullAddress, evmtypes.NullAddress, []byte{}, gasLimit, nil)
		}
	}

	// ante handler state changes should be applied always
	sdkCtx, err = b.app.AnteHandler()(sdkCtx, cosmosTx, false)
	if err != nil {
		return err
	}

	// create cache context for message handler and post handler
	sdkCtx, commit := sdkCtx.CacheContext()
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic: %v", r)
		}

		if err == nil {
			commit()
		}
	}()

	// run msgs with post handler
	for _, msg := range cosmosTx.GetMsgs() {
		_, err = b.app.MsgServiceRouter().Handler(msg)(sdkCtx, msg)
		if err != nil {
			break
		}
	}
	_, postErr := b.app.PostHandler()(sdkCtx, cosmosTx, false, err == nil)
	if err == nil && postErr != nil {
		err = postErr
	}

	// execute dummy OnExit and OnTxEnd
	if tracer != nil {
		gasUsed := sdkCtx.GasMeter().GasConsumedToLimit()
		if tracer.OnExit != nil {
			if revertErr, ok := err.(*evmtypes.RevertError); ok {
				tracer.OnExit(0, revertErr.Ret(), gasUsed, vm.ErrExecutionReverted, true)
			} else {
				tracer.OnExit(0, nil, gasUsed, err, false)
			}
		}
		if tracer.OnTxEnd != nil {
			tracer.OnTxEnd(&coretypes.Receipt{GasUsed: gasUsed}, err)
		}
	}

	return err
}
