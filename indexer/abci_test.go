package indexer_test

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"

	storetypes "cosmossdk.io/store/types"

	evmtypes "github.com/initia-labs/minievm/x/evm/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

func Test_ListenFinalizeBlock(t *testing.T) {
	app, indexer, addrs, privKeys := setupIndexer(t)

	tx, evmTxHash := generateCreateERC20Tx(t, app, privKeys[0])
	finalizeReq, finalizeRes := executeTxs(t, app, tx)
	checkTxResult(t, finalizeRes.TxResults[0], true)

	events := finalizeRes.TxResults[0].Events
	createEvent := events[len(events)-3]
	require.Equal(t, evmtypes.EventTypeContractCreated, createEvent.GetType())

	contractAddr, err := hexutil.Decode(createEvent.Attributes[0].Value)
	require.NoError(t, err)

	// listen finalize block
	ctx, err := app.CreateQueryContext(0, false)
	require.NoError(t, err)

	err = indexer.ListenFinalizeBlock(ctx.WithBlockGasMeter(storetypes.NewInfiniteGasMeter()), *finalizeReq, *finalizeRes)
	require.NoError(t, err)

	// check the tx is indexed
	evmTx, err := indexer.TxByHash(ctx, evmTxHash)
	require.NoError(t, err)
	require.NotNil(t, evmTx)

	// mint 1_000_000 tokens to the first address
	tx, evmTxHash = generateMintERC20Tx(t, app, privKeys[0], common.BytesToAddress(contractAddr), addrs[0], new(big.Int).SetUint64(1_000_000_000_000))
	finalizeReq, finalizeRes = executeTxs(t, app, tx)
	checkTxResult(t, finalizeRes.TxResults[0], true)

	// listen finalize block
	ctx, err = app.CreateQueryContext(0, false)
	require.NoError(t, err)

	err = indexer.ListenFinalizeBlock(ctx.WithBlockGasMeter(storetypes.NewInfiniteGasMeter()), *finalizeReq, *finalizeRes)
	require.NoError(t, err)

	// check the tx is indexed
	evmTx, err = indexer.TxByHash(ctx, evmTxHash)
	require.NoError(t, err)
	require.NotNil(t, evmTx)
}
