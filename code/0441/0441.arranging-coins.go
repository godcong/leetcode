package _0441

import (
	"sort"
)

func arrangeCoins(n int) int {
	return sort.Search(n, func(k int) bool { k++; return k*(k+1) > 2*n })
}
