package filters

import (
	"context"
	"errors"
	"fmt"
	"math/big"

	"cosmossdk.io/log"

	"github.com/ethereum/go-ethereum/common"
	coretypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/initia-labs/minievm/jsonrpc/backend"
)

// BloomIV represents the bit indexes and value inside the bloom filter that belong
// to some key.
type BloomIV struct {
	I [3]uint
	V [3]byte
}

// Filter can be used to retrieve and filter logs.
type Filter struct {
	logger  log.Logger
	backend *backend.JSONRPCBackend

	addresses []common.Address
	topics    [][]common.Hash

	block      *common.Hash
	begin, end int64
}

// newBlockFilter creates a new filter which directly inspects the contents of
// a block to figure out whether it is interesting or not.
func newBlockFilter(logger log.Logger, backend *backend.JSONRPCBackend, block common.Hash, addresses []common.Address, topics [][]common.Hash) *Filter {
	filter := newFilter(logger, backend, addresses, topics)
	filter.block = &block

	return filter
}

// newRangeFilter creates a new filter which uses a bloom filter on blocks to
// figure out whether a particular block is interesting or not.
func newRangeFilter(logger log.Logger, backend *backend.JSONRPCBackend, begin, end int64, addresses []common.Address, topics [][]common.Hash) *Filter {
	filter := newFilter(logger, backend, addresses, topics)
	filter.begin = begin
	filter.end = end

	return filter
}

// newFilter returns a new Filter
func newFilter(
	logger log.Logger,
	backend *backend.JSONRPCBackend,
	addresses []common.Address,
	topics [][]common.Hash,
) *Filter {
	return &Filter{
		logger:    logger,
		backend:   backend,
		addresses: addresses,
		topics:    topics,
	}
}

// Logs searches the blockchain for matching log entries, returning all from the
// first block that contains matches, updating the start of the filter accordingly.
func (f *Filter) Logs(ctx context.Context) ([]*coretypes.Log, error) {
	var err error

	// If we're doing singleton block filtering, execute and return
	if f.block != nil && *f.block != (common.Hash{}) {
		header, err := f.backend.GetHeaderByHash(*f.block)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch block header by hash %s: %w", f.block, err)
		}
		if header == nil {
			return nil, errors.New("unknown block")
		}
		return f.blockLogs(header)
	}

	// Figure out the limits of the filter range
	header, err := f.backend.GetHeaderByNumber(rpc.LatestBlockNumber)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch header by number (latest): %w", err)
	}
	if header == nil || header.Number == nil {
		f.logger.Debug("header not found or has no number")
		return nil, nil
	}

	head := header.Number.Int64()

	// resolve special
	if f.begin < 0 {
		f.begin = head
	} else if f.begin == 0 {
		f.begin = 1
	}
	if f.end < 0 {
		f.end = head
	} else if f.end == 0 {
		f.end = 1
	}
	if f.end < f.begin {
		return nil, fmt.Errorf("invalid range [%d, %d]", f.begin, f.end)
	}

	// check bounds
	if f.begin > head {
		return []*coretypes.Log{}, nil
	} else if f.end > head {
		f.end = head
	}

	logChan, errChan := f.rangeLogsAsync(ctx)
	var logs []*coretypes.Log
	for {
		select {
		case log := <-logChan:
			logs = append(logs, log)
		case err := <-errChan:
			if err != nil {
				// if an error occurs during extraction, we do return the extracted data
				return logs, err
			}
			return logs, nil
		}
	}
}

// rangeLogsAsync retrieves block-range logs that match the filter criteria asynchronously,
// it creates and returns two channels: one for delivering log data, and one for reporting errors.
func (f *Filter) rangeLogsAsync(ctx context.Context) (chan *coretypes.Log, chan error) {
	var (
		logChan = make(chan *coretypes.Log)
		errChan = make(chan error)
	)

	go func() {
		defer func() {
			close(errChan)
			close(logChan)
		}()

		// Gather all non indexed ones
		if err := f.unindexedLogs(ctx, uint64(f.end), logChan); err != nil {
			errChan <- err
			return
		}

		errChan <- nil
	}()

	return logChan, errChan
}

// unindexedLogs returns the logs matching the filter criteria based on raw block
// iteration and bloom matching.
func (f *Filter) unindexedLogs(ctx context.Context, end uint64, logChan chan *coretypes.Log) error {
	for ; f.begin <= int64(end); f.begin++ {
		header, err := f.backend.GetHeaderByNumber(rpc.BlockNumber(f.begin))
		if header == nil || err != nil {
			return err
		}
		found, err := f.blockLogs(header)
		if err != nil {
			return err
		}
		for _, log := range found {
			select {
			case logChan <- log:
			case <-ctx.Done():
				return ctx.Err()
			}
		}
	}
	return nil
}

// blockLogs returns the logs matching the filter criteria within a single block.
func (f *Filter) blockLogs(header *coretypes.Header) ([]*coretypes.Log, error) {
	if bloomFilter(header.Bloom, f.addresses, f.topics) {
		return f.checkMatches(header)
	}
	return nil, nil
}

// checkMatches checks if the receipts belonging to the given header contain any log events that
// match the filter criteria. This function is called when the bloom filter signals a potential match.
func (f *Filter) checkMatches(header *coretypes.Header) ([]*coretypes.Log, error) {
	logs, err := f.backend.GetLogsByHeight(header.Number.Uint64())
	if err != nil {
		return nil, err
	}

	logs = filterLogs(logs, nil, nil, f.addresses, f.topics)
	return logs, nil
}

// filterLogs creates a slice of logs matching the given criteria.
func filterLogs(logs []*coretypes.Log, fromBlock, toBlock *big.Int, addresses []common.Address, topics [][]common.Hash) []*coretypes.Log {
	var check = func(log *coretypes.Log) bool {
		if fromBlock != nil && fromBlock.Int64() >= 0 && fromBlock.Uint64() > log.BlockNumber {
			return false
		}
		if toBlock != nil && toBlock.Int64() >= 0 && toBlock.Uint64() < log.BlockNumber {
			return false
		}
		if len(addresses) > 0 && !includes(addresses, log.Address) {
			return false
		}
		// If the to filtered topics is greater than the amount of topics in logs, skip.
		if len(topics) > len(log.Topics) {
			return false
		}
		for i, sub := range topics {
			if len(sub) == 0 {
				continue // empty rule set == wildcard
			}
			if !includes(sub, log.Topics[i]) {
				return false
			}
		}
		return true
	}
	var ret []*coretypes.Log
	for _, log := range logs {
		if check(log) {
			ret = append(ret, log)
		}
	}
	return ret
}

// includes returns true if the element is present in the list.
func includes[T comparable](things []T, element T) bool {
	for _, thing := range things {
		if thing == element {
			return true
		}
	}
	return false
}

func bloomFilter(bloom coretypes.Bloom, addresses []common.Address, topics [][]common.Hash) bool {
	if len(addresses) > 0 {
		var included bool
		for _, addr := range addresses {
			if coretypes.BloomLookup(bloom, addr) {
				included = true
				break
			}
		}
		if !included {
			return false
		}
	}

	for _, sub := range topics {
		included := len(sub) == 0 // empty rule set == wildcard
		for _, topic := range sub {
			if coretypes.BloomLookup(bloom, topic) {
				included = true
				break
			}
		}
		if !included {
			return false
		}
	}
	return true
}
