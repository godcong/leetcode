package _3381

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_maxSubarraySum(t *testing.T) {
	type args struct {
		nums []int
		k    int
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
			if got := maxSubarraySum(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("maxSubarraySum() = %v, want %v", got, tt.want)
			}
		})
	}
}
