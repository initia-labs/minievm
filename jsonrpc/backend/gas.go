package backend

import (
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rpc"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"

	rpctypes "github.com/initia-labs/minievm/jsonrpc/types"
	"github.com/initia-labs/minievm/x/evm/keeper"
	"github.com/initia-labs/minievm/x/evm/types"
)

func (b *JSONRPCBackend) EstimateGas(args rpctypes.TransactionArgs, blockNrOrHash *rpc.BlockNumberOrHash, overrides *rpctypes.StateOverride) (hexutil.Uint64, error) {
	if overrides != nil {
		return hexutil.Uint64(0), errors.New("state overrides are not supported")
	}

	if args.Nonce == nil && args.From != nil {
		nonce, err := b.GetTransactionCount(*args.From, rpc.BlockNumberOrHashWithNumber(rpc.PendingBlockNumber))
		if err != nil {
			return hexutil.Uint64(0), err
		}

		args.Nonce = nonce
	}

	// set call defaults
	args.CallDefaults()

	// convert sender to string
	sender, err := b.app.AccountKeeper.AddressCodec().BytesToString(args.From[:])
	if err != nil {
		return hexutil.Uint64(0), err
	}

	// jsonrpc is not ready for querying
	if b.feeDenom == "" {
		return hexutil.Uint64(0), NewInternalError("jsonrpc is not ready")
	}

	sdkMsgs := []sdk.Msg{}
	if args.To == nil {
		sdkMsgs = append(sdkMsgs, &types.MsgCreate{
			Sender: sender,
			Code:   hexutil.Encode(args.GetData()),
			Value:  math.NewIntFromBigInt(types.FromEthersUnit(b.feeDecimals, args.Value.ToInt())),
		})
	} else {
		sdkMsgs = append(sdkMsgs, &types.MsgCall{
			Sender:       sender,
			ContractAddr: args.To.Hex(),
			Input:        hexutil.Encode(args.GetData()),
			Value:        math.NewIntFromBigInt(types.FromEthersUnit(b.feeDecimals, args.Value.ToInt())),
		})
	}

	txBuilder := b.app.TxConfig().NewTxBuilder()
	if err = txBuilder.SetMsgs(sdkMsgs...); err != nil {
		return hexutil.Uint64(0), err
	}
	if err = txBuilder.SetSignatures(signing.SignatureV2{
		PubKey: nil,
		Data: &signing.SingleSignatureData{
			SignMode:  keeper.SignMode_SIGN_MODE_ETHEREUM,
			Signature: nil,
		},
		Sequence: uint64(*args.Nonce),
	}); err != nil {
		return hexutil.Uint64(0), err
	}
	tx := txBuilder.GetTx()

	txBytes, err := b.app.TxConfig().TxEncoder()(tx)
	if err != nil {
		return hexutil.Uint64(0), err
	}

	gasInfo, _, err := b.app.Simulate(txBytes)
	if err != nil {
		b.logger.Error("failed to simulate tx", "err", err)
		return hexutil.Uint64(0), err
	}

	return hexutil.Uint64(gasInfo.GasUsed), nil
}

func (b *JSONRPCBackend) GasPrice() (*hexutil.Big, error) {
	queryCtx, err := b.getQueryCtx()
	if err != nil {
		return nil, err
	}

	params, err := b.app.OPChildKeeper.GetParams(queryCtx)
	if err != nil {
		return nil, err
	}

	// jsonrpc is not ready for querying
	if b.feeDenom == "" {
		return nil, NewInternalError("jsonrpc is not ready")
	}

	// multiply by 1e9 to prevent decimal drops
	gasPrice := params.MinGasPrices.AmountOf(b.feeDenom).
		MulTruncate(math.LegacyNewDec(1e9)).
		TruncateInt().BigInt()

	return (*hexutil.Big)(types.ToEthersUint(b.feeDecimals+9, gasPrice)), nil
}

func (b *JSONRPCBackend) MaxPriorityFeePerGas() (*hexutil.Big, error) {
	return (*hexutil.Big)(big.NewInt(0)), nil
}
