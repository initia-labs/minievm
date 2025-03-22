package checktx

import (
	"sync"
	"time"

	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/jellydator/ttlcache/v3"

	"cosmossdk.io/log"
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	cosmosante "github.com/cosmos/cosmos-sdk/x/auth/ante"

	"github.com/initia-labs/minievm/indexer"
	evmtypes "github.com/initia-labs/minievm/x/evm/types"

	"github.com/ethereum/go-ethereum/common"
	coretypes "github.com/ethereum/go-ethereum/core/types"

	blockchecktx "github.com/skip-mev/block-sdk/v2/abci/checktx"
)

type CheckTxWrapper struct {
	logger   log.Logger
	txConfig client.TxConfig
	cg       ContextGetter

	ng  NonceGetter
	bg  BalanceGetter
	fdg FeeDenomGetter
	ei  indexer.EVMIndexer

	txUtils evmtypes.TxUtils

	checkTx      blockchecktx.CheckTx
	txFeeChecker cosmosante.TxFeeChecker

	feeDenom       string
	feeDenomHeight uint64

	txQueue *ttlcache.Cache[txKey, txItem]

	responses       *sync.Map
	responsesHeight uint64

	stop chan struct{}
}

type txKey struct {
	sender common.Address
	nonce  uint64
}

type txItem struct {
	txBytes []byte
	tx      sdk.Tx
	ethTx   *coretypes.Transaction
}

func NewCheckTxWrapper(
	logger log.Logger,
	txConfig client.TxConfig,
	cg ContextGetter,
	ng NonceGetter,
	bg BalanceGetter,
	fdg FeeDenomGetter,
	ei indexer.EVMIndexer,
	txUtils evmtypes.TxUtils,
	checkTx blockchecktx.CheckTx,
	txFeeChecker cosmosante.TxFeeChecker,
) *CheckTxWrapper {
	w := &CheckTxWrapper{
		logger:   logger,
		txConfig: txConfig,

		cg:  cg,
		ng:  ng,
		bg:  bg,
		fdg: fdg,
		ei:  ei,

		txUtils: txUtils,

		checkTx:      checkTx,
		txFeeChecker: txFeeChecker,

		txQueue: ttlcache.New(ttlcache.WithTTL[txKey, txItem](time.Minute)),

		responses:       new(sync.Map),
		responsesHeight: 0,

		stop: make(chan struct{}),
	}

	// start the tx queue to evict expired txs
	go w.txQueue.Start()

	return w
}

func (w *CheckTxWrapper) Stop() {
	w.txQueue.Stop()
	close(w.stop)
}

// WrapCheckTx wrap the default checkTx handler to check the transaction is evm tx.
//
// - If the transaction is not evm tx, it will be passed to the default handler.
// - If the transaction is evm tx and sequence is equal to account sequence, it will be passed to the default handler.
// - If the transaction is evm tx and sequence is greater than account sequence, it will be minimally verified and kept in the cometbft mempool for recheck.
//
// After the above steps are finished, try to flush
func (w *CheckTxWrapper) CheckTx() blockchecktx.CheckTx {
	return func(req *abci.RequestCheckTx) (*abci.ResponseCheckTx, error) {
		isRecheck := req.Type == abci.CheckTxType_Recheck
		blockHeight, sdkTx, ethTx, sender, accNonce, err := w.getTxInfo(req)
		if err != nil {
			return sdkerrors.ResponseCheckTxWithEvents(err, 0, 0, nil, false), nil
		}

		// refresh responses map
		if w.responsesHeight != blockHeight {
			w.responsesHeight = blockHeight
			w.responses.Clear()
		}

		// check responses first
		if res, ok := w.responses.Load(ethTx.Hash()); ok {
			return res.(*abci.ResponseCheckTx), nil
		}

		isTxInQueue := false
		isEthTx := ethTx != nil && sender != nil
		if isEthTx {
			isTxInQueue = w.isTxInQueue(*sender, ethTx.Nonce())
		}

		// if not recheck, then pass to checkTx handler
		if !isRecheck || !isTxInQueue {
			res, err := w.checkTxHandler(req, sdkTx, ethTx, sender, accNonce)
			if err != nil {
				return sdkerrors.ResponseCheckTxWithEvents(err, 0, 0, nil, false), nil
			} else if res.Code != abci.CodeTypeOK || !isEthTx {
				return res, nil
			}

			// If the tx was passed to the default checkTx handler,
			// we need to increase the account sequence
			if res.Codespace != "txqueue" {
				accNonce++
			}

			// run flush queue
			w.flushQueue(sender, accNonce)

			return res, nil
		}

		// run flush queue
		w.flushQueue(sender, accNonce)

		// check responses
		if res, ok := w.responses.Load(ethTx.Hash()); ok {
			return res.(*abci.ResponseCheckTx), nil
		}

		// response okay to keep the tx in the mempool for recheck triggered by cometbft
		return &abci.ResponseCheckTx{
			Code:      abci.CodeTypeOK,
			Codespace: "txqueue",
		}, nil
	}
}

func (w *CheckTxWrapper) isTxInQueue(sender common.Address, nonce uint64) bool {
	return w.txQueue.Has(txKey{sender: sender, nonce: nonce})
}

func (w *CheckTxWrapper) appendToQueue(txBytes []byte, tx sdk.Tx, ethTx *coretypes.Transaction, sender common.Address) {
	w.ei.PushQueuedTx(ethTx)
	w.txQueue.Set(txKey{sender: sender, nonce: ethTx.Nonce()}, txItem{txBytes: txBytes, tx: tx, ethTx: ethTx}, ttlcache.DefaultTTL)
}

func (w *CheckTxWrapper) removeFromQueue(ethTxHash common.Hash, sender common.Address, nonce uint64) {
	w.ei.RemoveQueuedTx(ethTxHash)
	w.txQueue.Delete(txKey{sender: sender, nonce: nonce})
}
