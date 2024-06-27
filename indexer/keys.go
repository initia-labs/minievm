package indexer

import (
	"cosmossdk.io/collections"
)

var (
	// block indexes
	prefixBlockHeader       = collections.Prefix([]byte{0, 0, 1, 1})
	prefixBlockHashToNumber = collections.Prefix([]byte{0, 0, 1, 2})

	// tx indexes
	prefixTx                    = collections.Prefix([]byte{0, 0, 2, 1})
	prefixTxReceipt             = collections.Prefix([]byte{0, 0, 2, 2})
	prefixBlockAndIndexToTxHash = collections.Prefix([]byte{0, 0, 2, 3})
)
