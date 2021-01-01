package code

import "github.com/godcong/leetcode"

/*
111. 二叉树的最小深度
给定一个二叉树，找出其最小深度。

最小深度是从根节点到最近叶子节点的最短路径上的节点数量。

说明: 叶子节点是指没有子节点的节点。

示例:

给定二叉树 [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回它的最小深度  2.
*/
func minDepth(root *leetcode.TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	lval, rval := 0, 0
	if root.Left != nil {
		lval = minDepth(root.Left) + 1
	}
	if root.Right != nil {
		rval = minDepth(root.Right) + 1
	}
	if lval == 0 {
		return rval
	}
	if rval == 0 {
		return lval
	}
	if lval < rval {
		return lval
	}

	return rval
}
