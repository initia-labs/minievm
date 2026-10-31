package tests

import (
	"math/big"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	abci "github.com/cometbft/cometbft/abci/types"
	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"

	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ethereum/go-ethereum/common"

	minitiaapp "github.com/initia-labs/minievm/app"
	"github.com/initia-labs/minievm/x/evm/contracts/erc20"
	"github.com/initia-labs/minievm/x/evm/contracts/erc20_factory"
	evmtypes "github.com/initia-labs/minievm/x/evm/types"
)

type feePipeline struct {
	app *minitiaapp.MinitiaApp
	now time.Time
}

func (p *feePipeline) block(t *testing.T, txs [][]byte) *abci.ResponseFinalizeBlock {
	p.now = p.now.Add(2 * time.Second)
	h := p.app.LastBlockHeight() + 1
	proposer := make([]byte, 20)
	prep, err := p.app.PrepareProposal(&abci.RequestPrepareProposal{Height: h, Time: p.now, MaxTxBytes: 1 << 22, ProposerAddress: proposer, Txs: txs})
	require.NoError(t, err)
	proc, err := p.app.ProcessProposal(&abci.RequestProcessProposal{Height: h, Time: p.now, ProposerAddress: proposer, Txs: prep.Txs})
	require.NoError(t, err)
	require.Equal(t, abci.ResponseProcessProposal_ACCEPT, proc.Status)
	fin, err := p.app.FinalizeBlock(&abci.RequestFinalizeBlock{Height: h, Time: p.now, ProposerAddress: proposer, Txs: prep.Txs})
	require.NoError(t, err)
	_, err = p.app.Commit()
	require.NoError(t, err)
	p.app.EVMIndexer().Wait()
	require.Equal(t, len(prep.Txs), len(fin.TxResults))
	return fin
}

// land takes a tx the way a node does: CheckTx, simulation, proposal, block.
func (p *feePipeline) land(t *testing.T, txBz []byte) {
	res, err := p.app.CheckTx(&abci.RequestCheckTx{Tx: txBz, Type: abci.CheckTxType_New})
	require.NoError(t, err)
	require.Equal(t, abci.CodeTypeOK, res.Code, res.Log)
	_, _, err = p.app.Simulate(txBz)
	require.NoError(t, err)
	fin := p.block(t, [][]byte{txBz})
	require.Len(t, fin.TxResults, 1, "tx passed CheckTx but was left out of the block")
	require.Equal(t, abci.CodeTypeOK, fin.TxResults[0].Code, fin.TxResults[0].Log)
}

func Test_FeeDenoms_ReachBlock(t *testing.T) {
	app, addrs, privKeys := CreateApp(t)
	key := &ecdsaKey{k: privKeys[0], addr: addrs[0]}
	payer := sdk.AccAddress(key.addr.Bytes())
	p := &feePipeline{app: app, now: time.Now().UTC()}

	ctx := app.NewUncachedContext(false, cmtproto.Header{Height: app.LastBlockHeight(), ChainID: app.ChainID(), Time: p.now})
	factoryABI, err := erc20_factory.Erc20FactoryMetaData.GetAbi()
	require.NoError(t, err)
	erc20ABI, err := erc20.Erc20MetaData.GetAbi()
	require.NoError(t, err)
	factory, err := app.EVMKeeper.GetERC20FactoryAddr(ctx)
	require.NoError(t, err)
	newToken := func(symbol string) string {
		input, err := factoryABI.Pack("createERC20", symbol, symbol, uint8(18))
		require.NoError(t, err)
		ret, _, err := app.EVMKeeper.EVMCall(ctx, key.addr, factory, input, nil, nil, nil)
		require.NoError(t, err)
		token := common.BytesToAddress(ret[12:])
		amount, _ := new(big.Int).SetString("1000000000000000000000000000", 10)
		input, err = erc20ABI.Pack("mint", key.addr, amount)
		require.NoError(t, err)
		_, _, err = app.EVMKeeper.EVMCall(ctx, key.addr, token, input, nil, nil, nil)
		require.NoError(t, err)
		return evmtypes.DENOM_PREFIX + token.Hex()[2:]
	}
	feeDenom, chainDenom, localDenom := newToken("AAA"), newToken("BBB"), newToken("CCC")

	evmParams, err := app.EVMKeeper.Params.Get(ctx)
	require.NoError(t, err)
	evmParams.FeeDenom = feeDenom
	require.NoError(t, app.EVMKeeper.Params.Set(ctx, evmParams))
	opParams, err := app.OPChildKeeper.GetParams(ctx)
	require.NoError(t, err)
	opParams.MinGasPrices = sdk.NewDecCoinsFromCoins(sdk.NewInt64Coin(feeDenom, 20), sdk.NewInt64Coin(chainDenom, 15)).Sort()
	require.NoError(t, app.OPChildKeeper.SetParams(ctx, opParams))
	baseapp.SetMinGasPrices("10" + localDenom)(app.BaseApp)
	p.block(t, nil)
	p.block(t, nil)

	const gas = uint64(300_000)
	price := func(denom string, perGas int64) sdk.Coin { return sdk.NewCoin(denom, math.NewInt(perGas*int64(gas))) }
	balance := func(denom string) math.Int {
		q, closer, err := app.CreateQueryContext(0, false)
		require.NoError(t, err)
		if closer != nil {
			defer closer.Close()
		}
		return app.BankKeeper.GetBalance(q, payer, denom).Amount
	}
	paid := func(denom string, want sdk.Coins) {
		before := balance(denom)
		p.land(t, buildSignedFeeTx(t, app, key, gas, want))
		require.Equal(t, want.AmountOf(denom).String(), before.Sub(balance(denom)).String())
	}
	paid(feeDenom, sdk.NewCoins(price(feeDenom, 20)))
	paid(chainDenom, sdk.NewCoins(price(chainDenom, 15)))
	paid(localDenom, sdk.NewCoins(price(localDenom, 10)))
	paid(feeDenom, sdk.NewCoins(price(feeDenom, 20), price(chainDenom, 15)))

	// an eth tx, the path every JSON-RPC wallet takes
	contract := deployLoop(t, app, key)
	ethTx, _ := GenerateTx(t, app, key.k, &contract, loopInput(0, 0, 0, 10), new(big.Int), SetGasLimit(300_000))
	ethBz, err := app.TxConfig().TxEncoder()(ethTx)
	require.NoError(t, err)
	p.land(t, ethBz)

	// recheck keeps a valid tx
	valid := buildSignedFeeTx(t, app, key, gas, sdk.NewCoins(price(feeDenom, 20)))
	res, err := app.CheckTx(&abci.RequestCheckTx{Tx: valid, Type: abci.CheckTxType_New})
	require.NoError(t, err)
	require.Equal(t, abci.CodeTypeOK, res.Code, res.Log)
	p.now = p.now.Add(2 * time.Second)
	_, err = app.FinalizeBlock(&abci.RequestFinalizeBlock{Height: app.LastBlockHeight() + 1, Time: p.now})
	require.NoError(t, err)
	_, err = app.Commit()
	require.NoError(t, err)
	app.EVMIndexer().Wait()
	res, err = app.CheckTx(&abci.RequestCheckTx{Tx: valid, Type: abci.CheckTxType_Recheck})
	require.NoError(t, err)
	require.Equal(t, abci.CodeTypeOK, res.Code, res.Log)
	p.block(t, nil)

	// an unlisted evm/ denom is stopped before its code runs, alone or next to a valid fee
	spin := deployRuntime(t, app, key, []byte{0x5b, 0x60, 0x00, 0x56})
	spinDenom := evmtypes.DENOM_PREFIX + spin.Hex()[2:]
	for _, fee := range []sdk.Coins{
		sdk.NewCoins(sdk.NewCoin(spinDenom, math.OneInt())),
		sdk.NewCoins(price(feeDenom, 20), sdk.NewCoin(spinDenom, math.OneInt())),
	} {
		txBz := buildSignedFeeTx(t, app, key, gas, fee)
		res, err := app.CheckTx(&abci.RequestCheckTx{Tx: txBz, Type: abci.CheckTxType_New})
		require.NoError(t, err)
		require.NotEqual(t, abci.CodeTypeOK, res.Code)
		require.Contains(t, res.Log, "is not allowed")
		require.Less(t, res.GasUsed, int64(1_000_000))
		_, _, err = app.Simulate(txBz)
		require.ErrorContains(t, err, "is not allowed")
	}
}
