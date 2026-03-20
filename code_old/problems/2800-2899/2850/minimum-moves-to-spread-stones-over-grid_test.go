package _2850

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_minimumMoves(t *testing.T) {
	type args struct {
		grid [][]int
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
			if got := minimumMoves(tt.args.grid); got != tt.want {
				t.Errorf("minimumMoves() = %v, want %v", got, tt.want)
			}
		})
	}
}
