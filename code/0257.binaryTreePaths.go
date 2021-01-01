package code

import (
	"strconv"

	"github.com/godcong/leetcode"
)

/*
257. 二叉树的所有路径
给定一个二叉树，返回所有从根节点到叶子节点的路径。

说明: 叶子节点是指没有子节点的节点。

示例:

输入:

   1
 /   \
2     3
 \
  5

输出: ["1->2->5", "1->3"]

解释: 所有根节点到叶子节点的路径为: 1->2->5, 1->3
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *leetcode.TreeNode) []string {
	var res []string

	var binaryTreePathsDeep func(root *leetcode.TreeNode, path string)
	binaryTreePathsDeep = func(root *leetcode.TreeNode, path string) {
		if root != nil {
			pathSB := path + strconv.Itoa(root.Val)
			if root.Left == nil && root.Right == nil {
				res = append(res, pathSB)
			} else {
				pathSB += "->"
				binaryTreePathsDeep(root.Left, pathSB)
				binaryTreePathsDeep(root.Right, pathSB)
			}
		}
	}
	binaryTreePathsDeep(root, "")
	return res
}
