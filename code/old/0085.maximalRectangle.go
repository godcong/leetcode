package old

/*
85. 最大矩形
给定一个仅包含 0 和 1 、大小为 rows x cols 的二维二进制矩阵，找出只包含 1 的最大矩形，并返回其面积。



示例 1：


输入：matrix = [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]
输出：6
解释：最大矩形如上图所示。
示例 2：

输入：matrix = []
输出：0
示例 3：

输入：matrix = [["0"]]
输出：0
示例 4：

输入：matrix = [["1"]]
输出：1
示例 5：

输入：matrix = [["0","0"]]
输出：0


提示：

rows == matrix.length
cols == matrix[0].length
0 <= row, cols <= 200
matrix[i][j] 为 '0' 或 '1'
*/
func maximalRectangle(matrix [][]byte) int {
	var ans int
	if len(matrix) == 0 {
		return ans
	}
	m, n := len(matrix), len(matrix[0])
	left := make([][]int, m)
	for i, row := range matrix {
		left[i] = make([]int, n)
		for j, v := range row {
			if v == '0' {
				continue
			}
			if j == 0 {
				left[i][j] = 1
			} else {
				left[i][j] = left[i][j-1] + 1
			}
		}
	}
	for j := 0; j < n; j++ {
		up := make([]int, m)
		down := make([]int, m)
		stk := []int{}
		for i, l := range left {
			for len(stk) > 0 && left[stk[len(stk)-1]][j] >= l[j] {
				stk = stk[:len(stk)-1]
			}
			up[i] = -1
			if len(stk) > 0 {
				up[i] = stk[len(stk)-1]
			}
			stk = append(stk, i)
		}
		stk = nil
		for i := m - 1; i >= 0; i-- {
			for len(stk) > 0 && left[stk[len(stk)-1]][j] >= left[i][j] {
				stk = stk[:len(stk)-1]
			}
			down[i] = m
			if len(stk) > 0 {
				down[i] = stk[len(stk)-1]
			}
			stk = append(stk, i)
		}
		for i, l := range left {
			height := down[i] - up[i] - 1
			area := height * l[j]
			ans = maximalRectangleMax(ans, area)
		}
	}
	return ans
}

func maximalRectangleMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
