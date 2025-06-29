package main

func canFinish(numCourses int, prerequisites [][]int) bool {
	// Strategy
	// Walk the graph
	// See if any route can finish
	// back propagate it

	course := make(map[int][]int)
	view := make([]int, numCourses)

	for _, pre := range prerequisites {
		to, from := pre[0], pre[1]
		course[from] = append(course[from], to)
	}

	var dfs func(i int) bool
	dfs = func(i int) bool {
		if view[i] == 1 { // visiting, found cycle
			return false
		}

		if view[i] == 2 {
			return true // already visited, it's good.
		}

		view[i] = 1 // visiting
		for _, j := range course[i] {
			if !dfs(j) {
				return false
			}
		}
		view[i] = 2 // already visited
		return true
	}

	for i := range numCourses {
		if !dfs(i) {
			return false
		}
	}

	return true
}
