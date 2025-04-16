package main

func canFinish(numCourses int, prerequisites [][]int) bool {
	// Approach: Apply DFS and keep track of visited.
	preReqs := make(map[int][]int)
	for _, graph := range prerequisites {
		source, preReq := graph[0], graph[1]
		preReqs[source] = append(preReqs[source], preReq)
	}

	pathVisited := make(map[int]bool)
	var dfs func(int) bool
	dfs = func(course int) bool {
		if exists := pathVisited[course]; exists {
			return false
		}

		if preReqs[course] == nil {
			return true
		}

		pathVisited[course] = true
		coursePreReqs := preReqs[course]
		for _, preReq := range coursePreReqs {
			if !dfs(preReq) {
				return false
			}
		}

		pathVisited[course] = false
		preReqs[course] = nil // as in we can reach the whole rest of the graph from this
		// so we memoize it saying we can reach the rest.
		return true
	}

	for course := range numCourses {
		if !dfs(course) {
			return false
		}
	}

	return true
}
