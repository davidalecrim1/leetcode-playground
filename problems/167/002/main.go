package main

// Exactly one solution
// 1. Use two pointer. Sum them together, if result is greater than the target, move right, if not, move left.
func twoSum(numbers []int, target int) []int {
	l := 0
	r := len(numbers) - 1

	for l < r {
		// fmt.Println(l,r)

		if target == (numbers[r] + numbers[l]) {
			return []int{l + 1, r + 1}
		}

		if target < (numbers[r] + numbers[l]) {
			r-- // decrease
		} else {
			l++ // increase
		}
	}

	return []int{}
}
