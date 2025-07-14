package types

import (
	"context"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/codec"
	"github.com/cosmos/cosmos-sdk/x/authz"
)

const (
	gasCostPerIteration = uint64(10)
)

var (
	_ authz.Authorization = &CallAuthorization{}
)

func NewCallAuthorization(contracts []string) *CallAuthorization {
	return &CallAuthorization{
		Contracts: contracts,
	}
}

// MsgTypeURL implements Authorization.MsgTypeURL.
func (a CallAuthorization) MsgTypeURL() string {
	return sdk.MsgTypeURL(&MsgCall{})
}

func (a CallAuthorization) ValidateBasic() error {
	if err := validateChecksumHexAddrs(a.Contracts); err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid contract address: %s", err)
	}
	return nil
}

func (a CallAuthorization) Accept(ctx context.Context, msg sdk.Msg) (authz.AcceptResponse, error) {
	switch msg := msg.(type) {
	case *MsgCall:
		if len(a.Contracts) == 0 {
			return authz.AcceptResponse{Accept: true}, nil
		}

		sdkCtx := sdk.UnwrapSDKContext(ctx)

		// TODO - cannot retrieve address codec here
		ac := codec.NewBech32Codec(sdk.GetConfig().GetBech32AccountAddrPrefix())
		contractAddr, err := ContractAddressFromString(ac, msg.ContractAddr)
		if err != nil {
			return authz.AcceptResponse{}, sdkerrors.ErrUnauthorized.Wrap("invalid contract address")
		}

		for _, contract := range a.Contracts {
			sdkCtx.GasMeter().ConsumeGas(gasCostPerIteration, "decode contract address")
			if contractAddr.Hex() == contract {
				return authz.AcceptResponse{Accept: true}, nil
			}
		}
		return authz.AcceptResponse{}, sdkerrors.ErrUnauthorized.Wrap("unauthorized")
	default:
		return authz.AcceptResponse{}, sdkerrors.ErrInvalidRequest.Wrap("unknown msg type")
	}
}
