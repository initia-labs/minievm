package backend

// TxIndexingError is an API error that indicates the transaction indexing is not
// fully finished yet with JSON error code and a binary data blob.
type TxIndexingError struct{}

// NewTxIndexingError creates a TxIndexingError instance.
func NewTxIndexingError() *TxIndexingError { return &TxIndexingError{} }

// Error implement error interface, returning the error message.
func (e *TxIndexingError) Error() string {
	return "transaction indexing is in progress"
}

// ErrorCode returns the JSON error code for a revert.
// See: https://github.com/ethereum/wiki/wiki/JSON-RPC-Error-Codes-Improvement-Proposal
func (e *TxIndexingError) ErrorCode() int {
	return -32000 // to be decided
}

// ErrorData returns the hex encoded revert reason.
func (e *TxIndexingError) ErrorData() interface{} { return "transaction indexing is in progress" }

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
