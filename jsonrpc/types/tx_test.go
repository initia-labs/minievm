package types_test

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	coretypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	types "github.com/initia-labs/minievm/jsonrpc/types"
	"github.com/stretchr/testify/require"
)

func TestLegacyTxTypeRPCTransaction(t *testing.T) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		t.Fatalf("Failed to generate private key: %v", err)
	}

	toAddress := crypto.PubkeyToAddress(privateKey.PublicKey)
	chainID := big.NewInt(1)
	tx := coretypes.NewTx(&coretypes.LegacyTx{
		Nonce:    0,
		GasPrice: big.NewInt(1000),
		Gas:      1000,
		To:       &toAddress,
		Value:    big.NewInt(100),
		Data:     []byte{},
		V:        nil,
		R:        nil,
		S:        nil,
	})

	signedTx, err := coretypes.SignTx(tx, coretypes.NewCancunSigner(chainID), privateKey)
	if err != nil {
		t.Fatalf("Failed to sign transaction: %v", err)
	}

	rpcTx := types.NewRPCTransaction(signedTx, common.Hash{}, 0, 0, chainID)
	ethTx := rpcTx.ToTransaction()

	err = matchTx(signedTx, ethTx)
	require.NoError(t, err)

}

func TestAccessListTypeRPCTransaction(t *testing.T) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		t.Fatalf("Failed to generate private key: %v", err)
	}

	toAddress := crypto.PubkeyToAddress(privateKey.PublicKey)
	chainID := big.NewInt(1)
	tx := coretypes.NewTx(&coretypes.AccessListTx{
		ChainID:    chainID,
		Nonce:      0,
		GasPrice:   big.NewInt(1000),
		Gas:        1000,
		To:         &toAddress,
		Value:      big.NewInt(100),
		Data:       []byte{},
		AccessList: []coretypes.AccessTuple{},
		V:          nil,
		R:          nil,
		S:          nil,
	})

	signedTx, err := coretypes.SignTx(tx, coretypes.NewCancunSigner(chainID), privateKey)
	if err != nil {
		t.Fatalf("Failed to sign transaction: %v", err)
	}
	rpcTx := types.NewRPCTransaction(signedTx, common.Hash{}, 0, 0, chainID)
	ethTx := rpcTx.ToTransaction()
	ethTx.Hash()
	err = matchTx(signedTx, ethTx)
	require.NoError(t, err)
}

func TestDynamicFeeTxTypeRPCTransaction(t *testing.T) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		t.Fatalf("Failed to generate private key: %v", err)
	}

	toAddress := crypto.PubkeyToAddress(privateKey.PublicKey)
	chainID := big.NewInt(1)
	tx := coretypes.NewTx(&coretypes.DynamicFeeTx{
		ChainID:    chainID,
		Nonce:      0,
		GasTipCap:  big.NewInt(20),
		GasFeeCap:  big.NewInt(100),
		Gas:        21000,
		To:         &toAddress,
		Value:      big.NewInt(1000),
		Data:       []byte{},
		AccessList: []coretypes.AccessTuple{},
		V:          nil,
		R:          nil,
		S:          nil,
	})

	signedTx, err := coretypes.SignTx(tx, coretypes.NewCancunSigner(chainID), privateKey)
	if err != nil {
		t.Fatalf("Failed to sign transaction: %v", err)
	}
	rpcTx := types.NewRPCTransaction(signedTx, common.Hash{}, 0, 0, chainID)
	ethTx := rpcTx.ToTransaction()

	err = matchTx(signedTx, ethTx)
	require.NoError(t, err)
}

func matchTx(signedTx *coretypes.Transaction, ethTx *coretypes.Transaction) error {
	if signedTx.Type() != ethTx.Type() {
		return fmt.Errorf("Expected transaction type %v, got %v", signedTx.Type(), ethTx.Type())
	}

	if signedTx.Hash() != ethTx.Hash() {
		return fmt.Errorf("Expected hash %v, got %v", signedTx.Hash(), ethTx.Hash())
	}

	if signedTx.Nonce() != ethTx.Nonce() {
		return fmt.Errorf("Expected nonce %v, got %v", signedTx.Nonce(), ethTx.Nonce())
	}

	if signedTx.Gas() != ethTx.Gas() {
		return fmt.Errorf("Expected gas %v, got %v", signedTx.Gas(), ethTx.Gas())
	}

	if signedTx.GasFeeCapCmp(ethTx) != 0 {
		return fmt.Errorf("Expected gas price %v, got %v", signedTx.GasPrice(), ethTx.GasPrice())
	}

	if signedTx.GasFeeCapCmp(ethTx) != 0 {
		return fmt.Errorf("Expected gas fee cap %v, got %v", signedTx.GasFeeCap(), ethTx.GasFeeCap())
	}

	if signedTx.GasTipCapCmp(ethTx) != 0 {
		return fmt.Errorf("Expected gas tip cap %v, got %v", signedTx.GasTipCap(), ethTx.GasTipCap())
	}

	if signedTx.Value().Cmp(ethTx.Value()) != 0 {
		return fmt.Errorf("Expected value %v, got %v", signedTx.Value(), ethTx.Value())
	}

	if signedTx.To() == nil || ethTx.To() == nil || *signedTx.To() != *ethTx.To() {
		return fmt.Errorf("Expected to address %v, got %v", signedTx.To(), ethTx.To())
	}
	signedTxAccesList := signedTx.AccessList()
	ethTxAccessList := ethTx.AccessList()
	if len(signedTxAccesList) != len(ethTxAccessList) {
		return fmt.Errorf("Expected access list length %v, got %v", len(signedTxAccesList), len(ethTxAccessList))
	}
	for i := range signedTxAccesList {
		if signedTxAccesList[i].Address != ethTxAccessList[i].Address || len(signedTxAccesList[i].StorageKeys) != len(ethTxAccessList[i].StorageKeys) {
			return fmt.Errorf("Expected access list %v, got %v", signedTx.AccessList(), ethTxAccessList)
		}
		for j := range signedTxAccesList[i].StorageKeys {
			if signedTxAccesList[i].StorageKeys[j] != ethTxAccessList[i].StorageKeys[j] {
				return fmt.Errorf("Expected access list %v, got %v", signedTx.AccessList(), ethTxAccessList)
			}
		}
	}
	if string(signedTx.Data()) != string(ethTx.Data()) {
		return fmt.Errorf("Expected data %v, got %v", signedTx.Data(), ethTx.Data())
	}

	return nil
}
