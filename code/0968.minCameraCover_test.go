package code

import (
	"testing"

	"github.com/godcong/leetcode"
)

func Test_minCameraCover(t *testing.T) {
	type args struct {
		root *leetcode.TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				root: leetcode.strToTreeNode("[0,0,null,0,0]"),
			},
			want: 1,
		},
		{
			name: "",
			args: args{
				root: leetcode.strToTreeNode("[0,0,null,0,null,0,null,null,0]"),
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCameraCover(tt.args.root); got != tt.want {
				t.Errorf("minCameraCover() = %v, want %v", got, tt.want)
			}
		})
	}
}
