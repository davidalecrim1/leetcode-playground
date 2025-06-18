package main

import "math"

func lengthOfLIS(nums []int) int {
	// I have a decision tree to incluse or exclude each element.
	// That can be done as long as nums[i] < nums[i+1]

	// Strategy: DFS with Cache
	// Pure DFS = (2ˆN)
	// DFS with Cache = O(n^2) -- Not done right now.

	n := len(nums)

	var dfs func(lastVal int, i int, longestCount int) int
	dfs = func(lastVal int, i int, longestCount int) int {
		if i >= n {
			return longestCount
		}

		res := 0
		for j := i; j < n; j++ {
			if lastVal < nums[j] {
				newlongestCount := longestCount + 1
				res = max(res, dfs(nums[j], j+1, newlongestCount))
			}
		}

		return res
	}

	return dfs(math.MinInt64, 0, 0)
}
