package main

// The minimum of the problem is 1 for speed of eating
// The maximum is the largest element on the array
// We can do a binary search to find the optiomal minimum based on the input

func minEatingSpeed(piles []int, h int) int {
	maxSpeed := 0
	for i := range piles {
		maxSpeed = max(piles[i], maxSpeed)
	}

	left := 1
	right := maxSpeed
	for left < right {
		mid := left + (right-left)/2
		if isValid(piles, mid, h) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

func isValid(piles []int, k int, h int) bool {
	actual := 0
	for i := range piles {
		div := (piles[i] + k - 1) / k
		actual += div
	}

	return actual <= h
}
