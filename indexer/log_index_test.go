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
	evmtypes "github.com/initia-labs/minievm/x/evm/types"
)

// Test_TxStartLogIndex verifies that TxStartLogIndexByHash stores a block-scoped
// starting log index for each tx. When multiple txs in the same block each emit
// logs, the second tx's start index must be the total number of logs in all
// preceding txs — not zero.
func Test_TxStartLogIndex(t *testing.T) {
	app, addrs, privKeys := tests.CreateApp(t)
	indexer := app.EVMIndexer()
	defer app.Close()

	// deploy ERC20
	tx, _ := tests.GenerateCreateERC20Tx(t, app, privKeys[0])
	_, finalizeRes := tests.ExecuteTxs(t, app, tx)
	tests.CheckTxResult(t, finalizeRes.TxResults[0], true)

	events := finalizeRes.TxResults[0].Events
	createEvent := events[len(events)-3]
	require.Equal(t, evmtypes.EventTypeContractCreated, createEvent.GetType())
	contractAddr, err := hexutil.Decode(createEvent.Attributes[0].Value)
	require.NoError(t, err)
	contract := common.BytesToAddress(contractAddr)

	// mint tokens so addrs[0] can transfer
	tx, _ = tests.GenerateMintERC20Tx(t, app, privKeys[0], contract, addrs[0], new(big.Int).SetUint64(1_000_000_000))
	_, finalizeRes = tests.ExecuteTxs(t, app, tx)
	tests.CheckTxResult(t, finalizeRes.TxResults[0], true)

	// Execute two transfers in the SAME block so we can check block-scoped indices.
	// Each ERC20 Transfer emits exactly 1 log.
	tx1, evmHash1 := tests.GenerateTransferERC20Tx(t, app, privKeys[0], contract, addrs[1], new(big.Int).SetUint64(1_000), tests.SetNonce(2))
	tx2, evmHash2 := tests.GenerateTransferERC20Tx(t, app, privKeys[0], contract, addrs[1], new(big.Int).SetUint64(1_000), tests.SetNonce(3))
	_, finalizeRes = tests.ExecuteTxs(t, app, tx1, tx2)
	tests.CheckTxResult(t, finalizeRes.TxResults[0], true)
	tests.CheckTxResult(t, finalizeRes.TxResults[1], true)

	// tx1 is the first tx in the block — its start log index must be 0
	start1, err := indexer.TxStartLogIndexByHash(evmHash1)
	require.NoError(t, err)
	require.Equal(t, uint64(0), start1, "first tx start log index should be 0")

	// tx2 follows tx1 which emitted 1 log — its start log index must be 1
	receipt1, err := indexer.TxReceiptByHash(evmHash1)
	require.NoError(t, err)
	logsInTx1 := uint64(len(receipt1.Logs))
	require.Greater(t, logsInTx1, uint64(0), "tx1 must have emitted at least one log")

	start2, err := indexer.TxStartLogIndexByHash(evmHash2)
	require.NoError(t, err)
	require.Equal(t, logsInTx1, start2, "second tx start log index should equal number of logs in first tx")
}

// Test_TxStartLogIndex_Pruned verifies that TxStartLogIndexMap entries are
// removed together with the rest of the tx data when pruning runs.
func Test_TxStartLogIndex_Pruned(t *testing.T) {
	app, addrs, privKeys := tests.CreateApp(t)
	indexer := app.EVMIndexer().(*evmindexer.EVMIndexerImpl)
	defer app.Close()

	indexer.SetRetainHeight(1)

	// deploy ERC20
	tx, _ := tests.GenerateCreateERC20Tx(t, app, privKeys[0])
	_, finalizeRes := tests.ExecuteTxs(t, app, tx)
	tests.CheckTxResult(t, finalizeRes.TxResults[0], true)

	events := finalizeRes.TxResults[0].Events
	createEvent := events[len(events)-3]
	require.Equal(t, evmtypes.EventTypeContractCreated, createEvent.GetType())
	contractAddr, err := hexutil.Decode(createEvent.Attributes[0].Value)
	require.NoError(t, err)
	contract := common.BytesToAddress(contractAddr)

	// mint tokens
	tx, _ = tests.GenerateMintERC20Tx(t, app, privKeys[0], contract, addrs[0], new(big.Int).SetUint64(1_000_000_000))
	_, finalizeRes = tests.ExecuteTxs(t, app, tx)
	tests.CheckTxResult(t, finalizeRes.TxResults[0], true)

	// tx in block N (will be pruned)
	tx, evmHashPruned := tests.GenerateTransferERC20Tx(t, app, privKeys[0], contract, addrs[1], new(big.Int).SetUint64(1_000), tests.SetNonce(2))
	_, finalizeRes = tests.ExecuteTxs(t, app, tx)
	tests.CheckTxResult(t, finalizeRes.TxResults[0], true)

	// tx in block N+1 (retained)
	tx, evmHashKept := tests.GenerateTransferERC20Tx(t, app, privKeys[0], contract, addrs[1], new(big.Int).SetUint64(1_000), tests.SetNonce(3))
	finalizeReq, finalizeRes := tests.ExecuteTxs(t, app, tx)
	tests.CheckTxResult(t, finalizeRes.TxResults[0], true)

	require.Eventually(t, func() bool {
		return indexer.GetLastPruneTriggerHeight() >= uint64(finalizeReq.Height)
	}, 10*time.Second, 50*time.Millisecond, "timed out waiting for pruning to finish")

	// pruned tx's start log index should be gone
	_, err = indexer.TxStartLogIndexByHash(evmHashPruned)
	require.ErrorIs(t, err, collections.ErrNotFound, "start log index for pruned tx should be removed")

	// kept tx's start log index should still be present
	_, err = indexer.TxStartLogIndexByHash(evmHashKept)
	require.NoError(t, err, "start log index for retained tx should still exist")
}
