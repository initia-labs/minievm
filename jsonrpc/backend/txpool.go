package backend

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	rpctypes "github.com/initia-labs/minievm/jsonrpc/types"
)

func (b *JSONRPCBackend) TxPoolContent() (map[string]map[string]map[string]*rpctypes.RPCTransaction, error) {
	cache := b.app.EVMIndexer().MempoolCache()
	content := map[string]map[string]map[string]*rpctypes.RPCTransaction{
		"pending": cache.PendingContent(),
		"queued":  cache.QueuedContent(),
	}

	return content, nil
}

func (b *JSONRPCBackend) TxPoolContentFrom(addr common.Address) (map[string]map[string]*rpctypes.RPCTransaction, error) {
	content, err := b.TxPoolContent()
	if err != nil {
		return nil, err
	}

	accountContent := make(map[string]map[string]*rpctypes.RPCTransaction, 2)
	accountContent["pending"] = content["pending"][addr.Hex()]
	accountContent["queued"] = content["queued"][addr.Hex()]

	return accountContent, nil
}

// TxPoolStatus returns the number of pending and queued transactions in the pool.
func (b *JSONRPCBackend) TxPoolStatus() (map[string]hexutil.Uint, error) {
	cache := b.app.EVMIndexer().MempoolCache()
	return map[string]hexutil.Uint{
		"pending": hexutil.Uint(cache.PendingCount()),
		"queued":  hexutil.Uint(cache.QueuedCount()),
	}, nil
}

// TxPoolInspect retrieves the content of the transaction pool and flattens it into an
// easily inspectable list.
func (b *JSONRPCBackend) TxPoolInspect() (map[string]map[string]map[string]string, error) {
	inspectContent := map[string]map[string]map[string]string{
		"pending": make(map[string]map[string]string),
		"queued":  make(map[string]map[string]string),
	}

	content, err := b.TxPoolContent()
	if err != nil {
		return nil, err
	}

	// Define a formatter to flatten a transaction into a string
	var format = func(tx *rpctypes.RPCTransaction) string {
		if to := tx.To; to != nil {
			return fmt.Sprintf("%s: %v wei + %v gas × %v wei", tx.To.Hex(), tx.Value, tx.Gas, tx.GasPrice)
		}
		return fmt.Sprintf("contract creation: %v wei + %v gas × %v wei", tx.Value, tx.Gas, tx.GasPrice)
	}
	// Flatten the pending transactions
	for account, txs := range content["pending"] {
		dump := make(map[string]string)
		for _, tx := range txs {
			dump[fmt.Sprintf("%d", tx.Nonce)] = format(tx)
		}
		inspectContent["pending"][account] = dump
	}
	for account, txs := range content["queued"] {
		dump := make(map[string]string)
		for _, tx := range txs {
			dump[fmt.Sprintf("%d", tx.Nonce)] = format(tx)
		}
		inspectContent["queued"][account] = dump
	}
	return inspectContent, nil
}
