// Definition for singly-linked list.
// #[derive(PartialEq, Eq, Clone, Debug)]
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

struct Solution;
impl Solution {
    pub fn remove_nth_from_end(head: Option<Box<ListNode>>, n: i32) -> Option<Box<ListNode>> {
        let mut curr = &head;
        let mut len = 0;

        while let Some(node) = curr {
            len += 1;
            curr = &node.next;
        }

        let target = len - n as usize;
        let mut dummy = Box::new(ListNode { val: 0, next: head });
        let mut curr = &mut *dummy;

        for _ in 0..target {
            curr = curr.next.as_deref_mut().unwrap();
        }

        let removed = curr.next.take();
        if let Some(mut node) = removed {
            curr.next = node.next.take();
        }

        dummy.next
    }
}

fn main() {}
