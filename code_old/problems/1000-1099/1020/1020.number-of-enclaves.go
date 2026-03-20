package _1020

var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func numEnclaves(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	vis := make([][]bool, m)
	for i := range vis {
		vis[i] = make([]bool, n)
	}
	var dfs func(int, int)
	dfs = func(r, c int) {
		if r < 0 || r >= m || c < 0 || c >= n || grid[r][c] == 0 || vis[r][c] {
			return
		}
		vis[r][c] = true
		for _, d := range dirs {
			dfs(r+d.x, c+d.y)
		}
	}
	for i := range grid {
		dfs(i, 0)
		dfs(i, n-1)
	}
	for j := 1; j < n-1; j++ {
		dfs(0, j)
		dfs(m-1, j)
	}
	ans := 0
	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if grid[i][j] == 1 && !vis[i][j] {
				ans++
			}
		}
	}
	return ans
}
