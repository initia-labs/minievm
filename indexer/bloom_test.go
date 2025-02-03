package indexer_test

import (
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	coretypes "github.com/ethereum/go-ethereum/core/types"
	evmindexer "github.com/initia-labs/minievm/indexer"

	"github.com/initia-labs/minievm/tests"
	evmconfig "github.com/initia-labs/minievm/x/evm/config"
	evmtypes "github.com/initia-labs/minievm/x/evm/types"

	"github.com/stretchr/testify/require"
)

func Test_BloomIndexing(t *testing.T) {
	app, addrs, privKeys := tests.CreateApp(t)
	indexer := app.EVMIndexer().(*evmindexer.EVMIndexerImpl)
	defer app.Close()

	tx, _ := tests.GenerateCreateERC20Tx(t, app, privKeys[0])
	_, finalizeRes := tests.ExecuteTxs(t, app, tx)
	tests.CheckTxResult(t, finalizeRes.TxResults[0], true)

	events := finalizeRes.TxResults[0].Events
	createEvent := events[len(events)-3]
	require.Equal(t, evmtypes.EventTypeContractCreated, createEvent.GetType())

	contractAddr, err := hexutil.Decode(createEvent.Attributes[0].Value)
	require.NoError(t, err)

	// mint 1_000_000 tokens to the first address
	tx, _ = tests.GenerateMintERC20Tx(t, app, privKeys[0], common.BytesToAddress(contractAddr), addrs[0], new(big.Int).SetUint64(1_000_000_000_000))
	_, finalizeRes = tests.ExecuteTxs(t, app, tx)
	tests.CheckTxResult(t, finalizeRes.TxResults[0], true)

	for i := uint64(0); i < evmconfig.SectionSize; i++ {
		tests.IncreaseBlockHeight(t, app)
	}

	// wait for bloom indexing
	for {
		if indexer.IsBloomIndexingRunning() {
			time.Sleep(100 * time.Millisecond)
		} else {
			break
		}
	}

	// create a new block to trigger bloom indexing
	tests.IncreaseBlockHeight(t, app)

	// wait for bloom indexing
	for {
		if indexer.IsBloomIndexingRunning() {
			time.Sleep(100 * time.Millisecond)
		} else {
			break
		}
	}

	ctx, err := app.CreateQueryContext(0, false)
	require.NoError(t, err)

	for i := uint32(0); i < coretypes.BloomBitLength; i++ {
		bloomBits, err := indexer.ReadBloomBits(ctx, 0, i)
		require.NoError(t, err)
		require.NotNil(t, bloomBits)
	}

	nextSection, err := indexer.PeekBloomBitsNextSection(ctx)
	require.NoError(t, err)
	require.Equal(t, uint64(1), nextSection)
}
