package backend_test

import (
	"encoding/json"
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/eth/tracers/logger"
	"github.com/stretchr/testify/require"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/eth/tracers"

	"github.com/initia-labs/minievm/tests"
	"github.com/initia-labs/minievm/x/evm/contracts/counter"
)

type traceResult struct {
	Gas         uint64          `json:"gas"`
	Failed      bool            `json:"failed"`
	ReturnValue string          `json:"returnValue"`
	Error       string          `json:"error,omitempty"`
	StructLogs  json.RawMessage `json:"structLogs"`
}

func Test_TraceTransaction_OutOfGas(t *testing.T) {
	input := setupBackend(t)
	app, backend, addrs, privKeys := input.app, input.backend, input.addrs, input.privKeys

	// deploy ERC20 contract
	codeBz, err := hexutil.Decode(counter.CounterBin)
	require.NoError(t, err)

	ctx := app.BaseApp.NewContext(true)
	_, contractAddr, _, err := app.EVMKeeper.EVMCreate(ctx, addrs[0], codeBz, nil, nil)
	require.NoError(t, err)

	// out-of-gas transfer
	abi, err := counter.CounterMetaData.GetAbi()
	require.NoError(t, err)

	inputBz, err := abi.Pack("loop")
	require.NoError(t, err)
	tx, txHash := tests.GenerateTx(t, app, privKeys[0], &contractAddr, inputBz, new(big.Int).SetUint64(0), tests.SetGasLimit(200000))
	_, finalizeRes := tests.ExecuteTxs(t, app, tx)
	tests.CheckTxResult(t, finalizeRes.TxResults[0], true)

	cfg := &tracers.TraceConfig{
		Config: &logger.Config{
			EnableMemory:     false,
			EnableReturnData: true,
			DisableStack:     true,
			DisableStorage:   true,
		},
	}

	res, err := backend.TraceTransaction(txHash, cfg)
	var result traceResult
	err = json.Unmarshal(res.(json.RawMessage), &result)
	require.NoError(t, err,  string(res.(json.RawMessage)))

	prettyJSON, err := json.MarshalIndent(result, "", "  ") 
    require.NoError(t, err, "failed to create pretty JSON from result")

    fmt.Println(string(prettyJSON))
	// require.True(t, result.Failed, "expected transaction to be marked as failed")
	// require.Equal(t, "out of gas", result.Error, "expected error message to be 'out of gas'")

}
