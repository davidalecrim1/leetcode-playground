// Use sorting to define if they are anagrams.
// Create a hashmap based on the sorted key to group them.
// Use that to build the result

use std::collections::HashMap;

struct Solution;

// Time: O(n)
impl Solution {
    pub fn group_anagrams(strs: Vec<String>) -> Vec<Vec<String>> {
        let mut map: HashMap<Vec<char>, Vec<String>> = HashMap::new();

        for s in strs {
            let mut key: Vec<char> = s.chars().collect();
            key.sort_unstable();

            map.entry(key).or_insert_with(Vec::new).push(s);
        }

        map.into_values().collect()
    }
}

pub fn main() {}
