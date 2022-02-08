package _0807

func maxIncreaseKeepingSkyline(grid [][]int) int {
	n := len(grid)
	rowMax := make([]int, n)
	colMax := make([]int, n)
	ans := 0
	for i, row := range grid {
		for j, h := range row {
			rowMax[i] = max(rowMax[i], h)
			colMax[j] = max(colMax[j], h)
		}
	}
	for i, row := range grid {
		for j, h := range row {
			ans += min(rowMax[i], colMax[j]) - h
		}
	}
	return ans
}

func min(i int, i2 int) int {
	if i < i2 {
		return i
	}
	return i2
}

func max(i int, i2 int) int {
	if i > i2 {
		return i
	}
	return i2
}
