package _1716

func totalMoney(n int) int {
	week, day := 1, 1
	ans := 0
	for i := 0; i < n; i++ {
		ans += week + day - 1
		day++
		if day == 8 {
			day = 1
			week++
		}
	}
	return ans
}
