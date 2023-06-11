package Strategies

type Cache struct {
	storage      map[string]string
	evictionAlgo EvictionAlgo
	capacity     int
	maxCapacity  int
}

func InitCache(e EvictionAlgo) *Cache {
	storage := make(map[string]string)
	return &Cache{
		evictionAlgo: e,
		capacity:     0,
		maxCapacity:  2,
		storage:      storage,
	}
}

func (c *Cache) SetEvictionAlgo(e EvictionAlgo) {
	c.evictionAlgo = e
	//fmt.Println("")
}

func (c *Cache) Add(key, value string) {
	if c.capacity == c.maxCapacity {
		c.evictionAlgo.evict(c)
		c.capacity--
	}
	c.capacity++
	c.storage[key] = value
}
