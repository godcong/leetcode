package _0525

func findMaxLength(nums []int) int {
	m := make(map[int]int, 0)
	m[0] = -1
	maxLen, count := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			count++
		} else {
			count--
		}
		if v, exist := m[count]; exist {
			maxLen = max(maxLen, i-v)
		} else {
			m[count] = i
		}
	}
	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
