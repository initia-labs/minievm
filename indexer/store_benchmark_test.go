package indexer

import (
	dbm "github.com/cosmos/cosmos-db"
	"testing"
)

func BenchmarkCacheStore(b *testing.B) {
	// Create a memory DB for testing
	db := dbm.NewMemDB()
	store := NewCacheStoreWithBatch(db, 100*1024*1024) // 100MB cache
	defer store.Write()                                // Cleanup

	// Benchmark Set operations with batch
	b.Run("SetWithBatch", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			// Do 10 operations per batch
			for j := 0; j < 10; j++ {
				key := []byte{byte(i % 256), byte((i / 256) % 256), byte(j)}
				value := []byte{byte(i % 256), byte((i / 256) % 256), byte((i / 65536) % 256), byte(j)}
				_ = store.Set(key, value)
			}
			store.Write()
		}
	})

	// Benchmark Get operations with cache hits
	b.Run("GetWithCacheHits", func(b *testing.B) {
		// Pre-populate with a small dataset
		for i := 0; i < 100; i++ {
			key := []byte{byte(i % 256), byte((i / 256) % 256)}
			value := []byte{byte(i % 256), byte((i / 256) % 256), byte((i / 65536) % 256)}
			_ = store.Set(key, value)
		}
		store.Write()

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			key := []byte{byte(i % 100), byte((i / 100) % 256)}
			_, _ = store.Get(key)
		}
	})

	// Benchmark Get operations with cache misses
	b.Run("GetWithCacheMisses", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			key := []byte{byte(i % 256), byte((i / 256) % 256), byte((i / 65536) % 256)}
			_, _ = store.Get(key)
		}
	})

	// Benchmark concurrent operations
	b.Run("Concurrent", func(b *testing.B) {
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			i := 0
			for pb.Next() {
				key := []byte{byte(i % 256), byte((i / 256) % 256)}
				value := []byte{byte(i % 256), byte((i / 256) % 256), byte((i / 65536) % 256)}
				_ = store.Set(key, value)
				i++
			}
		})
		store.Write()
	})

	// Benchmark large dataset operations
	b.Run("LargeDataset", func(b *testing.B) {
		// Pre-populate with a large dataset
		for i := 0; i < 1000; i++ {
			key := []byte{byte(i % 256), byte((i / 256) % 256)}
			value := []byte{byte(i % 256), byte((i / 256) % 256), byte((i / 65536) % 256)}
			_ = store.Set(key, value)
		}
		store.Write()

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			// Mix of operations
			key := []byte{byte(i % 1000), byte((i / 1000) % 256)}
			_ = store.Set(key, []byte("new value"))
			_, _ = store.Get(key)
			iter, _ := store.Iterator(nil, nil)
			for ; iter.Valid(); iter.Next() {
				_ = iter.Key()
				_ = iter.Value()
			}
			iter.Close()
		}
	})
}
