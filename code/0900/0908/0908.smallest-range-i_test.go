package _0908

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_smallestRangeI(t *testing.T) {
	type args struct {
		nums []int
		k    int
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
			if got := smallestRangeI(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("smallestRangeI() = %v, want %v", got, tt.want)
			}
		})
	}
}
