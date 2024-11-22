package indexer_test

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

	abcitypes "github.com/cometbft/cometbft/abci/types"

	"github.com/initia-labs/initia/crypto/ethsecp256k1"

	minitiaapp "github.com/initia-labs/minievm/app"
	"github.com/initia-labs/minievm/x/evm/contracts/erc20"
	"github.com/initia-labs/minievm/x/evm/contracts/erc20_factory"
	"github.com/initia-labs/minievm/x/evm/contracts/initia_erc20"
	evmkeeper "github.com/initia-labs/minievm/x/evm/keeper"
	evmtypes "github.com/initia-labs/minievm/x/evm/types"
)

func generateKeys(t *testing.T, n int) ([]common.Address, []*ecdsa.PrivateKey) {
	addrs := make([]common.Address, n)
	privKeys := make([]*ecdsa.PrivateKey, n)
	for i := 0; i < n; i++ {
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

func generateTx(
	t *testing.T,
	app *minitiaapp.MinitiaApp,
	privKey *ecdsa.PrivateKey,
	to *common.Address,
	inputBz []byte,
	seqNum ...uint64,
) (sdk.Tx, common.Hash) {
	ctx, err := app.CreateQueryContext(0, false)
	require.NoError(t, err)

	gasLimit := new(big.Int).SetUint64(1_000_000)
	gasPrice := new(big.Int).SetUint64(1_000_000_000)

	nonce := uint64(0)
	if len(seqNum) == 0 {
		cosmosKey := ethsecp256k1.PrivKey{Key: crypto.FromECDSA(privKey)}
		addrBz := cosmosKey.PubKey().Address()
		nonce, err = app.AccountKeeper.GetSequence(ctx, sdk.AccAddress(addrBz))
		require.NoError(t, err)
	} else {
		nonce = seqNum[0]
	}

	ethChainID := evmtypes.ConvertCosmosChainIDToEthereumChainID(ctx.ChainID())
	ethTx := coretypes.NewTx(&coretypes.DynamicFeeTx{
		ChainID:    ethChainID,
		Nonce:      nonce,
		GasTipCap:  big.NewInt(0),
		GasFeeCap:  gasPrice,
		Gas:        gasLimit.Uint64(),
		To:         to,
		Data:       inputBz,
		Value:      new(big.Int).SetUint64(0),
		AccessList: coretypes.AccessList{},
	})

	signer := coretypes.LatestSignerForChainID(ethChainID)
	signedTx, err := coretypes.SignTx(ethTx, signer, privKey)
	require.NoError(t, err)

	// Convert to cosmos tx
	sdkTx, err := evmkeeper.NewTxUtils(app.EVMKeeper).ConvertEthereumTxToCosmosTx(ctx, signedTx)
	require.NoError(t, err)

	return sdkTx, signedTx.Hash()
}

func generateCreateInitiaERC20Tx(t *testing.T, app *minitiaapp.MinitiaApp, privKey *ecdsa.PrivateKey, seqNum ...uint64) (sdk.Tx, common.Hash) {
	abi, err := initia_erc20.InitiaErc20MetaData.GetAbi()
	require.NoError(t, err)

	bin, err := hexutil.Decode(initia_erc20.InitiaErc20MetaData.Bin)
	require.NoError(t, err)

	inputBz, err := abi.Pack("", "foo", "foo", uint8(6))
	require.NoError(t, err)

	return generateTx(t, app, privKey, nil, append(bin, inputBz...), seqNum...)
}

func generateCreateERC20Tx(t *testing.T, app *minitiaapp.MinitiaApp, privKey *ecdsa.PrivateKey, seqNum ...uint64) (sdk.Tx, common.Hash) {
	ctx, err := app.CreateQueryContext(0, false)
	require.NoError(t, err)

	ethFactoryAddr, err := app.EVMKeeper.GetERC20FactoryAddr(ctx)
	require.NoError(t, err)

	abi, err := erc20_factory.Erc20FactoryMetaData.GetAbi()
	require.NoError(t, err)

	inputBz, err := abi.Pack("createERC20", "foo", "foo", uint8(6))
	require.NoError(t, err)

	return generateTx(t, app, privKey, &ethFactoryAddr, inputBz, seqNum...)
}

func generateMintERC20Tx(t *testing.T, app *minitiaapp.MinitiaApp, privKey *ecdsa.PrivateKey, erc20Addr, recipient common.Address, amount *big.Int, seqNum ...uint64) (sdk.Tx, common.Hash) {
	abi, err := erc20.Erc20MetaData.GetAbi()
	require.NoError(t, err)

	inputBz, err := abi.Pack("mint", recipient, amount)
	require.NoError(t, err)

	return generateTx(t, app, privKey, &erc20Addr, inputBz, seqNum...)
}

func generateTransferERC20Tx(t *testing.T, app *minitiaapp.MinitiaApp, privKey *ecdsa.PrivateKey, erc20Addr, recipient common.Address, amount *big.Int, seqNum ...uint64) (sdk.Tx, common.Hash) {
	abi, err := erc20.Erc20MetaData.GetAbi()
	require.NoError(t, err)

	inputBz, err := abi.Pack("transfer", recipient, amount)
	require.NoError(t, err)

	return generateTx(t, app, privKey, &erc20Addr, inputBz, seqNum...)
}

// execute txs and finalize block and commit block
func executeTxs(t *testing.T, app *minitiaapp.MinitiaApp, txs ...sdk.Tx) (*abcitypes.RequestFinalizeBlock, *abcitypes.ResponseFinalizeBlock) {
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

	app.Commit()

	return finalizeReq, finalizeRes
}

func checkTxResult(t *testing.T, txResult *abcitypes.ExecTxResult, expectSuccess bool) {
	if expectSuccess {
		require.Equal(t, abcitypes.CodeTypeOK, txResult.Code)
	} else {
		require.NotEqual(t, abcitypes.CodeTypeOK, txResult.Code)
	}
}
