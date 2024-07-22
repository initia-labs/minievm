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

	// compute base fee from the opChild gas prices
	baseFee, err := e.baseFee(ctx)
	if err != nil {
		e.logger.Error("failed to get base fee", "err", err)
		return err
	}

	txIndex := uint(0)
	cumulativeGasUsed := uint64(0)
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

		txResult := res.TxResults[idx]
		txStatus := coretypes.ReceiptStatusSuccessful
		if txResult.Code != abci.CodeTypeOK {
			txStatus = coretypes.ReceiptStatusFailed
		}

		gasUsed := uint64(txResult.GasUsed)
		cumulativeGasUsed += gasUsed

		txIndex++
		ethTxs = append(ethTxs, ethTx)

		// extract logs and contract address from tx results
		ethLogs, contractAddr, err := extractLogsAndContractAddr(txStatus, txResult.Data, ethTx.To() == nil)
		if err != nil {
			e.logger.Error("failed to extract logs and contract address", "err", err)
			return err
		}

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

		// currently we do not support fee refund, so the effective gas price is the same as the gas price
		// if ethTx.Type() == coretypes.DynamicFeeTxType {
		// 	receipt.EffectiveGasPrice = new(big.Int).Add(ethTx.EffectiveGasTipValue(baseFee.ToInt()), baseFee.ToInt())
		// }

		// fill in contract address if it's a contract creation
		if contractAddr != nil {
			receipt.ContractAddress = *contractAddr
		}

		receipts = append(receipts, &receipt)
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
		BaseFee:     baseFee.ToInt(),

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
