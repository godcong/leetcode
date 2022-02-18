package _2006

func countKDifference(nums []int, k int) int {
	cnt := map[int]int{}
	c := 0

	for _, v := range nums {
		c += cnt[v+k] + cnt[v-k]
		cnt[v]++
	}
	return c
}
