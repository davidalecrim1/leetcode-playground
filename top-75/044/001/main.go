package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	left := dummy
	right := dummy

	for i := 0; i <= n && right != nil; i++ {
		right = right.Next
	}

	for right != nil {
		right = right.Next
		left = left.Next
	}

	left.Next = left.Next.Next
	return dummy.Next
}
