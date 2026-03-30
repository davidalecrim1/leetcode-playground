fn main() {
    println!("Hello, world!");
}
use std::collections::{HashMap, hash_map::Entry};

struct Solution;

struct Entries {
    value: String,
    timestamp: i32,
}
struct TimeMap {
    map: HashMap<String, Vec<Entries>>,
}

/**
 * `&self` means the method takes an immutable reference.
 * If you need a mutable reference, change it to `&mut self` instead.
 */
impl TimeMap {
    fn new() -> Self {
        Self {
            map: HashMap::new(),
        }
    }

    fn set(&mut self, key: String, value: String, timestamp: i32) {
        self.map
            .entry(key)
            .or_default()
            .push(Entries { value, timestamp });
    }

    fn get(&self, key: String, timestamp: i32) -> String {
        if let Some(entries) = self.map.get(&key) {
            match entries.binary_search_by_key(&timestamp, |e| e.timestamp) {
                Ok(i) => entries[i].value.clone(),
                Err(i) if i > 0 => entries[i - 1].value.clone(),
                _ => String::new(),
            }
        } else {
            String::new()
        }
    }
}
