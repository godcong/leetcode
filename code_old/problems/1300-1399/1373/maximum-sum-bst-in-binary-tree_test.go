package _1373

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_maxSumBST(t *testing.T) {
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
			if got := maxSumBST(tt.args.root); got != tt.want {
				t.Errorf("maxSumBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
