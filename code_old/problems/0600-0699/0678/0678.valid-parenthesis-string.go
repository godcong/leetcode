package _0678

func checkValidString(s string) bool {
	minCount, maxCount := 0, 0
	for _, ch := range s {
		if ch == '(' {
			minCount++
			maxCount++
		} else if ch == ')' {
			minCount = max(minCount-1, 0)
			maxCount--
			if maxCount < 0 {
				return false
			}
		} else {
			minCount = max(minCount-1, 0)
			maxCount++
		}
	}
	return minCount == 0
}

func max(i int, i2 int) int {
	if i > i2 {
		return i
	}
	return i2
}
