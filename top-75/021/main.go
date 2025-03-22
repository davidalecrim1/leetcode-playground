package main

/*
candidates = [2,3,6,7], target = 7

We need to search for all possible combinations that sum up to target. Since we can reuse numbers, this is a backtracking problem.

candidates[0] = 2
add 2
add 2
add 2 (end or remove)

add 3 (end or remove)
...

rollback to add 2 + 2
add 3 (so this should be based)


*/

func combinationSum(candidates []int, target int) [][]int {
	var res [][]int

	backtrack := func(start int, curr []int, remainingTarget int) {}
	backtrack = func(start int, curr []int, remainingTarget int) {
		if remainingTarget == 0 {
			combination := make([]int, len(curr))
			// make sure when the slice is changed, that the saved value is not altered
			copy(combination, curr)
			res = append(res, combination)
			return
		}

		if remainingTarget < 0 {
			return // the combination is over the target.
		}

		for i := start; i < len(candidates); i++ {
			curr = append(curr, candidates[i]) // choose a number. e.g. backtrack with 2 until it returns.

			// calculate target minus option until it reaches zero (condition above)
			remaining := remainingTarget - candidates[i]

			backtrack(i, curr, remaining)
			// when the backtrack returns, this will be returned with a valid combination or over in the ifs above

			curr = curr[:len(curr)-1] // pop the last to undo the last number added and keep trying
		}
	}

	backtrack(0, []int{}, target)
	return res
}
