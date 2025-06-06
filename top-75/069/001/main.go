package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// Time Complexity: O(n)
// Space Complexity: O(N)
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	queue := []*TreeNode{}

	var helper func(*TreeNode)
	helper = func(curr *TreeNode) {
		if curr == nil {
			return
		}

		helper(curr.Left)
		queue = append(queue, curr)
		helper(curr.Right)
	}

	helper(root)
	return queue[k-1].Val
}
