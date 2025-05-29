package types

import (
	"math"

	"github.com/ethereum/go-ethereum/params"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// CalGasUsed calculates gas used.
//
// NOTE: London enforced
func CalGasUsed(gasBalance, gasRemaining, gasRefunded uint64) uint64 {
	gasUsed := gasBalance - gasRemaining
	refund := min(gasUsed/params.RefundQuotientEIP3529, gasRefunded)

	return gasUsed - refund
}

func BlockGasLimit(ctx sdk.Context) uint64 {
	if blockGasMeter := ctx.BlockGasMeter(); blockGasMeter != nil && blockGasMeter.Limit() != math.MaxUint64 {
		return blockGasMeter.Limit()
	}

	cp := ctx.ConsensusParams()
	if cp.Block != nil {
		if cp.Block.MaxGas == -1 {
			return math.MaxUint64
		}

		return uint64(cp.Block.MaxGas)
	} else {
		return 0
	}
}
