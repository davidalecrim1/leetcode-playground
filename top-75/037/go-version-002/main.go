package main

import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
	// Strategy
	// Approach
	// Greedy one comparing the values on the end.
	// Count the ones that I remove.
	// Remove always the one that ends last.

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// fmt.Println(intervals)

	lastEnd := intervals[0][1]
	minRemoval := 0

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= lastEnd {
			// they are not overlapping
			lastEnd = intervals[i][1]
		} else {
			minRemoval++
			// always remove the one that ends last.
			lastEnd = min(lastEnd, intervals[i][1])
		}
	}

	return minRemoval
}
