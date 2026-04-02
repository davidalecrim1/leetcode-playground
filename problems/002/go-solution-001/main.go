package main

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	c1 := l1
	c2 := l2
	dummy := &ListNode{
		Val:  0,
		Next: nil,
	}
	prev := dummy

	carry := 0
	for c1 != nil || c2 != nil {
		sum := 0
		if c1 != nil {
			sum += c1.Val
			c1 = c1.Next
		}

		if c2 != nil {
			sum += c2.Val
			c2 = c2.Next
		}

		if carry >= 1 {
			sum += carry
		}

		out := &ListNode{}
		prev.Next = out

		out.Val = sum % 10
		if sum >= 10 {
			carry = 1
		} else {
			carry = 0
		}

		prev = prev.Next
	}

	if carry >= 1 {
		prev.Next = &ListNode{
			Val:  carry,
			Next: nil,
		}
	}

	return dummy.Next

}
