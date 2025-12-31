package _1970

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_latestDayToCross(t *testing.T) {
	type args struct {
		row   int
		col   int
		cells [][]int
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
			if got := latestDayToCross(tt.args.row, tt.args.col, tt.args.cells); got != tt.want {
				t.Errorf("latestDayToCross() = %v, want %v", got, tt.want)
			}
		})
	}
}
