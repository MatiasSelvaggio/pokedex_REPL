package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	data     map[string]cacheEntry
	mu       sync.RWMutex
	interval time.Duration
	stopChan chan struct{}
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		data:     make(map[string]cacheEntry),
		interval: interval,
		stopChan: make(chan struct{}),
	}

	go c.reapLoop()

	return c
}

func (c *Cache) Add(key string, value []byte) {
	c.mu.Lock()
	c.data[key] = cacheEntry{
		val:       value,
		createdAt: time.Now(),
	}
	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	data, ok := c.data[key]
	if !ok {
		return nil, false
	}
	return data.val, true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			c.reapOldEntries()
		case <-c.stopChan:
			return
		}
	}

}

func (c *Cache) reapOldEntries() {
	c.mu.Lock()
	defer c.mu.Unlock()

	now := time.Now()
	for key, entry := range c.data {
		if now.Sub(entry.createdAt) > c.interval {
			delete(c.data, key)
		}
	}
}

func (c *Cache) Stop() {
	close(c.stopChan)
}
