package go_s3_fifo

import (
	"container/list"
	"sync/atomic"
)

type Cache[K comparable, V any] interface {
	Set(key K, value V)
	Get(key K) (value V, ok bool)
	Remove(key K) (ok bool)
	Contains(key K) (ok bool)
	Capacity() int
}

type Entry[K comparable, V any] struct {
	k       K
	v       V
	element *list.Element
	freq    int32
}

func (e *Entry[K, V]) Incr() {
	atomic.AddInt32(&e.freq, 1)
}

type S3FIFOCacheWithG[K comparable, V any] struct {
	S3FIFOCache[K, V]
	g *FIFOCache[K, V]
}

type FIFOCache[K comparable, V any] struct {
	maxCapacity int
	capacity    int
	cache       map[K]*Entry[K, V]
	// cache map[K]*Entry[K, V]
	list *list.List
}

type S3FIFOCache[K comparable, V any] struct {
	capacity int

	s *FIFOCache[K, V]
	m *FIFOCache[K, V]
	g *FIFOCache[K, struct{}]
}
