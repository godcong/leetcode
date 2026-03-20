package _2859

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_sumIndicesWithKSetBits(t *testing.T) {
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
			if got := sumIndicesWithKSetBits(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("sumIndicesWithKSetBits() = %v, want %v", got, tt.want)
			}
		})
	}
}
