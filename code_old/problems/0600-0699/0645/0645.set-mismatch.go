package _0645

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func findErrorNums(nums []int) []int {
	var ret [2]int
	for _, v := range nums {
		var idx = abs(v)
		if nums[idx-1] < 0 {
			ret[0] = idx
		} else {
			nums[idx-1] *= -1
		}
	}

	for i, v := range nums {
		if v > 0 {
			ret[1] = i + 1
			break
		}
	}

	return ret[:]
}
