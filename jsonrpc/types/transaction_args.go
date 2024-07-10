package types

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// TransactionArgs represents the arguments to construct a new transaction
// or a message call.
type TransactionArgs struct {
	From  *common.Address `json:"from"`
	To    *common.Address `json:"to"`
	Value *hexutil.Big    `json:"value"`
	Nonce *hexutil.Uint64 `json:"nonce"`

	// We accept "data" and "input" for backwards-compatibility reasons.
	// "input" is the newer name and should be preferred by clients.
	// Issue detail: https://github.com/ethereum/go-ethereum/issues/15628
	Data  *hexutil.Bytes `json:"data"`
	Input *hexutil.Bytes `json:"input"`
}

// from retrieves the transaction sender address.
func (args *TransactionArgs) GetFrom() common.Address {
	if args.From == nil {
		return common.Address{}
	}
	return *args.From
}

// data retrieves the transaction calldata. Input field is preferred.
func (args *TransactionArgs) GetData() []byte {
	if args.Input != nil {
		return *args.Input
	}
	if args.Data != nil {
		return *args.Data
	}
	return nil
}

// CallDefaults sanitizes the transaction arguments, often filling in zero values,
// for the purpose of eth_call class of RPC methods.
func (args *TransactionArgs) CallDefaults() {
	if args.Nonce == nil {
		args.Nonce = new(hexutil.Uint64)
	}
	if args.Value == nil {
		args.Value = new(hexutil.Big)
	}
}

// String returns the struct in a string format.
func (args *TransactionArgs) String() string {
	return fmt.Sprintf(
		"TransactionArgs{From:%v, To:%v, Value:%v, Nonce:%v, Data:%v, Input:%v}",
		args.From,
		args.To,
		args.Value,
		args.Nonce,
		args.Data,
		args.Input,
	)
}
