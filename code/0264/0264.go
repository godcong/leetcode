package _0264

/*
264. 丑数 II
给你一个整数 n ，请你找出并返回第 n 个 丑数 。

丑数 就是只包含质因数 2、3 和/或 5 的正整数。



示例 1：

输入：n = 10
输出：12
解释：[1, 2, 3, 4, 5, 6, 8, 9, 10, 12] 是由前 10 个丑数组成的序列。
示例 2：

输入：n = 1
输出：1
解释：1 通常被视为丑数。


提示：

1 <= n <= 1690
*/
func nthUglyNumber(n int) int {
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}

	var a, b, c int
	for i := 1; i < n; i++ {
		t1, t2, t3 := dp[a]*2, dp[b]*3, dp[c]*5
		dp[i] = nthUglyNumberMin(t1, t2, t3)
		if dp[i] == t1 {
			a++
		}
		if dp[i] == t2 {
			b++
		}
		if dp[i] == t3 {
			c++
		}
	}
	return dp[n-1]
}

func nthUglyNumberMin(mins ...int) int {
	ret := mins[0]
	for i := 1; i < len(mins)-1; i++ {
		if ret > mins[i] {
			ret = mins[i]
		}
	}
	return ret
}
