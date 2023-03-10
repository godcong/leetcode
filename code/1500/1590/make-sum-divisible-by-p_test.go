package _1590

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_minSubarray(t *testing.T) {
	type args struct {
		nums []int
		p    int
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
			if got := minSubarray(tt.args.nums, tt.args.p); got != tt.want {
				t.Errorf("minSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}
