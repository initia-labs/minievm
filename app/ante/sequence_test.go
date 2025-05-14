package ante_test

import (
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authsign "github.com/cosmos/cosmos-sdk/x/auth/signing"

	"github.com/initia-labs/minievm/app/ante"
	evmtypes "github.com/initia-labs/minievm/x/evm/types"
)

func (suite *AnteTestSuite) TestIncrementSequenceDecorator() {
	suite.SetupTest() // setup
	suite.txBuilder = suite.clientCtx.TxConfig.NewTxBuilder()

	priv, _, addr := testdata.KeyTestPubAddr()
	acc := suite.app.AccountKeeper.NewAccountWithAddress(suite.ctx, addr)
	suite.NoError(acc.SetAccountNumber(uint64(50)))
	suite.app.AccountKeeper.SetAccount(suite.ctx, acc)

	msgs := []sdk.Msg{testdata.NewTestMsg(addr)}
	suite.NoError(suite.txBuilder.SetMsgs(msgs...))
	privs := []cryptotypes.PrivKey{priv}
	accNums := []uint64{suite.app.AccountKeeper.GetAccount(suite.ctx, addr).GetAccountNumber()}
	accSeqs := []uint64{suite.app.AccountKeeper.GetAccount(suite.ctx, addr).GetSequence()}
	feeAmount := testdata.NewTestFeeAmount()
	gasLimit := testdata.NewTestGasLimit()
	suite.txBuilder.SetFeeAmount(feeAmount)
	suite.txBuilder.SetGasLimit(gasLimit)

	defaultSignMode, err := authsign.APISignModeToInternal(suite.app.TxConfig().SignModeHandler().DefaultMode())
	suite.NoError(err)

	tx, err := suite.CreateTestTx(privs, accNums, accSeqs, suite.ctx.ChainID(), defaultSignMode)
	suite.NoError(err)

	isd := ante.NewIncrementSequenceDecorator(suite.app.AccountKeeper)
	antehandler := sdk.ChainAnteDecorators(isd)

	testCases := []struct {
		ctx         sdk.Context
		simulate    bool
		expectedSeq uint64
	}{
		{suite.ctx.WithIsReCheckTx(true), false, 1},
		{suite.ctx.WithIsCheckTx(true).WithIsReCheckTx(false), false, 2},
		{suite.ctx.WithIsReCheckTx(true), false, 3},
		{suite.ctx.WithIsReCheckTx(true), false, 4},
		{suite.ctx.WithIsReCheckTx(true), true, 5},
	}

	for i, tc := range testCases {
		ctx, err := antehandler(tc.ctx, tx, tc.simulate)
		suite.NoError(err, "unexpected error; tc #%d, %v", i, tc)
		suite.Equal(tc.expectedSeq, suite.app.AccountKeeper.GetAccount(suite.ctx, addr).GetSequence())

		// the flag should be set in the context
		suite.NotNil(ctx.Value(evmtypes.CONTEXT_KEY_SEQUENCE_INCREMENTED))
	}
}
