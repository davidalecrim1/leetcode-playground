package main

type Node struct {
	Val       int
	Neighbors []*Node
}

// Time: O(n)
// Space: O(n)
func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	queue := make([]*Node, 0)
	clones := map[*Node]*Node{}

	queue = append(queue, node)
	clones[node] = &Node{Val: node.Val}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		for _, n := range curr.Neighbors {
			if _, ok := clones[n]; !ok {
				clones[n] = &Node{Val: n.Val}
				queue = append(queue, n)
			}

			clones[curr].Neighbors = append(clones[curr].Neighbors, clones[n])
		}
	}

	return clones[node]
}
