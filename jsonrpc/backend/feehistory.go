package backend

import (
	"encoding/binary"
	"fmt"
	"math"
	"math/big"
	"slices"
	"sync/atomic"

	"github.com/ethereum/go-ethereum/common"
	coretypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"

	rpctypes "github.com/initia-labs/minievm/jsonrpc/types"
)

const (
	// maxBlockFetchers is the max number of goroutines to spin up to pull blocks
	// for the fee history calculation (mostly relevant for LES).
	maxBlockFetchers = 4
	// maxQueryLimit is the max number of requested percentiles.
	maxQueryLimit = 100
)

// blockFees represents a single block for processing
type blockFees struct {
	// set by the caller
	blockNumber uint64
	header      *coretypes.Header
	txs         coretypes.Transactions
	receipts    coretypes.Receipts
	// filled by processBlock
	results processedFees
	err     error
}

type cacheKey struct {
	number      uint64
	percentiles string
}

// processedFees contains the results of a processed block.
type processedFees struct {
	reward                       []*big.Int
	baseFee, nextBaseFee         *big.Int
	gasUsedRatio                 float64
	blobGasUsedRatio             float64
	blobBaseFee, nextBlobBaseFee *big.Int
}

// txGasAndReward is sorted in ascending order based on reward
type txGasAndReward struct {
	gasUsed uint64
	reward  *big.Int
}

// FeeHistory returns data relevant for fee estimation based on the specified range of blocks.
// The range can be specified either with absolute block numbers or ending with the latest
// or pending block. Backends may or may not support gathering data from the pending block
// or blocks older than a certain age (specified in maxHistory). The first block of the
// actually processed range is returned to avoid ambiguity when parts of the requested range
// are not available or when the head has changed during processing this request.
// Five arrays are returned based on the processed blocks:
//   - reward: the requested percentiles of effective priority fees per gas of transactions in each
//     block, sorted in ascending order and weighted by gas used.
//   - baseFee: base fee per gas in the given block
//   - gasUsedRatio: gasUsed/gasLimit in the given block
//   - blobBaseFee: the blob base fee per gas in the given block
//   - blobGasUsedRatio: blobGasUsed/blobGasLimit in the given block
//
// Note: baseFee and blobBaseFee both include the next block after the newest of the returned range,
// because this value can be derived from the newest block.
func (b *JSONRPCBackend) FeeHistory(blocks uint64, unresolvedLastBlock rpc.BlockNumber, rewardPercentiles []float64) (*big.Int, [][]*big.Int, []*big.Int, []float64, []*big.Int, []float64, error) {
	if blocks < 1 {
		return common.Big0, nil, nil, nil, nil, nil, nil
	}
	maxFeeHistory := uint64(b.cfg.FeeHistoryMaxHeaders)
	if len(rewardPercentiles) != 0 {
		maxFeeHistory = uint64(b.cfg.FeeHistoryMaxBlocks)
	}
	if len(rewardPercentiles) > maxQueryLimit {
		return common.Big0, nil, nil, nil, nil, nil, fmt.Errorf("%w: over the query limit %d", errInvalidPercentile, maxQueryLimit)
	}
	if blocks > maxFeeHistory {
		b.logger.Warn("Sanitizing fee history length", "requested", blocks, "truncated", maxFeeHistory)
		blocks = maxFeeHistory
	}
	for i, p := range rewardPercentiles {
		if p < 0 || p > 100 {
			return common.Big0, nil, nil, nil, nil, nil, fmt.Errorf("%w: %f", errInvalidPercentile, p)
		}
		if i > 0 && p <= rewardPercentiles[i-1] {
			return common.Big0, nil, nil, nil, nil, nil, fmt.Errorf("%w: #%d:%f >= #%d:%f", errInvalidPercentile, i-1, rewardPercentiles[i-1], i, p)
		}
	}
	lastBlock, blocks, err := b.resolveBlockRange(unresolvedLastBlock, blocks)
	if err != nil || blocks == 0 {
		return common.Big0, nil, nil, nil, nil, nil, err
	}
	oldestBlock := lastBlock + 1 - blocks

	var next atomic.Uint64
	next.Store(oldestBlock)
	results := make(chan *blockFees, blocks)

	percentileKey := make([]byte, 8*len(rewardPercentiles))
	for i, p := range rewardPercentiles {
		binary.LittleEndian.PutUint64(percentileKey[i*8:(i+1)*8], math.Float64bits(p))
	}
	for i := 0; i < maxBlockFetchers && i < int(blocks); i++ {
		go func() {
			for {
				// Retrieve the next block number to fetch with this goroutine
				blockNumber := next.Add(1) - 1
				if blockNumber > lastBlock {
					return
				}

				fees := &blockFees{blockNumber: blockNumber}
				cacheKey := cacheKey{number: blockNumber, percentiles: string(percentileKey)}

				if p, ok := b.historyCache.Get(cacheKey); ok {
					fees.results = p
					results <- fees
				} else {
					fees.header, fees.err = b.GetHeaderByNumber(rpc.BlockNumber(blockNumber))
					if len(rewardPercentiles) != 0 && fees.err == nil {
						fees.receipts, fees.err = b.getBlockReceipts(blockNumber)
						if fees.err == nil {
							var txs []*rpctypes.RPCTransaction
							txs, fees.err = b.getBlockTransactions(blockNumber)
							if fees.err == nil {
								for _, tx := range txs {
									fees.txs = append(fees.txs, tx.ToTransaction())
								}
							}
						}
					}
					if fees.header != nil && fees.err == nil {
						b.processBlock(fees, rewardPercentiles)
						if fees.err == nil {
							b.historyCache.Add(cacheKey, fees.results)
						}
					}
					// send to results even if empty to guarantee that blocks items are sent in total
					results <- fees
				}

			}
		}()
	}
	var (
		reward           = make([][]*big.Int, blocks)
		baseFee          = make([]*big.Int, blocks+1)
		gasUsedRatio     = make([]float64, blocks)
		blobGasUsedRatio = make([]float64, blocks)
		blobBaseFee      = make([]*big.Int, blocks+1)
		firstMissing     = blocks
	)
	for ; blocks > 0; blocks-- {
		fees := <-results
		if fees.err != nil {
			return common.Big0, nil, nil, nil, nil, nil, fees.err
		}
		i := fees.blockNumber - oldestBlock
		if fees.results.baseFee != nil {
			reward[i], baseFee[i], baseFee[i+1], gasUsedRatio[i] = fees.results.reward, fees.results.baseFee, fees.results.nextBaseFee, fees.results.gasUsedRatio
			blobGasUsedRatio[i], blobBaseFee[i], blobBaseFee[i+1] = fees.results.blobGasUsedRatio, fees.results.blobBaseFee, fees.results.nextBlobBaseFee
		} else {
			// getting no block and no error means we are requesting into the future (might happen because of a reorg)
			if i < firstMissing {
				firstMissing = i
			}
		}
	}
	if firstMissing == 0 {
		return common.Big0, nil, nil, nil, nil, nil, nil
	}
	if len(rewardPercentiles) != 0 {
		reward = reward[:firstMissing]
	} else {
		reward = nil
	}
	baseFee, gasUsedRatio = baseFee[:firstMissing+1], gasUsedRatio[:firstMissing]
	blobBaseFee, blobGasUsedRatio = blobBaseFee[:firstMissing+1], blobGasUsedRatio[:firstMissing]
	return new(big.Int).SetUint64(oldestBlock), reward, baseFee, gasUsedRatio, blobBaseFee, blobGasUsedRatio, nil
}

func (b *JSONRPCBackend) resolveBlockRange(reqEnd rpc.BlockNumber, blocks uint64) (uint64, uint64, error) {
	var (
		headBlock *coretypes.Header
		err       error
	)

	// Get the chain's current head.
	if headBlock, err = b.GetHeaderByNumber(rpc.LatestBlockNumber); err != nil {
		return 0, 0, err
	}
	head := rpc.BlockNumber(headBlock.Number.Uint64())

	// Fail if request block is beyond the chain's current head.
	if head < reqEnd {
		return 0, 0, fmt.Errorf("%w: requested %d, head %d", errRequestBeyondHead, reqEnd, head)
	}

	// return latest block if requested block is special
	if reqEnd < 0 {
		reqEnd = rpc.BlockNumber(headBlock.Number.Uint64())
	}

	// If there are no blocks to return, short circuit.
	if blocks == 0 {
		return 0, 0, nil
	}
	// Ensure not trying to retrieve before genesis.
	if uint64(reqEnd+1) < blocks {
		blocks = uint64(reqEnd + 1)
	}
	return uint64(reqEnd), blocks, nil
}

// processBlock takes a blockFees structure with the blockNumber, the header and optionally
// the block field filled in, retrieves the block from the backend if not present yet and
// fills in the rest of the fields.
func (b *JSONRPCBackend) processBlock(bf *blockFees, percentiles []float64) {
	// Fill in base fee and next base fee.
	if bf.results.baseFee = bf.header.BaseFee; bf.results.baseFee == nil {
		bf.results.baseFee = new(big.Int)
	}

	// NOTE: we don't have dynamic base fee calculation yet
	bf.results.nextBaseFee = bf.header.BaseFee
	bf.results.blobBaseFee = new(big.Int)
	bf.results.nextBlobBaseFee = new(big.Int)

	// Compute gas used ratio for normal and blob gas.
	bf.results.gasUsedRatio = float64(bf.header.GasUsed) / float64(bf.header.GasLimit)

	if len(percentiles) == 0 {
		// rewards were not requested, return null
		return
	}
	if bf.receipts == nil && len(bf.txs) == 0 {
		b.logger.Error("Block or receipts are missing while reward percentiles are requested")
		return
	}

	bf.results.reward = make([]*big.Int, len(percentiles))
	if len(bf.txs) == 0 {
		// return an all zero row if there are no transactions to gather data from
		for i := range bf.results.reward {
			bf.results.reward[i] = new(big.Int)
		}
		return
	}

	sorter := make([]txGasAndReward, len(bf.txs))
	for i, tx := range bf.txs {
		reward, _ := tx.EffectiveGasTip(bf.header.BaseFee)
		sorter[i] = txGasAndReward{gasUsed: bf.receipts[i].GasUsed, reward: reward}
	}
	slices.SortStableFunc(sorter, func(a, b txGasAndReward) int {
		return a.reward.Cmp(b.reward)
	})

	var txIndex int
	sumGasUsed := sorter[0].gasUsed

	for i, p := range percentiles {
		thresholdGasUsed := uint64(float64(bf.header.GasUsed) * p / 100)
		for sumGasUsed < thresholdGasUsed && txIndex < len(bf.txs)-1 {
			txIndex++
			sumGasUsed += sorter[txIndex].gasUsed
		}
		bf.results.reward[i] = sorter[txIndex].reward
	}
}
