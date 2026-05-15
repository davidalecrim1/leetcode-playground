struct Solution;

impl Solution {
    pub fn is_palindrome(s: String) -> bool {
        let filtered: Vec<u8> = s
            .bytes()
            .filter(|b| b.is_ascii_alphanumeric())
            .map(|b| b.to_ascii_lowercase())
            .collect();

        filtered.iter().eq(filtered.iter().rev())
    }
}

fn main() {
    println!("Hello, world!");
}
