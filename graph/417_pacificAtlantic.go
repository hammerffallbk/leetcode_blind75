package graph

func pacificAtlantic(heights [][]int) [][]int {
	rows, cols := len(heights), len(heights[0])
	pacific_reachable := make([][]bool, 0, rows)
	atlantic_reachable := make([][]bool, 0, rows)
	for range rows {
		pacific_reachable = append(pacific_reachable, make([]bool, cols))
		atlantic_reachable = append(atlantic_reachable, make([]bool, cols))
	}

	var dfs func(x, y int, reachable [][]bool)
	dfs = func(x, y int, reachable [][]bool) {
		reachable[x][y] = true
		for _, direction := range [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
			dx, dy := direction[0]+x, direction[1]+y
			if dx >= 0 && dx < rows && dy >= 0 && dy < cols && !reachable[dx][dy] &&
				heights[x][y] <= heights[dx][dy] {
				dfs(dx, dy, reachable)
			}
		}

		// look for 4 directions, if it wasn't reached, and if new r,c is higher than previous
	}
	for c := range cols {
		dfs(0, c, pacific_reachable)
		dfs(rows-1, c, atlantic_reachable)
	}

	for r := range rows {
		dfs(r, 0, pacific_reachable)
		dfs(r, cols-1, atlantic_reachable)
	}
	results := make([][]int, 0)
	for r := range rows {
		for c := range cols {
			if r < rows && c < cols {
				if pacific_reachable[r][c] && atlantic_reachable[r][c] {
					results = append(results, []int{r, c})
				}
			}
		}
	}
	return results
}

//
// {1, 2, 2, 3, 5}
// {3, 2, 3, 4, 4}
// {2, 4, 5, 3, 1}
// {6, 7, 1, 4, 5}
// {5, 1, 1, 2, 4}
