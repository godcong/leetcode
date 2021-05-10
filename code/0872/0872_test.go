package _0872

import (
	. "github.com/godcong/leetcode/common"
	"testing"
)

func Test_leafSimilar(t *testing.T) {
	type args struct {
		root1 *TreeNode
		root2 *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := leafSimilar(tt.args.root1, tt.args.root2); got != tt.want {
					t.Errorf("leafSimilar() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
