// This file contains AI generated code that has not been reviewed by a human

package main

import "sync"

// Store represents a thread-safe key-value store
type Store struct {
	data sync.Map
}

// NewStore creates a new key-value store
func NewStore() *Store {
	return &Store{}
}

// Get retrieves a value from the store
func (s *Store) Get(key string) (string, bool) {
	value, ok := s.data.Load(key)
	if !ok {
		return "", false
	}
	return value.(string), true
}

// Set stores a value in the store
func (s *Store) Set(key, value string) {
	s.data.Store(key, value)
}

// Delete removes a value from the store
func (s *Store) Delete(key string) bool {
	_, ok := s.data.LoadAndDelete(key)
	return ok
}
