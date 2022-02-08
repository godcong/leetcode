package old

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
	var line, column [9][9]bool
	var block [3][3][9]bool
	var spaces [81]bool
	var x, y byte
	var num byte
	for step := byte(0); step < 81; step++ {
		x, y = step/9, step%9
		if board[x][y] == '.' {
			spaces[step] = true
			continue
		}
		num = board[x][y] - '1'
		line[x][num] = true
		column[y][num] = true
		block[x/3][y/3][num] = true
	}

	var solveSudokuDFS func(byte) bool
	solveSudokuDFS = func(pos byte) bool {
		if pos == 81 {
			return true
		}
		if !spaces[pos] {
			if solveSudokuDFS(pos + 1) {
				return true
			}
			return false
		}
		x, y := pos/9, pos%9
		for digit := byte(0); digit < 9; digit++ {
			if !line[x][digit] && !column[y][digit] && !block[x/3][y/3][digit] {
				line[x][digit] = true
				column[y][digit] = true
				block[x/3][y/3][digit] = true
				board[x][y] = digit + '1'
				if solveSudokuDFS(pos + 1) {
					return true
				}
				line[x][digit] = false
				column[y][digit] = false
				block[x/3][y/3][digit] = false
			}
		}
		return false
	}
	solveSudokuDFS(0)
}
