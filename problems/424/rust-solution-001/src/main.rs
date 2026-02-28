fn main() {
    println!("Hello, world!");
}

struct Solution;

// Use a sliding window
// Start expanding the right pointer
// Keep track of all chars I see.
// Check the hashmap to fetch the char that is longest.
// When the window becomes invalid, move the left pointer until it's valid
// Keep track of the max valid window
impl Solution {
    pub fn character_replacement(s: String, k: i32) -> i32 {
        let mut res = 0;
        let mut seen = [0; 26];
        let mut l = 0;
        let bytes = s.as_bytes();

        for r in 0..(s.len()) {
            // increase the count on the hashmap
            let idx = (bytes[r] - b'A') as usize;
            seen[idx] += 1;

            // fetch the one that is longest on the hashmap
            let mut longest = 0;
            for &count in &seen {
                longest = longest.max(count);
            }

            // if the window is invalid
            while (longest + (k) as usize) < (r - l) + 1 {
                let idx = (bytes[l] - b'A') as usize;
                seen[idx] -= 1;
                l += 1;
            }

            // update the res
            res = std::cmp::max(res, (r - l) + 1);
        }

        res as i32
    }
}
