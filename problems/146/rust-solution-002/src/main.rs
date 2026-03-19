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

// Use a doubly linked list and a hashmap
use std::cell::RefCell;
use std::collections::HashMap;
use std::rc::{Rc, Weak};

struct Node {
    key: i32,
    val: i32,
    prev: Option<Weak<RefCell<Node>>>,
    next: Option<Rc<RefCell<Node>>>,
}

struct LRUCache {
    map: HashMap<i32, Rc<RefCell<Node>>>,
    head: Option<Rc<RefCell<Node>>>,
    tail: Option<Rc<RefCell<Node>>>,
    len: i32,
    cap: i32,
}

impl LRUCache {
    fn new(capacity: i32) -> Self {
        Self {
            map: HashMap::new(),
            head: None,
            tail: None,
            cap: capacity,
            len: 0,
        }
    }

    fn get(&mut self, key: i32) -> i32 {
        if let Some(node) = self.map.get(&key).cloned() {
            let val = node.borrow().val;
            self.remove_node(node.clone());
            self.insert_at_head(node);
            return val;
        }
        -1
    }

    fn put(&mut self, key: i32, value: i32) {
        if let Some(node) = self.map.get(&key).cloned() {
            node.borrow_mut().val = value;
            self.remove_node(node.clone());
            self.insert_at_head(node);
            return;
        }

        if self.len == self.cap {
            self.evict_lru();
        }

        let node = Rc::new(RefCell::new(Node {
            key,
            val: value,
            prev: None,
            next: None,
        }));

        self.map.insert(key, node.clone());
        self.insert_at_head(node);
    }

    fn remove_node(&mut self, node: Rc<RefCell<Node>>) {
        let prev = node.borrow().prev.clone();
        let next = node.borrow().next.clone();

        if let Some(prev_weak) = &prev {
            if let Some(prev_rc) = prev_weak.upgrade() {
                prev_rc.borrow_mut().next = next.clone();
            }
        } else {
            // node is the head
            self.head = next.clone();
        }

        if let Some(next_rc) = &next {
            next_rc.borrow_mut().prev = prev.clone();
        } else {
            // node is the tail
            self.tail = prev.and_then(|w| w.upgrade());
        }

        node.borrow_mut().prev = None;
        node.borrow_mut().next = None;
        self.len -= 1;
    }

    fn evict_lru(&mut self) {
        if let Some(tail) = self.tail.clone() {
            let key = tail.borrow().key;
            self.remove_node(tail);
            self.map.remove(&key);
        }
    }

    fn insert_at_head(&mut self, node: Rc<RefCell<Node>>) {
        node.borrow_mut().prev = None;
        node.borrow_mut().next = self.head.clone();

        if let Some(head) = &self.head {
            head.borrow_mut().prev = Some(Rc::downgrade(&node));
        }

        self.head = Some(node.clone());

        if self.tail.is_none() {
            self.tail = Some(node);
        }

        self.len += 1;
    }
}
