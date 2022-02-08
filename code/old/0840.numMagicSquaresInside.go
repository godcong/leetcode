package old

/*
840. 矩阵中的幻方
3 x 3 的幻方是一个填充有从 1 到 9 的不同数字的 3 x 3 矩阵，其中每行，每列以及两条对角线上的各数之和都相等。

给定一个由整数组成的 grid，其中有多少个 3 × 3 的 “幻方” 子矩阵？（每个子矩阵都是连续的）。



示例：

输入: [[4,3,8,4],
      [9,5,1,9],
      [2,7,6,2]]
输出: 1
解释:
下面的子矩阵是一个 3 x 3 的幻方：
438
951
276

而这一个不是：
384
519
762

总的来说，在本示例所给定的矩阵中只有一个 3 x 3 的幻方子矩阵。
提示:

1 <= grid.length <= 10
1 <= grid[0].length <= 10
0 <= grid[i][j] <= 15
*/
func numMagicSquaresInside(grid [][]int) int {
	ySize, xSize := len(grid), len(grid[0])
	ans := 0
	for y := 0; y < ySize-2; y++ {
		for x := 0; x < xSize-2; x++ {
			if grid[y+1][x+1] != 5 {
				continue
			}
			if magic(grid[y][x], grid[y][x+1], grid[y][x+2],
				grid[y+1][x], grid[y+1][x+1], grid[y+1][x+2],
				grid[y+2][x], grid[y+2][x+1], grid[y+2][x+2]) {
				ans++
			}
		}
	}

	return ans
}

func magic(vals ...int) bool {
	count := [16]int{}
	val := 0
	for _, val := range vals {
		count[val]++
	}

	for val = 1; val <= 9; val++ {
		if count[val] != 1 {
			return false
		}
	}

	return vals[0]+vals[1]+vals[2] == 15 &&
		vals[3]+vals[4]+vals[5] == 15 &&
		vals[6]+vals[7]+vals[8] == 15 &&
		vals[0]+vals[3]+vals[6] == 15 &&
		vals[1]+vals[4]+vals[7] == 15 &&
		vals[2]+vals[5]+vals[8] == 15 &&
		vals[0]+vals[4]+vals[8] == 15 &&
		vals[2]+vals[4]+vals[6] == 15
}
