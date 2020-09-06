package leetcode

/*
107. 二叉树的层次遍历 II
给定一个二叉树，返回其节点值自底向上的层次遍历。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）

例如：
给定二叉树 [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其自底向上的层次遍历为：

[
  [15,7],
  [9,20],
  [3]
]
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
	var ret [][]int
	if root == nil {
		return ret
	}
	var levelOrderBottomAppendSub func(node *TreeNode, index int)
	levelOrderBottomAppendSub = func(node *TreeNode, index int) {
		if len(ret) <= index {
			ret = append(ret, make([]int, 0))
		}
		if node.Left != nil {
			levelOrderBottomAppendSub(node.Left, index+1)
		}
		if node.Right != nil {
			levelOrderBottomAppendSub(node.Right, index+1)
		}
		ret[index] = append(ret[index], node.Val)
	}
	levelOrderBottomAppendSub(root, 0)

	for i, j := 0, len(ret)-1; i < j; i, j = i+1, j-1 {
		ret[i], ret[j] = ret[j], ret[i]
	}
	return ret
}
