package backend

import (
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	coretypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"

	rpctypes "github.com/initia-labs/minievm/jsonrpc/types"
	"github.com/initia-labs/minievm/x/evm/types"
)

func (b *JSONRPCBackend) ConvertCosmosTxToEthereumTx(sdkTx sdk.Tx) (*coretypes.Transaction, error) {
	chainID := b.clientCtx.ChainID
	ac := b.clientCtx.Codec.InterfaceRegistry().SigningContext().AddressCodec()
	msgs := sdkTx.GetMsgs()
	if len(msgs) != 1 {
		return nil, nil
	}

	authTx := sdkTx.(authsigning.Tx)
	sigs, err := authTx.GetSignaturesV2()
	if err != nil {
		return nil, err
	}
	if len(sigs) != 1 {
		return nil, nil
	}

	fees := authTx.GetFee()
	if len(fees) != 1 {
		return nil, nil
	}

	var tx *coretypes.Transaction

	sig := sigs[0]
	msg := msgs[0]
	fee := fees[0]
	gas := authTx.GetGas()
	typeUrl := sdk.MsgTypeURL(msg)

	switch typeUrl {
	case "/minievm.evm.v1.Call":
		callMsg := msg.(*types.MsgCall)
		contractAddr, err := types.ContractAddressFromString(ac, callMsg.ContractAddr)
		if err != nil {
			return nil, err
		}

		data, err := hexutil.Decode(callMsg.Input)
		if err != nil {
			return nil, err
		}

		tx = coretypes.NewTx(&coretypes.DynamicFeeTx{
			ChainID:    types.ConvertCosmosChainIDToEthereumChainID(chainID),
			Nonce:      sig.Sequence,
			Gas:        gas,
			To:         &contractAddr,
			Data:       data,
			GasTipCap:  big.NewInt(0),
			GasFeeCap:  big.NewInt(fee.Amount.Int64()),
			Value:      big.NewInt(0),
			AccessList: coretypes.AccessList{},
		})
		break
	case "/minievm.evm.v1.Create":
		createMsg := msg.(*types.MsgCreate)
		data, err := hexutil.Decode(createMsg.Code)
		if err != nil {
			return nil, err
		}

		tx = coretypes.NewTx(&coretypes.DynamicFeeTx{
			ChainID:    types.ConvertCosmosChainIDToEthereumChainID(chainID),
			Nonce:      sig.Sequence,
			Gas:        gas,
			To:         nil,
			Data:       data,
			GasTipCap:  big.NewInt(0),
			GasFeeCap:  big.NewInt(fee.Amount.Int64()),
			Value:      big.NewInt(0),
			AccessList: coretypes.AccessList{},
		})
		break
	case "/minievm.evm.v1.Create2":
		// create2 is not supported
		return nil, nil
	}

	return tx, nil
}

// newRPCTransaction returns a transaction that will serialize to the RPC
// representation, with the given location metadata set (if available).
func newRPCTransaction(tx *coretypes.Transaction, blockHash common.Hash, blockNumber uint64, blockTime uint64, index uint64, config *params.ChainConfig) *rpctypes.RPCTransaction {
	signer := coretypes.MakeSigner(config, new(big.Int).SetUint64(blockNumber), blockTime)
	from, _ := coretypes.Sender(signer, tx)
	v, r, s := tx.RawSignatureValues()
	result := &rpctypes.RPCTransaction{
		Type:     hexutil.Uint64(tx.Type()),
		From:     from,
		Gas:      hexutil.Uint64(tx.Gas()),
		GasPrice: (*hexutil.Big)(tx.GasPrice()),
		Hash:     tx.Hash(),
		Input:    hexutil.Bytes(tx.Data()),
		Nonce:    hexutil.Uint64(tx.Nonce()),
		To:       tx.To(),
		Value:    (*hexutil.Big)(tx.Value()),
		V:        (*hexutil.Big)(v),
		R:        (*hexutil.Big)(r),
		S:        (*hexutil.Big)(s),
	}
	if blockHash != (common.Hash{}) {
		result.BlockHash = &blockHash
		result.BlockNumber = (*hexutil.Big)(new(big.Int).SetUint64(blockNumber))
		result.TransactionIndex = (*hexutil.Uint64)(&index)
	}

	// only support dynamic fee tx
	al := tx.AccessList()
	yparity := hexutil.Uint64(v.Sign())
	result.Accesses = &al
	result.ChainID = (*hexutil.Big)(tx.ChainId())
	result.YParity = &yparity
	result.GasFeeCap = (*hexutil.Big)(tx.GasFeeCap())
	result.GasTipCap = (*hexutil.Big)(tx.GasTipCap())
	result.GasPrice = (*hexutil.Big)(tx.GasFeeCap())

	return result
}
