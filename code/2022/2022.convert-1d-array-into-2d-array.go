package _2022

func construct2DArray(original []int, m int, n int) [][]int {
	res := make([][]int, 0)
	if len(original) != m*n {
		return res
	}
	for i := 0; i < m; i++ {
		temp := make([]int, n)
		for j := 0; j < n; j++ {
			temp[j] = original[i*n+j]
		}
		res = append(res, temp)
	}

	return res
}
