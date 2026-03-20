package _0540

func singleNonDuplicate(nums []int) int {
	low, high := 0, len(nums)-1
	for low < high {
		mid := low + (high-low)/2
		if nums[mid] == nums[mid^1] {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return nums[low]
}
