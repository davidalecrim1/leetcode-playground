package main

import "math"

// Do this recursive with a DFS approach
// Ensure left is less and right is more
// Create a window of validation for the edge cases

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return validate(math.MinInt64, math.MaxInt64, root)
}

func validate(low, high int, curr *TreeNode) bool {
	if curr == nil {
		return true
	}

	if curr.Val <= low || curr.Val >= high {
		return false
	}

	return validate(low, curr.Val, curr.Left) && validate(curr.Val, high, curr.Right)
}
