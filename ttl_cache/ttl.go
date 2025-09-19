// DANGER: work with Gemini LLM model!!!

package ttl_cache

import (
	"sync"
	"sync/atomic"
	"time"
)

type Stats struct {
	Hits uint64
	Miss uint64
}

type cachedItem struct {
	value     interface{}
	expiresAt int64
}

type TTLCache struct {
	data map[string]cachedItem
	mu   sync.RWMutex

	hits atomic.Uint64
	miss atomic.Uint64

	cleanupInterval time.Duration
	stop            chan struct{}
}

type Option func(*TTLCache)

func WithCleanupInterval(interval time.Duration) Option {
	return func(c *TTLCache) {
		c.cleanupInterval = interval
	}
}

func NewTTLCache(options ...Option) *TTLCache {
	c := &TTLCache{
		data:            make(map[string]cachedItem),
		cleanupInterval: 1 * time.Minute,
		stop:            make(chan struct{}),
	}

	for _, opt := range options {
		opt(c)
	}

	if c.cleanupInterval > 0 {
		go c.startCleanup()
	}

	return c
}

func (c *TTLCache) Stop() {
	close(c.stop)
}

func (c *TTLCache) Set(key string, value interface{}, ttl time.Duration) {
	var expiresAt int64
	if ttl > 0 {
		expiresAt = time.Now().Add(ttl).UnixNano()
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[key] = cachedItem{
		value:     value,
		expiresAt: expiresAt,
	}
}

func (c *TTLCache) Get(key string) (interface{}, bool) {
	c.mu.RLock()

	item, exists := c.data[key]
	if !exists {
		c.mu.RUnlock()
		c.miss.Add(1)
		return nil, false
	}

	// запись в кеше протухла
	if item.expiresAt > 0 && time.Now().UnixNano() > item.expiresAt {
		c.mu.RUnlock()

		c.mu.Lock()
		delete(c.data, key)
		c.mu.Unlock()

		c.miss.Add(1)
		return nil, false
	}

	c.mu.RUnlock()
	c.hits.Add(1)
	return item.value, true
}

func (c *TTLCache) Delete(key string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	_, exists := c.data[key]
	delete(c.data, key)
	return exists
}

func (c *TTLCache) Stats() Stats {
	return Stats{
		Hits: c.hits.Load(),
		Miss: c.miss.Load(),
	}
}

func (c *TTLCache) ResetStats() {
	c.hits.Store(0)
	c.miss.Store(0)
}

func (c *TTLCache) Size() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return len(c.data)
}

func (c *TTLCache) startCleanup() {
	ticker := time.NewTicker(c.cleanupInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			c.deleteExpired()
		case <-c.stop:
			return
		}
	}
}

func (c *TTLCache) deleteExpired() {
	now := time.Now().UnixNano()

	// на время чистки, будет все заблокировано
	c.mu.Lock()
	defer c.mu.Unlock()

	for key, item := range c.data {
		if item.expiresAt > 0 && now > item.expiresAt {
			delete(c.data, key)
		}
	}
}
