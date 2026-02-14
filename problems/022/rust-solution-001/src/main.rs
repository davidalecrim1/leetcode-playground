fn main() {}

struct Solution;

impl Solution {
    pub fn generate_parenthesis(n: i32) -> Vec<String> {
        let mut generated: Vec<String> = Vec::new();
        Self::backtrack("".to_string(), 0, 0, n, &mut generated);
        generated
    }

    fn backtrack(curr: String, open: i32, close: i32, n: i32, generated: &mut Vec<String>) {
        if curr.len() == (n * 2) as usize {
            generated.push(curr);
            return;
        }

        if open < n {
            let mut next = curr.clone();
            next.push('(');
            Self::backtrack(next, open + 1, close, n, generated);
        }

        if close < open {
            let mut next = curr.clone();
            next.push(')');
            Self::backtrack(next, open, close + 1, n, generated);
        }
    }
}
