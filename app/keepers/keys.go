package keepers

import (
	storetypes "cosmossdk.io/store/types"
	"cosmossdk.io/x/feegrant"
	upgradetypes "cosmossdk.io/x/upgrade/types"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	authzkeeper "github.com/cosmos/cosmos-sdk/x/authz/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	consensusparamtypes "github.com/cosmos/cosmos-sdk/x/consensus/types"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	"github.com/cosmos/cosmos-sdk/x/group"

	// ibc imports
	packetforwardtypes "github.com/cosmos/ibc-apps/middleware/packet-forward-middleware/v8/packetforward/types"
	ratelimittypes "github.com/cosmos/ibc-apps/modules/rate-limiting/v8/types"
	capabilitytypes "github.com/cosmos/ibc-go/modules/capability/types"
	icacontrollertypes "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/controller/types"
	icahosttypes "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/host/types"
	ibcfeetypes "github.com/cosmos/ibc-go/v8/modules/apps/29-fee/types"
	ibctransfertypes "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"
	ibcexported "github.com/cosmos/ibc-go/v8/modules/core/exported"

	// initia imports
	ibchookstypes "github.com/initia-labs/initia/x/ibc-hooks/types"
	ibcnfttransfertypes "github.com/initia-labs/initia/x/ibc/nft-transfer/types"
	icaauthtypes "github.com/initia-labs/initia/x/intertx/types"

	// OPinit imports
	opchildtypes "github.com/initia-labs/OPinit/x/opchild/types"

	// skip imports
	auctiontypes "github.com/skip-mev/block-sdk/v2/x/auction/types"
	marketmaptypes "github.com/skip-mev/connect/v2/x/marketmap/types"
	oracletypes "github.com/skip-mev/connect/v2/x/oracle/types"

	// noble forwarding keeper
	forwardingtypes "github.com/noble-assets/forwarding/v2/types"

	// local imports
	evmtypes "github.com/initia-labs/minievm/x/evm/types"
)

// KVStoreKeys returns a list of the application's KV store keys.
func KVStoreKeys() []string {
	return []string{
		authtypes.StoreKey, banktypes.StoreKey, group.StoreKey, consensusparamtypes.StoreKey,
		crisistypes.StoreKey, ibcexported.StoreKey, upgradetypes.StoreKey,
		ibctransfertypes.StoreKey, ibcnfttransfertypes.StoreKey,
		capabilitytypes.StoreKey, authzkeeper.StoreKey, feegrant.StoreKey,
		icahosttypes.StoreKey, icacontrollertypes.StoreKey, icaauthtypes.StoreKey,
		ibcfeetypes.StoreKey, evmtypes.StoreKey, opchildtypes.StoreKey,
		auctiontypes.StoreKey, packetforwardtypes.StoreKey, ratelimittypes.StoreKey,
		oracletypes.StoreKey, marketmaptypes.StoreKey, ibchookstypes.StoreKey, forwardingtypes.StoreKey,
	}
}

// GenerateKeys generates the keys for the application.
func (appKeepers *AppKeepers) GenerateKeys() {
	// Define what keys will be used in the cosmos-sdk key/value store.
	// Cosmos-SDK modules each have a "key" that allows the application to reference what they've stored on the chain.
	appKeepers.keys = storetypes.NewKVStoreKeys(
		KVStoreKeys()...,
	)

	// Define transient store keys
	appKeepers.tkeys = storetypes.NewTransientStoreKeys(forwardingtypes.TransientStoreKey, ibchookstypes.TStoreKey)

	// MemKeys are for information that is stored only in RAM.
	appKeepers.memKeys = storetypes.NewMemoryStoreKeys(capabilitytypes.MemStoreKey)
}

func (appKeepers *AppKeepers) GetKVStoreKey() map[string]*storetypes.KVStoreKey {
	return appKeepers.keys
}

func (appKeepers *AppKeepers) GetTransientStoreKey() map[string]*storetypes.TransientStoreKey {
	return appKeepers.tkeys
}

func (appKeepers *AppKeepers) GetMemoryStoreKey() map[string]*storetypes.MemoryStoreKey {
	return appKeepers.memKeys
}

// GetKey returns the KVStoreKey for the provided store key.
//
// NOTE: This is solely to be used for testing purposes.
func (appKeepers *AppKeepers) GetKey(storeKey string) *storetypes.KVStoreKey {
	return appKeepers.keys[storeKey]
}

// GetTKey returns the TransientStoreKey for the provided store key.
//
// NOTE: This is solely to be used for testing purposes.
func (appKeepers *AppKeepers) GetTKey(storeKey string) *storetypes.TransientStoreKey {
	return appKeepers.tkeys[storeKey]
}

// GetMemKey returns the MemStoreKey for the provided mem key.
//
// NOTE: This is solely used for testing purposes.
func (appKeepers *AppKeepers) GetMemKey(storeKey string) *storetypes.MemoryStoreKey {
	return appKeepers.memKeys[storeKey]
}
