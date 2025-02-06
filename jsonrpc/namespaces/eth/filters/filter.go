package filters

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"slices"

	"cosmossdk.io/log"
	"golang.org/x/sync/errgroup"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/bloombits"
	coretypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/initia-labs/minievm/jsonrpc/backend"
	evmconfig "github.com/initia-labs/minievm/x/evm/config"
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

	matcher *bloombits.Matcher
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

	// Flatten the address and topic filter clauses into a single bloombits filter
	// system. Since the bloombits are not positional, nil topics are permitted,
	// which get flattened into a nil byte slice.
	var filters [][][]byte
	if len(addresses) > 0 {
		filter := make([][]byte, len(addresses))
		for i, address := range addresses {
			filter[i] = address.Bytes()
		}
		filters = append(filters, filter)
	}
	for _, topicList := range topics {
		filter := make([][]byte, len(topicList))
		for i, topic := range topicList {
			filter[i] = topic.Bytes()
		}
		filters = append(filters, filter)
	}
	filter.matcher = bloombits.NewMatcher(evmconfig.SectionSize, filters)

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
			return logs, err
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

		size, sections, err := f.backend.BloomStatus()
		if err != nil {
			errChan <- err
			return
		}

		end := uint64(f.end)
		if indexed := sections * size; indexed > uint64(f.begin) {
			if indexed > end {
				indexed = end + 1
			}
			if err = f.indexedLogs(ctx, indexed-1, logChan); err != nil {
				errChan <- err
				return
			}
		}

		// Gather all non indexed ones
		if err := f.unindexedLogs(ctx, logChan); err != nil {
			errChan <- err
			return
		}

		errChan <- nil
	}()

	return logChan, errChan
}

// indexedLogs returns the logs matching the filter criteria based on the bloom
// bits indexed available locally or via the network.
func (f *Filter) indexedLogs(ctx context.Context, end uint64, logChan chan *coretypes.Log) error {
	// Create a matcher session and request servicing from the backend
	matches := make(chan uint64, 64)

	session, err := f.matcher.Start(ctx, uint64(f.begin), end, matches)
	if err != nil {
		return err
	}
	defer session.Close()

	f.backend.ServiceFilter(session)

	for {
		select {
		case number, ok := <-matches:
			// Abort if all matches have been fulfilled
			if !ok {
				err := session.Error()
				if err == nil {
					f.begin = int64(end) + 1
				}
				return err
			}
			f.begin = int64(number) + 1

			// Retrieve the suggested block and pull any truly matching logs
			header, err := f.backend.GetHeaderByNumber(rpc.BlockNumber(number))
			if header == nil || err != nil {
				return err
			}
			found, err := f.checkMatches(header)
			if err != nil {
				return err
			}
			for _, log := range found {
				logChan <- log
			}

		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

// unindexedLogs returns the logs matching the filter criteria based on raw block
// iteration and bloom matching.
func (f *Filter) unindexedLogs(ctx context.Context, logChan chan *coretypes.Log) error {
	const batchSize = 500

	g, innerCtx := errgroup.WithContext(ctx)
	diff := f.end - f.begin + 1
	batchNum := diff / batchSize
	if diff%batchSize != 0 {
		batchNum++
	}

	logsArray := make([][]*coretypes.Log, batchNum)
	for i := int64(0); i < batchNum; i++ {

		// make local copy of i for goroutine
		idx := i
		begin := f.begin + i*batchSize
		end := begin + batchSize - 1
		if end > f.end {
			end = f.end
		}

		// fetch logs in parallel
		g.Go(func() error {
			logs, err := f.searchLogs(innerCtx, begin, end)
			if err != nil {
				return err
			}

			logsArray[idx] = logs
			return nil
		})
	}

	// wait for all goroutines to finish
	err := g.Wait()
	if err != nil {
		return err
	}

	// send logs to channel in order
	for _, logs := range logsArray {
		for _, log := range logs {
			select {
			case logChan <- log:
			case <-ctx.Done():
				return ctx.Err()
			}
		}
	}

	return nil
}

func (f *Filter) searchLogs(ctx context.Context, begin, end int64) ([]*coretypes.Log, error) {
	logs := make([]*coretypes.Log, 0)
	for ; begin <= int64(end); begin++ {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		header, err := f.backend.GetHeaderByNumber(rpc.BlockNumber(begin))
		if header == nil {
			continue
		} else if err != nil {
			return nil, err
		}

		found, err := f.blockLogs(header)
		if err != nil {
			return nil, err
		}

		logs = append(logs, found...)
	}

	return logs, nil
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
		if len(addresses) > 0 && !slices.Contains(addresses, log.Address) {
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
			if !slices.Contains(sub, log.Topics[i]) {
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
