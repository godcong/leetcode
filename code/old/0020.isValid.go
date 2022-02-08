package old

/*
20. 有效的括号
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。

示例 1:

输入: "()"
输出: true
示例 2:

输入: "()[]{}"
输出: true
示例 3:

输入: "(]"
输出: false
示例 4:

输入: "([)]"
输出: false
示例 5:

输入: "{[]}"
输出: true
*/
func isValid(s string) bool {
	if s == "" {
		return true
	}

	if len(s) < 2 {
		return false
	}

	count, pre := 0, make([]byte, len(s)/2+1)
	for i := 0; i < len(s); i++ {
		if count > len(s)/2 {
			return false
		}
		if s[i] == '[' || s[i] == '(' || s[i] == '{' {
			count, pre[count] = count+1, s[i]
			continue
		}
		if count < 1 || s[i]-pre[count-1] > 2 {
			return false
		}
		count--
	}
	if count > 0 {
		return false
	}
	return true
}
