package ante_test

import (
	"cosmossdk.io/math"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authsign "github.com/cosmos/cosmos-sdk/x/auth/signing"

	"github.com/initia-labs/minievm/app/ante"
	evmtypes "github.com/initia-labs/minievm/x/evm/types"
)

func (suite *AnteTestSuite) TestGasPricesDecorator() {
	suite.SetupTest() // setup
	suite.txBuilder = suite.clientCtx.TxConfig.NewTxBuilder()

	// keys and addresses
	priv1, _, _ := testdata.KeyTestPubAddr()

	feeAmount := sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, math.NewInt(100)))
	gasLimit := uint64(200_000)
	suite.txBuilder.SetFeeAmount(feeAmount)
	suite.txBuilder.SetGasLimit(gasLimit)

	privs, accNums, accSeqs := []cryptotypes.PrivKey{priv1}, []uint64{0}, []uint64{0}
	defaultSignMode, err := authsign.APISignModeToInternal(suite.app.TxConfig().SignModeHandler().DefaultMode())
	suite.NoError(err)
	tx, err := suite.CreateTestTx(privs, accNums, accSeqs, suite.ctx.ChainID(), defaultSignMode)
	suite.Require().NoError(err)

	decorator := ante.NewGasPricesDecorator()

	// in normal mode
	ctx, err := decorator.AnteHandle(suite.ctx, tx, false, nil)
	suite.Require().NoError(err)
	suite.Require().Equal(sdk.NewDecCoinsFromCoins(feeAmount...).QuoDec(math.LegacyNewDec(int64(gasLimit))), ctx.Value(evmtypes.CONTEXT_KEY_GAS_PRICES).(sdk.DecCoins))

	// in simulation mode
	ctx, err = decorator.AnteHandle(suite.ctx, tx, true, nil)
	suite.Require().NoError(err)
	suite.Require().Nil(ctx.Value(evmtypes.CONTEXT_KEY_GAS_PRICES))
}
