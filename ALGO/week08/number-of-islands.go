package week08

func numIslands(grid [][]byte) int {
	num := 0
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < len(grid) && i >= 0 && j < len(grid[0]) && j >= 0 && grid[i][j] == '1' {
			grid[i][j] = 0
			dfs(i, j+1)
			dfs(i, j-1)
			dfs(i-1, j)
			dfs(i+1, j)
		}
	}
	for i := 0; i < len(grid[0]); i++ {
		for j := 0; j < len(grid); j++ {
			if grid[j][i] == '1' {
				num++
				dfs(j, i)
			}
		}
	}
	return num
}
