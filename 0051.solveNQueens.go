package leetcode

/*
51. N 皇后
n 皇后问题研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。



上图为 8 皇后问题的一种解法。

给定一个整数 n，返回所有不同的 n 皇后问题的解决方案。

每一种解法包含一个明确的 n 皇后问题的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。



示例：

输入：4
输出：[
 [".Q..",  // 解法 1
  "...Q",
  "Q...",
  "..Q."],

 ["..Q.",  // 解法 2
  "Q...",
  "...Q",
  ".Q.."]
]
解释: 4 皇后问题存在两个不同的解法。


提示：

皇后彼此不能相互攻击，也就是说：任何两个皇后都不能处于同一条横行、纵行或斜线上。
*/
func solveNQueens(n int) [][]string {
	var solutions [][]string
	var backtrack func(queens []int, n, row int, columns, diagonals1, diagonals2 map[int]bool)
	backtrack = func(queens []int, n, row int, columns, diagonals1, diagonals2 map[int]bool) {
		if row == n {
			board := generateBoard(queens, n)
			solutions = append(solutions, board)
			return
		}
		for i := 0; i < n; i++ {
			if columns[i] {
				continue
			}
			diagonal1 := row - i
			if diagonals1[diagonal1] {
				continue
			}
			diagonal2 := row + i
			if diagonals2[diagonal2] {
				continue
			}
			queens[row] = i
			columns[i] = true
			diagonals1[diagonal1], diagonals2[diagonal2] = true, true
			backtrack(queens, n, row+1, columns, diagonals1, diagonals2)
			queens[row] = -1
			delete(columns, i)
			delete(diagonals1, diagonal1)
			delete(diagonals2, diagonal2)
		}
	}

	queens := make([]int, n)
	for i := 0; i < n; i++ {
		queens[i] = -1
	}
	columns := map[int]bool{}
	diagonals1, diagonals2 := map[int]bool{}, map[int]bool{}
	backtrack(queens, n, 0, columns, diagonals1, diagonals2)
	return solutions
}

func generateBoard(queens []int, n int) []string {
	var board []string
	for i := 0; i < n; i++ {
		row := make([]byte, n)
		for j := 0; j < n; j++ {
			row[j] = '.'
		}
		row[queens[i]] = 'Q'
		board = append(board, string(row))
	}
	return board
}
