package indexer

import (
	"encoding/json"
	"fmt"
	"math/big"

	abci "github.com/cometbft/cometbft/abci/types"
	comettypes "github.com/cometbft/cometbft/types"

	collcodec "cosmossdk.io/collections/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/gogoproto/proto"

	"github.com/ethereum/go-ethereum/common"
	coretypes "github.com/ethereum/go-ethereum/core/types"

	"github.com/initia-labs/minievm/x/evm/keeper"
	"github.com/initia-labs/minievm/x/evm/types"
)

type EthTxInfo struct {
	Tx           *coretypes.Transaction
	Logs         []*coretypes.Log
	ContractAddr *common.Address
	Status       uint64
	GasUsed      uint64
	CosmosTxHash []byte
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
	tx, err := txDecoder(txBytes)
	if err != nil {
		return nil, err
	}

	txStatus := coretypes.ReceiptStatusSuccessful
	if txResult.Code != abci.CodeTypeOK {
		txStatus = coretypes.ReceiptStatusFailed
	}

	// convert cosmos tx to ethereum tx
	ethTx, _, err := evmKeeper.TxUtils().ConvertCosmosTxToEthereumTx(ctx, tx)
	if err != nil {
		return nil, err
	}

	// if the tx is normal ethereum tx, then return a single EthTxInfo
	if ethTx != nil {
		ethLogs, contractAddr, err := extractLogsAndContractAddr(txStatus, txResult.Data, ethTx.To() == nil)
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

	var logs types.Logs
	for _, event := range txResult.Events {
		if event.Type != types.EventTypeEVM {
			continue
		}

		for _, attr := range event.Attributes {
			if attr.Key != types.AttributeKeyLog {
				continue
			}

			var log types.Log
			if err := json.Unmarshal([]byte(attr.Value), &log); err != nil {
				return nil, err
			}

			logs = append(logs, log)
		}
	}
	if len(logs) == 0 {
		return nil, nil
	}

	// if there is evm events, then create a fake eth tx
	cosmosTxHash := comettypes.Tx(txBytes).Hash()
	ethLogs := logs.ToEthLogs()
	ethTx = coretypes.NewTx(&coretypes.LegacyTx{
		Data:     cosmosTxHash,
		To:       nil,
		Nonce:    0,
		Gas:      0,
		GasPrice: new(big.Int),
		Value:    new(big.Int),
		V:        new(big.Int),
		R:        new(big.Int),
		S:        new(big.Int),
	})

	// check whether the tx has MsgCall or MsgCreate or MsgCreate2
	return &EthTxInfo{
		Tx:           ethTx,
		Logs:         ethLogs,
		Status:       txStatus,
		ContractAddr: nil,
		GasUsed:      uint64(txResult.GasUsed),
		CosmosTxHash: cosmosTxHash,
	}, nil
}

func extractLogsAndContractAddr(txStatus uint64, data []byte, isContractCreation bool) ([]*coretypes.Log, *common.Address, error) {
	if txStatus != coretypes.ReceiptStatusSuccessful {
		return []*coretypes.Log{}, nil, nil
	}

	var ethLogs []*coretypes.Log
	var contractAddr *common.Address

	if isContractCreation {
		var resp types.MsgCreateResponse
		if err := unpackData(data, &resp); err != nil {
			return nil, nil, err
		}

		ethLogs = types.Logs(resp.Logs).ToEthLogs()
		contractAddr_ := common.HexToAddress(resp.ContractAddr)
		contractAddr = &contractAddr_
	} else {
		var resp types.MsgCallResponse
		if err := unpackData(data, &resp); err != nil {
			return nil, nil, err
		}

		ethLogs = types.Logs(resp.Logs).ToEthLogs()
	}

	return ethLogs, contractAddr, nil
}

// unpackData extracts msg response from the data
func unpackData(data []byte, resp proto.Message) error {
	var txMsgData sdk.TxMsgData
	if err := proto.Unmarshal(data, &txMsgData); err != nil {
		return err
	}

	if len(txMsgData.MsgResponses) == 0 {
		return sdkerrors.ErrLogic.Wrap("failed to unpack data; got nil Msg response")
	}

	msgResp := txMsgData.MsgResponses[0]
	expectedTypeUrl := sdk.MsgTypeURL(resp)
	if msgResp.TypeUrl != expectedTypeUrl {
		return fmt.Errorf("unexpected type URL; got: %s, expected: %s", msgResp.TypeUrl, expectedTypeUrl)
	}

	// Unpack the response
	if err := proto.Unmarshal(msgResp.Value, resp); err != nil {
		return err
	}

	return nil
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
