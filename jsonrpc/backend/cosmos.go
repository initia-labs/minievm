package backend

import (
	"errors"

	"cosmossdk.io/collections"
	"github.com/ethereum/go-ethereum/common"
)

// CosmosTxHashByTxHash returns the Cosmos transaction hash by the Ethereum transaction hash.
func (b *JSONRPCBackend) CosmosTxHashByTxHash(hash common.Hash) ([]byte, error) {
	cosmosTxHash, err := b.app.EVMIndexer().CosmosTxHashByTxHash(b.ctx, hash)
	if err != nil && errors.Is(err, collections.ErrNotFound) {
		return nil, nil
	} else if err != nil {
		b.logger.Error("failed to get cosmos tx hash by tx hash", "err", err)
		return nil, NewInternalError("failed to get cosmos tx hash by tx hash")
	}

	return cosmosTxHash, nil
}

// TxHashByCosmosTxHash returns the Ethereum transaction hash by the Cosmos transaction hash.
func (b *JSONRPCBackend) TxHashByCosmosTxHash(hash []byte) (common.Hash, error) {
	txHash, err := b.app.EVMIndexer().TxHashByCosmosTxHash(b.ctx, hash)
	if err != nil && errors.Is(err, collections.ErrNotFound) {
		return common.Hash{}, nil
	} else if err != nil {
		b.logger.Debug("failed to get tx hash by cosmos tx hash", "err", err)
		return common.Hash{}, NewInternalError("failed to get tx hash by cosmos tx hash")
	}

	return txHash, nil
}
