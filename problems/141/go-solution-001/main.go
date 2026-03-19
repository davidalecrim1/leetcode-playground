package main

// Strategy:
// Use a slow and fast pointer
// If they end up on the same node, return true

/**
 * Definition for singly-linked list.
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	slow := head
	fast := head.Next

	for slow != nil && fast != nil {
		if slow == fast {
			return true
		}

		slow = slow.Next

		if fast.Next != nil {
			fast = fast.Next.Next
		} else {
			fast = fast.Next
		}

	}

	return false
}
