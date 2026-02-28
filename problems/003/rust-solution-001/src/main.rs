fn main() {
    println!("Hello, world!");
}

struct Solution;
// sliding window
// walk the input while the window is valid increasing the right pointer
// store the seen chars
// when the window becomes invalid, update to the seen char
// keep doing that until the end of the array, keep storing the max valid window on the way

impl Solution {
    pub fn length_of_longest_substring(s: String) -> i32 {
        let mut l = 0;
        let mut seen = std::collections::HashMap::new();
        let mut res = 0;
        let input: Vec<char> = s.chars().collect();

        for r in 0..input.len() {
            // removes if exists
            if let Some(value) = seen.remove(&input[r]) {
                // ensure we only look at the eixsting window
                l = l.max(value + 1);
            }

            seen.insert(input[r], r);
            res = res.max((r - l) + 1);
        }

        res as i32
    }
}
