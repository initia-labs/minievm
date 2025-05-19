package tests

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/rand"
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	coretypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	authsign "github.com/cosmos/cosmos-sdk/x/auth/signing"

	abcitypes "github.com/cometbft/cometbft/abci/types"

	"github.com/initia-labs/initia/crypto/ethsecp256k1"

	minitiaapp "github.com/initia-labs/minievm/app"
	"github.com/initia-labs/minievm/x/evm/contracts/erc20"
	"github.com/initia-labs/minievm/x/evm/contracts/erc20_factory"
	"github.com/initia-labs/minievm/x/evm/contracts/initia_erc20"
	evmkeeper "github.com/initia-labs/minievm/x/evm/keeper"
	evmtypes "github.com/initia-labs/minievm/x/evm/types"
)

func GenerateKeys(t *testing.T, n int) ([]common.Address, []*ecdsa.PrivateKey) {
	addrs := make([]common.Address, n)
	privKeys := make([]*ecdsa.PrivateKey, n)
	for i := range n {
		randBytes := make([]byte, 64)
		_, err := rand.Read(randBytes)
		require.NoError(t, err)
		reader := bytes.NewReader(randBytes)
		privKey, err := ecdsa.GenerateKey(crypto.S256(), reader)
		require.NoError(t, err)

		cosmosKey := ethsecp256k1.PrivKey{
			Key: crypto.FromECDSA(privKey),
		}
		addrBz := cosmosKey.PubKey().Address()
		addr := common.BytesToAddress(addrBz)

		addrs[i] = addr
		privKeys[i] = privKey
	}

	return addrs, privKeys
}

// Opt is a function that modifies the tx
type Opt func(tx *coretypes.DynamicFeeTx)

// SetNonce sets the nonce of the tx
func SetNonce(nonce uint64) Opt {
	return func(tx *coretypes.DynamicFeeTx) {
		tx.Nonce = nonce
	}
}

// SetGasFeeCap sets the gas fee cap of the tx
func SetGasFeeCap(gasFeeCap *big.Int) Opt {
	return func(tx *coretypes.DynamicFeeTx) {
		tx.GasFeeCap = gasFeeCap
	}
}

// SetGasLimit sets the gas limit of the tx
func SetGasLimit(gasLimit uint64) Opt {
	return func(tx *coretypes.DynamicFeeTx) {
		tx.Gas = gasLimit
	}
}

// SetGasTipCap sets the gas tip cap of the tx
func SetGasTipCap(gasTipCap *big.Int) Opt {
	return func(tx *coretypes.DynamicFeeTx) {
		tx.GasTipCap = gasTipCap
	}
}

// GenerateTx generates a tx with the given parameters
func GenerateTx(
	t *testing.T,
	app *minitiaapp.MinitiaApp,
	privKey *ecdsa.PrivateKey,
	to *common.Address,
	inputBz []byte,
	value *big.Int,
	opts ...Opt,
) (sdk.Tx, common.Hash) {
	ctx, err := app.CreateQueryContext(0, false)
	require.NoError(t, err)

	gasLimit := new(big.Int).SetUint64(1_000_000)
	gasPrice := new(big.Int).SetUint64(1_000_000_000)

	ethChainID := evmtypes.ConvertCosmosChainIDToEthereumChainID(ctx.ChainID())
	dynamicFeeTx := &coretypes.DynamicFeeTx{
		ChainID:    ethChainID,
		Nonce:      0,
		GasTipCap:  big.NewInt(0),
		GasFeeCap:  gasPrice,
		Gas:        gasLimit.Uint64(),
		To:         to,
		Data:       inputBz,
		Value:      value,
		AccessList: coretypes.AccessList{},
	}
	for _, opt := range opts {
		opt(dynamicFeeTx)
	}
	if dynamicFeeTx.Nonce == 0 {
		cosmosKey := ethsecp256k1.PrivKey{Key: crypto.FromECDSA(privKey)}
		addrBz := cosmosKey.PubKey().Address()
		dynamicFeeTx.Nonce, err = app.AccountKeeper.GetSequence(ctx, sdk.AccAddress(addrBz))
		require.NoError(t, err)
	}

	ethTx := coretypes.NewTx(dynamicFeeTx)
	signer := coretypes.LatestSignerForChainID(ethChainID)
	signedTx, err := coretypes.SignTx(ethTx, signer, privKey)
	require.NoError(t, err)

	// Convert to cosmos tx
	sdkTx, err := evmkeeper.NewTxUtils(app.EVMKeeper).ConvertEthereumTxToCosmosTx(ctx, signedTx)
	require.NoError(t, err)

	return sdkTx, signedTx.Hash()
}

func GenerateCreateInitiaERC20Tx(t *testing.T, app *minitiaapp.MinitiaApp, privKey *ecdsa.PrivateKey, opts ...Opt) (sdk.Tx, common.Hash) {
	abi, err := initia_erc20.InitiaErc20MetaData.GetAbi()
	require.NoError(t, err)

	bin, err := hexutil.Decode(initia_erc20.InitiaErc20MetaData.Bin)
	require.NoError(t, err)

	inputBz, err := abi.Pack("", "foo", "foo", uint8(6))
	require.NoError(t, err)

	return GenerateTx(t, app, privKey, nil, append(bin, inputBz...), new(big.Int).SetUint64(0), opts...)
}

func GenerateCreateERC20Tx(t *testing.T, app *minitiaapp.MinitiaApp, privKey *ecdsa.PrivateKey, opts ...Opt) (sdk.Tx, common.Hash) {
	ctx, err := app.CreateQueryContext(0, false)
	require.NoError(t, err)

	ethFactoryAddr, err := app.EVMKeeper.GetERC20FactoryAddr(ctx)
	require.NoError(t, err)

	abi, err := erc20_factory.Erc20FactoryMetaData.GetAbi()
	require.NoError(t, err)

	inputBz, err := abi.Pack("createERC20", "foo", "foo", uint8(6))
	require.NoError(t, err)

	return GenerateTx(t, app, privKey, &ethFactoryAddr, inputBz, new(big.Int).SetUint64(0), opts...)
}

func GenerateMintERC20Tx(t *testing.T, app *minitiaapp.MinitiaApp, privKey *ecdsa.PrivateKey, erc20Addr, recipient common.Address, amount *big.Int, opts ...Opt) (sdk.Tx, common.Hash) {
	abi, err := erc20.Erc20MetaData.GetAbi()
	require.NoError(t, err)

	inputBz, err := abi.Pack("mint", recipient, amount)
	require.NoError(t, err)

	return GenerateTx(t, app, privKey, &erc20Addr, inputBz, new(big.Int).SetUint64(0), opts...)
}

func GenerateTransferERC20Tx(t *testing.T, app *minitiaapp.MinitiaApp, privKey *ecdsa.PrivateKey, erc20Addr, recipient common.Address, amount *big.Int, opts ...Opt) (sdk.Tx, common.Hash) {
	abi, err := erc20.Erc20MetaData.GetAbi()
	require.NoError(t, err)

	inputBz, err := abi.Pack("transfer", recipient, amount)
	require.NoError(t, err)

	return GenerateTx(t, app, privKey, &erc20Addr, inputBz, new(big.Int).SetUint64(0), opts...)
}

// execute txs and finalize block and commit block
func ExecuteTxs(t *testing.T, app *minitiaapp.MinitiaApp, txs ...sdk.Tx) (*abcitypes.RequestFinalizeBlock, *abcitypes.ResponseFinalizeBlock) {
	txsBytes := make([][]byte, len(txs))
	for i, tx := range txs {
		txBytes, err := app.TxConfig().TxEncoder()(tx)
		require.NoError(t, err)

		txsBytes[i] = txBytes
	}

	finalizeReq := &abcitypes.RequestFinalizeBlock{
		Txs:    txsBytes,
		Height: app.LastBlockHeight() + 1,
	}
	resBlock, err := app.FinalizeBlock(finalizeReq)
	require.NoError(t, err)

	finalizeRes := &abcitypes.ResponseFinalizeBlock{
		TxResults: resBlock.TxResults,
	}

	_, err = app.Commit()
	require.NoError(t, err)

	// wait for indexing to complete
	app.EVMIndexer().Wait()

	return finalizeReq, finalizeRes
}

func CheckTxResult(t *testing.T, txResult *abcitypes.ExecTxResult, expectSuccess bool) {
	if expectSuccess {
		require.Equal(t, abcitypes.CodeTypeOK, txResult.Code)
	} else {
		require.NotEqual(t, abcitypes.CodeTypeOK, txResult.Code)
	}
}

func IncreaseBlockHeight(t *testing.T, app *minitiaapp.MinitiaApp) {
	_, err := app.FinalizeBlock(&abcitypes.RequestFinalizeBlock{
		Height: app.LastBlockHeight() + 1,
	})
	require.NoError(t, err)

	_, err = app.Commit()
	require.NoError(t, err)
}

func GenerateCosmosTx(t *testing.T, app *minitiaapp.MinitiaApp, privKey *ecdsa.PrivateKey, msgs []sdk.Msg) sdk.Tx {
	txConfig := app.TxConfig()
	txBuilder := txConfig.NewTxBuilder()
	_ = txBuilder.SetMsgs(msgs...)
	txBuilder.SetMemo("test")
	txBuilder.SetGasLimit(1000000)

	// build empty signature
	signMode, err := authsign.APISignModeToInternal(txConfig.SignModeHandler().DefaultMode())
	require.NoError(t, err)

	ethPrivKey := ethsecp256k1.PrivKey{Key: crypto.FromECDSA(privKey)}
	ethPubKey := ethPrivKey.PubKey()

	ctx, err := app.CreateQueryContext(0, false)
	require.NoError(t, err)

	account := app.AccountKeeper.GetAccount(ctx, sdk.AccAddress(ethPubKey.Address().Bytes()))
	require.NotNil(t, account)

	sequence := account.GetSequence()
	accountNumber := account.GetAccountNumber()

	sig := signing.SignatureV2{
		PubKey: ethPubKey,
		Data: &signing.SingleSignatureData{
			SignMode: signMode,
		},
		Sequence: sequence,
	}

	err = txBuilder.SetSignatures(sig)
	require.NoError(t, err)
	signerData := authsign.SignerData{
		Address:       sdk.AccAddress(ethPubKey.Address().Bytes()).String(),
		ChainID:       ctx.ChainID(),
		AccountNumber: accountNumber,
		Sequence:      sequence,
		PubKey:        ethPubKey,
	}

	signBytes, err := authsign.GetSignBytesAdapter(
		ctx, txConfig.SignModeHandler(), signMode, signerData, txBuilder.GetTx(),
	)
	require.NoError(t, err)

	sigBytes, err := ethPrivKey.Sign(signBytes)
	require.NoError(t, err)

	sig.Data.(*signing.SingleSignatureData).Signature = sigBytes
	err = txBuilder.SetSignatures(sig)
	require.NoError(t, err)

	return txBuilder.GetTx()
}
