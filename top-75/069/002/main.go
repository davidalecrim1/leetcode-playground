package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// Time Complexity: O(n)
// Space Complexity: O(1)
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	res := 0
	count := 0

	var helper func(*TreeNode)
	helper = func(curr *TreeNode) {
		if curr == nil || count >= k {
			return
		}

		helper(curr.Left)
		count++

		if count == k {
			res = curr.Val
			return
		}

		helper(curr.Right)
	}

	helper(root)
	return res
}

func main() {
	input := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 4,
				},
			},
		},
		Right: &TreeNode{
			Val: 6,
		},
	}

	got := kthSmallest(input, 1)
	fmt.Printf("expected %v got %v", 1, got)
}
