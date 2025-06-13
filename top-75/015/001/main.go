package main

import (
	"fmt"
	"strconv"
)

func reverseBits(num uint32) uint32 {
	return reverseBits_NotUsingBuitInConverter(num)
}

// Time complexity: O(1)
func reverseBits_NotUsingBuitInConverter(num uint32) uint32 {
	var res uint32

	for i := range 32 {
		// Extract the least significant bit (first bit on the right)
		bit := num & 1

		// Calculate the reversed position
		reversedPosition := 31 - i

		// Get the reverse bit
		bit = bit << reversedPosition

		// Place the bit in its reversed position
		res = res | bit

		// Shift the input number to the right
		num = num >> 1
	}

	return res
}

// Time complexity: O(1)
// Strategy: Use built-in functions for this.
func reverseBits_UsingBuitInConverter(num uint32) uint32 {
	binaryRep := fmt.Sprintf("%032b", num)

	reverseBinary := ""
	for i := len(binaryRep) - 1; i >= 0; i-- {
		reverseBinary += string(binaryRep[i])
	}

	intRepResult, _ := strconv.ParseUint(reverseBinary, 2, 32)
	return uint32(intRepResult)
}
