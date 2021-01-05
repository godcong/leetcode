package code

/*
144. 二叉树的前序遍历
给定一个二叉树，返回它的 前序 遍历。

 示例:

输入: [1,null,2,3]
   1
    \
     2
    /
   3

输出: [1,2,3]
进阶: 递归算法很简单，你可以通过迭代算法完成吗？
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
	var vals []int
	var preorderTraversalDFS func(*TreeNode)
	preorderTraversalDFS = func(node *TreeNode) {
		if node == nil {
			return
		}
		vals = append(vals, node.Val)
		preorderTraversalDFS(node.Left)
		preorderTraversalDFS(node.Right)
	}
	preorderTraversalDFS(root)
	return vals
}
