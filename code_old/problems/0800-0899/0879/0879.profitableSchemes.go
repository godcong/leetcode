package _0879

func profitableSchemes(n int, minProfit int, group []int, profit []int) int {
	g := len(group)
	dp := make([][]int, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, minProfit+1)
	}
	for i := 0; i < n+1; i++ {
		dp[i][0] = 1
	}
	for i := 0; i < g; i++ {
		manual := group[i]
		money := profit[i]
		for j := n; j >= manual; j-- {
			for k := minProfit; k >= 0; k-- {
				dp[j][k] += dp[j-manual][max(0, k-money)]
				dp[j][k] %= 1e9 + 7
			}
		}
	}
	return dp[n][minProfit]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
