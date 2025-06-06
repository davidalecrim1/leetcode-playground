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

func levelOrder(root *TreeNode) [][]int {
	var res [][]int

	dfs := func(*TreeNode, int) {}
	dfs = func(curr *TreeNode, level int) {
		if curr == nil {
			return
		}

		for len(res) <= level {
			res = append(res, []int{})
		}

		res[level] = append(res[level], curr.Val)

		nextLevel := level + 1

		dfs(curr.Left, nextLevel)
		dfs(curr.Right, nextLevel)
	}

	dfs(root, 0)

	return res
}
