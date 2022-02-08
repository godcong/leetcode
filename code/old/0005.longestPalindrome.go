package old

/*
5. 最长回文子串
给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。

示例 1：

输入: "babad"
输出: "bab"
注意: "aba" 也是一个有效答案。
示例 2：

输入: "cbbd"
输出: "bb"
*/

func longestPalindrome(s string) string {
	size := len(s)
	if size <= 1 {
		return s
	}
	gmin, gmax := 0, 0
	min, max := 0, 0
	for i := 0; i < size; i++ {
		if min, max = longestPalindromeMax(s, i, i); max-min > gmax-gmin {
			gmax = max
			gmin = min
		}
		if min, max = longestPalindromeMax(s, i, i+1); max-min > gmax-gmin {
			gmax = max
			gmin = min
		}
	}
	return s[gmin+1 : gmax]
}

func longestPalindromeMax(s string, i, j int) (int, int) {
	for ; i >= 0 && j < len(s); i, j = i-1, j+1 {
		if s[i] != s[j] {
			break
		}
	}
	return i, j
}
