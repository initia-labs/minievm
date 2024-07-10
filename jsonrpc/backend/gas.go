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

	if args.Nonce == nil {
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

	_, decimals, err := b.feeDenomWithDecimals()
	if err != nil {
		return hexutil.Uint64(0), err
	}

	sdkMsgs := []sdk.Msg{}
	if args.To == nil {
		sdkMsgs = append(sdkMsgs, &types.MsgCreate{
			Sender: sender,
			Code:   hexutil.Encode(args.GetData()),
			Value:  math.NewIntFromBigInt(types.FromEthersUnit(decimals, args.Value.ToInt())),
		})
	} else {
		sdkMsgs = append(sdkMsgs, &types.MsgCall{
			Sender:       sender,
			ContractAddr: args.To.Hex(),
			Input:        hexutil.Encode(args.GetData()),
			Value:        math.NewIntFromBigInt(types.FromEthersUnit(decimals, args.Value.ToInt())),
		})
	}

	txBuilder := b.app.TxConfig().NewTxBuilder()
	txBuilder.SetMsgs(sdkMsgs...)
	txBuilder.SetSignatures(signing.SignatureV2{
		PubKey: nil,
		Data: &signing.SingleSignatureData{
			SignMode:  keeper.SignMode_SIGN_MODE_ETHEREUM,
			Signature: nil,
		},
		Sequence: uint64(*args.Nonce),
	})
	tx := txBuilder.GetTx()
	txBytes, err := b.app.TxConfig().TxEncoder()(tx)
	if err != nil {
		return hexutil.Uint64(0), err
	}

	gasInfo, _, err := b.app.Simulate(txBytes)
	if err != nil {
		b.svrCtx.Logger.Error("failed to simulate tx", "err", err)
		return hexutil.Uint64(0), err
	}

	return hexutil.Uint64(gasInfo.GasUsed), nil
}

func (b *JSONRPCBackend) feeDenom() (string, error) {
	queryCtx, err := b.getQueryCtx()
	if err != nil {
		return "", err
	}

	params, err := b.app.EVMKeeper.Params.Get(queryCtx)
	if err != nil {
		return "", err
	}

	return params.FeeDenom, nil
}

func (b *JSONRPCBackend) feeDenomWithDecimals() (string, uint8, error) {
	feeDenom, err := b.feeDenom()
	if err != nil {
		return "", 0, err
	}

	queryCtx, err := b.getQueryCtx()
	if err != nil {
		return "", 0, err
	}

	decimals, err := b.app.EVMKeeper.ERC20Keeper().GetDecimals(queryCtx, feeDenom)
	if err != nil {
		return "", 0, err
	}

	return feeDenom, decimals, nil
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

	feeDenom, err := b.feeDenom()
	if err != nil {
		return nil, err
	}

	decimals, err := b.app.EVMKeeper.ERC20Keeper().GetDecimals(queryCtx, feeDenom)
	if err != nil {
		return nil, err
	}

	// multiply by 1e9 to prevent decimal drops
	gasPrice := params.MinGasPrices.AmountOf(feeDenom).
		MulTruncate(math.LegacyNewDec(1e9)).
		TruncateInt().BigInt()

	return (*hexutil.Big)(types.ToEthersUint(decimals+9, gasPrice)), nil
}

func (b *JSONRPCBackend) MaxPriorityFeePerGas() (*hexutil.Big, error) {
	return (*hexutil.Big)(big.NewInt(0)), nil
}
