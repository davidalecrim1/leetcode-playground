package main

import "fmt"

// My Default Solution. This is not so great.
// Time Complexity: O(nˆ2)
// Space Complexity: O(1)
func productExceptSelf_FirstSolution(nums []int) []int {
	result := make([]int, 0, len(nums))

	for i := 0; i < len(nums); i++ {
		var product *int

		for j := 0; j < len(nums); j++ {
			if !(i == j) {
				if product == nil {
					product = &nums[j]
				} else {
					newProduct := *product * nums[j]
					product = &newProduct
				}
			}
		}

		result = append(result, *product)
	}

	return result
}

// Time Complexity: O(n)
// Space Complexity: O(n). This could be O(1), but it's too hard for now.
func productExceptSelf(nums []int) []int {
	prefix := make([]int, len(nums))
	postfix := make([]int, len(nums))
	ans := make([]int, len(nums))

	length := len(nums)

	prefix[0] = 1
	for i := 1; i < length; i++ {
		prefix[i] = prefix[i-1] * nums[i-1]
	}

	postfix[length-1] = 1
	for i := length - 2; i >= 0; i-- {
		postfix[i] = postfix[i+1] * nums[i+1]
	}

	for i := range nums {
		ans[i] = prefix[i] * postfix[i]
	}

	return ans
}

func main() {
	fmt.Printf("productExceptSelf([]int{1, 2, 3, 4}): %v\n", productExceptSelf([]int{1, 2, 3, 4}))
}
