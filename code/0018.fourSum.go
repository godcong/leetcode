package code

import "sort"

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
	count := len(nums)
	sort.Ints(nums)
	var result [][]int
	for i := 0; i < count-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		if nums[i]+nums[i+1]+nums[i+2]+nums[i+3] > target {
			break
		}
		if nums[i]+nums[count-3]+nums[count-2]+nums[count-1] < target {
			continue
		}
		for j := i + 1; j < count-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			if nums[i]+nums[j]+nums[j+1]+nums[j+2] > target {
				break
			}
			if nums[i]+nums[j]+nums[count-2]+nums[count-1] < target {
				continue
			}
			left, right := j+1, count-1
			for left < right {
				if left > j+1 && nums[left] == nums[left-1] {
					left++
					continue
				}
				if right < count-1 && nums[right] == nums[right+1] {
					right--
					continue
				}
				sum := nums[i] + nums[j] + nums[left] + nums[right]
				if sum == target {
					result = append(result, append([]int{}, nums[i], nums[j], nums[left], nums[right]))
					left++
					right--
				} else if sum < target {
					left++
				} else if sum > target {
					right--
				}
			}
		}
	}
	return result
}
