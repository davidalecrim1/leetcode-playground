package main

// Not a good solution because it's O(n) in the worst case.
// Time Complexity: Worst - O(n)
// Space Complexity: O(1)
func findMin_FirstSolution(nums []int) int {
	// edge case where the number of rotations it the length of the array

	for i := 1; i < len(nums); i++ {
		if nums[i-1] > nums[i] {
			// the amount of rotation is i - 1
			return nums[i]
		}
	}

	return nums[0]
}

// Time Complexity:
// Space Complexity:
func findMin_BestSolution(nums []int) int {
	first, last := 0, len(nums)-1

	for first < last {
		mid := (last + first) / 2
		// the expression first + (last-first)/2 is the same as above,
		// but this is better to prevent integer overflow

		// 5,6,7,8,9,1,2,3,4
		if nums[mid] > nums[last] {
			first = mid + 1 // move right
		} else {
			last = mid // move left
		}
	}
	// always return the number in the left (because its sorted)
	return nums[first]
}

func main() {
}
