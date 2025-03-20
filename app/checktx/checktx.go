package checktx

import (
	"time"

	abci "github.com/cometbft/cometbft/abci/types"

	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/log"
	sdkmath "cosmossdk.io/math"
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

const (
	txTimeoutCode     = 999
	txTimeoutDuration = time.Minute * 10
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

	txCreationTimes map[common.Hash]time.Time
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
	return &CheckTxWrapper{
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

		txCreationTimes: make(map[common.Hash]time.Time),
	}
}

func (w *CheckTxWrapper) updateFeeDenomCache(sdkCtx sdk.Context) error {
	if w.feeDenomHeight != uint64(sdkCtx.BlockHeight()) {
		var err error
		w.feeDenom, err = w.fdg.GetFeeDenom(sdkCtx)
		if err != nil {
			return err
		}

		w.feeDenomHeight = uint64(sdkCtx.BlockHeight())
	}

	return nil
}

// sendToCheckTx sends the transaction to the original checkTx handler.
// - if it is recheck and tx in the queue, then it is fresh tx and we need to set the check type to new
// - if checkTx finished successfully, then move the tx to pending txs
// - if checkTx failed, then remove the tx from the queue
func (w *CheckTxWrapper) sendToCheckTx(req *abci.RequestCheckTx, ethTx *coretypes.Transaction, isRecheck bool) (*abci.ResponseCheckTx, error) {
	if ethTx != nil && isRecheck {
		queuedTx := w.ei.TxInQueued(ethTx.Hash())
		if queuedTx != nil {
			// queued txs are not checked by cosmos ante, so we need to set the check type to new
			req.Type = abci.CheckTxType_New
			defer func() {
				req.Type = abci.CheckTxType_Recheck
			}()
		}
	}

	res, err := w.checkTx(req)
	if ethTx != nil {
		if err == nil && res.Code == abci.CodeTypeOK {
			// push to pending txs
			w.ei.PushPendingTx(ethTx)

			// recheck failed; remove queued txs
			if isRecheck {
				w.ei.RemoveQueuedTx(ethTx.Hash())
			}
		} else {
			if err != nil {
				w.logger.Warn("failed to push pending tx", "ethTxHash", ethTx.Hash(), "err", err)
			} else if res.Code != abci.CodeTypeOK {
				w.logger.Warn("failed to push pending tx", "ethTxHash", ethTx.Hash(), "code", res.Code, "log", res.Log)
			}

			// recheck failed; remove queued txs
			if isRecheck {
				w.ei.RemoveQueuedTx(ethTx.Hash())
			}
		}
	}

	return res, err
}

func (w *CheckTxWrapper) customCheckTx(req *abci.RequestCheckTx) (*abci.ResponseCheckTx, error) {
	isRecheck := req.Type == abci.CheckTxType_Recheck
	sdkCtx := w.cg.GetContextForCheckTx(req.Tx)
	tx, err := w.txConfig.TxDecoder()(req.Tx)
	if err != nil {
		return nil, err
	}

	// check sequence and signature if tx is evm tx
	ethTx, expectedSender, err := w.txUtils.ConvertCosmosTxToEthereumTx(sdkCtx, tx)
	if err != nil {
		return nil, err
	} else if ethTx == nil {
		// normal cosmos tx, pass to the default checkTx handler
		return w.sendToCheckTx(req, ethTx, isRecheck)
	}

	// check sequence is greater than account sequence
	accSequence, err := w.ng.GetSequence(sdkCtx, sdk.AccAddress(expectedSender.Bytes()))
	if err != nil {
		return nil, err
	}

	// sequence must be greater than account sequence
	if ethTx.Nonce() < accSequence {
		return nil, errorsmod.Wrapf(
			sdkerrors.ErrWrongSequence,
			"account sequence mismatch, expected %d, got %d", accSequence, ethTx.Nonce(),
		)
	} else if ethTx.Nonce() == accSequence {
		// if sequence is equal to current sequence, it means the transaction is ready to be processed
		return w.sendToCheckTx(req, ethTx, req.Type == abci.CheckTxType_Recheck)
	}

	// load fee denom if it is not cached
	err = w.updateFeeDenomCache(sdkCtx)
	if err != nil {
		return nil, err
	}

	// check balance
	balance := w.bg.GetBalance(sdkCtx, sdk.AccAddress(expectedSender.Bytes()), w.feeDenom)
	if balance.Amount.LT(sdkmath.NewIntFromBigInt(ethTx.Cost())) {
		if isRecheck {
			// recheck failed; remove queued txs
			w.ei.RemoveQueuedTx(ethTx.Hash())
		}

		return nil, errorsmod.Wrapf(sdkerrors.ErrInsufficientFunds, "insufficient balance for tx: %s", ethTx.Hash().Hex())
	}

	// check timeout if tx is in recheck
	if isRecheck {
		if creationTime, ok := w.txCreationTimes[ethTx.Hash()]; ok {
			if time.Now().After(creationTime.Add(txTimeoutDuration)) {
				delete(w.txCreationTimes, ethTx.Hash())

				// recheck failed; remove queued txs
				w.ei.RemoveQueuedTx(ethTx.Hash())

				return &abci.ResponseCheckTx{
					Code: txTimeoutCode,
					Log:  "tx timeout",
				}, nil
			}
		}

		return &abci.ResponseCheckTx{
			Code:      abci.CodeTypeOK,
			Codespace: "txqueue",
		}, nil
	} else {
		w.txCreationTimes[ethTx.Hash()] = time.Now()
	}

	// if sequence is greater than account sequence, it means the transaction is not ready to be processed
	// so we just do minimal check and return nil to keep the tx in the cometbft mempool and retry at recheck

	// verify signature
	ethChainID := evmtypes.ConvertCosmosChainIDToEthereumChainID(sdkCtx.ChainID())
	signer := coretypes.LatestSignerForChainID(ethChainID)
	sender, err := signer.Sender(ethTx)
	if err != nil {
		return nil, errorsmod.Wrapf(sdkerrors.ErrorInvalidSigner, "failed to recover sender address: %v", err)
	}

	// check if the recovered sender matches the expected sender
	if expectedSender == nil || *expectedSender != sender {
		return nil, errorsmod.Wrapf(sdkerrors.ErrorInvalidSigner, "expected sender %s, got %s", expectedSender, sender)
	}

	// min gas prices check
	_, _, err = w.txFeeChecker(sdkCtx, tx)
	if err != nil {
		return nil, err
	}

	// push to queued txs
	w.ei.PushQueuedTx(ethTx)

	return &abci.ResponseCheckTx{
		Code:      abci.CodeTypeOK,
		Codespace: "txqueue",
	}, nil

}

// WrapCheckTx wrap the default checkTx handler to check the transaction is evm tx.
//
// - If the transaction is not evm tx, it will be passed to the default handler.
// - If the transaction is evm tx and sequence is equal to account sequence, it will be passed to the default handler.
// - If the transaction is evm tx and sequence is greater than account sequence, it will be minimally verified and kept in the cometbft mempool for recheck.
func (w *CheckTxWrapper) CheckTx() blockchecktx.CheckTx {
	return func(req *abci.RequestCheckTx) (*abci.ResponseCheckTx, error) {
		res, err := w.customCheckTx(req)
		if err != nil {
			return sdkerrors.ResponseCheckTxWithEvents(err, 0, 0, nil, false), nil
		}

		return res, nil
	}
}
