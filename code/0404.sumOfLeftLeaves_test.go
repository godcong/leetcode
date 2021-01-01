package code

import (
	"testing"

	"github.com/godcong/leetcode"
)

func Test_sumOfLeftLeaves(t *testing.T) {
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
				root: &leetcode.TreeNode{
					Val: 3,
					Left: &leetcode.TreeNode{
						Val: 9,
					},
					Right: &leetcode.TreeNode{
						Val:   20,
						Left:  &leetcode.TreeNode{Val: 15},
						Right: &leetcode.TreeNode{Val: 7},
					},
				},
			},
			want: 24,
		},
		{
			name: "",
			args: args{
				root: &leetcode.TreeNode{
					Val: 1,
					Left: &leetcode.TreeNode{
						Val: 2,
						Left: &leetcode.TreeNode{
							Val: 4,
						},
						Right: &leetcode.TreeNode{
							Val: 5,
						},
					},
					Right: &leetcode.TreeNode{
						Val: 3,
					},
				},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumOfLeftLeaves(tt.args.root); got != tt.want {
				t.Errorf("sumOfLeftLeaves() = %v, want %v", got, tt.want)
			}
		})
	}
}
