package main

import "sort"

// Time Complexity: O(n log n)
// Space Complexity: O(n)
func topKFrequent(nums []int, k int) []int {
	hm := make(map[int]int)

	for i := range nums {
		if _, ok := hm[nums[i]]; !ok {
			hm[nums[i]] = 1
			continue
		}

		hm[nums[i]]++
	}

	res := make([]value, 0, len(hm))
	for k, v := range hm {
		res = append(res, value{k, v})
	}

	sort.Slice(res, func(i, j int) bool {
		return res[i].Val > res[j].Val // desc
	})

	output := []int{}
	for i := range k {
		output = append(output, res[i].Key)
	}
	return output
}

type value struct {
	Key, Val int
}
