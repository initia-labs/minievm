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

const upgradeName = "0.6.6"

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
			params.HookMaxGas = opchildtypes.DefaultHookMaxGas

			err = app.OPChildKeeper.SetParams(ctx, params)
			if err != nil {
				return nil, err
			}

			//////////////////////////// MINIEVM ///////////////////////////////////

			// deploy and store erc20 factory contract address
			if err := app.EVMKeeper.DeployERC20Factory(ctx); err != nil &&
				// ignore contract address collision error (contract already deployed)
				!strings.Contains(err.Error(), vm.ErrContractAddressCollision.Error()) {
				return nil, err
			}

			// deploy and store erc20 wrapper contract address
			if err := app.EVMKeeper.DeployERC20Wrapper(ctx); err != nil &&
				// ignore contract address collision error (contract already deployed)
				!strings.Contains(err.Error(), vm.ErrContractAddressCollision.Error()) {
				return nil, err
			}

			code := hexutil.MustDecode(erc20.Erc20MetaData.Bin)

			// runtime code
			initCodeOP := common.Hex2Bytes("5ff3fe")
			initCodePos := bytes.Index(code, initCodeOP)
			code = code[initCodePos+3:]

			// code hash
			codeHash := crypto.Keccak256Hash(code).Bytes()

			// iterate all erc20 contracts and replace contract code to new version
			err = app.EVMKeeper.ERC20s.Walk(ctx, nil, func(contractAddr []byte) (bool, error) {
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
			if err != nil {
				return nil, err
			}

			return versionMap, nil
		},
	)
}

func uint64ToBytes(v uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, v)
	return bz
}
