package code

import (
	"testing"

	"github.com/godcong/leetcode"
)

func Test_convertBST(t *testing.T) {
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
				root: &leetcode.TreeNode{
					Val:   5,
					Left:  &leetcode.TreeNode{Val: 2},
					Right: &leetcode.TreeNode{Val: 13},
				},
			},
			want: &leetcode.TreeNode{
				Val:   18,
				Left:  &leetcode.TreeNode{Val: 20},
				Right: &leetcode.TreeNode{Val: 13},
			},
		},
		{
			name: "",
			args: args{
				root: &leetcode.TreeNode{
					Val: 2,
					Left: &leetcode.TreeNode{
						Val: 0,
						Left: &leetcode.TreeNode{
							Val: -4,
						},
						Right: &leetcode.TreeNode{
							Val: 1,
						},
					},
					Right: &leetcode.TreeNode{
						Val: 3,
					},
				},
			},
			want: &leetcode.TreeNode{
				Val: 5,
				Left: &leetcode.TreeNode{
					Val: 6,
					Left: &leetcode.TreeNode{
						Val: 2,
					},
					Right: &leetcode.TreeNode{
						Val: 6,
					},
				},
				Right: &leetcode.TreeNode{
					Val: 3,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertBST(tt.args.root); !leetcode.treeNodeDeepEqual(t, got, tt.want) {
				t.Errorf("convertBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
