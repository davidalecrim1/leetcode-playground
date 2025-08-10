package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	// Strategy
	// DFS
	// BFS

	if root == nil {
		return [][]int{}
	}

	type WrapNode struct {
		*TreeNode
		Level int
	}

	queue := []*WrapNode{
		{
			TreeNode: root,
			Level:    0,
		},
	}
	res := [][]int{}

	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]

		if node.Level == len(res) {
			res = append(res, []int{})
		}

		res[node.Level] = append(res[node.Level], node.Val)
		next := node.Level + 1

		if node.Left != nil {
			queue = append(queue, &WrapNode{node.Left, next})
		}

		if node.Right != nil {
			queue = append(queue, &WrapNode{node.Right, next})
		}
	}

	return res
}
