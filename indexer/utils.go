package indexer

import (
	"encoding/json"
	"fmt"

	collcodec "cosmossdk.io/collections/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/gogoproto/proto"

	"github.com/ethereum/go-ethereum/common"
	coretypes "github.com/ethereum/go-ethereum/core/types"

	"github.com/initia-labs/minievm/x/evm/types"
)

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
