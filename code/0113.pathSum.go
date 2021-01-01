package code

import "github.com/godcong/leetcode"

/*
113. 路径总和 II
给定一个二叉树和一个目标和，找到所有从根节点到叶子节点路径总和等于给定目标和的路径。

说明: 叶子节点是指没有子节点的节点。

示例:
给定如下二叉树，以及目标和 sum = 22，

              5
             / \
            4   8
           /   / \
          11  13  4
         /  \    / \
        7    2  5   1
返回:

[
   [5,4,11,2],
   [5,8,4,5]
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
func pathSum(root *leetcode.TreeNode, sum int) [][]int {
	var ret [][]int
	var pathSumDFS func(node *leetcode.TreeNode, t int, v []int)
	pathSumDFS = func(node *leetcode.TreeNode, t int, v []int) {
		if node == nil {
			return
		}

		t -= node.Val
		v = append(v, node.Val)
		if node.Right == nil && node.Left == nil && t == 0 {
			ret = append(ret, append([]int{}, v...))
			return
		}
		pathSumDFS(node.Left, t, v)
		pathSumDFS(node.Right, t, v)
	}
	pathSumDFS(root, sum, nil)

	return ret
}
