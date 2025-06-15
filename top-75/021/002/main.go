package main

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var res [][]int

	var dfs func(startIndex, currentSum int, path []int)
	dfs = func(startIndex, currentSum int, path []int) {
		for i := startIndex; i < len(candidates); i++ {
			num := candidates[i]
			newSum := currentSum + num

			if newSum > target {
				break // Stop exploring further as candidates are sorted
			}

			newPath := append(path, num)
			if newSum == target {
				res = append(res, append([]int{}, newPath...))
				continue
			}

			dfs(i, newSum, newPath)
		}
	}

	dfs(0, 0, []int{})
	return res
}
