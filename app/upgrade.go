package app

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"

	upgradetypes "cosmossdk.io/x/upgrade/types"
	"github.com/cosmos/cosmos-sdk/types/module"

	"github.com/initia-labs/minievm/x/evm/contracts/erc20_factory"
	"github.com/initia-labs/minievm/x/evm/contracts/erc20_wrapper"
	"github.com/initia-labs/minievm/x/evm/state"
	evmtypes "github.com/initia-labs/minievm/x/evm/types"
)

const upgradeName = "v1.1.0"

// RegisterUpgradeHandlers returns upgrade handlers
func (app *MinitiaApp) RegisterUpgradeHandlers(cfg module.Configurator) {
	app.UpgradeKeeper.SetUpgradeHandler(
		upgradeName,
		func(ctx context.Context, _ upgradetypes.Plan, versionMap module.VersionMap) (module.VersionMap, error) {

			// 1. update erc20 wrapper contract
			wrapperAddr, err := app.EVMKeeper.GetERC20WrapperAddr(ctx)
			if err != nil {
				return nil, err
			}
			wrapperCode, wrapperCodeHash := extractRuntimeCodeAndCodeHash(hexutil.MustDecode(erc20_wrapper.Erc20WrapperMetaData.Bin))
			if err := app.replaceCodeAndCodeHash(ctx, wrapperAddr.Bytes(), wrapperCode, wrapperCodeHash); err != nil {
				return nil, err
			}

			// 2. update erc20 factory contract
			factoryAddr, err := app.EVMKeeper.GetERC20FactoryAddr(ctx)
			if err != nil {
				return nil, err
			}
			factoryCode, factoryCodeHash := extractRuntimeCodeAndCodeHash(hexutil.MustDecode(erc20_factory.Erc20FactoryMetaData.Bin))
			if err := app.replaceCodeAndCodeHash(ctx, factoryAddr.Bytes(), factoryCode, factoryCodeHash); err != nil {
				return nil, err
			}

			// 3. run migrations
			return app.ModuleManager.RunMigrations(ctx, cfg, versionMap)
		},
	)
}

var (
	initCodeOP = common.Hex2Bytes("5ff3fe")
)

func (app *MinitiaApp) replaceCodeAndCodeHash(ctx context.Context, contractAddr []byte, code, codeHash []byte) error {
	// load account and replace code hash
	acc := app.AccountKeeper.GetAccount(ctx, contractAddr)
	if acc == nil {
		return fmt.Errorf("account not found for contract address %s", contractAddr)
	}
	contractAcc, ok := acc.(*evmtypes.ContractAccount)
	if !ok {
		return fmt.Errorf("account is not a contract account for contract address %s", contractAddr)
	}

	// if the code hash is the same, do nothing
	if bytes.Equal(contractAcc.CodeHash, codeHash) {
		return nil
	}

	// update the code hash
	contractAcc.CodeHash = codeHash
	app.AccountKeeper.SetAccount(ctx, contractAcc)

	// replace runtime code
	codeKey := append(contractAddr, append(state.CodeKeyPrefix, codeHash...)...)
	err := app.EVMKeeper.VMStore.Set(ctx, codeKey, code)
	if err != nil {
		return err
	}

	// replace code size
	codeSizeKey := append(contractAddr, append(state.CodeSizeKeyPrefix, codeHash...)...)
	err = app.EVMKeeper.VMStore.Set(ctx, codeSizeKey, uint64ToBytes(uint64(len(code))))
	if err != nil {
		return err
	}

	return nil
}

func extractRuntimeCodeAndCodeHash(code []byte) ([]byte, []byte) {
	initCodePos := bytes.Index(code, initCodeOP)
	code = code[initCodePos+3:]
	codeHash := crypto.Keccak256Hash(code).Bytes()
	return code, codeHash
}

func uint64ToBytes(v uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, v)
	return bz
}
