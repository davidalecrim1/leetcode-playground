package main

// Time Complexity: O(1)
func reverseBits(num uint32) uint32 {
	var res uint32 = 0

	for i := 0; i < 32; i++ {
		bit := num & 1     // this will be done in the least significant bit (LSB)
		res = res<<1 | bit // shift the result bits to the right and add the bit using the OR operator.
		num >>= 1          // shift the num to work with the next bit
	}

	return res
}
