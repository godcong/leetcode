package _0795

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_numSubarrayBoundedMax(t *testing.T) {
	type args struct {
		nums  []int
		left  int
		right int
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
			if got := numSubarrayBoundedMax(tt.args.nums, tt.args.left, tt.args.right); got != tt.want {
				t.Errorf("numSubarrayBoundedMax() = %v, want %v", got, tt.want)
			}
		})
	}
}
