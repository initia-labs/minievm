package indexer

import (
	"testing"

	abci "github.com/cometbft/cometbft/abci/types"
	coretypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/initia-labs/minievm/x/evm/types"
	"github.com/stretchr/testify/require"
)

func BenchmarkExtractLogsAndContractAddr(b *testing.B) {

	b.Run("smallEvent", func(b *testing.B) {
		smallEvent := []abci.Event{}

		for range 10 {
			smallEvent = append(smallEvent, abci.Event{
				Type: types.EventTypeEVM,
				Attributes: []abci.EventAttribute{
					{Key: types.AttributeKeyLog, Value: `{"address":"0x0000000000000000000000000000000000000000","topics":["0x0000000000000000000000000000000000000000000000000000000000000000"],"data":"0x0000000000000000000000000000000000000000000000000000000000000000"}`},
				},
			})
		}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_, _, err := extractLogsAndContractAddr(coretypes.ReceiptStatusSuccessful, smallEvent, false)
			require.NoError(b, err)
		}
	})

	b.Run("largeEvent", func(b *testing.B) {
		largeEvent := []abci.Event{}

		for range 10 {
			attrs := []abci.EventAttribute{}
			for range 100 {
				attrs = append(attrs, abci.EventAttribute{
					Key:   types.AttributeKeyLog,
					Value: `{"address":"0x0000000000000000000000000000000000000000","topics":["0x0000000000000000000000000000000000000000000000000000000000000000","0x0000000000000000000000000000000000000000000000000000000000000001","0x0000000000000000000000000000000000000000000000000000000000000002","0x0000000000000000000000000000000000000000000000000000000000000003"],"data":"0x0000000000000000000000000000000000000000000000000000000000000000"}`,
				})
			}
			largeEvent = append(largeEvent, abci.Event{
				Type:       types.EventTypeEVM,
				Attributes: attrs,
			})
		}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_, _, err := extractLogsAndContractAddr(coretypes.ReceiptStatusSuccessful, largeEvent, false)
			require.NoError(b, err)
		}
	})
}
