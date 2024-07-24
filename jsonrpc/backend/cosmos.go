package backend

import (
	"errors"

	"cosmossdk.io/collections"
	"github.com/ethereum/go-ethereum/common"
)

// CosmosTxHashByTxHash returns the Cosmos transaction hash by the Ethereum transaction hash.
func (b *JSONRPCBackend) CosmosTxHashByTxHash(hash common.Hash) ([]byte, error) {
	queryCtx, err := b.getQueryCtx()
	if err != nil {
		return nil, err
	}

	cosmosTxHash, err := b.app.EVMIndexer().CosmosTxHashByTxHash(queryCtx, hash)
	if err != nil && errors.Is(err, collections.ErrNotFound) {
		return nil, nil
	}

	return cosmosTxHash, nil
}

// TxHashByCosmosTxHash returns the Ethereum transaction hash by the Cosmos transaction hash.
func (b *JSONRPCBackend) TxHashByCosmosTxHash(hash []byte) (common.Hash, error) {
	queryCtx, err := b.getQueryCtx()
	if err != nil {
		return common.Hash{}, err
	}

	txHash, err := b.app.EVMIndexer().TxHashByCosmosTxHash(queryCtx, hash)
	if err != nil && errors.Is(err, collections.ErrNotFound) {
		return common.Hash{}, nil
	}

	return txHash, nil
}
