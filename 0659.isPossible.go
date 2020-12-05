package leetcode

/*
659. 分割数组为连续子序列
给你一个按升序排序的整数数组 num（可能包含重复数字），请你将它们分割成一个或多个子序列，其中每个子序列都由连续整数组成且长度至少为 3 。

如果可以完成上述分割，则返回 true ；否则，返回 false 。



示例 1：

输入: [1,2,3,3,4,5]
输出: True
解释:
你可以分割出这样两个连续子序列 :
1, 2, 3
3, 4, 5


示例 2：

输入: [1,2,3,3,4,4,5,5]
输出: True
解释:
你可以分割出这样两个连续子序列 :
1, 2, 3, 4, 5
3, 4, 5


示例 3：

输入: [1,2,3,4,4,5]
输出: False
*/
func isPossible(nums []int) bool {
	dp1, dp2, dp3 := 0, 0, 0
	i := 0
	for i < len(nums) {
		start := i
		x := nums[i]
		for i < len(nums) && nums[i] == x {
			i ++
		}
		cnt := i - start

		if start > 0 && x != nums[start-1]+1 {
			if dp1 + dp2 > 0 {
				return false
			} else {
				dp1 = cnt
				dp2, dp3 = 0, 0
			}
		} else {
			if dp1 + dp2 > cnt {
				return false
			}

			left := cnt - dp1 - dp2

			use3 := left
			if left > dp3 {
				use3 = dp3
			}

			dp3 = dp2 + use3
			dp2 = dp1
			dp1 = left - use3

		}

	}
	return dp1 == 0 && dp2 == 0
}
