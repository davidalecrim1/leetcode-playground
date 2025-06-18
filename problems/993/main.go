package main

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

func isCousins(root *TreeNode, x int, y int) bool {
	// Strategy
	// Walk the tree with DFS
	// Check the parent
	// Check the level

	xLevel := 0
	yLevel := 0
	var xParent *TreeNode
	var yParent *TreeNode

	var dfs func(level int, curr *TreeNode, prev *TreeNode)
	dfs = func(level int, curr *TreeNode, prev *TreeNode) {
		if curr == nil {
			return
		}

		if curr.Val == x {
			xParent = prev
			xLevel = level
		}

		if curr.Val == y {
			yParent = prev
			yLevel = level
		}

		dfs(level+1, curr.Left, curr)
		dfs(level+1, curr.Right, curr)
	}

	dfs(0, root, nil)

	// fmt.Println("x level", xLevel)
	// fmt.Println("x parent val", xParent.Val)

	// fmt.Println("y level", yLevel)
	// fmt.Println("y parent val", yParent.Val)

	if (xLevel == yLevel) && xParent != nil && yParent != nil && (xParent.Val != yParent.Val) {
		return true
	}

	return false
}
