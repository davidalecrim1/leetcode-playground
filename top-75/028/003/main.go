package main

func canFinish(numCourses int, prerequisites [][]int) bool {
	courses := make(map[int][]int)
	view := make(map[int]int)

	for _, pre := range prerequisites {
		to, from := pre[0], pre[1]
		courses[from] = append(courses[from], to)
	}

	var dfs func(i int) bool
	dfs = func(i int) bool {
		if view[i] == 1 {
			return false // found cycle
		}

		if view[i] == 2 {
			return true // already visited, we're good
		}

		view[i] = 1 // visiting
		for _, c := range courses[i] {
			if !dfs(c) {
				return false
			}
		}
		view[i] = 2 // visited
		return true
	}

	for i := range numCourses {
		if !dfs(i) {
			return false
		}
	}

	return true
}
