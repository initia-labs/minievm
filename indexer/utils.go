package indexer

import (
	"encoding/json"
	"fmt"
	"path/filepath"

	"github.com/spf13/cast"

	abci "github.com/cometbft/cometbft/abci/types"

	collcodec "cosmossdk.io/collections/codec"
	dbm "github.com/cosmos/cosmos-db"
	"github.com/cosmos/cosmos-sdk/server"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"

	coretypes "github.com/ethereum/go-ethereum/core/types"

	"github.com/initia-labs/minievm/x/evm/types"
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
		if event.Type == types.EventTypeEVM {
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

// CollJsonVal is used for protobuf values of the newest google.golang.org/protobuf API.
func CollJsonVal[T any]() collcodec.ValueCodec[T] {
	return &collJsonVal[T]{}
}

type collJsonVal[T any] struct{}

func (c collJsonVal[T]) Encode(value T) ([]byte, error) {
	return json.Marshal(value)
}

func (c collJsonVal[T]) Decode(b []byte) (T, error) {
	var value T

	err := json.Unmarshal(b, &value)
	return value, err
}

func (c collJsonVal[T]) EncodeJSON(value T) ([]byte, error) {
	return json.Marshal(value)
}

func (c collJsonVal[T]) DecodeJSON(b []byte) (T, error) {
	var value T

	err := json.Unmarshal(b, &value)
	return value, err
}

func (c collJsonVal[T]) Stringify(value T) string {
	return fmt.Sprintf("%v", value)
}

func (c collJsonVal[T]) ValueType() string {
	return "jsonvalue"
}
