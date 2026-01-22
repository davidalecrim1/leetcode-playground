// Use a set to define the non repeated num values
// Do not sort because it's O(n log n)
// Use the set to check if n-1 exists (this tells us it's the beginning of a sequence)
// Walk until the end, store the longest

impl Solution {
    pub fn longest_consecutive(nums: Vec<i32>) -> i32 {
        // will call the from_iter on the HashSet as the compiler can know the type
        let set: std::collections::HashSet<i32> = nums.into_iter().collect();
        let mut longest: i32 = 0;

        for n in &set {
            if set.contains(&(n - 1 as i32)) {
                continue;
            }

            let mut curr: i32 = *n;
            let mut sequence: i32 = 0;
            loop {
                if set.contains(&curr) {
                    sequence += 1;
                    curr += 1;

                    longest = std::cmp::max(longest, sequence)
                } else {
                    break;
                }
            }
        }

        longest
    }
}

struct Solution;

fn main() {
    println!("Hello, world!");
}
