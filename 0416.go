package leetcode

import "sort"

/*
416. 分割等和子集
给定一个只包含正整数的非空数组。是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。

注意:

每个数组中的元素不会超过 100
数组的大小不会超过 200
示例 1:

输入: [1, 5, 11, 5]

输出: true

解释: 数组可以分割成 [1, 5, 5] 和 [11].


示例 2:

输入: [1, 2, 3, 5]

输出: false

解释: 数组不能分割成两个元素和相等的子集.
*/
func canPartition(nums []int) bool {
	flag := false
	var canPartitionDFS func(sum int, nums []int, current, start int)
	canPartitionDFS = func(sum int, nums []int, current, start int) {
		if current > sum {
			return
		}
		if current == sum {
			flag = true
			return
		}
		if !flag {
			for i := start; i < len(nums); i++ {
				if current+nums[i] <= sum {
					current += nums[i]
					canPartitionDFS(sum, nums, current, i+1)
					current -= nums[i]
				}
			}
		}
	}
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	if sum%2 != 0 {
		return flag
	}

	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] > sum/2 {
			return false
		}
		if nums[i] < sum/2 {
			break
		}
	}
	canPartitionDFS(sum/2, nums, 0, 0)
	return flag
}
