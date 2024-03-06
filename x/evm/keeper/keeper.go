package keeper

import (
	"context"

	"cosmossdk.io/collections"
	"cosmossdk.io/core/address"
	corestoretypes "cosmossdk.io/core/store"
	"cosmossdk.io/log"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	evmconfig "github.com/initia-labs/minievm/x/evm/config"
	"github.com/initia-labs/minievm/x/evm/types"
)

type Keeper struct {
	ac           address.Codec
	cdc          codec.Codec
	storeService corestoretypes.KVStoreService

	config evmconfig.EVMConfig

	// the address capable of executing a MsgUpdateParams message. Typically, this
	// should be the x/gov module account.
	authority string

	Schema  collections.Schema
	Params  collections.Item[types.Params]
	VMRoot  collections.Item[[]byte]
	VMStore collections.Map[[]byte, []byte]
}

func NewKeeper(
	ac address.Codec,
	cdc codec.Codec,
	storeService corestoretypes.KVStoreService,
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
		config:       EVMConfig,

		Params:  collections.NewItem(sb, types.ParamsKey, "params", codec.CollValue[types.Params](cdc)),
		VMRoot:  collections.NewItem(sb, types.VMRootKey, "vm_root", collections.BytesValue),
		VMStore: collections.NewMap(sb, types.VMStorePrefix, "vm_store", collections.BytesKey, collections.BytesValue),
	}

	schema, err := sb.Build()
	if err != nil {
		panic(err)
	}
	k.Schema = schema
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
