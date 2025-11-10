package cache

import (
	"sync"
	"testing"
)

func TestNew(t *testing.T) {
	cache := New[string, int]()
	if cache == nil {
		t.Fatal("New() returned nil")
	}
	if cache.items == nil {
		t.Fatal("cache.items is nil")
	}
}

func TestSet(t *testing.T) {
	cache := New[string, int]()
	cache.Set("key1", 42)
	
	value, found := cache.Get("key1")
	if !found {
		t.Fatal("key1 not found after Set")
	}
	if value != 42 {
		t.Errorf("expected 42, got %d", value)
	}
}

func TestGet(t *testing.T) {
	cache := New[string, int]()
	
	// Test non-existent key
	_, found := cache.Get("nonexistent")
	if found {
		t.Error("found non-existent key")
	}
	
	// Test existing key
	cache.Set("key1", 100)
	value, found := cache.Get("key1")
	if !found {
		t.Error("key1 not found")
	}
	if value != 100 {
		t.Errorf("expected 100, got %d", value)
	}
}

func TestRemove(t *testing.T) {
	cache := New[string, int]()
	cache.Set("key1", 42)
	
	cache.Remove("key1")
	_, found := cache.Get("key1")
	if found {
		t.Error("key1 still exists after Remove")
	}
}

func TestPop(t *testing.T) {
	cache := New[string, int]()
	
	// Test non-existent key
	_, found := cache.Pop("nonexistent")
	if found {
		t.Error("popped non-existent key")
	}
	
	// Test existing key
	cache.Set("key1", 42)
	value, found := cache.Pop("key1")
	if !found {
		t.Error("key1 not found during Pop")
	}
	if value != 42 {
		t.Errorf("expected 42, got %d", value)
	}
	
	// Verify key is removed
	_, found = cache.Get("key1")
	if found {
		t.Error("key1 still exists after Pop")
	}
}

func TestConcurrentAccess(t *testing.T) {
	cache := New[int, string]()
	var wg sync.WaitGroup
	
	// Concurrent writes
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(key int) {
			defer wg.Done()
			cache.Set(key, "value")
		}(i)
	}
	
	// Concurrent reads
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(key int) {
			defer wg.Done()
			cache.Get(key)
		}(i)
	}
	
	wg.Wait()
}