package main

func maxProfit(prices []int) int {
	profit := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			profit = profit + prices[i] - prices[i-1]
		}
	}

	return profit
}
