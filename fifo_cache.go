package go_s3_fifo

import "container/list"

func NewFIFOCache[K comparable, V any](maxCapacity int) *FIFOCache[K, V] {
	return &FIFOCache[K, V]{
		maxCapacity: maxCapacity,
		capacity:    0,
		cache:       make(map[K]*Entry[K, V]),
		list:        list.New(),
	}
}

func (c *FIFOCache[K, V]) Set(key K, value V) *Entry[K, V] {
	if c.capacity == c.maxCapacity {
		e := c.list.Back()
		delete(c.cache, e.Value.(*Entry[K, V]).k)
		c.list.Remove(e)
		c.capacity--
		return e.Value.(*Entry[K, V])
	}

	et := Entry[K, V]{k: key, v: value}
	e := c.list.PushFront(&et)
	et.element = e
	c.cache[key] = e.Value.(*Entry[K, V])
	c.capacity++
	return nil
}
func (c *FIFOCache[K, V]) SetFromG(key K, value V, freq int32) *Entry[K, V] {
	if c.capacity == c.maxCapacity {
		e := c.list.Back()
		delete(c.cache, e.Value.(*Entry[K, V]).k)
		c.list.Remove(e)
		c.capacity--
		return e.Value.(*Entry[K, V])
	}

	et := Entry[K, V]{k: key, v: value, freq: freq}
	e := c.list.PushFront(&et)
	et.element = e
	c.cache[key] = e.Value.(*Entry[K, V])
	c.capacity++
	return nil
}

func (c *FIFOCache[K, V]) Get(key K) (value V, ok bool) {
	if e, ok := c.cache[key]; ok {
		return e.v, true
	}
	return value, false
}

func (c *FIFOCache[K, V]) Remove(key K) (ok bool) {
	if e, ok := c.cache[key]; ok {
		delete(c.cache, key)
		c.list.Remove(e.element)
		c.capacity--
		return true
	}
	return false
}

func (c *FIFOCache[K, V]) Contains(key K) (ok bool) {
	_, ok = c.cache[key]
	return ok
}

func (c *FIFOCache[K, V]) ContainsWithFreq(key K) (ok bool, freq int32) {
	e, ok := c.cache[key]
	if ok {
		freq = e.freq
	}
	return ok, freq
}

func (c *FIFOCache[K, V]) Capacity() int {
	return c.capacity
}

func (c *FIFOCache[K, V]) IncrFreq(key K) {
	e, ok := c.cache[key]
	if !ok {
		return
	}
	if e.freq < 3 {
		e.Incr()
	}
}
