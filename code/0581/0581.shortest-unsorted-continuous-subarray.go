package _0581

import (
	"math"
)

func findUnsortedSubarray(nums []int) int {
	curMin := math.MaxInt32
	curMax := math.MinInt32
	m := len(nums)
	for i := 1; i < m; i++ {
		if nums[i] < nums[i-1] {
			if nums[i] < curMin {
				curMin = nums[i]
			}
			if nums[i-1] > curMax {
				curMax = nums[i-1]
			}
		}
	}
	if curMax == math.MinInt32 && curMin == math.MaxInt32 {
		return 0
	}
	l, r := 0, m-1
	for ; l < m && nums[l] <= curMin; l++ {
	}
	for ; r >= 0 && nums[r] >= curMax; r-- {
	}

	return r - l + 1
}
