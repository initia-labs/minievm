package keeper_test

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/rand"
	"math/big"
	"testing"

	"cosmossdk.io/math"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	"github.com/ethereum/go-ethereum/common/hexutil"
	coretypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/initia-labs/initia/crypto/ethsecp256k1"
	"github.com/initia-labs/minievm/x/evm/contracts/erc20_factory"
	"github.com/initia-labs/minievm/x/evm/keeper"
	"github.com/initia-labs/minievm/x/evm/types"
)

func Test_DynamicFeeTxConversion(t *testing.T) {
	ctx, input := createDefaultTestInput(t)

	_, _, addr := keyPubAddr()
	input.Faucet.Mint(ctx, addr, sdk.NewCoin(sdk.DefaultBondDenom, math.NewInt(1000000000000000000)))

	feeAmount := int64(150000)
	gasLimit := uint64(1000000)
	ethFactoryAddr := types.ERC20FactoryAddress()

	abi, err := erc20_factory.Erc20FactoryMetaData.GetAbi()
	require.NoError(t, err)

	inputBz, err := abi.Pack("createERC20", "bar", "bar", uint8(6))
	require.NoError(t, err)

	gasFeeCap := types.ToEthersUint(0, big.NewInt(feeAmount))
	gasFeeCap = gasFeeCap.Quo(gasFeeCap, new(big.Int).SetUint64(gasLimit))
	value := types.ToEthersUint(0, big.NewInt(100))

	ethChainID := types.ConvertCosmosChainIDToEthereumChainID(ctx.ChainID())
	ethTx := coretypes.NewTx(&coretypes.DynamicFeeTx{
		ChainID:   types.ConvertCosmosChainIDToEthereumChainID(ctx.ChainID()),
		Nonce:     100,
		GasTipCap: big.NewInt(100),
		GasFeeCap: gasFeeCap,
		Gas:       gasLimit,
		To:        &ethFactoryAddr,
		Data:      inputBz,
		Value:     value,
	})

	signer := coretypes.LatestSignerForChainID(ethChainID)

	randBytes := make([]byte, 64)
	_, err = rand.Read(randBytes)
	require.NoError(t, err)
	reader := bytes.NewReader(randBytes)

	privKey, err := ecdsa.GenerateKey(crypto.S256(), reader)
	require.NoError(t, err)
	signedTx, err := coretypes.SignTx(ethTx, signer, privKey)
	require.NoError(t, err)

	cosmosKey := ethsecp256k1.PrivKey{
		Key: crypto.FromECDSA(privKey),
	}
	addrBz := cosmosKey.PubKey().Address()

	// Convert to cosmos tx
	sdkTx, err := keeper.NewTxUtils(&input.EVMKeeper).ConvertEthereumTxToCosmosTx(ctx, signedTx)
	require.NoError(t, err)

	msgs := sdkTx.GetMsgs()
	require.Len(t, msgs, 1)
	msg, ok := msgs[0].(*types.MsgCall)
	require.True(t, ok)
	require.Equal(t, msg, &types.MsgCall{
		Sender:       sdk.AccAddress(addrBz).String(),
		ContractAddr: ethFactoryAddr.Hex(),
		Input:        hexutil.Encode(inputBz),
		Value:        math.NewInt(100),
	})

	authTx := sdkTx.(authsigning.Tx)
	require.Equal(t, authTx.GetFee(), sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, math.NewInt(feeAmount+1))))

	sigs, err := authTx.GetSignaturesV2()
	require.NoError(t, err)
	require.Len(t, sigs, 1)

	sig := sigs[0]
	require.Equal(t, sig.PubKey, cosmosKey.PubKey())
	require.Equal(t, sig.Sequence, uint64(100))

	v, r, s := signedTx.RawSignatureValues()
	sigData := sig.Data.(*signing.SingleSignatureData)
	require.Equal(t, sigData.SignMode, keeper.SignMode_SIGN_MODE_ETHEREUM)
	require.Equal(t, sigData.Signature, append(append(r.Bytes(), s.Bytes()...), byte(v.Uint64())))

	// Convert back to ethereum tx
	ethTx2, _, err := keeper.NewTxUtils(&input.EVMKeeper).ConvertCosmosTxToEthereumTx(ctx, sdkTx)
	require.NoError(t, err)
	EqualEthTransaction(t, signedTx, ethTx2)
}

func Test_LegacyTxConversion(t *testing.T) {
	ctx, input := createDefaultTestInput(t)

	_, _, addr := keyPubAddr()
	input.Faucet.Mint(ctx, addr, sdk.NewCoin(sdk.DefaultBondDenom, math.NewInt(1000000000000000000)))

	feeAmount := int64(150000)
	gasLimit := uint64(1000000)
	ethFactoryAddr := types.ERC20FactoryAddress()

	abi, err := erc20_factory.Erc20FactoryMetaData.GetAbi()
	require.NoError(t, err)

	inputBz, err := abi.Pack("createERC20", "bar", "bar", uint8(6))
	require.NoError(t, err)

	gasFeeCap := types.ToEthersUint(0, big.NewInt(feeAmount))
	gasFeeCap = gasFeeCap.Quo(gasFeeCap, new(big.Int).SetUint64(gasLimit))
	value := types.ToEthersUint(0, big.NewInt(100))

	ethChainID := types.ConvertCosmosChainIDToEthereumChainID(ctx.ChainID())
	ethTx := coretypes.NewTx(&coretypes.LegacyTx{
		Nonce:    100,
		GasPrice: gasFeeCap,
		Gas:      gasLimit,
		To:       &ethFactoryAddr,
		Data:     inputBz,
		Value:    value,
	})

	signer := coretypes.LatestSignerForChainID(ethChainID)

	randBytes := make([]byte, 64)
	_, err = rand.Read(randBytes)
	require.NoError(t, err)
	reader := bytes.NewReader(randBytes)

	privKey, err := ecdsa.GenerateKey(crypto.S256(), reader)
	require.NoError(t, err)
	signedTx, err := coretypes.SignTx(ethTx, signer, privKey)
	require.NoError(t, err)

	cosmosKey := ethsecp256k1.PrivKey{
		Key: crypto.FromECDSA(privKey),
	}
	addrBz := cosmosKey.PubKey().Address()

	// Convert to cosmos tx
	sdkTx, err := keeper.NewTxUtils(&input.EVMKeeper).ConvertEthereumTxToCosmosTx(ctx, signedTx)
	require.NoError(t, err)

	msgs := sdkTx.GetMsgs()
	require.Len(t, msgs, 1)
	msg, ok := msgs[0].(*types.MsgCall)
	require.True(t, ok)
	require.Equal(t, msg, &types.MsgCall{
		Sender:       sdk.AccAddress(addrBz).String(),
		ContractAddr: ethFactoryAddr.Hex(),
		Input:        hexutil.Encode(inputBz),
		Value:        math.NewInt(100),
	})

	authTx := sdkTx.(authsigning.Tx)
	require.Equal(t, authTx.GetFee(), sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, math.NewInt(feeAmount+1))))

	sigs, err := authTx.GetSignaturesV2()
	require.NoError(t, err)
	require.Len(t, sigs, 1)

	sig := sigs[0]
	require.Equal(t, sig.PubKey, cosmosKey.PubKey())
	require.Equal(t, sig.Sequence, uint64(100))

	v, r, s := signedTx.RawSignatureValues()
	sigData := sig.Data.(*signing.SingleSignatureData)
	require.Equal(t, sigData.SignMode, keeper.SignMode_SIGN_MODE_ETHEREUM)

	sigBytes := make([]byte, 65)
	copy(sigBytes[32-len(r.Bytes()):32], r.Bytes())
	copy(sigBytes[64-len(s.Bytes()):64], s.Bytes())
	sigBytes[64] = byte(v.Uint64() - (35 + ethChainID.Uint64()*2))

	require.Equal(t, sigData.Signature, sigBytes)

	// Convert back to ethereum tx
	ethTx2, _, err := keeper.NewTxUtils(&input.EVMKeeper).ConvertCosmosTxToEthereumTx(ctx, sdkTx)
	require.NoError(t, err)
	EqualEthTransaction(t, signedTx, ethTx2)

}

func EqualEthTransaction(t *testing.T, expected, actual *coretypes.Transaction) {
	require.Equal(t, expected.ChainId(), actual.ChainId())
	require.Equal(t, expected.Nonce(), actual.Nonce())
	require.Equal(t, expected.GasTipCap(), actual.GasTipCap())
	require.Equal(t, expected.GasFeeCap(), actual.GasFeeCap())
	require.Equal(t, expected.Gas(), actual.Gas())
	require.Equal(t, expected.To(), actual.To())
	require.Equal(t, expected.Data(), actual.Data())
	require.Equal(t, expected.Value(), actual.Value())
	require.Equal(t, expected.Type(), actual.Type())
}
