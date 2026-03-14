fn main() {
    println!("Hello, world!");
}

struct Solution;
// Create a vec<i32> keeping track of s1 chars and total of them
// Create a sliding window of the len(s1)
// Use a shared vec<i32> and increase and decrease the count as the window moves
// If there is a match, return true, else false

impl Solution {
    pub fn check_inclusion(s1: String, s2: String) -> bool {
        if s1.len() > s2.len() {
            return false;
        }

        let sb1 = s1.as_bytes();
        let mut sv1 = vec![0_i32; 26];
        for i in 0..s1.len() {
            let idx = (sb1[i] - 'a' as u8) as usize;
            sv1[idx] += 1;
        }

        let mut left = 0;
        let mut right = s1.len() - 1;

        let sb2 = s2.as_bytes();

        let mut sv2 = vec![0_i32; 26];
        for i in left..right {
            let idx = (sb2[i] - 'a' as u8) as usize;
            sv2[idx] += 1;
        }

        while right < s2.len() {
            // update the new right
            let mut idx = (sb2[right] - 'a' as u8) as usize;
            sv2[idx] += 1;

            if sv1 == sv2 {
                return true;
            }

            // remove the old left
            idx = (sb2[left] - 'a' as u8) as usize;
            sv2[idx] -= 1;

            left += 1;
            right += 1;
        }

        return false;
    }
}
