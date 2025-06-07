package main

import "fmt"

func topKFrequent(nums []int, k int) []int {
	hm := make(map[int]int)

	for i := range nums {
		if _, ok := hm[nums[i]]; !ok {
			hm[nums[i]] = 0
		}

		hm[nums[i]]++
	}

	bucket := make([][]int, len(nums)+1)

	for num, freq := range hm {
		bucket[freq] = append(bucket[freq], num)
	}

	res := make([]int, 0, k)
	for i := len(bucket) - 1; i >= 0 && len(res) < k; i-- {
		for _, num := range bucket[i] {
			res = append(res, num)

			if len(res) == k {
				break
			}
		}
	}
	return res
}

func main() {
	fmt.Printf("topKFrequent([]int{1, 1, 1, 2, 2, 3}): %v\n", topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
}
