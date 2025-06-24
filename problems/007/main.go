package main

import "math"

func reverse(x int) int {
	// Strategy
	// I could parse this to a string an swap the houses
	// But this will be inefficient

	// I could use math operations to solve this

	sign := 1
	if x < 0 {
		sign = -1
	}

	y := 0
	x = int(math.Abs(float64(x)))
	for x != 0 {
		mod := x % 10
		x = x / 10

		if y > (math.MaxInt32-mod)/10 || y < (math.MinInt32-mod)/10 {
			return 0
		}

		y = y*10 + mod
	}

	return sign * y
}
