package leetcode

/*
31. 下一个排列
实现获取下一个排列的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。

如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。

必须原地修改，只允许使用额外常数空间。

以下是一些例子，输入位于左侧列，其相应输出位于右侧列。
1,2,3 → 1,3,2
3,2,1 → 1,2,3
1,1,5 → 1,5,1
*/
func nextPermutation(nums []int) {
	if nums == nil || len(nums) == 0 {
		return
	}
	i := len(nums) - 1
	for i > 0 && nums[i-1] >= nums[i] {
		i--
	}
	if i == 0 {
		for i := 0; i < len(nums)/2; i++ {
			nums[i], nums[len(nums)-i-1] = nums[len(nums)-i-1], nums[i]
		}
		return
	} else {
		j := len(nums) - 1
		for j >= i && nums[j] <= nums[i-1] {
			j--
		}
		nums[j], nums[i-1] = nums[i-1], nums[j]
		for j = i; j < (len(nums)+i)/2; j++ {
			nums[j], nums[len(nums)-j-1+i] = nums[len(nums)-j-1+i], nums[j]
		}
		return
	}
}
