package _3235

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_canReachCorner(t *testing.T) {
	type args struct {
		xCorner int
		yCorner int
		circles [][]int
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
			if got := canReachCorner(tt.args.xCorner, tt.args.yCorner, tt.args.circles); got != tt.want {
				t.Errorf("canReachCorner() = %v, want %v", got, tt.want)
			}
		})
	}
}
