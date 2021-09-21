package _0058

func lengthOfLastWord(s string) int {
	index := len(s) - 1
	ans := 0
	for s[index] == ' ' {
		index--
	}
	for index >= 0 && s[index] != ' ' {
		ans++
		index--
	}
	return ans
}
