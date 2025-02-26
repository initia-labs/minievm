package indexer

import (
	"context"
	"encoding/json"
	"fmt"

	"cosmossdk.io/collections"
	snapshot "cosmossdk.io/store/snapshots/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	coretypes "github.com/ethereum/go-ethereum/core/types"

	rpctypes "github.com/initia-labs/minievm/jsonrpc/types"
)

// SnapshotFormat format 1 is putting evm block header json bytes
const SnapshotFormat = 1

// SnapshotInfo is a struct to store snapshot info
type SnapshotInfo struct {
	Header         *coretypes.Header          `json:"header"`
	Receipts       []*coretypes.Receipt       `json:"receipts"`
	Txs            []*rpctypes.RPCTransaction `json:"txs"`
	CosmosTxHashes [][]byte                   `json:"cosmos_tx_hashes"`
}

// RestoreExtension implements types.ExtensionSnapshotter.
func (e *EVMIndexerImpl) RestoreExtension(height uint64, format uint32, payloadReader snapshot.ExtensionPayloadReader) error {
	logger := e.logger.With("module", "snapshotter")
	if format != SnapshotFormat {
		return fmt.Errorf("failed to restore extension: invalid format %d", format)
	}

	// read payload
	bz, err := payloadReader()
	if err != nil {
		logger.Error("failed to read snapshot extension", "err", err)
		return err
	}

	// unmarshal payload
	var s SnapshotInfo
	err = json.Unmarshal(bz, &s)
	if err != nil {
		logger.Error("failed to unmarshal snapshot extension", "err", err)
		return err
	}

	ctx := sdk.Context{}.WithContext(context.Background())
	for i, tx := range s.Txs {
		if tx == nil {
			continue
		}

		txHash := tx.Hash

		// store tx
		err = e.TxMap.Set(ctx, tx.Hash.Bytes(), *tx)
		if err != nil {
			logger.Error("failed to store tx", "err", err)
			return err
		}

		// store TxHashToCosmosTxHash
		cosmosTxHash := s.CosmosTxHashes[i]
		if err := e.TxHashToCosmosTxHash.Set(ctx, tx.Hash.Bytes(), cosmosTxHash); err != nil {
			logger.Error("failed to store txHashToCosmosTxHash", "err", err)
			return err
		}

		// store CosmosTxHashToTxHash
		if err := e.CosmosTxHashToTxHash.Set(ctx, cosmosTxHash, tx.Hash.Bytes()); err != nil {
			logger.Error("failed to store cosmosTxHashToTxHash", "err", err)
			return err
		}

		// store receipt
		receipt := s.Receipts[i]
		if receipt != nil {
			if err := e.TxReceiptMap.Set(ctx, txHash.Bytes(), *receipt); err != nil {
				logger.Error("failed to store tx receipt", "err", err)
				return err
			}
		}

		// store blockAndIndexToTxHash
		if err := e.BlockAndIndexToTxHashMap.Set(ctx, collections.Join(height, uint64(receipt.TransactionIndex)), txHash.Bytes()); err != nil {
			logger.Error("failed to store blockAndIndexToTxHash", "err", err)
			return err
		}
	}

	// store block header
	if s.Header != nil {
		if err := e.BlockHeaderMap.Set(ctx, height, *s.Header); err != nil {
			logger.Error("failed to marshal blockHeader", "err", err)
			return err
		}
		if err := e.BlockHashToNumberMap.Set(ctx, s.Header.Hash().Bytes(), height); err != nil {
			logger.Error("failed to store blockHashToNumber", "err", err)
			return err
		}
	}

	logger.Info("restored snapshot extension", "height", height)
	return nil
}

// SnapshotExtension implements types.ExtensionSnapshotter.
func (e *EVMIndexerImpl) SnapshotExtension(height uint64, payloadWriter snapshot.ExtensionPayloadWriter) error {
	logger := e.logger.With("module", "snapshotter")
	var s SnapshotInfo
	ctx := sdk.Context{}.WithContext(context.Background())

	header, err := e.BlockHeaderByNumber(ctx, height)
	if err != nil {
		logger.Error("failed to get block header", "err", err)
		return err
	}
	s.Header = header

	err = e.IterateBlockTxs(ctx, height, func(tx *rpctypes.RPCTransaction) (bool, error) {
		cosmosTxHash, err := e.CosmosTxHashByTxHash(ctx, tx.Hash)
		if err != nil {
			logger.Error("failed to get cosmos tx hash", "err", err)
			return true, err
		}

		receipt, err := e.TxReceiptByHash(ctx, tx.Hash)
		if err != nil {
			logger.Error("failed to get tx receipt", "err", err)
			return true, err
		}

		s.Txs = append(s.Txs, tx)
		s.Receipts = append(s.Receipts, receipt)
		s.CosmosTxHashes = append(s.CosmosTxHashes, cosmosTxHash)
		return false, nil
	})
	if err != nil {
		logger.Error("failed to iterate block txs", "err", err)
		return err
	}

	bz, err := json.Marshal(s)
	if err != nil {
		logger.Error("failed to marshal snapshot extension", "err", err)
		return err
	}

	return payloadWriter(bz)
}

// SnapshotFormat implements types.ExtensionSnapshotter.
func (e *EVMIndexerImpl) SnapshotFormat() uint32 {
	return SnapshotFormat
}

// SnapshotName implements types.ExtensionSnapshotter.
func (e *EVMIndexerImpl) SnapshotName() string {
	return "evm_indexer"
}

// SupportedFormats implements types.ExtensionSnapshotter.
func (e *EVMIndexerImpl) SupportedFormats() []uint32 {
	return []uint32{SnapshotFormat}
}
