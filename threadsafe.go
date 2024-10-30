package lru

import (
	"iter"
	"sync"
)

var _ LRU[struct{}, struct{}] = (*threadsafe[struct{}, struct{}])(nil)

type threadsafe[K comparable, V any] struct {
	lru lru[K, V]
	mu  sync.RWMutex
}

func NewThreadSafe[K comparable, V any](capacity int, evicted ...func(key K, val V)) LRU[K, V] {
	c := lru[K, V]{
		keys:    make([]K, 0, capacity),
		vals:    make([]V, 0, capacity),
		lastUse: make([]uint64, 0, capacity),
	}

	if len(evicted) > 0 {
		c.evicted = evicted[0]
	}

	return &threadsafe[K, V]{
		lru: c,
	}
}

// Cap implements LRU.
func (t *threadsafe[K, V]) Cap() int {
	t.mu.RLock()
	defer t.mu.RUnlock()

	return t.lru.Cap()
}

// Get implements LRU.
func (t *threadsafe[K, V]) Get(key K) (val V, ok bool) {
	t.mu.RLock()
	defer t.mu.RUnlock()

	return t.lru.Get(key)
}

// GetOrSet implements LRU.
func (t *threadsafe[K, V]) GetOrSet(key K, setter func(K) (V, error)) (val V, err error) {
	t.mu.Lock()
	defer t.mu.Unlock()

	return t.lru.GetOrSet(key, setter)
}

// Has implements LRU.
func (t *threadsafe[K, V]) Has(key K) (ok bool) {
	t.mu.RLock()
	defer t.mu.RUnlock()

	return t.lru.Has(key)
}

// Iterate all items in no particular order.
func (t *threadsafe[K, V]) Iterate() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		t.mu.RLock()
		defer t.mu.RUnlock()

		for i := range t.lru.keys {
			if !yield(t.lru.keys[i], t.lru.vals[i]) {
				return
			}
		}
	}
}

// Iterate all items in ascending order.
func (t *threadsafe[K, V]) IterateAsc() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		t.mu.RLock()
		defer t.mu.RUnlock()

		tick := t.lru.oldestTick()

		for range t.lru.keys {
			idx, ok := t.lru.find(tick)

			if !ok || !yield(t.lru.keys[idx], t.lru.vals[idx]) {
				return
			}

			tick++
		}
	}
}

// Iterate all items in descending order.
func (t *threadsafe[K, V]) IterateDesc() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		t.mu.RLock()
		defer t.mu.RUnlock()

		tick := t.lru.tick - 1

		for range t.lru.keys {
			idx, ok := t.lru.find(tick)

			if !ok || !yield(t.lru.keys[idx], t.lru.vals[idx]) {
				return
			}

			tick--
		}
	}
}

// Len implements LRU.
func (t *threadsafe[K, V]) Len() int {
	t.mu.RLock()
	defer t.mu.RUnlock()

	return t.lru.Len()
}

// Remove implements LRU.
func (t *threadsafe[K, V]) Remove(key K) (existed bool) {
	t.mu.Lock()
	defer t.mu.Unlock()

	return t.lru.Remove(key)
}

// RemoveAll implements LRU.
func (t *threadsafe[K, V]) RemoveAll() {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.lru.RemoveAll()
}

// Replace implements LRU.
func (t *threadsafe[K, V]) Replace(key K, val V) (existed bool) {
	t.mu.Lock()
	defer t.mu.Unlock()

	return t.lru.Replace(key, val)
}

// Reset implements LRU.
func (t *threadsafe[K, V]) Reset() {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.lru.Reset()
}

// Resize implements LRU.
func (t *threadsafe[K, V]) Resize(capacity int) {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.lru.Resize(capacity)
}

// Set implements LRU.
func (t *threadsafe[K, V]) Set(key K, val V) (ok bool) {
	t.mu.Lock()
	defer t.mu.Unlock()

	return t.lru.Set(key, val)
}
