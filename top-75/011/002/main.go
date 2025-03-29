package main

import "fmt"

func getSum(a int, b int) int {
	for b != 0 { // Keep going until carry is 0
		carry := a & b // Calculate carry
		a = a ^ b      // XOR to get the sum
		b = carry << 1 // Shift carry left
	}
	return a // Return final sum
}

func main() {
	fmt.Printf("getSum(1, 2): %v\n", getSum(1, 2))
}
