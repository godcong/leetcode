package old

import "sort"

/*
581. 最短无序连续子数组
给定一个整数数组，你需要寻找一个连续的子数组，如果对这个子数组进行升序排序，那么整个数组都会变为升序排序。

你找到的子数组应是最短的，请输出它的长度。

示例 1:

输入: [2, 6, 4, 8, 10, 9, 15]
输出: 5
解释: 你只需要对 [6, 4, 8, 10, 9] 进行升序排序，那么整个表都会变为升序排序。
说明 :

输入的数组长度范围在 [1, 10,000]。
输入的数组可能包含重复元素 ，所以升序的意思是<=。
*/
func findUnsortedSubarray(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	tmp := make([]int, len(nums))
	copy(tmp, nums)
	sort.Ints(tmp)
	i, j := 0, len(tmp)-1
	for ; i < j; i++ {
		if tmp[i] != nums[i] {
			break
		}
	}

	for ; j > i; j-- {
		if tmp[j] != nums[j] {
			break
		}
	}
	if i == j {
		return 0
	}

	return j - i + 1
}
