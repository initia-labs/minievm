package types

import (
	"context"

	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// AccountKeeper is expected keeper for auth module
type AccountKeeper interface {
	NewAccount(ctx context.Context, acc sdk.AccountI) sdk.AccountI
	GetAccount(ctx context.Context, addr sdk.AccAddress) sdk.AccountI
	SetAccount(ctx context.Context, acc sdk.AccountI)
	HasAccount(ctx context.Context, addr sdk.AccAddress) bool

	NewAccountWithAddress(ctx context.Context, addr sdk.AccAddress) sdk.AccountI
	NextAccountNumber(ctx context.Context) uint64
}

// BankKeeper is expected keeper for bank module
type BankKeeper interface {
	BlockedAddr(addr sdk.AccAddress) bool
}

type CommunityPoolKeeper interface {
	// FundCommunityPool allows an account to directly fund the community fund pool.
	FundCommunityPool(ctx context.Context, amount sdk.Coins, sender sdk.AccAddress) error
}

type IERC20StoresKeeper interface {
	Register(ctx context.Context, contractAddr common.Address) error
	RegisterFromFactory(ctx context.Context, caller, contractAddr common.Address) error
	RegisterStore(ctx context.Context, addr sdk.AccAddress, contractAddr common.Address) error
	IsStoreRegistered(ctx context.Context, addr sdk.AccAddress, contractAddr common.Address) (bool, error)
}

type IERC20Keeper interface {
	// balance
	GetBalance(ctx context.Context, addr sdk.AccAddress, denom string) (math.Int, error)
	GetPaginatedBalances(ctx context.Context, pageReq *query.PageRequest, addr sdk.AccAddress) (sdk.Coins, *query.PageResponse, error)
	GetPaginatedSupply(ctx context.Context, pageReq *query.PageRequest) (sdk.Coins, *query.PageResponse, error)
	IterateAccountBalances(ctx context.Context, addr sdk.AccAddress, cb func(sdk.Coin) (bool, error)) error
	IterateSupply(ctx context.Context, cb func(supply sdk.Coin) (bool, error)) error

	// operations
	SendCoins(ctx context.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) error
	MintCoins(ctx context.Context, addr sdk.AccAddress, amount sdk.Coins) error
	BurnCoins(ctx context.Context, addr sdk.AccAddress, amount sdk.Coins) error

	// supply
	GetSupply(ctx context.Context, denom string) (math.Int, error)
	HasSupply(ctx context.Context, denom string) (bool, error)

	// fungible asset
	GetMetadata(ctx context.Context, denom string) (banktypes.Metadata, error)

	// ABI
	GetERC20ABI() *abi.ABI

	// erc20 queries
	GetDecimals(ctx context.Context, denom string) (uint8, error)
}

type IERC721Keeper interface {
	CreateOrUpdateClass(ctx context.Context, classId, classUri, classData string) error
	Transfers(ctx context.Context, sender, escrowAddress sdk.AccAddress, classId string, tokenIds []string) error
	Burns(ctx context.Context, owner sdk.AccAddress, classId string, tokenIds []string) error
	Mints(ctx context.Context, receiver sdk.AccAddress, classId string, tokenIds, tokenUris []string, tokenData []string) error
	GetClassInfo(ctx context.Context, classId string) (className string, classUri string, classData string, err error)
	GetTokenInfos(ctx context.Context, classId string, tokenIds []string) (tokenUris []string, tokenData []string, err error)
}

type WithContext interface {
	WithContext(ctx context.Context) vm.PrecompiledContract
}

type GRPCRouter interface {
	Route(path string) baseapp.GRPCQueryHandler
}
