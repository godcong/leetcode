package leetcode

import "fmt"

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
	tmp := make([][]byte, len(board))
	if len(board) <= 2 {
		return
	}

	y, x := 0, 0
	for y = 0; y < len(board); y++ {
		tmp[y] = make([]byte, len(board[y]))
		for x = 0; x < len(board[y]); x++ {
			tmp[y][x] = board[y][x]
		}
	}

	for y = 1; y < len(board)-1; y++ {
		for x = 1; x < len(board[y])-1; x++ {
			board[y][x] = 'X'
		}
	}

	y = 0
	for x = 1; x < len(board[0]); x++ {
		if tmp[y][x] == 'O' {
			board[y][x] = 'O'
			check(tmp, board, x, y+1, len(board[0]))
		}
	}
	y = len(board) - 1
	for x = 1; x < len(board[0]); x++ {
		if tmp[y][x] == 'O' {
			board[y][x] = 'O'
			check(tmp, board, x, y-1, len(board[0]))
		}
	}

	x = 0
	for y = 1; y < len(board); y++ {
		if tmp[y][x] == 'O' {
			board[y][x] = 'O'
			check(tmp, board, x+1, y, len(board))
		}
	}

	x = len(board[0]) - 1
	for y = 1; y < len(board); y++ {
		if tmp[y][x] == 'O' {
			board[y][x] = 'O'
			check(tmp, board, x-1, y, len(board))
		}
	}
}

func check(tmp, board [][]byte, x, y int, max int) {
	fmt.Printf("x:%d,y:%d,tmp:%s,val:%s,max:%d\n", x, y, tmp, board, max)
	if max <= 0 {
		return
	}
	if (x < 1) || (y < 1) || (x >= len(board[y])-1) || (y >= len(board)-1) {
		return
	}

	max--
	if tmp[y][x] == 'O' {
		board[y][x] = 'O'
		check(tmp, board, x-1, y, max)
		check(tmp, board, x+1, y, max)
		check(tmp, board, x, y-1, max)
		check(tmp, board, x, y+1, max)
	}
}
