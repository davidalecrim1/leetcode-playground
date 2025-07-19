package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// Use two pointers
	// N+1 diff between them
	if head == nil {
		return nil
	}

	dummy := &ListNode{Next: head}

	prev := dummy
	curr := head
	for range n {
		curr = curr.Next
	}

	for curr != nil {
		curr = curr.Next
		prev = prev.Next
	}

	fmt.Println(prev.Val)
	prev.Next = prev.Next.Next
	return dummy.Next
}
