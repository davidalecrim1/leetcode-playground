package main

import (
	"container/heap"
	"fmt"
	"math"
)

type Edge struct {
	target int
	weight int
}

type Graph map[int][]Edge

type PriorityQueue []Node

/*
type Interface interface {
    sort.Interface
    Push(x any) // add x as element Len()
    Pop() any   // remove and return element Len() - 1.
}

The heap package provides an efficient way to keep the smallest (or largest) element at the top,
which is essential for Dijkstra's algorithm.
*/

type Node struct {
	id       int
	distance int
}

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool { return pq[i].distance < pq[j].distance }

func (pq PriorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

// x needs to be any because of heap.Interface
func (pq *PriorityQueue) Push(x any) {
	node, ok := x.(Node)
	if !ok {
		panic("expected x to be of type Node")
	}

	*pq = append(*pq, node)
}

func (pq *PriorityQueue) Pop() any {
	n := len(*pq)
	node := (*pq)[n-1]
	*pq = (*pq)[:n-1] // Remove the last element
	return node
}

func buildGraph(times [][]int) Graph {
	graph := make(Graph)
	for _, edge := range times {
		u, v, w := edge[0], edge[1], edge[2]
		graph[u] = append(graph[u], Edge{target: v, weight: w})
	}
	return graph
}

func networkDelayTime(times [][]int, n int, k int) int {
	graph := buildGraph(times)

	// Distance map to store the shortest path to each node
	distances := make(map[int]int)
	for i := 1; i <= n; i++ {
		distances[i] = math.MaxInt32
	}
	distances[k] = 0

	pq := &PriorityQueue{
		{id: k, distance: 0},
	}
	heap.Init(pq)

	for pq.Len() > 0 {
		curr := heap.Pop(pq).(Node)
		// Skip if a shorter path to current.id has been found
		if curr.distance > distances[curr.id] {
			continue
		}

		for _, edge := range graph[curr.id] {
			newDist := distances[curr.id] + edge.weight
			if newDist < distances[edge.target] {
				distances[edge.target] = newDist
				heap.Push(pq, Node{id: edge.target, distance: newDist})
			}
		}
	}

	maxTime := 0
	for _, d := range distances {
		if d == math.MaxInt32 {
			return -1
		}

		if d > maxTime {
			maxTime = d
		}
	}

	return maxTime
}

func main() {
	times := [][]int{
		{2, 1, 1},
		{2, 3, 1},
		{3, 4, 1},
	}
	n, k := 4, 2

	fmt.Printf("networkDelayTime(times, n, k): %v\n", networkDelayTime(times, n, k))
}
