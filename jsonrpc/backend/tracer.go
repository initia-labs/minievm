package backend

import (
	"fmt"

	"github.com/pkg/errors"

	"cosmossdk.io/collections"
	storetypes "cosmossdk.io/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/tracing"
	coretypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
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
	feeTx := cosmosTx.(sdk.FeeTx)
	gasLimit := feeTx.GetGas()
	sdkCtx = sdkCtx.WithGasMeter(storetypes.NewGasMeter(gasLimit)).WithExecMode(sdk.ExecModeFinalize)

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

	// setup tracing
	// execute OnTxStart and dummy OnEnter
	if tracer != nil {
		feePayer := common.BytesToAddress(feeTx.FeePayer())
		_, evm, _, err := b.app.EVMKeeper.CreateEVM(sdkCtx, feePayer, nil)
		if err != nil {
			return err
		}

		tracing := evmtypes.NewTracing(evm, tracer)
		sdkCtx = sdkCtx.WithValue(evmtypes.CONTEXT_KEY_TRACING, tracing)

		if tracer.OnTxStart != nil {
			tracer.OnTxStart(tracing.VMContext(), evmtypes.TracingTx(gasLimit), feePayer)
		}
		if tracer.OnEnter != nil {
			tracer.OnEnter(0, byte(vm.CALL), evmtypes.NullAddress, evmtypes.NullAddress, []byte{}, gasLimit, nil)
		}
	}

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
