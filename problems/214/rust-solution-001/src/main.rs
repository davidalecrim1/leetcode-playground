pub fn main() {
    let _ = Solution::is_anagram("david", "david");
}

struct Solution;

impl Solution {
    pub fn is_anagram(s: String, t: String) -> bool {
        let mut s1: Vec<char> = s.chars().collect();
        s1.sort_unstable();

        let mut t1: Vec<char> = t.chars().collect();
        t1.sort_unstable();

        s1 == t1
    }
}
