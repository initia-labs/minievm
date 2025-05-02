package backend

import (
	"errors"
	"fmt"

	"cosmossdk.io/collections"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/tracing"
	coretypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/eth/tracers"
	"github.com/ethereum/go-ethereum/eth/tracers/logger"
	_ "github.com/ethereum/go-ethereum/eth/tracers/native"
	"github.com/ethereum/go-ethereum/rpc"

	rpctypes "github.com/initia-labs/minievm/jsonrpc/types"
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
	for i, rcpTx := range rpcTxs {
		tx := rcpTx.ToTransaction()

		// Generate the next state snapshot fast without tracing
		txctx := &tracers.Context{
			BlockHash:   header.Hash(),
			BlockNumber: header.Number,
			TxIndex:     i,
			TxHash:      tx.Hash(),
		}
		res, err := b.traceTx(sdkCtx, tx, txctx, config)
		if err != nil {
			return nil, err
		}
		results[i] = &rpctypes.TxTraceResult{TxHash: tx.Hash(), Result: res}
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

	b.runTxWithTracer(sdkCtx, cosmosTx, tracer.Hooks)
	return tracer.GetResult()
}

func (b *JSONRPCBackend) runTxWithTracer(
	sdkCtx sdk.Context,
	cosmosTx sdk.Tx,
	tracer *tracing.Hooks,
) {
	// ante handler state changes should be applied always
	sdkCtx, err := b.app.AnteHandler()(sdkCtx, cosmosTx, false)
	if err != nil {
		return
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
	_, err = b.app.PostHandler()(sdkCtx, cosmosTx, false, err == nil)
	return
}
