package testing

import (
	"testing"

	"github.com/devisettymahidhar315/cache/multi_cache"
)

const len = 2

// benchmarking

//multi-cache
//becnchmarking for set method of the multi-cache

func BenchmarkSet(b *testing.B) {
	// Create a new multi-cache instance
	cache := multi_cache.NewMultiCache()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cache.Set("a", "1", len, -1)
	}
}

//benchmarking for get methods of the cache

func BenchmarkGet(b *testing.B) {
	// Create a new multi-cache instance
	cache := multi_cache.NewMultiCache()
	cache.Set("a", "1", len, -1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cache.Get("a")

	}
}

//benchmarking for del methods of the cache

func BenchmarkDel(b *testing.B) {
	// Create a new multi-cache instance
	cache := multi_cache.NewMultiCache()
	cache.Set("a", "1", len, -1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cache.Del("a")
	}
}

//benchmarking for printing redis methods of the cache

func BenchmarkPrintRedis(b *testing.B) {
	// Create a new multi-cache instance
	cache := multi_cache.NewMultiCache()
	cache.Set("a", "1", len, -1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cache.Print_redis()
	}
}

//benchmarking for printing in memory methods of the cache

func BenchmarkPrintinmemory(b *testing.B) {
	// Create a new multi-cache instance
	cache := multi_cache.NewMultiCache()
	cache.Set("a", "1", len, -1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cache.Print_in_mem()
	}
}

//benchmarking for del_all of the cache

func BenchmarkDel_All(b *testing.B) {
	// Create a new multi-cache instance
	cache := multi_cache.NewMultiCache()
	cache.Set("a", "1", len, -1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cache.Del_ALL()
	}
}
