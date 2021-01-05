package code

/*
767. 重构字符串
给定一个字符串S，检查是否能重新排布其中的字母，使得两相邻的字符不同。

若可行，输出任意可行的结果。若不可行，返回空字符串。

示例 1:

输入: S = "aab"
输出: "aba"
示例 2:

输入: S = "aaab"
输出: ""
注意:

S 只包含小写字母并且长度在[1, 500]区间内。
*/
func reorganizeString(S string) string {
	n := len(S)
	if n <= 1 {
		return S
	}

	cnt := [26]int{}
	maxCnt := 0
	for _, ch := range S {
		ch -= 'a'
		cnt[ch]++
		if cnt[ch] > maxCnt {
			maxCnt = cnt[ch]
		}
	}
	if maxCnt > (n+1)/2 {
		return ""
	}

	ans := make([]byte, n)
	evenIdx, oddIdx, halfLen := 0, 1, n/2
	for i, c := range cnt[:] {
		ch := byte('a' + i)
		for c > 0 && c <= halfLen && oddIdx < n {
			ans[oddIdx] = ch
			c--
			oddIdx += 2
		}
		for c > 0 {
			ans[evenIdx] = ch
			c--
			evenIdx += 2
		}
	}
	return string(ans)
}
