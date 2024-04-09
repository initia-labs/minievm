package cli

import (
	flag "github.com/spf13/pflag"
)

const (
	FlagTrace          = "trace"
	FlagWithStorage    = "with-storage"
	FlagWithMemory     = "with-memory"
	FlagWithStack      = "with-stack"
	FlagWithReturnData = "with-return-data"
)

func FlagTraceOptions() *flag.FlagSet {
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	fs.Bool(FlagTrace, false, `Trace the execution of the transaction`)
	fs.Bool(FlagWithStorage, false, `Trace the storage of the contract`)
	fs.Bool(FlagWithMemory, false, `Trace the memory of the contract`)
	fs.Bool(FlagWithStack, false, `Trace the stack of the contract`)
	fs.Bool(FlagWithReturnData, false, `Trace the return data of the contract`)
	return fs
}
