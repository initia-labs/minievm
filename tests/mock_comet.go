package tests

import (
	"context"

	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cometbft/cometbft/libs/bytes"
	"github.com/cometbft/cometbft/p2p"
	rpcclient "github.com/cometbft/cometbft/rpc/client"
	ctypes "github.com/cometbft/cometbft/rpc/core/types"
	"github.com/cometbft/cometbft/types"
	"github.com/ethereum/go-ethereum/common"

	"github.com/skip-mev/block-sdk/v2/block"

	minitiaapp "github.com/initia-labs/minievm/app"
	rpctypes "github.com/initia-labs/minievm/jsonrpc/types"
	evmkeeper "github.com/initia-labs/minievm/x/evm/keeper"

	"github.com/cosmos/cosmos-sdk/client"
)

var _ client.CometRPC = &MockCometRPC{}
var _ rpcclient.MempoolClient = &MockCometRPC{}
var _ rpcclient.NetworkClient = &MockCometRPC{}

type MockCometRPC struct {
	app *minitiaapp.MinitiaApp

	NPeers        int
	Listening     bool
	ClientVersion string

	txs [][]byte
}

func NewMockCometRPC(app *minitiaapp.MinitiaApp) *MockCometRPC {
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

	// save tx to be rechecked
	if res.Code == abci.CodeTypeOK && res.Codespace == "txqueue" {
		m.txs = append(m.txs, tx)
	}

	return &ctypes.ResultBroadcastTx{
		Code:      res.Code,
		Log:       res.Log,
		Data:      res.Data,
		Codespace: res.Codespace,
		Hash:      tx.Hash(),
	}, nil
}
func (m *MockCometRPC) RecheckTx() error {
	remainTxs := make([][]byte, 0)
	for _, tx := range m.txs {
		if res, err := m.app.CheckTx(&abci.RequestCheckTx{
			Tx:   tx,
			Type: abci.CheckTxType_Recheck,
		}); err != nil {
			return err
		} else if res.Code == 0 && res.Codespace == "txqueue" {
			remainTxs = append(remainTxs, tx)
		}
	}

	m.txs = remainTxs

	return nil
}
func (m *MockCometRPC) UnconfirmedTxs(ctx context.Context, limit *int) (*ctypes.ResultUnconfirmedTxs, error) {
	mempool := m.app.Mempool()
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

// for only mock comet rpc, use evm indexer db to get transactions
func (m *MockCometRPC) Block(ctx context.Context, height *int64) (*ctypes.ResultBlock, error) {
	h := m.app.LastBlockHeight()
	if height != nil {
		h = *height
	}

	txs := types.Txs{}
	err := m.app.EVMIndexer().IterateBlockTxs(ctx, uint64(h), func(tx *rpctypes.RPCTransaction) (bool, error) {
		ethTx := tx.ToTransaction()
		cosmosTx, err := evmkeeper.NewTxUtils(m.app.EVMKeeper).ConvertEthereumTxToCosmosTx(ctx, ethTx)
		if err != nil {
			return true, err
		}

		bz, err := m.app.TxConfig().TxEncoder()(cosmosTx)
		if err != nil {
			return true, err
		}

		txs = append(txs, bz)
		return false, nil
	})
	if err != nil {
		return nil, err
	}

	block := &types.Block{Data: types.Data{Txs: txs}}
	return &ctypes.ResultBlock{BlockID: types.BlockID{}, Block: block}, nil
}
func (m *MockCometRPC) BlockByHash(ctx context.Context, hash []byte) (*ctypes.ResultBlock, error) {
	blockNumber, err := m.app.EVMIndexer().BlockHashToNumber(ctx, common.BytesToHash(hash))
	if err != nil {
		return nil, err
	}

	h := int64(blockNumber)
	return m.Block(ctx, &h)
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
