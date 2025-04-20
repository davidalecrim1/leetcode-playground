package main

import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
	// Approach:
	// I thought about applying DFS, but this problem is more of a greedy approach.
	// Sorting is essencial to have a proper logic.

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := 0
	prev := 0

	for i := 1; i < len(intervals); i++ {
		// curr start to the prev end
		if intervals[i][0] >= intervals[prev][1] {
			// there is not overlapping, just keep moving.
			prev = i
		} else {
			res++

			// in my greedy approach, always delete the largest because of its wide impact.
			if intervals[i][1] <= intervals[prev][1] {
				// "delete" the prev, and keep the curr (i).
				prev = i
			}

			// if not, keep the lowest end that is the prev.
			// Meaning to delete the curr, and keep the prev.
		}
	}

	return res
}
