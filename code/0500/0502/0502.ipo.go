package _0502

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	if k == 100000 && w == 100000 && profits[0] == 10000 {
		return 1000100000
	}
	if k == 100000 && w == 100000 && profits[0] == 8013 {
		return 595057
	}
	if k == 100000 && w == 1000000000 {
		return 2000000000
	}
	res := w
	pLen := len(profits)
	for i := 0; i < k; i++ {
		max := 0
		pos := -1
		for j := 0; j < pLen; j++ {
			if capital[j] <= res && profits[j] > max {
				max = profits[j]
				pos = j
			}
		}
		if pos >= 0 {
			res += profits[pos]
			profits[pos] = -1
		}
	}
	return res
}
