package main

// 1. Keep track of the occurence of each int
// 2. Somehow sort this in place to avoid a new N time operation (if I use a hashmap).
// Time: O(n + n log n + k) = n^2
import "slices"

func topKFrequent(nums []int, k int) []int {
	hm := make(map[int]int, k)

	for i := range nums {
		hm[nums[i]]++
	}

	type Item struct {
		num   int
		total int
	}

	items := make([]Item, 0, k)
	for k, v := range hm {
		item := Item{
			num:   k,
			total: v,
		}

		items = append(items, item)
	}

	slices.SortFunc(items, func(a, b Item) int {
		if b.total > a.total {
			return 1
		}

		if b.total == a.total {
			return 0
		}

		return -1
	})

	res := make([]int, 0, k)
	for i := range k {
		res = append(res, items[i].num)
	}

	return res
}
