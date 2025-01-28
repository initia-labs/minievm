package backend

import (
	"errors"
	"time"

	"cosmossdk.io/collections"
	"github.com/ethereum/go-ethereum/common/bitutil"
	"github.com/ethereum/go-ethereum/core/bloombits"
	coretypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"

	evmconfig "github.com/initia-labs/minievm/x/evm/config"
)

// GetLogsByHeight returns all the logs from all the ethereum transactions in a block.
func (b *JSONRPCBackend) GetLogsByHeight(height uint64) ([]*coretypes.Log, error) {
	if blockLogs, ok := b.logsCache.Get(height); ok {
		return blockLogs, nil
	}

	blockHeader, err := b.GetHeaderByNumber(rpc.BlockNumber(height))
	if err != nil {
		return nil, err
	} else if blockHeader == nil {
		return nil, nil
	}

	txs, err := b.getBlockTransactions(height)
	if err != nil {
		return nil, err
	}
	receipts, err := b.getBlockReceipts(height)
	if err != nil {
		return nil, err
	}
	if len(txs) != len(receipts) {
		return nil, NewInternalError("mismatched number of transactions and receipts")
	}

	blockLogs := []*coretypes.Log{}
	for i, tx := range txs {
		receipt := receipts[i]
		logs := receipt.Logs
		for idx, log := range logs {
			log.BlockHash = blockHeader.Hash()
			log.BlockNumber = height
			log.TxHash = tx.Hash
			log.Index = uint(idx)
			log.TxIndex = receipt.TransactionIndex
		}
		blockLogs = append(blockLogs, logs...)
	}

	// cache the logs
	_ = b.logsCache.Add(height, blockLogs)
	return blockLogs, nil
}

const (
	// bloomServiceThreads is the number of goroutines used globally by an Ethereum
	// instance to service bloombits lookups for all running filters.
	bloomServiceThreads = 16

	// bloomFilterThreads is the number of goroutines used locally per filter to
	// multiplex requests onto the global servicing goroutines.
	bloomFilterThreads = 3

	// bloomRetrievalBatch is the maximum number of bloom bit retrievals to service
	// in a single batch.
	bloomRetrievalBatch = 16

	// bloomRetrievalWait is the maximum time to wait for enough bloom bit requests
	// to accumulate request an entire batch (avoiding hysteresis).
	bloomRetrievalWait = time.Duration(0)
)

func (b *JSONRPCBackend) BloomStatus() (uint64, uint64, error) {
	queryCtx, err := b.getQueryCtx()
	if err != nil {
		return 0, 0, err
	}

	sections, err := b.app.EVMIndexer().PeekBloomBitsNextSection(queryCtx)
	if err != nil {
		return 0, 0, err
	}

	return evmconfig.SectionSize, sections, nil
}

func (b *JSONRPCBackend) ServiceFilter(session *bloombits.MatcherSession) {
	for i := 0; i < bloomFilterThreads; i++ {
		go session.Multiplex(bloomRetrievalBatch, bloomRetrievalWait, b.bloomRequests)
	}
}

// startBloomHandlers starts a batch of goroutines to accept bloom bit database
// retrievals from possibly a range of filters and serving the data to satisfy.
func (b *JSONRPCBackend) startBloomHandlers(sectionSize uint64) {
	for i := 0; i < bloomServiceThreads; i++ {
		go func() {
			for {
				select {
				case <-b.ctx.Done():
					return

				case request := <-b.bloomRequests:
					task := <-request
					task.Bitsets = make([][]byte, len(task.Sections))

					queryCtx, err := b.getQueryCtx()
					if err != nil {
						task.Error = err
						request <- task
						continue
					}

					for i, section := range task.Sections {
						compVector, err := b.app.EVMIndexer().ReadBloomBits(queryCtx, section, uint32(task.Bit))
						if errors.Is(err, collections.ErrNotFound) {
							// pruned section, return empty bitset
							task.Bitsets[i] = make([]byte, evmconfig.SectionSize/8)
							continue
						} else if err != nil {
							task.Error = err
							break
						}

						blob, err := bitutil.DecompressBytes(compVector, int(sectionSize/8))
						if err != nil {
							task.Error = err
							break
						}

						task.Bitsets[i] = blob
					}
					request <- task
				}
			}
		}()
	}
}
