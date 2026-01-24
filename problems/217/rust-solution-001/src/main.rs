// To avoid N*N on the array we need a set with good lookup
// HashMap will help with that.

use std::collections::HashMap;

impl Solution {
    pub fn contains_duplicate(nums: Vec<i32>) -> bool {
        let mut hm: HashMap<i32, bool> = HashMap::new();

        for i in nums {
            if hm.contains_key(&i) {
                return true;
            }

            hm.insert(i, true);
        }

        return false;
    }
}

// To avoid N*N on the array we need a set with good lookup
// HashMap will help with that.

use std::collections::HashSet;

impl Solution {
    pub fn contains_duplicate(nums: Vec<i32>) -> bool {
        let mut hm: HashSet<i32> = HashSet::new();

        for i in nums {
            if !hm.insert(i) {
                return true;
            }
        }

        false
    }
}
