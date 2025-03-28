package ante_test

import (
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"

	"github.com/initia-labs/minievm/app/ante"
	evmkeeper "github.com/initia-labs/minievm/x/evm/keeper"

	coretypes "github.com/ethereum/go-ethereum/core/types"
)

func (suite *AnteTestSuite) Test_SkipSequenceCheck() {
	suite.SetupTest() // setup
	suite.txBuilder = suite.clientCtx.TxConfig.NewTxBuilder()

	priv1, _, addr1 := testdata.KeyTestPubAddr()
	acc1 := suite.app.AccountKeeper.NewAccountWithAddress(suite.ctx, addr1)
	acc1.SetPubKey(priv1.PubKey())
	suite.app.AccountKeeper.SetAccount(suite.ctx, acc1)

	suite.txBuilder.SetMsgs(testdata.NewTestMsg(addr1))
	suite.txBuilder.SetSignatures()

	sigV2 := signing.SignatureV2{
		PubKey: priv1.PubKey(),
		Data: &signing.SingleSignatureData{
			SignMode:  evmkeeper.SignMode_SIGN_MODE_ETHEREUM,
			Signature: nil,
		},
		// invalid sequence
		Sequence: 100,
	}

	err := suite.txBuilder.SetSignatures(sigV2)
	suite.NoError(err)

	// 1. simulate should skip sequence check
	suite.ctx = suite.ctx.WithValue(ante.ContextKeyEthTx, &coretypes.Transaction{})
	sigVerifyAnte := ante.NewSigVerificationDecorator(suite.app.AccountKeeper, suite.app.TxConfig().SignModeHandler())
	_, err = sigVerifyAnte.AnteHandle(suite.ctx, suite.txBuilder.GetTx(), true, func(ctx sdk.Context, tx sdk.Tx, simulate bool) (newCtx sdk.Context, err error) { return ctx, nil })
	suite.NoError(err)

	// 2. simulate should check sequence when it is not ethereum tx
	suite.ctx = suite.ctx.WithValue(ante.ContextKeyEthTx, nil)
	err = suite.txBuilder.SetSignatures(sigV2)
	suite.NoError(err)
	_, err = sigVerifyAnte.AnteHandle(suite.ctx, suite.txBuilder.GetTx(), true, func(ctx sdk.Context, tx sdk.Tx, simulate bool) (newCtx sdk.Context, err error) { return ctx, nil })
	suite.ErrorIs(err, sdkerrors.ErrWrongSequence)

	// 3. non-simulate should check sequence
	suite.ctx = suite.ctx.WithValue(ante.ContextKeyEthTx, &coretypes.Transaction{})
	err = suite.txBuilder.SetSignatures(sigV2)
	suite.NoError(err)
	_, err = sigVerifyAnte.AnteHandle(suite.ctx, suite.txBuilder.GetTx(), false, func(ctx sdk.Context, tx sdk.Tx, simulate bool) (newCtx sdk.Context, err error) { return ctx, nil })
	suite.ErrorIs(err, sdkerrors.ErrWrongSequence)
}
