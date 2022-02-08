package _1995

func countQuadruplets(nums []int) int {
	cnt := map[int]int{}
	ans := 0
	for i := 1; i < len(nums)-2; i++ {
		for j := 0; j < i; j++ {
			cnt[nums[j]+nums[i]]++
		}
		for j := i + 2; j < len(nums); j++ {
			ans += cnt[nums[j]-nums[i+1]]
		}
	}

	return ans
}
