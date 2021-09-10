package _1894

import (
	"sort"
)

func chalkReplacer(chalk []int, k int) int {
	if chalk[0] > k {
		return 0
	}
	n := len(chalk)
	for i := 1; i < n; i++ {
		chalk[i] += chalk[i-1]
		if chalk[i] > k {
			return i
		}
	}
	k %= chalk[n-1]
	return sort.SearchInts(chalk, k+1)
}
