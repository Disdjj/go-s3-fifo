package go_s3_fifo

func NewS3FIFOCache[K comparable, V any](capacity int) *S3FIFOCache[K, V] {
	return &S3FIFOCache[K, V]{
		capacity: capacity,
		s:        NewFIFOCache[K, V](capacity / 10),
		m:        NewFIFOCache[K, V](capacity - capacity/10),
		g:        NewFIFOCache[K, struct{}](capacity),
	}
}

func (c *S3FIFOCache[K, V]) Set(key K, value V) {
	// if key exists in g, remove it from g and set it in m
	if ok, f := c.g.ContainsWithFreq(key); ok {
		c.g.Remove(key)
		var freq = f
		e := c.m.SetFromG(key, value, freq)
		// e.state flow
		if e.freq > 1 {
			e.freq -= 1
			c.m.list.MoveToFront(e.element)
		}
		return
	}

	// if key exists in s, Incr freq
	if c.s.Contains(key) {
		c.s.IncrFreq(key)
		return
	}

	// if key exists in m, Incr freq
	if c.m.Contains(key) {
		c.m.IncrFreq(key)
		return
	}

	// add to s
	e := c.s.Set(key, value)
	if e != nil {
		// e.state flow
		if e.freq > 1 {
			c.m.list.MoveToFront(e.element)
		}
	}
}

func (c *S3FIFOCache[K, V]) Get(key K) (value V, ok bool) {
	if value, ok = c.s.Get(key); ok {
		c.s.IncrFreq(key)
		return
	}

	if value, ok = c.m.Get(key); ok {
		c.m.IncrFreq(key)
		return
	}

	if ok = c.g.Contains(key); ok {
		c.g.IncrFreq(key)
	}
	return value, ok
}
