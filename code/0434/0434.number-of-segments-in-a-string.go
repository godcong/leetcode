package _0434

func countSegments(s string) int {
	var ans int
	for i, ch := range s {
		if (i == 0 || s[i-1] == ' ') && ch != ' ' {
			ans++
		}
	}
	return ans
}
