package _0930

func numSubarraysWithSum(nums []int, goal int) int {
	freq, sum, res := make([]int, len(nums)+1), 0, 0
	freq[0] = 1
	for _, v := range nums {
		t := sum + v - goal
		if t >= 0 {
			res += freq[t]
		}
		sum += v
		freq[sum]++
	}
	return res
}
