package types

import (
	"context"
	"encoding/binary"
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/params"
	"golang.org/x/crypto/sha3"
)

func ConvertCosmosChainIDToEthereumChainID(chainID string) *big.Int {
	hasher := sha3.NewLegacyKeccak256()
	hasher.Write([]byte(chainID))
	hash := hasher.Sum(nil)

	// metamask max
	metamaskMax := uint64(4503599627370476)
	ethChainID := binary.BigEndian.Uint64(hash[:8]) % metamaskMax
	return new(big.Int).SetUint64(ethChainID)
}

func chainID(ctx context.Context) *big.Int {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	chainID := sdkCtx.ChainID()

	return ConvertCosmosChainIDToEthereumChainID(chainID)
}

func DefaultChainConfig(ctx context.Context) *params.ChainConfig {
	shanghaiTime := uint64(0)
	cancunTime := uint64(0)
	pragueTime := uint64(0)
	verkleTime := uint64(0)

	return &params.ChainConfig{
		ChainID:             chainID(ctx),
		HomesteadBlock:      big.NewInt(0),
		DAOForkBlock:        big.NewInt(0),
		DAOForkSupport:      true,
		EIP150Block:         big.NewInt(0),
		EIP155Block:         big.NewInt(0),
		EIP158Block:         big.NewInt(0),
		ByzantiumBlock:      big.NewInt(0),
		ConstantinopleBlock: big.NewInt(0),
		PetersburgBlock:     big.NewInt(0),
		IstanbulBlock:       big.NewInt(0),
		MuirGlacierBlock:    big.NewInt(0),
		BerlinBlock:         big.NewInt(0),
		LondonBlock:         big.NewInt(0),
		ArrowGlacierBlock:   big.NewInt(0),
		GrayGlacierBlock:    big.NewInt(0),
		MergeNetsplitBlock:  big.NewInt(0),
		ShanghaiTime:        &shanghaiTime,
		CancunTime:          &cancunTime,
		PragueTime:          &pragueTime,
		VerkleTime:          &verkleTime,
	}
}
