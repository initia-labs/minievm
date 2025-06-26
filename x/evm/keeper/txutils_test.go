package keeper_test

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/rand"
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"

	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"

	"github.com/ethereum/go-ethereum/common"
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
	decimals := input.Decimals
	gasLimit := uint64(1_000_000)
	feeAmount := new(big.Int).Mul(
		big.NewInt(int64(gasLimit)),
		new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(decimals)-8), nil), // gas price is 1e-8
	)

	ethFactoryAddr, err := input.EVMKeeper.GetERC20FactoryAddr(ctx)
	require.NoError(t, err)

	abi, err := erc20_factory.Erc20FactoryMetaData.GetAbi()
	require.NoError(t, err)

	inputBz, err := abi.Pack("createERC20", "bar", "bar", uint8(6))
	require.NoError(t, err)

	gasFeeCap := types.ToEthersUnit(decimals, feeAmount)
	gasFeeCap = gasFeeCap.Quo(gasFeeCap, new(big.Int).SetUint64(gasLimit))
	value := types.ToEthersUnit(decimals, big.NewInt(100))

	ethChainID := types.ConvertCosmosChainIDToEthereumChainID(ctx.ChainID())
	dynTx := &coretypes.DynamicFeeTx{
		ChainID:   types.ConvertCosmosChainIDToEthereumChainID(ctx.ChainID()),
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
	}
	ethTx := coretypes.NewTx(dynTx)
	randBytes := make([]byte, 64)
	_, err = rand.Read(randBytes)
	require.NoError(t, err)
	reader := bytes.NewReader(randBytes)
	privKey, err := ecdsa.GenerateKey(crypto.S256(), reader)
	require.NoError(t, err)
	signer := coretypes.LatestSignerForChainID(ethChainID)
	signedTx, err := coretypes.SignTx(ethTx, signer, privKey)
	require.NoError(t, err)

	cosmosKey := ethsecp256k1.PrivKey{
		Key: crypto.FromECDSA(privKey),
	}
	addrBz := cosmosKey.PubKey().Address()

	// 1. Create a dynamic fee tx without max gas configuration
	// Convert to cosmos tx
	sdkTx, err := keeper.NewTxUtils(&input.EVMKeeper).ConvertEthereumTxToCosmosTx(ctx, signedTx)
	require.NoError(t, err)
	expectedMsg := &types.MsgCall{
		Sender:       sdk.AccAddress(addrBz).String(),
		ContractAddr: ethFactoryAddr.Hex(),
		Input:        hexutil.Encode(inputBz),
		Value:        math.NewInt(100),
		AccessList: []types.AccessTuple{
			{
				Address: ethFactoryAddr.String(),
				StorageKeys: []string{
					common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000").Hex(),
					common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000001").Hex(),
					common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000002").Hex()},
			},
		},
	}
	msgs := sdkTx.GetMsgs()
	require.Len(t, msgs, 1)
	msg, ok := msgs[0].(*types.MsgCall)
	require.True(t, ok)
	require.Equal(t, msg, expectedMsg)

	authTx := sdkTx.(authsigning.Tx)
	expectedFeeAmount := sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, math.NewIntFromBigInt(feeAmount).AddRaw(1)))
	require.Equal(t, authTx.GetFee(), expectedFeeAmount)

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
	sigBytes[64] = byte(v.Uint64())

	require.Equal(t, sigData.Signature, sigBytes)

	// Convert back to ethereum tx
	ethTx2, _, err := keeper.NewTxUtils(&input.EVMKeeper).ConvertCosmosTxToEthereumTx(ctx, sdkTx)
	require.NoError(t, err)
	EqualEthTransaction(t, signedTx, ethTx2)

	// manipulate the fee amount
	txBuilder := sdkTx.(client.TxBuilder)
	txBuilder.SetFeeAmount(expectedFeeAmount.Add(sdk.NewCoin(sdk.DefaultBondDenom, math.NewInt(1))))
	_, _, err = keeper.NewTxUtils(&input.EVMKeeper).ConvertCosmosTxToEthereumTx(ctx, txBuilder.GetTx())
	require.ErrorIs(t, err, types.ErrTxConversionFailed)

	// 2. Set the max params of gas configuration
	// Set the gas enforcement parameters with unlimited sender
	maxGasLimit := gasLimit / 2
	params, err := input.EVMKeeper.Params.Get(ctx)
	require.NoError(t, err)
	maxGasFeeCap := math.NewIntFromBigInt(gasFeeCap.Div(gasFeeCap, big.NewInt(2)))

	gasEnforcement := &types.GasEnforcement{
		MaxGasLimit:  maxGasLimit,
		MaxGasFeeCap: &maxGasFeeCap,
		UnlimitedGasSenders: []string{
			common.BytesToAddress(addrBz.Bytes()).String(),
		},
	}
	params.GasEnforcement = gasEnforcement
	err = input.EVMKeeper.Params.Set(ctx, params)
	require.NoError(t, err)

	// Convert to cosmos tx
	// Since the sender is in the unlimited gas senders, the gas limit will be still the original gas limit
	sdkTx, err = keeper.NewTxUtils(&input.EVMKeeper).ConvertEthereumTxToCosmosTx(ctx, signedTx)
	require.NoError(t, err)

	msgs = sdkTx.GetMsgs()
	require.Len(t, msgs, 1)
	msg, ok = msgs[0].(*types.MsgCall)
	require.True(t, ok)

	require.Equal(t, msg, expectedMsg)

	authTx = sdkTx.(authsigning.Tx)
	expectedFeeAmount = sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, math.NewIntFromBigInt(feeAmount).AddRaw(1)))
	require.Equal(t, authTx.GetFee(), expectedFeeAmount)
	require.Equal(t, authTx.GetGas(), gasLimit)

	sigs, err = authTx.GetSignaturesV2()
	require.NoError(t, err)
	require.Len(t, sigs, 1)

	sig = sigs[0]
	require.Equal(t, sig.PubKey, cosmosKey.PubKey())
	require.Equal(t, sig.Sequence, uint64(100))

	v, r, s = signedTx.RawSignatureValues()
	sigData = sig.Data.(*signing.SingleSignatureData)
	require.Equal(t, sigData.SignMode, keeper.SignMode_SIGN_MODE_ETHEREUM)

	sigBytes = make([]byte, 65)
	copy(sigBytes[32-len(r.Bytes()):32], r.Bytes())
	copy(sigBytes[64-len(s.Bytes()):64], s.Bytes())
	sigBytes[64] = byte(v.Uint64())

	require.Equal(t, sigData.Signature, sigBytes)

	// Convert back to ethereum tx
	ethTx2, _, err = keeper.NewTxUtils(&input.EVMKeeper).ConvertCosmosTxToEthereumTx(ctx, sdkTx)
	require.NoError(t, err)

	// Verify the signature to check sender is correct
	sender, err := signer.Sender(ethTx2)
	require.NoError(t, err)
	require.Equal(t, common.BytesToAddress(addrBz), sender)
	EqualEthTransaction(t, signedTx, ethTx2)

	// 3. Set without unlimited sender
	// Set the gas enforcement parameters without unlimited sender
	params, err = input.EVMKeeper.Params.Get(ctx)
	require.NoError(t, err)
	gasEnforcement.UnlimitedGasSenders = []string{}
	params.GasEnforcement = gasEnforcement
	err = input.EVMKeeper.Params.Set(ctx, params)
	require.NoError(t, err)

	// feeAmount will be quarter of the original fee amount
	expectedFeeAmount = sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, math.NewIntFromBigInt(feeAmount.Div(feeAmount, big.NewInt(4))).AddRaw(1)))

	// Convert to cosmos tx
	// Since the sender is in the unlimited gas senders, the gas limit will be set to maxGasLimit
	sdkTx, err = keeper.NewTxUtils(&input.EVMKeeper).ConvertEthereumTxToCosmosTx(ctx, signedTx)
	require.NoError(t, err)

	msgs = sdkTx.GetMsgs()
	require.Len(t, msgs, 1)
	msg, ok = msgs[0].(*types.MsgCall)
	require.True(t, ok)

	expectedMsg.Sender = sdk.AccAddress(addrBz).String()
	require.Equal(t, msg, expectedMsg)

	authTx = sdkTx.(authsigning.Tx)
	require.Equal(t, authTx.GetGas(), maxGasLimit)
	require.Equal(t, authTx.GetFee(), expectedFeeAmount)

	sigs, err = authTx.GetSignaturesV2()
	require.NoError(t, err)
	require.Len(t, sigs, 1)

	sig = sigs[0]
	require.Equal(t, sig.PubKey, cosmosKey.PubKey())
	require.Equal(t, sig.Sequence, uint64(100))

	v, r, s = signedTx.RawSignatureValues()
	sigData = sig.Data.(*signing.SingleSignatureData)
	require.Equal(t, sigData.SignMode, keeper.SignMode_SIGN_MODE_ETHEREUM)

	sigBytes = make([]byte, 65)
	copy(sigBytes[32-len(r.Bytes()):32], r.Bytes())
	copy(sigBytes[64-len(s.Bytes()):64], s.Bytes())
	sigBytes[64] = byte(v.Uint64())

	require.Equal(t, sigData.Signature, sigBytes)

	// Convert back to ethereum tx
	ethTx2, _, err = keeper.NewTxUtils(&input.EVMKeeper).ConvertCosmosTxToEthereumTx(ctx, sdkTx)
	require.NoError(t, err)

	// Verify the signature to check sender is correct
	sender, err = signer.Sender(ethTx2)
	require.NoError(t, err)
	require.Equal(t, common.BytesToAddress(addrBz), sender)
	EqualEthTransaction(t, signedTx, ethTx2)
}

func Test_AccessTxConversion(t *testing.T) {
	ctx, input := createDefaultTestInput(t)

	_, _, addr := keyPubAddr()
	input.Faucet.Mint(ctx, addr, sdk.NewCoin(sdk.DefaultBondDenom, math.NewInt(1000000000000000000)))

	decimals := input.Decimals
	gasLimit := uint64(1_000_000)
	feeAmount := new(big.Int).Mul(
		big.NewInt(int64(gasLimit)),
		new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(decimals)-8), nil), // gas price is 1e-8
	)

	ethFactoryAddr, err := input.EVMKeeper.GetERC20FactoryAddr(ctx)
	require.NoError(t, err)

	abi, err := erc20_factory.Erc20FactoryMetaData.GetAbi()
	require.NoError(t, err)

	inputBz, err := abi.Pack("createERC20", "bar", "bar", uint8(6))
	require.NoError(t, err)

	value := types.ToEthersUnit(decimals, big.NewInt(100))

	gasFeeCap := types.ToEthersUnit(decimals, feeAmount)
	gasFeeCap = gasFeeCap.Quo(gasFeeCap, new(big.Int).SetUint64(gasLimit))

	ethChainID := types.ConvertCosmosChainIDToEthereumChainID(ctx.ChainID())
	// 1. Test with non AccessList but type is AccessListTx
	ethTx := coretypes.NewTx(&coretypes.AccessListTx{
		ChainID:    types.ConvertCosmosChainIDToEthereumChainID(ctx.ChainID()),
		Nonce:      100,
		GasPrice:   gasFeeCap,
		Gas:        gasLimit,
		To:         &ethFactoryAddr,
		Data:       inputBz,
		Value:      value,
		AccessList: nil,
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
		AccessList:   nil,
	})

	authTx := sdkTx.(authsigning.Tx)
	require.Equal(t, authTx.GetFee(), sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, math.NewIntFromBigInt(feeAmount).AddRaw(1))))

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
	sigBytes[64] = byte(v.Uint64())

	require.Equal(t, sigData.Signature, sigBytes)

	// Convert back to ethereum tx
	ethTx2, _, err := keeper.NewTxUtils(&input.EVMKeeper).ConvertCosmosTxToEthereumTx(ctx, sdkTx)
	require.NoError(t, err)
	EqualEthTransaction(t, signedTx, ethTx2)

	// 2. Test Normal Case
	ethTx = coretypes.NewTx(&coretypes.AccessListTx{
		ChainID:  types.ConvertCosmosChainIDToEthereumChainID(ctx.ChainID()),
		Nonce:    100,
		GasPrice: gasFeeCap,
		Gas:      gasLimit,
		To:       &ethFactoryAddr,
		Data:     inputBz,
		Value:    value,
		AccessList: coretypes.AccessList{
			coretypes.AccessTuple{Address: ethFactoryAddr,
				StorageKeys: []common.Hash{
					common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
					common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000001"),
					common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000002"),
				}},
		},
	})

	randBytes = make([]byte, 64)
	_, err = rand.Read(randBytes)
	require.NoError(t, err)
	reader = bytes.NewReader(randBytes)

	privKey, err = ecdsa.GenerateKey(crypto.S256(), reader)
	require.NoError(t, err)
	signedTx, err = coretypes.SignTx(ethTx, signer, privKey)
	require.NoError(t, err)

	cosmosKey = ethsecp256k1.PrivKey{
		Key: crypto.FromECDSA(privKey),
	}
	addrBz = cosmosKey.PubKey().Address()

	// Convert to cosmos tx
	sdkTx, err = keeper.NewTxUtils(&input.EVMKeeper).ConvertEthereumTxToCosmosTx(ctx, signedTx)
	require.NoError(t, err)

	msgs = sdkTx.GetMsgs()
	require.Len(t, msgs, 1)
	msg, ok = msgs[0].(*types.MsgCall)
	require.True(t, ok)
	require.Equal(t, msg, &types.MsgCall{
		Sender:       sdk.AccAddress(addrBz).String(),
		ContractAddr: ethFactoryAddr.Hex(),
		Input:        hexutil.Encode(inputBz),
		Value:        math.NewInt(100),
		AccessList: []types.AccessTuple{
			{
				Address: ethFactoryAddr.String(),
				StorageKeys: []string{
					common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000").Hex(),
					common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000001").Hex(),
					common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000002").Hex()},
			},
		},
	})

	authTx = sdkTx.(authsigning.Tx)
	expectedFeeAmount := sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, math.NewIntFromBigInt(feeAmount).AddRaw(1)))
	require.Equal(t, authTx.GetFee(), expectedFeeAmount)

	sigs, err = authTx.GetSignaturesV2()
	require.NoError(t, err)
	require.Len(t, sigs, 1)

	sig = sigs[0]
	require.Equal(t, sig.PubKey, cosmosKey.PubKey())
	require.Equal(t, sig.Sequence, uint64(100))

	v, r, s = signedTx.RawSignatureValues()
	sigData = sig.Data.(*signing.SingleSignatureData)
	require.Equal(t, sigData.SignMode, keeper.SignMode_SIGN_MODE_ETHEREUM)

	sigBytes = make([]byte, 65)
	copy(sigBytes[32-len(r.Bytes()):32], r.Bytes())
	copy(sigBytes[64-len(s.Bytes()):64], s.Bytes())
	sigBytes[64] = byte(v.Uint64())

	require.Equal(t, sigData.Signature, sigBytes)

	// Convert back to ethereum tx
	ethTx2, _, err = keeper.NewTxUtils(&input.EVMKeeper).ConvertCosmosTxToEthereumTx(ctx, sdkTx)
	require.NoError(t, err)
	EqualEthTransaction(t, signedTx, ethTx2)

	// manipulate the fee amount
	txBuilder := sdkTx.(client.TxBuilder)
	txBuilder.SetFeeAmount(expectedFeeAmount.Add(sdk.NewCoin(sdk.DefaultBondDenom, math.NewInt(1))))
	_, _, err = keeper.NewTxUtils(&input.EVMKeeper).ConvertCosmosTxToEthereumTx(ctx, txBuilder.GetTx())
	require.ErrorIs(t, err, types.ErrTxConversionFailed)
}

func Test_LegacyTxConversion(t *testing.T) {
	ctx, input := createDefaultTestInput(t)

	_, _, addr := keyPubAddr()
	input.Faucet.Mint(ctx, addr, sdk.NewCoin(sdk.DefaultBondDenom, math.NewInt(1000000000000000000)))

	decimals := input.Decimals
	gasLimit := uint64(1_000_000)
	feeAmount := new(big.Int).Mul(
		big.NewInt(int64(gasLimit)),
		new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(decimals)-8), nil), // gas price is 1e-8
	)

	ethFactoryAddr, err := input.EVMKeeper.GetERC20FactoryAddr(ctx)
	require.NoError(t, err)

	abi, err := erc20_factory.Erc20FactoryMetaData.GetAbi()
	require.NoError(t, err)

	inputBz, err := abi.Pack("createERC20", "bar", "bar", uint8(6))
	require.NoError(t, err)

	gasFeeCap := types.ToEthersUnit(decimals, feeAmount)
	gasFeeCap = gasFeeCap.Quo(gasFeeCap, new(big.Int).SetUint64(gasLimit))
	value := types.ToEthersUnit(decimals, big.NewInt(100))

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
	expectedMsg := &types.MsgCall{
		Sender:       sdk.AccAddress(addrBz).String(),
		ContractAddr: ethFactoryAddr.Hex(),
		Input:        hexutil.Encode(inputBz),
		Value:        math.NewInt(100),
	}
	require.Equal(t, msg, expectedMsg)

	authTx := sdkTx.(authsigning.Tx)
	expectedFeeAmount := sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, math.NewIntFromBigInt(feeAmount).AddRaw(1)))
	require.Equal(t, authTx.GetFee(), expectedFeeAmount)

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

	// manipulate the fee amount
	txBuilder := sdkTx.(client.TxBuilder)
	txBuilder.SetFeeAmount(expectedFeeAmount.Add(sdk.NewCoin(sdk.DefaultBondDenom, math.NewInt(1))))
	_, _, err = keeper.NewTxUtils(&input.EVMKeeper).ConvertCosmosTxToEthereumTx(ctx, txBuilder.GetTx())
	require.ErrorIs(t, err, types.ErrTxConversionFailed)

	// 2. Set the max params of gas configuration
	maxGasLimit := gasLimit / 2
	params, err := input.EVMKeeper.Params.Get(ctx)
	require.NoError(t, err)
	maxGasFeeCap := math.NewIntFromBigInt(gasFeeCap.Div(gasFeeCap, big.NewInt(2)))

	gasEnforcement := &types.GasEnforcement{
		MaxGasLimit:  maxGasLimit,
		MaxGasFeeCap: &maxGasFeeCap,
	}
	params.GasEnforcement = gasEnforcement
	err = input.EVMKeeper.Params.Set(ctx, params)
	require.NoError(t, err)

	// Convert to cosmos tx
	sdkTx, err = keeper.NewTxUtils(&input.EVMKeeper).ConvertEthereumTxToCosmosTx(ctx, signedTx)
	require.NoError(t, err)

	msgs = sdkTx.GetMsgs()
	require.Len(t, msgs, 1)
	msg, ok = msgs[0].(*types.MsgCall)
	require.True(t, ok)
	require.Equal(t, msg, expectedMsg)

	authTx = sdkTx.(authsigning.Tx)
	expectedFeeAmount = sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, math.NewIntFromBigInt(feeAmount.Div(feeAmount, big.NewInt(4))).AddRaw(1)))
	require.Equal(t, authTx.GetFee(), expectedFeeAmount)
	sigs, err = authTx.GetSignaturesV2()
	require.NoError(t, err)
	require.Len(t, sigs, 1)

	sig = sigs[0]
	require.Equal(t, sig.PubKey, cosmosKey.PubKey())
	require.Equal(t, sig.Sequence, uint64(100))

	v, r, s = signedTx.RawSignatureValues()
	sigData = sig.Data.(*signing.SingleSignatureData)
	require.Equal(t, sigData.SignMode, keeper.SignMode_SIGN_MODE_ETHEREUM)

	sigBytes = make([]byte, 65)
	copy(sigBytes[32-len(r.Bytes()):32], r.Bytes())
	copy(sigBytes[64-len(s.Bytes()):64], s.Bytes())
	sigBytes[64] = byte(v.Uint64() - (35 + ethChainID.Uint64()*2))

	require.Equal(t, sigData.Signature, sigBytes)

	// Convert back to ethereum tx
	ethTx2, _, err = keeper.NewTxUtils(&input.EVMKeeper).ConvertCosmosTxToEthereumTx(ctx, sdkTx)
	require.NoError(t, err)

	// Verify the signature to check sender is correct
	sender, err := signer.Sender(ethTx2)
	require.NoError(t, err)
	require.Equal(t, common.BytesToAddress(addrBz), sender)
	EqualEthTransaction(t, signedTx, ethTx2)

}

func Test_IsEthereumTx(t *testing.T) {
	ctx, input := createDefaultTestInput(t)

	ok := keeper.NewTxUtils(&input.EVMKeeper).IsEthereumTx(ctx)
	require.False(t, ok)

	ctx = ctx.WithValue(types.CONTEXT_KEY_ETH_TX, true)
	ok = keeper.NewTxUtils(&input.EVMKeeper).IsEthereumTx(ctx)
	require.True(t, ok)
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
