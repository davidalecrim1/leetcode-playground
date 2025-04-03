package main

// Approach: Some kind of Breadth First Search (BFS) with recursion.
// Using a level control.
// Time Complexity: O(N)

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	levels := [][]int{}
	traverse(root, 0, &levels)
	return levels
}

func traverse(node *Node, depth int, levels *[][]int) {
	if node == nil {
		return
	}

	if len(*levels) == depth {
		// increase the size of levels
		*levels = append(*levels, []int{})
	}

	(*levels)[depth] = append((*levels)[depth], node.Val)

	for _, child := range node.Children {
		traverse(child, depth+1, levels)
	}
}
