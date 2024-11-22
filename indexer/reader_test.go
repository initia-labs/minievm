package indexer_test

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"

	cmttypes "github.com/cometbft/cometbft/types"

	storetypes "cosmossdk.io/store/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	coretypes "github.com/ethereum/go-ethereum/core/types"

	rpctypes "github.com/initia-labs/minievm/jsonrpc/types"
	evmtypes "github.com/initia-labs/minievm/x/evm/types"
)

func Test_Reader(t *testing.T) {
	app, indexer, addrs, privKeys := setupIndexer(t)
	defer app.Close()

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
	tx2, evmTxHash2 := generateTransferERC20Tx(t, app, privKeys[0], common.BytesToAddress(contractAddr), addrs[1], new(big.Int).SetUint64(1_000_000), 2)
	finalizeReq, finalizeRes = executeTxs(t, app, tx, tx2)
	checkTxResult(t, finalizeRes.TxResults[0], true)
	checkTxResult(t, finalizeRes.TxResults[1], true)

	txBytes, err := app.TxEncode(tx)
	require.NoError(t, err)
	txBytes2, err := app.TxEncode(tx2)
	require.NoError(t, err)

	cmtTx := cmttypes.Tx(txBytes)
	cosmosTxHash := cmtTx.Hash()
	cmtTx2 := cmttypes.Tx(txBytes2)
	cosmosTxHash2 := cmtTx2.Hash()

	// listen finalize block
	ctx, err = app.CreateQueryContext(0, false)
	require.NoError(t, err)

	err = indexer.ListenFinalizeBlock(ctx.WithBlockGasMeter(storetypes.NewInfiniteGasMeter()), *finalizeReq, *finalizeRes)
	require.NoError(t, err)

	// check the tx is indexed
	evmTx, err = indexer.TxByHash(ctx, evmTxHash)
	require.NoError(t, err)
	require.NotNil(t, evmTx)
	evmTx, err = indexer.TxByHash(ctx, evmTxHash2)
	require.NoError(t, err)
	require.NotNil(t, evmTx)

	// check the block header is indexed
	header, err := indexer.BlockHeaderByNumber(ctx, uint64(finalizeReq.Height))
	require.NoError(t, err)
	require.NotNil(t, header)
	require.Equal(t, finalizeReq.Height, header.Number.Int64())

	// check tx hash by block and index
	txHash, err := indexer.TxHashByBlockAndIndex(ctx, uint64(finalizeReq.Height), 1)
	require.NoError(t, err)
	require.Equal(t, evmTxHash, txHash)

	txHash, err = indexer.TxHashByBlockAndIndex(ctx, uint64(finalizeReq.Height), 2)
	require.NoError(t, err)
	require.Equal(t, evmTxHash2, txHash)

	// iterate block txs
	count := 0
	err = indexer.IterateBlockTxs(ctx, uint64(finalizeReq.Height), func(tx *rpctypes.RPCTransaction) (bool, error) {
		count++
		if count == 1 {
			require.Equal(t, evmTxHash, tx.Hash)
		} else if count == 2 {
			require.Equal(t, evmTxHash2, tx.Hash)
		}
		return false, nil
	})
	require.NoError(t, err)
	require.Equal(t, 2, count)

	// receipt by hash
	receipt1, err := indexer.TxReceiptByHash(ctx, evmTxHash)
	require.NoError(t, err)
	require.NotNil(t, receipt1)

	// receipt by hash
	receipt2, err := indexer.TxReceiptByHash(ctx, evmTxHash2)
	require.NoError(t, err)
	require.NotNil(t, receipt2)

	// iterate block tx receipts
	count = 0
	err = indexer.IterateBlockTxReceipts(ctx, uint64(finalizeReq.Height), func(receipt *coretypes.Receipt) (bool, error) {
		count++
		if count == 1 {
			require.Equal(t, receipt1, receipt)
		} else if count == 2 {
			require.Equal(t, receipt2, receipt)
		}
		return false, nil
	})
	require.NoError(t, err)

	// block hash to number
	blockNumber, err := indexer.BlockHashToNumber(ctx, header.Hash())
	require.NoError(t, err)
	require.Equal(t, uint64(finalizeReq.Height), blockNumber)

	// cosmos tx hash
	hash, err := indexer.CosmosTxHashByTxHash(ctx, evmTxHash)
	require.NoError(t, err)
	require.Equal(t, cosmosTxHash, hash)

	hash, err = indexer.CosmosTxHashByTxHash(ctx, evmTxHash2)
	require.NoError(t, err)
	require.Equal(t, cosmosTxHash2, hash)

	// tx hash by cosmos tx hash
	txHash, err = indexer.TxHashByCosmosTxHash(ctx, cosmosTxHash)
	require.NoError(t, err)
	require.Equal(t, evmTxHash, txHash)

	txHash, err = indexer.TxHashByCosmosTxHash(ctx, cosmosTxHash2)
	require.NoError(t, err)
	require.Equal(t, evmTxHash2, txHash)
}
