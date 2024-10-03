package graph

const (
	zero = '0'
	one  = '1'
)

func numIslands(grid [][]byte) int {
	visited := make([][]bool, len(grid))
	rows, cols := len(grid), len(grid[0])
	for r := range grid {
		visited[r] = make([]bool, len(grid[r]))
	}

	var result int
	var dfs func(r, c int, visited [][]bool)
	dfs = func(r, c int, visited [][]bool) {
		if r == rows || c == cols || r < 0 || c < 0 {
			return
		}
		if visited[r][c] {
			return
		}
		visited[r][c] = true
		if grid[r][c] == zero {
			return
		}
		dfs(r+1, c, visited)
		dfs(r, c+1, visited)
		dfs(r-1, c, visited)
		dfs(r, c-1, visited)
	}

	for x := range grid {
		for y := range grid[x] {
			if grid[x][y] == one && !visited[x][y] {
				result++
				dfs(x, y, visited)
			}
		}
	}

	return result
}
