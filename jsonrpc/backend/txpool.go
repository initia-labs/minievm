package backend

import (
	"fmt"

	rpcclient "github.com/cometbft/cometbft/rpc/client"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	rpctypes "github.com/initia-labs/minievm/jsonrpc/types"
	"github.com/initia-labs/minievm/x/evm/keeper"
)

func (b *JSONRPCBackend) TxPoolContent() (map[string]map[string]map[string]*rpctypes.RPCTransaction, error) {
	content := map[string]map[string]map[string]*rpctypes.RPCTransaction{
		"pending": make(map[string]map[string]*rpctypes.RPCTransaction),
		"queued":  make(map[string]map[string]*rpctypes.RPCTransaction),
	}

	limit := int(100)
	pending, err := b.clientCtx.Client.(rpcclient.MempoolClient).UnconfirmedTxs(b.ctx, &limit)
	if err != nil {
		return nil, err
	}

	ctx, err := b.getQueryCtx()
	if err != nil {
		return nil, err
	}

	txUtils := keeper.NewTxUtils(b.app.EVMKeeper)
	for _, tx := range pending.Txs {
		cosmosTx, err := b.app.TxDecode(tx)
		if err != nil {
			return nil, err
		}

		ethTx, account, err := txUtils.ConvertCosmosTxToEthereumTx(ctx, cosmosTx)
		if err != nil {
			return nil, err
		}
		if ethTx == nil {
			continue
		}

		dump, ok := content["pending"][account.Hex()]
		if !ok {
			dump = make(map[string]*rpctypes.RPCTransaction)
			content["pending"][account.Hex()] = dump
		}

		dump[fmt.Sprintf("%d", ethTx.Nonce())] = rpctypes.NewRPCTransaction(ethTx, common.Hash{}, 0, 0, ethTx.ChainId())
	}

	// load queued txs
	b.mut.Lock()
	queuedTxs := b.queuedTxs.Values()
	b.mut.Unlock()

	for _, txQueueItem := range queuedTxs {
		ethTx := txQueueItem.body
		sender := txQueueItem.sender

		dump, ok := content["queued"][sender]
		if !ok {
			dump = make(map[string]*rpctypes.RPCTransaction)
			content["queued"][sender] = dump
		}

		dump[fmt.Sprintf("%d", ethTx.Nonce())] = rpctypes.NewRPCTransaction(ethTx, common.Hash{}, 0, 0, ethTx.ChainId())
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

// Status returns the number of pending and queued transaction in the pool.
func (b *JSONRPCBackend) TxPoolStatus() (map[string]hexutil.Uint, error) {
	numPendingTxs, err := b.clientCtx.Client.(rpcclient.MempoolClient).NumUnconfirmedTxs(b.ctx)
	if err != nil {
		return nil, err
	}

	b.mut.Lock()
	numQueuedTxs := b.queuedTxs.Len()
	b.mut.Unlock()

	return map[string]hexutil.Uint{
		"pending": hexutil.Uint(numPendingTxs.Count),
		"queued":  hexutil.Uint(numQueuedTxs),
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
