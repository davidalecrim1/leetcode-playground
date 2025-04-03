package main

// Approach: Actual Breadth First Search (BFS) approach.
// Level Order Traversal is BFS.

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	res := [][]int{}
	queue := []*Node{root} // FIFO

	for len(queue) > 0 {
		levelSize := len(queue)
		level := []int{}

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			level = append(level, node.Val)
			queue = append(queue, node.Children...)
		}

		res = append(res, level)
	}

	return res
}
