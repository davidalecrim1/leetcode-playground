fn main() {}

use std::collections::HashMap;
struct Solution;

// Use a stack to append when isOpen
// Pop from the stack when is close,
// if don't match, return false.
// When end up the stack, return true.
impl Solution {
    pub fn is_valid(s: String) -> bool {
        let mut map: HashMap<char, char> = HashMap::new();
        map.insert('}', '{');
        map.insert(']', '[');
        map.insert(')', '(');

        let mut stack = Vec::with_capacity(s.len());
        for c in s.chars() {
            if let Some(open) = map.get(&c) {
                if stack.len() == 0 {
                    return false;
                }

                if let Some(value) = stack.pop() {
                    if value != *open {
                        return false;
                    }
                } else {
                    return false;
                }
            } else {
                stack.push(c);
            }
        }

        return stack.len() == 0;
    }
}
