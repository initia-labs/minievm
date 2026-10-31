package tests

import (
	"testing"

	"github.com/stretchr/testify/require"

	abcitypes "github.com/cometbft/cometbft/abci/types"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	authsign "github.com/cosmos/cosmos-sdk/x/auth/signing"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	"github.com/ethereum/go-ethereum/crypto"

	"github.com/initia-labs/initia/crypto/ethsecp256k1"

	minitiaapp "github.com/initia-labs/minievm/app"
	evmtypes "github.com/initia-labs/minievm/x/evm/types"
)

// An `evm/<address>` fee coin makes fee deduction call that contract, before
// signatures are checked. Outside block execution such fee denoms must be
// rejected before any contract code runs.
func Test_CheckTx_RejectsUnlistedERC20FeeDenom(t *testing.T) {
	app, addrs, privKeys := CreateApp(t)
	key := &ecdsaKey{k: privKeys[0], addr: addrs[0]}

	// every call into this contract loops until it runs out of gas
	spin := deployRuntime(t, app, key, []byte{0x5b, 0x60, 0x00, 0x56}) // JUMPDEST PUSH1 0 JUMP
	spinDenom := evmtypes.DENOM_PREFIX + spin.Hex()[2:]

	ctx, closer, err := app.CreateQueryContext(0, false)
	require.NoError(t, err)
	feeDenom, err := app.EVMKeeper.GetFeeDenom(ctx)
	require.NoError(t, err)
	minGasPrices, err := app.OPChildKeeper.MinGasPrices(ctx)
	require.NoError(t, err)
	if closer != nil {
		closer.Close()
	}

	const gas = uint64(10_000_000)
	requiredFee := sdk.NewCoin(feeDenom, minGasPrices.AmountOf(feeDenom).MulInt64(int64(gas)).Ceil().TruncateInt())

	payer := sdk.AccAddress(key.addr.Bytes())
	state := func() (sdk.Coin, uint64) {
		ctx, closer, err := app.CreateQueryContext(0, false)
		require.NoError(t, err)
		if closer != nil {
			defer closer.Close()
		}
		seq, err := app.AccountKeeper.GetSequence(ctx, payer)
		require.NoError(t, err)
		return app.BankKeeper.GetBalance(ctx, payer, feeDenom), seq
	}

	// CheckTx rejects the tx before deducting fees
	txBz := buildSignedFeeTx(t, app, key, gas, sdk.NewCoins(sdk.NewCoin(spinDenom, math.NewInt(1)), requiredFee))
	balBefore, seqBefore := state()
	res, err := app.CheckTx(&abcitypes.RequestCheckTx{Tx: txBz, Type: abcitypes.CheckTxType_New})
	require.NoError(t, err)
	require.NotEqual(t, abcitypes.CodeTypeOK, res.Code)
	require.Contains(t, res.Log, "is not allowed")
	require.Less(t, res.GasUsed, int64(1_000_000), "contract code ran during fee deduction")
	balAfter, seqAfter := state()
	require.Equal(t, balBefore, balAfter)
	require.Equal(t, seqBefore, seqAfter)

	// simulation rejects it too
	_, _, err = app.Simulate(txBz)
	require.ErrorContains(t, err, "is not allowed")

	// paying in the fee denom alone is still accepted
	okBz := buildSignedFeeTx(t, app, key, gas, sdk.NewCoins(requiredFee))
	res, err = app.CheckTx(&abcitypes.RequestCheckTx{Tx: okBz, Type: abcitypes.CheckTxType_New})
	require.NoError(t, err)
	require.Equal(t, abcitypes.CodeTypeOK, res.Code, res.Log)
}

// buildSignedFeeTx builds a signed bank send-to-self with the given gas and fee.
func buildSignedFeeTx(t *testing.T, app *minitiaapp.MinitiaApp, key *ecdsaKey, gas uint64, fee sdk.Coins) []byte {
	txConfig := app.TxConfig()
	b := txConfig.NewTxBuilder()
	payer := sdk.AccAddress(key.addr.Bytes())
	qctx, qcloser, err := app.CreateQueryContext(0, false)
	require.NoError(t, err)
	if qcloser != nil {
		defer qcloser.Close()
	}
	feeDenom, err := app.EVMKeeper.GetFeeDenom(qctx)
	require.NoError(t, err)
	require.NoError(t, b.SetMsgs(banktypes.NewMsgSend(payer, payer, sdk.NewCoins(sdk.NewInt64Coin(feeDenom, 1)))))
	b.SetGasLimit(gas)
	b.SetFeeAmount(fee)

	signMode, err := authsign.APISignModeToInternal(txConfig.SignModeHandler().DefaultMode())
	require.NoError(t, err)
	priv := ethsecp256k1.PrivKey{Key: crypto.FromECDSA(key.k)}
	pub := priv.PubKey()

	ctx, closer, err := app.CreateQueryContext(0, false)
	require.NoError(t, err)
	if closer != nil {
		defer closer.Close()
	}
	acc := app.AccountKeeper.GetAccount(ctx, payer)
	require.NotNil(t, acc)

	sig := signing.SignatureV2{PubKey: pub, Data: &signing.SingleSignatureData{SignMode: signMode}, Sequence: acc.GetSequence()}
	require.NoError(t, b.SetSignatures(sig))
	signBytes, err := authsign.GetSignBytesAdapter(ctx, txConfig.SignModeHandler(), signMode, authsign.SignerData{
		Address: payer.String(), ChainID: ctx.ChainID(), AccountNumber: acc.GetAccountNumber(), Sequence: acc.GetSequence(), PubKey: pub,
	}, b.GetTx())
	require.NoError(t, err)
	sigBz, err := priv.Sign(signBytes)
	require.NoError(t, err)
	sig.Data.(*signing.SingleSignatureData).Signature = sigBz
	require.NoError(t, b.SetSignatures(sig))

	bz, err := txConfig.TxEncoder()(b.GetTx())
	require.NoError(t, err)
	return bz
}
