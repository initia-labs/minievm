package backend

import (
	"strconv"
	"strings"

	rpcclient "github.com/cometbft/cometbft/rpc/client"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

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

func (b *JSONRPCBackend) PeerCount() (hexutil.Uint, error) {
	netInfo, err := b.clientCtx.Client.(rpcclient.NetworkClient).NetInfo(b.ctx)
	if err != nil {
		return hexutil.Uint(0), err
	}

	return hexutil.Uint(netInfo.NPeers), nil
}

func (b *JSONRPCBackend) Listening() (bool, error) {
	netInfo, err := b.clientCtx.Client.(rpcclient.NetworkClient).NetInfo(b.ctx)
	if err != nil {
		return false, err
	}

	return netInfo.Listening, nil
}
