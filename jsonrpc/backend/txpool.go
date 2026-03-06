package backend

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/initia-labs/initia/abcipp"

	sdk "github.com/cosmos/cosmos-sdk/types"

	rpctypes "github.com/initia-labs/minievm/jsonrpc/types"
	"github.com/initia-labs/minievm/x/evm/keeper"
)

func (b *JSONRPCBackend) TxPoolContent() (map[string]map[string]map[string]*rpctypes.RPCTransaction, error) {
	content := map[string]map[string]map[string]*rpctypes.RPCTransaction{
		"pending": make(map[string]map[string]*rpctypes.RPCTransaction),
		"queued":  make(map[string]map[string]*rpctypes.RPCTransaction),
	}

	queryCtx, closer, err := b.getQueryCtx()
	if closer != nil {
		defer closer.Close()
	}
	if err != nil {
		return nil, err
	}
	sdkCtx := sdk.UnwrapSDKContext(queryCtx)
	txUtils := keeper.NewTxUtils(b.app.EVMKeeper)

	mempool, ok := b.app.Mempool().(abcipp.Mempool)
	if !ok {
		return content, nil
	}

	mempool.IteratePendingTxs(func(_ string, _ uint64, tx sdk.Tx) bool {
		ethTx, _, err := txUtils.ConvertCosmosTxToEthereumTx(sdkCtx, tx)
		if err != nil || ethTx == nil {
			return true
		}
		rpcTx := rpctypes.NewRPCTransaction(ethTx, common.Hash{}, 0, 0, ethTx.ChainId())
		dump, ok := content["pending"][rpcTx.From.Hex()]
		if !ok {
			dump = make(map[string]*rpctypes.RPCTransaction)
			content["pending"][rpcTx.From.Hex()] = dump
		}
		dump[fmt.Sprintf("%d", rpcTx.Nonce)] = rpcTx
		return true
	})

	mempool.IterateQueuedTxs(func(_ string, _ uint64, tx sdk.Tx) bool {
		ethTx, _, err := txUtils.ConvertCosmosTxToEthereumTx(sdkCtx, tx)
		if err != nil || ethTx == nil {
			return true
		}
		rpcTx := rpctypes.NewRPCTransaction(ethTx, common.Hash{}, 0, 0, ethTx.ChainId())
		dump, ok := content["queued"][rpcTx.From.Hex()]
		if !ok {
			dump = make(map[string]*rpctypes.RPCTransaction)
			content["queued"][rpcTx.From.Hex()] = dump
		}
		dump[fmt.Sprintf("%d", rpcTx.Nonce)] = rpcTx
		return true
	})

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

// Status returns the number of pending and queued transaction in the pool.
func (b *JSONRPCBackend) TxPoolStatus() (map[string]hexutil.Uint, error) {
	numPending := 0
	numQueued := 0

	mempool, ok := b.app.Mempool().(abcipp.Mempool)
	if ok {
		mempool.IteratePendingTxs(func(_ string, _ uint64, _ sdk.Tx) bool {
			numPending++
			return true
		})
		mempool.IterateQueuedTxs(func(_ string, _ uint64, _ sdk.Tx) bool {
			numQueued++
			return true
		})
	}

	return map[string]hexutil.Uint{
		"pending": hexutil.Uint(numPending),
		"queued":  hexutil.Uint(numQueued),
	}, nil
}

// Inspect retrieves the content of the transaction pool and flattens it into an
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
