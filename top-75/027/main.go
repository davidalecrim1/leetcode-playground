package main

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	// Approach 1: Breadth First Search with visited graphs.

	if node == nil {
		return nil
	}

	queue := []*Node{node}

	cloned := make(map[*Node]*Node)
	cloned[node] = &Node{Val: node.Val}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		for _, neighbor := range curr.Neighbors {
			if _, ok := cloned[neighbor]; !ok {
				cloned[neighbor] = &Node{Val: neighbor.Val}
				queue = append(queue, neighbor)
			}

			cloned[curr].Neighbors = append(cloned[curr].Neighbors, cloned[neighbor])
		}
	}

	return cloned[node]
}
