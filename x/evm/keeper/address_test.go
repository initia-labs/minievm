package keeper_test

import (
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/initia-labs/minievm/x/evm/keeper"
	"github.com/initia-labs/minievm/x/evm/types"
)

func (s *ERC20TestSuite) Test_AllowLongCosmosAddress() {
	evmAddr, addr, addr2 := s.evmAddrs[0], s.accAddrs[0], s.accAddrs[1]

	addr3 := append([]byte{0}, addr2.Bytes()...)
	addr4 := append([]byte{1}, addr2.Bytes()...)

	erc20Keeper, err := keeper.NewERC20Keeper(&s.input.EVMKeeper)
	s.Require().NoError(err)

	// deploy erc20 contract
	_, fooDenom := s.createContract("foo", evmAddr)

	// mint erc20
	s.mintERC20(evmAddr, evmAddr, sdk.NewCoin(fooDenom, math.NewInt(100)), false)

	// mint native coin
	err = erc20Keeper.MintCoins(s.ctx, addr, sdk.NewCoins(
		sdk.NewCoin("bar", math.NewInt(200)),
	))
	s.Require().NoError(err)

	// long address should be allowed
	err = erc20Keeper.SendCoins(s.ctx, addr, addr3, sdk.NewCoins(
		sdk.NewCoin("bar", math.NewInt(100)),
		sdk.NewCoin(fooDenom, math.NewInt(50)),
	))
	s.Require().NoError(err)

	// should be allowed because the address is not taken yet
	err = erc20Keeper.SendCoins(s.ctx, addr, addr4, sdk.NewCoins(
		sdk.NewCoin("bar", math.NewInt(100)),
		sdk.NewCoin(fooDenom, math.NewInt(50)),
	))
	s.Require().NoError(err)

	// take the address ownership
	err = erc20Keeper.SendCoins(s.ctx, addr3, addr, sdk.NewCoins(
		sdk.NewCoin("bar", math.NewInt(100)),
		sdk.NewCoin(fooDenom, math.NewInt(50)),
	))
	s.Require().NoError(err)

	// then other account can't use this address
	err = erc20Keeper.SendCoins(s.ctx, addr4, addr, sdk.NewCoins(
		sdk.NewCoin("bar", math.NewInt(100)),
		sdk.NewCoin(fooDenom, math.NewInt(50)),
	))
	s.Require().ErrorContains(err, types.ErrAddressAlreadyExists.Error())

	// also can't use the address as a receive
	err = erc20Keeper.SendCoins(s.ctx, addr, addr4, sdk.NewCoins(
		sdk.NewCoin("bar", math.NewInt(100)),
		sdk.NewCoin(fooDenom, math.NewInt(50)),
	))
	s.Require().ErrorContains(err, types.ErrAddressAlreadyExists.Error())
}

func (s *ERC20TestSuite) Test_AllowLongCosmosAddress_ConvertEmptyAccount() {
	evmAddr, evmAddr2, addr, addr2 := s.evmAddrs[0], s.evmAddrs[1], s.accAddrs[0], s.accAddrs[1]
	addr3 := append([]byte{0}, addr2.Bytes()...)

	// deploy erc20 contract
	_, fooDenom := s.createContract("foo", evmAddr)

	// mint erc20
	s.mintERC20(evmAddr, evmAddr, sdk.NewCoin(fooDenom, math.NewInt(100)), false)

	// create empty account
	s.mintERC20(evmAddr, evmAddr2, sdk.NewCoin(fooDenom, math.NewInt(100)), false)
	expectedAccNum := s.input.AccountKeeper.GetAccount(s.ctx, addr2).GetAccountNumber()

	// take the address ownership
	err := s.erc20Keeper.SendCoins(s.ctx, addr3, addr, sdk.NewCoins(
		sdk.NewCoin(fooDenom, math.NewInt(50)),
	))
	s.Require().NoError(err)

	// account number should be the same
	accNum := s.input.AccountKeeper.GetAccount(s.ctx, addr2).GetAccountNumber()
	s.Require().Equal(expectedAccNum, accNum)
}
