package leetcode

import (
	"strconv"
	"strings"
	"testing"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */
type Node struct {
	Val       int
	Neighbors []*Node
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func treeNodeDeepEqual(t *testing.T, root *TreeNode, want *TreeNode) bool {
	if root == nil {
		return root == want
	}

	if root.Left != nil {
		treeNodeDeepEqual(t, root.Left, want.Left)
	}
	if root.Right != nil {
		treeNodeDeepEqual(t, root.Right, want.Right)
	}
	if root.Val != want.Val {
		t.Errorf("recoverTree() = %v, want %v", root.Val, want.Val)
	}
	return true
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
	for i := range sNums {
		strToTreeNodeDFS(nodes[i], i*2+1, nodes)
	}
	return nodes[0]
}

func strToTreeNodeDFS(node *TreeNode, idx int, nodes []*TreeNode) {
	if node == nil || idx > len(nodes)-1 {
		return
	}
	node.Left = nodes[idx]
	node.Right = nodes[idx+1]
}

func throughErrorPanic(err error) {
	if err != nil {
		panic(err)
	}
}
