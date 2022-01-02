package _0390

func lastRemaining(n int) int {
	a1 := 1
	k, cnt, step := 0, n, 1
	for cnt > 1 {
		if k%2 == 0 {
			a1 += step
		} else {
			if cnt%2 == 1 {
				a1 += step
			}
		}
		k++
		cnt >>= 1
		step <<= 1
	}
	return a1
}
