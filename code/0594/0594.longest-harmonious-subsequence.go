package _0594

func findLHS(nums []int) int {
	m := make(map[int]int)
	for _, num := range nums {
		m[num] += 1
	}

	var res int
	for key, value := range m {
		if nextValue, ok := m[key+1]; ok {
			res = max(res, value+nextValue)
		}
	}
	return res
}

func max(res int, i int) int {
	if res < i {
		return i
	}
	return res
}
