package _1803

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_countPairs(t *testing.T) {
	type args struct {
		nums []int
		low  int
		high int
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
			if got := countPairs(tt.args.nums, tt.args.low, tt.args.high); got != tt.want {
				t.Errorf("countPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
