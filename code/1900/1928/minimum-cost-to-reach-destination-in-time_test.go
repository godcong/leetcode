package _1928

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_minCost(t *testing.T) {
	type args struct {
		maxTime     int
		edges       [][]int
		passingFees []int
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
			if got := minCost(tt.args.maxTime, tt.args.edges, tt.args.passingFees); got != tt.want {
				t.Errorf("minCost() = %v, want %v", got, tt.want)
			}
		})
	}
}
