package _0304

import (
	"fmt"
	"strconv"
	"testing"
)

func Test_strToListNode(t *testing.T) {
	type args struct {
		nums string
		pos  int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
		deep int
	}{
		{
			name: "",
			args: args{
				nums: "[3,2,0,-4]",
				pos:  1,
			},
			want: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 0,
						Next: &ListNode{
							Val: -4,
							Next: &ListNode{
								Val:  2,
								Next: nil,
							},
						},
					},
				},
			},
			deep: 5,
		},
		{
			name: "",
			args: args{
				nums: "[1,2]",
				pos:  0,
			},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val:  1,
						Next: nil,
					},
				},
			},
			deep: 3,
		},
		{
			name: "",
			args: args{
				nums: "[1]",
				pos:  -1,
			},
			want: &ListNode{
				Val:  1,
				Next: nil,
			},
			deep: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strToListNode(tt.args.nums, tt.args.pos); !listNodeDeepEqual(t, got, tt.want, tt.deep) {
				t.Errorf("strToListNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

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

func listNodeDeepEqual(t *testing.T, root *ListNode, want *ListNode, deep int) bool {
	checked := make(map[*ListNode]bool)
	left, right := root, want
	deepCount := 1
	for deepCount <= deep {
		if left == nil || right == nil {
			return left == right
		}

		if checked[left] {
			return true
		}

		if left.Val != right.Val {
			t.Errorf("count:%+v,got:%+v,want:%+v\n", deepCount, left.Val, right.Val)
			return false
		}

		checked[left] = true
		left = root.Next
		right = want.Next
		deepCount++
	}

	return true
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
