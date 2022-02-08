package _1104

func pathInZigZagTree(label int) []int {
	row, rowStart := 1, 1
	var path []int
	for rowStart*2 <= label {
		row++
		rowStart *= 2
	}
	if row%2 == 0 {
		label = getReverse(label, row)
	}
	for row > 0 {
		if row%2 == 0 {
			path = append(path, getReverse(label, row))
		} else {
			path = append(path, label)
		}
		row--
		label >>= 1
	}
	for i, n := 0, len(path); i < n/2; i++ {
		path[i], path[n-1-i] = path[n-1-i], path[i]
	}
	return path
}

func getReverse(label, row int) int {
	return 1<<(row-1) + 1<<row - 1 - label
}
