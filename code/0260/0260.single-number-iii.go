package _0260

func singleNumber(nums []int) []int {
	var a = 0
	for _, v := range nums {
		a ^= v
	}
	var t = 1
	for (t & a) == 0 {
		t <<= 1
	}
	var res1, res2 int
	for _, v := range nums {
		if (v & t) > 0 {
			res1 ^= v
		} else {
			res2 ^= v
		}
	}
	return []int{res1, res2}
}
