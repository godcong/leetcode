package code

import (
	"testing"

	"github.com/godcong/leetcode"
)

func Test_sumNumbers(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				root: leetcode.strToTreeNode("[1,2,3]"),
			},
			want: 25,
		},
		{
			name: "",
			args: args{
				root: leetcode.strToTreeNode("[4,9,0,5,1]"),
			},
			want: 1026,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumNumbers(tt.args.root); got != tt.want {
				t.Errorf("sumNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
