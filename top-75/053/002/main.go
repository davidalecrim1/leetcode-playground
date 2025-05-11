package main

import (
	"slices"
)

func isAnagram(s string, t string) bool {
	bs := []byte(s)
	slices.Sort(bs)
	ts := []byte(t)
	slices.Sort(ts)

	return string(bs) == string(ts)
}
