package _2617

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_minimumVisitedCells(t *testing.T) {
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
			if got := minimumVisitedCells(tt.args.grid); got != tt.want {
				t.Errorf("minimumVisitedCells() = %v, want %v", got, tt.want)
			}
		})
	}
}
