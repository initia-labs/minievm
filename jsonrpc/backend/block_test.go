package backend_test

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rpc"

	rpctypes "github.com/initia-labs/minievm/jsonrpc/types"
	"github.com/initia-labs/minievm/tests"
	evmtypes "github.com/initia-labs/minievm/x/evm/types"
)

func Test_BlockNumber(t *testing.T) {
	input := setupBackend(t)
	app, _, backend, _, _ := input.app, input.addrs, input.backend, input.addrs, input.privKeys

	bn, err := backend.BlockNumber()
	require.NoError(t, err)
	require.Equal(t, uint64(app.LastBlockHeight()), uint64(bn))
}

func Test_GetHeaderByNumber(t *testing.T) {
	input := setupBackend(t)
	app, _, backend, _, privKeys := input.app, input.addrs, input.backend, input.addrs, input.privKeys

	tx, _ := tests.GenerateCreateERC20Tx(t, app, privKeys[0])
	finalizeReq, finalizeRes := tests.ExecuteTxs(t, app, tx)
	tests.CheckTxResult(t, finalizeRes.TxResults[0], true)

	header, err := backend.GetHeaderByNumber(rpc.BlockNumber(finalizeReq.Height))
	require.NoError(t, err)
	require.Equal(t, uint64(finalizeReq.Height), header.Number.Uint64())

	// get latest block header
	header2, err := backend.GetHeaderByNumber(rpc.LatestBlockNumber)
	require.NoError(t, err)
	require.Equal(t, header, header2)

	// pending block number should return the latest block number
	header3, err := backend.GetHeaderByNumber(rpc.PendingBlockNumber)
	require.NoError(t, err)
	require.Equal(t, header, header3)

	// safe block number should return the latest block number
	header4, err := backend.GetHeaderByNumber(rpc.SafeBlockNumber)
	require.NoError(t, err)
	require.Equal(t, header, header4)

	// finalized block number should return the latest block number
	header5, err := backend.GetHeaderByNumber(rpc.FinalizedBlockNumber)
	require.NoError(t, err)
	require.Equal(t, header, header5)

	// 0 block number should return genesis block
	header6, err := backend.GetHeaderByNumber(0)
	require.NoError(t, err)

	genesisHeader, err := backend.GetHeaderByNumber(rpc.BlockNumber(1))
	require.NoError(t, err)
	require.Equal(t, genesisHeader, header6)

	// invalid block number should return nil
	header7, err := backend.GetHeaderByNumber(rpc.BlockNumber(1000000))
	require.NoError(t, err)
	require.Nil(t, header7)
}

func Test_GetHeaderByHash(t *testing.T) {
	input := setupBackend(t)
	app, _, backend, _, privKeys := input.app, input.addrs, input.backend, input.addrs, input.privKeys

	tx, _ := tests.GenerateCreateERC20Tx(t, app, privKeys[0])
	finalizeReq, finalizeRes := tests.ExecuteTxs(t, app, tx)
	tests.CheckTxResult(t, finalizeRes.TxResults[0], true)

	header, err := backend.GetHeaderByNumber(rpc.BlockNumber(finalizeReq.Height))
	require.NoError(t, err)

	header2, err := backend.GetHeaderByHash(header.Hash())
	require.NoError(t, err)
	require.Equal(t, header, header2)
}

func Test_GetBlockByNumber(t *testing.T) {
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
	tx, evmTxHash := tests.GenerateMintERC20Tx(t, app, privKeys[0], common.BytesToAddress(contractAddr), addrs[0], new(big.Int).SetUint64(1_000_000_000_000))
	tx2, evmTxHash2 := tests.GenerateTransferERC20Tx(t, app, privKeys[0], common.BytesToAddress(contractAddr), addrs[1], new(big.Int).SetUint64(1_000_000), tests.SetNonce(2))
	finalizeReq, finalizeRes := tests.ExecuteTxs(t, app, tx, tx2)
	tests.CheckTxResult(t, finalizeRes.TxResults[0], true)
	tests.CheckTxResult(t, finalizeRes.TxResults[1], true)

	header, err := backend.GetHeaderByNumber(rpc.BlockNumber(finalizeReq.Height))
	require.NoError(t, err)
	evmTx, err := backend.GetTransactionByHash(evmTxHash)
	require.NoError(t, err)
	evmTx2, err := backend.GetTransactionByHash(evmTxHash2)
	require.NoError(t, err)

	block, err := backend.GetBlockByNumber(rpc.BlockNumber(finalizeReq.Height), true)
	require.NoError(t, err)

	require.Equal(t, (*hexutil.Big)(header.Number), block["number"])
	require.Equal(t, header.Hash(), block["hash"])
	require.Equal(t, hexutil.Uint64(header.GasUsed), block["gasUsed"])
	require.Equal(t, hexutil.Uint64(header.GasLimit), block["gasLimit"])
	require.Equal(t, hexutil.Uint64(header.Time), block["timestamp"])
	require.Equal(t, []*rpctypes.RPCTransaction{evmTx, evmTx2}, block["transactions"])

	block, err = backend.GetBlockByNumber(rpc.BlockNumber(finalizeReq.Height), false)
	require.NoError(t, err)
	require.Equal(t, []common.Hash{evmTx.Hash, evmTx2.Hash}, block["transactions"])
}

func Test_GetBlockByHash(t *testing.T) {
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
	tx, evmTxHash := tests.GenerateMintERC20Tx(t, app, privKeys[0], common.BytesToAddress(contractAddr), addrs[0], new(big.Int).SetUint64(1_000_000_000_000))
	tx2, evmTxHash2 := tests.GenerateTransferERC20Tx(t, app, privKeys[0], common.BytesToAddress(contractAddr), addrs[1], new(big.Int).SetUint64(1_000_000), tests.SetNonce(2))
	finalizeReq, finalizeRes := tests.ExecuteTxs(t, app, tx, tx2)
	tests.CheckTxResult(t, finalizeRes.TxResults[0], true)
	tests.CheckTxResult(t, finalizeRes.TxResults[1], true)

	header, err := backend.GetHeaderByNumber(rpc.BlockNumber(finalizeReq.Height))
	require.NoError(t, err)
	evmTx, err := backend.GetTransactionByHash(evmTxHash)
	require.NoError(t, err)
	evmTx2, err := backend.GetTransactionByHash(evmTxHash2)
	require.NoError(t, err)

	block, err := backend.GetBlockByHash(header.Hash(), true)
	require.NoError(t, err)

	require.Equal(t, (*hexutil.Big)(header.Number), block["number"])
	require.Equal(t, header.Hash(), block["hash"])
	require.Equal(t, hexutil.Uint64(header.GasUsed), block["gasUsed"])
	require.Equal(t, hexutil.Uint64(header.GasLimit), block["gasLimit"])
	require.Equal(t, hexutil.Uint64(header.Time), block["timestamp"])
	require.Equal(t, []*rpctypes.RPCTransaction{evmTx, evmTx2}, block["transactions"])

	block, err = backend.GetBlockByHash(header.Hash(), false)
	require.NoError(t, err)
	require.Equal(t, []common.Hash{evmTx.Hash, evmTx2.Hash}, block["transactions"])
}
