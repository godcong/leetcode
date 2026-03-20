package _0313

import (
	"math"
)

func nthSuperUglyNumber(n int, primes []int) int {
	pIndex := make([]int, len(primes))
	dp := make([]int, n)
	dp[0] = 1

	for i := 1; i < n; i++ {
		minVal := math.MaxInt64
		for j := range pIndex {
			if dp[pIndex[j]]*primes[j] < minVal {
				minVal = dp[pIndex[j]] * primes[j]
			}
		}

		for j := range pIndex {
			if dp[pIndex[j]]*primes[j] == minVal {
				pIndex[j]++
			}
		}

		dp[i] = minVal
	}

	return dp[n-1]
}
