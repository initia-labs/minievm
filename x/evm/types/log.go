package types

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	coretypes "github.com/ethereum/go-ethereum/core/types"
)

type Logs []Log

func NewLogs(ethLogs []*coretypes.Log) Logs {
	logs := make([]Log, len(ethLogs))
	for i, ethLog := range ethLogs {
		logs[i] = NewLog(ethLog)
	}

	return logs
}

func NewLog(ethLog *coretypes.Log) Log {
	topics := make([]string, len(ethLog.Topics))
	for i, topic := range ethLog.Topics {
		topics[i] = topic.Hex()
	}

	return Log{
		Address: ethLog.Address.Hex(),
		Topics:  topics,
		Data:    hexutil.Encode(ethLog.Data),
	}
}

func (l Logs) ToEthLogs() []*coretypes.Log {
	logs := make([]*coretypes.Log, len(l))
	for i, log := range l {
		logs[i] = log.ToEthLog()
	}

	return logs
}

func (l Log) ToEthLog() *coretypes.Log {
	topics := make([]common.Hash, len(l.Topics))
	for i, topic := range l.Topics {
		topics[i] = common.HexToHash(topic)
	}

	return &coretypes.Log{
		Address: common.HexToAddress(l.Address),
		Topics:  topics,
		Data:    hexutil.MustDecode(l.Data),
	}
}
