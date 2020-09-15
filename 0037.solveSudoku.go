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
	var spaces [81]bool
	for step := 0; step < 81; step++ {
		b := board[step/9][step%9]
		if b == '.' {
			spaces[step] = true
			continue
		}
		num := b - '1'
		blockNum := (step / 9 / 3 * 3) + (step%9)/3
		line[step/9][num] = true
		column[step%9][num] = true
		block[blockNum][num] = true
	}

	var solveSudokuDFS func(int) bool
	solveSudokuDFS = func(pos int) bool {
		if pos == len(spaces) {
			return true
		}
		if !spaces[pos] {
			if solveSudokuDFS(pos + 1) {
				return true
			}
			return false
		}

		x, y := pos/9, pos%9
		blockNum := (pos / 9 / 3 * 3) + (pos%9)/3
		for digit := byte(0); digit < 9; digit++ {
			if !line[x][digit] && !column[y][digit] && !block[blockNum][digit] {
				line[x][digit] = true
				column[y][digit] = true
				block[blockNum][digit] = true
				board[x][y] = digit + '1'
				if solveSudokuDFS(pos + 1) {
					return true
				}
				line[x][digit] = false
				column[y][digit] = false
				block[blockNum][digit] = false
			}
		}
		return false
	}
	solveSudokuDFS(0)
}
