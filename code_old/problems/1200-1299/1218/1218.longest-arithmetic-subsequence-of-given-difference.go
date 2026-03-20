package _1218

func longestSubsequence(arr []int, difference int) int {
	ans := 0
	dp := make(map[int]int, len(arr))
	for _, v := range arr {
		dp[v] = dp[v-difference] + 1
		if dp[v] > ans {
			ans = dp[v]
		}
	}
	return ans
}
