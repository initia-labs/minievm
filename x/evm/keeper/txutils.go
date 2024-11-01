package keeper

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"math/big"

	"cosmossdk.io/math"
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
	"github.com/initia-labs/minievm/x/evm/types"
)

const SignMode_SIGN_MODE_ETHEREUM = signing.SignMode(9999)

type TxUtils struct {
	*Keeper
}

func NewTxUtils(k *Keeper) *TxUtils {
	return &TxUtils{
		Keeper: k,
	}
}

func computeGasFeeAmount(gasFeeCap *big.Int, gas uint64, decimals uint8) *big.Int {
	if gasFeeCap.Cmp(big.NewInt(0)) == 0 {
		return big.NewInt(0)
	}

	gasFeeCap = new(big.Int).Mul(gasFeeCap, new(big.Int).SetUint64(gas))
	gasFeeAmount := types.FromEthersUnit(decimals, gasFeeCap)

	// add 1 to the gas fee amount to avoid rounding errors
	return new(big.Int).Add(gasFeeAmount, big.NewInt(1))
}

func (u *TxUtils) ConvertEthereumTxToCosmosTx(ctx context.Context, ethTx *coretypes.Transaction) (sdk.Tx, error) {
	params, err := u.Params.Get(ctx)
	if err != nil {
		return nil, err
	}

	feeDenom := params.FeeDenom
	decimals, err := u.ERC20Keeper().GetDecimals(ctx, feeDenom)
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
	gasFeeAmount := computeGasFeeAmount(gasFeeCap, gasLimit, decimals)
	feeAmount := sdk.NewCoins(sdk.NewCoin(params.FeeDenom, math.NewIntFromBigInt(gasFeeAmount)))

	// convert value unit from wei to cosmos fee unit
	value := types.FromEthersUnit(decimals, ethTx.Value())

	// check if the value is correctly converted without dropping any precision
	if types.ToEthersUint(decimals, value).Cmp(ethTx.Value()) != 0 {
		return nil, types.ErrInvalidValue.Wrap("failed to convert value to token unit without dropping precision")
	}

	// signer
	chainID := sdk.UnwrapSDKContext(ctx).ChainID()
	ethChainID := types.ConvertCosmosChainIDToEthereumChainID(chainID)
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
	sender, err := u.ac.BytesToString(ethSender.Bytes())
	if err != nil {
		return nil, err
	}

	// convert access list
	accessList := types.ConvertEthAccessListToCosmos(ethTx.AccessList())

	sdkMsgs := []sdk.Msg{}
	if ethTx.To() == nil {
		sdkMsgs = append(sdkMsgs, &types.MsgCreate{
			Sender:     sender,
			Code:       hexutil.Encode(ethTx.Data()),
			Value:      math.NewIntFromBigInt(value),
			AccessList: accessList,
		})
	} else {
		sdkMsgs = append(sdkMsgs, &types.MsgCall{
			Sender:       sender,
			ContractAddr: ethTx.To().String(),
			Input:        hexutil.Encode(ethTx.Data()),
			Value:        math.NewIntFromBigInt(value),
			AccessList:   accessList,
		})
	}

	txBuilder := authtx.NewTxConfig(u.cdc, authtx.DefaultSignModes).NewTxBuilder()
	if err = txBuilder.SetMsgs(sdkMsgs...); err != nil {
		return nil, err
	}
	txBuilder.SetFeeAmount(feeAmount)
	txBuilder.SetGasLimit(gasLimit)
	if err = txBuilder.SetSignatures(sig); err != nil {
		return nil, err
	}

	// set memo
	memo, err := json.Marshal(metadata{
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

type metadata struct {
	Type      uint8  `json:"type"`
	GasFeeCap string `json:"gas_fee_cap"`
	GasTipCap string `json:"gas_tip_cap"`
}

// ConvertCosmosTxToEthereumTx converts a Cosmos SDK transaction to an Ethereum transaction.
// It returns nil if the transaction is not an EVM transaction.
func (u *TxUtils) ConvertCosmosTxToEthereumTx(ctx context.Context, sdkTx sdk.Tx) (*coretypes.Transaction, *common.Address, error) {
	msgs := sdkTx.GetMsgs()
	if len(msgs) != 1 {
		return nil, nil, nil
	}

	authTx := sdkTx.(authsigning.Tx)
	memo := authTx.GetMemo()
	if len(memo) == 0 {
		return nil, nil, nil
	}
	md := metadata{}
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

	fees := authTx.GetFee()
	params, err := u.Params.Get(ctx)
	if err != nil {
		return nil, nil, err
	}
	decimals, err := u.ERC20Keeper().GetDecimals(ctx, params.FeeDenom)
	if err != nil {
		return nil, nil, err
	}

	if !(len(fees) == 0 || (len(fees) == 1 && fees[0].Denom == params.FeeDenom)) {
		return nil, nil, nil
	}

	var tx *coretypes.Transaction

	msg := msgs[0]
	gas := authTx.GetGas()
	typeUrl := sdk.MsgTypeURL(msg)

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
		return nil, nil, nil
	}

	gasFeeCap, ok := new(big.Int).SetString(md.GasFeeCap, 10)
	if !ok {
		return nil, nil, err
	}

	gasTipCap, ok := new(big.Int).SetString(md.GasTipCap, 10)
	if !ok {
		return nil, nil, err
	}

	var to *common.Address
	var input []byte
	var value *big.Int
	var accessList coretypes.AccessList
	switch typeUrl {
	case "/minievm.evm.v1.MsgCall":
		callMsg := msg.(*types.MsgCall)
		contractAddr, err := types.ContractAddressFromString(u.ac, callMsg.ContractAddr)
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
		value = types.ToEthersUint(decimals, callMsg.Value.BigInt())
		accessList = types.ConvertCosmosAccessListToEth(callMsg.AccessList)

	case "/minievm.evm.v1.MsgCreate":
		createMsg := msg.(*types.MsgCreate)
		data, err := hexutil.Decode(createMsg.Code)
		if err != nil {
			return nil, nil, err
		}

		to = nil
		input = data
		// Same as above (MsgCall)
		value = types.ToEthersUint(decimals, createMsg.Value.BigInt())
		accessList = types.ConvertCosmosAccessListToEth(createMsg.AccessList)
	case "/minievm.evm.v1.MsgCreate2":
		// create2 is not supported
		return nil, nil, nil
	}

	chainID := sdk.UnwrapSDKContext(ctx).ChainID()
	ethChainID := types.ConvertCosmosChainIDToEthereumChainID(chainID)

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

// IsEthereumTx checks if the given Cosmos SDK transaction is an Ethereum transaction.
func (u *TxUtils) IsEthereumTx(ctx context.Context, sdkTx sdk.Tx) (bool, error) {
	msgs := sdkTx.GetMsgs()
	if len(msgs) != 1 {
		return false, nil
	}

	authTx := sdkTx.(authsigning.Tx)
	memo := authTx.GetMemo()
	if len(memo) == 0 {
		return false, nil
	}
	md := metadata{}
	decoder := json.NewDecoder(bytes.NewReader([]byte(memo)))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&md); err != nil {
		return false, nil
	}

	sigs, err := authTx.GetSignaturesV2()
	if err != nil {
		return false, err
	}
	if len(sigs) != 1 {
		return false, nil
	}

	fees := authTx.GetFee()
	params, err := u.Params.Get(ctx)
	if err != nil {
		return false, err
	}

	if !(len(fees) == 0 || (len(fees) == 1 && fees[0].Denom == params.FeeDenom)) {
		return false, nil
	}

	msg := msgs[0]
	typeUrl := sdk.MsgTypeURL(msg)
	if typeUrl != "/minievm.evm.v1.MsgCall" && typeUrl != "/minievm.evm.v1.MsgCreate" {
		return false, nil
	}

	sig := sigs[0]
	cosmosSender := sig.PubKey.Address()
	if len(cosmosSender.Bytes()) != common.AddressLength {
		return false, nil
	}

	sigData, ok := sig.Data.(*signing.SingleSignatureData)
	if !ok {
		return false, nil
	}

	// filter out non-EVM transactions
	if sigData.SignMode != SignMode_SIGN_MODE_ETHEREUM {
		return false, nil
	}

	return true, nil
}
