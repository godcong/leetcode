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
		{
			name: "",
			args: args{
				root1: StrToTreeNode("[3,5,1,6,2,9,8,null,null,7,4]"),
				root2: StrToTreeNode("[3,5,1,6,7,4,2,null,null,null,null,null,null,9,8]"),
			},
			want: true,
		},
		{
			name: "",
			args: args{
				root1: StrToTreeNode("[1]"),
				root2: StrToTreeNode("[1]"),
			},
			want: true,
		},
		{
			name: "",
			args: args{
				root1: StrToTreeNode("[1]"),
				root2: StrToTreeNode("[2]"),
			},
			want: false,
		},
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
