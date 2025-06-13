package main

import (
	"fmt"
)

// Time Complexity: O(n * m)
// Space Complexity: O(n)
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)

	for i := range amount + 1 {
		dp[i] = amount + 1
	}

	dp[0] = 0

	for a := range amount + 1 {
		for c := range coins {
			if a-coins[c] >= 0 {
				dp[a] = min(dp[a], 1+dp[a-coins[c]])
			}
		}
	}

	if dp[amount] == amount+1 {
		return -1
	}

	return dp[amount]
}

func main() {
	fmt.Printf("coinChange([]int{2}, 3): %v\n", coinChange([]int{2}, 3))
}
