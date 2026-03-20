package _0268

func missingNumber(nums []int) int {
	var ret int
	for i, num := range nums {
		ret ^= i ^ num
	}
	return ret ^ len(nums)
}
