package main

import "math"

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
	return innerIsValidBST(root, math.MinInt64, math.MaxInt64)
}

func innerIsValidBST(head *TreeNode, min, max int) bool {
	if head == nil {
		return true
	}

	if head.Val <= min || head.Val >= max {
		return false
	}

	return innerIsValidBST(head.Left, min, head.Val) && innerIsValidBST(head.Right, head.Val, max)
}
