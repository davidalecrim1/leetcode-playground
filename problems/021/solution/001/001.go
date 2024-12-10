/*
Considering that the lists are sorted already.

Time complexity: O(n+m), where
n is the number of elements in the first list (l1) and
m is the number of elements in the second list (l2).

*/

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	resultList := &ListNode{}

	if list1.Val < list2.Val {
		resultList.Val = list1.Val
		resultList.Next = mergeTwoLists(list1.Next, list2)
	} else {
		resultList.Val = list2.Val
		resultList.Next = mergeTwoLists(list1, list2.Next)
	}

	return resultList
}
