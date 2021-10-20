package _0453

func minMoves(nums []int) int {
	var min = nums[0]
	var sum = 0
	for i := 0; i < len(nums); i++ {
		if nums[i] < min {
			min = nums[i]
		}
		sum += nums[i]
	}
	return sum - len(nums)*min
}
