package _1145

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_btreeGameWinningMove(t *testing.T) {
	type args struct {
		root *TreeNode
		n    int
		x    int
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
			if got := btreeGameWinningMove(tt.args.root, tt.args.n, tt.args.x); got != tt.want {
				t.Errorf("btreeGameWinningMove() = %v, want %v", got, tt.want)
			}
		})
	}
}
