package main

import "reflect"

// Strategy HashMap
// Time complexity: O(s + t)
// Space: O(s + t)
func isAnagram(s string, t string) bool {
	given := make(map[rune]int)

	for _, char := range s {
		if val, ok := given[char]; ok {
			given[char] = val + 1
		} else {
			given[char] = 1
		}
	}

	want := make(map[rune]int)

	for _, char := range t {
		if val, ok := want[char]; ok {
			want[char] = val + 1
		} else {
			want[char] = 1
		}
	}

	return reflect.DeepEqual(given, want)
}
