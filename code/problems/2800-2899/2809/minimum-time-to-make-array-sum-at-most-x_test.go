package _2809

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_minimumTime(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
		x     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumTime(tt.args.nums1, tt.args.nums2, tt.args.x); got != tt.want {
				t.Errorf("minimumTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
