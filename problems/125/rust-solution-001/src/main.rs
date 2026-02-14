fn main() {
    println!("Hello, world!");
}

struct Solution;

impl Solution {
    pub fn is_palindrome(s: String) -> bool {
        let chars: Vec<char> = s.chars().collect();

        let mut left = 0;
        let mut right = s.len() - 1;

        while left < right {
            if !chars[left].is_alphanumeric() {
                left += 1;
                continue;
            }

            if !chars[right].is_alphanumeric() {
                right -= 1;
                continue;
            }

            if chars[left].to_ascii_lowercase() == chars[right].to_ascii_lowercase() {
                left += 1;
                right -= 1;
            } else {
                return false;
            }
        }

        return true;
    }
}
