package indexer

import (
	"encoding/json"
	"path/filepath"

	abci "github.com/cometbft/cometbft/abci/types"
	dbm "github.com/cosmos/cosmos-db"
	"github.com/cosmos/cosmos-sdk/server"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	coretypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/initia-labs/minievm/x/evm/types"
	"github.com/spf13/cast"
)

// helper function to make config creation independent of root dir
func rootify(path, root string) string {
	if filepath.IsAbs(path) {
		return path
	}
	return filepath.Join(root, path)
}

// getDBConfig returns the database configuration for the EVM indexer
func getDBConfig(appOpts servertypes.AppOptions) (string, dbm.BackendType) {
	rootDir := cast.ToString(appOpts.Get("home"))
	dbDir := cast.ToString(appOpts.Get("db_dir"))
	dbBackend := server.GetAppDBBackend(appOpts)

	return rootify(dbDir, rootDir), dbBackend
}

// extractLogsFromEvents extracts logs from the events
func (e *EVMIndexerImpl) extractLogsFromEvents(events []abci.Event) []*coretypes.Log {
	var ethLogs []*coretypes.Log
	for _, event := range events {
		if event.Type == types.EventTypeCall || event.Type == types.EventTypeCreate {
			logs := make(types.Logs, 0, len(event.Attributes))

			for _, attr := range event.Attributes {
				if attr.Key == types.AttributeKeyLog {
					var log types.Log
					err := json.Unmarshal([]byte(attr.Value), &log)
					if err != nil {
						e.logger.Error("failed to unmarshal log", "err", err)
						continue
					}

					logs = append(logs, log)
				}
			}

			ethLogs = logs.ToEthLogs()
			break
		}
	}

	return ethLogs
}
