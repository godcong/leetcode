package _0516

func longestPalindromeSubseq(s string) int {
	n := len(s)
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		m := 0
		for j := i - 1; j >= 0; j-- {
			t := dp[j]
			if s[i] == s[j] {
				dp[j] = m + 2
			}
			if t > m {
				m = t
			}
		}
		dp[i] = 1
	}
	max := 0
	for _, v := range dp {
		if v > max {
			max = v
		}
	}
	return max
}
