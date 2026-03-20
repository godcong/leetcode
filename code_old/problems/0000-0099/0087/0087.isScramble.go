package _0087

/*
87. 扰乱字符串
使用下面描述的算法可以扰乱字符串 s 得到字符串 t ：
如果字符串的长度为 1 ，算法停止
如果字符串的长度 > 1 ，执行下述步骤：
在一个随机下标处将字符串分割成两个非空的子字符串。即，如果已知字符串 s ，则可以将其分成两个子字符串 x 和 y ，且满足 s = x + y 。
随机 决定是要「交换两个子字符串」还是要「保持这两个子字符串的顺序不变」。即，在执行这一步骤之后，s 可能是 s = x + y 或者 s = y + x 。
在 x 和 y 这两个子字符串上继续从步骤 1 开始递归执行此算法。
给你两个 长度相等 的字符串 s1 和 s2，判断 s2 是否是 s1 的扰乱字符串。如果是，返回 true ；否则，返回 false 。



示例 1：

输入：s1 = "great", s2 = "rgeat"
输出：true
解释：s1 上可能发生的一种情形是：
"great" --> "gr/eat" // 在一个随机下标处分割得到两个子字符串
"gr/eat" --> "gr/eat" // 随机决定：「保持这两个子字符串的顺序不变」
"gr/eat" --> "g/r / e/at" // 在子字符串上递归执行此算法。两个子字符串分别在随机下标处进行一轮分割
"g/r / e/at" --> "r/g / e/at" // 随机决定：第一组「交换两个子字符串」，第二组「保持这两个子字符串的顺序不变」
"r/g / e/at" --> "r/g / e/ a/t" // 继续递归执行此算法，将 "at" 分割得到 "a/t"
"r/g / e/ a/t" --> "r/g / e/ a/t" // 随机决定：「保持这两个子字符串的顺序不变」
算法终止，结果字符串和 s2 相同，都是 "rgeat"
这是一种能够扰乱 s1 得到 s2 的情形，可以认为 s2 是 s1 的扰乱字符串，返回 true
示例 2：

输入：s1 = "abcde", s2 = "caebd"
输出：false
示例 3：

输入：s1 = "a", s2 = "a"
输出：true


提示：

s1.length == s2.length
1 <= s1.length <= 30
s1 和 s2 由小写英文字母组成
*/
func isScramble(s1 string, s2 string) bool {
	n := len(s1)
	dp := make([][][]int8, n)
	for i := range dp {
		dp[i] = make([][]int8, n)
		for j := range dp[i] {
			dp[i][j] = make([]int8, n+1)
			for k := range dp[i][j] {
				dp[i][j][k] = -1
			}
		}
	}

	var dfs func(i1, i2, length int) int8
	dfs = func(i1, i2, length int) (res int8) {
		d := &dp[i1][i2][length]
		if *d != -1 {
			return *d
		}
		defer func() { *d = res }()

		x, y := s1[i1:i1+length], s2[i2:i2+length]
		if x == y {
			return 1
		}

		freq := [26]int{}
		for i, ch := range x {
			freq[ch-'a']++
			freq[y[i]-'a']--
		}
		for _, f := range freq[:] {
			if f != 0 {
				return 0
			}
		}

		for i := 1; i < length; i++ {

			if dfs(i1, i2, i) == 1 && dfs(i1+i, i2+i, length-i) == 1 {
				return 1
			}

			if dfs(i1, i2+length-i, i) == 1 && dfs(i1+i, i2, length-i) == 1 {
				return 1
			}
		}

		return 0
	}
	return dfs(0, 0, n) == 1
}
