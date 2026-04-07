package app

import (
	cmtmempool "github.com/cometbft/cometbft/mempool"
	"github.com/ethereum/go-ethereum/common"

	"github.com/initia-labs/initia/abcipp"

	rpctypes "github.com/initia-labs/minievm/jsonrpc/types"
)

// ConnectMempoolEvents wires cometbft ProxyMempool event channel to the app mempool.
func (app *MinitiaApp) ConnectMempoolEvents(eventCh chan cmtmempool.AppMempoolEvent) {
	pm, ok := app.Mempool().(*abcipp.PriorityMempool)
	if !ok {
		return
	}

	pm.SetEventCh(eventCh)

	appCh := make(chan cmtmempool.AppMempoolEvent, 8192)
	pm.SetAppEventCh(appCh)

	cache := app.evmIndexer.MempoolCache()
	cache.StartEventConsumer(appCh, app.convertTxBytesToRPCTx)
}

// convertTxBytesToRPCTx decodes raw cosmos tx bytes and converts to an eth RPCTransaction.
// Returns nil for non-EVM transactions or on error.
func (app *MinitiaApp) convertTxBytesToRPCTx(txBytes []byte) *rpctypes.RPCTransaction {
	sdkTx, err := app.txConfig.TxDecoder()(txBytes)
	if err != nil {
		return nil
	}

	ctx := app.GetContextForCheckTx(nil)
	ethTx, _, err := app.EVMKeeper.TxUtils().ConvertCosmosTxToEthereumTx(ctx, sdkTx)
	if err != nil || ethTx == nil {
		return nil
	}

	return rpctypes.NewRPCTransaction(ethTx, common.Hash{}, 0, 0, ethTx.ChainId())
}
