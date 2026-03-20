package _2304

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_minPathCost(t *testing.T) {
	type args struct {
		grid     [][]int
		moveCost [][]int
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
			if got := minPathCost(tt.args.grid, tt.args.moveCost); got != tt.want {
				t.Errorf("minPathCost() = %v, want %v", got, tt.want)
			}
		})
	}
}
