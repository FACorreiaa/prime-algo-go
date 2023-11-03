package day_maps

type CacheNode[T comparable] struct {
	Value T
	Next  *CacheNode[T]
	Prev  *CacheNode[T]
}

func NewCacheNode[T comparable](value T) CacheNode[T] {
	return CacheNode[T]{Value: value}
}

type LRU[K comparable, V comparable] struct {
	Length        int
	Head, Tail    *CacheNode[V]
	Lookup        map[K]*CacheNode[V]
	ReverseLookup map[CacheNode[V]]*K
	Capacity      int
}

func NewLRU[K, V comparable](capacity int) *LRU[K, V] {
	return &LRU[K, V]{
		Length:        0,
		Head:          nil,
		Tail:          nil,
		Lookup:        make(map[K]*CacheNode[V]),
		ReverseLookup: make(map[CacheNode[V]]*K),
		Capacity:      capacity,
	}
}

func (lru *LRU[K, V]) Set(key K, value V) {
	node := lru.Lookup[key]
	if node == nil {
		node := NewCacheNode(value)
		lru.Length++
		lru.trimCache()
		lru.Lookup[key] = &node
		lru.ReverseLookup[node] = &key
	} else {
		lru.detach(node)
		lru.prepend(node)
		node.Value = value
	}
	//if it doesn't
	// - check capacity and evict if over
	//if it does, we need to update to the front of the list
	//update the value
}

func (lru *LRU[K, V]) Get(key K) V {
	var zero V
	node := lru.Lookup[key]
	if node == nil {
		return zero
	}
	//update the value we found and move it to the front
	lru.detach(node)
	lru.prepend(node)

	return node.Value
}

func (lru *LRU[K, V]) detach(node *CacheNode[V]) {
	if node.Prev != nil {
		node.Prev.Next = node.Next
	}

	if node.Next != nil {
		node.Next.Prev = node.Prev
	}

	if lru.Length == 1 {
		lru.Tail, lru.Head = nil, nil
	}

	if lru.Head == node {
		lru.Head = lru.Head.Next
	}

	if lru.Tail == node {
		lru.Tail = lru.Tail.Prev
	}

	node.Next, node.Prev = nil, nil
}

func (lru *LRU[K, V]) prepend(node *CacheNode[V]) {
	if lru.Head == nil {
		lru.Head, lru.Tail = node, node
		return
	}

	node.Next = lru.Head
	lru.Head.Prev = node
	lru.Head = node
}

func (lru *LRU[K, V]) trimCache() {
	if lru.Length <= lru.Capacity {
		return
	}

	tail := lru.Tail
	lru.detach(lru.Tail)

	key := lru.ReverseLookup[*tail]
	delete(lru.Lookup, *key)
	delete(lru.ReverseLookup, *tail)
	lru.Length--
}
