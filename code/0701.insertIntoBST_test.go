package code

import (
	"reflect"
	"testing"

	"github.com/godcong/leetcode"
)

func Test_insertIntoBST(t *testing.T) {
	type args struct {
		root *leetcode.TreeNode
		val  int
	}
	tests := []struct {
		name string
		args args
		want *leetcode.TreeNode
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				root: &leetcode.TreeNode{
					Val: 4,
					Left: &leetcode.TreeNode{
						Val: 2,
						Left: &leetcode.TreeNode{
							Val:   1,
							Left:  nil,
							Right: nil,
						},
						Right: &leetcode.TreeNode{
							Val:   3,
							Left:  nil,
							Right: nil,
						},
					},
					Right: &leetcode.TreeNode{
						Val:   7,
						Left:  nil,
						Right: nil,
					},
				},
				val: 5,
			},
			want: &leetcode.TreeNode{
				Val: 4,
				Left: &leetcode.TreeNode{
					Val: 2,
					Left: &leetcode.TreeNode{
						Val:   1,
						Left:  nil,
						Right: nil,
					},
					Right: &leetcode.TreeNode{
						Val:   3,
						Left:  nil,
						Right: nil,
					},
				},
				Right: &leetcode.TreeNode{
					Val: 7,
					Left: &leetcode.TreeNode{
						Val:   5,
						Left:  nil,
						Right: nil,
					},
					Right: nil,
				},
			},
		},
		{
			name: "",
			args: args{
				root: leetcode.strToTreeNode(""),
				val:  5,
			},
			want: &leetcode.TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insertIntoBST(tt.args.root, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insertIntoBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
