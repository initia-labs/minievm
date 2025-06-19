package checktx

import (
	abci "github.com/cometbft/cometbft/abci/types"

	errorsmod "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	evmtypes "github.com/initia-labs/minievm/x/evm/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	coretypes "github.com/ethereum/go-ethereum/core/types"
)

// moveToPending moves the transaction to the pending pool.
func (w *CheckTxWrapper) moveToPending(
	req *abci.RequestCheckTx,
	ethTx *coretypes.Transaction,
) (*abci.ResponseCheckTx, error) {
	res, err := w.checkTx(req)
	if err == nil && res.Code == abci.CodeTypeOK && ethTx != nil {
		w.ei.PushPendingTx(ethTx)
	}

	return res, err
}

// updateFeeDenomCache updates the fee denom cache.
//
// - If the fee denom height is not the same as the current block height, it will update the fee denom.
// - It will then update the fee denom height.
func (w *CheckTxWrapper) updateFeeDenomCache(sdkCtx sdk.Context) error {
	w.feeDenomMut.Lock()
	defer w.feeDenomMut.Unlock()

	if w.feeDenomHeight < uint64(sdkCtx.BlockHeight()) {
		var err error
		w.feeDenom, err = w.fdg.GetFeeDenom(sdkCtx)
		if err != nil {
			return err
		}

		w.feeDenomHeight = uint64(sdkCtx.BlockHeight())
	}

	return nil
}

// checkTxHandler is the handler for the checkTx request.
//
// - If the transaction is not evm tx, it will be passed to the default handler.
// - If the transaction is evm tx and sequence is equal to account sequence, it will be passed to the default handler.
// - If the transaction is evm tx and sequence is greater than account sequence, it will be minimally verified and kept in the cometbft mempool for recheck.
func (w *CheckTxWrapper) checkTxHandler(
	req *abci.RequestCheckTx,
	sdkTx sdk.Tx,
	ethTx *coretypes.Transaction,
	sender *common.Address,
	accSequence uint64,
	checkReplace bool,
) (*abci.ResponseCheckTx, error) {
	// normal cosmos tx, pass to the default checkTx handler
	if ethTx == nil || sender == nil {
		return w.moveToPending(req, ethTx)
	}

	// sequence must be greater than account sequence
	if ethTx.Nonce() < accSequence {
		return nil, errorsmod.Wrapf(
			sdkerrors.ErrWrongSequence,
			"account sequence mismatch, expected %d, got %d", accSequence, ethTx.Nonce(),
		)
	} else if ethTx.Nonce() == accSequence {
		// if sequence is equal to current sequence, it means the transaction is ready to be processed
		return w.moveToPending(req, ethTx)
	}

	// if the sequence is greater than the account sequence, check if we have a transaction
	// with the same (sender, nonce) in the queue, and try to replace it with the new one if conditions are met
	sdkCtx := w.cg.GetContextForCheckTx(req.Tx)
	res, err := w.validateTx(sdkCtx, sdkTx, ethTx, *sender)
	if err != nil {
		return nil, err
	} else if res.Code != abci.CodeTypeOK {
		return res, nil
	}

	// check if there's an existing transaction with the same sender and nonce
	if checkReplace {
		if existingTx, ok := w.getTxFromQueue(*sender, ethTx.Nonce()); ok {
			existingEthTx := existingTx.ethTx

			// compare gas prices - if the new tx has higher gas price, replace the old one
			if ethTx.GasPrice().Cmp(existingEthTx.GasPrice()) > 0 {
				// remove the old transaction
				w.removeFromQueue(existingEthTx.Hash(), *sender, ethTx.Nonce())
			} else {
				// if the new tx has lower or equal gas price, reject it
				return nil, errorsmod.Wrapf(
					sdkerrors.ErrInsufficientFee,
					"existing transaction with the same sequence has higher or equal gas price",
				)
			}
		}
	}

	// proceeds with appending tx to the queue
	w.appendToQueue(req.Tx, sdkTx, ethTx, *sender)
	return res, nil
}

// flushQueue flushes the queue of the sender.
//
// - If the sender is nil, it will return immediately.
// - It will load the nonce from the current checkTx state and remove the tx from the queue.
// - It will then pass the tx to the default checkTx handler.
func (w *CheckTxWrapper) flushQueue(sender *common.Address, nonce uint64) {
	if sender == nil {
		return
	}

	for {
		select {
		case <-w.stop:
			return
		default:
		}

		txItem := w.txQueue.Get(txKey{sender: *sender, nonce: nonce})
		if txItem == nil {
			break
		}

		// remove from the queue
		ethTx := txItem.Value().ethTx
		txHash := ethTx.Hash()
		w.removeFromQueue(txHash, *sender, nonce)

		// run default checkTx handler and save the response
		if res, err := w.moveToPending(&abci.RequestCheckTx{
			Tx:   txItem.Value().txBytes,
			Type: abci.CheckTxType_New,
		}, ethTx); err != nil {
			w.responsesMut.Lock()
			w.responses[txHash] = sdkerrors.ResponseCheckTxWithEvents(err, 0, 0, nil, false)
			w.responsesMut.Unlock()

			w.logger.Error("failed to check tx", "error", err)
			break
		} else if res.Code != abci.CodeTypeOK {
			w.responsesMut.Lock()
			w.responses[txHash] = res
			w.responsesMut.Unlock()

			w.logger.Error("failed to check tx", "code", res.Code, "log", res.Log)
			break
		} else {
			w.responsesMut.Lock()
			w.responses[txHash] = res
			w.responsesMut.Unlock()
		}

		nonce++
	}
}

func (w *CheckTxWrapper) validateTx(sdkCtx sdk.Context, tx sdk.Tx, ethTx *coretypes.Transaction, expectedSender common.Address) (*abci.ResponseCheckTx, error) {
	// check intrinsic gas
	intrGas, err := core.IntrinsicGas(ethTx.Data(), ethTx.AccessList(), []coretypes.SetCodeAuthorization{}, ethTx.To() == nil, true, true, true)
	if err != nil {
		return nil, err
	}
	if ethTx.Gas() < intrGas {
		return nil, errorsmod.Wrapf(core.ErrIntrinsicGas, "gas %v, minimum needed %v", ethTx.Gas(), intrGas)
	}

	// verify signature
	ethChainID := evmtypes.ConvertCosmosChainIDToEthereumChainID(sdkCtx.ChainID())
	signer := coretypes.LatestSignerForChainID(ethChainID)
	recoveredSender, err := signer.Sender(ethTx)
	if err != nil {
		return nil, errorsmod.Wrapf(sdkerrors.ErrorInvalidSigner, "failed to recover sender address: %v", err)
	}

	// check if the recovered sender matches the expected sender
	if expectedSender != recoveredSender {
		return nil, errorsmod.Wrapf(sdkerrors.ErrorInvalidSigner, "expected sender %s, got %s", expectedSender, recoveredSender)
	}

	// min gas prices check
	_, _, err = w.txFeeChecker(sdkCtx, tx)
	if err != nil {
		return nil, err
	}

	// load fee denom if it is not cached
	err = w.updateFeeDenomCache(sdkCtx)
	if err != nil {
		return nil, err
	}

	// check balance
	balance := w.bg.GetBalance(sdkCtx, sdk.AccAddress(expectedSender.Bytes()), w.feeDenom)
	if balance.Amount.LT(sdkmath.NewIntFromBigInt(ethTx.Cost())) {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInsufficientFunds, "insufficient balance for tx: %s", ethTx.Hash().Hex())
	}

	return &abci.ResponseCheckTx{
		Code:      abci.CodeTypeOK,
		Codespace: "txqueue",
	}, nil
}
