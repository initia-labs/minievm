package indexer

import (
	"context"
	"math/big"

	abci "github.com/cometbft/cometbft/abci/types"

	"cosmossdk.io/collections"
	storetypes "cosmossdk.io/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ethereum/go-ethereum/common"
	coretypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/trie"

	rpctypes "github.com/initia-labs/minievm/jsonrpc/types"
	"github.com/initia-labs/minievm/x/evm/keeper"
	"github.com/initia-labs/minievm/x/evm/types"
)

func (e *EVMIndexerImpl) ListenCommit(ctx context.Context, res abci.ResponseCommit, changeSet []*storetypes.StoreKVPair) error {
	e.store.Write()
	return nil
}

// IndexBlock implements EVMIndexer.
func (e *EVMIndexerImpl) ListenFinalizeBlock(ctx context.Context, req abci.RequestFinalizeBlock, res abci.ResponseFinalizeBlock) error {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	batch := e.db.NewBatch()
	defer batch.Close()

	txIndex := uint(0)
	usedGas := uint64(0)
	ethTxs := make([]*coretypes.Transaction, 0, len(req.Txs))
	receipts := make([]*coretypes.Receipt, 0, len(req.Txs))
	for idx, txBytes := range req.Txs {
		tx, err := e.txConfig.TxDecoder()(txBytes)
		if err != nil {
			e.logger.Error("failed to decode tx", "err", err)
			continue
		}

		ethTx, _, err := keeper.NewTxUtils(e.evmKeeper).ConvertCosmosTxToEthereumTx(sdkCtx, tx)
		if err != nil {
			e.logger.Error("failed to convert CosmosTx to EthTx", "err", err)
			return err
		}
		if ethTx == nil {
			continue
		}

		txIndex++
		usedGas += ethTx.Gas()

		txResults := res.TxResults[idx]
		txStatus := coretypes.ReceiptStatusSuccessful
		if txResults.Code != abci.CodeTypeOK {
			txStatus = coretypes.ReceiptStatusFailed
		}

		ethTxs = append(ethTxs, ethTx)
		ethLogs := e.extractLogsFromEvents(txResults.Events)
		receipts = append(receipts, &coretypes.Receipt{
			PostState:         nil,
			Status:            txStatus,
			CumulativeGasUsed: usedGas,
			Bloom:             coretypes.Bloom(coretypes.LogsBloom(ethLogs)),
			Logs:              ethLogs,
			TransactionIndex:  txIndex,
		})
	}

	chainId := types.ConvertCosmosChainIDToEthereumChainID(sdkCtx.ChainID())
	blockGasMeter := sdkCtx.BlockGasMeter()
	blockHeight := sdkCtx.BlockHeight()

	hasher := trie.NewStackTrie(nil)
	blockHeader := coretypes.Header{
		TxHash:      coretypes.DeriveSha(coretypes.Transactions(ethTxs), hasher),
		ReceiptHash: coretypes.DeriveSha(coretypes.Receipts(receipts), hasher),
		Bloom:       coretypes.CreateBloom(receipts),
		GasLimit:    blockGasMeter.Limit(),
		GasUsed:     blockGasMeter.GasConsumedToLimit(),
		Number:      big.NewInt(blockHeight),
		Time:        uint64(sdkCtx.BlockTime().Unix()),

		// empty values
		Root:            coretypes.EmptyRootHash,
		UncleHash:       coretypes.EmptyUncleHash,
		WithdrawalsHash: &coretypes.EmptyWithdrawalsHash,
		ParentHash:      common.Hash{},
		MixDigest:       common.Hash{},
		Difficulty:      big.NewInt(0),
		Nonce:           coretypes.EncodeNonce(0),
		Coinbase:        common.Address{},
		Extra:           []byte{},
	}

	blockHash := blockHeader.Hash()
	for txIndex, ethTx := range ethTxs {
		txHash := ethTx.Hash()
		receipt := receipts[txIndex]

		// store tx
		rpcTx := rpctypes.NewRPCTransaction(ethTx, blockHash, uint64(blockHeight), uint64(receipt.TransactionIndex), chainId)
		if err := e.TxMap.Set(ctx, txHash.Bytes(), *rpcTx); err != nil {
			e.logger.Error("failed to store rpcTx", "err", err)
			return err
		}
		if err := e.TxReceiptMap.Set(ctx, txHash.Bytes(), *receipt); err != nil {
			e.logger.Error("failed to store tx receipt", "err", err)
			return err
		}

		// store index
		if err := e.BlockAndIndexToTxHashMap.Set(ctx, collections.Join(uint64(blockHeight), uint64(receipt.TransactionIndex)), txHash.Bytes()); err != nil {
			e.logger.Error("failed to store blockAndIndexToTxHash", "err", err)
			return err
		}

		// emit log events
		if e.logsChan != nil {
			for idx, log := range receipt.Logs {
				// fill in missing fields before emitting
				log.Index = uint(idx)
				log.BlockHash = blockHash
				log.BlockNumber = uint64(blockHeight)
				log.TxHash = txHash
				log.TxIndex = uint(txIndex)
			}

			// emit logs event
			e.logsChan <- receipt.Logs
		}
	}

	// index block header
	if err := e.BlockHeaderMap.Set(ctx, uint64(blockHeight), blockHeader); err != nil {
		e.logger.Error("failed to marshal blockHeader", "err", err)
		return err
	}
	if err := e.BlockHashToNumberMap.Set(ctx, blockHash.Bytes(), uint64(blockHeight)); err != nil {
		e.logger.Error("failed to store blockHashToNumber", "err", err)
		return err
	}

	// emit new block events
	if e.blockChan != nil {
		e.blockChan <- &blockHeader
	}

	return nil
}
