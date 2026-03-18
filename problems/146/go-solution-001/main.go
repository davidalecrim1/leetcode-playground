package main

// Use a double linked list for the ordering and O(1) sort on usage
// Use a hashmap for O(1) lookup
type Node struct {
	key  int
	val  int
	prev *Node
	next *Node
}

type LRUCache struct {
	hm   map[int]*Node
	head *Node
	tail *Node
	cap  int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		hm:  make(map[int]*Node, capacity),
		cap: capacity,
	}
}

func (c *LRUCache) remove(node *Node) {
	if node.prev != nil {
		node.prev.next = node.next
	} else {
		c.head = node.next
	}

	if node.next != nil {
		node.next.prev = node.prev
	} else {
		c.tail = node.prev
	}

	node.prev = nil
	node.next = nil
}

func (c *LRUCache) insertAtHead(node *Node) {
	node.next = c.head
	node.prev = nil

	if c.head != nil {
		c.head.prev = node
	}

	c.head = node

	if c.tail == nil {
		c.tail = node
	}
}

func (c *LRUCache) Get(key int) int {
	if node, ok := c.hm[key]; ok {
		c.remove(node)
		c.insertAtHead(node)
		return node.val
	}
	return -1
}

func (c *LRUCache) Put(key int, value int) {
	if node, ok := c.hm[key]; ok {
		node.val = value
		c.remove(node)
		c.insertAtHead(node)
		return
	}

	if len(c.hm) == c.cap {
		delete(c.hm, c.tail.key)
		c.remove(c.tail)
	}

	node := &Node{key: key, val: value}
	c.insertAtHead(node)
	c.hm[key] = node
}
