package leetcode

/*
18. 四数之和
给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c + d 的值与 target 相等？找出所有满足条件且不重复的四元组。

注意：

答案中不可以包含重复的四元组。

示例：

给定数组 nums = [1, 0, -1, 0, -2, 2]，和 target = 0。

满足要求的四元组集合为：
[
  [-1,  0, 0, 1],
  [-2, -1, 1, 2],
  [-2,  0, 0, 2]
]
*/
func fourSum(nums []int, target int) [][]int {
	var ret [][]int
	var fourSum func(n []int, idx, t int)
	fourSum = func(n []int, idx, t int) {
		if t == 0 {
			ret = append(ret, append([]int{}, n...))
			return
		}
		for ; idx < len(nums); idx++ {
			n = append(n, nums[idx])
			fourSum(n, idx, target-nums[idx])
		}
	}
	fourSum(nil, 0, target)

	return ret
}
