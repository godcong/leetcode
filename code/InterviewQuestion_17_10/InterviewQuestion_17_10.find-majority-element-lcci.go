package _InterviewQuestion_17_10

func majorityElement(nums []int) int {
	candidate := -1
	count := 0
	for _, num := range nums {
		if count == 0 {
			candidate = num
		}
		if num == candidate {
			count++
		} else {
			count--
		}
	}
	count = 0
	for _, num := range nums {
		if num == candidate {
			count++
		}
	}
	if count*2 > len(nums) {
		return candidate
	}
	return -1
}
