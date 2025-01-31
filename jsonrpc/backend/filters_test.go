package backend_test

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/bloombits"
	coretypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/initia-labs/minievm/tests"
	evmconfig "github.com/initia-labs/minievm/x/evm/config"
	evmtypes "github.com/initia-labs/minievm/x/evm/types"
	"github.com/stretchr/testify/require"
)

func Test_GetLogsByHeight(t *testing.T) {
	input := setupBackend(t)
	app, _, backend, addrs, privKeys := input.app, input.addrs, input.backend, input.addrs, input.privKeys

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
	tx2, _ := tests.GenerateMintERC20Tx(t, app, privKeys[0], common.BytesToAddress(contractAddr), addrs[0], new(big.Int).SetUint64(1_000_000_000_000), tests.SetNonce(2))
	_, finalizeRes = tests.ExecuteTxs(t, app, tx, tx2)
	tests.CheckTxResult(t, finalizeRes.TxResults[0], true)
	tests.CheckTxResult(t, finalizeRes.TxResults[1], true)

	logs, err := backend.GetLogsByHeight(uint64(app.LastBlockHeight()))
	require.NoError(t, err)

	receipts, err := backend.GetBlockReceipts(rpc.BlockNumberOrHashWithNumber(rpc.BlockNumber(app.LastBlockHeight())))
	require.NoError(t, err)

	blockLogs := []*coretypes.Log{}
	logs0, ok := receipts[0]["logs"].([]*coretypes.Log)
	require.True(t, ok)
	logs1, ok := receipts[1]["logs"].([]*coretypes.Log)
	require.True(t, ok)

	blockLogs = append(blockLogs, logs0...)
	blockLogs = append(blockLogs, logs1...)
	require.Equal(t, blockLogs, logs)
}

func Test_BloomStatus(t *testing.T) {
	input := setupBackend(t)
	app, indexer, backend, addrs, privKeys := input.app, input.indexer, input.backend, input.addrs, input.privKeys

	tx1, _ := tests.GenerateCreateERC20Tx(t, app, privKeys[0])
	_, finalizeRes1 := tests.ExecuteTxs(t, app, tx1)
	tests.CheckTxResult(t, finalizeRes1.TxResults[0], true)
	height1 := app.LastBlockHeight()

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

	tx2, _ := tests.GenerateCreateERC20Tx(t, app, privKeys[1])
	_, finalizeRes2 := tests.ExecuteTxs(t, app, tx2)
	tests.CheckTxResult(t, finalizeRes2.TxResults[0], true)
	height2 := app.LastBlockHeight()

	size, section, err := backend.BloomStatus()
	require.NoError(t, err)
	require.Equal(t, evmconfig.SectionSize, size)
	require.Equal(t, uint64(1), section)

	filters := make([][][]byte, 1)
	filters[0] = [][]byte{addrs[0].Bytes()}
	matcher := bloombits.NewMatcher(evmconfig.SectionSize, filters)

	// start find matches
	matches := make(chan uint64, 64)
	session1, err := matcher.Start(context.Background(), uint64(1), uint64(height2), matches)
	require.NoError(t, err)
	defer session1.Close()

	backend.ServiceFilter(session1)

LOOP1:
	for {
		number, ok := <-matches
		if !ok {
			err := session1.Error()
			require.NoError(t, err)
			break LOOP1
		}

		require.True(t, ok)
		require.Equal(t, height1, number)
	}

	filters = make([][][]byte, 1)
	filters[0] = [][]byte{addrs[1].Bytes()}
	matcher2 := bloombits.NewMatcher(evmconfig.SectionSize, filters)

	// start find matches
	matches = make(chan uint64, 64)
	session2, err := matcher2.Start(context.Background(), uint64(1), uint64(height2), matches)
	require.NoError(t, err)
	defer session2.Close()

	backend.ServiceFilter(session2)

LOOP2:
	for {
		number, ok := <-matches
		if !ok {
			err := session2.Error()
			require.NoError(t, err)

			break LOOP2
		}

		require.True(t, ok)
		require.Equal(t, height2, number)
	}
}
