package _0413

func numberOfArithmeticSlices(nums []int) int {
	n := len(nums)
	ans := 0
	if n == 1 {
		return ans
	}

	d, t := nums[0]-nums[1], 0
	for i := 2; i < n; i++ {
		if nums[i-1]-nums[i] == d {
			t++
		} else {
			d, t = nums[i-1]-nums[i], 0
		}
		ans += t
	}
	return ans
}
