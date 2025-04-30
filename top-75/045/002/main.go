package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time Complexity: O(n)
// Space Complexity: O(1)
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	var prev *ListNode
	curr := slow.Next
	slow.Next = nil // break the linked list

	// on the second half of the list, reverse it
	for curr != nil {
		nextTemp := curr.Next
		curr.Next = prev

		prev = curr
		curr = nextTemp
	}

	first, second := head, prev
	for second != nil {
		tmp1, tmp2 := first.Next, second.Next
		first.Next = second
		second.Next = tmp1

		first = tmp1
		second = tmp2
	}
}
