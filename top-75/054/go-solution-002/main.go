package main

import "slices"

func groupAnagrams(strs []string) [][]string {
	/*
	   Use a hashmap to keep track of the groups
	   Op 1: Sort the word an use as a key in the map.
	   Op 2: Use the byte representation of the array (num);
	   PS: Can be flawed.

	   m ==> len(strs)
	   n ==> max(len(str))
	   Time: O(m * n log n)

	   m == > keys
	   n -==> len(strs)
	   Space: O(m * n)
	*/

	hm := map[string][]string{}

	for _, str := range strs {
		sliceStr := []byte(str)
		slices.Sort(sliceStr)

		sortedStr := string(sliceStr)
		hm[sortedStr] = append(hm[sortedStr], str)
	}

	res := [][]string{}
	for _, v := range hm {
		res = append(res, v)
	}

	return res
}
