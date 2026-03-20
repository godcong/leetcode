package _0496

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	hash := make(map[int]int, len(nums2))
	var stack []int
	for _, num := range nums2 {
		for len(stack) > 0 && stack[len(stack)-1] <= num {
			stackTop := stack[len(stack)-1]
			hash[stackTop] = num
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, num)
	}

	res := make([]int, 0, len(nums1))
	for _, num := range nums1 {
		if value, ok := hash[num]; ok {
			res = append(res, value)
		} else {
			res = append(res, -1)
		}

	}

	return res
}
