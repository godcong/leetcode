package _0583

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([]int, n+1)
	for i := 1; i <= m; i++ {
		lt := 0
		for j := 1; j <= n; j++ {
			tmp := dp[j]
			if word1[i-1] == word2[j-1] {
				dp[j] = lt + 1
			} else {
				dp[j] = max(dp[j], dp[j-1])
			}
			lt = tmp
		}
	}
	return m + n - 2*dp[n]
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
