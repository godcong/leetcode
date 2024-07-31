package _3111

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_minRectanglesToCoverPoints(t *testing.T) {
	type args struct {
		points [][]int
		w      int
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
			if got := minRectanglesToCoverPoints(tt.args.points, tt.args.w); got != tt.want {
				t.Errorf("minRectanglesToCoverPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
