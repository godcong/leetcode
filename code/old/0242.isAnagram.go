package old

/*
242. 有效的字母异位词
给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。

示例 1:

输入: s = "anagram", t = "nagaram"
输出: true
示例 2:

输入: s = "rat", t = "car"
输出: false
说明:
你可以假设字符串只包含小写字母。

进阶:
如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？
*/
func isAnagram(s string, t string) bool {
	m, n := len(s), len(t)
	if m != n {
		return false
	}
	var table [26]int
	for i := 0; i < m; i++ {
		table[s[i]-'a']++
	}
	for i := 0; i < n; i++ {
		table[t[i]-'a']--
		if table[t[i]-'a'] < 0 {
			return false
		}
	}
	return true
}
