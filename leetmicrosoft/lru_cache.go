package leetmicrosoft

type LRUCache struct {
	capacity int
	cache    map[int]int
	accesses []int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{capacity: capacity, cache: make(map[int]int), accesses: make([]int, 0)}
}

func (l *LRUCache) Get(key int) int {
	val, ok := l.cache[key]

	if ok {
		l.visit(key)

		return val
	}

	return -1
}

func (l *LRUCache) Put(key int, value int) {
	l.cache[key] = value

	l.visit(key)
}

func (l *LRUCache) visit(key int) {

	// is last access
	if len(l.accesses) > 0 && l.accesses[0] == key {
		return
	}

	if len(l.cache) > l.capacity {
		last := l.accesses[l.capacity-1]
		delete(l.cache, last)
	}

	filtered := make([]int, 1)
	filtered[0] = key

	for i := 0; i < len(l.accesses); i++ {
		if l.accesses[i] != key {
			filtered = append(filtered, l.accesses[i])
		}
	}

	l.accesses = filtered
}
