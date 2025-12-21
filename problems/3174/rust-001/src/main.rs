// Use an array as a stack to pop elements
// Walk the array while it has numbers

impl Solution {
    pub fn clear_digits(s: String) -> String {
        let chars = s.chars();
        let mut stack = Vec::with_capacity(s.len());

        for char in chars {
            if char.is_numeric() {
                if stack.len() < 1 {
                    return stack.iter().collect();
                }

                let _ = stack.pop();
            } else {
                stack.push(char);
            }
        }

        return stack.iter().collect();
    }
}

struct Solution;

pub fn main() {
    let res = Solution::clear_digits("cb34".to_string());
    assert_eq!(res, "")
}
