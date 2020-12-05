package leetcode

import (
	"testing"
)

func Test_convertBST(t *testing.T) {
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
					Val:   5,
					Left:  &TreeNode{Val: 2},
					Right: &TreeNode{Val: 13},
				},
			},
			want: &TreeNode{
				Val:   18,
				Left:  &TreeNode{Val: 20},
				Right: &TreeNode{Val: 13},
			},
		},
		{
			name: "",
			args: args{
				root: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 0,
						Left: &TreeNode{
							Val: -4,
						},
						Right: &TreeNode{
							Val: 1,
						},
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
			},
			want: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 6,
					Left: &TreeNode{
						Val: 2,
					},
					Right: &TreeNode{
						Val: 6,
					},
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertBST(tt.args.root); !treeNodeDeepEqual(t, got, tt.want) {
				t.Errorf("convertBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
