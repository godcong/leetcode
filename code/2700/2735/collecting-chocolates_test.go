package _2735

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_minCost(t *testing.T) {
	type args struct {
		nums []int
		x    int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCost(tt.args.nums, tt.args.x); got != tt.want {
				t.Errorf("minCost() = %v, want %v", got, tt.want)
			}
		})
	}
}
