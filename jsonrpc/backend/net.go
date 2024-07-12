package backend

import (
	"strconv"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
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
