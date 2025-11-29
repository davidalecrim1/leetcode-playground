package main

// Apply a BFS or DFS with level tracking.
// order level traversal left, save, right, save, next level

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
	if root == nil {
		return nil
	}

	type NodeWrapper struct {
		Node  *TreeNode
		Level int
	}

	queue := []*NodeWrapper{&NodeWrapper{
		Node:  root,
		Level: 0,
	}}

	res := [][]int{}

	for len(queue) != 0 {
		w := queue[0]
		queue = queue[1:]

		if len(res) == w.Level {
			res = append(res, []int{})
		}

		res[w.Level] = append(res[w.Level], w.Node.Val)
		next := w.Level + 1

		if w.Node.Left != nil {
			queue = append(queue, &NodeWrapper{
				Node:  w.Node.Left,
				Level: next,
			})
		}
		if w.Node.Right != nil {
			queue = append(queue, &NodeWrapper{
				Node:  w.Node.Right,
				Level: next,
			})
		}
	}

	return res
}
