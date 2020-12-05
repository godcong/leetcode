package leetcode

/*
47. 全排列 II
给定一个可包含重复数字的序列，返回所有不重复的全排列。

示例:

输入: [1,1,2]
输出:
[
  [1,1,2],
  [1,2,1],
  [2,1,1]
]
*/
func permuteUnique(nums []int) [][]int {
	var ret [][]int
	var permuteUniqueDFS func(int)
	permuteUniqueDFS = func(cur int) {
		if cur == len(nums) {
			ret = append(ret, append([]int{}, nums...))
			return
		}
	Loop:
		for i := cur; i < len(nums); i++ {
			for j := cur; j < i; j++ {
				if nums[j] == nums[i] {
					continue Loop
				}
			}
			nums[i], nums[cur] = nums[cur], nums[i]
			permuteUniqueDFS(cur + 1)
			nums[i], nums[cur] = nums[cur], nums[i]
		}
	}
	permuteUniqueDFS(0)
	return ret
}
