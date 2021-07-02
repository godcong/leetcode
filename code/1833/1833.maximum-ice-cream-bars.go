package _1833

import (
	"runtime/debug"
)

const max int = 1e5

func maxIceCream(costs []int, coins int) int {
	ans := 0
	freq := [max + 1]int{}
	for _, c := range costs {
		freq[c]++
	}
	for i := 1; i <= max && coins >= i; i++ {
		if freq[i] == 0 {
			continue
		}
		c := min(freq[i], coins/i)
		ans += c
		coins -= i * c
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func init() { debug.SetGCPercent(-1) }
