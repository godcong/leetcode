package _1031

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_maxSumTwoNoOverlap(t *testing.T) {
	type args struct {
		nums      []int
		firstLen  int
		secondLen int
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
			if got := maxSumTwoNoOverlap(tt.args.nums, tt.args.firstLen, tt.args.secondLen); got != tt.want {
				t.Errorf("maxSumTwoNoOverlap() = %v, want %v", got, tt.want)
			}
		})
	}
}
