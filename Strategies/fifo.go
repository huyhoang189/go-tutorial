package Strategies

import "fmt"

type Fifo struct {
}

func (f *Fifo) evict(c *Cache) {
	fmt.Println("Eviction by fifo strategy")
}
