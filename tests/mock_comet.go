package tests

import (
	"context"

	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cometbft/cometbft/libs/bytes"
	"github.com/cometbft/cometbft/p2p"
	rpcclient "github.com/cometbft/cometbft/rpc/client"
	ctypes "github.com/cometbft/cometbft/rpc/core/types"
	"github.com/cometbft/cometbft/types"

	"github.com/skip-mev/block-sdk/v2/block"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types/mempool"

	"github.com/initia-labs/minievm/indexer"
)

var _ client.CometRPC = &MockCometRPC{}
var _ rpcclient.MempoolClient = &MockCometRPC{}
var _ rpcclient.NetworkClient = &MockCometRPC{}

type MockCometRPC struct {
	app *baseapp.BaseApp

	NPeers        int
	Listening     bool
	ClientVersion string
}

func NewMockCometRPC(app *baseapp.BaseApp) *MockCometRPC {
	return &MockCometRPC{app: app}
}

// setters
func (m *MockCometRPC) WithNPeers(n int) *MockCometRPC {
	m.NPeers = n
	return m
}
func (m *MockCometRPC) WithListening(listening bool) *MockCometRPC {
	m.Listening = listening
	return m
}
func (m *MockCometRPC) WithClientVersion(version string) *MockCometRPC {
	m.ClientVersion = version
	return m
}

// CometRPC methods
func (m *MockCometRPC) Status(context.Context) (*ctypes.ResultStatus, error) {
	return &ctypes.ResultStatus{
		NodeInfo: p2p.DefaultNodeInfo{
			Version: m.ClientVersion,
		},
	}, nil
}
func (m *MockCometRPC) BroadcastTxSync(ctx context.Context, tx types.Tx) (*ctypes.ResultBroadcastTx, error) {
	res, err := m.app.CheckTx(&abci.RequestCheckTx{
		Tx:   tx,
		Type: abci.CheckTxType_New,
	})
	if err != nil {
		return nil, err
	}
	return &ctypes.ResultBroadcastTx{
		Code:      res.Code,
		Log:       res.Log,
		Data:      res.Data,
		Codespace: res.Codespace,
		Hash:      tx.Hash(),
	}, nil
}
func (m *MockCometRPC) UnconfirmedTxs(ctx context.Context, limit *int) (*ctypes.ResultUnconfirmedTxs, error) {
	var ok bool
	var mempool mempool.Mempool
	if mempool, ok = m.app.Mempool().(*indexer.MempoolWrapper); ok {
		mempool = mempool.(*indexer.MempoolWrapper).Inner()
	} else {
		mempool = m.app.Mempool()
	}

	laneMempool := mempool.(*block.LanedMempool)
	lanes := laneMempool.Registry()
	txs := make([]types.Tx, 0)
	for _, lane := range lanes {
		iter := lane.Select(ctx, nil)
		for ; iter != nil; iter = iter.Next() {
			tx, err := m.app.TxEncode(iter.Tx())
			if err != nil {
				return nil, err
			}

			txs = append(txs, tx)
		}
	}

	return &ctypes.ResultUnconfirmedTxs{
		Txs: txs,
	}, nil
}
func (m *MockCometRPC) NumUnconfirmedTxs(context.Context) (*ctypes.ResultUnconfirmedTxs, error) {
	count := m.app.Mempool().CountTx()
	return &ctypes.ResultUnconfirmedTxs{
		Count: count,
		Total: count,
	}, nil
}
func (m *MockCometRPC) NetInfo(context.Context) (*ctypes.ResultNetInfo, error) {
	return &ctypes.ResultNetInfo{
		NPeers:    m.NPeers,
		Listening: m.Listening,
	}, nil
}

// unused methods
func (m *MockCometRPC) DumpConsensusState(context.Context) (*ctypes.ResultDumpConsensusState, error) {
	panic("implement me")
}
func (m *MockCometRPC) ConsensusState(context.Context) (*ctypes.ResultConsensusState, error) {
	panic("implement me")
}
func (m *MockCometRPC) ConsensusParams(ctx context.Context, height *int64) (*ctypes.ResultConsensusParams, error) {
	panic("implement me")
}
func (m *MockCometRPC) Health(context.Context) (*ctypes.ResultHealth, error) {
	panic("implement me")
}
func (m *MockCometRPC) CheckTx(context.Context, types.Tx) (*ctypes.ResultCheckTx, error) {
	panic("implement me")
}
func (m *MockCometRPC) BroadcastTxCommit(context.Context, types.Tx) (*ctypes.ResultBroadcastTxCommit, error) {
	panic("implement me")
}
func (m *MockCometRPC) BroadcastTxAsync(context.Context, types.Tx) (*ctypes.ResultBroadcastTx, error) {
	panic("implement me")
}
func (m *MockCometRPC) ABCIInfo(context.Context) (*ctypes.ResultABCIInfo, error) {
	panic("implement me")
}
func (m *MockCometRPC) ABCIQuery(ctx context.Context, path string, data bytes.HexBytes) (*ctypes.ResultABCIQuery, error) {
	panic("implement me")
}
func (m *MockCometRPC) ABCIQueryWithOptions(ctx context.Context, path string, data bytes.HexBytes,
	opts rpcclient.ABCIQueryOptions) (*ctypes.ResultABCIQuery, error) {
	panic("implement me")
}
func (m *MockCometRPC) Validators(ctx context.Context, height *int64, page, perPage *int) (*ctypes.ResultValidators, error) {
	panic("implement me")
}

func (m *MockCometRPC) Block(ctx context.Context, height *int64) (*ctypes.ResultBlock, error) {
	panic("implement me")
}
func (m *MockCometRPC) BlockByHash(ctx context.Context, hash []byte) (*ctypes.ResultBlock, error) {
	panic("implement me")
}
func (m *MockCometRPC) BlockResults(ctx context.Context, height *int64) (*ctypes.ResultBlockResults, error) {
	panic("implement me")
}
func (m *MockCometRPC) BlockchainInfo(ctx context.Context, minHeight, maxHeight int64) (*ctypes.ResultBlockchainInfo, error) {
	panic("implement me")
}
func (m *MockCometRPC) Commit(ctx context.Context, height *int64) (*ctypes.ResultCommit, error) {
	panic("implement me")
}
func (m *MockCometRPC) Tx(ctx context.Context, hash []byte, prove bool) (*ctypes.ResultTx, error) {
	panic("implement me")
}
func (m *MockCometRPC) TxSearch(
	ctx context.Context,
	query string,
	prove bool,
	page, perPage *int,
	orderBy string,
) (*ctypes.ResultTxSearch, error) {
	panic("implement me")
}
func (m *MockCometRPC) BlockSearch(
	ctx context.Context,
	query string,
	page, perPage *int,
	orderBy string,
) (*ctypes.ResultBlockSearch, error) {
	panic("implement me")
}
