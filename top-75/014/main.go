package main

import (
	"fmt"
	"strconv"
)

// Time complexity: O(1)
func reverseBits(num uint32) uint32 {
	binaryRep := fmt.Sprintf("%032b", num)

	reverseBinary := ""
	for i := len(binaryRep) - 1; i >= 0; i-- {
		reverseBinary += string(binaryRep[i])
	}

	intRepResult, _ := strconv.ParseUint(reverseBinary, 2, 32)
	return uint32(intRepResult)
}
