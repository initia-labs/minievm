package posthandler_test

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/rand"
	"math/big"

	"cosmossdk.io/math"
	storetypes "cosmossdk.io/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/ethereum/go-ethereum/common"
	coretypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/initia-labs/initia/crypto/ethsecp256k1"
	"github.com/initia-labs/minievm/app/posthandler"
	"github.com/initia-labs/minievm/x/evm/contracts/erc20_factory"
	"github.com/initia-labs/minievm/x/evm/keeper"
	"github.com/initia-labs/minievm/x/evm/types"
	evmtypes "github.com/initia-labs/minievm/x/evm/types"
)

func (suite *PostHandlerTestSuite) Test_NotSpendingGasForTxWithFeeDenom() {
	suite.SetupTest() // setup
	suite.txBuilder = suite.clientCtx.TxConfig.NewTxBuilder()

	gasRefundPostHandler := posthandler.NewGasRefundDecorator(suite.app.Logger(), suite.app.EVMKeeper)

	params, err := suite.app.EVMKeeper.Params.Get(suite.ctx)
	suite.NoError(err)

	// create fee token
	decimals := uint8(18)
	feeCollectorAddr := authtypes.NewModuleAddress(authtypes.FeeCollectorName)
	suite.app.EVMKeeper.InitializeWithDecimals(suite.ctx, decimals)

	// mint fee token to fee collector
	gasPrice := math.NewInt(1_000_000_000)
	gasLimit := uint64(1_000_000)
	paidFeeAmount := sdk.NewCoins(sdk.NewCoin(params.FeeDenom, gasPrice.Mul(math.NewIntFromUint64(gasLimit))))
	err = suite.app.EVMKeeper.ERC20Keeper().MintCoins(suite.ctx, feeCollectorAddr, paidFeeAmount)
	suite.NoError(err)

	feeAmount := new(big.Int).Mul(
		big.NewInt(int64(gasLimit)),
		new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(decimals)-8), nil), // gas price is 1e-8
	)

	ethFactoryAddr, err := suite.app.EVMKeeper.GetERC20FactoryAddr(suite.ctx)
	suite.NoError(err)

	abi, err := erc20_factory.Erc20FactoryMetaData.GetAbi()
	suite.NoError(err)

	inputBz, err := abi.Pack("createERC20", "bar", "bar", uint8(6))
	suite.NoError(err)

	gasFeeCap := types.ToEthersUnit(decimals, feeAmount)
	gasFeeCap = gasFeeCap.Quo(gasFeeCap, new(big.Int).SetUint64(gasLimit))
	value := types.ToEthersUnit(decimals, big.NewInt(100))

	ethChainID := types.ConvertCosmosChainIDToEthereumChainID(suite.ctx.ChainID())
	ethTx := coretypes.NewTx(&coretypes.DynamicFeeTx{
		ChainID:   types.ConvertCosmosChainIDToEthereumChainID(suite.ctx.ChainID()),
		Nonce:     100,
		GasTipCap: big.NewInt(100),
		GasFeeCap: gasFeeCap,
		Gas:       gasLimit,
		To:        &ethFactoryAddr,
		Data:      inputBz,
		Value:     value,
		AccessList: coretypes.AccessList{
			coretypes.AccessTuple{Address: ethFactoryAddr,
				StorageKeys: []common.Hash{
					common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
					common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000001"),
					common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000002"),
				}},
		},
	})

	randBytes := make([]byte, 64)
	_, err = rand.Read(randBytes)
	suite.NoError(err)
	reader := bytes.NewReader(randBytes)
	privKey, err := ecdsa.GenerateKey(crypto.S256(), reader)
	suite.NoError(err)
	signer := coretypes.LatestSignerForChainID(ethChainID)
	signedTx, err := coretypes.SignTx(ethTx, signer, privKey)
	suite.NoError(err)

	// Compute sender address
	cosmosKey := ethsecp256k1.PrivKey{
		Key: crypto.FromECDSA(privKey),
	}
	addrBz := cosmosKey.PubKey().Address()

	// Convert to cosmos tx
	sdkTx, err := keeper.NewTxUtils(suite.app.EVMKeeper).ConvertEthereumTxToCosmosTx(suite.ctx, signedTx)
	suite.NoError(err)

	// Spend half of the gas
	gasMeter := storetypes.NewGasMeter(gasLimit)
	gasMeter.ConsumeGas(gasLimit/2-1114 /* extra gas for params loading */, "test")
	gasPrices := sdk.DecCoins{sdk.NewDecCoin(params.FeeDenom, gasPrice)}

	ctx := sdk.UnwrapSDKContext(suite.ctx).WithValue(evmtypes.CONTEXT_KEY_GAS_PRICES, gasPrices)
	ctx = ctx.WithGasMeter(gasMeter).WithExecMode(sdk.ExecModeFinalize)
	ctx = ctx.WithValue(types.CONTEXT_KEY_ETH_TX, true)
	ctx, err = gasRefundPostHandler.PostHandle(ctx, sdkTx, false, true, func(ctx sdk.Context, tx sdk.Tx, simulate, success bool) (newCtx sdk.Context, err error) {
		return ctx, nil
	})
	suite.NoError(err)

	gasRefundRatio := params.GasRefundRatio
	sender := sdk.AccAddress(addrBz.Bytes())

	// Check the gas refund
	amount := suite.app.BankKeeper.GetBalance(ctx, sender, params.FeeDenom)
	refunds, _ := gasPrices.MulDec(gasRefundRatio.MulInt64(int64(gasLimit / 2))).TruncateDecimal()
	suite.Equal(amount, refunds[0])
}
