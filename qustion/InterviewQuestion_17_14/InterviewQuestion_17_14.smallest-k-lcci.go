package _InterviewQuestion_17_14

import (
	"sort"
)

func smallestK(arr []int, k int) []int {
	sort.Ints(arr)
	return arr[:k]
}
