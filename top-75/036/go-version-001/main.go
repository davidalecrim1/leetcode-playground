package main

import "sort"

// Time Complexity: O(n) + O(n log n) = O(n log n)
func merge(intervals [][]int) [][]int {
	// Approach sort the array and compare the elements.

	// O(n log n)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// O(n)
	res := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		currStart := intervals[i][0]
		end := intervals[i][1]

		lastEnd := res[len(res)-1][1]
		if currStart <= lastEnd {
			res[len(res)-1][1] = max(lastEnd, end)
		} else {
			res = append(res, intervals[i])
		}
	}

	return res
}
