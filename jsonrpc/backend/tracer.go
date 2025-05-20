package backend

import (
	"fmt"

	"github.com/pkg/errors"

	"cosmossdk.io/collections"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/tracing"
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

	bn := int64(blockNumber)
	cosmosBlock, err := b.clientCtx.Client.Block(ctx, &bn)
	if err != nil {
		return nil, err
	}

	var (
		cosmosTxs = cosmosBlock.Block.Data.Txs
		results   = []*rpctypes.TxTraceResult{}
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
		isCosmosTx := err == nil
		if !isCosmosTx {
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
		for _, result := range res {
			results = append(results, &rpctypes.TxTraceResult{TxHash: txHash, Result: result})
		}
		if err != nil {
			results = append(results, &rpctypes.TxTraceResult{TxHash: txHash, Error: err.Error()})
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

// traceTx configures a new tracer according to the provided configuration, and
// executes the given message in the provided environment. The return value will
// be tracer dependent.
func (b *JSONRPCBackend) traceTx(
	sdkCtx sdk.Context,
	cosmosTx sdk.Tx,
	txctx *tracers.Context,
	config *tracers.TraceConfig,
) ([]any, error) {
	var (
		tracer *tracers.Tracer
		err    error
	)
	if config == nil {
		config = &tracers.TraceConfig{}
	}

	tracerFn := func() (*tracers.Tracer, error) {
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

		return tracer, nil
	}

	// check tracerFn returns a valid tracer
	if _, err := tracerFn(); err != nil {
		return nil, err
	}

	var tracers []*tracers.Tracer
	tracerGenerator := func() *tracing.Hooks {
		tracer, _ := tracerFn()
		tracers = append(tracers, tracer)
		return tracer.Hooks
	}

	execErr := b.runTxWithTracer(sdkCtx, cosmosTx, tracerGenerator)
	results := make([]any, len(tracers))
	for i, tracer := range tracers {
		result, err := tracer.GetResult()
		if err != nil {
			if execErr != nil {
				return nil, errors.Wrap(err, execErr.Error())
			}

			return nil, err
		}

		results[i] = result
	}

	return results, nil
}

func (b *JSONRPCBackend) runTxWithTracer(
	sdkCtx sdk.Context,
	cosmosTx sdk.Tx,
	tracerGenerator evmtypes.TracingHooks,
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
	_, err = b.app.MsgServiceRouter().Handler(msg)(sdkCtx.WithValue(evmtypes.CONTEXT_KEY_TRACER, tracerGenerator), msg)
	_, postErr := b.app.PostHandler()(sdkCtx, cosmosTx, false, err == nil)
	if err == nil && postErr != nil {
		err = postErr
	}

	return err
}
