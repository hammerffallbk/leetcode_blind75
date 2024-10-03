package graph

const (
	notVisited = iota
	visiting
	visited
)

func canFinish(numCourses int, prerequisites [][]int) bool {
	dependencies := make(map[int][]int)
	for _, prerequisite := range prerequisites {
		dependencies[prerequisite[0]] = append(dependencies[prerequisite[0]], prerequisite[1])
	}

	courseState := map[int]int{}
	var dfs func(course int) bool
	dfs = func(course int) bool {
		if courseState[course] == visiting {
			return false
		}
		if courseState[course] == visited {
			return true
		}
		courseState[course] = visiting
		for _, prerequisite := range dependencies[course] {
			if !dfs(prerequisite) {
				return false
			}
		}
		courseState[course] = visited
		return true
	}

	for i := range numCourses {
		if !dfs(i) {
			return false
		}
	}
	return true
}
