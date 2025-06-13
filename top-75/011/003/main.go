package main

func getSum(a int, b int) int {
	// Use logical ports
	// XOR and AND with bitwise manipulation

	// A sum operation can be simulated with a loop using XOR to sum all the values,
	// but it doesnt take care of the carry.

	// For the 1 of carry when we have the sum of 1 AND 1. We do and AND operator in both A and B,
	// than we place the bit in the house to the left (<<).

	for b != 0 {
		carry := a & b // AND to see if we have a carry
		a = a ^ b      // XOR to the get the sum result
		b = carry << 1 // If carry is not zero, we shift it to the left to it's house.
	}

	return a
}
