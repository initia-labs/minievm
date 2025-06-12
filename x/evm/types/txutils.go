package types

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/big"

	"cosmossdk.io/core/address"
	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	coretypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/initia-labs/initia/crypto/ethsecp256k1"
)

// SignMode_SIGN_MODE_ETHEREUM is the sign mode for Ethereum transactions.
const SignMode_SIGN_MODE_ETHEREUM = signing.SignMode(9999)

// txMetadata is the metadata of a Cosmos SDK transaction.
type txMetadata struct {
	Type      uint8  `json:"type"`
	GasFeeCap string `json:"gas_fee_cap"`
	GasTipCap string `json:"gas_tip_cap"`
}

// LazyArgsGetterForConvertEthereumTxToCosmosTx is a function that returns the arguments for ConvertEthereumTxToCosmosTx.
// use lazy args getter to avoid unnecessary params and decimals fetching
type LazyArgsGetterForConvertEthereumTxToCosmosTx func() (chainID string, ac address.Codec, cdc codec.Codec, params Params, feeDecimals uint8, ethTx *coretypes.Transaction, err error)

// LazyArgsGetterForConvertCosmosTxToEthereumTx is a function that returns the arguments for ConvertCosmosTxToEthereumTx.
// use lazy args getter to avoid unnecessary params and decimals fetching
type LazyArgsGetterForConvertCosmosTxToEthereumTx func() (chainID string, ac address.Codec, params Params, feeDecimals uint8, sdkTx sdk.Tx, err error)

// ConvertEthereumTxToCosmosTx converts an Ethereum transaction to a Cosmos SDK transaction.
func ConvertEthereumTxToCosmosTx(lazyArgsGetter LazyArgsGetterForConvertEthereumTxToCosmosTx) (sdk.Tx, error) {
	chainID, ac, cdc, params, feeDecimals, ethTx, err := lazyArgsGetter()
	if err != nil {
		return nil, err
	}

	gasFeeCap := ethTx.GasFeeCap()
	if gasFeeCap == nil {
		gasFeeCap = big.NewInt(0)
	}
	gasTipCap := ethTx.GasTipCap()
	if gasTipCap == nil {
		gasTipCap = big.NewInt(0)
	}

	// convert gas fee unit from wei to cosmos fee unit
	gasLimit := ethTx.Gas()
	gasFeeAmount := computeGasFeeAmount(gasFeeCap, gasLimit, feeDecimals)
	feeAmount := sdk.NewCoins(sdk.NewCoin(params.FeeDenom, math.NewIntFromBigInt(gasFeeAmount)))

	// convert value unit from wei to cosmos fee unit
	value := FromEthersUnit(feeDecimals, ethTx.Value())

	// check if the value is correctly converted without dropping any precision
	if ToEthersUnit(feeDecimals, value).Cmp(ethTx.Value()) != 0 {
		return nil, ErrInvalidValue.Wrap("failed to convert value to token unit without dropping precision")
	}

	// signer
	ethChainID := ConvertCosmosChainIDToEthereumChainID(chainID)
	signer := coretypes.LatestSignerForChainID(ethChainID)

	// get tx sender
	ethSender, err := coretypes.Sender(signer, ethTx)
	if err != nil {
		return nil, err
	}
	// sig bytes
	v, r, s := ethTx.RawSignatureValues()
	sigBytes := make([]byte, 65)
	switch ethTx.Type() {
	case coretypes.LegacyTxType:
		sigBytes[64] = byte(new(big.Int).Sub(v, new(big.Int).Add(new(big.Int).Add(ethChainID, ethChainID), big.NewInt(35))).Uint64())
	case coretypes.AccessListTxType, coretypes.DynamicFeeTxType:
		sigBytes[64] = byte(v.Uint64())
	default:
		return nil, sdkerrors.ErrorInvalidSigner.Wrapf("unsupported tx type: %d", ethTx.Type())
	}

	copy(sigBytes[32-len(r.Bytes()):32], r.Bytes())
	copy(sigBytes[64-len(s.Bytes()):64], s.Bytes())

	sigData := &signing.SingleSignatureData{
		SignMode:  SignMode_SIGN_MODE_ETHEREUM,
		Signature: sigBytes,
	}

	// recover pubkey
	pubKeyBz, err := crypto.Ecrecover(signer.Hash(ethTx).Bytes(), sigBytes)
	if err != nil {
		return nil, sdkerrors.ErrorInvalidSigner.Wrapf("failed to recover pubkey: %v", err.Error())
	}

	// compress pubkey
	compressedPubKey, err := ethsecp256k1.NewPubKeyFromBytes(pubKeyBz)
	if err != nil {
		return nil, sdkerrors.ErrorInvalidSigner.Wrapf("failed to create pubkey: %v", err.Error())
	}

	// construct signature
	sig := signing.SignatureV2{
		PubKey:   compressedPubKey,
		Data:     sigData,
		Sequence: ethTx.Nonce(),
	}

	// convert sender to string
	sender, err := ac.BytesToString(ethSender.Bytes())
	if err != nil {
		return nil, err
	}

	// convert access list
	accessList := ConvertEthAccessListToCosmos(ethTx.AccessList())

	sdkMsgs := []sdk.Msg{}
	if ethTx.To() == nil {
		sdkMsgs = append(sdkMsgs, &MsgCreate{
			Sender:     sender,
			Code:       hexutil.Encode(ethTx.Data()),
			Value:      math.NewIntFromBigInt(value),
			AccessList: accessList,
		})
	} else {
		sdkMsgs = append(sdkMsgs, &MsgCall{
			Sender:       sender,
			ContractAddr: ethTx.To().String(),
			Input:        hexutil.Encode(ethTx.Data()),
			Value:        math.NewIntFromBigInt(value),
			AccessList:   accessList,
		})
	}

	txBuilder := authtx.NewTxConfig(cdc, authtx.DefaultSignModes).NewTxBuilder()
	if err = txBuilder.SetMsgs(sdkMsgs...); err != nil {
		return nil, err
	}
	txBuilder.SetFeeAmount(feeAmount)
	txBuilder.SetGasLimit(gasLimit)
	if err = txBuilder.SetSignatures(sig); err != nil {
		return nil, err
	}

	// set memo
	memo, err := json.Marshal(txMetadata{
		Type:      ethTx.Type(),
		GasFeeCap: gasFeeCap.String(),
		GasTipCap: gasTipCap.String(),
	})
	if err != nil {
		return nil, err
	}
	txBuilder.SetMemo(string(memo))

	return txBuilder.GetTx(), nil
}

// ConvertCosmosTxToEthereumTx converts a Cosmos SDK transaction to an Ethereum transaction.
// It returns nil if the transaction is not an EVM transaction.
func ConvertCosmosTxToEthereumTx(lazyArgsGetter LazyArgsGetterForConvertCosmosTxToEthereumTx) (*coretypes.Transaction, *common.Address, error) {
	chainID, ac, params, feeDecimals, sdkTx, err := lazyArgsGetter()
	if err != nil {
		return nil, nil, err
	}

	msgs := sdkTx.GetMsgs()
	if len(msgs) != 1 {
		return nil, nil, nil
	}

	msg := msgs[0]
	typeUrl := sdk.MsgTypeURL(msg)
	if typeUrl != "/minievm.evm.v1.MsgCall" && typeUrl != "/minievm.evm.v1.MsgCreate" {
		return nil, nil, nil
	}

	authTx := sdkTx.(authsigning.Tx)
	memo := authTx.GetMemo()
	if len(memo) == 0 {
		return nil, nil, nil
	}
	md := txMetadata{}
	decoder := json.NewDecoder(bytes.NewReader([]byte(memo)))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&md); err != nil {
		return nil, nil, nil
	}

	sigs, err := authTx.GetSignaturesV2()
	if err != nil {
		return nil, nil, err
	}
	if len(sigs) != 1 {
		return nil, nil, nil
	}

	feeAmount := authTx.GetFee()
	if !(len(feeAmount) == 0 || (len(feeAmount) == 1 && feeAmount[0].Denom == params.FeeDenom)) {
		return nil, nil, nil
	}

	var tx *coretypes.Transaction

	sig := sigs[0]
	cosmosSender := sig.PubKey.Address()
	if len(cosmosSender.Bytes()) != common.AddressLength {
		return nil, nil, nil
	}

	sender := common.BytesToAddress(sig.PubKey.Address())
	sigData, ok := sig.Data.(*signing.SingleSignatureData)
	if !ok {
		return nil, nil, nil
	}

	// filter out non-EVM transactions
	if sigData.SignMode != SignMode_SIGN_MODE_ETHEREUM {
		return nil, nil, nil
	}

	var v, r, s []byte
	if len(sigData.Signature) == 65 {
		v, r, s = sigData.Signature[64:], sigData.Signature[:32], sigData.Signature[32:64]
	} else if len(sigData.Signature) == 64 {
		v, r, s = []byte{}, sigData.Signature[:32], sigData.Signature[32:64]
	} else {
		return nil, nil, ErrTxConversionFailed.Wrap("invalid signature length")
	}

	gasFeeCap, ok := new(big.Int).SetString(md.GasFeeCap, 10)
	if !ok {
		return nil, nil, err
	}

	gasTipCap, ok := new(big.Int).SetString(md.GasTipCap, 10)
	if !ok {
		return nil, nil, err
	}

	// check if the fee amount is correctly converted
	gas := authTx.GetGas()
	computedFeeAmount := sdk.NewCoins(sdk.NewCoin(params.FeeDenom, math.NewIntFromBigInt(computeGasFeeAmount(gasFeeCap, gas, feeDecimals))))
	if !feeAmount.Equal(computedFeeAmount) {
		return nil, nil, ErrTxConversionFailed.Wrap("fee amount manipulation detected")
	}

	var to *common.Address
	var input []byte
	var value *big.Int
	var accessList coretypes.AccessList
	switch typeUrl {
	case "/minievm.evm.v1.MsgCall":
		callMsg := msg.(*MsgCall)
		contractAddr, err := ContractAddressFromString(ac, callMsg.ContractAddr)
		if err != nil {
			return nil, nil, err
		}

		data, err := hexutil.Decode(callMsg.Input)
		if err != nil {
			return nil, nil, err
		}

		to = &contractAddr
		input = data
		// When ethereum tx is converted into cosmos tx by ConvertEthereumTxToCosmosTx,
		// the value is converted to cosmos fee unit from wei.
		// So we need to convert it back to wei to get original ethereum tx and verify signature.
		value = ToEthersUnit(feeDecimals, callMsg.Value.BigInt())
		accessList = ConvertCosmosAccessListToEth(callMsg.AccessList)

	case "/minievm.evm.v1.MsgCreate":
		createMsg := msg.(*MsgCreate)
		data, err := hexutil.Decode(createMsg.Code)
		if err != nil {
			return nil, nil, err
		}

		to = nil
		input = data
		// Same as above (MsgCall)
		value = ToEthersUnit(feeDecimals, createMsg.Value.BigInt())
		accessList = ConvertCosmosAccessListToEth(createMsg.AccessList)
	}

	ethChainID := ConvertCosmosChainIDToEthereumChainID(chainID)

	var txData coretypes.TxData
	switch md.Type {
	case coretypes.LegacyTxType:
		txData = &coretypes.LegacyTx{
			Nonce:    sig.Sequence,
			Gas:      gas,
			To:       to,
			Data:     input,
			GasPrice: gasFeeCap,
			Value:    value,
			R:        new(big.Int).SetBytes(r),
			S:        new(big.Int).SetBytes(s),
			V:        new(big.Int).Add(new(big.Int).SetBytes(v), new(big.Int).SetUint64(35+ethChainID.Uint64()*2)),
		}
	case coretypes.AccessListTxType:
		txData = &coretypes.AccessListTx{
			ChainID:    ethChainID,
			Nonce:      sig.Sequence,
			GasPrice:   gasFeeCap,
			Gas:        gas,
			Value:      value,
			To:         to,
			Data:       input,
			AccessList: accessList,
			R:          new(big.Int).SetBytes(r),
			S:          new(big.Int).SetBytes(s),
			V:          new(big.Int).SetBytes(v),
		}

	case coretypes.DynamicFeeTxType:
		txData = &coretypes.DynamicFeeTx{
			ChainID:    ethChainID,
			Nonce:      sig.Sequence,
			GasTipCap:  gasTipCap,
			GasFeeCap:  gasFeeCap,
			Gas:        gas,
			To:         to,
			Data:       input,
			Value:      value,
			AccessList: accessList,
			R:          new(big.Int).SetBytes(r),
			S:          new(big.Int).SetBytes(s),
			V:          new(big.Int).SetBytes(v),
		}

	default:
		return nil, nil, fmt.Errorf("unsupported tx type: %d", md.Type)
	}

	tx = coretypes.NewTx(txData)

	return tx, &sender, nil
}

func computeGasFeeAmount(gasFeeCap *big.Int, gas uint64, decimals uint8) *big.Int {
	if gasFeeCap.Cmp(big.NewInt(0)) == 0 {
		return big.NewInt(0)
	}

	gasFeeCap = new(big.Int).Mul(gasFeeCap, new(big.Int).SetUint64(gas))
	gasFeeAmount := FromEthersUnit(decimals, gasFeeCap)

	// add 1 to the gas fee amount to avoid rounding errors
	return new(big.Int).Add(gasFeeAmount, big.NewInt(1))
}

// ConvertCosmosAccessListToEth converts a Cosmos SDK access list to an Ethereum access list.
func ConvertCosmosAccessListToEth(cosmosAccessList []AccessTuple) coretypes.AccessList {
	if len(cosmosAccessList) == 0 {
		return nil
	}
	coreAccessList := make(coretypes.AccessList, len(cosmosAccessList))
	for i, a := range cosmosAccessList {
		storageKeys := make([]common.Hash, len(a.StorageKeys))
		for j, s := range a.StorageKeys {
			storageKeys[j] = common.HexToHash(s)
		}
		coreAccessList[i] = coretypes.AccessTuple{
			Address:     common.HexToAddress(a.Address),
			StorageKeys: storageKeys,
		}
	}
	return coreAccessList
}

// ConvertEthAccessListToCosmos converts an Ethereum access list to a Cosmos SDK access list.
func ConvertEthAccessListToCosmos(ethAccessList coretypes.AccessList) []AccessTuple {
	if len(ethAccessList) == 0 {
		return nil
	}
	accessList := make([]AccessTuple, len(ethAccessList))
	for i, al := range ethAccessList {
		storageKeys := make([]string, len(al.StorageKeys))
		for j, s := range al.StorageKeys {
			storageKeys[j] = s.String()
		}
		accessList[i] = AccessTuple{
			Address:     al.Address.String(),
			StorageKeys: storageKeys,
		}
	}
	return accessList
}
