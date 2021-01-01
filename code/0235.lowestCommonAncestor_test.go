package code

import (
	"testing"

	"github.com/godcong/leetcode"
)

func Test_lowestCommonAncestor(t *testing.T) {
	type args struct {
		root *leetcode.TreeNode
		p    *leetcode.TreeNode
		q    *leetcode.TreeNode
	}
	var tests = []struct {
		name string
		args args
		want *leetcode.TreeNode
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				root: leetcode.strToTreeNode("[6,2,8,0,4,7,9,null,null,3,5]"),
				p:    &leetcode.TreeNode{Val: 2},
				q:    &leetcode.TreeNode{Val: 8},
			},
			want: &leetcode.TreeNode{Val: 6},
		},
		{
			name: "",
			args: args{
				root: leetcode.strToTreeNode("[6,2,8,0,4,7,9,null,null,3,5]"),
				p:    &leetcode.TreeNode{Val: 2},
				q:    &leetcode.TreeNode{Val: 4},
			},
			want: &leetcode.TreeNode{Val: 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lowestCommonAncestor(tt.args.root, tt.args.p, tt.args.q); got.Val != tt.want.Val {
				t.Errorf("lowestCommonAncestor() = %v, want %v", got, tt.want)
			}
		})
	}
}
