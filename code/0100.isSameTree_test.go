package code

import (
	"testing"

	"github.com/godcong/leetcode"
)

func Test_isSameTree(t *testing.T) {
	type args struct {
		p *leetcode.TreeNode
		q *leetcode.TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{
				p: &leetcode.TreeNode{
					Val:   1,
					Left:  &leetcode.TreeNode{Val: 2},
					Right: &leetcode.TreeNode{Val: 3},
				},
				q: &leetcode.TreeNode{
					Val:   1,
					Left:  &leetcode.TreeNode{Val: 2},
					Right: &leetcode.TreeNode{Val: 3},
				},
			},
			want: true,
		},
		{
			name: "",
			args: args{
				p: &leetcode.TreeNode{
					Val:  1,
					Left: &leetcode.TreeNode{Val: 2},
				},
				q: &leetcode.TreeNode{
					Val:   1,
					Right: &leetcode.TreeNode{Val: 2},
				},
			},
			want: false,
		},
		{
			name: "",
			args: args{
				p: &leetcode.TreeNode{
					Val:   1,
					Left:  &leetcode.TreeNode{Val: 2},
					Right: &leetcode.TreeNode{Val: 1},
				},
				q: &leetcode.TreeNode{
					Val:   1,
					Left:  &leetcode.TreeNode{Val: 1},
					Right: &leetcode.TreeNode{Val: 2},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSameTree(tt.args.p, tt.args.q); got != tt.want {
				t.Errorf("isSameTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
