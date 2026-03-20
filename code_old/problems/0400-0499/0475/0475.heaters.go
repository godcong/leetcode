package _0475

import (
	"math"
	"sort"
)

func findRadius(houses []int, heaters []int) int {
	sort.Ints(heaters)
	ans := 0
	for _, house := range houses {
		j := sort.SearchInts(heaters, house+1)
		minDis := math.MaxInt32
		if j < len(heaters) {
			minDis = heaters[j] - house
		}
		i := j - 1
		if i >= 0 && house-heaters[i] < minDis {
			minDis = house - heaters[i]
		}
		if minDis > ans {
			ans = minDis
		}
	}

	return ans
}
