package main

import (
	"strconv"
	"strings"
)

func hammingWeight(n int) int {
	binRepr := strconv.FormatInt(int64(n), 2)
	return strings.Count(binRepr, "1")
}
