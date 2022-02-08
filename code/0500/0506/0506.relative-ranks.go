package _0506

import (
	"sort"
	"strconv"
)

var desc = [3]string{"Gold Medal", "Silver Medal", "Bronze Medal"}

func findRelativeRanks(score []int) []string {
	const (
		SHIFT = 20
		MASK  = (1 << SHIFT) - 1
	)

	for i, v := range score {
		score[i] = (v << SHIFT) | i
	}

	sort.Ints(score)

	var ret = make([]string, len(score))

	var all = len(score) - 1
	for r, v := range score {
		ret[v&MASK] = getMedal(all - r)
	}
	return ret
}

func getMedal(rank int) string {
	if rank >= 3 {
		return strconv.Itoa(rank + 1)
	}
	return desc[rank]
}
