package main

// Time Complexity: O(n)
// Space Complexity: O(n)
func productExceptSelf(nums []int) []int {
	prefix := make([]int, len(nums))
	postfix := make([]int, len(nums))

	for i := range nums {
		if i == 0 {
			prefix[i] = nums[i]
		} else {
			prefix[i] = nums[i] * prefix[i-1]
		}
	}

	for j := len(nums) - 1; j >= 0; j-- {
		if j == len(nums)-1 {
			postfix[j] = nums[j]
		} else {
			postfix[j] = nums[j] * postfix[j+1]
		}
	}

	res := make([]int, len(nums))

	for k := range res {
		var pre int
		if k-1 < 0 {
			pre = 1
		} else {
			pre = prefix[k-1]
		}

		var post int
		if k+1 == len(postfix) {
			post = 1
		} else {
			post = postfix[k+1]
		}

		res[k] = pre * post
	}

	return res
}
