package leetcode

/*
37. 解数独

编写一个程序，通过已填充的空格来解决数独问题。

一个数独的解法需遵循如下规则：

数字 1-9 在每一行只能出现一次。
数字 1-9 在每一列只能出现一次。
数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。
空白格用 '.' 表示。



一个数独。



答案被标成红色。

Note:

给定的数独序列只包含数字 1-9 和字符 '.' 。
你可以假设给定的数独只有唯一解。
给定数独永远是 9x9 形式的。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sudoku-solver
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func solveSudoku(board [][]byte) {
	var line, column, block [9][9]bool
	var spaces [][2]int
	for x := range board {
		for y := range board[x] {
			if board[x][y] == '.' {
				spaces = append(spaces, [2]int{x, y})
			} else {
				digit := board[x][y] - '1'
				line[x][digit] = true
				column[y][digit] = true
				block[solveSudokuBlockIndex(x, y)][digit] = true
			}
		}
	}

	var solveSudokuDFS func(int) bool
	solveSudokuDFS = func(pos int) bool {
		if pos == len(spaces) {
			return true
		}
		x, y := spaces[pos][0], spaces[pos][1]
		for digit := byte(0); digit < 9; digit++ {
			if !line[x][digit] && !column[y][digit] && !block[solveSudokuBlockIndex(x, y)][digit] {
				line[x][digit] = true
				column[y][digit] = true
				block[solveSudokuBlockIndex(x, y)][digit] = true
				board[x][y] = digit + '1'
				if solveSudokuDFS(pos + 1) {
					return true
				}
				line[x][digit] = false
				column[y][digit] = false
				block[solveSudokuBlockIndex(x, y)][digit] = false
			}
		}
		return false
	}
	solveSudokuDFS(0)
}

func solveSudokuBlockIndex(x, y int) int {
	switch {
	case x < 3 && y < 3:
		return 0
	case x >= 3 && x < 6 && y < 3:
		return 1
	case x >= 6 && y < 3:
		return 2
	case x < 3 && y >= 3 && y < 6:
		return 3
	case x >= 3 && x < 6 && y >= 3 && y < 6:
		return 4
	case x >= 6 && y >= 3 && y < 6:
		return 5
	case x < 3 && y >= 6:
		return 6
	case x >= 3 && x < 6 && y >= 6:
		return 7
	case x >= 6 && y >= 6:
		return 8
	}
	return -1
}
