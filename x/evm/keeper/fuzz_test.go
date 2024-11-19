package keeper_test

import (
	"bytes"
	"sync"
	"sync/atomic"
	"testing"

	storetypes "cosmossdk.io/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	coretypes "github.com/ethereum/go-ethereum/core/types"

	"github.com/holiman/uint256"
	"github.com/stretchr/testify/require"

	"github.com/initia-labs/minievm/x/evm/contracts/counter"
	"github.com/initia-labs/minievm/x/evm/types"
)

func Fuzz_Concurrent_Counter(f *testing.F) {
	f.Add(uint8(100), uint8(100))
	f.Fuzz(func(t *testing.T, numThread uint8, numCount uint8) {
		if numThread == 0 || numCount == 0 {
			t.Skip("skip invalid input")
		}

		ctx, input := createDefaultTestInput(t)
		_, _, addr := keyPubAddr()

		counterBz, err := hexutil.Decode(counter.CounterBin)
		require.NoError(t, err)

		caller := common.BytesToAddress(addr.Bytes())
		retBz, contractAddr, _, err := input.EVMKeeper.EVMCreate(ctx, caller, counterBz, nil, nil)
		require.NoError(t, err)
		require.NotEmpty(t, retBz)
		require.Len(t, contractAddr, 20)

		parsed, err := counter.CounterMetaData.GetAbi()
		require.NoError(t, err)

		count := getCount(t, ctx, input, contractAddr)
		require.Equal(t, uint256.NewInt(0), count)

		inputBz, err := parsed.Pack("increase_for_fuzz", uint64(numCount))
		require.NoError(t, err)

		atomicBloomBytes := atomic.Pointer[[]byte]{}
		atomicBloomBytes.Store(nil)

		var wg sync.WaitGroup
		cacheCtxes := make([]sdk.Context, numThread)
		for i := uint8(0); i < numThread; i++ {
			wg.Add(1)
			cacheCtx, _ := ctx.CacheContext()
			cacheCtx = cacheCtx.WithGasMeter(storetypes.NewInfiniteGasMeter())
			cacheCtxes[i] = cacheCtx
			go func(ctx sdk.Context) {
				defer wg.Done()

				// call with value
				res, logs, err := input.EVMKeeper.EVMCall(ctx, caller, contractAddr, inputBz, nil, nil)
				require.NoError(t, err)
				require.Empty(t, res)
				bloomBytes := coretypes.LogsBloom(logs.ToEthLogs())
				prev := atomicBloomBytes.Swap(&bloomBytes)
				require.True(t, prev == nil || bytes.Equal(*prev, bloomBytes))
			}(cacheCtx)
		}
		wg.Wait()

		for i := uint8(0); i < numThread; i++ {
			count := getCount(t, cacheCtxes[i], input, contractAddr)
			require.Equal(t, uint256.NewInt(uint64(numCount)), count)
			require.NotEmpty(t, atomicBloomBytes.Load())
		}
	})
}

func getCount(t *testing.T, ctx sdk.Context, input TestKeepers, contractAddr common.Address) *uint256.Int {
	parsed, err := counter.CounterMetaData.GetAbi()
	require.NoError(t, err)

	queryInputBz, err := parsed.Pack("count")
	require.NoError(t, err)

	queryRes, err := input.EVMKeeper.EVMStaticCall(ctx, types.StdAddress, contractAddr, queryInputBz, nil)
	require.NoError(t, err)

	return uint256.NewInt(0).SetBytes32(queryRes)
}
