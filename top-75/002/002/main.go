package main

// Time: O(nˆ2)
func maxProfit(prices []int) int {
	maxProfit := 0
	for i := 0; i < len(prices); i++ {
		for j := 1; j < len(prices); j++ {
			if j > i {
				maxProfit = max(maxProfit, prices[j]-prices[i])
			}
		}
	}

	return maxProfit
}
