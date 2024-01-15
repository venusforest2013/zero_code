package dfs

func GetIslandCount(grid [][]byte) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	var dfs func([][]byte, int, int)
	var ret int

	dfs = func(grid [][]byte, i, j int) {
		visited[i][j] = true
		if i-1 >= 0 && grid[i-1][j] == '1' && visited[i-1][j] == false {
			dfs(grid, i-1, j)
		}
		if i+1 < m && grid[i+1][j] == '1' && visited[i+1][j] == false {
			dfs(grid, i+1, j)
		}
		if j-1 >= 0 && grid[i][j-1] == '1' && visited[i][j-1] == false {
			dfs(grid, i, j-1)
		}
		if j+1 < n && grid[i][j+1] == '1' && visited[i][j+1] == false {
			dfs(grid, i, j+1)
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' && visited[i][j] == false {
				dfs(grid, i, j)
				ret++
			}
		}
	}
	return ret
}
