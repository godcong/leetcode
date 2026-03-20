package _2596

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_checkValidGrid(t *testing.T) {
	type args struct {
		grid [][]int
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
			if got := checkValidGrid(tt.args.grid); got != tt.want {
				t.Errorf("checkValidGrid() = %v, want %v", got, tt.want)
			}
		})
	}
}
