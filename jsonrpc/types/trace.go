package types

import "github.com/ethereum/go-ethereum/common"

// TxTraceResult is the result of a single transaction trace.
type TxTraceResult struct {
	TxHash common.Hash `json:"txHash"`           // transaction hash
	Result any         `json:"result,omitempty"` // Trace results produced by the tracer
	Error  string      `json:"error,omitempty"`  // Trace failure produced by the tracer
}

type StorageRangeResult struct {
	Storage StorageMap   `json:"storage"`
	NextKey *common.Hash `json:"nextKey"` // nil if Storage includes the last key
}

type StorageMap map[common.Hash]StorageEntry

type StorageEntry struct {
	Key   *common.Hash `json:"key"`
	Value common.Hash  `json:"value"`
}
