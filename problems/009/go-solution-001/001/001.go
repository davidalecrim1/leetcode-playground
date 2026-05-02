package main

import (
	"strconv"
)

func isPalindrome(x int) bool {
	if x <= 0 {
		return false
	}
	value := strconv.Itoa(x)
	return value == reverse(value)
}

func reverse(text string) (result string) {
	for _, char := range text {
		result = string(char) + result
	}
	return
}
