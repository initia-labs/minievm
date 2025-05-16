package backend

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/lru"
	"github.com/ethereum/go-ethereum/core/bloombits"
	coretypes "github.com/ethereum/go-ethereum/core/types"

	"cosmossdk.io/log"
	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/server"

	"github.com/initia-labs/minievm/app"
	"github.com/initia-labs/minievm/jsonrpc/config"
	rpctypes "github.com/initia-labs/minievm/jsonrpc/types"
	evmconfig "github.com/initia-labs/minievm/x/evm/config"
)

type JSONRPCBackend struct {
	app    *app.MinitiaApp
	logger log.Logger

	historyCache *lru.Cache[cacheKey, processedFees]

	// per block caches
	headerCache        *lru.Cache[uint64, *coretypes.Header]
	blockTxsCache      *lru.Cache[uint64, []*rpctypes.RPCTransaction]
	blockReceiptsCache *lru.Cache[uint64, []*coretypes.Receipt]
	blockHashCache     *lru.Cache[common.Hash, uint64]
	logsCache          *lru.Cache[uint64, []*coretypes.Log]

	// per tx caches
	txLookupCache *lru.Cache[common.Hash, *rpctypes.RPCTransaction]
	receiptCache  *lru.Cache[common.Hash, *coretypes.Receipt]

	// fee cache
	feeDenom    string
	feeDecimals uint8
	feeMutex    sync.RWMutex

	ctx       context.Context
	svrCtx    *server.Context
	clientCtx client.Context

	cfg           config.JSONRPCConfig
	gasMultiplier math.LegacyDec

	// Channel receiving bloom data retrieval requests
	bloomRequests chan chan *bloombits.Retrieval
}

const (
	feeHistoryCacheSize = 2048
	blockCacheLimit     = 256
	txLookupCacheLimit  = 1024
)

// NewJSONRPCBackend creates a new JSONRPCBackend instance
func NewJSONRPCBackend(
	ctx context.Context,
	app *app.MinitiaApp,
	logger log.Logger,
	svrCtx *server.Context,
	clientCtx client.Context,
	cfg config.JSONRPCConfig,
) (*JSONRPCBackend, error) {
	if cfg.LogCacheSize == 0 {
		cfg.LogCacheSize = config.DefaultLogCacheSize
	}
	if cfg.GasMultiplier == "" {
		cfg.GasMultiplier = config.DefaultGasMultiplier
	}
	if cfg.FilterMaxBlockRange == 0 {
		cfg.FilterMaxBlockRange = config.DefaultFilterMaxBlockRange
	}
	if cfg.FilterMaxAddresses == 0 {
		cfg.FilterMaxAddresses = config.DefaultFilterMaxAddresses
	}

	gasMultiplier, err := math.LegacyNewDecFromStr(cfg.GasMultiplier)
	if err != nil {
		return nil, err
	}

	b := &JSONRPCBackend{
		app:    app,
		logger: logger.With("module", "jsonrpc"),

		historyCache: lru.NewCache[cacheKey, processedFees](feeHistoryCacheSize),

		// per block caches
		headerCache:        lru.NewCache[uint64, *coretypes.Header](blockCacheLimit),
		blockTxsCache:      lru.NewCache[uint64, []*rpctypes.RPCTransaction](blockCacheLimit),
		blockReceiptsCache: lru.NewCache[uint64, []*coretypes.Receipt](blockCacheLimit),
		blockHashCache:     lru.NewCache[common.Hash, uint64](blockCacheLimit),
		logsCache:          lru.NewCache[uint64, []*coretypes.Log](cfg.LogCacheSize),

		// per tx caches
		txLookupCache: lru.NewCache[common.Hash, *rpctypes.RPCTransaction](txLookupCacheLimit),
		receiptCache:  lru.NewCache[common.Hash, *coretypes.Receipt](txLookupCacheLimit),

		ctx:       ctx,
		svrCtx:    svrCtx,
		clientCtx: clientCtx,

		cfg:           cfg,
		gasMultiplier: gasMultiplier,

		bloomRequests: make(chan chan *bloombits.Retrieval),
	}

	// start fee fetcher
	go b.feeFetcher()

	// Start the bloom bits servicing goroutines
	b.startBloomHandlers(evmconfig.SectionSize)

	return b, nil
}

func (b *JSONRPCBackend) feeInfo() (string, uint8, error) {
	b.feeMutex.RLock()
	defer b.feeMutex.RUnlock()

	if b.feeDenom == "" {
		return "", 0, NewInternalError("jsonrpc is not ready")
	}

	return b.feeDenom, b.feeDecimals, nil
}

func (b *JSONRPCBackend) feeFetcher() {
	fetcher := func() (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = fmt.Errorf("feeFetcher panic: %v", r)
			}
		}()

		if b.app.LastBlockHeight() <= 1 {
			return nil
		}

		queryCtx, err := b.getQueryCtx()
		if err != nil {
			return err
		}

		params, err := b.app.EVMKeeper.Params.Get(queryCtx)
		if err != nil {
			return err
		}

		feeDenom := params.FeeDenom
		decimals, err := b.app.EVMKeeper.ERC20Keeper().GetDecimals(queryCtx, feeDenom)
		if err != nil {
			return err
		}

		b.feeMutex.Lock()
		b.feeDenom = feeDenom
		b.feeDecimals = decimals
		b.feeMutex.Unlock()

		return nil
	}

	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if err := fetcher(); err != nil {
				b.logger.Error("failed to fetch fee", "err", err)
			}
		case <-b.ctx.Done():
			return
		}
	}
}

func (b *JSONRPCBackend) FilterTimeout() time.Duration {
	return b.cfg.FilterTimeout
}

func (b *JSONRPCBackend) FilterMaxBlockRange() int {
	return b.cfg.FilterMaxBlockRange
}

func (b *JSONRPCBackend) FilterMaxAddresses() int {
	return b.cfg.FilterMaxAddresses
}
