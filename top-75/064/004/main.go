package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
Strategy
- DFS
- BFS
- Level tracking to save the nodes
- O(n)
*/
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	res := [][]int{}

	var dfs func(node *TreeNode, level int)
	dfs = func(node *TreeNode, level int) {
		if node == nil {
			return
		}

		if level >= len(res) {
			res = append(res, []int{})
		}

		res[level] = append(res[level], node.Val)
		next := level + 1

		dfs(node.Left, next)
		dfs(node.Right, next)
	}

	dfs(root, 0)

	return res
}
