package _0911

import (
	"sort"
)

type TopVotedCandidate struct {
	tops, times []int
}

func Constructor(persons []int, times []int) TopVotedCandidate {
	tops := []int{}
	top := -1
	voteCounts := map[int]int{-1: -1}
	for _, p := range persons {
		voteCounts[p]++
		if voteCounts[p] >= voteCounts[top] {
			top = p
		}
		tops = append(tops, top)
	}
	return TopVotedCandidate{tops, times}

}

func (this *TopVotedCandidate) Q(t int) int {
	return this.tops[sort.SearchInts(this.times, t+1)-1]
}

/**
 * Your TopVotedCandidate object will be instantiated and called as such:
 * obj := Constructor(persons, times);
 * param_1 := obj.Q(t);
 */
