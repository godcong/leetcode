package leetcode

import (
	"fmt"
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
	if root == nil {
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
	fmt.Println(sNums)
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
		if nodes[i] == nil && i+1 < len(sNums) {
			strToTreeNodeDFS(nodes[i+1], i*2+1, nodes)
			continue
		}
		strToTreeNodeDFS(nodes[i], i*2+1, nodes)
	}
	return nodes[0]
}

func strToTreeNodeDFS(node *TreeNode, idx int, nodes []*TreeNode) {
	if node == nil {
		return
	}
	if idx < len(nodes) {
		node.Left = nodes[idx]
		idx++
	}
	if idx < len(nodes) {
		node.Right = nodes[idx]
		idx++
	}
}

func throughErrorPanic(err error) {
	if err != nil {
		panic(err)
	}
}
