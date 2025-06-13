package main

import "math"

func coinChange(coins []int, amount int) int {
	// Strategy 1:
	// Greedy Strategy
	// Sort the coins
	// Walk in reverse adding the values
	// But this will have unhandled edge cases e.g. remove a previous added coin.
	// Time: O(n)

	// Strategy 2:
	// I fell here I could apply dynamic programming
	// Let me think this.
	// This will end up being a decision tree 2^(len(coins)).
	// sub problem: 1,5,5 = 5,5,1 => 3
	// sub problem: 5,1 = 1,5 => 2

	// Strategy 3:
	// Brute force this with two loops
	// Time: O(n^2)

	// Choosen: DFS + Memo
	// Time: O(n)
	// Space: O(n)
	memo := make(map[int]int)

	var dfs func(remaining int) int
	dfs = func(remaining int) int {
		if remaining == 0 {
			return 0
		}

		if remaining < 0 {
			return -1
		}

		if val, ok := memo[remaining]; ok {
			return val
		}

		min := math.MaxInt64
		for _, coin := range coins {
			res := dfs(remaining - coin)
			if res >= 0 && res < min {
				min = res + 1
			}

		}
		if min == math.MaxInt64 {
			memo[remaining] = -1
		} else {
			memo[remaining] = min
		}

		return memo[remaining]
	}

	return dfs(amount)
}
