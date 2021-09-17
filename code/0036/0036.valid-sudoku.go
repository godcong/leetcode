package _0036

func isValidSudoku(board [][]byte) bool {
	var rows, columns [9][9]int
	var boxes [3][3][9]int
	for i, row := range board {
		for j, c := range row {
			if c == '.' {
				continue
			}
			index := c - '1'
			rows[i][index]++
			columns[j][index]++
			boxes[i/3][j/3][index]++
			if rows[i][index] > 1 || columns[j][index] > 1 || boxes[i/3][j/3][index] > 1 {
				return false
			}
		}
	}
	return true
}
