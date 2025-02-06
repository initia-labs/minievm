package app

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/crypto"

	upgradetypes "cosmossdk.io/x/upgrade/types"
	"github.com/cosmos/cosmos-sdk/types/module"

	"github.com/initia-labs/minievm/x/evm/contracts/erc20"
	"github.com/initia-labs/minievm/x/evm/state"
	evmtypes "github.com/initia-labs/minievm/x/evm/types"

	opchildtypes "github.com/initia-labs/OPinit/x/opchild/types"
)

const upgradeName = "0.7.1"

// RegisterUpgradeHandlers returns upgrade handlers
func (app *MinitiaApp) RegisterUpgradeHandlers(cfg module.Configurator) {
	app.UpgradeKeeper.SetUpgradeHandler(
		upgradeName,
		func(ctx context.Context, _ upgradetypes.Plan, versionMap module.VersionMap) (module.VersionMap, error) {
			//////////////////////////// OPINIT ////////////////////////////////////

			// opchild params update
			params, err := app.OPChildKeeper.GetParams(ctx)
			if err != nil {
				return nil, err
			}

			// set non-zero default values for new params
			if params.HookMaxGas == 0 {
				params.HookMaxGas = opchildtypes.DefaultHookMaxGas

				err = app.OPChildKeeper.SetParams(ctx, params)
				if err != nil {
					return nil, err
				}
			}

			//////////////////////////// MINIEVM ///////////////////////////////////

			// try to deploy and store erc20 wrapper contract address
			if err := app.EVMKeeper.DeployERC20Wrapper(ctx); err != nil &&
				// ignore contract address collision error (contract already deployed)
				!strings.Contains(err.Error(), vm.ErrContractAddressCollision.Error()) {
				return nil, err
			}

			// try to deploy and store erc20 factory contract address
			if err := app.EVMKeeper.DeployERC20Factory(ctx); err != nil &&
				// ignore contract address collision error (contract already deployed)
				!strings.Contains(err.Error(), vm.ErrContractAddressCollision.Error()) {
				return nil, err
			} else if err == nil {
				// update erc20 contracts only if erc20 factory contract has been deployed without error.
				//
				// address collision error is ignored because it means that the contract has already been deployed
				// and the erc20 contracts have already been updated.
				//
				err = app.updateERC20s(ctx)
				if err != nil {
					return nil, err
				}
			}

			// try to deploy and store connect oracle contract address
			if err := app.EVMKeeper.DeployConnectOracle(ctx); err != nil &&
				// ignore contract address collision error (contract already deployed)
				!strings.Contains(err.Error(), vm.ErrContractAddressCollision.Error()) {
				return nil, err
			}

			return versionMap, nil
		},
	)
}

// updateERC20s updates all erc20 contracts to the new version
// - update contract code
// - update contract code hash
// - update contract code size
func (app *MinitiaApp) updateERC20s(ctx context.Context) error {
	code := hexutil.MustDecode(erc20.Erc20MetaData.Bin)

	// runtime code
	initCodeOP := common.Hex2Bytes("5ff3fe")
	initCodePos := bytes.Index(code, initCodeOP)
	code = code[initCodePos+3:]

	// code hash
	codeHash := crypto.Keccak256Hash(code).Bytes()

	// iterate all erc20 contracts and replace contract code to new version
	return app.EVMKeeper.ERC20s.Walk(ctx, nil, func(contractAddr []byte) (bool, error) {
		acc := app.AccountKeeper.GetAccount(ctx, contractAddr)
		if acc == nil {
			return true, fmt.Errorf("account not found for contract address %s", contractAddr)
		}

		contractAcc, ok := acc.(*evmtypes.ContractAccount)
		if !ok {
			return true, fmt.Errorf("account is not a contract account for contract address %s", contractAddr)
		}

		contractAcc.CodeHash = codeHash
		app.AccountKeeper.SetAccount(ctx, contractAcc)

		// set code
		codeKey := append(contractAddr, append(state.CodeKeyPrefix, codeHash...)...)
		err := app.EVMKeeper.VMStore.Set(ctx, codeKey, code)
		if err != nil {
			return true, err
		}

		// set code size
		codeSizeKey := append(contractAddr, append(state.CodeSizeKeyPrefix, codeHash...)...)
		err = app.EVMKeeper.VMStore.Set(ctx, codeSizeKey, uint64ToBytes(uint64(len(code))))
		if err != nil {
			return true, err
		}

		return false, nil
	})
}

func uint64ToBytes(v uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, v)
	return bz
}
