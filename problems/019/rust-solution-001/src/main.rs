fn main() {
    // Test: [1,2,3,4,5], n=2 -> [1,2,3,5]
    let head = build_list(&[1, 2, 3, 4, 5]);
    let result = Solution::remove_nth_from_end(head, 2);
    println!("{:?}", collect_list(result)); // [1, 2, 3, 5]

    // Test: [1], n=1 -> []
    let head = build_list(&[1]);
    let result = Solution::remove_nth_from_end(head, 1);
    println!("{:?}", collect_list(result)); // []

    // Test: [1,2], n=1 -> [1]
    let head = build_list(&[1, 2]);
    let result = Solution::remove_nth_from_end(head, 1);
    println!("{:?}", collect_list(result)); // [1]

    // Test: [1,2], n=2 -> [2] (remove head)
    let head = build_list(&[1, 2]);
    let result = Solution::remove_nth_from_end(head, 2);
    println!("{:?}", collect_list(result)); // [2]
}

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

struct Solution;

// Walk n times on the list with pointer right.
// Then start a new pointer called left and move them together until right is nil.
// left should end up in (3) prev to the actual one to remove it, then remove it.
impl Solution {
    pub fn remove_nth_from_end(head: Option<Box<ListNode>>, n: i32) -> Option<Box<ListNode>> {
        // Compute list length with a shared reference.
        let mut len = 0;
        let mut cur = &head;
        while let Some(node) = cur {
            len += 1;
            cur = &node.next;
        }

        // Index of the node to remove (0-based from the front).
        let target = len - n as usize;

        // Sentinel node simplifies the edge case of removing the head.
        let mut dummy = Box::new(ListNode { val: 0, next: head });
        let mut cur = &mut *dummy;

        for _ in 0..target {
            cur = cur.next.as_deref_mut().unwrap();
        }

        // cur is now the node just before the target; splice it out.
        let removed = cur.next.take();
        if let Some(mut removed_node) = removed {
            cur.next = removed_node.next.take();
        }

        dummy.next
    }
}

fn build_list(vals: &[i32]) -> Option<Box<ListNode>> {
    let mut head = None;
    for &v in vals.iter().rev() {
        let mut node = Box::new(ListNode::new(v));
        node.next = head;
        head = Some(node);
    }
    head
}

fn collect_list(head: Option<Box<ListNode>>) -> Vec<i32> {
    let mut result = vec![];
    let mut cur = head;
    while let Some(node) = cur {
        result.push(node.val);
        cur = node.next;
    }
    result
}
