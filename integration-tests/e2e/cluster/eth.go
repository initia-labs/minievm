package cluster

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"

	sdk "github.com/cosmos/cosmos-sdk/types"

	evmtypes "github.com/initia-labs/minievm/x/evm/types"
)

// EthAccount holds an Ethereum keypair generated for JSON-RPC benchmarks.
type EthAccount struct {
	Name       string
	PrivateKey *ecdsa.PrivateKey
	Address    common.Address
}

// SignedEthTx holds a pre-signed Ethereum transaction ready for JSON-RPC broadcast.
type SignedEthTx struct {
	Account string
	Nonce   uint64
	TxHash  string
	RawHex  string
}

// EthAccounts returns all generated Ethereum accounts.
func (c *Cluster) EthAccounts() []EthAccount {
	return c.ethAccounts
}

// EthAccountNames returns the names of all Ethereum accounts.
func (c *Cluster) EthAccountNames() []string {
	names := make([]string, len(c.ethAccounts))
	for i, a := range c.ethAccounts {
		names[i] = a.Name
	}

	return names
}

// EthAccountByName looks up an Ethereum account by name.
func (c *Cluster) EthAccountByName(name string) (EthAccount, bool) {
	for _, a := range c.ethAccounts {
		if a.Name == name {
			return a, true
		}
	}

	return EthAccount{}, false
}

// EthChainID returns the EVM chain ID derived from the Cosmos chain ID.
func (c *Cluster) EthChainID() *big.Int {
	return evmtypes.ConvertCosmosChainIDToEthereumChainID(c.opts.ChainID)
}

// ValidatorEthAddress returns the validator address converted to an Ethereum address.
func (c *Cluster) ValidatorEthAddress() (common.Address, error) {
	return bech32ToEthAddress(c.ValidatorAddress())
}

// generateEthAccounts creates Ethereum key pairs for JSON-RPC benchmarks.
func (c *Cluster) generateEthAccounts() error {
	c.ethAccounts = make([]EthAccount, 0, c.opts.AccountCount)
	for i := 1; i <= c.opts.AccountCount; i++ {
		key, err := ethcrypto.GenerateKey()
		if err != nil {
			return fmt.Errorf("generate eth key %d: %w", i, err)
		}
		c.ethAccounts = append(c.ethAccounts, EthAccount{
			Name:       fmt.Sprintf("eth%d", i),
			PrivateKey: key,
			Address:    ethcrypto.PubkeyToAddress(key.PublicKey),
		})
	}

	return nil
}

// addEthGenesisAccounts adds ETH accounts to genesis with the standard balance.
func (c *Cluster) addEthGenesisAccounts(ctx context.Context, baseHome string) error {
	for _, acct := range c.ethAccounts {
		bech32Addr, err := sdk.Bech32ifyAddressBytes("init", acct.Address.Bytes())
		if err != nil {
			return fmt.Errorf("bech32 encode eth account %s: %w", acct.Name, err)
		}
		if _, err := c.exec(ctx,
			"genesis", "add-genesis-account", bech32Addr, "1000000000000000GAS",
			"--home", baseHome,
		); err != nil {
			return fmt.Errorf("add eth genesis account %s: %w", acct.Name, err)
		}
	}

	return nil
}

// SignEthTransfer signs a native ETH value transfer transaction.
func (c *Cluster) SignEthTransfer(from EthAccount, to common.Address, value *big.Int, nonce, gasLimit uint64) (SignedEthTx, error) {
	chainID := c.EthChainID()
	signer := ethtypes.LatestSignerForChainID(chainID)

	tx := ethtypes.NewTx(&ethtypes.LegacyTx{
		Nonce:    nonce,
		To:       &to,
		Value:    value,
		Gas:      gasLimit,
		GasPrice: big.NewInt(0),
	})

	signedTx, err := ethtypes.SignTx(tx, signer, from.PrivateKey)
	if err != nil {
		return SignedEthTx{}, fmt.Errorf("sign eth transfer: %w", err)
	}

	rawBytes, err := signedTx.MarshalBinary()
	if err != nil {
		return SignedEthTx{}, fmt.Errorf("marshal eth tx: %w", err)
	}

	return SignedEthTx{
		Account: from.Name,
		Nonce:   nonce,
		TxHash:  signedTx.Hash().Hex(),
		RawHex:  "0x" + hex.EncodeToString(rawBytes),
	}, nil
}

// SignEthContractCall signs an EVM contract call transaction.
func (c *Cluster) SignEthContractCall(from EthAccount, contractAddr common.Address, data []byte, nonce, gasLimit uint64) (SignedEthTx, error) {
	chainID := c.EthChainID()
	signer := ethtypes.LatestSignerForChainID(chainID)

	tx := ethtypes.NewTx(&ethtypes.LegacyTx{
		Nonce:    nonce,
		To:       &contractAddr,
		Value:    big.NewInt(0),
		Gas:      gasLimit,
		GasPrice: big.NewInt(0),
		Data:     data,
	})

	signedTx, err := ethtypes.SignTx(tx, signer, from.PrivateKey)
	if err != nil {
		return SignedEthTx{}, fmt.Errorf("sign eth contract call: %w", err)
	}

	rawBytes, err := signedTx.MarshalBinary()
	if err != nil {
		return SignedEthTx{}, fmt.Errorf("marshal eth tx: %w", err)
	}

	return SignedEthTx{
		Account: from.Name,
		Nonce:   nonce,
		TxHash:  signedTx.Hash().Hex(),
		RawHex:  "0x" + hex.EncodeToString(rawBytes),
	}, nil
}

func bech32ToEthAddress(addr string) (common.Address, error) {
	bz, err := sdk.GetFromBech32(addr, "init")
	if err != nil {
		return common.Address{}, fmt.Errorf("decode bech32 address %s: %w", addr, err)
	}

	return common.BytesToAddress(bz), nil
}
