package leetcode

import "testing"

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
