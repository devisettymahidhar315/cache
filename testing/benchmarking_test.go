package testing

import (
	"testing"

	"github.com/devisettymahidhar315/cache/in_memory"
	"github.com/devisettymahidhar315/cache/multi_cache"
	"github.com/devisettymahidhar315/cache/redis"
)

const len = 2

var cache = multi_cache.NewMultiCache()
var redis_cache = redis.NewLRUCache()
var inmemory_cache = in_memory.NewLRUCache(1)

// benchmarking

//multi-cache

//becnchmarking for set method of the multi-cache

func BenchmarkSet(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		cache.Set("a", "1", len, -1)
	}
}

//benchmarking for get methods of the cache

func BenchmarkGet(b *testing.B) {
	cache.Set("a", "1", len, -1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cache.Get("a")
	}
}

//benchmarking for del methods of the cache

func BenchmarkDel(b *testing.B) {
	cache.Set("a", "1", len, -1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cache.Del("a")
	}
}

// benchmarking for printing redis methods of the cache

func BenchmarkPrint(b *testing.B) {
	cache.Set("a", "1", len, -1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cache.Print()
	}
}

//benchmarking for del_all of the cache

func BenchmarkDel_All(b *testing.B) {
	cache.Set("a", "1", len, -1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cache.Del_ALL()
	}
}

// redis
func BenchmarkRedisSet(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		redis_cache.Put("a", "1", len, -1)
	}
}

//benchmarking for get methods of the cache

func BenchmarkRedisGet(b *testing.B) {
	redis_cache.Put("a", "1", len, -1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		redis_cache.Get("a")
	}
}

//benchmarking for del methods of the cache

func BenchmarkRedisDel(b *testing.B) {
	redis_cache.Put("a", "1", len, -1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		redis_cache.Del("a")
	}
}

//benchmarking for printing redis methods of the cache

func BenchmarkRedisPrint(b *testing.B) {
	redis_cache.Put("a", "1", len, -1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		redis_cache.Print()
	}
}

//benchmarking for del_all of the cache

func BenchmarkRedisDel_All(b *testing.B) {
	redis_cache.Put("a", "1", len, -1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		redis_cache.DEL_ALL()
	}
}

// inmemory
func BenchmarkInmemorySet(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		inmemory_cache.Put("a", "1", len, -1)
	}
}

//benchmarking for get methods of the cache

func BenchmarkInmemoryGet(b *testing.B) {
	inmemory_cache.Put("a", "1", len, -1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		inmemory_cache.Get("a")
	}
}

//benchmarking for del methods of the cache

func BenchmarkInmemoryDel(b *testing.B) {
	inmemory_cache.Put("a", "1", len, -1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		inmemory_cache.Del("a")
	}
}

//benchmarking for printing redis methods of the cache

func BenchmarkInmemoryPrint(b *testing.B) {
	inmemory_cache.Put("a", "1", len, -1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		inmemory_cache.Print()
	}
}

//benchmarking for del_all of the cache

func BenchmarkInmemoryDel_All(b *testing.B) {
	inmemory_cache.Put("a", "1", len, -1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		inmemory_cache.DEL_ALL()
	}
}
