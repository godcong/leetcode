package _1894

func chalkReplacer(chalk []int, k int) int {
	total := 0
	var i int
	for i = range chalk {
		total += chalk[i]
	}
	k %= total
	for i = range chalk {
		if k < chalk[i] {
			return i
		}
		k -= chalk[i]
	}
	return 0
}
