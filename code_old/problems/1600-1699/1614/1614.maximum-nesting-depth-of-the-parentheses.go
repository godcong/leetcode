package _1614

func maxDepth(s string) int {
	size := 0
	ans := 0
	for _, ch := range s {
		if ch == '(' {
			size++
			if size > ans {
				ans = size
			}
		} else if ch == ')' {
			size--
		}
	}
	return ans
}
