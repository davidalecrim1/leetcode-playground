package main

func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	res := make([]int, n)

	for i := range temperatures {
		diff := 0

		for j := i + 1; j < n; j++ {
			if temperatures[j] > temperatures[i] {
				diff = j - i
				break
			}
		}

		res[i] = diff
	}

	return res
}
