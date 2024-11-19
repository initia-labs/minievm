package keeper

import (
	"context"
	"sync/atomic"

	"cosmossdk.io/collections"
	"cosmossdk.io/core/address"
	corestoretypes "cosmossdk.io/core/store"
	"cosmossdk.io/log"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"

	evmconfig "github.com/initia-labs/minievm/x/evm/config"
	"github.com/initia-labs/minievm/x/evm/contracts/i_cosmos_callback"
	"github.com/initia-labs/minievm/x/evm/types"
)

type Keeper struct {
	ac           address.Codec
	cdc          codec.Codec
	storeService corestoretypes.KVStoreService

	accountKeeper       types.AccountKeeper
	bankKeeper          types.BankKeeper
	communityPoolKeeper types.CommunityPoolKeeper
	ibcHookKeeper       types.IBCHookKeeper
	gasPriceKeeper      types.GasPriceKeeper
	erc20Keeper         types.IERC20Keeper
	erc20StoresKeeper   types.IERC20StoresKeeper
	erc721Keeper        types.IERC721Keeper
	txUtils             types.TxUtils

	// grpc routers
	msgRouter  baseapp.MessageRouter
	grpcRouter types.GRPCRouter

	config evmconfig.EVMConfig

	// the address capable of executing a MsgUpdateParams message. Typically, this
	// should be the x/gov module account.
	authority string

	Schema  collections.Schema
	Params  collections.Item[types.Params]
	VMStore collections.Map[[]byte, []byte]

	// erc20 stores of users
	ERC20FactoryAddr          collections.Item[[]byte]
	ERC20WrapperAddr          collections.Item[[]byte]
	ERC20s                    collections.KeySet[[]byte]
	ERC20Stores               collections.KeySet[collections.Pair[[]byte, []byte]]
	ERC20DenomsByContractAddr collections.Map[[]byte, string]
	ERC20ContractAddrsByDenom collections.Map[string, []byte]

	// erc721 stores of users
	ERC721ClassURIs              collections.Map[[]byte, string]
	ERC721ClassIdsByContractAddr collections.Map[[]byte, string]
	ERC721ContractAddrsByClassId collections.Map[string, []byte]

	// evm stores
	EVMBlockHashes collections.Map[uint64, []byte]

	precompiles          precompiles
	queryCosmosWhitelist types.QueryCosmosWhitelist
	cosmosCallbackABI    *abi.ABI
}

func NewKeeper(
	ac address.Codec,
	cdc codec.Codec,
	storeService corestoretypes.KVStoreService,
	accountKeeper types.AccountKeeper,
	bankKeeper types.BankKeeper,
	communityPoolKeeper types.CommunityPoolKeeper,
	ibcHookKeeper types.IBCHookKeeper,
	gasPriceKeeper types.GasPriceKeeper,
	msgRouter baseapp.MessageRouter,
	grpcRouter *baseapp.GRPCQueryRouter,
	authority string,
	evmConfig evmconfig.EVMConfig,
	queryCosmosWhitelist types.QueryCosmosWhitelist,
) *Keeper {
	sb := collections.NewSchemaBuilder(storeService)

	if evmConfig.ContractSimulationGasLimit == 0 {
		evmConfig.ContractSimulationGasLimit = evmconfig.DefaultContractSimulationGasLimit
	}

	cosmosCallbackABI, err := i_cosmos_callback.ICosmosCallbackMetaData.GetAbi()
	if err != nil {
		panic(err)
	}

	execIndex := &atomic.Uint64{}
	execIndex.Store(0)
	k := &Keeper{
		ac:           ac,
		cdc:          cdc,
		storeService: storeService,

		accountKeeper:       accountKeeper,
		bankKeeper:          bankKeeper,
		communityPoolKeeper: communityPoolKeeper,
		ibcHookKeeper:       ibcHookKeeper,
		gasPriceKeeper:      gasPriceKeeper,
		authority:           authority,

		msgRouter:  msgRouter,
		grpcRouter: grpcRouter,

		config: evmConfig,

		Params:  collections.NewItem(sb, types.ParamsKey, "params", codec.CollValue[types.Params](cdc)),
		VMStore: collections.NewMap(sb, types.VMStorePrefix, "vm_store", collections.BytesKey, collections.BytesValue),

		ERC20WrapperAddr:          collections.NewItem(sb, types.ERC20WrapperAddrKey, "erc20_wrapper_addr", collections.BytesValue),
		ERC20FactoryAddr:          collections.NewItem(sb, types.ERC20FactoryAddrKey, "erc20_factory_addr", collections.BytesValue),
		ERC20s:                    collections.NewKeySet(sb, types.ERC20sPrefix, "erc20s", collections.BytesKey),
		ERC20Stores:               collections.NewKeySet(sb, types.ERC20StoresPrefix, "erc20_stores", collections.PairKeyCodec(collections.BytesKey, collections.BytesKey)),
		ERC20DenomsByContractAddr: collections.NewMap(sb, types.ERC20DenomsByContractAddrPrefix, "erc20_denoms_by_contract_addr", collections.BytesKey, collections.StringValue),
		ERC20ContractAddrsByDenom: collections.NewMap(sb, types.ERC20ContractAddrsByDenomPrefix, "erc20_contract_addrs_by_denom", collections.StringKey, collections.BytesValue),

		ERC721ClassURIs:              collections.NewMap(sb, types.ERC721ClassURIPrefix, "erc721_class_uris", collections.BytesKey, collections.StringValue),
		ERC721ClassIdsByContractAddr: collections.NewMap(sb, types.ERC721ClassIdsByContractAddrPrefix, "erc721_class_ids_by_contract_addr", collections.BytesKey, collections.StringValue),
		ERC721ContractAddrsByClassId: collections.NewMap(sb, types.ERC721ContractAddrsByClassIdPrefix, "erc721_contract_addrs_by_class_id", collections.StringKey, collections.BytesValue),

		EVMBlockHashes: collections.NewMap(sb, types.EVMBlockHashPrefix, "evm_block_hashes", collections.Uint64Key, collections.BytesValue),

		precompiles:          []precompile{},
		queryCosmosWhitelist: queryCosmosWhitelist,
		cosmosCallbackABI:    cosmosCallbackABI,
	}

	// setup schema
	schema, err := sb.Build()
	if err != nil {
		panic(err)
	}

	k.Schema = schema
	k.erc20StoresKeeper = NewERC20StoresKeeper(k)
	k.erc20Keeper, err = NewERC20Keeper(k)
	if err != nil {
		panic(err)
	}

	k.erc721Keeper, err = NewERC721Keeper(k)
	if err != nil {
		panic(err)
	}

	k.txUtils = NewTxUtils(k)

	// setup precompiles
	if err := k.loadPrecompiles(); err != nil {
		panic(err)
	}

	return k
}

// GetAuthority returns the x/evm module's authority.
func (ak Keeper) GetAuthority() string {
	return ak.authority
}

// Logger returns a module-specific logger.
func (k *Keeper) Logger(ctx context.Context) log.Logger {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	return sdkCtx.Logger().With("module", "x/"+types.ModuleName)
}

// StoreService returns the KVStoreService.
func (k Keeper) StoreService() corestoretypes.KVStoreService {
	return k.storeService
}

// Config returns the x/evm configuration.
func (k Keeper) Config() evmconfig.EVMConfig {
	return k.config
}

// ERC20Keeper returns the ERC20Keeper
func (k Keeper) ERC20Keeper() types.IERC20Keeper {
	return k.erc20Keeper
}

// ERC20StoresKeeper returns the ERC20StoresKeeper
func (k Keeper) ERC20StoresKeeper() types.IERC20StoresKeeper {
	return k.erc20StoresKeeper
}

// ERC721Keeper returns the ERC721Keeper
func (k Keeper) ERC721Keeper() types.IERC721Keeper {
	return k.erc721Keeper
}

// TxUtils returns the TxUtils
func (k Keeper) TxUtils() types.TxUtils {
	return k.txUtils
}

// GetContractAddrByDenom returns contract address by denom
func (k Keeper) GetContractAddrByDenom(ctx context.Context, denom string) (common.Address, error) {
	bz, err := k.ERC20ContractAddrsByDenom.Get(ctx, denom)
	if err != nil {
		return common.Address{}, err
	}

	return common.BytesToAddress(bz), nil
}

// GetDenomByContractAddr returns denom by contract address
func (k Keeper) GetDenomByContractAddr(ctx context.Context, contractAddr common.Address) (string, error) {
	return k.ERC20DenomsByContractAddr.Get(ctx, contractAddr.Bytes())
}

// GetContractAddrByDenom returns contract address by denom
func (k Keeper) GetContractAddrByClassId(ctx context.Context, classId string) (common.Address, error) {
	bz, err := k.ERC721ContractAddrsByClassId.Get(ctx, classId)
	if err != nil {
		return common.Address{}, err
	}

	return common.BytesToAddress(bz), nil
}

// GetDenomByContractAddr returns denom by contract address
func (k Keeper) GetClassIdByContractAddr(ctx context.Context, contractAddr common.Address) (string, error) {
	return k.ERC721ClassIdsByContractAddr.Get(ctx, contractAddr.Bytes())
}

func (k Keeper) GetERC20FactoryAddr(ctx context.Context) (common.Address, error) {
	factoryAddr, err := k.ERC20FactoryAddr.Get(ctx)
	if err != nil {
		return common.Address{}, types.ErrFailedToGetERC20FactoryAddr.Wrap(err.Error())
	}

	return common.BytesToAddress(factoryAddr), nil
}

func (k Keeper) GetERC20WrapperAddr(ctx context.Context) (common.Address, error) {
	wrapperAddr, err := k.ERC20WrapperAddr.Get(ctx)
	if err != nil {
		return common.Address{}, types.ErrFailedToGetERC20WrapperAddr.Wrap(err.Error())
	}

	return common.BytesToAddress(wrapperAddr), nil
}

// keep track recent `NumRetainBlockHashes` block hashes
// - https://www.ethervm.io/#40 (default action is keep `256“ block hashes)
func (k Keeper) TrackBlockHash(ctx context.Context, blockHeight uint64, hash common.Hash) error {
	num, err := k.NumRetainBlockHashes(ctx)
	if err != nil {
		return err
	}

	// keep all block hashes if `NumRetainBlockHashes` is 0
	if num != 0 && blockHeight > num {
		err := k.EVMBlockHashes.Remove(ctx, blockHeight-num)
		if err != nil {
			return err
		}
	}

	return k.EVMBlockHashes.Set(ctx, blockHeight, hash.Bytes())
}
