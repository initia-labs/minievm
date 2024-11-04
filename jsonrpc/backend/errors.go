package backend

import "errors"

type InternalError struct {
	msg string
}

func NewInternalError(msg string) *InternalError {
	return &InternalError{msg: msg}
}

func (e *InternalError) Error() string {
	return "internal jsonrpc error: " + e.msg
}

func (e *InternalError) ErrorCode() int {
	// Internal JSON-RPC error
	return -32603
}

var (
	errInvalidPercentile = errors.New("invalid reward percentile")
	errRequestBeyondHead = errors.New("request beyond head block")
)
