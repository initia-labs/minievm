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

// StaticGasForCosmos is some buffer for gas used in cosmos part of the tx
const StaticGasForCosmos = 50_000

func (b *JSONRPCBackend) EstimateGas(args rpctypes.TransactionArgs, blockNrOrHash *rpc.BlockNumberOrHash, overrides *rpctypes.StateOverride) (hexutil.Uint64, error) {
	if overrides != nil {
		return hexutil.Uint64(0), errors.New("state overrides are not supported")
	}
	if args.From == nil {
		return hexutil.Uint64(0), errors.New("from address cannot be empty in the estimated gas")
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

	_, feeDecimals, err := b.feeInfo()
	if err != nil {
		return hexutil.Uint64(0), err
	}

	var list []types.AccessTuple
	if args.AccessList != nil {
		list = types.ConvertEthAccessListToCosmos(*args.AccessList)
	}

	sdkMsgs := []sdk.Msg{}
	if args.To == nil {
		sdkMsgs = append(sdkMsgs, &types.MsgCreate{
			Sender:     sender,
			Code:       hexutil.Encode(args.GetData()),
			Value:      math.NewIntFromBigInt(types.FromEthersUnit(feeDecimals, args.Value.ToInt())),
			AccessList: list,
		})
	} else {
		sdkMsgs = append(sdkMsgs, &types.MsgCall{
			Sender:       sender,
			ContractAddr: args.To.Hex(),
			Input:        hexutil.Encode(args.GetData()),
			Value:        math.NewIntFromBigInt(types.FromEthersUnit(feeDecimals, args.Value.ToInt())),
			AccessList:   list,
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

	// apply gas multiplier
	gasUsed := b.gasMultiplier.MulInt(math.NewIntFromUint64(gasInfo.GasUsed)).TruncateInt().Uint64() + StaticGasForCosmos
	return hexutil.Uint64(gasUsed), nil
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

	feeDenom, feeDecimals, err := b.feeInfo()
	if err != nil {
		return nil, err
	}

	// Multiply by 1e9 to maintain precision during conversion
	// This adds 9 decimal places to prevent truncation errors
	const precisionMultiplier = 1e9

	// multiply by 1e9 to prevent decimal drops
	gasPrice := params.MinGasPrices.AmountOf(feeDenom).
		MulTruncate(math.LegacyNewDec(precisionMultiplier)).
		TruncateInt().BigInt()

	// Verify the result is within safe bounds
	if gasPrice.BitLen() > 256 {
		return nil, NewInternalError("gas price overflow")
	}

	return (*hexutil.Big)(types.ToEthersUnit(feeDecimals+9, gasPrice)), nil
}

func (b *JSONRPCBackend) MaxPriorityFeePerGas() (*hexutil.Big, error) {
	return (*hexutil.Big)(big.NewInt(0)), nil
}
