package main

import (
	"fmt"
)

func addBinary(a, b string) string {
	result := ""
	carry := 0

	// Pointers to walk over the binary
	i, j := len(a)-1, len(b)-1

	for i >= 0 || j >= 0 || carry > 0 {
		sum := carry

		if i >= 0 {
			sum += int(a[i] - '0') // Convert '0' or '1' to integer
			i--
		}

		if j >= 0 {
			sum += int(b[j] - '0')
			j--
		}

		// Compute current bit and carry in the sum
		result = string((sum%2)+'0') + result
		carry = sum / 2
	}

	return result
}

func main() {
	a := "11"
	b := "1"
	fmt.Printf("Sum of %s and %s is %s\n", a, b, addBinary(a, b))
}
