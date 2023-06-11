package Strategies

type EvictionAlgo interface {
	evict(c *Cache)
}
