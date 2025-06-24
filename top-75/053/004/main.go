package main

import "slices"

func isAnagram(s string, t string) bool {
	// Strategy
	// Sorting and comparing
	// Time: O(n log n)
	sByte := []byte(s)
	tByte := []byte(t)

	slices.Sort(sByte)
	slices.Sort(tByte)

	return string(sByte) == string(tByte)
}
