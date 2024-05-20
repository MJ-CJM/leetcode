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
	head := &LinkNode{
		key:  0,
		val:  0,
		pre:  nil,
		next: nil,
	}
	tail := &LinkNode{
		key:  0,
		val:  0,
		pre:  nil,
		next: nil,
	}
	head.next = tail
	tail.pre = head
	lru := LRUCache{
		m:    map[int]*LinkNode{},
		cap:  capacity,
		head: head,
		tail: tail,
	}
	return lru
}

func (this *LRUCache) Get(key int) int {
	cache := this.m
	if v, ok := cache[key]; ok {
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
	if v, ok := cache[key]; ok {
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
	// 移到头部
	head.next.pre = node
	node.next = head.next
	head.next = node
	node.pre = head
}

/*
这段代码实现了LRU（Least Recently Used）缓存的功能，LRU是一种缓存淘汰策略，它会优先淘汰最近最少使用的数据。下面是对代码实现的详细解释：

定义结构体：

LinkNode结构体：用于构建双向链表的节点，包含了键值对中的key和val，以及指向前一个节点和后一个节点的指针pre和next。
LRUCache结构体：LRU缓存结构体，包含了一个map用于存储键值对，缓存的容量cap，以及指向双向链表头尾的指针head和tail。
构造函数 Constructor(capacity int) LRUCache：

初始化LRU缓存，传入参数为容量capacity，创建头部和尾部的哑节点，初始化LRUCache结构体，并返回。
Get(key int) int方法：

根据给定的键key获取对应的值val。
首先从LRU缓存的map中查找键key，如果存在，则将对应的节点移动到链表头部（表示最近访问过），然后返回对应的值val；如果不存在，则返回-1表示未找到。
Put(key int, value int)方法：

存入键值对到LRU缓存中。
首先检查给定的键key是否已存在于LRU缓存中，如果存在，则更新对应节点的值val，并将节点移动到链表头部；如果不存在，则创建新的节点，将其插入到链表头部，并将其加入LRU缓存的map中。若LRU缓存已满（即缓存的大小达到了容量上限），则删除最近最少使用的节点，即链表尾部的节点，并在map中也删除对应的键值对。
MoveToHead(node *LinkNode)方法：

将指定节点移动到链表头部。
首先从链表中移除该节点，然后将其插入到链表头部，表示最近被访问过。
通过这样的实现，LRU缓存能够实现快速存取数据，并在容量不足时按照最近使用情况淘汰数据，保持缓存的有效性。
 */