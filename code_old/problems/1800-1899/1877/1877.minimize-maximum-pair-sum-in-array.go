package _1877

import (
	"runtime/debug"
)

const NumberMax int = 1e5

func minPairSum(nums []int) int {
	cnt := [NumberMax + 1]int{}
	for _, v := range nums {
		cnt[v]++
	}

	ans := 2
	i, j := 1, NumberMax
	for {
		for ; cnt[i] == 0 && i <= j; i++ {
		}
		for ; cnt[j] == 0 && j >= i; j-- {
		}
		if i >= j {
			if i == j && cnt[i] > 1 {
				ans = max(ans, i*2)
			}
			break
		}
		ans = max(ans, i+j)
		cnt[i]--
		cnt[j]--
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func init() { debug.SetGCPercent(-1) }
