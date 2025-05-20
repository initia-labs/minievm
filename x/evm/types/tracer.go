package types

import (
	"github.com/ethereum/go-ethereum/core/tracing"
)

// TracingHooks is a function that returns a tracing.Hooks
// It is used to generate a new tracing.Hooks for each transaction
// and to stop the tracing.Hooks after the transaction is executed
type TracingHooks func() *tracing.Hooks
