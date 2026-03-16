fn main() {
    println!("Hello, world!");
}

struct Solution;
// Definition for singly-linked list.
#[derive(PartialEq, Eq, Clone, Debug)]
pub struct ListNode {
    pub val: i32,
    pub next: Option<Box<ListNode>>,
}

impl ListNode {
    #[inline]
    fn new(val: i32) -> Self {
        ListNode { next: None, val }
    }
}

// Make the linked list a vector with one pass
// Walk with two pointers (beginning and end)
// Updates the links until left == right
impl Solution {
    pub fn reorder_list(head: &mut Option<Box<ListNode>>) {
        let mut list = Vec::new();

        let mut curr = head.take();
        while let Some(mut node) = curr {
            curr = node.next.take();
            list.push(node);
        }

        let mut left = 0;
        let mut right = list.len().saturating_sub(1);

        let mut dummy = ListNode::new(0);
        let mut tail = &mut dummy;

        while left <= right {
            if left == right {
                tail.next = Some(list[left].clone());
                break;
            }

            tail.next = Some(list[left].clone());
            tail = tail.next.as_mut().unwrap();

            tail.next = Some(list[right].clone());
            tail = tail.next.as_mut().unwrap();

            left += 1;
            right -= 1;
        }

        *head = dummy.next;
    }
}
