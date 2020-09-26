package leetcode

import (
	"fmt"
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
		{
			name: "",
			args: args{
				nums: "[5,4,8,11,null,13,4,7,2,null,null,5,1]",
			},
			want: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 11,
						Left: &TreeNode{
							Val:   7,
							Left:  nil,
							Right: nil,
						},
						Right: &TreeNode{
							Val:   2,
							Left:  nil,
							Right: nil,
						},
					},
					Right: nil,
				},
				Right: &TreeNode{
					Val: 8,
					Left: &TreeNode{
						Val:   13,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val:   5,
							Left:  nil,
							Right: nil,
						},
						Right: &TreeNode{
							Val:   1,
							Left:  nil,
							Right: nil,
						},
					},
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
	size := len(sNums)
	nodes := make([]*TreeNode, size)
	for idx, num := range sNums {
		if num == strings.ToLower("null") || num == "" {
			continue
		}

		i, err := strconv.Atoi(num)
		throughErrorPanic(err)
		nodes[idx] = &TreeNode{Val: i}
	}

	i := 1
	queue := make([]*TreeNode, 1, size*2)
	queue[0] = nodes[0]
	for i < len(sNums) {
		node := queue[0]
		queue = queue[1:]

		if i < len(sNums) && nodes[i] != nil {
			node.Left = nodes[i]
			queue = append(queue, node.Left)
		}
		i++

		if i < len(sNums) && nodes[i] != nil {
			node.Right = nodes[i]
			queue = append(queue, node.Right)
		}
		i++
	}
	return nodes[0]
}

func throughErrorPanic(err error) {
	if err != nil {
		panic(err)
	}
}

func treeNodeDeepPrint(root *TreeNode) {
	fmt.Println(getTreeNodeIntArray(root))
}

func getTreeNodeIntArray(root *TreeNode) []string {
	var ret []string
	ret = append(ret, strconv.Itoa(root.Val))
	if root.Left != nil {
		ret = append(ret, getTreeNodeIntArray(root.Left)...)
	} else {
		ret = append(ret, "null")
	}
	if root.Right != nil {
		ret = append(ret, getTreeNodeIntArray(root.Right)...)
	} else {
		//ret = append(ret, "null")
	}
	return ret
}

func treeNodeDeepEqual(t *testing.T, root *TreeNode, want *TreeNode) bool {
	if root == nil || want == nil {
		return root == want
	}

	if !treeNodeDeepEqual(t, root.Left, want.Left) {
		t.Errorf("recoverTree(Left) = %v, want %v", root.Left, want.Left)
		return false
	}
	if !treeNodeDeepEqual(t, root.Right, want.Right) {
		t.Errorf("recoverTree(Right) = %v, want %v", root.Right, want.Right)
		return false
	}
	if root.Val != want.Val {
		t.Errorf("recoverTree() = %v, want %v", root.Val, want.Val)
	}
	return true
}

func printTreeNodeToArray(nodes []*TreeNode) {
	for i := range nodes {
		if nodes[i] == nil {
			fmt.Print("null")
		} else {
			fmt.Print(nodes[i].Val)
		}
		fmt.Print(" ")
	}
	fmt.Println()
}
