package fifo

import (
	"container/list"
	"sync"
)

// Cache is a thread-safe list fixed size (LRU FIFO)
type Cache struct {
	size      int
	evictList *list.List
	lock      sync.RWMutex
	onEvicted func(value interface{})
}

// New creates an LRU of the given size
func new(size int) *Cache {
	return newWithOnEvicted(size, nil)
}

func newWithOnEvicted(size int, onEvicted func(value interface{})) *Cache {
	if size <= 0 {
		return nil
	}
	c := &Cache{
		size:      size,
		evictList: list.New(),
		onEvicted: onEvicted,
	}
	return c
}

// Purge is used to completely clear the fifo
func (c *Cache) purge() {
	c.lock.Lock()
	defer c.lock.Unlock()

	if c.onEvicted != nil {
		for _, v := range c.getElements() {
			c.onEvicted(v.Value)
		}
	}
	c.evictList = list.New()
}

// Add adds a value to the fifo.  Returns true if an eviction occured.
func (c *Cache) add(value interface{}) bool {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.evictList.PushFront(value)
	evict := c.evictList.Len() > c.size
	// Verify size not exceeded
	if evict {
		c.removeOldest()
	}
	return evict
}

// Keys returns a slice of the keys in the fifo, from oldest to newest.
func (c *Cache) getElements() []*list.Element {
	c.lock.RLock()
	defer c.lock.RUnlock()

	keys := make([]*list.Element, c.evictList.Len())
	ent := c.evictList.Back()
	i := 0
	for ent != nil {
		keys[i] = ent
		ent = ent.Prev()
		i++
	}
	return keys
}

// Len returns the number of items in the fifo
func (c *Cache) len() int {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.evictList.Len()
}

// removeOldest removes the oldest item from the fifo
func (c *Cache) removeOldest() *list.Element {
	c.lock.Lock()
	defer c.lock.Unlock()
	ent := c.evictList.Back()
	if ent != nil {
		c.removeElement(ent)
	}
	return ent
}

// removeElement is used to remove a given list element from the fifo
func (c *Cache) removeElement(e *list.Element) {
	c.evictList.Remove(e)
	if c.onEvicted != nil {
		c.onEvicted(e.Value)
	}
}
