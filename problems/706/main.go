package main

type Node struct {
	key  int
	val  int
	next *Node
}

type MyHashMap struct {
	buckets []*Node
	size    int
}

func Constructor() MyHashMap {
	return MyHashMap{
		buckets: make([]*Node, 4093), // use prime numbers to avoid colision and create multiple nodes in the linked list
		size:    4093,
	}
}

func (m *MyHashMap) Put(key int, value int) {
	index := key % m.size

	// means no value is stored in this bucket
	if m.buckets[index] == nil {
		m.buckets[index] = &Node{
			key:  key,
			val:  value,
			next: nil,
		}
		return
	}

	curr := m.buckets[index]
	for curr != nil {
		if curr.key == key {
			curr.val = value // update if the key exists
			return
		}

		if curr.next == nil {
			break
		}

		// move to the next node
		curr = curr.next
	}

	// if the key is not found, add to the end of the linked list
	curr.next = &Node{
		key:  key,
		val:  value,
		next: nil,
	}
}

func (m *MyHashMap) Get(key int) int {
	index := key % m.size

	curr := m.buckets[index]

	for curr != nil {
		if curr.key == key {
			return curr.val
		}

		curr = curr.next
	}

	return -1
}

func (m *MyHashMap) Remove(key int) {
	index := key % m.size
	curr := m.buckets[index]
	if curr == nil {
		return
	}

	if curr.key == key {
		// since i dropped the first node pointer, the GB will delete it from the heap later.
		m.buckets[index] = curr.next
		return
	}

	prev := curr
	curr = curr.next // Move curr to the second node
	for curr != nil {
		if curr.key == key {
			prev.next = curr.next // Remove the node by skipping it
			return                // Return after removal
		}
		prev = curr
		curr = curr.next
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
