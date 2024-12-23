package main

func maxProfit(prices []int) int {
	return maxProfit_MyDefaultSolution(prices)
}

// This solution is the same performance as the below,
// but I think is hard to see it with my mind
func maxProfit_SlidingWindowFromNeetCode(prices []int) int {
	left := 0
	right := 1
	maxProfit := 0

	for right < len(prices) {
		if prices[left] < prices[right] {
			profit := prices[right] - prices[left]

			if profit > maxProfit {
				maxProfit = profit
			}
		} else {
			left = right
		}

		right++
	}

	return maxProfit
}

// Approach: One pass with two pointers
// Time: O(n)
// Space: O(1)
func maxProfit_MyDefaultSolution(prices []int) int {
	low := 0
	high := 0
	maxProfit := 0

	for i := 1; i < len(prices)-1; i++ {
		if prices[i] < prices[low] {
			low = i
		}

		if prices[i] > prices[low] {
			high = i

			profit := prices[high] - prices[low]

			if profit > maxProfit {
				maxProfit = profit
			}
		}
	}

	return maxProfit
}
