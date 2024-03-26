package keeper

import (
	"context"

	"cosmossdk.io/collections"
	"cosmossdk.io/core/address"
	corestoretypes "cosmossdk.io/core/store"
	"cosmossdk.io/log"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"

	evmconfig "github.com/initia-labs/minievm/x/evm/config"
	"github.com/initia-labs/minievm/x/evm/types"
)

type Keeper struct {
	ac           address.Codec
	cdc          codec.Codec
	storeService corestoretypes.KVStoreService

	accountKeeper       types.AccountKeeper
	communityPoolKeeper types.CommunityPoolKeeper
	erc20Keeper         types.IERC20Keeper
	erc20StoresKeeper   types.IERC20StoresKeeper

	config evmconfig.EVMConfig

	// the address capable of executing a MsgUpdateParams message. Typically, this
	// should be the x/gov module account.
	authority string

	Schema  collections.Schema
	Params  collections.Item[types.Params]
	VMRoot  collections.Item[[]byte]
	VMStore collections.Map[[]byte, []byte]

	// erc20 stores of users
	ERC20s                    collections.KeySet[[]byte]
	ERC20Stores               collections.KeySet[collections.Pair[[]byte, []byte]]
	ERC20DenomsByContractAddr collections.Map[[]byte, string]
	ERC20ContractAddrsByDenom collections.Map[string, []byte]

	precompiles precompiles
}

func NewKeeper(
	ac address.Codec,
	cdc codec.Codec,
	storeService corestoretypes.KVStoreService,
	accountKeeper types.AccountKeeper,
	communityPoolKeeper types.CommunityPoolKeeper,
	authority string,
	EVMConfig evmconfig.EVMConfig,
) *Keeper {
	sb := collections.NewSchemaBuilder(storeService)

	if EVMConfig.ContractSimulationGasLimit == 0 {
		EVMConfig.ContractSimulationGasLimit = evmconfig.DefaultContractSimulationGasLimit
	}

	if EVMConfig.ContractQueryGasLimit == 0 {
		EVMConfig.ContractQueryGasLimit = evmconfig.DefaultContractQueryGasLimit
	}

	k := &Keeper{
		ac:           ac,
		cdc:          cdc,
		storeService: storeService,

		accountKeeper:       accountKeeper,
		communityPoolKeeper: communityPoolKeeper,

		config: EVMConfig,

		Params:  collections.NewItem(sb, types.ParamsKey, "params", codec.CollValue[types.Params](cdc)),
		VMRoot:  collections.NewItem(sb, types.VMRootKey, "vm_root", collections.BytesValue),
		VMStore: collections.NewMap(sb, types.VMStorePrefix, "vm_store", collections.BytesKey, collections.BytesValue),

		ERC20s:                    collections.NewKeySet(sb, types.ERC20sPrefix, "erc20s", collections.BytesKey),
		ERC20Stores:               collections.NewKeySet(sb, types.ERC20StoresPrefix, "erc20_stores", collections.PairKeyCodec(collections.BytesKey, collections.BytesKey)),
		ERC20DenomsByContractAddr: collections.NewMap(sb, types.ERC20DenomsByContractAddrPrefix, "erc20_denoms_by_contract_addr", collections.BytesKey, collections.StringValue),
		ERC20ContractAddrsByDenom: collections.NewMap(sb, types.ERC20ContractAddrsByDenomPrefix, "erc20_contract_addrs_by_denom", collections.StringKey, collections.BytesValue),

		precompiles: []precompile{},
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

	// setup precompiles
	if err := k.loadPrecompiles(); err != nil {
		panic(err)
	}

	return k
}

// GetAuthority returns the x/move module's authority.
func (ak Keeper) GetAuthority() string {
	return ak.authority
}

// Logger returns a module-specific logger.
func (k *Keeper) Logger(ctx context.Context) log.Logger {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	return sdkCtx.Logger().With("module", "x/"+types.ModuleName)
}

// ERC20Keeper returns the ERC20Keeper
func (k Keeper) ERC20Keeper() types.IERC20Keeper {
	return k.erc20Keeper
}

// ERC20StoresKeeper returns the ERC20StoresKeeper
func (k Keeper) ERC20StoresKeeper() types.IERC20StoresKeeper {
	return k.erc20StoresKeeper
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
