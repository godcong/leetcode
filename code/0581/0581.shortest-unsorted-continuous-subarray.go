package _0581

import (
	"sort"
)

func findUnsortedSubarray(nums []int) int {
	if sort.IntsAreSorted(nums) {
		return 0
	}
	numsSorted := append([]int(nil), nums...)
	sort.Ints(numsSorted)
	left, right := 0, len(nums)-1
	for nums[left] == numsSorted[left] {
		left++
	}
	for nums[right] == numsSorted[right] {
		right--
	}
	return right - left + 1
}
