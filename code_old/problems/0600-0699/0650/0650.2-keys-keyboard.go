package _0650

import (
	"math"
)

func minSteps(n int) int {
	f := make([]int, n+1)
	for i := 2; i <= n; i++ {
		f[i] = math.MaxInt32
		for j := 1; j*j <= i; j++ {
			if i%j == 0 {
				f[i] = min(f[i], f[j]+i/j)
				f[i] = min(f[i], f[i/j]+j)
			}
		}
	}
	return f[n]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
