package main

type LinkNode struct {
	key, val  int
	pre, next *LinkNode
}

type LRUCache struct {
	m          map[int]*LinkNode
	cap        int
	head, tail *LinkNode
}

func Constructor(capacity int) LRUCache {
	head := &LinkNode{0, 0, nil, nil}
	tail := &LinkNode{0, 0, nil, nil}
	head.next = tail
	tail.pre = head
	lruCache := LRUCache{
		m:    make(map[int]*LinkNode),
		cap:  capacity,
		head: head,
		tail: tail,
	}
	return lruCache
}

func (this *LRUCache) Get(key int) int {
	cache := this.m
	if v, exist := cache[key]; exist {
		this.MoveToHead(v)
		return v.val
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	head := this.head
	tail := this.tail
	cache := this.m
	if v, exist := cache[key]; exist {
		v.val = value
		this.MoveToHead(v)
	} else {
		v := &LinkNode{
			key:  key,
			val:  value,
			pre:  nil,
			next: nil,
		}
		if len(cache) == this.cap {
			//删除最后元素，map中删除，链表中也删除
			delete(cache, tail.pre.key)
			tail.pre.pre.next = tail
			tail.pre = tail.pre.pre
		}
		v.next = head.next
		v.pre = head
		head.next.pre = v
		head.next = v
		cache[key] = v
	}
}

func (this *LRUCache) MoveToHead(node *LinkNode) {
	head := this.head
	node.pre.next = node.next
	node.next.pre = node.pre
	// 移到首位
	node.next = head.next
	head.next.pre = node
	node.pre = head
	head.next = node
}
