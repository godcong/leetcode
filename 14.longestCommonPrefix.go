package leetcode

/*
14. 最长公共前缀
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

示例 1:

输入: ["flower","flow","flight"]
输出: "fl"
示例 2:

输入: ["dog","racecar","car"]
输出: ""
解释: 输入不存在公共前缀。
说明:

所有输入只包含小写字母 a-z 。
*/
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	i, max := 1, len(strs[0])
	str, pre := "", byte(0)

	for ; i < len(strs); i++ {
		if max > len(strs[i]) {
			max = len(strs[i])
		}
	}

Loop:
	for i = 0; i < max; i++ {
		pre = strs[0][i]
		for _, str = range strs[1:] {
			if pre != str[i] {
				break Loop
			}
		}
	}

	return strs[0][:i]
}
