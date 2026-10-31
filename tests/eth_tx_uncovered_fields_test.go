package tests

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"

	abcitypes "github.com/cometbft/cometbft/abci/types"

	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// The ethereum signature of a converted eth tx only covers the fields used to
// rebuild the eth tx. Changing any other Cosmos tx field (e.g. lowering the gas
// limit so the tx runs out of gas while the sender still pays the full fee)
// must make the tx invalid instead of executing it.
func Test_EthTx_RejectsChangedUncoveredFields(t *testing.T) {
	app, addrs, privKeys := CreateApp(t)
	victim := &ecdsaKey{k: privKeys[0], addr: addrs[0]}
	contract := deployLoop(t, app, victim)

	// victim signs gas 1,000,000 for ~800k gas of work
	input := loopInput(0, 0, 0, 20_000)
	tx, _ := GenerateTx(t, app, victim.k, &contract, input, new(big.Int), SetGasLimit(1_000_000))
	txBz, err := app.TxConfig().TxEncoder()(tx)
	require.NoError(t, err)

	// WrapTxBuilder edits the tx in place, so every case starts from a fresh decode
	freshTx := func() sdk.Tx {
		decoded, err := app.TxConfig().TxDecoder()(txBz)
		require.NoError(t, err)
		return decoded
	}

	state := func() (sdk.Coin, uint64) {
		ctx, closer, err := app.CreateQueryContext(0, false)
		require.NoError(t, err)
		if closer != nil {
			defer closer.Close()
		}
		feeDenom, err := app.EVMKeeper.GetFeeDenom(ctx)
		require.NoError(t, err)
		seq, err := app.AccountKeeper.GetSequence(ctx, sdk.AccAddress(victim.addr.Bytes()))
		require.NoError(t, err)
		return app.BankKeeper.GetBalance(ctx, sdk.AccAddress(victim.addr.Bytes()), feeDenom), seq
	}

	mutations := []struct {
		name   string
		mutate func(client.TxBuilder)
	}{
		{"lower gas limit", func(b client.TxBuilder) { b.SetGasLimit(120_000) }},
		{"higher gas limit", func(b client.TxBuilder) { b.SetGasLimit(2_000_000) }},
		{"fee granter", func(b client.TxBuilder) { b.SetFeeGranter(sdk.AccAddress(addrs[1].Bytes())) }},
		{"timeout height", func(b client.TxBuilder) { b.SetTimeoutHeight(1_000_000) }},
	}
	for _, m := range mutations {
		t.Run(m.name, func(t *testing.T) {
			builder, err := app.TxConfig().WrapTxBuilder(freshTx())
			require.NoError(t, err)
			m.mutate(builder)
			mutated := builder.GetTx()

			// rejected by CheckTx
			bz, err := app.TxConfig().TxEncoder()(mutated)
			require.NoError(t, err)
			res, err := app.CheckTx(&abcitypes.RequestCheckTx{Tx: bz, Type: abcitypes.CheckTxType_New})
			require.NoError(t, err)
			require.NotEqual(t, abcitypes.CodeTypeOK, res.Code)

			// rejected in a block without touching the sender
			balBefore, seqBefore := state()
			_, fin := ExecuteTxs(t, app, mutated)
			require.NotEqual(t, abcitypes.CodeTypeOK, fin.TxResults[0].Code)
			require.Contains(t, fin.TxResults[0].Log, "Tx conversion failed")
			balAfter, seqAfter := state()
			require.Equal(t, balBefore, balAfter)
			require.Equal(t, seqBefore, seqAfter)
		})
	}

	// the untouched tx still executes
	_, fin := ExecuteTxs(t, app, freshTx())
	CheckTxResult(t, fin.TxResults[0], true)
}
