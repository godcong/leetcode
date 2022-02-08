package old

import "math"

/*
491. 递增子序列
给定一个整型数组, 你的任务是找到所有该数组的递增子序列，递增子序列的长度至少是2。

示例:

输入: [4, 6, 7, 7]
输出: [[4, 6], [4, 7], [4, 6, 7], [4, 6, 7, 7], [6, 7], [6, 7, 7], [7,7], [4,7,7]]
说明:

给定数组的长度不会超过15。
数组中的整数范围是 [-100,100]。
给定数组中可能包含重复数字，相等的数字应该被视为递增的一种情况。
*/
func findSubsequences(nums []int) [][]int {
	//todo:optimize
	var ans [][]int
	var tmp []int
	var dfs func(cur int, last int, nums []int)
	dfs = func(cur int, last int, nums []int) {
		if cur == len(nums) {
			if len(tmp) >= 2 {
				t := make([]int, len(tmp))
				copy(t, tmp)
				ans = append(ans, t)
			}
			return

		}
		if nums[cur] >= last {
			tmp = append(tmp, nums[cur])
			dfs(cur+1, nums[cur], nums)
			tmp = tmp[:len(tmp)-1]
		}
		if nums[cur] != last {
			dfs(cur+1, last, nums)
		}
	}
	dfs(0, math.MinInt32, nums)
	return ans
}
