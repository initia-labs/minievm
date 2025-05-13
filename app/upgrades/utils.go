package upgrades

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"

	evmstate "github.com/initia-labs/minievm/x/evm/state"
	evmtypes "github.com/initia-labs/minievm/x/evm/types"
)

// ReplaceCodeAndCodeHash replaces the code and code hash for a contract.
func ReplaceCodeAndCodeHash(ctx context.Context, app MinitiaApp, contractAddr []byte, code, codeHash []byte) error {
	// load account and replace code hash
	acc := app.GetAccountKeeper().GetAccount(ctx, contractAddr)
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
	app.GetAccountKeeper().SetAccount(ctx, contractAcc)

	// replace runtime code
	codeKey := append(contractAddr, append(evmstate.CodeKeyPrefix, codeHash...)...)
	err := app.GetEVMKeeper().VMStore.Set(ctx, codeKey, code)
	if err != nil {
		return err
	}

	// replace code size
	codeSizeKey := append(contractAddr, append(evmstate.CodeSizeKeyPrefix, codeHash...)...)
	err = app.GetEVMKeeper().VMStore.Set(ctx, codeSizeKey, uint64ToBytes(uint64(len(code))))
	if err != nil {
		return err
	}

	return nil
}

func CodeHash(code []byte) []byte {
	return crypto.Keccak256Hash(code).Bytes()
}

func uint64ToBytes(v uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, v)
	return bz
}
