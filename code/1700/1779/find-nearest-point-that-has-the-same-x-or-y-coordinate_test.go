package _1779

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_nearestValidPoint(t *testing.T) {
	type args struct {
		x      int
		y      int
		points [][]int
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
			if got := nearestValidPoint(tt.args.x, tt.args.y, tt.args.points); got != tt.want {
				t.Errorf("nearestValidPoint() = %v, want %v", got, tt.want)
			}
		})
	}
}
