package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	coretypes "github.com/ethereum/go-ethereum/core/types"
)

// RPCTransaction represents a transaction that will serialize to the RPC representation of a transaction
type RPCTransaction struct {
	BlockHash           *common.Hash          `json:"blockHash"`
	BlockNumber         *hexutil.Big          `json:"blockNumber"`
	From                common.Address        `json:"from"`
	Gas                 hexutil.Uint64        `json:"gas"`
	GasPrice            *hexutil.Big          `json:"gasPrice"`
	GasFeeCap           *hexutil.Big          `json:"maxFeePerGas,omitempty"`
	GasTipCap           *hexutil.Big          `json:"maxPriorityFeePerGas,omitempty"`
	MaxFeePerBlobGas    *hexutil.Big          `json:"maxFeePerBlobGas,omitempty"`
	Hash                common.Hash           `json:"hash"`
	Input               hexutil.Bytes         `json:"input"`
	Nonce               hexutil.Uint64        `json:"nonce"`
	To                  *common.Address       `json:"to"`
	TransactionIndex    *hexutil.Uint64       `json:"transactionIndex"`
	Value               *hexutil.Big          `json:"value"`
	Type                hexutil.Uint64        `json:"type"`
	Accesses            *coretypes.AccessList `json:"accessList,omitempty"`
	ChainID             *hexutil.Big          `json:"chainId,omitempty"`
	BlobVersionedHashes []common.Hash         `json:"blobVersionedHashes,omitempty"`
	V                   *hexutil.Big          `json:"v"`
	R                   *hexutil.Big          `json:"r"`
	S                   *hexutil.Big          `json:"s"`
	YParity             *hexutil.Uint64       `json:"yParity,omitempty"`
}

func NewRPCTransaction(tx *coretypes.Transaction, blockHash common.Hash, blockNumber uint64, index uint64, chainID *big.Int) *RPCTransaction {
	signer := coretypes.LatestSignerForChainID(chainID)
	from, _ := coretypes.Sender(signer, tx)
	v, r, s := tx.RawSignatureValues()
	al := tx.AccessList()
	var yparity *hexutil.Uint64 = new(hexutil.Uint64)
	switch tx.Type() {
	case coretypes.LegacyTxType:
		yparity = nil
	default: // Dynamic and Access List use yParity
		*yparity = hexutil.Uint64(v.Sign())
	}

	result := &RPCTransaction{
		Type:      hexutil.Uint64(tx.Type()),
		From:      from,
		Gas:       hexutil.Uint64(tx.Gas()),
		GasPrice:  (*hexutil.Big)(tx.GasPrice()),
		GasFeeCap: (*hexutil.Big)(tx.GasFeeCap()),
		GasTipCap: (*hexutil.Big)(tx.GasTipCap()),
		Hash:      tx.Hash(),
		Input:     hexutil.Bytes(tx.Data()),
		Nonce:     hexutil.Uint64(tx.Nonce()),
		To:        tx.To(),
		Value:     (*hexutil.Big)(tx.Value()),
		V:         (*hexutil.Big)(v),
		R:         (*hexutil.Big)(r),
		S:         (*hexutil.Big)(s),
		ChainID:   (*hexutil.Big)(chainID),
		Accesses:  &al,
		YParity:   yparity,
	}
	if blockHash != (common.Hash{}) {
		result.BlockHash = &blockHash
		result.BlockNumber = (*hexutil.Big)(new(big.Int).SetUint64(blockNumber))
		result.TransactionIndex = (*hexutil.Uint64)(&index)
	}

	return result
}

func (rpcTx RPCTransaction) ToTransaction() *coretypes.Transaction {
	return coretypes.NewTx(&coretypes.LegacyTx{
		Nonce:    uint64(rpcTx.Nonce),
		GasPrice: rpcTx.GasPrice.ToInt(),
		Gas:      uint64(rpcTx.Gas),
		To:       rpcTx.To,
		Value:    rpcTx.Value.ToInt(),
		Data:     rpcTx.Input,
		V:        rpcTx.V.ToInt(),
		R:        rpcTx.R.ToInt(),
		S:        rpcTx.S.ToInt(),
	})
}
