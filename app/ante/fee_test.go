package ante_test

import (
	"cosmossdk.io/math"
	storetypes "cosmossdk.io/store/types"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authsign "github.com/cosmos/cosmos-sdk/x/auth/signing"

	"github.com/initia-labs/minievm/app/ante"
)

func (suite *AnteTestSuite) Test_NotSpendingGasForTxWithFeeDenom() {
	suite.SetupTest() // setup
	suite.txBuilder = suite.clientCtx.TxConfig.NewTxBuilder()

	feeAnte := ante.NewGasFreeFeeDecorator(suite.app.AccountKeeper, suite.app.BankKeeper, suite.app.FeeGrantKeeper, suite.app.EVMKeeper, nil)

	// keys and addresses
	priv1, _, addr1 := testdata.KeyTestPubAddr()

	msg := testdata.NewTestMsg(addr1)
	feeAmount := sdk.NewCoins(sdk.NewCoin(feeDenom, math.NewInt(100)))
	gasLimit := uint64(200_000)
	atomFeeAmount := sdk.NewCoins(sdk.NewCoin("atom", math.NewInt(200)))

	suite.app.EVMKeeper.ERC20Keeper().MintCoins(suite.ctx, addr1, feeAmount.MulInt(math.NewInt(10)))
	suite.app.EVMKeeper.ERC20Keeper().MintCoins(suite.ctx, addr1, atomFeeAmount.MulInt(math.NewInt(10)))

	// Case 1. only fee denom
	suite.Require().NoError(suite.txBuilder.SetMsgs(msg))
	suite.txBuilder.SetFeeAmount(feeAmount)
	suite.txBuilder.SetGasLimit(gasLimit)

	privs, accNums, accSeqs := []cryptotypes.PrivKey{priv1}, []uint64{0}, []uint64{0}
	defaultSignMode, err := authsign.APISignModeToInternal(suite.app.TxConfig().SignModeHandler().DefaultMode())
	suite.NoError(err)
	tx, err := suite.CreateTestTx(privs, accNums, accSeqs, suite.ctx.ChainID(), defaultSignMode)
	suite.Require().NoError(err)

	gasMeter := storetypes.NewGasMeter(500000)
	feeAnte.AnteHandle(suite.ctx.WithGasMeter(gasMeter), tx, false, nil)
	suite.Require().Zero(gasMeter.GasConsumed(), "should not consume gas for fee deduction")

	// Case 2. fee denom and other denom
	suite.txBuilder.SetFeeAmount(feeAmount.Add(atomFeeAmount...))

	gasMeter = storetypes.NewGasMeter(500000)
	feeAnte.AnteHandle(suite.ctx.WithGasMeter(gasMeter), tx, false, nil)
	suite.Require().NotZero(gasMeter.GasConsumed(), "should consume gas for fee deduction")

	// Case 3. other denom
	suite.txBuilder.SetFeeAmount(feeAmount.Add(atomFeeAmount...))

	gasMeter = storetypes.NewGasMeter(500000)
	feeAnte.AnteHandle(suite.ctx.WithGasMeter(gasMeter), tx, false, nil)
	suite.Require().NotZero(gasMeter.GasConsumed(), "should consume gas for fee deduction")

	// Case 4. no fee
	suite.txBuilder.SetFeeAmount(sdk.NewCoins())

	gasMeter = storetypes.NewGasMeter(500000)
	feeAnte.AnteHandle(suite.ctx.WithGasMeter(gasMeter), tx, false, nil)
	suite.Require().NotZero(gasMeter.GasConsumed(), "should consume gas for fee deduction")

	// Case 5. simulate gas consumption
	suite.txBuilder.SetFeeAmount(sdk.NewCoins())

	gasMeter = storetypes.NewGasMeter(500000)
	feeAnte.AnteHandle(suite.ctx.WithGasMeter(gasMeter), tx, true, nil)
	suite.Require().Greater(gasMeter.GasConsumed(), uint64(250000), "should consume gas for fee deduction")
}
