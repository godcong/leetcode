package leetcode

/*
58. 最后一个单词的长度
给定一个仅包含大小写字母和空格 ' ' 的字符串 s，返回其最后一个单词的长度。如果字符串从左向右滚动显示，那么最后一个单词就是最后出现的单词。

如果不存在最后一个单词，请返回 0 。

说明：一个单词是指仅由字母组成、不包含任何空格字符的 最大子字符串。



示例:

输入: "Hello World"
输出: 5
*/
func lengthOfLastWord(s string) int {
	i := len(s) - 1
	for i >= 0 && s[i] == ' ' {
		s = s[:i]
		i--
	}
	for i >= 0 && s[0] == ' ' {
		s = s[1:]
		i--
	}
	for ; i > 0; i-- {
		if !lengthOfLastWordIsWord(s[i]) {
			break
		}
	}
	if i <= 0 {
		return len(s)
	}
	return len(s) - (i + 1)
}

func lengthOfLastWordIsWord(b byte) bool {
	return (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z')
}
