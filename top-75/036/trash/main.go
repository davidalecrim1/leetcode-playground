package main

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	hm := make(map[[2]int]struct{})
	for i := 0; i < len(intervals); i++ {
		for j := 1; j < len(intervals); j++ {
			prev := intervals[i]
			curr := intervals[j]

			if overlap(prev, curr) || overlap(curr, prev) {
				idx := [2]int{
					min(prev[0], curr[0]),
					max(prev[1], curr[1]),
				}

				hm[idx] = struct{}{}
				continue
			}

			hm[[2]int{prev[0], prev[1]}] = struct{}{}
			hm[[2]int{curr[0], curr[1]}] = struct{}{}
		}
	}

	res := [][]int{}
	for k := range hm {
		res = append(res, k[:])
	}

	return res
}

func overlap(first, second []int) bool {
	return (first[0] >= second[0] && first[0] <= second[1]) || (first[1] >= second[0] && first[1] <= second[1])
}
