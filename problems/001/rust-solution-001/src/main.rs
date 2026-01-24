fn main() {
    println!("Hello, world!");
}
// Use a hashmap to store the num and idx
// If found, return both of them

// HashMap has a O(1) lookup

use std::collections::HashMap;

struct Solution;

impl Solution {
    pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
        let mut hashmap: HashMap<i32, i32> = HashMap::new();

        for (i, num) in nums.iter().enumerate() {
            let remainder = target - num;
            if let Some(other) = hashmap.get(&remainder) {
                return vec![i as i32, *other];
            }

            hashmap.insert(*num, i as i32);
        }

        return Vec::new();
    }
}
