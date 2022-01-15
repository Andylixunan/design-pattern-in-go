package strategy

import "fmt"

type EvictionAlgo interface {
	Evict()
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

func (f FIFO) Evict() {
	fmt.Println("Evicting by fifo strategy")
}

type LRU struct{}

func (l LRU) Evict() {
	fmt.Println("Evicting by lru strategy")
}
