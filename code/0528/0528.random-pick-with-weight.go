package _0528

import (
	"sort"
)

type Solution struct {
	pre   []int
	total uint
	rd    uint
}

func Constructor(w []int) Solution {
	for i := 1; i < len(w); i++ {
		w[i] += w[i-1]
	}
	return Solution{
		pre:   w,
		total: uint(w[len(w)-1]),
		rd:    1,
	}
}

func (this *Solution) PickIndex() int {
	this.rd ^= this.rd << 13
	this.rd ^= this.rd >> 17
	this.rd ^= this.rd << 5
	return sort.SearchInts(this.pre, int(this.rd%this.total)+1)
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */
