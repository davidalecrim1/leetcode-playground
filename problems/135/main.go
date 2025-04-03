package main

// Time Complexity: O(n)
func candy(ratings []int) int {
	// Approach: Greedy (Brute Force) passing the array two times.
	candy := make([]int, len(ratings))

	// start everybody with one
	for i := range candy {
		candy[i] = 1
	}

	// skip first
	for i := 1; i < len(ratings); i++ {
		if ratings[i-1] < ratings[i] {
			candy[i] = candy[i-1] + 1
			// 2 because everybody should have 1. This is an extra 1.
		}
	}

	// skip last
	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			res := max(candy[i], candy[i+1]+1)
			candy[i] = res
		}
	}

	sum := 0
	for _, value := range candy {
		sum += value
	}
	return sum
}
