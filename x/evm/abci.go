package evm

import (
	"time"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/initia-labs/minievm/x/evm/keeper"
	"github.com/initia-labs/minievm/x/evm/types"
)

// PreBlock track latest 256 block hashes
func PreBlock(ctx sdk.Context, k *keeper.Keeper) (sdk.ResponsePreBlock, error) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyEndBlocker)
	return sdk.ResponsePreBlock{}, k.TrackBlockHash(ctx, uint64(ctx.BlockHeight()-1), common.BytesToHash(ctx.HeaderHash()))
}
