package main

import "slices"

// Time Complexity: O(m * n log n)
// m is the len of strs.
// n log n is the sorting complexity of each string.
func groupAnagrams(strs []string) [][]string {
	// Approach: Sort the strings to group them together.
	hm := make(map[string][]string)

	for _, str := range strs {
		sortedStr := sortedString(str)
		hm[sortedStr] = append(hm[sortedStr], str)
	}

	var res [][]string

	for _, v := range hm {
		res = append(res, v)
	}

	return res
}

func sortedString(s string) string {
	b := []byte(s)
	slices.Sort(b)
	return string(b)
}
