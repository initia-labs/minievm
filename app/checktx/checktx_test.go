package checktx_test

import (
	"crypto/ecdsa"
	"math/big"
	"math/rand/v2"
	"testing"

	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ethereum/go-ethereum/common"

	"github.com/initia-labs/minievm/tests"
	evmkeeper "github.com/initia-labs/minievm/x/evm/keeper"
)

func Test_CheckTxWrapper(t *testing.T) {
	input := setupBackend(t)
	app, _, backend, addrs, privKeys := input.app, input.addrs, input.backend, input.addrs, input.privKeys

	defer app.Close()

	action := func(addr common.Address, privKey *ecdsa.PrivateKey) {
		nonce := uint64(0)
		txs := make([]sdk.Tx, 30)

		ctx, err := app.CreateQueryContext(0, false)
		require.NoError(t, err)

		for i := range 10 {
			txs[i*3+0], _ = tests.GenerateCreateERC20Tx(t, app, privKey, tests.SetNonce(nonce))
			txs[i*3+1], _ = tests.GenerateMintERC20Tx(t, app, privKey, addr, addr, big.NewInt(1000), tests.SetNonce(nonce+1))
			txs[i*3+2], _ = tests.GenerateTransferERC20Tx(t, app, privKey, addr, addr, big.NewInt(100), tests.SetNonce(nonce+2))
			nonce += 3
		}

		rand.Shuffle(len(txs), func(i, j int) {
			txs[i], txs[j] = txs[j], txs[i]
		})
		for _, tx := range txs {
			evmTx, _, err := evmkeeper.NewTxUtils(app.EVMKeeper).ConvertCosmosTxToEthereumTx(ctx, tx)
			require.NoError(t, err)
			require.NotNil(t, evmTx)

			txBz, err := evmTx.MarshalBinary()
			require.NoError(t, err)

			_, err = backend.SendRawTransaction(txBz)
			require.NoError(t, err)
		}
	}

	action(addrs[1], privKeys[1])

	require.Equal(t, 30, app.EVMIndexer().NumPendingTxs())
	require.Equal(t, 0, app.EVMIndexer().NumQueuedTxs())
}
