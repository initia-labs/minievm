package types

import (
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
		Address: ethLog.Address.String(),
		Topics:  topics,
		Data:    ethLog.Data,
	}
}
