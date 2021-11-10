package _0495

func findPoisonedDuration(timeSeries []int, duration int) int {
	if len(timeSeries) == 0 {
		return 0
	}
	var res int
	for i := 1; i < len(timeSeries); i++ {
		res += min(timeSeries[i]-timeSeries[i-1], duration)
	}
	res += duration
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
