package _1221

func balancedStringSplit(s string) int {
	d := 0
	ans := 0
	for _, ch := range s {
		if ch == 'L' {
			d++
		} else {
			d--
		}
		if d == 0 {
			ans++
		}
	}
	return ans
}
