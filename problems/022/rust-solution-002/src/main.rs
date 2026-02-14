// This will create a decision tree.
// Walk every path. Add rules to the formation.
// e.g. cannot start with ), so if close < open
// open < n
// backtrack this
pub fn main() {}

struct Solution;

impl Solution {
    pub fn generate_parenthesis(n: i32) -> Vec<String> {
        let mut output = Vec::new();
        Self::backtrack(&mut output, n, &mut String::new(), 0, 0);
        return output;
    }

    fn backtrack(output: &mut Vec<String>, n: i32, curr: &mut String, open: i32, close: i32) {
        if curr.len() == (n * 2) as usize {
            output.push(curr.clone());
            return;
        }

        if open < n {
            curr.push('(');
            Self::backtrack(output, n, &mut curr.clone(), open + 1, close);
            curr.pop();
        }

        if close < open {
            curr.push(')');
            Self::backtrack(output, n, &mut curr.clone(), open, close + 1);
            curr.pop();
        }
    }
}
