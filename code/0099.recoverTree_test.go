package code

import (
	"testing"

	"github.com/godcong/leetcode"
)

func Test_recoverTree(t *testing.T) {
	type args struct {
		root *leetcode.TreeNode
	}
	tests := []struct {
		name string
		args args
		want *leetcode.TreeNode
	}{
		{
			name: "",
			args: args{
				root: leetcode.strToTreeNode("[1,3,null,null,2]"),
			},
			want: leetcode.strToTreeNode("[3,1,null,null,2]"),
		},
		{
			name: "",
			args: args{
				root: leetcode.strToTreeNode("[3,1,4,null,null,2]"),
			},
			want: leetcode.strToTreeNode("[2,1,4,null,null,3]"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if recoverTree(tt.args.root); !leetcode.treeNodeDeepEqual(t, tt.args.root, tt.want) {
				t.Errorf("recoverTree() = %v, want %v", tt.args.root, tt.want)
			}
		})
	}
}
