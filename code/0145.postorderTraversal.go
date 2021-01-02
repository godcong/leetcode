package code

/*
145. 二叉树的后序遍历
给定一个二叉树，返回它的 后序 遍历。

示例:

输入: [1,null,2,3]
   1
    \
     2
    /
   3

输出: [3,2,1]
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
func postorderTraversal(root *TreeNode) []int {
	var ret []int
	var postorderTraversalDFS func(*TreeNode)
	postorderTraversalDFS = func(node *TreeNode) {
		if node == nil {
			return
		}
		postorderTraversalDFS(node.Left)
		postorderTraversalDFS(node.Right)
		ret = append(ret, node.Val)
	}
	postorderTraversalDFS(root)
	return ret
}
