package strategy

import "fmt"

type EvictionAlgo interface {
	Evict(c *Cache)
}

type Cache struct {
	EvictionAlgo EvictionAlgo
	Storage      map[string]string
}

// To initialize a new cache, we need our clients to specify
// a concrete eviction algorithm
func NewCache(e EvictionAlgo) *Cache {
	return &Cache{
		EvictionAlgo: e,
		Storage:      make(map[string]string),
	}
}

func (c *Cache) SetEvictionAlgo(e EvictionAlgo) {
	c.EvictionAlgo = e
}

// FIFO and LRU struct implement EvictionAlgo interface

type FIFO struct{}

var _ EvictionAlgo = FIFO{}

func (f FIFO) Evict(c *Cache) {
	fmt.Printf("Evicting by fifo strategy: %+v \n", c.Storage)
}

type LRU struct{}

var _ EvictionAlgo = LRU{}

func (l LRU) Evict(c *Cache) {
	fmt.Printf("Evicting by lru strategy: %+v \n", c.Storage)
}
