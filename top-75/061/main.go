package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Depth For Search
// Time Complexity: O(n)
// Space Complexity: O(n)
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil {
		return false
	}

	if q == nil {
		return false
	}

	if p.Val == q.Val {
		return true && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	}

	return false
}
