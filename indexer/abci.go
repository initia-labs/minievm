package indexer

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"time"

	abci "github.com/cometbft/cometbft/abci/types"

	"cosmossdk.io/collections"
	storetypes "cosmossdk.io/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ethereum/go-ethereum/common"
	coretypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/trie"

	rpctypes "github.com/initia-labs/minievm/jsonrpc/types"
)

func (e *EVMIndexerImpl) ListenCommit(ctx context.Context, res abci.ResponseCommit, changeSet []*storetypes.StoreKVPair) error {
	return nil
}

// IndexBlock implements EVMIndexer.
func (e *EVMIndexerImpl) ListenFinalizeBlock(ctx context.Context, req abci.RequestFinalizeBlock, res abci.ResponseFinalizeBlock) error {
	if !e.enabled {
		return nil
	}

	// avoid passing the context passed to the abci listener to the indexing goroutine
	// because it will be cleared after the abci listener returns by the sdk.
	//
	// so load all the state dependent args before passing this to the indexing goroutine
	params, err := e.evmKeeper.Params.Get(sdk.UnwrapSDKContext(ctx))
	if err != nil {
		return fmt.Errorf("failed to get params: %w", err)
	}
	feeDecimals, err := e.evmKeeper.ERC20Keeper().GetDecimals(sdk.UnwrapSDKContext(ctx), params.FeeDenom)
	if err != nil {
		return fmt.Errorf("failed to get fee decimals: %w", err)
	}
	baseFee, err := e.evmKeeper.BaseFee(ctx)
	if err != nil {
		return fmt.Errorf("failed to load fee: %w", err)
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	blockTime := sdkCtx.BlockTime()
	blockHeight := sdkCtx.BlockHeight()
	blockGasMeter := sdkCtx.BlockGasMeter()

	// add to the indexing wait group
	e.indexingWg.Add(1)
	e.indexingChan <- &indexingTask{
		req: &req,
		res: &res,

		// state dependent args for indexing
		args: &indexingArgs{
			chainID: sdk.UnwrapSDKContext(ctx).ChainID(),

			ac:        e.ac,
			txDecoder: e.txConfig.TxDecoder(),

			params:      params,
			baseFee:     baseFee,
			feeDecimals: feeDecimals,

			blockHeight:   blockHeight,
			blockTime:     blockTime,
			blockGasMeter: blockGasMeter,
		},
	}
	return nil
}

// indexingLoop is the main loop for indexing.
func (e *EVMIndexerImpl) indexingLoop() {
	for task := range e.indexingChan {
		needBackfill, err := e.doIndexing(task.args, task.req, task.res)
		if err != nil {
			e.logger.Error("indexingLoop error", "err", err)
		} else if needBackfill {
			lastIndexedHeight, err := e.GetLastIndexedHeight(context.Background())
			if err != nil {
				e.logger.Error("failed to get last indexed height", "err", err)
				continue
			}
			err = e.Backfill(uint64(lastIndexedHeight+1), uint64(task.args.blockHeight-1))
			if err != nil {
				e.logger.Error("failed to backfill", "err", err)
				continue
			}

			// retry the indexing
			_, err = e.doIndexing(task.args, task.req, task.res)
			if err != nil {
				e.logger.Error("indexingLoop error", "err", err)
				continue
			}
		}

		// done with the indexing
		e.indexingWg.Done()
	}
}

// doIndexing is the main function for indexing.
func (e *EVMIndexerImpl) doIndexing(args *indexingArgs, req *abci.RequestFinalizeBlock, res *abci.ResponseFinalizeBlock) (needBackfill bool, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("doIndexing panic: %v", r)
		}
	}()

	// TODO: Consider removing context usage across all getter and setter methods since they are only used for consistency.
	// Currently keeping it to maintain uniform API patterns across the collections.Map interface and other storage operations.
	ctx := context.Background()
	ethTxInfos, err_ := extractEthTxInfos(e.logger, args, req, res)
	if err_ != nil {
		err = fmt.Errorf("failed to extract eth tx infos: %w", err_)
		return
	}

	txIndex := uint(0)
	cumulativeGasUsed := uint64(0)
	ethTxs := make([]*coretypes.Transaction, len(ethTxInfos))
	receipts := make([]*coretypes.Receipt, len(ethTxInfos))
	for idx, ethTxInfo := range ethTxInfos {
		ethTx := ethTxInfo.Tx
		txStatus := ethTxInfo.Status
		gasUsed := ethTxInfo.GasUsed
		cosmosTxHash := ethTxInfo.CosmosTxHash
		ethLogs := ethTxInfo.Logs
		contractAddr := ethTxInfo.ContractAddr

		// index tx hash
		if err_ := e.TxHashToCosmosTxHash.Set(ctx, ethTx.Hash().Bytes(), cosmosTxHash); err_ != nil {
			err = fmt.Errorf("failed to store tx hash to cosmos tx hash: %w", err_)
			return
		}
		if err_ := e.CosmosTxHashToTxHash.Set(ctx, cosmosTxHash, ethTx.Hash().Bytes()); err_ != nil {
			err = fmt.Errorf("failed to store cosmos tx hash to tx hash: %w", err_)
			return
		}

		cumulativeGasUsed += gasUsed

		txIndex++
		ethTxs[idx] = ethTx

		receipt := coretypes.Receipt{
			PostState:         nil,
			Status:            txStatus,
			CumulativeGasUsed: cumulativeGasUsed,
			GasUsed:           gasUsed,
			Bloom:             coretypes.Bloom(coretypes.LogsBloom(ethLogs)),
			Logs:              ethLogs,
			TransactionIndex:  txIndex,
			EffectiveGasPrice: ethTx.GasPrice(),
		}

		// currently we do not support fee refund for gas tip, so the effective gas price is the same as the gas price
		// if ethTx.Type() == coretypes.DynamicFeeTxType {
		// 	receipt.EffectiveGasPrice = new(big.Int).Add(ethTx.EffectiveGasTipValue(baseFee.ToInt()), baseFee.ToInt())
		// }

		// fill in contract address if it's a contract creation
		if contractAddr != nil {
			receipt.ContractAddress = *contractAddr
		}

		receipts[idx] = &receipt
	}

	blockGasMeter := args.blockGasMeter
	blockHeight := args.blockHeight

	// load parent hash
	parentHash := common.Hash{}
	if blockHeight > 1 {
		parentNumber := uint64(blockHeight - 1)
		parentHeader, err_ := e.BlockHeaderByNumber(ctx, parentNumber)
		if err_ != nil && errors.Is(err_, collections.ErrNotFound) {
			return true, nil
		} else if err_ != nil {
			err = fmt.Errorf("failed to get parent header: %w", err_)
			return
		}

		parentHash = parentHeader.Hash()
	}

	hasher := trie.NewStackTrie(nil)
	blockHeader := coretypes.Header{
		TxHash:      coretypes.DeriveSha(coretypes.Transactions(ethTxs), hasher),
		ReceiptHash: coretypes.DeriveSha(coretypes.Receipts(receipts), hasher),
		Bloom:       coretypes.CreateBloom(receipts),
		GasLimit:    blockGasMeter.Limit(),
		GasUsed:     blockGasMeter.GasConsumedToLimit(),
		Number:      big.NewInt(blockHeight),
		Time:        uint64(args.blockTime.Unix()),
		BaseFee:     args.baseFee,

		// empty values
		Root:            coretypes.EmptyRootHash,
		UncleHash:       coretypes.EmptyUncleHash,
		WithdrawalsHash: &coretypes.EmptyWithdrawalsHash,
		ParentHash:      parentHash,
		MixDigest:       common.Hash{},
		Difficulty:      big.NewInt(0),
		Nonce:           coretypes.EncodeNonce(0),
		Coinbase:        common.Address{},
		Extra:           []byte{},
	}

	blockHash := blockHeader.Hash()
	blockLogs := make([][]*coretypes.Log, 0, len(ethTxs))
	for idx, ethTx := range ethTxs {
		txHash := ethTx.Hash()
		receipt := receipts[idx]

		// store tx
		rpcTx := rpctypes.NewRPCTransaction(ethTx, blockHash, uint64(blockHeight), uint64(receipt.TransactionIndex), ethTx.ChainId())
		if err_ := e.TxMap.Set(ctx, txHash.Bytes(), *rpcTx); err_ != nil {
			err = fmt.Errorf("failed to store rpcTx: %w", err_)
			return
		}
		if err_ := e.TxReceiptMap.Set(ctx, txHash.Bytes(), *receipt); err_ != nil {
			err = fmt.Errorf("failed to store tx receipt: %w", err_)
			return
		}

		// store index
		if err_ := e.BlockAndIndexToTxHashMap.Set(ctx, collections.Join(uint64(blockHeight), uint64(receipt.TransactionIndex)), txHash.Bytes()); err_ != nil {
			err = fmt.Errorf("failed to store blockAndIndexToTxHash: %w", err_)
			return
		}

		// remove tx from the pending and queued after indexing
		e.pendingTxs.Delete(ethTx.Hash())
		e.queuedTxs.Delete(ethTx.Hash())

		if len(e.logsChans) > 0 && len(receipt.Logs) > 0 {
			for idx, log := range receipt.Logs {
				// fill in missing fields before emitting
				log.Index = uint(idx)
				log.BlockHash = blockHash
				log.BlockNumber = uint64(blockHeight)
				log.TxHash = txHash
				log.TxIndex = uint(receipt.TransactionIndex)
			}

			blockLogs = append(blockLogs, receipt.Logs)
		}
	}

	// index block header
	if err_ := e.BlockHeaderMap.Set(ctx, uint64(blockHeight), blockHeader); err_ != nil {
		err = fmt.Errorf("failed to marshal blockHeader: %w", err_)
		return
	}
	if err_ := e.BlockHashToNumberMap.Set(ctx, blockHash.Bytes(), uint64(blockHeight)); err_ != nil {
		err = fmt.Errorf("failed to store blockHashToNumber: %w", err_)
		return
	}

	// emit block event in a goroutine
	done := make(chan struct{})
	go e.blockEventsEmitter(&blockEvents{header: &blockHeader, logs: blockLogs}, done)
	go func() {
		select {
		case <-done:
		case <-time.After(5 * time.Second):
			e.logger.Error("block event emitter timed out")
		}
	}()

	e.lastIndexedHeight.Store(uint64(blockHeight))
	// TODO - currently state changes are not supported in abci listener, so we track cosmos block hash at x/evm preblocker.
	// - https://github.com/cosmos/cosmos-sdk/issues/22246
	//
	// err = e.evmKeeper.TrackBlockHash(sdkCtx, uint64(blockHeight), blockHash)
	// if err != nil {
	// 	e.logger.Error("failed to track block hash", "err", err)
	// 	return err
	// }

	// execute pruning only if retain height is set
	if e.retainHeight > 0 {
		e.doPrune(ctx, uint64(blockHeight))
	}

	// trigger bloom indexing
	e.doBloomIndexing(ctx, uint64(blockHeight))

	e.logger.Info("evm indexer indexed", "blockHeight", blockHeight)

	return false, nil
}
