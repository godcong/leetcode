package _1711

import (
	"math/bits"
	"sort"
)

func countPairs(deliciousness []int) int {
	num := len(deliciousness)
	sort.Ints(deliciousness)

	numMap := map[int]int{}
	prevCnt := 0

	numMap[deliciousness[0]] = 1
	for i := 1; i < num; i++ {
		n := deliciousness[i]
		n0 := 1<<(32-bits.LeadingZeros32(uint32(n))) - n
		cnt := numMap[n0]

		if deliciousness[0] == 0 && (n > 0 && (n&(n-1)) == 0) {
			cnt += numMap[0]
		}
		prevCnt = (prevCnt + cnt) % (1e9 + 7)
		numMap[n] += 1
	}
	return prevCnt
}
