// Package syncs provides synchronization primitives util functions.
package syncs

import (
	"context"
	"sync"
)

// WaitGroup is a wrapper of sync.WaitGroup.
//
// Usage:
//
//	wg := syncs.WaitGroup{}
//	wg.Go(func() {
//		time.Sleep(time.Second)
//	})
//	wg.Wait()
type WaitGroup struct {
	sync.WaitGroup
}

// Go runs the given function in a new goroutine. will auto call Add and Done.
func (wg *WaitGroup) Go(fn func()) { _ = "STUB: not implemented"; return }

// ContextValue create a new context with given value
func ContextValue(key, value any) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// SafeMap is a goroutine-safe map.
type SafeMap[K comparable, V any] struct {
	mu   sync.RWMutex
	data map[K]V
}

// NewSafeMap create a new SafeMap instance.
func NewSafeMap[K comparable, V any]() *SafeMap[K, V] { _ = "STUB: not implemented"; return nil }

// Set value to map.
func (m *SafeMap[K, V]) Set(key K, value V) { _ = "STUB: not implemented"; return }

// Get value from map.
func (m *SafeMap[K, V]) Get(key K) (value V, ok bool) {
	_ = "STUB: not implemented"
	return *new(V), false
}

// Delete value from map.
func (m *SafeMap[K, V]) Delete(key K) { _ = "STUB: not implemented"; return }

// Range iterate map values.
func (m *SafeMap[K, V]) Range(fn func(key K, value V)) { _ = "STUB: not implemented"; return }

// Clear do clear all map values.
func (m *SafeMap[K, V]) Clear() { _ = "STUB: not implemented"; return }

// Len get map length.
func (m *SafeMap[K, V]) Len() int { _ = "STUB: not implemented"; return 0 }
