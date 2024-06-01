package backend

import (
	"context"
	"errors"
	"math/big"

	tmrpcclient "github.com/cometbft/cometbft/rpc/client"
	tmtypes "github.com/cometbft/cometbft/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	coretypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"

	rpctypes "github.com/initia-labs/minievm/jsonrpc/types"
	"github.com/initia-labs/minievm/x/evm/types"
)

func (b *JSONRPCBackend) BlockNumber(ctx context.Context) (hexutil.Uint64, error) {
	res, err := b.clientCtx.Client.Status(ctx)
	if err != nil {
		return 0, err
	}

	return hexutil.Uint64(res.SyncInfo.LatestBlockHeight), nil
}

func (b *JSONRPCBackend) GetBlockByNumber(ctx context.Context, ethBlockNum rpc.BlockNumber, fullTx bool) (map[string]interface{}, error) {
	blockNum := ethBlockNum.Int64()
	tmBlock, err := b.clientCtx.Client.Block(ctx, &blockNum)
	if err != nil {
		return nil, err
	}

	return b.convertTmBlockToEthBlock(ctx, tmBlock.Block)
}

func (b *JSONRPCBackend) GetBlockByHash(ctx context.Context, hash common.Hash, fullTx bool) (map[string]interface{}, error) {
	sc, ok := b.clientCtx.Client.(tmrpcclient.SignClient)
	if !ok {
		return nil, errors.New("invalid rpc client")
	}

	tmBlock, err := sc.BlockByHash(ctx, hash.Bytes())
	if err != nil {
		b.svrCtx.Logger.Debug("tendermint client failed to get block", "blockHash", hash.Hex(), "error", err.Error())
		return nil, err
	}

	return b.convertTmBlockToEthBlock(ctx, tmBlock.Block)
}

func (b *JSONRPCBackend) GetBlockTransactionCountByHash(ctx context.Context, hash common.Hash) (*hexutil.Uint, error) {
	b.GetBlockByHash(ctx, hash, false)
	sc, ok := b.clientCtx.Client.(tmrpcclient.SignClient)
	if !ok {
		return nil, errors.New("invalid rpc client")
	}

}

// formatBlock creates an ethereum block from a tendermint header and ethereum-formatted
// transactions.
//
// TODO: bloom filter, logs, and receipts if there is a need to support them.
func formatBlock(
	header tmtypes.Header, size int, gasLimit int64,
	gasUsed *big.Int, transactions []*rpctypes.RPCTransaction,
	validatorAddr common.Address, baseFee *big.Int,
) map[string]interface{} {
	var transactionsRoot common.Hash
	if len(transactions) == 0 {
		transactionsRoot = coretypes.EmptyRootHash
	} else {
		transactionsRoot = common.BytesToHash(header.DataHash)
	}

	result := map[string]interface{}{
		"number":           hexutil.Uint64(header.Height),
		"hash":             hexutil.Bytes(header.Hash()),
		"parentHash":       common.BytesToHash(header.LastBlockID.Hash.Bytes()),
		"nonce":            coretypes.BlockNonce{},   // PoW specific
		"sha3Uncles":       coretypes.EmptyUncleHash, // No uncles in Tendermint
		"stateRoot":        hexutil.Bytes(header.AppHash),
		"miner":            validatorAddr,
		"mixHash":          common.Hash{},
		"difficulty":       (*hexutil.Big)(big.NewInt(0)),
		"extraData":        "0x",
		"size":             hexutil.Uint64(size),
		"gasLimit":         hexutil.Uint64(gasLimit), // Static gas limit
		"gasUsed":          (*hexutil.Big)(gasUsed),
		"timestamp":        hexutil.Uint64(header.Time.Unix()),
		"transactionsRoot": transactionsRoot,
		"receiptsRoot":     coretypes.EmptyRootHash,

		"uncles":          []common.Hash{},
		"transactions":    transactions,
		"totalDifficulty": (*hexutil.Big)(big.NewInt(0)),
	}

	if baseFee != nil {
		result["baseFeePerGas"] = (*hexutil.Big)(baseFee)
	}

	return result
}

// blockMaxGasFromConsensusParams returns the gas limit for the current block from the chain consensus params.
func (s *JSONRPCBackend) blockMaxGasFromConsensusParams(goCtx context.Context, blockHeight int64) (int64, error) {
	tmrpcClient, ok := s.clientCtx.Client.(tmrpcclient.Client)
	if !ok {
		return 0, errors.New("incorrect tm rpc client")
	}

	resConsParams, err := tmrpcClient.ConsensusParams(goCtx, &blockHeight)
	defaultGasLimit := int64(^uint32(0)) // #nosec G701
	if err != nil {
		return defaultGasLimit, err
	}

	gasLimit := resConsParams.ConsensusParams.Block.MaxGas
	if gasLimit == -1 {
		// Sets gas limit to max uint32 to not error with javascript dev tooling
		// This -1 value indicating no block gas limit is set to max uint64 with geth hexutils
		// which errors certain javascript dev tooling which only supports up to 53 bits
		gasLimit = defaultGasLimit
	}

	return gasLimit, nil
}

func (b *JSONRPCBackend) convertTmBlockToEthBlock(ctx context.Context, block *tmtypes.Block) (map[string]interface{}, error) {
	chainID, _ := types.ConvertCosmosChainIDToEthereumChainID(b.clientCtx.ChainID)
	gasLimit, err := b.blockMaxGasFromConsensusParams(ctx, block.Height)
	if err != nil {
		b.svrCtx.Logger.Error("failed to query consensus params", "error", err.Error())
	}

	gasUsed := uint64(0)

	rpcTxs := []*rpctypes.RPCTransaction{}
	decoder := b.clientCtx.TxConfig.TxDecoder()
	for index, txBz := range block.Txs {
		tx, err := decoder(txBz)
		if err != nil {
			continue
		}

		// update gas used
		gasTx := tx.(sdk.FeeTx)
		gasUsed += gasTx.GetGas()

		ethTx, err := b.ConvertCosmosTxToEthereumTx(tx)
		if ethTx == nil || err != nil {
			continue
		}

		rpcTxs = append(rpcTxs, newRPCTransaction(
			ethTx,
			common.BytesToHash(block.Hash()),
			uint64(block.Height),
			uint64(index),
			chainID,
		))
	}

	return formatBlock(
		block.Header,
		block.Size(),
		gasLimit,
		new(big.Int).SetUint64(gasUsed),
		rpcTxs,
		common.BytesToAddress(block.Header.ProposerAddress),
		nil,
	), nil
}
