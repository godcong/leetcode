package _3086

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_minimumMoves(t *testing.T) {
	type args struct {
		nums       []int
		k          int
		maxChanges int
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
			if got := minimumMoves(tt.args.nums, tt.args.k, tt.args.maxChanges); got != tt.want {
				t.Errorf("minimumMoves() = %v, want %v", got, tt.want)
			}
		})
	}
}
