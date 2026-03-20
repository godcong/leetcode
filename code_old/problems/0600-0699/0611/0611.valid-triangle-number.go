package _0611

import (
	"sort"
)

func triangleNumber(nums []int) int {
	sort.Sort(sort.IntSlice(nums))
	res := 0
	for i := len(nums) - 1; i >= 2; i-- {
		left := 0
		right := i - 1
		for left < right {
			if nums[left]+nums[right] > nums[i] {
				// left .... right
				res += right - left
				right--
			} else {
				left++
			}
		}
	}
	return res
}
