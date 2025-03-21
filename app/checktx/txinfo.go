package checktx

import (
	abci "github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	coretypes "github.com/ethereum/go-ethereum/core/types"
)

// getTxInfo returns the block height, sdk tx, eth tx, expected sender, and account sequence.
func (w *CheckTxWrapper) getTxInfo(req *abci.RequestCheckTx) (uint64, sdk.Tx, *coretypes.Transaction, *common.Address, uint64, error) {
	sdkCtx := w.cg.GetContextForCheckTx(req.Tx)
	sdkTx, err := w.txConfig.TxDecoder()(req.Tx)
	if err != nil {
		return 0, nil, nil, nil, 0, err
	}

	// check sequence and signature if tx is evm tx
	ethTx, expectedSender, err := w.txUtils.ConvertCosmosTxToEthereumTx(sdkCtx, sdkTx)
	if err != nil {
		return 0, nil, nil, nil, 0, err
	} else if ethTx == nil || expectedSender == nil {
		// normal cosmos tx, pass to the default checkTx handler
		return 0, nil, nil, nil, 0, nil
	}

	// check sequence is greater than account sequence
	accSequence, err := w.ng.GetSequence(sdkCtx, sdk.AccAddress(expectedSender.Bytes()))
	if err != nil {
		return 0, nil, nil, nil, 0, err
	}

	blockHeight := uint64(sdkCtx.BlockHeight())
	return blockHeight, sdkTx, ethTx, expectedSender, accSequence, nil
}
