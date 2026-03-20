package _0687

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_longestUnivaluePath(t *testing.T) {
	type args struct {
		root *TreeNode
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
			if got := longestUnivaluePath(tt.args.root); got != tt.want {
				t.Errorf("longestUnivaluePath() = %v, want %v", got, tt.want)
			}
		})
	}
}
