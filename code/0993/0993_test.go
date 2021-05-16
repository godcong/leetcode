package _0993

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_isCousins(t *testing.T) {
	type args struct {
		root *TreeNode
		x    int
		y    int
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
				root: StrToTreeNode("[1,2,3,4]"),
				x:    4,
				y:    3,
			},
			want: false,
		},
		{
			name: "",
			args: args{
				root: StrToTreeNode("[1,2,3,null,4,null,5]"),
				x:    5,
				y:    4,
			},
			want: true,
		},
		{
			name: "",
			args: args{
				root: StrToTreeNode("[1,2,3,null,4]"),
				x:    2,
				y:    3,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isCousins(tt.args.root, tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("isCousins() = %v, want %v", got, tt.want)
			}
		})
	}
}
