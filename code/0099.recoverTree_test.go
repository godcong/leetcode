package code

import (
	"testing"
)

func Test_recoverTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "",
			args: args{
				root: strToTreeNode("[1,3,null,null,2]"),
			},
			want: strToTreeNode("[3,1,null,null,2]"),
		},
		{
			name: "",
			args: args{
				root: strToTreeNode("[3,1,4,null,null,2]"),
			},
			want: strToTreeNode("[2,1,4,null,null,3]"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if recoverTree(tt.args.root); !treeNodeDeepEqual(t, tt.args.root, tt.want) {
				t.Errorf("recoverTree() = %v, want %v", tt.args.root, tt.want)
			}
		})
	}
}
