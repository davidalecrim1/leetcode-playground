package main

// Approach 2: Apply memoization in the brute force
/*

The walk tree is like:
robHouse(0)
├── robHouse(2)
│   ├── robHouse(4)
│   │   ├── robHouse(6) → 0
│   │   └── robHouse(5) → 0
│   └── robHouse(3)
│       ├── robHouse(5) → 0
│       └── robHouse(4) (already computed!)
├── robHouse(1)
│   ├── robHouse(3) (already computed!)
│   └── robHouse(2) (already computed!)

*/

// Time Complexity: O(n)
// Space Complexity: O(n)
func rob(nums []int) int {
	memo := make(map[int]int)
	return robMemo(nums, 0, memo)
}

func robMemo(nums []int, index int, memo map[int]int) int {
	if index >= len(nums) {
		return 0
	}

	if val, exists := memo[index]; exists {
		return val
	}

	option1 := nums[index] + robMemo(nums, index+2, memo)
	option2 := robMemo(nums, index+1, memo)

	memo[index] = max(option1, option2)

	return memo[index]
}
