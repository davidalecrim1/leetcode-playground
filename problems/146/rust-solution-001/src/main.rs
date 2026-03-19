fn main() {
    let mut cache = LRUCache::new(2);
    cache.put(1, 1);
    cache.put(2, 2);
    println!("{}", cache.get(1)); // 1
    cache.put(3, 3); // evicts key 2
    println!("{}", cache.get(2)); // -1
    cache.put(4, 4); // evicts key 1
    println!("{}", cache.get(1)); // -1
    println!("{}", cache.get(3)); // 3
    println!("{}", cache.get(4)); // 4
}

use std::collections::HashMap;

const NULL: usize = usize::MAX;

struct Node {
    key: i32,
    val: i32,
    prev: usize,
    next: usize,
}

struct LRUCache {
    nodes: Vec<Node>,
    map: HashMap<i32, usize>, // key -> index into nodes
    head: usize,
    tail: usize,
    cap: usize,
}

impl LRUCache {
    fn new(capacity: i32) -> Self {
        Self {
            nodes: Vec::with_capacity(capacity as usize),
            map: HashMap::new(),
            head: NULL,
            tail: NULL,
            cap: capacity as usize,
        }
    }

    fn get(&mut self, key: i32) -> i32 {
        if let Some(&idx) = self.map.get(&key) {
            let val = self.nodes[idx].val;
            self.move_to_head(idx);
            return val;
        }
        -1
    }

    fn put(&mut self, key: i32, value: i32) {
        if let Some(&idx) = self.map.get(&key) {
            self.nodes[idx].val = value;
            self.move_to_head(idx);
            return;
        }

        if self.map.len() == self.cap {
            self.evict_lru();
        }

        let idx = self.nodes.len();
        self.nodes.push(Node { key, val: value, prev: NULL, next: NULL });
        self.map.insert(key, idx);
        self.insert_at_head(idx);
    }

    fn remove_node(&mut self, idx: usize) {
        let prev = self.nodes[idx].prev;
        let next = self.nodes[idx].next;

        if prev == NULL {
            self.head = next;
        } else {
            self.nodes[prev].next = next;
        }

        if next == NULL {
            self.tail = prev;
        } else {
            self.nodes[next].prev = prev;
        }

        self.nodes[idx].prev = NULL;
        self.nodes[idx].next = NULL;
    }

    fn insert_at_head(&mut self, idx: usize) {
        self.nodes[idx].next = self.head;
        self.nodes[idx].prev = NULL;

        if self.head != NULL {
            self.nodes[self.head].prev = idx;
        }

        self.head = idx;

        if self.tail == NULL {
            self.tail = idx;
        }
    }

    fn move_to_head(&mut self, idx: usize) {
        self.remove_node(idx);
        self.insert_at_head(idx);
    }

    fn evict_lru(&mut self) {
        if self.tail == NULL {
            return;
        }
        let tail = self.tail;
        let key = self.nodes[tail].key;
        self.remove_node(tail);
        self.map.remove(&key);
        // Swap-remove to reclaim the slot; fix up the swapped node's index in map/list
        let last = self.nodes.len() - 1;
        if tail != last {
            self.nodes.swap(tail, last);
            let swapped_key = self.nodes[tail].key;
            self.map.insert(swapped_key, tail);
            let prev = self.nodes[tail].prev;
            let next = self.nodes[tail].next;
            if prev == NULL {
                self.head = tail;
            } else {
                self.nodes[prev].next = tail;
            }
            if next == NULL {
                self.tail = tail;
            } else {
                self.nodes[next].prev = tail;
            }
        }
        self.nodes.pop();
    }
}
