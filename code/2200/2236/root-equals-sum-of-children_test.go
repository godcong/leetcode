package _2236

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_checkTree(t *testing.T) {
	type args struct {
		root *TreeNode
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
			if got := checkTree(tt.args.root); got != tt.want {
				t.Errorf("checkTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
