package _3013

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_minimumCost(t *testing.T) {
	type args struct {
		nums []int
		k    int
		dist int
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
			if got := minimumCost(tt.args.nums, tt.args.k, tt.args.dist); got != tt.want {
				t.Errorf("minimumCost() = %v, want %v", got, tt.want)
			}
		})
	}
}
