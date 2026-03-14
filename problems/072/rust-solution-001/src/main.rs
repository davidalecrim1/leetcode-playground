fn main() {
    println!("Hello, world!");
}

struct Solution;
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

impl Solution {
    pub fn merge_two_lists(
        list1: Option<Box<ListNode>>,
        list2: Option<Box<ListNode>>,
    ) -> Option<Box<ListNode>> {
        let mut curr1 = list1;
        let mut curr2 = list2;
        let mut dummy = Box::new(ListNode::new(0));
        let mut tail = &mut dummy;

        while curr1.is_some() && curr2.is_some() {
            let take_curr_1 = curr1.as_ref().unwrap().val < curr2.as_ref().unwrap().val;

            if take_curr_1 {
                let mut node = curr1.take().unwrap();
                curr1 = node.next.take();
                tail.next = Some(node);
            } else {
                let mut node = curr2.take().unwrap();
                curr2 = node.next.take();
                tail.next = Some(node);
            }

            tail = tail.next.as_mut().unwrap();
        }

        tail.next = if curr1.is_some() { curr1 } else { curr2 };
        dummy.next
    }
}
