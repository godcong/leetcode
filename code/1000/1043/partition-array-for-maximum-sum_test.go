package _1043

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_maxSumAfterPartitioning(t *testing.T) {
	type args struct {
		arr []int
		k   int
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
			if got := maxSumAfterPartitioning(tt.args.arr, tt.args.k); got != tt.want {
				t.Errorf("maxSumAfterPartitioning() = %v, want %v", got, tt.want)
			}
		})
	}
}
