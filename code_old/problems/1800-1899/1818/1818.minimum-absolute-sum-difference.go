package _1818

import (
	"runtime/debug"
)

func minAbsoluteSumDiff(nums1 []int, nums2 []int) int {
	sum := 0
	m := 0
	for i, n2 := range nums2 {
		d := abs(nums1[i] - n2)
		sum += d
		if m < d {
			t := 100001
			for _, n1 := range nums1 {
				t = min(t, abs(n1-n2))
			}
			m = max(m, d-t)
		}
	}
	return (sum - m) % (1e9 + 7)
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func init() { debug.SetGCPercent(-1) }
