package _0149

func maxPoints(points [][]int) int {
	var ans int
	if len(points) < 3 {
		return len(points)
	}

	checked := make([][]bool, len(points))
	for i := range checked {
		checked[i] = make([]bool, len(points))
	}
	for i := range points {
		for ii := i + 1; ii < len(points); ii++ {
			if checked[i][ii] {
				continue
			}
			checked[i][ii] = true
			temp_max := 2
			for iii := ii + 1; iii < len(points); iii++ {
				if oneLine(points[i], points[ii], points[iii]) {
					temp_max++
					checked[i][iii] = true
					checked[ii][iii] = true
				}
			}
			if temp_max > ans {
				ans = temp_max
			}
		}
	}
	return ans
}

func oneLine(a, b, c []int) bool {
	return (a[0]-b[0])*(b[1]-c[1]) == (b[0]-c[0])*(a[1]-b[1])
}
