package leetcode

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
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val:  3,
						Left: nil,
						Right: &TreeNode{
							Val:   2,
							Left:  nil,
							Right: nil,
						},
					},
					Right: nil,
				},
			},
			want: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val:  1,
					Left: nil,
					Right: &TreeNode{
						Val:   2,
						Left:  nil,
						Right: nil,
					},
				},
				Right: nil,
			},
		},
		{
			name: "",
			args: args{
				root: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val:   1,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val:   2,
							Left:  nil,
							Right: nil,
						},
						Right: nil,
					},
				},
			},
			want: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val:   1,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val:   3,
						Left:  nil,
						Right: nil,
					},
					Right: nil,
				},
			},
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
