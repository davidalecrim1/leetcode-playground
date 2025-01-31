package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	memo := make(map[*ListNode]int)

	for head != nil {
		_, ok := memo[head]

		if ok {
			return true
		}

		memo[head] = head.Val
		head = head.Next
	}

	return false
}
