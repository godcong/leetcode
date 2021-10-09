package _0352

import (
	"sort"
)

type SummaryRanges struct {
	Intervals [][]int
}

func Constructor() SummaryRanges {
	return SummaryRanges{
		Intervals: [][]int{
			{-5, -5},
			{10005, 10005},
		},
	}
}

func (this *SummaryRanges) AddNum(val int) {
	i := sort.Search(len(this.Intervals), func(i int) bool {
		return this.Intervals[i][0] >= val
	})
	if this.Intervals[i][0] == val {
		return
	}
	i--
	if this.Intervals[i][1] >= val {
		return
	} else if this.Intervals[i][1]+1 == val {
		if this.Intervals[i+1][0]-1 == val {
			this.Intervals[i+1][0] = this.Intervals[i][0]
			this.Intervals = append(this.Intervals[:i], this.Intervals[i+1:]...)
		} else {
			this.Intervals[i][1] = val
		}
	} else {
		if i < len(this.Intervals)-1 && this.Intervals[i+1][0]-1 == val {
			this.Intervals[i+1][0] = val
		} else {
			this.Intervals = append(this.Intervals[:i+1], this.Intervals[i:]...)
			this.Intervals[i+1] = []int{val, val}
		}
	}
}

func (this *SummaryRanges) GetIntervals() [][]int {
	return this.Intervals[1 : len(this.Intervals)-1]
}

/**
 * Your SummaryRanges object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(val);
 * param_2 := obj.GetIntervals();
 */
