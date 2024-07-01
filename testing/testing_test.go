package testing

import (
	"testing"
	"time"

	"github.com/devisettymahidhar315/cache/in_memory"
	"github.com/devisettymahidhar315/cache/multi_cache"
	"github.com/devisettymahidhar315/cache/redis"
)

const len1 = 2

// Initialize a new inmemory cache
var cache_inmemory = in_memory.NewLRUCache(1 * time.Second)
var cache_redis = redis.NewLRUCache()

// Length of the cache for testing
// in memory
// TestPut_inmemory tests the Put method of the inmemory cache
func TestPut_inmemory(t *testing.T) {
	cache_inmemory.DEL_ALL()

	// Insert key-value pairs into the cache
	cache_inmemory.Put("a1", "1", len1, -1)
	cache_inmemory.Put("b1", "2", len1, -1)

	// Retrieve the value for key "a1"
	result := cache_inmemory.Get("a1")

	// Check if the value is as expected
	if result != "1" {
		t.Error("Expected value '1', got", result)
	}
	cache_inmemory.DEL_ALL()
	result = cache_inmemory.Get("a1")
	if result != "" {
		t.Error("Expected value '', got", result)
	}
	// checking the time to live
	cache_inmemory.Put("a", "1", len1, 5)
	time.Sleep(6 * time.Second)
	r := cache_inmemory.Get("a")
	if r != "" {
		t.Error("Expected value ''. got ", r)
	}
}

// TestGet_inmemory tests the Get method of the inmemory cache

func TestGet_inmemory(t *testing.T) {
	cache_inmemory.DEL_ALL()

	// Insert key-value pairs into the cache
	cache_inmemory.Put("a1", "1", len1, -1)
	cache_inmemory.Put("b1", "2", len1, -1)

	// Retrieve the value for key "a1"
	result := cache_inmemory.Get("a1")
	if result != "1" {
		t.Error("Expected value '1', got", result)
	}

	// Attempt to retrieve a value for a non-existent key "v"
	result = cache_inmemory.Get("v")
	if result != "" {
		t.Error("Expected empty string for non-existent key, got", result)
	}
}

// TestPrint_inmemory tests the Print method of the inmemory cache
func TestPrint_inmemory(t *testing.T) {
	cache_inmemory.DEL_ALL()

	// Insert key-value pairs into the cache
	cache_inmemory.Put("a1", "1", len1, -1)
	cache_inmemory.Put("b1", "2", len1, -1)

	// Print the current state of the cache
	result := cache_inmemory.Print()
	expected_result := "b1:2, a1:1"

	// Check if the printed result matches the expected result
	if result != expected_result {
		t.Error("Expected", expected_result, "got", result)
	}

	// Insert another key-value pair to exceed the cache length
	cache_inmemory.Put("c1", "3", len1, -1)

	// Print the current state of the cache
	result = cache_inmemory.Print()
	expected_result = "c1:3, b1:2"

	// Check if the printed result matches the expected result
	if result != expected_result {
		t.Error("Expected", expected_result, "got", result)
	}
}

// TestDel_inmemory tests the Del method of the inmemory cache
func TestDel_inmemory(t *testing.T) {
	cache_inmemory.DEL_ALL()
	// Insert a key-value pair into the cache
	cache_inmemory.Put("a1", "1", len1, -1)

	// Delete the key "a1"
	cache_inmemory.Del("a1")

	// Print the current state of the cache
	result := cache_inmemory.Print()
	expected_result := ""

	// Check if the printed result matches the expected result
	if result != expected_result {
		t.Error("Expected", expected_result, "got", result)
	}

	// Insert key-value pairs into the cache
	cache_inmemory.Put("a1", "1", len1, -1)
	cache_inmemory.Put("b1", "2", len1, -1)

	// Delete the key "a1"
	cache_inmemory.Del("a1")

	// Attempt to retrieve the value for the deleted key "a1"
	result = cache_inmemory.Get("a1")
	expected_result = ""

	// Check if the result matches the expected result
	if result != expected_result {
		t.Error("Expected empty string for deleted key, got", result)
	}
}

//TestDelAll_inmemory tests the Del_All method of the inmemory cache

func TestDelAll_inmemory(t *testing.T) {
	cache_inmemory.DEL_ALL()

	// Insert a key-value pair into the cache

	cache_inmemory.Put("a", "1", len1, -1)
	cache_inmemory.Put("b", "2", len1, -1)

	// Delete all key
	cache_inmemory.DEL_ALL()

	// Attempt to retrieve the value for the deleted key "a1"
	result := cache_inmemory.Get("a1")
	expected_result := ""

	// Check if the result matches the expected result
	if result != expected_result {
		t.Error("Expected empty string for deleted key, got", result)
	}

}

// redis
// TestPut_redis tests the Put method of the Redis cache
func TestPut_redis(t *testing.T) {
	cache_redis.DEL_ALL()

	// Insert key-value pairs into the cache
	cache_redis.Put("a1", "1", len1, -1)
	cache_redis.Put("b1", "2", len1, -1)

	// Retrieve the value for key "a1"
	result := cache_redis.Get("a1")

	// Check if the value is as expected
	if result != "1" {
		t.Error("Expected value '1', got", result)
	}
}

// TestGet_redis tests the Get method of the Redis cache
func TestGet_redis(t *testing.T) {
	cache_redis.DEL_ALL()

	// Insert key-value pairs into the cache
	cache_redis.Put("a1", "1", len1, -1)
	cache_redis.Put("b1", "2", len1, -1)

	// Retrieve the value for key "a1"
	result := cache_redis.Get("a1")
	if result != "1" {
		t.Error("Expected value '1', got", result)
	}

	// Attempt to retrieve a value for a non-existent key "v"
	result = cache_redis.Get("v")
	if result != "" {
		t.Error("Expected empty string for non-existent key, got", result)
	}
}

// TestPrint_redis tests the Print method of the Redis cache
func TestPrint_redis(t *testing.T) {
	cache_redis.DEL_ALL()

	// Insert key-value pairs into the cache
	cache_redis.Put("a1", "1", len1, -1)
	cache_redis.Put("b1", "2", len1, -1)

	// Print the current state of the cache
	result := cache_redis.Print()

	expected_result := "b1:2, a1:1"

	// Check if the printed result matches the expected result
	if result != expected_result {
		t.Error("Expected", expected_result, "got", result)
	}

	// Insert another key-value pair to exceed the cache length
	cache_redis.Put("c1", "3", len1, -1)

	// Print the current state of the cache
	result = cache_redis.Print()
	expected_result = "c1:3, b1:2"

	// Check if the printed result matches the expected result
	if result != expected_result {
		t.Error("Expected", expected_result, "got", result)
	}
}

// TestDel_redis tests the Del method of the Redis cache
func TestDel_redis(t *testing.T) {
	cache_redis.DEL_ALL()

	// Insert a key-value pair into the cache
	cache_redis.Put("a1", "1", len1, -1)

	// Delete the key "a1"
	cache_redis.Del("a1")

	// Print the current state of the cache
	result := cache_redis.Print()
	expected_result := ""

	// Check if the printed result matches the expected result
	if result != expected_result {
		t.Error("Expected", expected_result, "got", result)
	}

}

// TestDelAll_redis tests the Del_All method of the Redis cache
func TestDelAll_redis(t *testing.T) {
	// Insert a key-value pair into the cache

	cache_redis.Put("a", "1", len1, -1)
	cache_redis.Put("b", "2", len1, -1)

	// Delete all key
	cache_redis.DEL_ALL()

	// Attempt to retrieve the value for the deleted key "a1"
	result := cache_redis.Get("a1")
	expected_result := ""

	// Check if the result matches the expected result
	if result != expected_result {
		t.Error("Expected empty string for deleted key, got", result)
	}

}

// integration testing
// testing related to multi_cache folder
func TestGET(t *testing.T) {
	// Create a new multi-cache instance
	cache := multi_cache.NewMultiCache()
	// Set key-value pairs in the cache
	cache.Set("a", "1", len1, -1)
	cache.Set("b", "2", len1, -1)

	// Test case 1: Get a non-existent key
	res1 := cache.Get("c")
	if res1 != "key is not present" {
		t.Error("expected '' got ", res1)
	}

	// Test case 2: Get an existing key
	res2 := cache.Get("a")
	if res2 != "1" {
		t.Error("expected '1' got ", res2)
	}
	cache.Del_ALL()

	//storing the empty string
	s := cache.Set("k", "", len1, -1)
	if s != "empty string" {
		t.Error("expected 'empty string' got", s)
	}
	//try to retrive the key
	s = cache.Get("k")
	if s != "key is not present" {
		t.Error("expected 'key is not present' got ", s)
	}

	// printing
	s = cache.Print()
	if s != "no key value pair is present" {
		t.Error("expected 'no key value pair is present' got ", s)
	}
	// setting the key
	s = cache.Set("k", "1", len1, -1)
	if s != "successfully added" {
		t.Error("expected 'successfully added' got ", s)
	}
	cache.Del_ALL()

	//storing the data only in in memory
	cache_redis.Put("k", "1", len1, -1)
	cache.Set("a", "1", len1, -1)
	s = cache.Print()
	if s != "error: 'data is not same'" {
		t.Error("expected 'error: 'data is not same' got ", s)
	}
	cache.Del_ALL()

	//storing the data only in in memory
	cache_redis.Put("k", "1", len1, -1)

	s = cache.Get("k")
	if s != "value is not same in inmemory and redis" {
		t.Error("expected 'value is not same in inmemory and redis' got ", s)
	}
	cache.Del_ALL()

	//time to live
	cache.Del_ALL()
	cache.Set("a", "1", len1, 5)
	time.Sleep(6 * time.Second)
	res3 := cache.Get("a")
	if res3 != "key is not present" {
		t.Error("expected '' got ", res3)
	}
}

// TestDel tests the Del method of the cache
func TestDel(t *testing.T) {
	// Create a new multi-cache instance
	cache := multi_cache.NewMultiCache()
	// Set key-value pairs in the cache
	cache.Set("a", "1", len1, -1)
	cache.Set("b", "2", len1, -1)

	// Delete a key from the cache
	cache.Del("a")
	// Check if the deleted key returns an empty string
	result := cache.Get("a")
	if result != "key is not present" {
		t.Error("expected empty string for deleted key 'a'")
	}
	// Check if an existing key returns the correct value
	result = cache.Get("b")
	if result != "2" {
		t.Error("expected '2' for key 'b'")
	}
}
