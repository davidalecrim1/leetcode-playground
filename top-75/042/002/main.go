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

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// Strategy
	// Compare the values
	// Only walk the less one between both
	// Time O(n)

	dummy := &ListNode{Next: list1}
	head := dummy
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			dummy.Next = list1
			list1 = list1.Next
		} else {
			dummy.Next = list2
			list2 = list2.Next
		}
		dummy = dummy.Next
	}

	if list1 != nil {
		dummy.Next = list1
	}

	if list2 != nil {
		dummy.Next = list2
	}

	return head.Next
}
