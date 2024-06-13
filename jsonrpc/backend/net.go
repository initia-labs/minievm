package backend

import (
	"math/big"
	"strconv"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/initia-labs/minievm/x/evm/types"
)

func (b *JSONRPCBackend) ChainId() (*hexutil.Big, error) {
	chainID, err := b.ChainID()
	return (*hexutil.Big)(chainID), err
}

func (b *JSONRPCBackend) ChainID() (*big.Int, error) {
	queryCtx, err := b.getQueryCtx()
	if err != nil {
		return nil, err
	}

	sdkCtx := sdk.UnwrapSDKContext(queryCtx)
	return types.ConvertCosmosChainIDToEthereumChainID(sdkCtx.ChainID()), nil
}

func (b *JSONRPCBackend) Version() (string, error) {
	queryCtx, err := b.getQueryCtx()
	if err != nil {
		return "", err
	}

	sdkCtx := sdk.UnwrapSDKContext(queryCtx)
	items := strings.Split(sdkCtx.ChainID(), "-")
	version := items[len(items)-1]

	if _, err = strconv.Atoi(version); err != nil {
		return "1", nil
	}

	return version, err
}
