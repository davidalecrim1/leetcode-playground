package main

// Time: O(n)
func maxProfit(prices []int) int {
	maxProfit := 0

	bestBuy := prices[0]
	for i := 1; i < len(prices); i++ {
		bestBuy = min(bestBuy, prices[i])
		sellProfit := prices[i] - bestBuy
		maxProfit = max(maxProfit, sellProfit)
	}

	return maxProfit
}
