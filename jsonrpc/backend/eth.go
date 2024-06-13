package backend

import (
	"errors"

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

	feeDenom, decimals, err := b.feeDenomWithDecimals()
	if err != nil {
		return nil, err
	}

	balance, err := b.app.EVMKeeper.ERC20Keeper().GetBalance(queryCtx, sdk.AccAddress(address[:]), feeDenom)
	if err != nil {
		return nil, err
	}

	return (*hexutil.Big)(types.ToEthersUint(decimals, balance.BigInt())), nil
}

func (b *JSONRPCBackend) Call(args rpctypes.TransactionArgs, blockNrOrHash *rpc.BlockNumberOrHash, overrides *rpctypes.StateOverride, blockOverrides *rpctypes.BlockOverrides) (hexutil.Bytes, error) {
	if overrides != nil {
		return nil, errors.New("state overrides are not supported")
	}
	if blockOverrides != nil {
		return nil, errors.New("block overrides are not supported")
	}

	// if blockNrOrHash is nil, use the latest block
	if blockNrOrHash != nil {
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

	res, err := keeper.NewQueryServer(b.app.EVMKeeper).Call(queryCtx, &types.QueryCallRequest{
		Sender:       sender,
		ContractAddr: contractAddr,
		Input:        hexutil.Encode(args.GetData()),
	})

	if err != nil {
		return nil, err
	}

	if res.Error != "" {
		return nil, errors.New(res.Error)
	}

	return hexutil.MustDecode(res.Response), nil
}
