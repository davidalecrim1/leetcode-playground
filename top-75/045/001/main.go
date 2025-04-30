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

// Time Complexity: O(n)
// Space Complexity: O(n)
func reorderList(head *ListNode) {
	n := 0

	curr := head
	for curr != nil {
		curr = curr.Next
		n++
	}

	var first []*ListNode // pop from the beginning
	curr = head.Next
	for range n / 2 {
		first = append(first, curr)
		curr = curr.Next
	}

	var second []*ListNode // pop from the end
	for curr != nil {
		second = append(second, curr)
		curr = curr.Next
	}

	curr = head
	i := len(first)
	j := len(second)

	for i > 0 || j > 0 {
		if j > 0 {
			node := second[j-1]
			second = second[:j-1]
			curr.Next = node
			curr = curr.Next
			j--
		}

		if i > 0 {
			node := first[0]
			first = first[1:]
			curr.Next = node
			curr = curr.Next
			i--
		}
	}

	curr.Next = nil
}
