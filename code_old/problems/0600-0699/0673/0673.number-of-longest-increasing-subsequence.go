package _0673

import (
	"sort"
)

func findNumberOfLIS(nums []int) int {
	var d [][]int
	var cnt [][]int
	for _, v := range nums {
		i := sort.Search(len(d), func(i int) bool { return d[i][len(d[i])-1] >= v })
		c := 1
		if i > 0 {
			k := sort.Search(len(d[i-1]), func(k int) bool { return d[i-1][k] < v })
			c = cnt[i-1][len(cnt[i-1])-1] - cnt[i-1][k]
		}
		if i == len(d) {
			d = append(d, []int{v})
			cnt = append(cnt, []int{0, c})
		} else {
			d[i] = append(d[i], v)
			cnt[i] = append(cnt[i], cnt[i][len(cnt[i])-1]+c)
		}
	}
	c := cnt[len(cnt)-1]
	return c[len(c)-1]

}
