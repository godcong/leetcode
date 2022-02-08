package _1725

func countGoodRectangles(rectangles [][]int) int {
	maxLen := 0
	ans := 0
	for _, rect := range rectangles {
		k := min(rect[0], rect[1])
		if k == maxLen {
			ans++
		} else if k > maxLen {
			maxLen, ans = k, 1
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
