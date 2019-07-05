package main

type Cache struct {
	buckets map[int]int
	refMap map[int]int
	size int
}

func NewCache(size int) *Cache{
	cache := new(Cache)
	cache.buckets = make(map[int]int)
	cache.refMap = make(map[int]int)
	cache.size = size

	return cache
}

func (cache *Cache) Put(index, value int) {
	if cache.isFull() {
		leastUsedIndex := cache.findLeastUsed()
		delete(cache.buckets, leastUsedIndex)
	}

	cache.buckets[index] = value
	cache.refMap[index] = 1 
}

func (cache *Cache) Get(index int) int {
	if _,ok := cache.refMap[index]; !ok {
		return -1 
	} else {
		cache.refMap[index] += 1
		return cache.buckets[index]
	}
}

func (cache *Cache) isFull() bool {
	if len(cache.buckets) == cache.size {
		return true
	}

	return false
}

func (cache *Cache) findLeastUsed() int {
	var leastUsedCount int
	var leastUsedIndex int

	initialized := false
	for k, v := range cache.refMap {
		if !initialized {
			leastUsedCount = v
			leastUsedIndex = k
			initialized = true
		}

		if v < leastUsedCount {
			leastUsedCount = v
			leastUsedIndex = k
		}
	}

	return leastUsedIndex
}