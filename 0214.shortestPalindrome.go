package leetcode

/*
214. 最短回文串
给定一个字符串 s，你可以通过在字符串前面添加字符将其转换为回文串。找到并返回可以用这种方式转换的最短回文串。

示例 1:

输入: "aacecaaa"
输出: "aaacecaaa"
示例 2:

输入: "abcd"
输出: "dcbabcd"
*/
func shortestPalindrome(s string) string {
	n := len(s)
	base, mod := 131, 1000000007
	left, right, mul := 0, 0, 1
	best := -1
	for i := 0; i < n; i++ {
		left = (left*base + int(s[i]-'0')) % mod
		right = (right + mul*int(s[i]-'0')) % mod
		if left == right {
			best = i
		}
		mul = mul * base % mod
	}
	add := ""
	if best != n-1 {
		add = s[best+1:]
	}
	b := []byte(add)
	for i := 0; i < len(b)/2; i++ {
		b[i], b[len(b)-1-i] = b[len(b)-1-i], b[i]
	}
	return string(b) + s
}
