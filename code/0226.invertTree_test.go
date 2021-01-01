package code

import (
	"reflect"
	"testing"

	"github.com/godcong/leetcode"
)

func Test_invertTree(t *testing.T) {
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
				root: leetcode.strToTreeNode("[4,2,7,1,3,6,9]"),
			},
			want: &leetcode.TreeNode{
				Val: 4,
				Left: &leetcode.TreeNode{
					Val:   7,
					Left:  &leetcode.TreeNode{Val: 9},
					Right: &leetcode.TreeNode{Val: 6},
				},
				Right: &leetcode.TreeNode{
					Val:   2,
					Left:  &leetcode.TreeNode{Val: 3},
					Right: &leetcode.TreeNode{Val: 1},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := invertTree(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("invertTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
