package _0598

func maxCount(m int, n int, ops [][]int) int {
	mina, minb := m, n
	for _, op := range ops {
		mina = min(mina, op[0])
		minb = min(minb, op[1])
	}
	return mina * minb
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
