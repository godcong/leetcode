package _0930

func numSubarraysWithSum(nums []int, goal int) int {
	cnt := map[int]int{}
	ans := 0
	sum := 0
	for _, num := range nums {
		cnt[sum]++
		sum += num
		ans += cnt[sum-goal]
	}
	return ans
}
