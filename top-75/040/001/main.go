package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode // prev is nil
	curr := head

	for curr != nil {
		// save next
		next := curr.Next

		// swap arrow
		curr.Next = prev

		// move one
		prev = curr
		curr = next
	}

	return prev
}
