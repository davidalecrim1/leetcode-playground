package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	dfs(root, 0, &res)
	return res
}

func dfs(curr *TreeNode, level int, res *[][]int) {
	if curr == nil {
		return
	}

	if len(*res) <= level {
		(*res) = append((*res), []int{})
	}

	(*res)[level] = append((*res)[level], curr.Val)

	next := level + 1
	dfs(curr.Left, next, res)
	dfs(curr.Right, next, res)
}
