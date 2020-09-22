package leetcode

import (
	"strconv"
	"strings"
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

func strToTreeNode(nums string) *TreeNode {
	sta := 0
	end := len(nums)
	if strings.Index(nums, "[") == 0 {
		sta += 1
	}

	if strings.Index(nums, "]") == end-1 {
		end -= 1
	}

	sNums := strings.Split(nums[sta:end], ",")
	nodes := make([]*TreeNode, len(sNums))
	for idx, num := range sNums {
		if num == strings.ToLower("null") || num == "" {
			continue
		}
		i, err := strconv.Atoi(num)
		throughErrorPanic(err)
		nodes[idx] = &TreeNode{Val: i}
	}
	//printTreeNodeToArray(nodes)

	idx := 0
	next := 0
	for i := range sNums {
		if nodes[i] == nil {
			continue
		}
		idx = i*2 + 1 - next
		if idx < len(sNums) {
			if nodes[idx] == nil {
				next += 2
			}
			nodes[i].Left = nodes[idx]
		}
		idx++
		if idx < len(sNums) {
			if nodes[idx] == nil && nodes[idx-1] == nil {
				next -= 2
			}
			nodes[i].Right = nodes[idx]
		}
	}

	return nodes[0]
}

func throughErrorPanic(err error) {
	if err != nil {
		panic(err)
	}
}
