package types

import "github.com/ethereum/go-ethereum/params"

// CalGasUsed calculates gas used.
//
// NOTE: London enforced
func CalGasUsed(gasBalance, gasRemaining, gasRefunded uint64) uint64 {
	gasUsed := gasBalance - gasRemaining
	refund := gasUsed / params.RefundQuotientEIP3529
	if refund > gasRefunded {
		refund = gasRefunded
	}

	return gasUsed - refund
}
