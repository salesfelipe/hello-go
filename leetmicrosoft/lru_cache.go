package leetmicrosoft

type LRUCache struct {
	capacity int
	cache    map[int]int
	queue    Queue
}

type Queue struct {
	length int
	head   *DoubleLinkedNode
	tail   *DoubleLinkedNode
}

type DoubleLinkedNode struct {
	value int
	prev  *DoubleLinkedNode
	next  *DoubleLinkedNode
}

func (q *Queue) Enqueue(item int) {
	newNode := DoubleLinkedNode{value: item}

	if q.tail != nil {
		newNode.prev = q.tail
		q.tail.next = &newNode
	}

	q.tail = &newNode

	if q.head == nil {
		q.head = &newNode
	}

	q.length += 1
}

func (q *Queue) Dequeue() int {
	lastItemAdded := q.head

	q.head = lastItemAdded.next
	q.head.prev = nil
	q.length -= 1

	return lastItemAdded.value
}

func (q *Queue) SearchNRemove(item int) {
	currentItem := q.head

	for currentItem != nil {
		next := currentItem.next

		if currentItem.value == item {
			prev := currentItem.prev
			q.length -= 1

			if prev != nil {
				prev.next = next
			}

			if next != nil {
				next.prev = prev
			}

			if q.head.value == item {
				q.head = next
			}

			if q.tail.value == item {
				q.tail = prev
			}

			return
		}

		currentItem = next
	}
}

func Constructor(capacity int) LRUCache {
	return LRUCache{capacity: capacity, cache: make(map[int]int), queue: Queue{}}
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
	if l.queue.tail != nil && l.queue.tail.value == key {
		return
	}

	if len(l.cache) > l.capacity {
		last := l.queue.Dequeue()
		delete(l.cache, last)
	}

	l.queue.SearchNRemove(key)
	l.queue.Enqueue(key)
}
