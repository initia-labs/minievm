package backend

import (
	"errors"
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rpc"

	rpctypes "github.com/initia-labs/minievm/jsonrpc/types"
	"github.com/initia-labs/minievm/x/evm/keeper"
	"github.com/initia-labs/minievm/x/evm/types"
)

func (b *JSONRPCBackend) GetBalance(address common.Address, blockNrOrHash rpc.BlockNumberOrHash) (*hexutil.Big, error) {
	blockNumber, err := b.resolveBlockNrOrHash(blockNrOrHash)
	if err != nil {
		return nil, err
	}

	queryCtx, err := b.getQueryCtxWithHeight(blockNumber)
	if err != nil {
		return nil, err
	}

	feeDenom, feeDecimals, err := b.feeInfo()
	if err != nil {
		return nil, err
	}

	balance, err := b.app.EVMKeeper.ERC20Keeper().GetBalance(queryCtx, sdk.AccAddress(address[:]), feeDenom)
	if err != nil {
		return nil, err
	}

	return (*hexutil.Big)(types.ToEthersUnit(feeDecimals, balance.BigInt())), nil
}

func (b *JSONRPCBackend) Call(args rpctypes.TransactionArgs, blockNrOrHash *rpc.BlockNumberOrHash, overrides *rpctypes.StateOverride, blockOverrides *rpctypes.BlockOverrides) (hexutil.Bytes, error) {
	if overrides != nil && len(*overrides) > 0 {
		return nil, errors.New("state overrides are not supported")
	}
	if blockOverrides != nil {
		return nil, errors.New("block overrides are not supported")
	}

	// if blockNrOrHash is nil, use the latest block
	if blockNrOrHash == nil {
		latest := rpc.BlockNumberOrHashWithNumber(rpc.LatestBlockNumber)
		blockNrOrHash = &latest
	}

	blockNumber, err := b.resolveBlockNrOrHash(*blockNrOrHash)
	if err != nil {
		return nil, err
	}

	queryCtx, err := b.getQueryCtxWithHeight(blockNumber)
	if err != nil {
		return nil, err
	}

	// set call defaults
	args.CallDefaults()

	// convert sender to string
	sender := ""
	if args.From != nil {
		senderStr, err := b.app.AccountKeeper.AddressCodec().BytesToString(args.From[:])
		if err != nil {
			return nil, err
		}

		sender = senderStr
	}

	contractAddr := ""
	if args.To != nil {
		contractAddr = args.To.Hex()
	}

	var list []types.AccessTuple
	if args.AccessList != nil {
		list = types.ConvertEthAccessListToCosmos(*args.AccessList)
	}

	res, err := keeper.NewQueryServer(b.app.EVMKeeper).Call(queryCtx, &types.QueryCallRequest{
		Sender:       sender,
		ContractAddr: contractAddr,
		Input:        hexutil.Encode(args.GetData()),
		AccessList:   list,
	})

	if err != nil {
		return nil, err
	}

	if res.Error != "" {
		return nil, errors.New(res.Error)
	}

	return hexutil.MustDecode(res.Response), nil
}

func (b *JSONRPCBackend) GetStorageAt(address common.Address, key common.Hash, blockNrOrHash rpc.BlockNumberOrHash) (hexutil.Bytes, error) {
	blockNumber, err := b.resolveBlockNrOrHash(blockNrOrHash)
	if err != nil {
		return nil, err
	}

	queryCtx, err := b.getQueryCtxWithHeight(blockNumber)
	if err != nil {
		return nil, err
	}

	res, err := keeper.NewQueryServer(b.app.EVMKeeper).State(queryCtx, &types.QueryStateRequest{
		ContractAddr: address.Hex(),
		Key:          key.Hex(),
	})

	if err != nil {
		return nil, err
	}

	return hexutil.MustDecode(res.Value), nil
}

func (b *JSONRPCBackend) GetCode(address common.Address, blockNrOrHash rpc.BlockNumberOrHash) (hexutil.Bytes, error) {
	blockNumber, err := b.resolveBlockNrOrHash(blockNrOrHash)
	if err != nil {
		return nil, err
	}

	queryCtx, err := b.getQueryCtxWithHeight(blockNumber)
	if err != nil {
		return nil, err
	}

	res, err := keeper.NewQueryServer(b.app.EVMKeeper).Code(queryCtx, &types.QueryCodeRequest{
		ContractAddr: address.Hex(),
	})

	if err != nil {
		return nil, err
	}

	return hexutil.MustDecode(res.Code), nil
}

func (b *JSONRPCBackend) ChainId() (*hexutil.Big, error) {
	chainID, err := b.ChainID()
	return (*hexutil.Big)(chainID), err
}

func (b *JSONRPCBackend) ChainID() (*big.Int, error) {
	queryCtx, err := b.getQueryCtx()
	if err != nil {
		return nil, err
	}

	sdkCtx := sdk.UnwrapSDKContext(queryCtx)
	return types.ConvertCosmosChainIDToEthereumChainID(sdkCtx.ChainID()), nil
}

func (b *JSONRPCBackend) Syncing() (interface{}, error) {
	status, err := b.clientCtx.Client.Status(b.ctx)
	if err != nil {
		return nil, err
	}

	if !status.SyncInfo.CatchingUp {
		return false, nil
	}

	latestHeight := status.SyncInfo.LatestBlockHeight

	// Otherwise gather the block sync stats
	return map[string]interface{}{
		"startingBlock":          hexutil.Uint64(status.SyncInfo.EarliestBlockHeight),
		"currentBlock":           hexutil.Uint64(latestHeight),
		"highestBlock":           hexutil.Uint64(latestHeight),
		"syncedAccounts":         hexutil.Uint64(latestHeight),
		"syncedAccountBytes":     hexutil.Uint64(latestHeight),
		"syncedBytecodes":        hexutil.Uint64(latestHeight),
		"syncedBytecodeBytes":    hexutil.Uint64(latestHeight),
		"syncedStorage":          hexutil.Uint64(latestHeight),
		"syncedStorageBytes":     hexutil.Uint64(latestHeight),
		"healedTrienodes":        hexutil.Uint64(latestHeight),
		"healedTrienodeBytes":    hexutil.Uint64(latestHeight),
		"healedBytecodes":        hexutil.Uint64(latestHeight),
		"healedBytecodeBytes":    hexutil.Uint64(latestHeight),
		"healingTrienodes":       hexutil.Uint64(latestHeight),
		"healingBytecode":        hexutil.Uint64(latestHeight),
		"txIndexFinishedBlocks":  hexutil.Uint64(latestHeight),
		"txIndexRemainingBlocks": hexutil.Uint64(latestHeight),
	}, nil
}
