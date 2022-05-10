package _1728

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_canMouseWin(t *testing.T) {
	type args struct {
		grid      []string
		catJump   int
		mouseJump int
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
			if got := canMouseWin(tt.args.grid, tt.args.catJump, tt.args.mouseJump); got != tt.want {
				t.Errorf("canMouseWin() = %v, want %v", got, tt.want)
			}
		})
	}
}
