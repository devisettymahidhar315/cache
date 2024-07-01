package multi_cache

import (
	"strings"
	"sync"
	"time"

	"github.com/devisettymahidhar315/cache/in_memory"
	"github.com/devisettymahidhar315/cache/redis"
)

// MultiCache struct manages both Redis and in-memory caches.
type MultiCache struct {
	redisCache    *redis.LRUCache
	inMemoryCache *in_memory.LRUCache
}

// NewMultiCache initializes a new MultiCache with Redis and in-memory LRU caches.
func NewMultiCache() *MultiCache {
	return &MultiCache{
		redisCache:    redis.NewLRUCache(),
		inMemoryCache: in_memory.NewLRUCache(1 * time.Second),
	}
}

// multi cache
// Set stores the key-value pair in both Redis and in-memory caches concurrently.
func (c *MultiCache) Set(key, value string, length int, t int) string {
	if strings.TrimSpace(value) == "" {
		return "empty string"
	}

	var wg sync.WaitGroup
	value = strings.TrimSpace(value)
	wg.Add(2) // Add two goroutines to the wait group
	// Store in Redis cache concurrently
	go func() {
		defer wg.Done()

		c.redisCache.Put(key, value, length, t)
	}()
	// Store in inmemory cache concurrently
	go func() {
		defer wg.Done()
		c.inMemoryCache.Put(key, value, length, t)
	}()
	wg.Wait() // Wait for both goroutines to finish
	return "successfully added"
}

// Get retrieves the value for a key from both Redis and in-memory caches concurrently and compares them.
func (c *MultiCache) Get(key string) string {

	var redis_value, inmemory_value string
	var wg sync.WaitGroup
	wg.Add(2) // Add two goroutines to the wait group
	// Retrieve value from in-memory cache concurrently
	go func() {
		defer wg.Done()
		inmemory_value = c.inMemoryCache.Get(key)
	}()
	// Retrieve value from redis cache concurrently
	go func() {
		defer wg.Done()
		redis_value = c.redisCache.Get(key)
	}()
	wg.Wait() // Wait for both goroutines to finish

	// Return the value if they match, otherwise return an empty string
	if redis_value == inmemory_value {
		if strings.TrimSpace(redis_value) == "" {
			return "key is not present"
		}
		return redis_value
	} else {
		return "value is not same in inmemory and redis"
	}

}

func (c *MultiCache) Print() string {

	var wg sync.WaitGroup
	wg.Add(2) // Add one goroutine to the wait group
	var result1, result2 string

	// Print the in-memory cache contents concurrently
	go func() {
		defer wg.Done()
		result1 = c.inMemoryCache.Print()
	}()
	// Print the Redis cache contents concurrently
	go func() {
		defer wg.Done()
		result2 = c.redisCache.Print()
	}()

	wg.Wait() // Wait for the goroutine to finish
	if result1 == result2 {
		if strings.TrimSpace(result1) == "" {
			return "no key value pair is present"
		}
		return result2
	} else {
		return "error: 'data is not same'"
	}

}

// Del deletes the key-value pair from both Redis and in-memory caches concurrently.
func (c *MultiCache) Del(key string) {

	var wg sync.WaitGroup
	wg.Add(2) // Add two goroutines to the wait group

	// Delete from in-memory cache concurrently
	go func() {
		defer wg.Done()
		c.inMemoryCache.Del(key)
	}()

	// Delete from Redis cache concurrently
	go func() {
		defer wg.Done()
		c.redisCache.Del(key)
	}()

	wg.Wait() // Wait for both goroutines to finish

}

// Del_ALL deletes the entire data from both Redis and in-memory caches concurrently.
func (c *MultiCache) Del_ALL() {
	var wg sync.WaitGroup

	wg.Add(2) // Add two goroutines to the wait group

	// Delete from in-memory cache concurrently
	go func() {
		defer wg.Done()
		c.inMemoryCache.DEL_ALL()
	}()

	// Delete from Redis cache concurrently
	go func() {
		defer wg.Done()
		c.redisCache.DEL_ALL()
	}()

	wg.Wait() // Wait for both goroutines to finish

}

// redis
func (c *MultiCache) Redis_Get(key string) string {
	var wg sync.WaitGroup
	wg.Add(1) // Add two goroutines to the wait group
	var st string
	// Delete from in-memory cache concurrently
	go func() {
		defer wg.Done()
		st = c.redisCache.Get(key)
	}()

	wg.Wait() // Wait for both goroutines to finish
	if strings.TrimSpace(st) == "" {
		return "no key is present"
	}
	return st
}

func (c *MultiCache) Redis_Print() string {

	var wg sync.WaitGroup
	wg.Add(1) // Add one goroutine to the wait group
	var result string

	// Print the Redis cache contents concurrently
	go func() {
		defer wg.Done()
		result = c.redisCache.Print()
	}()
	wg.Wait()
	if strings.TrimSpace(result) == "" {
		return "no key value pair is present"
	}
	return result
}

func (c *MultiCache) Redis_Del_ALL() {
	var wg sync.WaitGroup
	wg.Add(1) // Add one goroutines to the wait group

	// Delete from Redis cache concurrently
	go func() {
		defer wg.Done()
		c.redisCache.DEL_ALL()
	}()

	wg.Wait() // Wait for both goroutines to finish
}

func (c *MultiCache) Redis_Del(key string) {

	var wg sync.WaitGroup
	wg.Add(1) // Add one goroutines to the wait group

	// Delete from Redis cache concurrently
	go func() {
		defer wg.Done()
		c.redisCache.Del(key)
	}()

	wg.Wait() // Wait for both goroutines to finish
}

func (c *MultiCache) Redis_Set(key, value string, length int, t int) string {
	if strings.TrimSpace(value) == "" {
		return "empty string"
	}

	var wg sync.WaitGroup
	value = strings.TrimSpace(value)
	wg.Add(1) // Add two goroutines to the wait group
	// Store in Redis cache concurrently
	go func() {
		defer wg.Done()

		c.redisCache.Put(key, value, length, t)
	}()

	wg.Wait() // Wait for both goroutines to finish
	return "successfully added"
}

// inmemory
// Set stores the key-value pair in both Redis and in-memory caches concurrently.
func (c *MultiCache) Inmemory_Set(key, value string, length int, t int) string {
	if strings.TrimSpace(value) == "" {
		return "empty string"
	}

	var wg sync.WaitGroup
	value = strings.TrimSpace(value)
	wg.Add(1) // Add two goroutines to the wait group

	// Store in inmemory cache concurrently
	go func() {
		defer wg.Done()
		c.inMemoryCache.Put(key, value, length, t)
	}()
	wg.Wait() // Wait for both goroutines to finish
	return "successfully added"
}

// Get retrieves the value for a key from both Redis and in-memory caches concurrently and compares them.
func (c *MultiCache) Inmemory_Get(key string) string {

	var inmemory_value string
	var wg sync.WaitGroup
	wg.Add(1) // Add two goroutines to the wait group
	// Retrieve value from in-memory cache concurrently
	go func() {
		defer wg.Done()
		inmemory_value = c.inMemoryCache.Get(key)
	}()

	wg.Wait() // Wait for both goroutines to finish
	if strings.TrimSpace(inmemory_value) == "" {
		return "no key is present"
	}
	return inmemory_value

}

func (c *MultiCache) Inmemory_Print() string {

	var wg sync.WaitGroup
	wg.Add(1) // Add one goroutine to the wait group
	var result1 string

	// Print the in-memory cache contents concurrently
	go func() {
		defer wg.Done()
		result1 = c.inMemoryCache.Print()
	}()

	wg.Wait()
	if strings.TrimSpace(result1) == "" {
		return "no key value is present"
	}
	return result1

}

// Del deletes the key-value pair from both Redis and in-memory caches concurrently.
func (c *MultiCache) Inmemory_Del(key string) {

	var wg sync.WaitGroup
	wg.Add(1) // Add two goroutines to the wait group

	// Delete from in-memory cache concurrently
	go func() {
		defer wg.Done()
		c.inMemoryCache.Del(key)
	}()

	wg.Wait() // Wait for both goroutines to finish
}

// Del_ALL deletes the entire data from both Redis and in-memory caches concurrently.
func (c *MultiCache) Inmemory_Del_ALL() {
	var wg sync.WaitGroup
	wg.Add(1) // Add two goroutines to the wait group

	// Delete from in-memory cache concurrently
	go func() {
		defer wg.Done()
		c.inMemoryCache.DEL_ALL()
	}()
	wg.Wait() // Wait for both goroutines to finish
}
