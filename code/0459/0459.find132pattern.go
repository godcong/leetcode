package _0459

/*
456. 132 模式

给你一个整数数组 nums ，数组中共有 n 个整数。132 模式的子序列 由三个整数 nums[i]、nums[j] 和 nums[k] 组成，并同时满足：i < j < k 和 nums[i] < nums[k] < nums[j] 。

如果 nums 中存在 132 模式的子序列 ，返回 true ；否则，返回 false 。



进阶：很容易想到时间复杂度为 O(n^2) 的解决方案，你可以设计一个时间复杂度为 O(n logn) 或 O(n) 的解决方案吗？



示例 1：

输入：nums = [1,2,3,4]
输出：false
解释：序列中不存在 132 模式的子序列。

示例 2：

输入：nums = [3,1,4,2]
输出：true
解释：序列中有 1 个 132 模式的子序列： [1, 4, 2] 。

示例 3：

输入：nums = [-1,3,2,0]
输出：true
解释：序列中有 3 个 132 模式的的子序列：[-1, 3, 2]、[-1, 3, 0] 和 [-1, 2, 0] 。



提示：

    n == nums.length
    1 <= n <= 104
    -109 <= nums[i] <= 109
*/
func find132pattern(nums []int) bool {
	n := len(nums)
	if n < 3 {
		return false
	}
	minNums := make([]int, n)
	minNums[0] = nums[0]
	for i := 1; i < n; i++ {
		if nums[i] < minNums[i-1] {
			minNums[i] = nums[i]
		} else {
			minNums[i] = minNums[i-1]
		}
	}
	stack := make([]int, 0)
	for j := len(nums) - 1; j >= 0; j-- {
		if minNums[j] < nums[j] {
			for len(stack) > 0 && minNums[j] >= stack[len(stack)-1] {
				stack = stack[:len(stack)-1]
			}
			if len(stack) > 0 && stack[len(stack)-1] < nums[j] {
				return true
			}
			stack = append(stack, nums[j])
		}
	}
	return false
}
