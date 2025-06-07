package main

import (
	"container/heap"
)

// Time Complexity: O(n log n)
// Space Complexity: O(n)
func topKFrequent(nums []int, k int) []int {
	freqMap := make(map[int]int)
	for _, num := range nums {
		freqMap[num]++
	}

	h := &MaxHeap{}
	heap.Init(h)
	for num, freq := range freqMap {
		heap.Push(h, Element{num, freq})
	}

	res := make([]int, 0, k)
	for i := 0; i < k && h.Len() > 0; i++ {
		res = append(res, heap.Pop(h).(Element).number)
	}

	return res
}

type MaxHeap []Element

func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(Element))
}

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h MaxHeap) Less(i, j int) bool {
	return h[i].count > h[j].count // inverted to transform min heap in max heap
}

func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	elem := old[n-1]
	*h = old[:n-1]
	return elem
}

type Element struct {
	number int
	count  int
}
