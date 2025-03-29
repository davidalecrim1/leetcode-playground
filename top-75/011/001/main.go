package main

func getSum(a int, b int) int {
	for b != 0 {
		sum := a ^ b          // XOR operator
		carry := (a & b) << 1 // AND operator and shift to the left

		a = sum
		b = carry
	}
	return a
}
