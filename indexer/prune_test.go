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
	finalizeReq, finalizeRes := tests.ExecuteTxs(t, app, tx)
	tests.CheckTxResult(t, finalizeRes.TxResults[0], true)

	events := finalizeRes.TxResults[0].Events
	createEvent := events[len(events)-3]
	require.Equal(t, evmtypes.EventTypeContractCreated, createEvent.GetType())

	contractAddr, err := hexutil.Decode(createEvent.Attributes[0].Value)
	require.NoError(t, err)

	// check the tx is indexed
	// check the tx is indexed
	evmTx, err := indexer.TxByHash(evmTxHash)
	require.NoError(t, err)
	require.NotNil(t, evmTx)

	require.Eventually(t, func() bool {
		return indexer.GetLastPruneTriggerHeight() >= uint64(finalizeReq.Height)
	}, 10*time.Second, 50*time.Millisecond)

	// mint 1_000_000 tokens to the first address
	tx, evmTxHash2 := tests.GenerateMintERC20Tx(t, app, privKeys[0], common.BytesToAddress(contractAddr), addrs[0], new(big.Int).SetUint64(1_000_000_000_000))
	finalizeReq, finalizeRes = tests.ExecuteTxs(t, app, tx)
	tests.CheckTxResult(t, finalizeRes.TxResults[0], true)

	require.Eventually(t, func() bool {
		return indexer.GetLastPruneTriggerHeight() >= uint64(finalizeReq.Height)
	}, 10*time.Second, 50*time.Millisecond)

	// check the block header is indexed
	header, err := indexer.BlockHeaderByNumber(uint64(finalizeReq.Height))
	require.NoError(t, err)
	require.NotNil(t, header)
	require.Equal(t, finalizeReq.Height, header.Number.Int64())

	// previous block should be pruned
	header, err = indexer.BlockHeaderByNumber(uint64(finalizeReq.Height - 1))
	require.ErrorIs(t, err, collections.ErrNotFound)
	require.Nil(t, header)

	// check the tx is indexed
	evmTx, err = indexer.TxByHash(evmTxHash2)
	require.NoError(t, err)
	require.NotNil(t, evmTx)

	// but the first tx should be pruned
	evmTx, err = indexer.TxByHash(evmTxHash)
	require.ErrorIs(t, err, collections.ErrNotFound)
	require.Nil(t, evmTx)

	// check the receipt is indexed
	receipt, err := indexer.TxReceiptByHash(evmTxHash2)
	require.NoError(t, err)
	require.NotNil(t, receipt)

	// check the receipt is pruned
	_, err = indexer.TxReceiptByHash(evmTxHash)
	require.ErrorIs(t, err, collections.ErrNotFound)

	// check cosmos tx hash is indexed
	cosmosTxHash, err := indexer.CosmosTxHashByTxHash(evmTxHash2)
	require.NoError(t, err)
	require.NotNil(t, cosmosTxHash)
	evmTxHash3, err := indexer.TxHashByCosmosTxHash(cosmosTxHash)
	require.NoError(t, err)
	require.Equal(t, evmTxHash2, evmTxHash3)

	// check cosmos tx hash is pruned
	_, err = indexer.CosmosTxHashByTxHash(evmTxHash)
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

	require.Eventually(t, func() bool {
		if indexer.GetLastBloomIndexedHeight() < nextSectionHeight {
			tests.IncreaseBlockHeight(t, app)
			return false
		}
		return true
	}, 20*time.Second, 100*time.Millisecond)

	require.Eventually(t, func() bool {
		if indexer.GetLastPruneTriggerHeight() < nextSectionHeight {
			tests.IncreaseBlockHeight(t, app)
			return false
		}
		return true
	}, 20*time.Second, 100*time.Millisecond)

	// increase block height to trigger bloom indexing and pruning
	tests.IncreaseBlockHeight(t, app)
	postTriggerHeight := nextSectionHeight + 1
	for uint64(app.LastBlockHeight()) < postTriggerHeight {
		tests.IncreaseBlockHeight(t, app)
	}

	require.Eventually(t, func() bool {
		return indexer.GetLastPruneTriggerHeight() >= postTriggerHeight
	}, 20*time.Second, 100*time.Millisecond)

	// check the bloom bits are pruned
	prunedSection := nextSectionHeight/evmconfig.SectionSize - 1
	require.Eventually(t, func() bool {
		ctx, closer, err := app.CreateQueryContext(0, false)
		if err != nil {
			return false
		}
		if closer != nil {
			defer closer.Close()
		}

		found := false
		err = indexer.BloomBits.Walk(ctx, collections.NewPrefixedPairRange[uint64, uint32](prunedSection), func(key collections.Pair[uint64, uint32], value []byte) (bool, error) {
			found = true
			return true, nil
		})
		if err != nil {
			return false
		}
		return !found
	}, 20*time.Second, 100*time.Millisecond, "section %d bloom bits should be pruned", prunedSection)
}
