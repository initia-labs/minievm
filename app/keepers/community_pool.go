package keepers

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type bankKeeperForCommunityPoolKeeper interface {
	SendCoinsFromAccountToModule(ctx context.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
}

type CommunityPoolKeeper struct {
	bk               bankKeeperForCommunityPoolKeeper
	feeCollectorName string
}

func NewCommunityPoolKeeper(bk bankKeeperForCommunityPoolKeeper, feeCollectorName string) CommunityPoolKeeper {
	return CommunityPoolKeeper{
		bk:               bk,
		feeCollectorName: feeCollectorName,
	}
}

func (k CommunityPoolKeeper) FundCommunityPool(ctx context.Context, amount sdk.Coins, sender sdk.AccAddress) error {
	return k.bk.SendCoinsFromAccountToModule(ctx, sender, k.feeCollectorName, amount)
}
