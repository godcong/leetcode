package old

import "sort"

/*
354. 俄罗斯套娃信封问题

给你一个二维整数数组 envelopes ，其中 envelopes[i] = [wi, hi] ，表示第 i 个信封的宽度和高度。

当另一个信封的宽度和高度都比这个信封大的时候，这个信封就可以放进另一个信封里，如同俄罗斯套娃一样。

请计算 最多能有多少个 信封能组成一组“俄罗斯套娃”信封（即可以把一个信封放到另一个信封里面）。

注意：不允许旋转信封。


示例 1：

输入：envelopes = [[5,4],[6,4],[6,7],[2,3]]
输出：3
解释：最多信封的个数为 3, 组合为: [2,3] => [5,4] => [6,7]。

示例 2：

输入：envelopes = [[1,1],[1,1],[1,1]]
输出：1



提示：

    1 <= envelopes.length <= 5000
    envelopes[i].length == 2
    1 <= wi, hi <= 104


*/
func maxEnvelopes(envelopes [][]int) int {
	n := len(envelopes)
	if n < 2 {
		return n
	}
	sort.Slice(envelopes, func(i, j int)bool{return envelopes[i][0] < envelopes[j][0] || (envelopes[i][0] == envelopes[j][0] && envelopes[i][1] > envelopes[j][1])})
	dp := make([]int, n + 1)
	for i := range dp[2:] {
		dp[i+2] = - 1
	}
	for i := 1; i < n; i++ {
		left, right := 1, i
		for left < right {
			mid := (left + right + 1) >> 1
			if dp[mid] == -1 || envelopes[i][1] <= envelopes[dp[mid]][1] ||  envelopes[i][0] <= envelopes[dp[mid]][0] {
				right = mid - 1
			} else {
				left = mid
			}
		}
		if envelopes[i][1] > envelopes[dp[right]][1] && envelopes[i][0] > envelopes[dp[right]][0]{
			if dp[right+1] == -1 || envelopes[dp[right+1]][1] > envelopes[i][1] {
				dp[right+1] = i
			}
		} else {
			if envelopes[dp[right]][1] > envelopes[i][1] {
				dp[right] = i
			}
		}
	}
	for i := len(dp) - 1; i >= 0; i-- {
		if dp[i] > -1 {
			return i
		}
	}
	return 0
}
