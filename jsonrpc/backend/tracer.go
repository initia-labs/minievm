package backend

import (
	"fmt"

	"github.com/holiman/uint256"
	"github.com/pkg/errors"

	"cosmossdk.io/collections"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/tracing"
	coretypes "github.com/ethereum/go-ethereum/core/types"
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
	}

	rpcTxs, err := b.getBlockTransactions(blockNumber)
	if err != nil {
		return nil, err
	} else if len(rpcTxs) == 0 {
		return nil, nil
	}

	var (
		results = make([]*rpctypes.TxTraceResult, len(rpcTxs))
	)

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	for i, rpcTx := range rpcTxs {
		tx := rpcTx.ToTransaction()

		// Generate the next state snapshot fast without tracing
		txctx := &tracers.Context{
			BlockHash:   header.Hash(),
			BlockNumber: header.Number,
			TxIndex:     i,
			TxHash:      tx.Hash(),
		}
		res, err := b.traceTx(sdkCtx, tx, txctx, config)
		results[i] = &rpctypes.TxTraceResult{TxHash: tx.Hash(), Result: res}
		if err != nil {
			results[i].Error = err.Error()
		}
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

func (b *JSONRPCBackend) TraceTransaction(hash common.Hash, config *tracers.TraceConfig) (*rpctypes.TxTraceResult, error) {
	queryCtx, err := b.getQueryCtx()
	if err != nil {
		return nil, err
	}

	tx, err := b.app.EVMIndexer().TxByHash(queryCtx, hash)
	if err != nil {
		return nil, err
	} else if tx == nil {
		return nil, errors.New("transaction not found")
	}

	blockNumber := tx.BlockNumber.ToInt().Uint64()

	ctx, err := b.getQueryCtxWithHeight(blockNumber - 1)
	if err != nil {
		return nil, err
	}

	header, err := b.app.EVMIndexer().BlockHeaderByNumber(ctx, blockNumber)
	if err != nil {
		return nil, err
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	txIndex := tx.TransactionIndex
	ethTx := tx.ToTransaction()
	txctx := &tracers.Context{
		BlockHash:   header.Hash(),
		BlockNumber: header.Number,
		TxHash:      ethTx.Hash(),
		TxIndex:     int(*txIndex),
	}
	res, err := b.traceTx(sdkCtx, ethTx, txctx, config)
	if err != nil {
		return nil, err
	}

	return &rpctypes.TxTraceResult{TxHash: ethTx.Hash(), Result: res}, nil
}

func (b *JSONRPCBackend) StorageRangeAt(blockNrOrHash rpc.BlockNumberOrHash, txIndex int, contractAddress common.Address, keyStart hexutil.Bytes, maxResult int) (rpctypes.StorageRangeResult, error) {
	blockNumber, err := b.resolveBlockNrOrHash(blockNrOrHash)
	if err != nil {
		return rpctypes.StorageRangeResult{}, err
	}

	traceCtx, err := b.getQueryCtxWithHeight(blockNumber - 1)
	if err != nil {
		return rpctypes.StorageRangeResult{}, err
	}

	rpcTxs, err := b.getBlockTransactions(blockNumber)
	if err != nil {
		return rpctypes.StorageRangeResult{}, err
	} else if len(rpcTxs) == 0 {
		return rpctypes.StorageRangeResult{}, nil
	}

	// replay all transactions in the block before the given txIndex
	for idx, rpcTx := range rpcTxs {
		if idx >= txIndex {
			break
		}
		value, _ := uint256.FromBig(rpcTx.Value.ToInt())
		if rpcTx.To != nil && *rpcTx.To != (common.Address{}) {
			_, _, err = b.app.EVMKeeper.EVMCall(traceCtx, rpcTx.From, *rpcTx.To, rpcTx.Input, value, *rpcTx.Accesses)
		} else {
			_, _, _, err = b.app.EVMKeeper.EVMCreate(traceCtx, rpcTx.From, rpcTx.Input, value, *rpcTx.Accesses)
		}
		if err != nil {
			return rpctypes.StorageRangeResult{}, err
		}
	}

	parseStateKey := func(key []byte) (addr common.Address, slot common.Hash) {
		copy(addr[:], key[:20])
		prefixLen := 20 + len(state.StateKeyPrefix)
		copy(slot[:], key[prefixLen:])
		return addr, slot
	}

	result := rpctypes.StorageRangeResult{Storage: rpctypes.StorageMap{}}
	prefix := append(contractAddress.Bytes(), state.StateKeyPrefix...)
	startKey := append(prefix, keyStart...)
	iter, err := b.app.EVMKeeper.VMStore.Iterate(traceCtx, new(collections.Range[[]byte]).Prefix(prefix).StartInclusive(startKey))
	if err != nil {
		return rpctypes.StorageRangeResult{}, err
	}

	for i := 0; i < maxResult && iter.Valid(); i++ {
		keyValue, err := iter.KeyValue()
		if err != nil {
			return rpctypes.StorageRangeResult{}, err
		}
		key := keyValue.Key
		_, slot := parseStateKey(key)
		content := keyValue.Value
		e := rpctypes.StorageEntry{Value: common.BytesToHash(content)}
		result.Storage[slot] = e
		iter.Next()
	}

	if iter.Valid() {
		nextKey, err := iter.Key()
		if err != nil {
			result.NextKey = nil
		} else {
			_, nextSlot := parseStateKey(nextKey)
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
	tx *coretypes.Transaction,
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

	cosmosTx, err := b.app.EVMKeeper.TxUtils().ConvertEthereumTxToCosmosTx(sdkCtx, tx)
	if err != nil {
		return nil, err
	}

	execErr := b.runTxWithTracer(sdkCtx, cosmosTx, tracer.Hooks)
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
	tracer *tracing.Hooks,
) (err error) {
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

	// run msg with post handler
	msg := cosmosTx.GetMsgs()[0]
	_, err = b.app.MsgServiceRouter().Handler(msg)(sdkCtx.WithValue(evmtypes.CONTEXT_KEY_TRACER, tracer), msg)
	_, postErr := b.app.PostHandler()(sdkCtx, cosmosTx, false, err == nil)
	if err == nil && postErr != nil {
		err = postErr
	}

	return err
}
