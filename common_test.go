package leetcode

import (
	"testing"
)

func Test_strToTreeNode(t *testing.T) {
	type args struct {
		nums string
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "",
			args: args{
				nums: "[0,0,null,0,0]",
			},
			want: &TreeNode{
				Val: 0,
				Left: &TreeNode{
					Val: 0,
					Left: &TreeNode{
						Val:   0,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:   0,
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
				nums: "0,0,null,0,0]",
			},
			want: &TreeNode{
				Val: 0,
				Left: &TreeNode{
					Val: 0,
					Left: &TreeNode{
						Val:   0,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:   0,
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
				nums: "0,0,null,0,0",
			},
			want: &TreeNode{
				Val: 0,
				Left: &TreeNode{
					Val: 0,
					Left: &TreeNode{
						Val:   0,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:   0,
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
				nums: "0,0,null,0,0,",
			},
			want: &TreeNode{
				Val: 0,
				Left: &TreeNode{
					Val: 0,
					Left: &TreeNode{
						Val:   0,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:   0,
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
				nums: "0,0,null,0,0,",
			},
			want: &TreeNode{
				Val: 0,
				Left: &TreeNode{
					Val: 0,
					Left: &TreeNode{
						Val:   0,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:   0,
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
				nums: "[1,null,2,3]",
			},
			want: &TreeNode{
				Val:  1,
				Left: nil,
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val:   3,
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
				nums: "[1,null,2,null,3]",
			},
			want: &TreeNode{
				Val:  1,
				Left: nil,
				Right: &TreeNode{
					Val:  2,
					Left: nil,
					Right: &TreeNode{
						Val:   3,
						Left:  nil,
						Right: nil,
					},
				},
			},
		},
		{
			name: "",
			args: args{
				nums: "[1,null,2,null,3,null,4,null,5]",
			},
			want: &TreeNode{
				Val:  1,
				Left: nil,
				Right: &TreeNode{
					Val:  2,
					Left: nil,
					Right: &TreeNode{
						Val:  3,
						Left: nil,
						Right: &TreeNode{
							Val:  4,
							Left: nil,
							Right: &TreeNode{
								Val:   5,
								Left:  nil,
								Right: nil,
							},
						},
					},
				},
			},
		},
		{
			name: "",
			args: args{
				nums: "[3,1,4,null,null,2]",
			},
			want: &TreeNode{
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strToTreeNode(tt.args.nums); !treeNodeDeepEqual(t, got, tt.want) {
				t.Errorf("strToTreeNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
