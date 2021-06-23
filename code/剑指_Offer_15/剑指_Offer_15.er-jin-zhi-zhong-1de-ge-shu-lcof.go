package _剑指_Offer_15

func hammingWeight(num uint32) int {
	var ones int
	for i := 0; i < 32; i++ {
		if 1<<i&num > 0 {
			ones++
		}
	}
	return ones
}
