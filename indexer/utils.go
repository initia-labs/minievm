package indexer

import (
	"encoding/json"
	"fmt"
	"math/big"
	"sync"

	abci "github.com/cometbft/cometbft/abci/types"
	comettypes "github.com/cometbft/cometbft/types"

	collcodec "cosmossdk.io/collections/codec"
	"cosmossdk.io/log"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	coretypes "github.com/ethereum/go-ethereum/core/types"

	"github.com/initia-labs/minievm/x/evm/keeper"
	"github.com/initia-labs/minievm/x/evm/types"

	"github.com/valyala/fastjson"
)

const (
	workerCount = 4
)

// EthTxInfo is the information of an Ethereum transaction.
type EthTxInfo struct {
	Tx           *coretypes.Transaction
	Logs         []*coretypes.Log
	ContractAddr *common.Address
	Status       uint64
	GasUsed      uint64
	CosmosTxHash []byte
}

// extractEthTxInfos extracts Ethereum transaction information from the finalize block request and response.
// It uses a worker pool to extract the transaction information in parallel.
func extractEthTxInfos(
	sdkCtx sdk.Context,
	logger log.Logger,
	txDecoder sdk.TxDecoder,
	evmKeeper keeper.Keeper,
	req abci.RequestFinalizeBlock,
	res abci.ResponseFinalizeBlock,
) ([]*EthTxInfo, error) {
	numTxs := len(req.Txs)
	if numTxs == 0 {
		return nil, nil
	}

	// use min of workerCount or numTxs
	workerCount := min(workerCount, numTxs)

	ethTxInfos := make([]*EthTxInfo, numTxs)
	wg := sync.WaitGroup{}
	wg.Add(numTxs)

	// Buffered channel to avoid blocking
	jobs := make(chan struct {
		idx     int
		txBytes []byte
		result  *abci.ExecTxResult
	}, numTxs)

	// Start worker pool
	for range workerCount {
		go func() {
			for job := range jobs {
				ethTxInfo, err := extractEthTxInfo(sdkCtx, txDecoder, evmKeeper, job.txBytes, job.result)
				if err != nil {
					logger.Error("failed to extract tx info",
						"height", req.Height,
						"index", job.idx,
						"error", err,
					)
				} else if ethTxInfo != nil {
					ethTxInfos[job.idx] = ethTxInfo
				}

				wg.Done()
			}
		}()
	}

	// Send jobs to workers
	for idx, txBytes := range req.Txs {
		jobs <- struct {
			idx     int
			txBytes []byte
			result  *abci.ExecTxResult
		}{
			idx:     idx,
			txBytes: txBytes,
			result:  res.TxResults[idx],
		}
	}

	close(jobs)
	wg.Wait()

	// remove nil from ethTxInfos
	filteredTxInfos := make([]*EthTxInfo, 0, len(ethTxInfos))
	for _, txInfo := range ethTxInfos {
		if txInfo != nil {
			filteredTxInfos = append(filteredTxInfos, txInfo)
		}
	}
	ethTxInfos = filteredTxInfos
	return ethTxInfos, nil
}

// extractEthTxInfo extracts Ethereum transaction information from a Cosmos SDK transaction.
// It decodes the transaction, checks its status, and converts it to an Ethereum transaction format.
// For successful Ethereum transactions, it extracts logs and contract address information.
// For non-Ethereum Cosmos transactions, it processes EVM events if the transaction was successful and
// creates a fake Ethereum transaction for indexing purposes.
func extractEthTxInfo(
	ctx sdk.Context,
	txDecoder sdk.TxDecoder,
	evmKeeper keeper.Keeper,
	txBytes []byte,
	txResult *abci.ExecTxResult,
) (*EthTxInfo, error) {
	sdkTx, err := txDecoder(txBytes)
	if err != nil {
		return nil, err
	}

	txStatus := coretypes.ReceiptStatusSuccessful
	if txResult.Code != abci.CodeTypeOK {
		txStatus = coretypes.ReceiptStatusFailed
	}

	// convert cosmos tx to ethereum tx
	ethTx, _, err := evmKeeper.TxUtils().ConvertCosmosTxToEthereumTx(ctx, sdkTx)
	if err != nil {
		return nil, err
	}

	// if the tx is normal ethereum tx, then return a single EthTxInfo
	if ethTx != nil {
		ethLogs, contractAddr, err := extractLogsAndContractAddr(txStatus, txResult.Events, ethTx.To() == nil)
		if err != nil {
			return nil, err
		}

		return &EthTxInfo{
			Tx:           ethTx,
			Logs:         ethLogs,
			ContractAddr: contractAddr,
			Status:       txStatus,
			GasUsed:      uint64(txResult.GasUsed),
			CosmosTxHash: comettypes.Tx(txBytes).Hash(),
		}, nil
	}

	// if the tx is not successful, then we don't need to create fake eth tx for cosmos tx
	if txStatus != coretypes.ReceiptStatusSuccessful {
		return nil, nil
	}

	// extract the logs from the events
	ethLogs, _, err := extractLogsAndContractAddr(txStatus, txResult.Events, false)
	if err != nil {
		return nil, err
	} else if len(ethLogs) == 0 {
		return nil, nil
	}

	// if there is evm events, then create a fake eth tx
	cosmosTxHash := comettypes.Tx(txBytes).Hash()
	ethTx = coretypes.NewTx(&coretypes.LegacyTx{
		Data:     cosmosTxHash,
		To:       nil,
		Nonce:    0,
		Gas:      0,
		GasPrice: new(big.Int),
		Value:    new(big.Int),
		R:        new(big.Int),
		S:        new(big.Int),
		V:        new(big.Int),
	})

	return &EthTxInfo{
		Tx:           ethTx,
		Logs:         ethLogs,
		Status:       txStatus,
		ContractAddr: nil,
		GasUsed:      uint64(txResult.GasUsed),
		CosmosTxHash: cosmosTxHash,
	}, nil
}

// extractLogsAndContractAddr extracts logs and contract address from the events
func extractLogsAndContractAddr(txStatus uint64, events []abci.Event, isContractCreation bool) ([]*coretypes.Log, *common.Address, error) {
	if txStatus != coretypes.ReceiptStatusSuccessful {
		return []*coretypes.Log{}, nil, nil
	}

	logs := []*coretypes.Log{}
	var contractAddr *common.Address

	parser := fastjson.Parser{}
	for _, event := range events {
		// if the tx is a contract creation, then extract the contract address
		if isContractCreation && event.Type == types.EventTypeCreate {
			for _, attr := range event.Attributes {
				if attr.Key == types.AttributeKeyContract {
					contractAddr_ := common.HexToAddress(attr.Value)
					contractAddr = &contractAddr_
					break
				}
			}
			continue
		}

		// filter out the events that are not EVM events
		if event.Type != types.EventTypeEVM {
			continue
		}

		// extract the logs
		for _, attr := range event.Attributes {
			if attr.Key != types.AttributeKeyLog {
				continue
			}

			var log coretypes.Log
			val, err := parser.Parse(attr.Value)
			if err != nil {
				return []*coretypes.Log{}, nil, err
			}

			log.Address = common.HexToAddress(string(val.GetStringBytes("address")))
			topics := val.GetArray("topics")
			for _, topic := range topics {
				log.Topics = append(log.Topics, common.HexToHash(string(topic.GetStringBytes())))
			}
			log.Data = hexutil.MustDecode(string(val.GetStringBytes("data")))
			logs = append(logs, &log)
		}
	}

	return logs, contractAddr, nil
}

// CollJsonVal is used for protobuf values of the newest google.golang.org/protobuf API.
func CollJsonVal[T any]() collcodec.ValueCodec[T] {
	return &collJsonVal[T]{}
}

type collJsonVal[T any] struct{}

func (c collJsonVal[T]) Encode(value T) ([]byte, error) {
	return c.EncodeJSON(value)
}

func (c collJsonVal[T]) Decode(b []byte) (T, error) {
	return c.DecodeJSON(b)
}

func (c collJsonVal[T]) EncodeJSON(value T) ([]byte, error) {
	return json.Marshal(value)
}

func (c collJsonVal[T]) DecodeJSON(b []byte) (T, error) {
	var value T

	err := json.Unmarshal(b, &value)
	return value, err
}

func (c collJsonVal[T]) Stringify(value T) string {
	return fmt.Sprintf("%v", value)
}

func (c collJsonVal[T]) ValueType() string {
	return "jsonvalue"
}
