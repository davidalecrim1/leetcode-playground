impl Solution {
    pub fn count_substrings(s: String) -> i32 {
        let chars: Vec<char> = s.chars().collect();
        let n = chars.len();

        let palindrome_count = |mut i: isize, mut j: isize| {
            let mut count = 0;

            while i >= 0 && (j as usize) < n && chars[i as usize] == chars[j as usize] {
                count+=1;
                i-=1;
                j+=1;
            }

            count
        };

        let mut res = 0;
        for i in 0..n{
            res += palindrome_count(i as isize, i as isize);
        }

        for i in 1..n{
            res += palindrome_count((i - 1) as isize, i as isize);
        }

        res
    }
}