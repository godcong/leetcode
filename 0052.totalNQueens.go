package leetcode

/*
52. N皇后 II
n 皇后问题研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。



上图为 8 皇后问题的一种解法。

给定一个整数 n，返回 n 皇后不同的解决方案的数量。

示例:

输入: 4
输出: 2
解释: 4 皇后问题存在如下两个不同的解法。
[
 [".Q..",  // 解法 1
  "...Q",
  "Q...",
  "..Q."],

 ["..Q.",  // 解法 2
  "Q...",
  "...Q",
  ".Q.."]
]


提示：

皇后，是国际象棋中的棋子，意味着国王的妻子。皇后只做一件事，那就是“吃子”。当她遇见可以吃的棋子时，就迅速冲上去吃掉棋子。当然，她横、竖、斜都可走一或 N-1 步，可进可退。（引用自 百度百科 - 皇后 ）
*/
func totalNQueens(n int) int {
	var solutions [][]string
	var backtrack func(queens []int, n, row int, columns, diagonals1, diagonals2 map[int]bool)
	backtrack = func(queens []int, n, row int, columns, diagonals1, diagonals2 map[int]bool) {
		if row == n {
			board := generateBoard2(queens, n)
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
	return len(solutions)
}

func generateBoard2(queens []int, n int) []string {
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
