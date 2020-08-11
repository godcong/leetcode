package leetcode

/*
130. 被围绕的区域
给定一个二维的矩阵，包含 'X' 和 'O'（字母 O）。

找到所有被 'X' 围绕的区域，并将这些区域里所有的 'O' 用 'X' 填充。

示例:

X X X X
X O O X
X X O X
X O X X
运行你的函数后，矩阵变为：

X X X X
X X X X
X X X X
X O X X
解释:

被围绕的区间不会存在于边界上，换句话说，任何边界上的 'O' 都不会被填充为 'X'。 任何不在边界上，或不与边界上的 'O' 相连的 'O' 最终都会被填充为 'X'。如果两个元素在水平或垂直方向相邻，则称它们是“相连”的。
*/
func solve(board [][]byte) {
	if len(board) < 3 || len(board[0]) < 3 {
		return
	}
	y, x := 0, 0
	for y = 0; y < len(board); y++ {
		for x = 0; x < len(board[y]); x++ {
			if (x == 0) || (y == 0) || (x == len(board[0])-1) || (y == len(board)-1) {
				if board[y][x] == 'O' {
					board[y][x] = '#'
					replaceO(board, x-1, y)
					replaceO(board, x+1, y)
					replaceO(board, x, y-1)
					replaceO(board, x, y+1)
				}
			}
		}
	}

	for y = range board {
		for x = range board[y] {
			if board[y][x] == 'O' {
				board[y][x] = 'X'
			}
			if board[y][x] == '#' {
				board[y][x] = 'O'
			}
		}
	}
}

func replaceO(board [][]byte, x, y int) {
	if (x <= 0) || (y <= 0) || (x >= len(board[0])-1) || (y >= len(board)-1) {
		return
	}
	if board[y][x] == 'X' {
		return
	}

	if board[y][x] == '#' {
		return
	}

	if board[y][x] == 'O' {
		board[y][x] = '#'
		replaceO(board, x-1, y)
		replaceO(board, x+1, y)
		replaceO(board, x, y-1)
		replaceO(board, x, y+1)
	}
}
