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

// TimeoutError indicates that the request has timed out.
type TimeoutError struct {
	msg string
}

func NewTimeoutError(msg string) *TimeoutError {
	return &TimeoutError{msg: msg}
}

func (e *TimeoutError) Error() string {
	return "jsonrpc timeout error: " + e.msg
}

func (e *TimeoutError) ErrorCode() int {
	return 4
}

// ReadinessError indicates that the backend is not ready to process requests.
type ReadinessError struct {
	msg string
}

func NewReadinessError(msg string) *ReadinessError {
	return &ReadinessError{msg: msg}
}

func (e *ReadinessError) Error() string {
	return "jsonrpc readiness error: " + e.msg
}

func (e *ReadinessError) ErrorCode() int {
	return 5
}
