package app

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/module"
	icacontrollertypes "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/controller/types"
	icagenesistypes "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/genesis/types"
	icahosttypes "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/host/types"
	icatypes "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/types"
	ibcexported "github.com/cosmos/ibc-go/v8/modules/core/exported"
	ibctypes "github.com/cosmos/ibc-go/v8/modules/core/types"

	opchildtypes "github.com/initia-labs/OPinit/x/opchild/types"

	auctiontypes "github.com/skip-mev/block-sdk/x/auction/types"
)

// GenesisState - The genesis state of the blockchain is represented here as a map of raw json
// messages key'd by a identifier string.
// The identifier is used to determine which module genesis information belongs
// to so it may be appropriately routed during init chain.
// Within this application default genesis information is retrieved from
// the ModuleBasicManager which populates json from each BasicModule
// object provided to it during init.
type GenesisState map[string]json.RawMessage

// NewDefaultGenesisState generates the default state for the application.
func NewDefaultGenesisState(cdc codec.JSONCodec, mbm module.BasicManager, denom string) GenesisState {
	return GenesisState(mbm.DefaultGenesis(cdc)).
		ConfigureMinGasPrices(cdc).
		ConfigureICA(cdc).
		ConfigureIBCAllowedClients(cdc).
		ConfigureAuctionFee(cdc, denom)
}

func (genState GenesisState) ConfigureAuctionFee(cdc codec.JSONCodec, denom string) GenesisState {
	var auctionGenState auctiontypes.GenesisState
	cdc.MustUnmarshalJSON(genState[auctiontypes.ModuleName], &auctionGenState)
	auctionGenState.Params.ReserveFee.Denom = denom
	auctionGenState.Params.MinBidIncrement.Denom = denom
	genState[auctiontypes.ModuleName] = cdc.MustMarshalJSON(&auctionGenState)

	return genState
}

// ConfigureMinGasPrices generates the default state for the application.
func (genState GenesisState) ConfigureMinGasPrices(cdc codec.JSONCodec) GenesisState {
	var opChildGenState opchildtypes.GenesisState
	cdc.MustUnmarshalJSON(genState[opchildtypes.ModuleName], &opChildGenState)
	opChildGenState.Params.MinGasPrices = nil
	genState[opchildtypes.ModuleName] = cdc.MustMarshalJSON(&opChildGenState)

	return genState
}

func (genState GenesisState) ConfigureICA(cdc codec.JSONCodec) GenesisState {
	// create ICS27 Controller submodule params
	controllerParams := icacontrollertypes.Params{
		ControllerEnabled: true,
	}

	// create ICS27 Host submodule params
	hostParams := icahosttypes.Params{
		HostEnabled: true,
		AllowMessages: []string{
			authzMsgExec,
			authzMsgGrant,
			authzMsgRevoke,
			bankMsgSend,
			bankMsgMultiSend,
			feegrantMsgGrantAllowance,
			feegrantMsgRevokeAllowance,
			groupCreateGroup,
			groupCreateGroupPolicy,
			groupExec,
			groupLeaveGroup,
			groupSubmitProposal,
			groupUpdateGroupAdmin,
			groupUpdateGroupMember,
			groupUpdateGroupPolicyAdmin,
			groupUpdateGroupPolicyDecisionPolicy,
			groupVote,
			groupWithdrawProposal,
			transferMsgTransfer,
			nftTransferMsgTransfer,
			sftTransferMsgTransfer,
			moveMsgPublishModuleBundle,
			moveMsgExecuteEntryFunction,
			moveMsgExecuteScript,
		},
	}

	var icaGenState icagenesistypes.GenesisState
	cdc.MustUnmarshalJSON(genState[icatypes.ModuleName], &icaGenState)
	icaGenState.ControllerGenesisState.Params = controllerParams
	icaGenState.HostGenesisState.Params = hostParams
	genState[icatypes.ModuleName] = cdc.MustMarshalJSON(&icaGenState)

	return genState
}

func (genState GenesisState) ConfigureIBCAllowedClients(cdc codec.JSONCodec) GenesisState {
	var ibcGenesis ibctypes.GenesisState
	cdc.MustUnmarshalJSON(genState[ibcexported.ModuleName], &ibcGenesis)

	allowedClients := ibcGenesis.ClientGenesis.Params.AllowedClients
	for i, client := range allowedClients {
		if client == ibcexported.Localhost {
			allowedClients = append(allowedClients[:i], allowedClients[i+1:]...)
			break
		}
	}

	ibcGenesis.ClientGenesis.Params.AllowedClients = allowedClients
	genState[ibcexported.ModuleName] = cdc.MustMarshalJSON(&ibcGenesis)

	return genState
}
