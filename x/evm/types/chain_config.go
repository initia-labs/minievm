package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/params"
)

func DefaultChainConfig() *params.ChainConfig {
	shanghaiTime := uint64(0)
	cancunTime := uint64(0)
	pragueTime := uint64(0)
	verkleTime := uint64(0)

	return &params.ChainConfig{
		ChainID:             nil,
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
