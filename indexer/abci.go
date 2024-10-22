package indexer

import (
	"context"
	"math/big"

	abci "github.com/cometbft/cometbft/abci/types"
	comettypes "github.com/cometbft/cometbft/types"

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

	// load base fee from evm keeper
	baseFee, err := e.evmKeeper.BaseFee(sdkCtx)
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

		// index tx hash
		cosmosTxHash := comettypes.Tx(txBytes).Hash()
		if err := e.TxHashToCosmosTxHash.Set(sdkCtx, ethTx.Hash().Bytes(), cosmosTxHash); err != nil {
			e.logger.Error("failed to store tx hash to cosmos tx hash", "err", err)
			return err
		}
		if err := e.CosmosTxHashToTxHash.Set(sdkCtx, cosmosTxHash, ethTx.Hash().Bytes()); err != nil {
			e.logger.Error("failed to store cosmos tx hash to tx hash", "err", err)
			return err
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

		// currently we do not support fee refund for gas tip, so the effective gas price is the same as the gas price
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
		BaseFee:     baseFee,

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

	// fill parent hash
	if blockHeight > 1 {
		parentHeader, err := e.BlockHeaderMap.Get(sdkCtx, uint64(blockHeight-1))
		if err == nil {
			blockHeader.ParentHash = parentHeader.Hash()
		}
	}

	blockHash := blockHeader.Hash()
	for txIndex, ethTx := range ethTxs {
		txHash := ethTx.Hash()
		receipt := receipts[txIndex]

		// store tx
		rpcTx := rpctypes.NewRPCTransaction(ethTx, blockHash, uint64(blockHeight), uint64(receipt.TransactionIndex), chainId)
		if err := e.TxMap.Set(sdkCtx, txHash.Bytes(), *rpcTx); err != nil {
			e.logger.Error("failed to store rpcTx", "err", err)
			return err
		}
		if err := e.TxReceiptMap.Set(sdkCtx, txHash.Bytes(), *receipt); err != nil {
			e.logger.Error("failed to store tx receipt", "err", err)
			return err
		}

		// store index
		if err := e.BlockAndIndexToTxHashMap.Set(sdkCtx, collections.Join(uint64(blockHeight), uint64(receipt.TransactionIndex)), txHash.Bytes()); err != nil {
			e.logger.Error("failed to store blockAndIndexToTxHash", "err", err)
			return err
		}

		if len(e.logsChans) > 0 && len(receipt.Logs) > 0 {
			for idx, log := range receipt.Logs {
				// fill in missing fields before emitting
				log.Index = uint(idx)
				log.BlockHash = blockHash
				log.BlockNumber = uint64(blockHeight)
				log.TxHash = txHash
				log.TxIndex = uint(txIndex)
			}

			// emit logs event
			for _, logsChan := range e.logsChans {
				logsChan <- receipt.Logs
			}
		}
	}

	// emit empty logs event to confirm all logs are emitted and consumed, so the logs are
	// available for querying.
	if len(e.logsChans) > 0 {
		for _, logsChan := range e.logsChans {
			logsChan <- []*coretypes.Log{}
		}
	}

	// index block header
	if err := e.BlockHeaderMap.Set(sdkCtx, uint64(blockHeight), blockHeader); err != nil {
		e.logger.Error("failed to marshal blockHeader", "err", err)
		return err
	}
	if err := e.BlockHashToNumberMap.Set(sdkCtx, blockHash.Bytes(), uint64(blockHeight)); err != nil {
		e.logger.Error("failed to store blockHashToNumber", "err", err)
		return err
	}

	// emit new block events
	if len(e.blockChans) > 0 {
		for _, blockChan := range e.blockChans {
			blockChan <- &blockHeader
		}
	}

	// TODO - currently state changes are not supported in abci listener, so we track cosmos block hash at x/evm preblocker.
	// - https://github.com/cosmos/cosmos-sdk/issues/22246
	//
	// err = e.evmKeeper.TrackBlockHash(sdkCtx, uint64(blockHeight), blockHash)
	// if err != nil {
	// 	e.logger.Error("failed to track block hash", "err", err)
	// 	return err
	// }

	return nil
}
