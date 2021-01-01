package code

import (
	"testing"

	"github.com/godcong/leetcode"
)

func Test_minDepth(t *testing.T) {
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
						Val:   9,
						Left:  nil,
						Right: nil,
					},
					Right: &leetcode.TreeNode{
						Val: 20,
						Left: &leetcode.TreeNode{
							Val:   15,
							Left:  nil,
							Right: nil,
						},
						Right: &leetcode.TreeNode{
							Val:   7,
							Left:  nil,
							Right: nil,
						},
					},
				},
			},
			want: 2,
		},
		{
			name: "",
			args: args{
				root: &leetcode.TreeNode{
					Val: 1,
					Left: &leetcode.TreeNode{
						Val:   2,
						Left:  nil,
						Right: nil,
					},
					Right: nil,
				},
			},
			want: 2,
		},
		{
			name: "",
			args: args{
				root: &leetcode.TreeNode{
					Val: 1,
					Left: &leetcode.TreeNode{
						Val:  2,
						Left: nil,
						Right: &leetcode.TreeNode{
							Val:  3,
							Left: nil,
							Right: &leetcode.TreeNode{
								Val:  4,
								Left: nil,
								Right: &leetcode.TreeNode{
									Val:   5,
									Left:  nil,
									Right: nil,
								},
							},
						},
					},
					Right: nil,
				},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDepth(tt.args.root); got != tt.want {
				t.Errorf("minDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
