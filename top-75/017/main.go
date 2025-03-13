package main

import (
	"fmt"
	"math"
)

func coinChange(coins []int, amount int) int {
	// Use dynamic programming to store the calculated value
	dp := make([]int, (amount + 1)) // will start at zero until it reaches the amount

	// We are going to use a bottom up strategy

	for i := range dp {
		dp[i] = math.MaxInt64 // just a dumb value to fill the array
	}

	// To reach the amount of 0 takes 0 steps
	dp[0] = 0

	// We brute force the steps from 0 to the amount
	for a := 1; a <= amount; a++ {
		for _, coin := range coins {
			// if the result of amount minus the coin is less than zero, we should keep looking

			if coin <= a {
				// The 1 is because we are counting the coin
				// Consider itself because the current amount may have a larger solution, like using 11 coins of 1
				dp[a] = min(dp[a], dp[a-coin]+1)

				// if the amount is 7 and the coin is 4, it would be:
				// dp[7] = 1 + dp[7-4]
			}
		}
	}

	// handle the edge case of -1
	// if the default value is stored, than we havent found a solution
	if dp[amount] == math.MaxInt64 {
		return -1
	}

	return dp[amount]
}

func main() {
	coins := []int{1, 2, 5}
	amount := 11
	fmt.Println(coinChange(coins, amount)) // Output: 3
}
