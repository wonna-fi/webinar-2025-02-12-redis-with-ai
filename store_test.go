// This file contains AI generated code that has not been reviewed by a human

package main

import (
	"fmt"
	"testing"
)

func TestStore(t *testing.T) {
	store := NewStore()

	// Test Set and Get
	store.Set("key1", "value1")
	value, ok := store.Get("key1")
	if !ok {
		t.Error("Get returned not ok for existing key")
	}
	if value != "value1" {
		t.Errorf("Get returned wrong value: got %v want %v", value, "value1")
	}

	// Test Get non-existent key
	_, ok = store.Get("nonexistent")
	if ok {
		t.Error("Get returned ok for non-existent key")
	}

	// Test overwrite value
	store.Set("key1", "newvalue")
	value, ok = store.Get("key1")
	if !ok {
		t.Error("Get returned not ok for existing key after overwrite")
	}
	if value != "newvalue" {
		t.Errorf("Get returned wrong value after overwrite: got %v want %v", value, "newvalue")
	}

	// Test Delete
	ok = store.Delete("key1")
	if !ok {
		t.Error("Delete returned not ok for existing key")
	}
	_, ok = store.Get("key1")
	if ok {
		t.Error("Get returned ok for deleted key")
	}

	// Test Delete non-existent key
	ok = store.Delete("nonexistent")
	if ok {
		t.Error("Delete returned ok for non-existent key")
	}
}

func TestStoreConcurrent(t *testing.T) {
	store := NewStore()
	done := make(chan bool)

	// Start multiple goroutines to write
	for i := 0; i < 10; i++ {
		go func(n int) {
			key := fmt.Sprintf("key%d", n)
			store.Set(key, "value")
			done <- true
		}(i)
	}

	// Wait for all writes to complete
	for i := 0; i < 10; i++ {
		<-done
	}

	// Verify all writes succeeded
	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("key%d", i)
		_, ok := store.Get(key)
		if !ok {
			t.Errorf("Get returned not ok for key %v", key)
		}
	}
}
