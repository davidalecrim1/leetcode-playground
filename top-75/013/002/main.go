package main

// Time: O(n log n) because each number may take up to log n bit operations.
// Space: O(n)
func countBits(n int) []int {
	ans := make([]int, n+1)

	for i := 0; i <= n; i++ {
		count := 0
		x := i

		for x > 0 {
			count = count + x&1
			x = x >> 1
		}

		ans[i] = count
	}

	return ans
}
