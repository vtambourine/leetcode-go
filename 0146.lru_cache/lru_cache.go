package main

import "fmt"

type LRUCache struct {
	cache    map[int]int
	queue    []int
	capacity int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		cache:    make(map[int]int),
		queue:    make([]int, capacity),
		capacity: capacity,
	}
}

func (c *LRUCache) Get(key int) int {
	if v, ok := c.cache[key]; ok {
		for i, k := range c.queue {
			if k == key {
				c.queue = append(c.queue[:i], c.queue[i+1:]...)
				c.queue = append(c.queue, key)
				break
			}
		}
		return v
	} else {
		return -1
	}
}

func (c *LRUCache) Put(key int, value int) {
	//fmt.Println("Put", key, value)
	if _, ok := c.cache[key]; ok {
		//fmt.Println("+ hit cache", c.cache)
		c.cache[key] = value
		for i, k := range c.queue {
			if k == key {
				c.queue = append(c.queue[:i], c.queue[i+1:]...)
				c.queue = append(c.queue, key)
				break
			}
		}
	} else {
		//fmt.Println("+ miss cache", c.cache)
		c.cache[key] = value
		c.queue = append(c.queue, key)
		if len(c.queue) > c.capacity {
			old := c.queue[0]
			delete(c.cache, old)
			c.queue = c.queue[1:]
		}
	}
}

func main() {
	cache := Constructor(2)
	//cache.Put(1, 1)
	//cache.Put(2, 2)
	//fmt.Println(cache.Get(1)) // returns 1
	//cache.Put(3, 3)  // evicts key 2
	//fmt.Println(cache.Get(2)) // returns -1 (not found)
	//cache.Put(4, 4)  // evicts key 1
	//fmt.Println(cache.Get(1)) // returns -1 (not found)
	//fmt.Println(cache.Get(3)) // returns 3

	cache.Put(2, 1)
	cache.Put(2, 2)
	fmt.Println(cache.Get(2)) // 2
	cache.Put(1, 1)
	cache.Put(4, 1)
	fmt.Println(cache.Get(2)) // -1
}
