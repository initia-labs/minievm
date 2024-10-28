package indexer

import (
	"context"

	"github.com/jellydator/ttlcache/v3"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/mempool"
	"github.com/ethereum/go-ethereum/common"

	rpctypes "github.com/initia-labs/minievm/jsonrpc/types"
	evmkeeper "github.com/initia-labs/minievm/x/evm/keeper"
	evmtypes "github.com/initia-labs/minievm/x/evm/types"
)

var _ mempool.Mempool = (*MempoolWrapper)(nil)

type MempoolWrapper struct {
	mempool mempool.Mempool
	indexer *EVMIndexerImpl
}

// MempoolWrapper returns a mempool wrapper that emits transactions to the filters.
func (indexer *EVMIndexerImpl) MempoolWrapper(mempool mempool.Mempool) mempool.Mempool {
	return &MempoolWrapper{mempool: mempool, indexer: indexer}
}

// TxInMempool returns true if the transaction with the given hash is in the mempool.
func (indexer *EVMIndexerImpl) TxInMempool(hash common.Hash) bool {
	return indexer.txPendingMap.Has(hash)
}

// CountTx implements mempool.Mempool.
func (m *MempoolWrapper) CountTx() int {
	return m.mempool.CountTx()
}

// Insert implements mempool.Mempool.
func (m *MempoolWrapper) Insert(ctx context.Context, tx sdk.Tx) error {
	txUtils := evmkeeper.NewTxUtils(m.indexer.evmKeeper)
	ethTx, _, err := txUtils.ConvertCosmosTxToEthereumTx(ctx, tx)
	if err != nil {
		m.indexer.logger.Error("failed to convert CosmosTx to EthTx", "err", err)
		return err
	}

	if ethTx != nil {
		sdkCtx := sdk.UnwrapSDKContext(ctx)
		chainId := evmtypes.ConvertCosmosChainIDToEthereumChainID(sdkCtx.ChainID())
		rpcTx := rpctypes.NewRPCTransaction(ethTx, common.Hash{}, 0, 0, chainId)

		ethTxHash := ethTx.Hash()

		m.indexer.logger.Debug("inserting tx into mempool", "pending len", m.indexer.txPendingMap.Len(), "ethTxHash", ethTxHash)
		m.indexer.txPendingMap.Set(ethTxHash, true, ttlcache.DefaultTTL)

		go func() {
			// emit the transaction to all pending channels
			for _, pendingChan := range m.indexer.pendingChans {
				pendingChan <- rpcTx
			}
		}()
	}

	return m.mempool.Insert(ctx, tx)
}

// Remove implements mempool.Mempool.
func (m *MempoolWrapper) Remove(tx sdk.Tx) error {
	return m.mempool.Remove(tx)
}

// Select implements mempool.Mempool.
func (m *MempoolWrapper) Select(ctx context.Context, txs [][]byte) mempool.Iterator {
	return m.mempool.Select(ctx, txs)
}
