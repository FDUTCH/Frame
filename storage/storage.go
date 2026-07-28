package storage

import (
	"reflect"
)

// Storage stores typed data.
type Storage struct {
	storage map[reflect.Type]any
}

// NewStorage create new storage.
func NewStorage() *Storage {
	return &Storage{storage: make(map[reflect.Type]any)}
}

// Set sets value.
func Set[T any](s *Storage, val T) {
	s.storage[reflect.TypeOf(val)] = val
}

// Get returns the value.
func Get[T any](s *Storage) (T, bool) {
	val, ok := s.storage[reflect.TypeOf((*T)(nil))]
	return val, ok
}

// GetOr returns the value or a value 'or'.
func GetOr[T any](s *Storage, or T) T {
	val, ok := s.storage[reflect.TypeOf((*T)(nil))]
	if !ok {
		return or
	}
	return val
}
