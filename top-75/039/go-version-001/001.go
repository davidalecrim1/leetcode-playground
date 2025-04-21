/*
This uses a priority queue data structure to solve the problem.
The h[0] is the minimum heap.
The h[len(h-1) is the maximum heap.
*/
package main

import (
	"container/heap"
	"fmt"
	"sort"
)

/**
 * Definition of Interval:
 * type Interval struct {
 *     Start, End int
 * }
 */

type Interval struct {
	Start, End int
}

/**
 * @param intervals: an array of meeting time intervals
 * @return: the minimum number of conference rooms required
 */
func MinMeetingRooms(intervals []*Interval) int {
	// O(n log n)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})

	h := &MinHeap{}

	// The Space complexity will be O(n).
	heap.Init(h)

	// O(log n) for push and pop given it uses a binary tree as data structure.
	heap.Push(h, intervals[0].End)

	for i := 1; i < len(intervals); i++ {
		// if the room is free, reuse it.
		if intervals[i].Start >= (*h)[0] {
			heap.Pop(h)
		}

		// allocate a new room or reuse the one above if the condition is met
		heap.Push(h, intervals[i].End)
	}

	return h.Len()
}

type MinHeap []int

func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h MinHeap) Len() int {
	return len(h)
}

func (h *MinHeap) Push(x any) {
	val, ok := x.(int)
	if !ok {
		panic("unexpected value received")
	}

	*h = append(*h, val)
}

func (h *MinHeap) Pop() any {
	n := len(*h) - 1
	x := (*h)[n]
	*h = (*h)[:n]
	return x
}

func main() {
	res1 := MinMeetingRooms([]*Interval{
		{Start: 0, End: 30},
		{Start: 5, End: 10},
		{Start: 15, End: 20},
	})

	fmt.Printf("Expected 2 got %v\n", res1)

	res2 := MinMeetingRooms([]*Interval{
		{Start: 65, End: 424},
		{Start: 351, End: 507},
		{Start: 314, End: 807},
		{Start: 387, End: 722},
		{Start: 19, End: 797},
		{Start: 259, End: 722},
		{Start: 165, End: 221},
		{Start: 136, End: 897},
	})

	fmt.Printf("Expected 7 got %v\n", res2)
}
