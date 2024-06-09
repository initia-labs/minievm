package indexer

import (
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	coretypes "github.com/ethereum/go-ethereum/core/types"

	rpctypes "github.com/initia-labs/minievm/jsonrpc/types"
	"github.com/initia-labs/minievm/x/evm/types"
)

func (e *EVMIndexerImpl) ConvertCosmosTxToEthereumTx(ctx sdk.Context, sdkTx sdk.Tx) (*coretypes.Transaction, error) {
	chainID := ctx.ChainID()
	ac := e.appCodec.InterfaceRegistry().SigningContext().AddressCodec()
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
	params, err := e.evmKeeper.Params.Get(ctx)
	if err != nil {
		return nil, err
	}

	if len(fees) != 1 && fees[0].Denom != params.FeeDenom {
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
		chainId := types.ConvertCosmosChainIDToEthereumChainID(chainID)
		tx = coretypes.NewTx(&coretypes.DynamicFeeTx{
			ChainID:    chainId,
			Nonce:      sig.Sequence,
			Gas:        gas,
			To:         &contractAddr,
			Data:       data,
			GasTipCap:  big.NewInt(0),
			GasFeeCap:  big.NewInt(fee.Amount.Int64()),
			Value:      big.NewInt(0),
			AccessList: coretypes.AccessList{},
		})
	case "/minievm.evm.v1.Create":
		createMsg := msg.(*types.MsgCreate)
		data, err := hexutil.Decode(createMsg.Code)
		if err != nil {
			return nil, err
		}
		chainId := types.ConvertCosmosChainIDToEthereumChainID(chainID)
		tx = coretypes.NewTx(&coretypes.DynamicFeeTx{
			ChainID:    chainId,
			Nonce:      sig.Sequence,
			Gas:        gas,
			To:         nil,
			Data:       data,
			GasTipCap:  big.NewInt(0),
			GasFeeCap:  big.NewInt(fee.Amount.Int64()),
			Value:      big.NewInt(0),
			AccessList: coretypes.AccessList{},
		})
	case "/minievm.evm.v1.Create2":
		// create2 is not supported
		return nil, nil
	}

	return tx, nil
}

func newRPCTransaction(tx *coretypes.Transaction, blockHash common.Hash, blockNumber uint64, index uint64, chainID *big.Int) *rpctypes.RPCTransaction {
	signer := coretypes.LatestSignerForChainID(chainID)
	from, _ := coretypes.Sender(signer, tx)
	v, r, s := tx.RawSignatureValues()
	al := tx.AccessList()
	yparity := hexutil.Uint64(v.Sign())

	result := &rpctypes.RPCTransaction{
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
		YParity:   &yparity,
	}
	if blockHash != (common.Hash{}) {
		result.BlockHash = &blockHash
		result.BlockNumber = (*hexutil.Big)(new(big.Int).SetUint64(blockNumber))
		result.TransactionIndex = (*hexutil.Uint64)(&index)
	}

	return result
}
