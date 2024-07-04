// Copyright 2021 Evmos Foundation
// This file is part of Evmos' Ethermint library.
//
// The Ethermint library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The Ethermint library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the Ethermint library. If not, see https://github.com/evmos/ethermint/blob/main/LICENSE
package filters

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/cosmos/cosmos-sdk/client"

	"cosmossdk.io/log"

	coretypes "github.com/cometbft/cometbft/rpc/core/types"
	rpcclient "github.com/cometbft/cometbft/rpc/jsonrpc/client"
	tmtypes "github.com/cometbft/cometbft/types"

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/eth/filters"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/initia-labs/minievm/jsonrpc/backend"

	sdk "github.com/cosmos/cosmos-sdk/types"
	evmtypes "github.com/initia-labs/minievm/x/evm/types"
)

// consider a filter inactive if it has not been polled for within deadline
var deadline = 5 * time.Minute

// filter is a helper struct that holds meta information over the filter type
// and associated subscription in the event system.
type filter struct {
	typ      filters.Type
	deadline *time.Timer // filter is inactive when deadline triggers
	hashes   []common.Hash
	crit     filters.FilterCriteria
	logs     []*ethtypes.Log
	s        *Subscription // associated subscription in event system
}

// PublicFilterAPI offers support to create and manage filters. This will allow external clients to retrieve various
// information related to the Ethereum protocol such as blocks, transactions and logs.
type PublicFilterAPI struct {
	logger    log.Logger
	clientCtx client.Context
	backend   *backend.JSONRPCBackend
	events    *EventSystem
	filtersMu sync.Mutex
	filters   map[rpc.ID]*filter
}

// NewFiltersAPI returns a new instance
func NewFilterAPI(logger log.Logger, backend *backend.JSONRPCBackend, clientCtx client.Context, tmWSClient *rpcclient.WSClient) *PublicFilterAPI {
	logger = logger.With("api", "filter")
	api := &PublicFilterAPI{
		logger:    logger,
		clientCtx: clientCtx,
		backend:   backend,
		filters:   make(map[rpc.ID]*filter),
		events:    NewEventSystem(logger, tmWSClient),
	}

	go api.timeoutLoop()

	return api
}

// timeoutLoop runs every 5 minutes and deletes filters that have not been recently used.
// Tt is started when the api is created.
func (api *PublicFilterAPI) timeoutLoop() {
	ticker := time.NewTicker(deadline)
	defer ticker.Stop()

	for {
		<-ticker.C
		api.filtersMu.Lock()
		for id, f := range api.filters {
			select {
			case <-f.deadline.C:
				f.s.Unsubscribe(api.events)
				delete(api.filters, id)
			default:
				continue
			}
		}
		api.filtersMu.Unlock()
	}
}

// TODO:  Implement eth_newPendingTransactionFilter
// func (api *PublicFilterAPI) NewPendingTransactionFilter() rpc.ID {}

// NewBlockFilter creates a filter that fetches blocks that are imported into the chain.
// It is part of the filter package since polling goes with eth_getFilterChanges.
//
// https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_newblockfilter
func (api *PublicFilterAPI) NewBlockFilter() rpc.ID {
	api.filtersMu.Lock()
	defer api.filtersMu.Unlock()

	if len(api.filters) >= int(api.backend.RPCFilterCap()) {
		return rpc.ID("error creating block filter: max limit reached")
	}

	headerSub, cancelSubs, err := api.events.SubscribeNewHeads()
	if err != nil {
		// wrap error on the ID
		return rpc.ID(fmt.Sprintf("error creating block filter: %s", err.Error()))
	}

	api.filters[headerSub.ID()] = &filter{typ: filters.BlocksSubscription, deadline: time.NewTimer(deadline), hashes: []common.Hash{}, s: headerSub}

	go func(headersCh <-chan coretypes.ResultEvent, errCh <-chan error) {
		defer cancelSubs()

		for {
			select {
			case ev, ok := <-headersCh:
				if !ok {
					api.filtersMu.Lock()
					delete(api.filters, headerSub.ID())
					api.filtersMu.Unlock()
					return
				}

				data, ok := ev.Data.(tmtypes.EventDataNewBlockHeader)
				if !ok {
					api.logger.Debug("event data type mismatch", "type", fmt.Sprintf("%T", ev.Data))
					continue
				}

				api.filtersMu.Lock()
				if f, found := api.filters[headerSub.ID()]; found {
					f.hashes = append(f.hashes, common.BytesToHash(data.Header.Hash()))
				}
				api.filtersMu.Unlock()
			case <-errCh:
				api.filtersMu.Lock()
				delete(api.filters, headerSub.ID())
				api.filtersMu.Unlock()
				return
			}
		}
	}(headerSub.eventCh, headerSub.Err())

	return headerSub.ID()
}

// NewFilter creates a new filter and returns the filter id. It can be
// used to retrieve logs when the state changes. This method cannot be
// used to fetch logs that are already stored in the state.
//
// Default criteria for the from and to block are "latest".
// Using "latest" as block number will return logs for mined blocks.
// Using "pending" as block number returns logs for not yet mined (pending) blocks.
// In case logs are removed (chain reorg) previously returned logs are returned
// again but with the removed property set to true.
//
// In case "fromBlock" > "toBlock" an error is returned.
//
// https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_newfilter
func (api *PublicFilterAPI) NewFilter(criteria filters.FilterCriteria) (rpc.ID, error) {
	api.filtersMu.Lock()
	defer api.filtersMu.Unlock()

	if len(api.filters) >= int(api.backend.RPCFilterCap()) {
		return rpc.ID(""), fmt.Errorf("error creating filter: max limit reached")
	}

	var (
		filterID = rpc.ID("")
		err      error
	)

	logsSub, cancelSubs, err := api.events.SubscribeLogs(criteria)
	if err != nil {
		return rpc.ID(""), err
	}

	filterID = logsSub.ID()

	api.filters[filterID] = &filter{
		typ:      filters.LogsSubscription,
		crit:     criteria,
		deadline: time.NewTimer(deadline),
		hashes:   []common.Hash{},
		s:        logsSub,
	}

	go func(eventCh <-chan coretypes.ResultEvent) {
		defer cancelSubs()

		for {
			select {
			case ev, ok := <-eventCh:
				if !ok {
					api.filtersMu.Lock()
					delete(api.filters, filterID)
					api.filtersMu.Unlock()
					return
				}
				dataTx, ok := ev.Data.(tmtypes.EventDataTx)
				if !ok {
					api.logger.Debug("event data type mismatch", "type", fmt.Sprintf("%T", ev.Data))
					continue
				}

				var txResp sdk.TxMsgData
				var ethResp evmtypes.MsgCallResponse

				err := txResp.Unmarshal(dataTx.TxResult.Result.Data)
				if err != nil {
					api.logger.Error("failed to decode tx response", "error", err.Error())
					return
				}

				if len(txResp.MsgResponses) == 0 {
					continue
				}

				err = ethResp.Unmarshal(txResp.MsgResponses[0].Value)
				if len(ethResp.Logs) == 0 {
					continue
				}

				logs := evmtypes.Logs.ToEthLogs(ethResp.Logs)
				if len(logs) == 0 {
					continue
				}

				api.filtersMu.Lock()
				if f, found := api.filters[filterID]; found {
					f.logs = append(f.logs, logs...)
				}
				api.filtersMu.Unlock()

			case <-logsSub.Err():
				api.filtersMu.Lock()
				delete(api.filters, filterID)
				api.filtersMu.Unlock()
				return
			}
		}
	}(logsSub.eventCh)

	return filterID, err
}

// GetLogs returns logs matching the given argument that are stored within the state.
//
// https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_getlogs
func (api *PublicFilterAPI) GetLogs(ctx context.Context, crit filters.FilterCriteria) ([]*ethtypes.Log, error) {
	var filter *Filter
	if crit.BlockHash != nil {
		// Block filter requested, construct a single-shot filter
		filter = NewBlockFilter(api.logger, *api.backend, crit)
	} else {
		// Convert the RPC block numbers into internal representations
		begin := rpc.LatestBlockNumber.Int64()
		if crit.FromBlock != nil {
			begin = crit.FromBlock.Int64()
		}
		end := rpc.LatestBlockNumber.Int64()
		if crit.ToBlock != nil {
			end = crit.ToBlock.Int64()
		}
		// Construct the range filter
		filter = NewRangeFilter(api.logger, *api.backend, begin, end, crit.Addresses, crit.Topics)
	}

	// Run the filter and return all the logs
	logs, err := filter.Logs(ctx, int(api.backend.RPCLogsCap()), int64(api.backend.RPCBlockRangeCap()))
	if err != nil {
		return nil, err
	}

	return returnLogs(logs), err
}

// UninstallFilter removes the filter with the given filter id.
//
// https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_uninstallfilter
func (api *PublicFilterAPI) UninstallFilter(id rpc.ID) bool {
	api.filtersMu.Lock()
	f, found := api.filters[id]
	if found {
		delete(api.filters, id)
	}
	api.filtersMu.Unlock()

	if !found {
		return false
	}
	f.s.Unsubscribe(api.events)
	return true
}

// GetFilterLogs returns the logs for the filter with the given id.
// If the filter could not be found an empty array of logs is returned.
//
// https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_getfilterlogs
func (api *PublicFilterAPI) GetFilterLogs(ctx context.Context, id rpc.ID) ([]*ethtypes.Log, error) {
	api.filtersMu.Lock()
	f, found := api.filters[id]
	api.filtersMu.Unlock()

	if !found {
		return returnLogs(nil), fmt.Errorf("filter %s not found", id)
	}

	if f.typ != filters.LogsSubscription {
		return returnLogs(nil), fmt.Errorf("filter %s doesn't have a LogsSubscription type: got %d", id, f.typ)
	}

	var filter *Filter
	if f.crit.BlockHash != nil {
		// Block filter requested, construct a single-shot filter
		filter = NewBlockFilter(api.logger, *api.backend, f.crit)
	} else {
		// Convert the RPC block numbers into internal representations
		begin := rpc.LatestBlockNumber.Int64()
		if f.crit.FromBlock != nil {
			begin = f.crit.FromBlock.Int64()
		}
		end := rpc.LatestBlockNumber.Int64()
		if f.crit.ToBlock != nil {
			end = f.crit.ToBlock.Int64()
		}
		// Construct the range filter
		filter = NewRangeFilter(api.logger, *api.backend, begin, end, f.crit.Addresses, f.crit.Topics)
	}
	// Run the filter and return all the logs
	logs, err := filter.Logs(ctx, int(api.backend.RPCLogsCap()), int64(api.backend.RPCBlockRangeCap()))
	if err != nil {
		return nil, err
	}
	return returnLogs(logs), nil
}

// GetFilterChanges returns the logs for the filter with the given id since
// last time it was called. This can be used for polling.
//
// For pending transaction and block filters the result is []common.Hash.
// (pending)Log filters return []Log.
//
// https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_getfilterchanges
func (api *PublicFilterAPI) GetFilterChanges(id rpc.ID) (interface{}, error) {
	api.filtersMu.Lock()
	defer api.filtersMu.Unlock()

	f, found := api.filters[id]
	if !found {
		return nil, fmt.Errorf("filter %s not found", id)
	}

	if !f.deadline.Stop() {
		// timer expired but filter is not yet removed in timeout loop
		// receive timer value and reset timer
		<-f.deadline.C
	}
	f.deadline.Reset(deadline)

	switch f.typ {
	case filters.PendingTransactionsSubscription, filters.BlocksSubscription:
		hashes := f.hashes
		f.hashes = nil
		return returnHashes(hashes), nil
	case filters.LogsSubscription, filters.MinedAndPendingLogsSubscription:
		logs := make([]*ethtypes.Log, len(f.logs))
		copy(logs, f.logs)
		f.logs = []*ethtypes.Log{}
		return returnLogs(logs), nil
	default:
		return nil, fmt.Errorf("invalid filter %s type %d", id, f.typ)
	}
}
