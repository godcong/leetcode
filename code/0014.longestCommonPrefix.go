package code

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

	j := len(strs) - 1
	i, max := 1, len(strs[0])

	//for ; i <= j; i++ {
	//	if max > (len(strs[i])) {
	//		max = len(strs[i])
	//	}
	//}

Loop:
	for i = 0; i < max; i++ {
		for j = len(strs) - 1; j > 0; j-- {
			if i > len(strs[j])-1 || strs[0][i] != strs[j][i] {
				break Loop
			}
		}
	}

	return strs[0][:i]
}
