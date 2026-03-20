package _3349

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_hasIncreasingSubarrays(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasIncreasingSubarrays(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("hasIncreasingSubarrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
