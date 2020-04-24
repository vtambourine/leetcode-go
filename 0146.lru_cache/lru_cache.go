package lru_cache

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
	}
	return -1
}

func (c *LRUCache) Put(key int, value int) {
	if _, ok := c.cache[key]; ok {
		c.cache[key] = value
		for i, k := range c.queue {
			if k == key {
				c.queue = append(c.queue[:i], c.queue[i+1:]...)
				break
			}
		}
	} else {
		c.cache[key] = value
		if len(c.queue) == c.capacity {
			old := c.queue[0]
			delete(c.cache, old)
			c.queue = c.queue[1:]
		}
	}
	c.queue = append(c.queue, key)
}
