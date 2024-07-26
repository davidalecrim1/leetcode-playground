/*
	Time Complexity: O(n) or O(n/2) if the number is even digit count.
	Space Complexity: O(1)
*/

package main

func isPalindrome(x int) bool {
	if x <= 0 || (x%10 == 0) {
		return false
	}

	half := 0

	for x > half {
		half = (half * 10) + (x % 10)
		x = x / 10
	}

	return x == half || x == half/10

}
