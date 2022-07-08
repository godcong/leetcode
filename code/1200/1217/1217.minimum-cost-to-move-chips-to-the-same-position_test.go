package _1217

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_minCostToMoveChips(t *testing.T) {
	type args struct {
		position []int
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
			if got := minCostToMoveChips(tt.args.position); got != tt.want {
				t.Errorf("minCostToMoveChips() = %v, want %v", got, tt.want)
			}
		})
	}
}
