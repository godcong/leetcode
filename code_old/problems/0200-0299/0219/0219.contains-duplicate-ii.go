package _0219

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int, 0)
	for idx, num := range nums {
		if get, ok := m[num]; ok {
			if idx-get <= k {
				return true
			} else {
				m[num] = idx
			}
		} else {
			m[num] = idx
		}
	}
	return false
}
