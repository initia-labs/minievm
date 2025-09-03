package indexer

import (
	"fmt"

	storetypes "cosmossdk.io/store/types"

	abci "github.com/cometbft/cometbft/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Backfill backfills the EVM indexer.
func (e *EVMIndexerImpl) Backfill(startHeight uint64, endHeight uint64) error {
	e.logger.Info("backfilling", "startHeight", startHeight, "endHeight", endHeight)
	for startHeight <= endHeight {
		if startHeight%100 == 0 {
			e.logger.Info("backfilling", "height", startHeight)
		}

		ctx, err := e.contextCreator(int64(startHeight), false)
		if err != nil {
			return fmt.Errorf("failed to create context: %w", err)
		}

		params, err := e.evmKeeper.Params.Get(sdk.UnwrapSDKContext(ctx))
		if err != nil {
			return fmt.Errorf("failed to get params: %w", err)
		}
		feeDecimals, err := e.evmKeeper.ERC20Keeper().GetDecimals(sdk.UnwrapSDKContext(ctx), params.FeeDenom)
		if err != nil {
			return fmt.Errorf("failed to get fee decimals: %w", err)
		}
		baseFee, err := e.evmKeeper.BaseFee(ctx)
		if err != nil {
			return fmt.Errorf("failed to load fee: %w", err)
		}

		height := int64(startHeight)
		block, err := e.clientCtx.Client.Block(ctx, &height)
		if err != nil {
			return err
		}
		blockResults, err := e.clientCtx.Client.BlockResults(ctx, &height)
		if err != nil {
			return err
		}

		txs := make([][]byte, len(block.Block.Data.Txs))
		for i, tx := range block.Block.Data.Txs {
			txs[i] = []byte(tx)
		}

		gasUsed := int64(0)
		for _, txResult := range blockResults.TxsResults {
			gasUsed += txResult.GasUsed
		}

		req := abci.RequestFinalizeBlock{
			Height: height,
			Txs:    txs,
		}
		res := abci.ResponseFinalizeBlock{
			TxResults: blockResults.TxsResults,
		}

		blockGasMeter := storetypes.NewInfiniteGasMeter()
		blockGasMeter.ConsumeGas(uint64(gasUsed), "block gas")

		task := &indexingTask{
			req: &req,
			res: &res,

			// state dependent args for indexing
			args: &indexingArgs{
				chainID: sdk.UnwrapSDKContext(ctx).ChainID(),

				ac:        e.ac,
				txDecoder: e.txConfig.TxDecoder(),

				params:      params,
				baseFee:     baseFee,
				feeDecimals: feeDecimals,

				blockHeight:   height,
				blockTime:     block.Block.Header.Time,
				blockGasMeter: blockGasMeter,
			},
		}

		// do backfill
		_, err = e.doIndexing(task.args, task.req, task.res)
		if err != nil {
			return err
		}

		startHeight++
	}
	return nil
}
