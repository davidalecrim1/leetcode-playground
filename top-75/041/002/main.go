package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	// Strategy:
	// 1. Use a hashmap to store the seen values (Space O(n))

	// 2. Two pointers (fast and slow)

	fast := head
	slow := head

	for fast != nil {
		if fast.Next != nil && fast.Next.Next != nil {
			fast = fast.Next.Next
			slow = slow.Next

			if fast == slow {
				return true
			}
		} else {
			return false
		}
	}

	return false
}
