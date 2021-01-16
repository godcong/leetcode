package code

/*
803. 打砖块
有一个 m x n 的二元网格，其中 1 表示砖块，0 表示空白。砖块 稳定（不会掉落）的前提是：

一块砖直接连接到网格的顶部，或者
至少有一块相邻（4 个方向之一）砖块 稳定 不会掉落时
给你一个数组 hits ，这是需要依次消除砖块的位置。每当消除 hits[i] = (rowi, coli) 位置上的砖块时，对应位置的砖块（若存在）会消失，然后其他的砖块可能因为这一消除操作而掉落。一旦砖块掉落，它会立即从网格中消失（即，它不会落在其他稳定的砖块上）。

返回一个数组 result ，其中 result[i] 表示第 i 次消除操作对应掉落的砖块数目。

注意，消除可能指向是没有砖块的空白位置，如果发生这种情况，则没有砖块掉落。



示例 1：

输入：grid = [[1,0,0,0],[1,1,1,0]], hits = [[1,0]]
输出：[2]
解释：
网格开始为：
[[1,0,0,0]，
 [1,1,1,0]]
消除 (1,0) 处加粗的砖块，得到网格：
[[1,0,0,0]
 [0,1,1,0]]
两个加粗的砖不再稳定，因为它们不再与顶部相连，也不再与另一个稳定的砖相邻，因此它们将掉落。得到网格：
[[1,0,0,0],
 [0,0,0,0]]
因此，结果为 [2] 。
示例 2：

输入：grid = [[1,0,0,0],[1,1,0,0]], hits = [[1,1],[1,0]]
输出：[0,0]
解释：
网格开始为：
[[1,0,0,0],
 [1,1,0,0]]
消除 (1,1) 处加粗的砖块，得到网格：
[[1,0,0,0],
 [1,0,0,0]]
剩下的砖都很稳定，所以不会掉落。网格保持不变：
[[1,0,0,0],
 [1,0,0,0]]
接下来消除 (1,0) 处加粗的砖块，得到网格：
[[1,0,0,0],
 [0,0,0,0]]
剩下的砖块仍然是稳定的，所以不会有砖块掉落。
因此，结果为 [0,0] 。


提示：

m == grid.length
n == grid[i].length
1 <= m, n <= 200
grid[i][j] 为 0 或 1
1 <= hits.length <= 4 * 104
hits[i].length == 2
0 <= xi <= m - 1
0 <= yi <= n - 1
所有 (xi, yi) 互不相同
*/
func hitBricks(grid [][]int, hits [][]int) []int {
	h, w := len(grid), len(grid[0])
	fa := make([]int, h*w+1)
	size := make([]int, h*w+1)
	for i := range fa {
		fa[i] = i
		size[i] = 1
	}
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}
	union := func(from, to int) {
		from, to = find(from), find(to)
		if from != to {
			size[to] += size[from]
			fa[from] = to
		}
	}

	status := make([][]int, h)
	for i, row := range grid {
		status[i] = append([]int(nil), row...)
	}
	// 遍历 hits 得到最终状态
	for _, p := range hits {
		status[p[0]][p[1]] = 0
	}

	// 根据最终状态建立并查集
	root := h * w
	for i, row := range status {
		for j, v := range row {
			if v == 0 {
				continue
			}
			if i == 0 {
				union(i*w+j, root)
			}
			if i > 0 && status[i-1][j] == 1 {
				union(i*w+j, (i-1)*w+j)
			}
			if j > 0 && status[i][j-1] == 1 {
				union(i*w+j, i*w+j-1)
			}
		}
	}

	type pair struct{ x, y int }
	directions := []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // 上下左右

	ans := make([]int, len(hits))
	for i := len(hits) - 1; i >= 0; i-- {
		p := hits[i]
		r, c := p[0], p[1]
		if grid[r][c] == 0 {
			continue
		}

		preSize := size[find(root)]
		if r == 0 {
			union(c, root)
		}
		for _, d := range directions {
			if newR, newC := r+d.x, c+d.y; 0 <= newR && newR < h && 0 <= newC && newC < w && status[newR][newC] == 1 {
				union(r*w+c, newR*w+newC)
			}
		}
		curSize := size[find(root)]
		if cnt := curSize - preSize - 1; cnt > 0 {
			ans[i] = cnt
		}
		status[r][c] = 1
	}
	return ans
}