package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/legacy"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	// this line is used by starport scaffolding # 1
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	legacy.RegisterAminoMsg(cdc, &MsgCall{}, "evm/MsgCall")
	legacy.RegisterAminoMsg(cdc, &MsgCreate{}, "evm/MsgCreate")
	legacy.RegisterAminoMsg(cdc, &MsgCreate2{}, "evm/MsgCreate2")
	legacy.RegisterAminoMsg(cdc, &MsgUpdateParams{}, "evm/MsgUpdateParams")

	cdc.RegisterConcrete(&ContractAccount{}, "evm/ContractAccount", nil)
	cdc.RegisterConcrete(&ShorthandAccount{}, "evm/ShorthandAccount", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgCall{},
		&MsgCreate{},
		&MsgCreate2{},
		&MsgUpdateParams{},
	)

	// auth account registration
	registry.RegisterImplementations(
		(*sdk.AccountI)(nil),
		&ContractAccount{},
	)
	registry.RegisterImplementations(
		(*authtypes.GenesisAccount)(nil),
		&ContractAccount{},
	)
	registry.RegisterImplementations(
		(*sdk.AccountI)(nil),
		&ShorthandAccount{},
	)
	registry.RegisterImplementations(
		(*authtypes.GenesisAccount)(nil),
		&ShorthandAccount{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
