package indexer_test

import (
	"math/big"
	"testing"
	"time"

	"cosmossdk.io/collections"
	"github.com/stretchr/testify/require"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	evmindexer "github.com/initia-labs/minievm/indexer"
	"github.com/initia-labs/minievm/tests"
	evmconfig "github.com/initia-labs/minievm/x/evm/config"
	evmtypes "github.com/initia-labs/minievm/x/evm/types"
)

func Test_PruneIndexer(t *testing.T) {
	app, addrs, privKeys := tests.CreateApp(t)
	indexer := app.EVMIndexer().(*evmindexer.EVMIndexerImpl)
	defer app.Close()

	// set retain height to 1, only last block is indexed
	indexer.SetRetainHeight(1)

	tx, evmTxHash := tests.GenerateCreateERC20Tx(t, app, privKeys[0])
	_, finalizeRes := tests.ExecuteTxs(t, app, tx)
	tests.CheckTxResult(t, finalizeRes.TxResults[0], true)

	events := finalizeRes.TxResults[0].Events
	createEvent := events[len(events)-3]
	require.Equal(t, evmtypes.EventTypeContractCreated, createEvent.GetType())

	contractAddr, err := hexutil.Decode(createEvent.Attributes[0].Value)
	require.NoError(t, err)

	// listen finalize block
	ctx, err := app.CreateQueryContext(0, false)
	require.NoError(t, err)

	// check the tx is indexed
	evmTx, err := indexer.TxByHash(ctx, evmTxHash)
	require.NoError(t, err)
	require.NotNil(t, evmTx)

	// wait for pruning
	for {
		time.Sleep(100 * time.Millisecond)
		if indexer.IsPruningRunning() {
			continue
		} else {
			break
		}
	}

	// mint 1_000_000 tokens to the first address
	tx, evmTxHash2 := tests.GenerateMintERC20Tx(t, app, privKeys[0], common.BytesToAddress(contractAddr), addrs[0], new(big.Int).SetUint64(1_000_000_000_000))
	finalizeReq, finalizeRes := tests.ExecuteTxs(t, app, tx)
	tests.CheckTxResult(t, finalizeRes.TxResults[0], true)

	// wait for pruning
	for {
		time.Sleep(100 * time.Millisecond)

		if indexer.IsPruningRunning() {
			continue
		} else {
			break
		}
	}

	// listen finalize block
	ctx, err = app.CreateQueryContext(0, false)
	require.NoError(t, err)

	// check the block header is indexed
	header, err := indexer.BlockHeaderByNumber(ctx, uint64(finalizeReq.Height))
	require.NoError(t, err)
	require.NotNil(t, header)
	require.Equal(t, finalizeReq.Height, header.Number.Int64())

	// previous block should be pruned
	header, err = indexer.BlockHeaderByNumber(ctx, uint64(finalizeReq.Height-1))
	require.ErrorIs(t, err, collections.ErrNotFound)
	require.Nil(t, header)

	// check the tx is indexed
	evmTx, err = indexer.TxByHash(ctx, evmTxHash2)
	require.NoError(t, err)
	require.NotNil(t, evmTx)

	// but the first tx should be pruned
	evmTx, err = indexer.TxByHash(ctx, evmTxHash)
	require.ErrorIs(t, err, collections.ErrNotFound)
	require.Nil(t, evmTx)

	// check the receipt is indexed
	receipt, err := indexer.TxReceiptByHash(ctx, evmTxHash2)
	require.NoError(t, err)
	require.NotNil(t, receipt)

	// check the receipt is pruned
	_, err = indexer.TxReceiptByHash(ctx, evmTxHash)
	require.ErrorIs(t, err, collections.ErrNotFound)

	// check cosmos tx hash is indexed
	cosmosTxHash, err := indexer.CosmosTxHashByTxHash(ctx, evmTxHash2)
	require.NoError(t, err)
	require.NotNil(t, cosmosTxHash)
	evmTxHash3, err := indexer.TxHashByCosmosTxHash(ctx, cosmosTxHash)
	require.NoError(t, err)
	require.Equal(t, evmTxHash2, evmTxHash3)

	// check cosmos tx hash is pruned
	_, err = indexer.CosmosTxHashByTxHash(ctx, evmTxHash)
	require.ErrorIs(t, err, collections.ErrNotFound)
}

func Test_PruneIndexer_BloomBits(t *testing.T) {
	app, _, _ := tests.CreateApp(t)
	indexer := app.EVMIndexer().(*evmindexer.EVMIndexerImpl)
	defer app.Close()

	// set retain height to 1, only last block is indexed
	indexer.SetRetainHeight(1)

	startHeight := uint64(app.LastBlockHeight())
	nextSectionHeight := (startHeight/evmconfig.SectionSize + 1) * evmconfig.SectionSize
	for range evmconfig.SectionSize + 1 {
		tests.IncreaseBlockHeight(t, app)
	}

	// wait for bloom indexing
	for {
		time.Sleep(100 * time.Millisecond)
		if indexer.GetLastBloomIndexedHeight() < nextSectionHeight {
			continue
		} else {
			break
		}
	}

	// wait for pruning
	for {
		time.Sleep(100 * time.Millisecond)
		if indexer.GetLastPrunedHeight() < nextSectionHeight {
			continue
		} else {
			break
		}
	}

	// increase block height to trigger bloom indexing and pruning
	tests.IncreaseBlockHeight(t, app)

	// wait for bloom indexing
	for {
		time.Sleep(100 * time.Millisecond)
		if indexer.GetLastBloomIndexedHeight() < nextSectionHeight {
			continue
		} else {
			break
		}
	}

	// wait for pruning
	for {
		time.Sleep(100 * time.Millisecond)
		if indexer.GetLastPrunedHeight() < nextSectionHeight {
			continue
		} else {
			break
		}
	}

	// check the bloom bits are pruned
	ctx, err := app.CreateQueryContext(0, false)
	require.NoError(t, err)

	err = indexer.BloomBits.Walk(ctx, nil, func(key collections.Pair[uint64, uint32], value []byte) (bool, error) {
		require.Fail(t, "bloom bits should be pruned")
		return true, nil
	})
	require.NoError(t, err)
}
